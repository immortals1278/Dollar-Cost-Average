// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract DCARegistry {
    struct Plan {
        address user;
        address depositToken; 
        uint256 amount;
        address targetZRC20;
        uint256 interval;
        uint256 lastExec;
        bool active;
    }

    mapping(uint256 => Plan) public plans;
    uint256 public nextId;

    event PlanCreated(uint256 planId, address user);
    event PlanExecuted(uint256 planId);
    event PlanCancelled(uint256 planId);

    function createPlan(
        address depositToken,
        uint256 amount,
        address targetZRC20,
        uint256 interval
    ) external returns (uint256 id) {
        id = nextId++;

        plans[id] = Plan({
            user: msg.sender,
            depositToken: depositToken,
            amount: amount,
            targetZRC20: targetZRC20,
            interval: interval,
            lastExec: block.timestamp,
            active: true
        });

        emit PlanCreated(id, msg.sender);
    }

    function cancelPlan(uint256 id) external {
        require(plans[id].user == msg.sender, "not owner");
        plans[id].active = false;
        emit PlanCancelled(id);
    }
}
