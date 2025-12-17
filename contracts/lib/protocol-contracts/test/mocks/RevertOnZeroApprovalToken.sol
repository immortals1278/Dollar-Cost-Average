// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../utils/TestERC20.sol";

contract RevertOnZeroApprovalToken is TestERC20 {
    constructor(string memory name, string memory symbol) TestERC20(name, symbol) { }

    function approve(address spender, uint256 amount) public override returns (bool) {
        if (amount <= 0) revert();
        return super.approve(spender, amount);
    }
}
