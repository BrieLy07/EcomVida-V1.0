from app.db import obtener_conexion
from app.models.producto import Producto

class ProductoRepository:

    @staticmethod
    def obtener_todos():
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute("SELECT id, nombre, descripcion, precio, stock FROM productos")
        rows = cursor.fetchall()
        conn.close()
        return [Producto(*row) for row in rows]

    @staticmethod
    def obtener_por_id(producto_id):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute("SELECT id, nombre, descripcion, precio, stock FROM productos WHERE id = %s", (producto_id,))
        row = cursor.fetchone()
        conn.close()
        return Producto(*row) if row else None

    @staticmethod
    def crear(producto: Producto):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute(
            "INSERT INTO productos (nombre, descripcion, precio, stock) VALUES (%s, %s, %s, %s) RETURNING id",
            (producto.nombre, producto.descripcion, producto.precio, producto.stock)
        )
        producto.id = cursor.fetchone()[0]
        conn.commit()
        conn.close()
        return producto

    @staticmethod
    def actualizar(producto_id, producto: Producto):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute(
            "UPDATE productos SET nombre = %s, descripcion = %s, precio = %s, stock = %s WHERE id = %s",
            (producto.nombre, producto.descripcion, producto.precio, producto.stock, producto_id)
        )
        conn.commit()
        conn.close()

    @staticmethod
    def eliminar(producto_id):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute("DELETE FROM productos WHERE id = %s", (producto_id,))
        conn.commit()
        conn.close()
