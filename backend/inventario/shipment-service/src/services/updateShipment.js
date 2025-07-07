const esClient = require('../utils/elasticsearch');

module.exports = async function updateShipment(id, data) {
  await esClient.update({
    index: 'shipments',
    id,
    body: { doc: data }
  });
  return { id, ...data };
};
