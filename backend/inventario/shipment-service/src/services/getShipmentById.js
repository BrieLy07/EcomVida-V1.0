const esClient = require('../utils/elasticsearch');

module.exports = async function getShipmentById(id) {
  const result = await esClient.get({
    index: 'shipments',
    id
  });
  return result.body._source;
};
