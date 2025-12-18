// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

// ============= ZetaChain 官方接口 =============
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/UniversalContract.sol";
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/IZRC20.sol";


// ============= UniswapV2 风格 DEX (Zeta DEX) =============
interface IUniswapV2Router {
    function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    function getAmountsIn(
        uint amountOut, 
        address[] calldata path
    ) external view returns (uint[] memory amounts);
}

contract ZetaDCAExecutionBTC is UniversalContract {
    IUniswapV2Router public immutable dex;
    
    mapping(bytes => uint256) public userBalances; 

    /// @notice ZRC20 地址
    address public immutable USDC_ZRC20;
    address public immutable BTC_ZRC20;
    

    event BoughtAndStored(
        bytes indexed userId,       // 源链用户自定义 ID
        address indexed tokenIn,    // 输入的 ZRC20
        uint256 amountIn,
        uint256 btcOut,
        uint256 totalStored
    );
    event UserBalance(
        bytes indexed userId,
        uint256 balance
    );
    event WithdrawExecuted(
        bytes indexed userId,
        uint256 btcRequired,
        uint256 usdtOut,
        address indexed recipientEVM
    );

    constructor(
        address _dex,
        address _usdc,
        address _btc
    ) {
        dex = IUniswapV2Router(_dex);
        USDC_ZRC20 = _usdc;
        BTC_ZRC20  = _btc;
    }

    struct SwapMsg {
        bytes userId;       // 源链传过来的用户标识
        address tokenIn;    // 输入 token (USDT/USDC 的 ZRC20)
    }


    /// @notice ZetaChain 跨链调用入口
    function onCall(
        MessageContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    ) external override onlyGateway {
        (uint8 opType, bytes memory params) = abi.decode(message, (uint8, bytes));
        if (opType == 1) {
            // Swap
            (bytes memory userId, address tokenIn) = abi.decode(params, (bytes, address));
            _executeSwap(userId, tokenIn, amount);

        } else if (opType == 2){
            // Query
            (bytes memory userId) = abi.decode(params, (bytes));
            getUserBalance(userId);
        } else if (opType == 3) {
            // Withdraw
            (bytes memory userId, uint256 wantUSDT, bytes memory recipient) =
                abi.decode(params, (bytes, uint256, bytes));

            _executeWithdraw(userId, wantUSDT, recipient);
        }
    }

    function onCall(
        MessageContext calldata context,
        bytes calldata message
    ) external payable override onlyGateway {
        revert("Native ZETA transfers not supported");
    }

    /// @notice 查询用户在 ZetaChain 上累计买到多少 BTC（ZRC20）
    function getUserBalance(bytes memory userId) internal {
            uint256 bal = userBalances[userId];
            emit UserBalance(userId, bal);
    }


    function _executeSwap(
        bytes memory userId, 
        address tokenIn, 
        uint256 amountIn
    ) internal returns (uint256 btcOut) {
        require(amountIn > 0, "Zero amount");
        require(
            tokenIn == USDC_ZRC20,
            "Unsupported token"
        );

        // 1. 先 approve 给 DEX
        IZRC20(tokenIn).approve(address(dex), amountIn);

        // 2. 走 ZetaChain 内置 DEX（总是 ZRC20 → ZRC20）
        address[] memory path = new address[](2);
        path[0] = tokenIn;
        path[1] = BTC_ZRC20;

        uint[] memory amounts = dex.swapExactTokensForTokens(
            amountIn,
            0,                      // 不设最小值
            path,
            address(this),
            block.timestamp + 300
        );

        btcOut = amounts[1];// 输出的 SOL（ZRC20）
        require(btcOut > 0, "Swap returned zero");

        // 3. 存入 VAULT
        userBalances[userId] += btcOut;

        emit BoughtAndStored(
            userId,
            tokenIn,
            amountIn,
            btcOut,
            userBalances[userId]
        );

        return btcOut;
    }

    /// @notice 用户提现（输入希望收到的 USDT 数量）
    function _executeWithdraw(
        bytes memory userId,
        uint256 usdcDesired,
        bytes memory recipient//用户接受地址
    ) internal returns (uint256 btcSpent, uint256 usdtOut) {

        require(recipient.length == 20, "Recipient must be EVM address (20 bytes)");
        address recipientEVM = address(bytes20(recipient));//将bytes变成能用的地址

        // 1. 先反算需要多少 btc（ZRC20）才能换出 usdcDesired
        address[] memory path = new address[](2);
        path[0] = BTC_ZRC20;
        path[1] = USDC_ZRC20;

        uint[] memory quote = dex.getAmountsIn(usdcDesired, path);
        uint256 btcRequired = quote[0];

        require(userBalances[userId] >= btcRequired, "Not enough balance");

         // 扣减余额
        userBalances[userId] -= btcRequired;

        // 2. 用 BTC → USDT swap
        IZRC20(BTC_ZRC20).approve(address(dex), btcRequired);

        uint[] memory results = dex.swapExactTokensForTokens(
            btcRequired,
            usdcDesired,           // 用户希望至少收到这么多 U
            path,
            address(this),
            block.timestamp + 300
        );

        usdtOut = results[1];
        require(usdtOut >= usdcDesired, "Slippage too high");

        // Gas fee 处理
        (address gasZRC20, uint256 gasFee) = IZRC20(USDC_ZRC20).withdrawGasFee();
        require(IZRC20(gasZRC20).balanceOf(address(this)) >= gasFee, "Insufficient gas fee balance");

        // 直接 withdraw（官方 swap 示例常用这个）
        IZRC20(USDC_ZRC20).withdraw(abi.encodePacked(recipientEVM), usdtOut);

        emit WithdrawExecuted(userId, btcRequired, usdtOut, recipientEVM);

        return (btcRequired, usdtOut);
    }

}
