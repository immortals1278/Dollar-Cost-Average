import { ethers } from 'ethers';
import ABI from '../contracts/DCAControllerABI.json';

const CONTRACT_ADDRESS = '0x35aA5A3bOEf7739D953A72ca4480dB88b52DEd70';

export const getContract = async (walletClient) => {
  if (!walletClient) throw new Error('Wallet not connected');
  // wagmi v2 的 walletClient 转 ethers signer
  const provider = new ethers.providers.Web3Provider(walletClient.transport);
  const signer = provider.getSigner(await walletClient.getAddresses()[0]);
  return new ethers.Contract(CONTRACT_ADDRESS, ABI, signer);
};

export const createSchedule = async (walletClient, tokenIn, amount, interval, targetType) => {
  const contract = await getContract(walletClient);
  return await contract.createSchedule(tokenIn, amount, interval, targetType);
};

export const executeDCA = async (walletClient, id) => {
  const contract = await getContract(walletClient);
  return await contract.executeDCA(id);
};