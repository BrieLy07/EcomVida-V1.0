const { PutCommand } = require("@aws-sdk/lib-dynamodb");
const ddb = require("../database/dynamoClient");
const { v4: uuidv4 } = require("uuid");

class CrearReembolso {
  constructor(datos) {
    this.datos = datos;
  }

  async ejecutar() {
    const reembolso = {
      id: uuidv4(),
      orderId: this.datos.orderId,
      motivo: this.datos.motivo,
      monto: this.datos.monto,
      estado: "pendiente",
      fecha: new Date().toISOString(),
    };

    const cmd = new PutCommand({
      TableName: process.env.DYNAMO_TABLE,
      Item: reembolso,
    });

    await ddb.send(cmd);
    return reembolso;
  }
}

module.exports = CrearReembolso;
