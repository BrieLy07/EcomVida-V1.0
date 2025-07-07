import { removeFromWishlist } from "../models/wishlistModel.js";
import { observador } from "../observers/wishlistObserver.js";

export const eliminarDeWishlist = async (req, res) => {
  try {
    const usuario_id = req.headers["x-user-id"];
    const { productId } = req.params;

    if (!usuario_id || !productId)
      return res.status(400).json({ mensaje: "Datos incompletos" });

    await removeFromWishlist(usuario_id, productId);
    observador.notificarAccion(usuario_id, "eliminado", productId);
    res.json({ mensaje: "Producto eliminado de wishlist" });
  } catch (error) {
    res.status(500).json({ mensaje: "Error al eliminar producto", error });
  }
};
