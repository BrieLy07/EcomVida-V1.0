const esClient = require('../utils/elasticsearch');

module.exports = async function getAllShipments() {
  const result = await esClient.search({
    index: 'shipments',
    size: 1000,
    body: { query: { match_all: {} } }
  });
  return result.body.hits.hits.map(hit => ({
    id: hit._id,
    ...hit._source
  }));
};
