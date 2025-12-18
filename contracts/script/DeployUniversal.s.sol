// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "forge-std/console.sol";
import {Script} from "forge-std/Script.sol";
import {ZetaDCAExecution} from "../src/ZetaDCASwap.sol";

contract DeployUniversal is Script {
    function run() external {
        vm.startBroadcast();
        address dex = 0x2ca7d64A7EFE2D62A725E2B35Cf7230D6677FfEe;            // Uniswap v2 router on testnet
        address usdc = 0xd0eFed75622e7AA4555EE44F296dA3744E3ceE19;
        address sol = 0xADF73ebA3Ebaa7254E859549A44c74eF7cff7501;

        ZetaDCAExecution universal = new ZetaDCAExecution(dex, usdc, sol);

        vm.stopBroadcast();
        console.log("Universal Contract on ZetaChain:", address(universal));
    }
}