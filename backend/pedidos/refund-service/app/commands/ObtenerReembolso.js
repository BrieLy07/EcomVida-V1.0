const { GetCommand } = require("@aws-sdk/lib-dynamodb");
const ddb = require("../database/dynamoClient");

class ObtenerReembolso {
  constructor(id) {
    this.id = id;
  }

  async ejecutar() {
    const cmd = new GetCommand({
      TableName: process.env.DYNAMO_TABLE,
      Key: { id: this.id },
    });

    const { Item } = await ddb.send(cmd);
    return Item;
  }
}

module.exports = ObtenerReembolso;
