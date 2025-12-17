// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/evm/ERC20Custody.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployERC20Custody is Script {
    function run() external {
        address payable tss = payable(vm.envAddress("TSS_ADDRESS"));
        address admin = vm.envAddress("ERC20_CUSTODY_ADMIN_ADDRESS_EVM");
        address gateway = vm.envAddress("GATEWAY_PROXY_EVM");

        address expectedImplAddress;
        address expectedProxyAddress;

        bytes32 implSalt = keccak256("ERC20Custody");
        bytes32 proxySalt = keccak256("ERC20CustodyProxy");

        // Add this specific check to ensure contract is not deployed at deterministic address with wrong admin
        require(admin != address(0), "Environment variable ERC20_CUSTODY_ADMIN_ADDRESS_EVM is not set");

        vm.startBroadcast();

        // Compute expected implementation address
        expectedImplAddress = vm.computeCreate2Address(
            implSalt,
            hashInitCode(type(ERC20Custody).creationCode)
        );

        // Deploy the implementation contract using CREATE2
        ERC20Custody erc20CustodyImpl = new ERC20Custody{salt: implSalt}();
        require(address(erc20CustodyImpl) != address(0), "erc20CustodyImpl deployment failed");
        require(expectedImplAddress == address(erc20CustodyImpl), "impl address doesn't match expected address");

        // Compute expected proxy address
        expectedProxyAddress = vm.computeCreate2Address(
            proxySalt,
            hashInitCode(
                type(ERC1967Proxy).creationCode,
                abi.encode(
                    address(erc20CustodyImpl),
                    abi.encodeWithSelector(ERC20Custody.initialize.selector, gateway, tss, msg.sender)
                )
            )
        );

        // Deploy the proxy contract using CREATE2, with deployer as initial admin
        ERC1967Proxy erc20CustodyProxy = new ERC1967Proxy{salt: proxySalt}(
            address(erc20CustodyImpl),
            abi.encodeWithSelector(ERC20Custody.initialize.selector, gateway, tss, msg.sender)
        );
        require(address(erc20CustodyProxy) != address(0), "erc20CustodyProxy deployment failed");
        require(expectedProxyAddress == address(erc20CustodyProxy), "proxy address doesn't match expected address");

        ERC20Custody erc20Custody = ERC20Custody(address(erc20CustodyProxy));

        // Verify initial configuration
        require(erc20Custody.tssAddress() == tss, "tss not set");
        require(address(erc20Custody.gateway()) == gateway, "gateway not set");

        // Transfer admin role from deployer to admin
        if (msg.sender != admin) {
            transferAdmin(erc20Custody, msg.sender, admin);
        }

        vm.stopBroadcast();
    }

    function transferAdmin(ERC20Custody erc20Custody, address deployer, address newAdmin) internal {
        // Grant roles to specified admin and renounce deployer's roles
        erc20Custody.grantRole(erc20Custody.PAUSER_ROLE(), newAdmin);
        erc20Custody.grantRole(erc20Custody.WHITELISTER_ROLE(), newAdmin);
        erc20Custody.grantRole(erc20Custody.DEFAULT_ADMIN_ROLE(), newAdmin);

        // Renounce deployer's roles
        erc20Custody.renounceRole(erc20Custody.PAUSER_ROLE(), deployer);
        erc20Custody.renounceRole(erc20Custody.WHITELISTER_ROLE(), deployer);
        erc20Custody.renounceRole(erc20Custody.DEFAULT_ADMIN_ROLE(), deployer);

        // Assert that the roles are correctly transitioned
        require(erc20Custody.hasRole(erc20Custody.PAUSER_ROLE(), newAdmin), "admin does not have PAUSER_ROLE");
        require(erc20Custody.hasRole(erc20Custody.WHITELISTER_ROLE(), newAdmin), "admin does not have WHITELISTER_ROLE");
        require(erc20Custody.hasRole(erc20Custody.DEFAULT_ADMIN_ROLE(), newAdmin), "admin does not have DEFAULT_ADMIN_ROLE");
        require(!erc20Custody.hasRole(erc20Custody.PAUSER_ROLE(), deployer), "deployer still has PAUSER_ROLE");
        require(!erc20Custody.hasRole(erc20Custody.WHITELISTER_ROLE(), deployer), "deployer still has WHITELISTER_ROLE");
        require(!erc20Custody.hasRole(erc20Custody.DEFAULT_ADMIN_ROLE(), deployer), "deployer still has DEFAULT_ADMIN_ROLE");
    }
}
