import React, { useState } from 'react';
import { useWalletClient } from 'wagmi';
import { createSchedule } from '../utils/contract';
import { toast } from 'react-toastify';

const CreateSchedule = () => {
  const { data: walletClient } = useWalletClient();

  // 状态变量，必须有这些
  const [tokenIn, setTokenIn] = useState('');
  const [amount, setAmount] = useState('');        // ← 这里定义 amount
  const [interval, setInterval] = useState('');    // ← 这里定义 interval
  const [targetType, setTargetType] = useState('1');

  const handleCreate = async () => {
    if (!walletClient) {
      toast.error('Please connect wallet first');
      return;
    }

    if (!tokenIn || !amount || !interval) {
      toast.error('Please fill all fields');
      return;
    }

    try {
      // eslint-disable-next-line no-undef
      const bigAmount = BigInt(amount || 0);
      // eslint-disable-next-line no-undef
      const bigInterval = BigInt(interval || 0);

      const tx = await createSchedule(
        walletClient,
        tokenIn,
        bigAmount,
        bigInterval,
        Number(targetType)
      );

      // 等待交易确认
      const receipt = await tx.wait();

      // 简单提示成功（黑客松演示不需要精确解析 ID，后面手动输入也行）
      toast.success('Schedule created successfully! Transaction hash: ' + tx.hash);

      // 可选：提示用户手动复制 hash 去 explorer 查看事件获取 ID
      toast.info('Check Base Sepolia explorer for ScheduleCreated event to get your ID');

    } catch (err) {
      console.error(err);
      toast.error(err?.reason || err?.message || 'Transaction failed');
    }
  };

  return (
    <div style={{ marginBottom: '40px' }}>
      <h2>Create DCA Schedule</h2>
      <div style={{ display: 'grid', gap: '10px', maxWidth: '400px' }}>
        <input
          placeholder="Token In Address (e.g. USDC on Base Sepolia)"
          value={tokenIn}
          onChange={(e) => setTokenIn(e.target.value)}
          style={{ padding: '8px' }}
        />
        <input
          type="number"
          placeholder="Amount (in smallest unit, e.g. 1000000 for 1 USDC)"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          style={{ padding: '8px' }}
        />
        <input
          type="number"
          placeholder="Interval in seconds (e.g. 86400 for 1 day)"
          value={interval}
          onChange={(e) => setInterval(e.target.value)}
          style={{ padding: '8px' }}
        />
        <select
          value={targetType}
          onChange={(e) => setTargetType(e.target.value)}
          style={{ padding: '8px' }}
        >
          <option value="1">SOL</option>
          <option value="2">BTC</option>
        </select>
        <button onClick={handleCreate} style={{ padding: '10px', fontSize: '16px' }}>
          Create Schedule
        </button>
      </div>
    </div>
  );
};

export default CreateSchedule;