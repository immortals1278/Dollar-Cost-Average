// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@zetachain/protocol-contracts/contracts/zevm/interfaces/IZRC20.sol";
import "@zetachain/protocol-contracts/contracts/interfaces/IMessageRecipient.sol";
import "@zetachain/protocol-contracts/contracts/interfaces/IMessageSender.sol";
import "@zetachain/protocol-contracts/contracts/ZetaInterfaces.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title DCAManager
 * @notice 管理用户的 DCA（定投）计划，负责发跨链消息
 */
contract DCAManager is Ownable, IMessageSender {

    struct DCAPlan {
        address user;
        uint256 amountPerCycle;     // 例如 10 USDT
        uint256 frequency;          // 例如 86400 秒（1天）
        uint256 lastExecuted;       // 上次执行时间戳
        uint256 targetChainId;      // 定投目标链
        bytes targetToken;          // 目标链 token 地址（编码）
        uint256 balance;            // 用户存进去用于定投的金额
        bool active;
    }

    uint256 public nextPlanId = 1;
    mapping(uint256 => DCAPlan) public plans;

    IERC20 public usdt;               // 源链 USDT（你也可以换成其他 token）
    address public zetaConnector;     // ZetaChain connector（每条链有不同地址）

    event PlanCreated(uint256 planId, address user);
    event PlanExecuted(uint256 planId, uint256 amount);
    event PlanCancelled(uint256 planId);

    constructor(address _usdt, address _connector) {
        usdt = IERC20(_usdt);
        zetaConnector = _connector;
    }

    /**
     * 用户创建一个 DCA 计划
     */
    function createDCAPlan(
        uint256 amountPerCycle,
        uint256 frequency,
        uint256 targetChainId,
        bytes calldata targetToken
    ) external returns (uint256) {

        require(amountPerCycle > 0, "amount=0");
        require(frequency > 0, "frequency=0");

        uint256 planId = nextPlanId++;

        plans[planId] = DCAPlan({
            user: msg.sender,
            amountPerCycle: amountPerCycle,
            frequency: frequency,
            lastExecuted: 0,
            targetChainId: targetChainId,
            targetToken: targetToken,
            balance: 0,
            active: true
        });

        emit PlanCreated(planId, msg.sender);
        return planId;
    }

    /**
     * 用户充值 USDT，为他们的定投计划提供资金
     */
    function fundPlan(uint256 planId, uint256 amount) external {
        DCAPlan storage p = plans[planId];
        require(msg.sender == p.user, "not owner");
        require(p.active, "inactive");

        usdt.transferFrom(msg.sender, address(this), amount);
        p.balance += amount;
    }

    /**
     * 后台定时脚本调用：执行一次定投
     */
    function executeDCA(uint256 planId) external {
        DCAPlan storage p = plans[planId];
        require(p.active, "inactive");
        require(
            block.timestamp >= p.lastExecuted + p.frequency,
            "too early"
        );
        require(p.balance >= p.amountPerCycle, "insufficient balance");

        p.balance -= p.amountPerCycle;
        p.lastExecuted = block.timestamp;

        // -------------------------------
        //       发送跨链消息
        // -------------------------------
        bytes memory message = abi.encode(
            p.user,              // 接收者
            p.amountPerCycle,    // 购买数量
            p.targetToken        // 目标 token 地址
        );

        IMessageSender(zetaConnector).send(
            p.targetChainId,     // 目标链
            message
        );

        emit PlanExecuted(planId, p.amountPerCycle);
    }

    /**
     * 用户取消 DCA
     */
    function cancelPlan(uint256 planId) external {
        DCAPlan storage p = plans[planId];
        require(msg.sender == p.user, "not owner");
        p.active = false;

        emit PlanCancelled(planId);
    }
}
