import React from 'react';
import './App.css';
import {useSuspenseQuery} from '@connectrpc/connect-query';
import {getTasks} from "../../gen/proto/task/v1/service-TaskService_connectquery";

function App() {
  const {data} = useSuspenseQuery(getTasks)
  return (
    <div>
      {JSON.stringify(data, null, 2)}
    </div>
  );
}

export default App;
