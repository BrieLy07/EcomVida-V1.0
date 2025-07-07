const express = require("express");
const router = express.Router();

const getAll = require("../controllers/get_all_variants");
const getById = require("../controllers/get_variant_by_id");
const create = require("../controllers/create_variant");
const update = require("../controllers/update_variant");
const remove = require("../controllers/delete_variant");

router.get("/variants", getAll);
router.get("/variants/:id", getById);
router.post("/products/:productId/variants", create);
router.put("/variants/:id", update);
router.delete("/variants/:id", remove);

module.exports = router;
