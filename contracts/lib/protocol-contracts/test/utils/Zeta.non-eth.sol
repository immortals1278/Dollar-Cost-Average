// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

/**
 * @dev Common custom errors
 */
interface ZetaErrors {
    // @dev Thrown when caller is not the address defined as TSS address
    error CallerIsNotTss(address caller);

    // @dev Thrown when caller is not the address defined as connector address
    error CallerIsNotConnector(address caller);

    // @dev Thrown when caller is not the address defined as TSS Updater address
    error CallerIsNotTssUpdater(address caller);

    // @dev Thrown when caller is not the address defined as TSS or TSS Updater address
    error CallerIsNotTssOrUpdater(address caller);

    // @dev Thrown when a contract receives an invalid address param, mostly zero address validation
    error InvalidAddress();

    // @dev Thrown when Zeta can't be transferred for some reason
    error ZetaTransferError();
}

/**
 * @dev ZetaNonEthInterface is a mintable / burnable version of IERC20
 */
interface ZetaNonEthInterface is IERC20 {
    function burnFrom(address account, uint256 amount) external;

    function mint(address mintee, uint256 value, bytes32 internalSendHash) external;
}

contract ZetaNonEth is ZetaNonEthInterface, ERC20Burnable, ZetaErrors {
    address public connectorAddress;

    /**
     * @dev Collectively held by Zeta blockchain validators
     */
    address public tssAddress;

    /**
     * @dev Initially a multi-sig, eventually held by Zeta blockchain validators (via renounceTssAddressUpdater)
     */
    address public tssAddressUpdater;

    event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash);

    event Burnt(address indexed burnee, uint256 amount);

    event TSSAddressUpdated(address callerAddress, address newTssAddress);

    event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress);

    event ConnectorAddressUpdated(address callerAddress, address newConnectorAddress);

    constructor(address tssAddress_, address tssAddressUpdater_) ERC20("Zeta", "ZETA") {
        if (tssAddress_ == address(0) || tssAddressUpdater_ == address(0)) revert InvalidAddress();

        tssAddress = tssAddress_;
        tssAddressUpdater = tssAddressUpdater_;
    }

    function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) external {
        if (msg.sender != tssAddressUpdater && msg.sender != tssAddress) revert CallerIsNotTssOrUpdater(msg.sender);
        if (tssAddress_ == address(0) || connectorAddress_ == address(0)) revert InvalidAddress();

        tssAddress = tssAddress_;
        connectorAddress = connectorAddress_;

        emit TSSAddressUpdated(msg.sender, tssAddress_);
        emit ConnectorAddressUpdated(msg.sender, connectorAddress_);
    }

    /**
     * @dev Sets tssAddressUpdater to be tssAddress
     */
    function renounceTssAddressUpdater() external {
        if (msg.sender != tssAddressUpdater) revert CallerIsNotTssUpdater(msg.sender);
        if (tssAddress == address(0)) revert InvalidAddress();

        tssAddressUpdater = tssAddress;
        emit TSSAddressUpdaterUpdated(msg.sender, tssAddress);
    }

    function mint(address mintee, uint256 value, bytes32 internalSendHash) external override {
        /**
         * @dev Only Connector can mint. Minting requires burning the equivalent amount on another chain
         */
        if (msg.sender != connectorAddress) revert CallerIsNotConnector(msg.sender);

        _mint(mintee, value);

        emit Minted(mintee, value, internalSendHash);
    }

    function burnFrom(address account, uint256 amount) public override(ZetaNonEthInterface, ERC20Burnable) {
        /**
         * @dev Only Connector can burn.
         */
        if (msg.sender != connectorAddress) revert CallerIsNotConnector(msg.sender);

        ERC20Burnable.burnFrom(account, amount);

        emit Burnt(account, amount);
    }
}
