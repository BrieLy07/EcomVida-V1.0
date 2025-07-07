from app.db import obtener_conexion
from app.models.categoria import Categoria
from app.models.producto import Producto

class CategoriaRepository:

    @staticmethod
    def obtener_todas():
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute("SELECT id, nombre, descripcion FROM categorias")
        rows = cursor.fetchall()
        conn.close()
        return [Categoria(*row) for row in rows]

    @staticmethod
    def obtener_por_id(id_categoria):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute("SELECT id, nombre, descripcion FROM categorias WHERE id = %s", (id_categoria,))
        row = cursor.fetchone()
        conn.close()
        return Categoria(*row) if row else None

    @staticmethod
    def crear(categoria: Categoria):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute(
            "INSERT INTO categorias (nombre, descripcion) VALUES (%s, %s) RETURNING id",
            (categoria.nombre, categoria.descripcion)
        )
        categoria.id = cursor.fetchone()[0]
        conn.commit()
        conn.close()
        return categoria

    @staticmethod
    def actualizar(id_categoria, categoria: Categoria):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute(
            "UPDATE categorias SET nombre = %s, descripcion = %s WHERE id = %s",
            (categoria.nombre, categoria.descripcion, id_categoria)
        )
        conn.commit()
        conn.close()

    @staticmethod
    def eliminar(id_categoria):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute("DELETE FROM categorias WHERE id = %s", (id_categoria,))
        conn.commit()
        conn.close()

    @staticmethod
    def obtener_productos_por_categoria(id_categoria):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute(
            "SELECT id, nombre, descripcion, precio, stock FROM productos WHERE categoria_id = %s",
            (id_categoria,)
        )
        rows = cursor.fetchall()
        conn.close()
        return [Producto(*row) for row in rows]
