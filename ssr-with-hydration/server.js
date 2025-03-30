require('@babel/register')({
    extensions: ['.js', '.jsx'],
    ignore: [/node_modules/],
  });
  
  const express = require('express');
  const path = require('path');
  const bodyParser = require('body-parser');
  const { renderApp } = require('./src/index.jsx');
  
  const app = express();
  const PORT = 3000;
  
  let todos = [];
  
  app.use(express.static(path.join(__dirname, 'public')));
  app.use(bodyParser.urlencoded({ extended: true }));
  
  // Home page render
  app.get('/', (req, res) => {
    const appHtml = renderApp(todos);
  
    const htmlTemplate = `
      <!DOCTYPE html>
      <html lang="en">
      <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
        <title>SSR TODO App</title>
      </head>
      <body>
        <div id="root">${appHtml}</div>
      </body>
      </html>
    `;
  
    res.send(htmlTemplate);
  });
  
  // Add Todo
  app.post('/add', (req, res) => {
    const { title } = req.body;
    todos.push({ title });
    res.redirect('/');
  });
  
  // Update Todo
  app.post('/update', (req, res) => {
    const { id, title } = req.body;
    if (todos[id]) {
      todos[id].title = title;
    }
    res.redirect('/');
  });
  
  // Delete Todo
  app.post('/delete', (req, res) => {
    const { id } = req.body;
    todos.splice(id, 1);
    res.redirect('/');
  });
  
  app.listen(PORT, () => {
    console.log(`✅ TODO App running at http://localhost:${PORT}`);
  });