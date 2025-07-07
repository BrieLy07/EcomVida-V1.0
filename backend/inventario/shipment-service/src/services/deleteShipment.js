const esClient = require('../utils/elasticsearch');

module.exports = async function deleteShipment(id) {
  await esClient.delete({
    index: 'shipments',
    id
  });
  return { mensaje: 'Envío eliminado' };
};
