const express = require("express");
const app = express();
const rutasCategoria = require("./routes/categoryRoutes");

app.use(express.json());
app.use("/api", rutasCategoria);

module.exports = app;
