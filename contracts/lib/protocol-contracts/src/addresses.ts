import { getChainId } from "@zetachain/networks";

import mainnet from "../data/addresses.mainnet.json";
import testnet from "../data/addresses.testnet.json";
import { ParamChainName, ParamSymbol, ParamType } from "./types";

export const getAddress = (type: ParamType, network: ParamChainName, symbol?: ParamSymbol) => {
  const networks = [...testnet, ...mainnet];
  let address;
  if (type !== "zrc20" && symbol) {
    throw new Error("Symbol is only supported when ParamType is zrc20");
  }
  if (type === "zrc20" && !symbol) {
    // for backwards compatibility
    const chainId = getChainId(network);
    address = networks.find((n: any) => {
      return n.foreign_chain_id === chainId?.toString() && n.type === type && n.coin_type === "gas";
    });
  } else {
    address = networks.find((n: any) => {
      return n.chain_name === network && n.type === type && n.symbol === symbol;
    });
  }
  return address?.address;
};
