import React, { useEffect } from 'react';
import './App.css';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { WagmiProvider } from 'wagmi';
import { 
  RainbowKitProvider, 
  getDefaultConfig 
} from '@rainbow-me/rainbowkit';
import { baseSepolia } from 'wagmi/chains';
import '@rainbow-me/rainbowkit/styles.css';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from 'react-toastify';

import WalletConnect from './components/WalletConnect';
import CreateSchedule from './components/CreateSchedule';
import ExecuteDCA from './components/ExecuteDCA';

const queryClient = new QueryClient();

const config = getDefaultConfig({
  appName: 'DCA App',
  projectId: '8bf805dc8d790d1d91b4463496b9dd66',
  chains: [baseSepolia],
  ssr: false,
});

// 添加一个组件来处理错误
function ErrorBoundary({ children }) {
  useEffect(() => {
    // 隐藏错误弹出窗口（如果有的话）
    const handleKeyPress = (e) => {
      if (e.key === 'Escape') {
        // 可以添加关闭错误窗口的逻辑
      }
    };
    
    window.addEventListener('keydown', handleKeyPress);
    
    return () => {
      window.removeEventListener('keydown', handleKeyPress);
    };
  }, []);

  return children;
}

function App() {
  return (
    <ErrorBoundary>
      <WagmiProvider config={config}>
        <QueryClientProvider client={queryClient}>
          <RainbowKitProvider>
            <div style={{ padding: '40px', fontFamily: 'sans-serif', maxWidth: '800px', margin: '0 auto' }}>
              <h1>DCA Controller Demo</h1>
              <p><strong>Network:</strong> Base Sepolia</p>
              <WalletConnect />
              <hr style={{ margin: '40px 0' }} />
              <CreateSchedule />
              <hr style={{ margin: '40px 0' }} />
              <ExecuteDCA />
              <ToastContainer position="bottom-right" />
            </div>
          </RainbowKitProvider>
        </QueryClientProvider>
      </WagmiProvider>
    </ErrorBoundary>
  );
}

export default App;