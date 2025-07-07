const { ScanCommand } = require("@aws-sdk/lib-dynamodb");
const ddb = require("../database/dynamoClient");

class ReembolsosPorOrden {
  constructor(orderId) {
    this.orderId = orderId;
  }

  async ejecutar() {
    const cmd = new ScanCommand({
      TableName: process.env.DYNAMO_TABLE,
      FilterExpression: "orderId = :o",
      ExpressionAttributeValues: { ":o": this.orderId },
    });

    const { Items } = await ddb.send(cmd);
    return Items;
  }
}

module.exports = ReembolsosPorOrden;
