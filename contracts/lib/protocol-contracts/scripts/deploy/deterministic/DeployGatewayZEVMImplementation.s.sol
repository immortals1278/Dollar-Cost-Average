// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/zevm/GatewayZEVM.sol";

contract DeployGatewayZEVM is Script {
    function run() external {
        address expectedImplAddress;
        bytes32 implSalt = keccak256("GatewayZEVM");

        vm.startBroadcast();

        // Compute expected implementation address
        expectedImplAddress = vm.computeCreate2Address(
            implSalt,
            hashInitCode(type(GatewayZEVM).creationCode)
        );

        // Deploy the implementation contract using CREATE2
        GatewayZEVM gatewayImpl = new GatewayZEVM{salt: implSalt}();
        require(address(gatewayImpl) != address(0), "gatewayImpl deployment failed");
        require(expectedImplAddress == address(gatewayImpl), "impl address doesn't match expected address");

        vm.stopBroadcast();
    }
}