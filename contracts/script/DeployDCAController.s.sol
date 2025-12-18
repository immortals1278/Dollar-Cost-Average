// script/DeployDCAController.s.sol
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import { DCAController } from "../src/DCAController.sol";

contract DeployDCAController is Script {
    function run() external {
        vm.startBroadcast();

        address gateway = 0x0c487a766110c85d301D96E33579C5B317Fa4995;
        address universalApp = 0x35aA5A3b0Ef7739D953A72ca4480dB88b52DEd70;     // USDC 版本（或你主要用的那个）
        address universalAppBTC = 0x2359a48a8F6253e50ab657eF977E3068370244C2;    // BTC 版本
        uint256 zetaChainId = 7001;

        DCAController controller = new DCAController(gateway, universalApp, universalAppBTC, zetaChainId);

        console.log("DCAController deployed on Base Sepolia:");
        console.logAddress(address(controller));

        vm.stopBroadcast();
    }
}