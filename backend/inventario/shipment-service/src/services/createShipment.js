const esClient = require('../utils/elasticsearch');

module.exports = async function createShipment(data) {
  const response = await esClient.index({
    index: 'shipments',
    body: data
  });
  return { id: response.body._id, ...data };
};
