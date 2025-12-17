pragma solidity ^0.8.0;

import "../contracts/zevm/WZETA.sol";
import "forge-std/Test.sol";

contract WZETATest is Test {
    WETH9 wzeta;

    address alice = address(0xabc);
    address bob = address(0xdef);

    function setUp() public {
        wzeta = new WETH9();
        vm.deal(alice, 10 ether);
        vm.deal(bob, 10 ether);
    }

    function testDeposit() public {
        vm.startPrank(alice);
        wzeta.deposit{ value: 1 ether }();

        assertEq(wzeta.balanceOf(alice), 1 ether);
        assertEq(wzeta.totalSupply(), 1 ether);

        vm.stopPrank();
    }

    function testWithdraw() public {
        vm.startPrank(alice);
        wzeta.deposit{ value: 2 ether }();

        wzeta.withdraw(1 ether);

        assertEq(wzeta.balanceOf(alice), 1 ether);
        assertEq(alice.balance, 9 ether);
        assertEq(wzeta.totalSupply(), 1 ether);

        vm.stopPrank();
    }

    function testTransfer() public {
        vm.startPrank(alice);
        wzeta.deposit{ value: 2 ether }();

        wzeta.transfer(bob, 1 ether);

        assertEq(wzeta.balanceOf(alice), 1 ether);
        assertEq(wzeta.balanceOf(bob), 1 ether);

        vm.stopPrank();
    }

    function testApproveAndTransferFrom() public {
        vm.startPrank(alice);
        wzeta.deposit{ value: 2 ether }();

        wzeta.approve(bob, 1 ether);
        assertEq(wzeta.allowance(alice, bob), 1 ether);

        vm.stopPrank();

        vm.startPrank(bob);
        wzeta.transferFrom(alice, bob, 1 ether);

        assertEq(wzeta.balanceOf(alice), 1 ether);
        assertEq(wzeta.balanceOf(bob), 1 ether);
        assertEq(wzeta.allowance(alice, bob), 0);

        vm.stopPrank();
    }

    function testDepositAndReceiveFallback() public {
        vm.prank(alice);
        (bool success,) = address(wzeta).call{ value: 1 ether }("");
        assertTrue(success);

        assertEq(wzeta.balanceOf(alice), 1 ether);
        assertEq(wzeta.totalSupply(), 1 ether);
    }

    function testWithdrawRevertsIfInsufficientBalance() public {
        vm.startPrank(alice);
        wzeta.deposit{ value: 1 ether }();

        vm.expectRevert();
        wzeta.withdraw(2 ether);

        vm.stopPrank();
    }

    function testTransferRevertsIfInsufficientBalance() public {
        vm.startPrank(alice);
        wzeta.deposit{ value: 1 ether }();

        vm.expectRevert();
        wzeta.transfer(bob, 2 ether);

        vm.stopPrank();
    }

    function testTransferFromRevertsIfInsufficientAllowance() public {
        vm.startPrank(alice);
        wzeta.deposit{ value: 1 ether }();

        wzeta.approve(bob, 0.5 ether);

        vm.stopPrank();

        vm.startPrank(bob);
        vm.expectRevert();
        wzeta.transferFrom(alice, bob, 1 ether);

        vm.stopPrank();
    }

    function testTotalSupply() public {
        vm.startPrank(alice);
        wzeta.deposit{ value: 1 ether }();

        vm.stopPrank();

        vm.startPrank(bob);
        wzeta.deposit{ value: 2 ether }();

        vm.stopPrank();

        assertEq(wzeta.totalSupply(), 3 ether);
    }
}
