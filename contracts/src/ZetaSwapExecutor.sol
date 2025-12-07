// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./interfaces/IZetaSwap.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ZetaSwapExecutor {

    address public immutable zetaSwapAddress;

    constructor(address _zetaSwapAddress) {
        zetaSwapAddress = _zetaSwapAddress;
    }

    /// 实际执行 swap（调用 ZetaChain 系统合约）
    function executeSwap(
        address depositToken,
        uint256 amount,
        address targetZRC20,
        address user
    ) external returns (uint256 amountOut) {

        IERC20(depositToken).transferFrom(msg.sender, address(this), amount);

        IERC20(depositToken).approve(zetaSwapAddress, amount);

        amountOut = IZetaSwap(zetaSwapAddress).swap(
            depositToken,
            amount,
            targetZRC20,
            user   // 用户直接收到对应 ZRC20
        );
    }
}
