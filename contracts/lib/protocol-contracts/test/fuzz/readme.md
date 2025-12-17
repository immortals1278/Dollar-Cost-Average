## Setup echidna

```
brew install echidna
solc-select use 0.8.20
```

## Execute contract tests

```
echidna test/fuzz/ERC20CustodyEchidnaTest.sol --contract ERC20CustodyEchidnaTest  --config echidna.yaml
echidna test/fuzz/GatewayEVMEchidnaTest.sol --contract GatewayEVMEchidnaTest  --config echidna.yaml
```