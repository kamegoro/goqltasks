import type { NextPage } from 'next';

import { useQuery } from '@apollo/client';

import { GetTodoDocument } from '@/generated/graphql';
import { GetTodoQuery } from '@/generated/graphql';

const Home: NextPage = () => {
  const { data } = useQuery<GetTodoQuery>(GetTodoDocument);
  return (
    <div>
      <div style={{ margin: '0 auto', width: '1000px' }}>
        <p>Id</p>
        <input type="text" />
        <p>Text</p>
        <input type="text" />
      </div>
      <div style={{ margin: '0 auto', width: '1000px' }}>
        {data?.todos?.map((todo) => (
          <div key={todo.id}>
            <h1>{todo.text}</h1>
            <p>id:{todo.id}</p>
            <p>text:{todo.text}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Home;
