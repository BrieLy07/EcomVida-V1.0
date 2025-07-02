const express = require('express');
const cors = require('cors'); // 👈 importar cors
const app = express();
require('dotenv').config();

app.use(cors()); // 👈 habilitar CORS para todos los orígenes

const registerRoute = require('./routes/register');
const loginRoute = require('./routes/login');
const logoutRoute = require('./routes/logout');

app.use(express.json());
app.use('/register', registerRoute);
app.use('/login', loginRoute);
app.use('/logout', logoutRoute);

module.exports = app;
