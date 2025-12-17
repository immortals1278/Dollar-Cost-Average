// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Script} from "forge-std/Script.sol";
import {DCAController} from "../src/DCAController.sol";

contract DeployController is Script {
    function run() external {
        vm.startBroadcast();

        address gatewayOnBase = 0x0c487a766110c85d301d96e33579c5b317fa4995;  // 确认地址
        address universalApp = 0xYOUR_UNIVERSAL_FROM_PREVIOUS;
        uint256 zetaChainId = 7001;

        DCAController controller = new DCAController(gatewayOnBase, universalApp, zetaChainId);

        vm.stopBroadcast();
        console.log("DCAController on Base Sepolia:", address(controller));
    }
}