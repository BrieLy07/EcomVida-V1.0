const { Client } = require('@elastic/elasticsearch');
require('dotenv').config();

const esClient = new Client({
  node: process.env.ELASTIC_NODE || 'http://localhost:9200'
});

esClient.ping()
  .then(() => console.log('✅ Conectado a Elasticsearch'))
  .catch(err => console.error('❌ Error al conectar a Elasticsearch', err));

module.exports = esClient;
