const Categoria = require("../models/category");

exports.crearCategoria = async (req, res) => {
  try {
    const categoria = await Categoria.create(req.body);
    res.status(201).json(categoria);
  } catch (err) {
    res.status(500).json({ error: "Error al crear categoría" });
  }
};

exports.obtenerCategorias = async (_req, res) => {
  try {
    const categorias = await Categoria.findAll();
    res.json(categorias);
  } catch (err) {
    res.status(500).json({ error: "Error al listar categorías" });
  }
};

exports.obtenerCategoriaPorId = async (req, res) => {
  try {
    const categoria = await Categoria.findByPk(req.params.id);
    if (!categoria) return res.status(404).json({ error: "No encontrada" });
    res.json(categoria);
  } catch (err) {
    res.status(500).json({ error: "Error al buscar categoría" });
  }
};
