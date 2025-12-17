// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/// @title TestERC20
/// @notice This contract is used just for testing.
/// @dev A simple ERC20 token contract with a mint function.
contract TestERC20 is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) { }

    /// @notice Mints new tokens to the specified address.
    /// @param to The address to which the tokens will be minted.
    /// @param amount The amount of tokens to mint.
    /// @dev This function can be called by anyone for testing purposes.
    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}
