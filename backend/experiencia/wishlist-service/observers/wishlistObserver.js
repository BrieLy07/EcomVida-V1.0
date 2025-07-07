export class WishlistObserver {
  notificarAccion(usuario_id, accion, producto_id) {
    console.log(`[OBSERVADOR] Usuario ${usuario_id} ha ${accion} el producto ${producto_id} a la wishlist.`);
    // Aquí puedes agregar lógica adicional como eventos, logs, etc.
  }
}

export const observador = new WishlistObserver();
