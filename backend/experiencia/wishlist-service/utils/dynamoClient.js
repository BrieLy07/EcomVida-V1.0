import AWS from 'aws-sdk';
import dotenv from 'dotenv';

dotenv.config();

const dynamo = new AWS.DynamoDB.DocumentClient({
  region: process.env.AWS_REGION,
  endpoint: process.env.DYNAMO_ENDPOINT
});

export default dynamo;
