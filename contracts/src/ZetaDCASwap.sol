// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

// ============= ZetaChain 官方接口 =============
import "@zetachain/protocol-contracts/contracts/evm/interfaces/UniversalContract.sol";
import "@zetachain/protocol-contracts/contracts/evm/interfaces/IZRC20.sol";
import "@zetachain/protocol-contracts/contracts/evm/interfaces/IGatewayEVM.sol";

// ============= UniswapV2 风格 DEX (Zeta DEX) =============
interface IUniswapV2Router {
    function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);
}

contract ZetaDCAExecution is UniversalContract {
    IGatewayEVM public immutable gateway;
    IUniswapV2Router public immutable dex;

    
    mapping(bytes => uint256) public userBalances; 

    /// @notice ZRC20 地址
    address public immutable USDT_ZRC20;
    address public immutable USDC_ZRC20;
    address public immutable SOL_ZRC20;

    event BoughtAndStored(
        bytes indexed userId,       // 源链用户自定义 ID
        address indexed tokenIn,    // 输入的 ZRC20
        uint256 amountIn,
        uint256 solOut,
        uint256 totalStored
    );
    event UserBalance(
        bytes indexed userId,
        uint256 balance
    );

    constructor(
        address _gateway,
        address _dex,
        address _usdt,
        address _usdc,
        address _sol
    ) {
        gateway = IGatewayEVM(_gateway);
        dex = IUniswapV2Router(_dex);
        USDT_ZRC20 = _usdt;
        USDC_ZRC20 = _usdc;
        SOL_ZRC20  = _sol;
    }

    struct SwapMsg {
        bytes userId;       // 源链传过来的用户标识（比如 keccak(user address) 或钱包地址）
        address tokenIn;    // 输入 token (USDT/USDC/ETH 的 ZRC20)
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

    /// @notice 查询用户在 ZetaChain 上累计买到多少 SOL（ZRC20）
    function getUserBalance(bytes calldata userId) internal {
            uint256 bal = userBalances[userId];
            emit UserBalance(userId, bal);
    }

    function _executeSwap(
        bytes memory userId, 
        address tokenIn, 
        uint256 amountIn
    ) internal returns (uint256 solOut) {
        require(amountIn > 0, "Zero amount");
        require(
            tokenIn == USDT_ZRC20 || tokenIn == USDC_ZRC20,
            "Unsupported token"
        );

        // 1. 先 approve 给 DEX
        IZRC20(tokenIn).approve(address(dex), amountIn);

        // 2. 走 ZetaChain 内置 DEX（总是 ZRC20 → ZRC20）
        address;
        path[0] = tokenIn;
        path[1] = SOL_ZRC20;

        uint[] memory amounts = dex.swapExactTokensForTokens(
            amountIn,
            0,                      // 不设最小值（你可以加）
            path,
            address(this),
            block.timestamp + 300
        );

        solOut = amounts[1];
        require(solOut > 0, "Swap returned zero");

        // 3. 存入 VAULT
        userBalances[userId] += solOut;

        emit BoughtAndStored(
            userId,
            tokenIn,
            amountIn,
            solOut,
            userBalances[userId]
        );

        return solOut;
    }

    /// @notice 用户提现（输入希望收到的 USDT 数量）
    function _executeWithdraw(
        bytes memory userId,
        uint256 usdtDesired,
        bytes memory recipient
    ) internal returns (uint256 solSpent, uint256 usdtOut) {

        require(recipient.length == 20, "Recipient must be EVM address (20 bytes)");
        address recipientEVM = address(bytes20(recipient));

        // 1. 先反算需要多少 SOL（ZRC20）才能换出 usdtDesired
        address;
        path[0] = SOL_ZRC20;
        path[1] = USDT_ZRC20;

        uint[] memory quote = dex.getAmountsIn(usdtDesired, path);
        uint256 solRequired = quote[0];

        require(userBalances[userId] >= solRequired, "Not enough balance");

        // 扣减余额
        userBalances[userId] -= solRequired;

        // 2. 用 SOL → USDT swap
        IZRC20(SOL_ZRC20).approve(address(dex), solRequired);

        uint[] memory results = dex.swapExactTokensForTokens(
            solRequired,
            usdtDesired,           // 用户希望至少收到这么多 U
            path,
            address(this),
            block.timestamp + 300
        );

        usdtOut = results[1];
        require(usdtOut >= usdtDesired, "Slippage too high");

        // 3. approve USDT 给 gateway
        IZRC20(USDT_ZRC20).approve(address(gateway), usdtOut);

        // 4. 从 ZetaChain withdraw 回源链（用户钱包）
        gateway.withdraw(
            recipient,           // 20 bytes user EVM address
            usdtOut,
            USDT_ZRC20,
            RevertOptions({
                revertAddress: address(0),
                callOnRevert: false,
                abortAddress: address(0),
                revertMessage: "",
                onRevertGasLimit: 0
            })
        );

        emit WithdrawExecuted(userId, solRequired, usdtOut, recipientEVM);

        return (solRequired, usdtOut);
    }

}
