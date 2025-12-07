// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// 这是 Zeta 官方 swap 接口的最小化抽象版，你可以直接用
interface IZetaSwap {
    function swap(
        address inputToken,
        uint256 amountIn,
        address outputZRC20,
        address recipient
    ) external returns (uint256 amountOut);
}
