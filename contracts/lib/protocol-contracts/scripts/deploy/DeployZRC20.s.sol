// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "../../contracts/zevm/ZRC20.sol";
import "../../contracts/zevm/interfaces/IZRC20.sol";

contract DeployZRC20 is Script {
    function run() external {
        address gateway = vm.envAddress("GATEWAY_PROXY_ZEVM");
        address systemContract = vm.envAddress("SYSTEM_CONTRACT");
        string memory name = vm.envString("ZRC20_NAME");
        string memory symbol = vm.envString("ZRC20_SYMBOL");
        uint8 decimals = uint8(vm.envUint("ZRC20_DECIMALS"));
        uint chainID = vm.envUint("ZRC20_CHAIN_ID");
        CoinType coinType = CoinType(vm.envUint("ZRC20_COIN_TYPE"));
        uint gasLimit = vm.envUint("ZRC20_GAS_LIMIT");

        vm.startBroadcast();

        ZRC20 zrc20 = new ZRC20(
            name,
            symbol,
            decimals,
            chainID,
            coinType,
            gasLimit,
            systemContract,
            gateway
        );

        require(address(zrc20) != address(0), "deployment failed");

        require(keccak256(abi.encodePacked(zrc20.name())) == keccak256(abi.encodePacked(name)), "name not set");
        require(keccak256(abi.encodePacked(zrc20.symbol())) == keccak256(abi.encodePacked(symbol)), "symbol not set");
        require(zrc20.decimals() == decimals, "decimals not set");
        require(zrc20.CHAIN_ID() == chainID, "chain id not set");
        require(zrc20.GAS_LIMIT() == gasLimit, "gas limit not set");
        require(zrc20.COIN_TYPE() == coinType, "coin type not set");
        require(zrc20.gatewayAddress() == gateway, "gateway not set");
        require(zrc20.SYSTEM_CONTRACT_ADDRESS() == systemContract, "system contract not set");

        vm.stopBroadcast();
    }
}
