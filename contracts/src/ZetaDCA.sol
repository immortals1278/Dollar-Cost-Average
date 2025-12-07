// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./DCARegistry.sol";
import "./ZetaSwapExecutor.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ZetaDCA is Ownable {

    DCARegistry public immutable registry;
    ZetaSwapExecutor public immutable executor;

    constructor(address _registry, address _executor) {
        registry = DCARegistry(_registry);
        executor = ZetaSwapExecutor(_executor);
    }

    /// keeper / 后台定时执行 DCA
    function execDCA(uint256 planId) external {
        DCARegistry.Plan memory p = registry.plans(planId);

        require(p.active, "inactive");
        require(block.timestamp >= p.lastExec + p.interval, "not time");

        uint256 amountOut = executor.executeSwap(
            p.depositToken,
            p.amount,
            p.targetZRC20,
            p.user
        );

        registry.plans(planId).lastExec = block.timestamp;

        emit DCAExecuted(planId, amountOut);
    }

    event DCAExecuted(uint256 planId, uint256 amountOut);
}
