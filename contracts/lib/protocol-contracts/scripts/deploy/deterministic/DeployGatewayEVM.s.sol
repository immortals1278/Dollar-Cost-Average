// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/evm/GatewayEVM.sol";
import "test/utils/TestERC20.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployGatewayEVM is Script {
    function run() external {
        address payable tss = payable(vm.envAddress("TSS_ADDRESS"));
        address admin = vm.envAddress("GATEWAY_ADMIN_ADDRESS_EVM");
        address zeta = vm.envAddress("ZETA_ERC20_EVM");

        address expectedImplAddress;
        address expectedProxyAddress;

        bytes32 implSalt = keccak256("GatewayEVM");
        bytes32 proxySalt = keccak256("GatewayEVMProxy");

        // Add this specific check to ensure contract is not deployed at deterministic address with wrong admin
        require(admin != address(0), "Environment variable GATEWAY_ADMIN_ADDRESS_EVM is not set");

        vm.startBroadcast();

        // Compute expected implementation address
        expectedImplAddress = vm.computeCreate2Address(
            implSalt,
            hashInitCode(type(GatewayEVM).creationCode)
        );

        // Deploy the implementation contract using CREATE2
        GatewayEVM gatewayImpl = new GatewayEVM{salt: implSalt}();
        require(address(gatewayImpl) != address(0), "gatewayImpl deployment failed");
        require(expectedImplAddress == address(gatewayImpl), "impl address doesn't match expected address");

        // Compute expected proxy address
        expectedProxyAddress = vm.computeCreate2Address(
            proxySalt,
            hashInitCode(
                type(ERC1967Proxy).creationCode,
                abi.encode(
                    address(gatewayImpl),
                    abi.encodeWithSelector(GatewayEVM.initialize.selector, tss, address(zeta), msg.sender)
                )
            )
        );

        // Deploy the proxy contract using CREATE2, with deployer as initial admin
        ERC1967Proxy gatewayProxy = new ERC1967Proxy{salt: proxySalt}(
            address(gatewayImpl),
            abi.encodeWithSelector(GatewayEVM.initialize.selector, tss, address(zeta), msg.sender)
        );
        require(address(gatewayProxy) != address(0), "gatewayProxy deployment failed");
        require(expectedProxyAddress == address(gatewayProxy), "proxy address doesn't match expected address");

        GatewayEVM gateway = GatewayEVM(address(gatewayProxy));

        // Verify initial configuration
        require(gateway.tssAddress() == tss, "tss not set");
        require(gateway.zetaToken() == address(zeta), "zeta token not set");

        // Transfer admin role from deployer to admin if the address is different
        if (msg.sender != admin) {
            transferAdmin(gateway, msg.sender, admin);
        }
        vm.stopBroadcast();
    }

    function transferAdmin(GatewayEVM gateway, address deployer, address newAdmin) internal {
        // Grant roles to specified admin and renounce deployer's roles
        gateway.grantRole(gateway.PAUSER_ROLE(), newAdmin);
        gateway.grantRole(gateway.DEFAULT_ADMIN_ROLE(), newAdmin);

        // Renounce deployer's roles
        gateway.renounceRole(gateway.PAUSER_ROLE(), deployer);
        gateway.renounceRole(gateway.DEFAULT_ADMIN_ROLE(), deployer);

        // Assert that the roles are correctly transitioned
        require(gateway.hasRole(gateway.PAUSER_ROLE(), newAdmin), "admin does not have PAUSER_ROLE");
        require(gateway.hasRole(gateway.DEFAULT_ADMIN_ROLE(), newAdmin), "admin does not have DEFAULT_ADMIN_ROLE");
        require(!gateway.hasRole(gateway.PAUSER_ROLE(), deployer), "deployer still has PAUSER_ROLE");
        require(!gateway.hasRole(gateway.DEFAULT_ADMIN_ROLE(), deployer), "deployer still has DEFAULT_ADMIN_ROLE");
    }
}
