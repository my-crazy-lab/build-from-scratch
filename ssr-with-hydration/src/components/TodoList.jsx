import React from 'react';

const TodoList = ({ todos }) => {
  return (
    <div>
      <h1>TODO List</h1>
      <form method="POST" action="/add">
        <input type="text" name="title" placeholder="Add new todo" required />
        <button type="submit">Add</button>
      </form>
      <ul>
        {todos.map((todo, index) => (
          <li key={index}>
            <form method="POST" action="/update" style={{ display: 'inline' }}>
              <input type="hidden" name="id" value={index} />
              <input type="text" name="title" defaultValue={todo.title} />
              <button type="submit">Update</button>
            </form>
            <form method="POST" action="/delete" style={{ display: 'inline', marginLeft: '10px' }}>
              <input type="hidden" name="id" value={index} />
              <button type="submit">Delete</button>
            </form>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default TodoList;