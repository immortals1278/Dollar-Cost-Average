// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/evm/GatewayEVM.sol";
import "test/utils/GatewayEVMUpgradeTest.sol";
import "test/utils/TestERC20.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract UpgradeGatewayEVM is Script {
    function run() external {
        address proxy = vm.envAddress("GATEWAY_EVM_PROXY");

        GatewayEVM prevImpl = GatewayEVM(proxy);

        vm.startBroadcast(deployer);

        Upgrades.upgradeProxy(proxy, "GatewayEVMUpgradeTest.sol", "");

        GatewayEVM newImpl = GatewayEVM(proxy);

        require(prevImpl.tssAddress() == newImpl.tssAddress(), "tss addr changed");
        require(prevImpl.zetaToken() == newImpl.zetaToken(), "zeta token addr changed");

        vm.stopBroadcast();
    }
}