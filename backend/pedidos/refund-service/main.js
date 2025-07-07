const express = require("express");
const app = express();
require("dotenv").config();

const createRefund = require("./app/controllers/createRefund");
const getRefundById = require("./app/controllers/getRefundById");
const getRefundsByOrder = require("./app/controllers/getRefundsByOrder");

app.use(express.json());

// Rutas
app.post("/refunds", createRefund);
app.get("/refunds/:id", getRefundById);
app.get("/refunds/order/:orderId", getRefundsByOrder);

const PORT = process.env.PORT || 3014;
app.listen(PORT, () => {
  console.log(`Refund service corriendo en el puerto ${PORT}`);
});
