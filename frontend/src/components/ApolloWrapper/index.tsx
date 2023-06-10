import { ReactNode } from 'react';

import { ApolloClient, InMemoryCache, createHttpLink } from '@apollo/client';
import { ApolloProvider } from '@apollo/client';

interface ApolloWrapperType {
  children: ReactNode;
}

const ApolloWrapper: React.FC<ApolloWrapperType> = ({ children }) => {
  const link = createHttpLink({
    uri: 'http://localhost:8080/query',
    credentials: 'include',
  });

  const client = new ApolloClient({
    cache: new InMemoryCache(),
    link: link,
  });

  return <ApolloProvider client={client}>{children}</ApolloProvider>;
};

export default ApolloWrapper;
