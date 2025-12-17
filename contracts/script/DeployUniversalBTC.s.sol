// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Script} from "forge-std/Script.sol";
import {ZetaDCAExecutionBTC} from "../src/ZetaDCAExecutionBTC.sol";  // 改成你的BTC合约路径

contract DeployUniversalBTC is Script {
    function run() external {
        vm.startBroadcast();
        address dex = Ox2ca7dc5c9b07e5c3e3d3f5c9c9e5c3e3d3f5c9c;            // 同SOL（测试网DEX通常通用）
        address usdc = 0xYOUR_USDC_ZRC20;            // 只用一个稳定币地址测试
        address btc = 0xYOUR_BTC_ZRC20;              // BTC ZRC20测试网地址

        ZetaDCAExecutionBTC universalBTC = new ZetaDCAExecutionBTC(gateway, dex, usdc, address(0), btc);  // USDT填0或同一个USDC

        vm.stopBroadcast();
        console.log("Universal BTC Contract on ZetaChain:", address(universalBTC));
    }
}