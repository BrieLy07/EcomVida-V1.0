const express = require("express");
const router = express.Router();
const controller = require("../controllers/categoryController");

router.post("/categorias", controller.crearCategoria);
router.get("/categorias", controller.obtenerCategorias);
router.get("/categorias/:id", controller.obtenerCategoriaPorId);

module.exports = router;
