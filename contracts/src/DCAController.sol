// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

interface IGatewayEVM {
    function callContract(
        uint256 destChainId,
        address destContract,
        bytes calldata message,
        bytes calldata callOptions
    ) external payable;
}

contract DCAController {

    IGatewayEVM public gateway;
    address public universalApp;         // 通用合约地址
    address public universalAppBTC;  // BTC通用合约地址
    uint256 public zetaChainId;          

    
    struct DCASchedule {
        address user;            // 用户
        uint8 targetType;  // 1 = SOL, 2 = BTC
        address tokenIn;         //投入币种
        uint256 amount;          // 每次投入金额（源链 token）
        uint256 interval;        // 间隔时间
        uint256 nextExec;        // 下次执行时间
        uint256 totalInvested;   // 累计投入金额
        bool active;
    }

    // scheduleId => DCASchedule
    mapping(uint256 => DCASchedule) public schedules;

    uint256 public nextScheduleId = 1;

    event ScheduleCreated(uint256 indexed id, address user, uint256 amount);
    event ScheduleExecuted(uint256 indexed id, uint256 amount, uint256 totalInvested);
    event WithdrawRequested(uint256 indexed id, uint256 amount);
    event ScheduleStopped(uint256 indexed id);
    event QueryRequested(uint256 indexed id);

    constructor(
        address _gateway,
        address _universalApp,
        address _universalAppBTC,
        uint256 _zetaChainId
    ) {
        gateway = IGatewayEVM(_gateway);
        universalApp = _universalApp;
        universalAppBTC = _universalAppBTC;
        zetaChainId = _zetaChainId;
    }

    // 创建定投计划
    
    function createSchedule(
        address tokenIn,
        uint256 amount,
        uint256 interval,
        uint8 targetType
    ) external returns (uint256 id) {
        id = nextScheduleId++;
        schedules[id] = DCASchedule({
            user: msg.sender,
            targetType: targetType,
            tokenIn: tokenIn,
            amount: amount,
            interval: interval,
            nextExec: block.timestamp + interval,
            totalInvested: 0,
            active: true
        });
        require(targetType == 1 || targetType == 2, "Invalid target");
        emit ScheduleCreated(id, msg.sender, amount);
    }

    // 执行定投
    
    function executeDCA(uint256 id) external {
        DCASchedule storage s = schedules[id];
        require(s.active, "inactive");
        require(block.timestamp >= s.nextExec, "too early");

        // 源链 tokenIn 由前端或脚本先 transferFrom 到 controller

        // 累计投入金额
        s.totalInvested += s.amount;

        // 安排下次执行
        s.nextExec = block.timestamp + s.interval;

        
        // opType = 1 (swap)
        // params = (bytes userId, address tokenIn)
        bytes memory userId = abi.encodePacked(id);

        bytes memory message = abi.encode(
            uint8(1),                 // opType = 1 = swap
            abi.encode(
                userId,
                s.tokenIn
            )
        );

        address targetApp = s.targetType == 1 ? universalApp : universalAppBTC;
        
        // 送到 ZetaChain 执行 swap
        gateway.callContract(
            zetaChainId,
            targetApp,
            message,
            "" 
        );

        emit ScheduleExecuted(id, s.amount, s.totalInvested);
    }

    // ============================================================
    // 🔹 3. 发起提现 —— 源链调用通用合约（swap back + withdraw）
    // ============================================================
    function requestWithdraw(
        uint256 id,
        uint256 wantUSDT
    ) external {
        DCASchedule storage s = schedules[id];
        require(msg.sender == s.user, "not owner");
        require(s.active, "inactive");

        // 更新累计投入金额（减少）
        require(s.totalInvested >= wantUSDT, "exceeds total invested");
        s.totalInvested -= wantUSDT;

        // opType = 3 withdraw
        bytes memory userId = abi.encodePacked(id);

        // params = (bytes userId, uint256 wantUSDT, bytes recipient)
        bytes memory message = abi.encode(
            uint8(3),
            abi.encode(
                userId,
                wantUSDT,
                abi.encodePacked(msg.sender) //用户接受地址
            )
        );

        address targetApp = s.targetType == 1 ? universalApp : universalAppBTC;

        gateway.callContract(
            zetaChainId,
            targetApp,
            message,
            ""
        );

        emit WithdrawRequested(id, wantUSDT);
    }

    // ============================================================
    // 🔹 4. 查询用户累计投资（跨链从 ZetaChain 查询）
    // ============================================================
    function queryUserBalance(uint256 id) external {
        require(schedules[id].active, "inactive");

        bytes memory userId = abi.encodePacked(id);

        // opType = 2 query
        bytes memory message = abi.encode(
            uint8(2),
            abi.encode(userId)
        );

        DCASchedule storage s = schedules[id];
        address targetApp = s.targetType == 1 ? universalApp : universalAppBTC;

        gateway.callContract(
            zetaChainId,
            targetApp,
            message,
            ""
        );

        emit QueryRequested(id);
    }

    // ============================================================
    // 🔹 5. 停止定投
    // ============================================================
    function stopDCA(uint256 id) external {
        require(msg.sender == schedules[id].user);
        schedules[id].active = false;
        emit ScheduleStopped(id);
    }
}
