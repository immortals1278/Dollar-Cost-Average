import React, { useState, useEffect } from 'react';
import { useWalletClient } from 'wagmi';
import { executeDCA } from '../utils/contract';
import { toast } from 'react-toastify';

const ExecuteDCA = () => {
  const { data: walletClient } = useWalletClient();  // 改成 walletClient
  const [id, setId] = useState(localStorage.getItem('scheduleId') || '');
  const [manualLoading, setManualLoading] = useState(false);

  const handleExecute = async () => {
    if (!walletClient) {
      toast.error('Please connect wallet first');
      return;
    }
    if (!id) {
      toast.error('Please enter Schedule ID');
      return;
    }

    setManualLoading(true);
    try {
      // eslint-disable-next-line no-undef
      const bigId = BigInt(id);

      await executeDCA(walletClient, bigId);
      toast.success(`DCA executed for ID ${id}`);
    } catch (err) {
      console.error(err);
      toast.error(err.message || 'Execute failed');
    }
    setManualLoading(false);
  };

  return (
    <div>
      <h2>Execute DCA</h2>
      <input 
        placeholder="Schedule ID" 
        value={id} 
        onChange={e => {
          setId(e.target.value);
          localStorage.setItem('scheduleId', e.target.value);
        }} 
      />
      <br /><br />
      <button onClick={handleExecute} disabled={manualLoading}>
        {manualLoading ? 'Executing...' : 'Manual Execute DCA'}
      </button>
      <p><small>Tip: Create a schedule first to get ID</small></p>
    </div>
  );
};

export default ExecuteDCA;