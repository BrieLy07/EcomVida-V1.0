from app.db import obtener_conexion
from app.models.marca import Marca
from app.adapters.product_adapter import ProductAdapter

class MarcaRepository:

    @staticmethod
    def obtener_todas():
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute("SELECT id, nombre, descripcion FROM marcas")
        rows = cursor.fetchall()
        conn.close()
        return [Marca(*row) for row in rows]

    @staticmethod
    def obtener_por_id(id_marca):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute("SELECT id, nombre, descripcion FROM marcas WHERE id = %s", (id_marca,))
        row = cursor.fetchone()
        conn.close()
        return Marca(*row) if row else None

    @staticmethod
    def crear(marca: Marca):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute(
            "INSERT INTO marcas (nombre, descripcion) VALUES (%s, %s) RETURNING id",
            (marca.nombre, marca.descripcion)
        )
        marca.id = cursor.fetchone()[0]
        conn.commit()
        conn.close()
        return marca

    @staticmethod
    def actualizar(id_marca, marca: Marca):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute(
            "UPDATE marcas SET nombre = %s, descripcion = %s WHERE id = %s",
            (marca.nombre, marca.descripcion, id_marca)
        )
        conn.commit()
        conn.close()

    @staticmethod
    def eliminar(id_marca):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute("DELETE FROM marcas WHERE id = %s", (id_marca,))
        conn.commit()
        conn.close()

    @staticmethod
    def obtener_productos_por_marca(id_marca):
        conn = obtener_conexion()
        cursor = conn.cursor()
        cursor.execute(
            "SELECT id, nombre, descripcion, precio, stock FROM productos WHERE marca_id = %s",
            (id_marca,)
        )
        rows = cursor.fetchall()
        conn.close()
        return [ProductAdapter.transformar_fila(row) for row in rows]
