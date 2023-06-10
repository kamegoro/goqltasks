import React from 'react';

import type { AppProps } from 'next/app';

import ApolloWrapper from '@/components/ApolloWrapper';
import '@/styles/globals.css';

const App: React.FC<AppProps> = ({ Component, pageProps }) => {
  return (
    <ApolloWrapper>
      <Component {...pageProps} />
    </ApolloWrapper>
  );
};

export default App;
