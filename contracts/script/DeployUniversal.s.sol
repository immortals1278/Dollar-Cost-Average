// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Script} from "forge-std/Script.sol";
import {ZetaDCAExecution} from "../src/ZetaDCAExecution.sol";

contract DeployUniversal is Script {
    function run() external {
        vm.startBroadcast();

        address dex = Ox2ca7dc5c9b07e5c3e3d3f5c9c9e5c3e3d3f5c9c;            // Uniswap v2 router on testnet
        address usdc = 0xYOUR_USDC_ZRC20;  
        address sol = 0xYOUR_SOL_ZRC20;

        ZetaDCAExecution universal = new ZetaDCAExecution(gateway, dex, usdt, usdc, sol);

        vm.stopBroadcast();
        console.log("Universal Contract on ZetaChain:", address(universal));
    }
}