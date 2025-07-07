import { getWishlist } from "../models/wishlistModel.js";

export const obtenerWishlist = async (req, res) => {
  try {
    const usuario_id = req.headers["x-user-id"];
    if (!usuario_id) return res.status(400).json({ mensaje: "Falta ID de usuario" });

    const lista = await getWishlist(usuario_id);
    res.json(lista);
  } catch (error) {
    res.status(500).json({ mensaje: "Error al obtener wishlist", error });
  }
};
