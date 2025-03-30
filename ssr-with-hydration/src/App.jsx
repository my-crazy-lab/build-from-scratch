import React from 'react';
import TodoList from './components/TodoList';

const App = ({ todos }) => {
  return <TodoList todos={todos} />;
};

export default App;