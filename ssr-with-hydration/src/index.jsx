import React from 'react';
import ReactDOMServer from 'react-dom/server';
import App from './App';

export const renderApp = (todos) => {
  const content = ReactDOMServer.renderToString(<App todos={todos} />);
  return content;
};