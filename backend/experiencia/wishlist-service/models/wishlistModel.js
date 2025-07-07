import dynamo from "../utils/dynamoClient.js";
import { v4 as uuidv4 } from "uuid";
import dotenv from "dotenv";
dotenv.config();

const TABLE = process.env.DYNAMO_TABLE;

export const getWishlist = async (usuario_id) => {
  const params = {
    TableName: TABLE,
    KeyConditionExpression: "usuario_id = :uid",
    ExpressionAttributeValues: {
      ":uid": usuario_id
    }
  };
  const result = await dynamo.query(params).promise();
  return result.Items;
};

export const addToWishlist = async (usuario_id, producto_id) => {
  const item = {
    usuario_id,
    producto_id,
    id: uuidv4(),
    fecha: new Date().toISOString()
  };
  const params = {
    TableName: TABLE,
    Item: item
  };
  await dynamo.put(params).promise();
  return item;
};

export const removeFromWishlist = async (usuario_id, producto_id) => {
  const params = {
    TableName: TABLE,
    Key: {
      usuario_id,
      producto_id
    }
  };
  await dynamo.delete(params).promise();
  return true;
};
