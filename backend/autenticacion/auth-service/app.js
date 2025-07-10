const express = require('express');
const cors = require('cors');
const dotenv = require('dotenv');
dotenv.config();

const registerRoute = require('./routes/register');
const loginRoute = require('./routes/login');

const app = express();
app.use(cors());
app.use(express.json());

app.use('/register', registerRoute);
app.use('/login', loginRoute);

const PORT = process.env.PORT || 3001;
app.listen(PORT, () => {
  console.log(`Auth service corriendo en puerto ${PORT}`);
});
