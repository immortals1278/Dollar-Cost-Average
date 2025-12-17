// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/evm/ZetaConnectorNonNative.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployZetaConnectorNonNative is Script {
    function run() external {
        address payable tss = payable(vm.envAddress("TSS_ADDRESS"));
        address admin = vm.envAddress("ZETA_CONNECTOR_ADMIN_ADDRESS_EVM");
        address gateway = vm.envAddress("GATEWAY_PROXY_EVM");
        address zeta = vm.envAddress("ZETA_ERC20_EVM");

        bytes32 implSalt = keccak256("ZetaConnectorNonNative");
        bytes32 proxySalt = keccak256("ZetaConnectorNonNativeProxy");

        vm.startBroadcast();

        ZetaConnectorNonNative connectorImpl = new ZetaConnectorNonNative{salt: implSalt}();
        require(address(connectorImpl) != address(0), "connectorImpl deployment failed");
        require(
            vm.computeCreate2Address(implSalt, hashInitCode(type(ZetaConnectorNonNative).creationCode))
            == address(connectorImpl),
            "impl address doesn't match expected address"
        );

        bytes memory initData = abi.encodeWithSelector(
            ZetaConnectorNonNative.initialize.selector,
            gateway,
            zeta,
            tss,
            admin
        );

        ERC1967Proxy connectorProxy = new ERC1967Proxy{salt: proxySalt}(address(connectorImpl), initData);
        require(address(connectorProxy) != address(0), "connectorProxy deployment failed");
        require(
            vm.computeCreate2Address(
                proxySalt,
                hashInitCode(type(ERC1967Proxy).creationCode, abi.encode(address(connectorImpl), initData))
            ) == address(connectorProxy),
            "proxy address doesn't match expected address"
        );

        // Verify deployment
        ZetaConnectorNonNative connector = ZetaConnectorNonNative(address(connectorProxy));
        require(connector.tssAddress() == tss, "tss not set");
        require(address(connector.gateway()) == gateway, "gateway not set");
        require(address(connector.zetaToken()) == zeta, "zeta not set");

        vm.stopBroadcast();
    }
}