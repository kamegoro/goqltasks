import React from 'react';

import type { AppProps } from 'next/app';

import ApolloWrapper from '@/components/ApolloWrapper';

const App: React.FC<AppProps> = ({ Component, pageProps }) => {
  return (
    <ApolloWrapper>
      <Component {...pageProps} />
    </ApolloWrapper>
  );
};

export default App;
