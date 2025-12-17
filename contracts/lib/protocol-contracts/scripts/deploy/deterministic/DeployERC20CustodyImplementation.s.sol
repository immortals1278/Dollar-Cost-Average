// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/evm/ERC20Custody.sol";

contract DeployERC20CustodyImplementation is Script {
    function run() external {
        address expectedImplAddress;
        bytes32 implSalt = keccak256("ERC20Custody");

        vm.startBroadcast();

        // Compute expected implementation address
        expectedImplAddress = vm.computeCreate2Address(
            implSalt,
            hashInitCode(type(ERC20Custody).creationCode)
        );

        // Deploy the implementation contract using CREATE2
        ERC20Custody erc20CustodyImpl = new ERC20Custody{salt: implSalt}();
        require(address(erc20CustodyImpl) != address(0), "erc20CustodyImpl deployment failed");
        require(expectedImplAddress == address(erc20CustodyImpl), "impl address doesn't match expected address");

        vm.stopBroadcast();
    }
}