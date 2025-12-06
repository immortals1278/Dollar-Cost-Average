// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@zetachain/protocol-contracts/contracts/interfaces/IMessageRecipient.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IDEX {
    function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);
}

/**
 * @title TargetChainSwapper
 * @notice 收到跨链消息后执行 swap
 */
contract TargetChainSwapper is IMessageRecipient {

    address public dex;               // 如 PancakeSwap Router
    address public baseToken;         // 比如 WETH、WMATIC
    address public zrc20;             // 对应 Zeta 支付 token

    constructor(
        address _dex,
        address _baseToken,
        address _zrc20
    ) {
        dex = _dex;
        baseToken = _baseToken;
        zrc20 = _zrc20;
    }

    /**
     * 跨链消息入口：ZetaChain 会调用这个
     */
    function onZetaMessage(
        bytes calldata zetaMessage
    ) external override {

        (address user, uint256 amountIn, bytes memory targetTokenBytes)
            = abi.decode(zetaMessage, (address, uint256, bytes));

        address targetToken = abi.decode(targetTokenBytes, (address));

        // ------------------------
        //   1. 将 USDT 换成目标资产
        // ------------------------
        IERC20(zrc20).approve(dex, amountIn);

        address;
        path[0] = zrc20;
        path[1] = targetToken;

        uint[] memory result = IDEX(dex).swapExactTokensForTokens(
            amountIn,
            0,                  // no slippage protection in MVP
            path,
            user,               // 直接发给用户
            block.timestamp + 300
        );

        // 结果 token 已经发给用户
    }
}
