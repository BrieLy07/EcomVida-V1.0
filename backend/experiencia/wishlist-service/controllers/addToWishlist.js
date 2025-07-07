import { addToWishlist } from "../models/wishlistModel.js";
import { observador } from "../observers/wishlistObserver.js";

export const agregarAWishlist = async (req, res) => {
  try {
    const usuario_id = req.headers["x-user-id"];
    const { producto_id } = req.body;

    if (!usuario_id || !producto_id)
      return res.status(400).json({ mensaje: "Datos incompletos" });

    const nuevo = await addToWishlist(usuario_id, producto_id);
    observador.notificarAccion(usuario_id, "agregado", producto_id);
    res.status(201).json(nuevo);
  } catch (error) {
    res.status(500).json({ mensaje: "Error al agregar a wishlist", error });
  }
};
