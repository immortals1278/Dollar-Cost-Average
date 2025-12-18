// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "forge-std/console.sol";
import {Script} from "forge-std/Script.sol";
import {ZetaDCAExecutionBTC} from "../src/ZetaDCASwapBTC.sol";  // 改成你的BTC合约路径

contract DeployUniversalBTC is Script {
    function run() external {
        vm.startBroadcast();
        address dex = 0x2ca7d64A7EFE2D62A725E2B35Cf7230D6677FfEe;            // 同SOL（测试网DEX通常通用）
        address usdc = 0xd0eFed75622e7AA4555EE44F296dA3744E3ceE19;
        address btc = 0xfC9201f4116aE6b054722E10b98D904829b469c3;              // BTC ZRC20测试网地址

        ZetaDCAExecutionBTC universalBTC = new ZetaDCAExecutionBTC(dex, usdc, btc);  // USDT填0或同一个USDC

        vm.stopBroadcast();
        console.log("Universal BTC Contract on ZetaChain:", address(universalBTC));
    }
}