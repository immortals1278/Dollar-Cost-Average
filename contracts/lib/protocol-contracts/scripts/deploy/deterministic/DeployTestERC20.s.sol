// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "test/utils/TestERC20.sol";

// This is just to deploy test erc20 tokens for testing deployments
contract DeployTestERC20 is Script {
    function run() external {
        bytes32 erc20Salt = keccak256("TestERC20");

        vm.startBroadcast();

        TestERC20 zeta = new TestERC20{salt: erc20Salt}("zeta", "ZETA");
        require(address(zeta) != address(0), "deployment failed");

        address expectedAddr = vm.computeCreate2Address(
            erc20Salt,
            hashInitCode(type(TestERC20).creationCode, abi.encode("zeta", "ZETA"))
        );

        require(expectedAddr == address(zeta), "zeta address doesn't match expected address");

        vm.stopBroadcast();
    }
}