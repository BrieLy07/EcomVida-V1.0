import { useEffect, useState } from "react";
import { obtenerPerfil, actualizarPerfil } from "../services/profileService";
import {
  obtenerDirecciones,
  crearDireccion,
  eliminarDireccion,
} from "../services/direccionService";

const Perfil = ({ onVolver, userId }) => {
  const token = localStorage.getItem("token");

  const [modoEditar, setModoEditar] = useState(false);
  const [perfil, setPerfil] = useState(null);
  const [direcciones, setDirecciones] = useState([]);
  const [nuevaDireccion, setNuevaDireccion] = useState({
    direccion: "",
    ciudad: "",
    provincia: "",
    pais: "",
    codigo_postal: "",
    telefono: "",
  });

  const [form, setForm] = useState({
    nombre: "",
    apellido: "",
    usuario: "",
    correo: "",
    numero: "",
  });

  useEffect(() => {
    const cargarDatos = async () => {
      try {
        const datos = await obtenerPerfil(userId);
        setPerfil(datos);
        setForm(datos);

        const direccionesData = await obtenerDirecciones(userId, token);
        if (Array.isArray(direccionesData)) {
          setDirecciones(direccionesData);
        } else {
          console.warn("La respuesta de direcciones no es un array:", direccionesData);
          setDirecciones([]);
        }
      } catch (error) {
        console.error("Error al cargar datos:", error);
        setDirecciones([]);
      }
    };

    if (userId && token) {
      cargarDatos();
    }
  }, [userId, token]);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleGuardar = async () => {
    await actualizarPerfil(userId, form);
    setPerfil(form);
    setModoEditar(false);
  };

  const handleNuevaDireccion = (e) => {
    setNuevaDireccion({ ...nuevaDireccion, [e.target.name]: e.target.value });
  };

  const handleCrearDireccion = async () => {
    const ok = await crearDireccion(userId, nuevaDireccion, token);
    if (ok) {
      const actualizadas = await obtenerDirecciones(userId, token);
      setDirecciones(Array.isArray(actualizadas) ? actualizadas : []);
      setNuevaDireccion({
        direccion: "",
        ciudad: "",
        provincia: "",
        pais: "",
        codigo_postal: "",
        telefono: "",
      });
    }
  };

  const handleEliminarDireccion = async (id) => {
    const ok = await eliminarDireccion(id, token);
    if (ok) {
      setDirecciones(direcciones.filter((d) => d.id !== id));
    }
  };

  if (!perfil) return <p>Cargando perfil...</p>;

  return (
    <div className="bg-white shadow p-6 rounded-lg w-full max-w-2xl">
      <h2 className="text-xl font-bold mb-4">Mi Perfil</h2>

      {!modoEditar ? (
        <>
          <p><strong>Nombre:</strong> {perfil.nombre}</p>
          <p><strong>Apellido:</strong> {perfil.apellido}</p>
          <p><strong>Usuario:</strong> {perfil.usuario}</p>
          <p><strong>Correo:</strong> {perfil.correo}</p>
          <p><strong>Número:</strong> {perfil.numero}</p>

          <div className="mt-4 space-x-2">
            <button onClick={() => setModoEditar(true)} className="bg-yellow-500 text-white px-4 py-1 rounded">
              Editar
            </button>
            <button onClick={onVolver} className="bg-gray-400 text-white px-4 py-1 rounded">
              Volver
            </button>
          </div>
        </>
      ) : (
        <>
          <input name="nombre" value={form.nombre} onChange={handleChange} className="input" />
          <input name="apellido" value={form.apellido} onChange={handleChange} className="input" />
          <input name="usuario" value={form.usuario} onChange={handleChange} className="input" />
          <input name="correo" value={form.correo} onChange={handleChange} className="input" />
          <input name="numero" value={form.numero} onChange={handleChange} className="input" />

          <div className="mt-4 space-x-2">
            <button onClick={handleGuardar} className="bg-blue-600 text-white px-4 py-1 rounded">
              Guardar
            </button>
            <button onClick={() => setModoEditar(false)} className="bg-gray-400 text-white px-4 py-1 rounded">
              Cancelar
            </button>
          </div>
        </>
      )}

      <hr className="my-6" />
      <h3 className="text-lg font-semibold mb-2">Direcciones</h3>
      {direcciones.length === 0 ? (
        <p>No tienes direcciones registradas.</p>
      ) : (
        <ul className="space-y-2">
          {direcciones.map((dir) => (
            <li key={dir.id} className="border p-2 rounded">
              <p>{dir.direccion}, {dir.ciudad}, {dir.provincia}</p>
              <p>{dir.pais}, CP: {dir.codigo_postal}</p>
              <p>Tel: {dir.telefono}</p>
              <button onClick={() => handleEliminarDireccion(dir.id)} className="mt-1 bg-red-500 text-white px-2 py-1 text-sm rounded">
                Eliminar
              </button>
            </li>
          ))}
        </ul>
      )}

      <h4 className="mt-6 font-medium">Agregar nueva dirección</h4>
      <div className="grid grid-cols-1 gap-2">
        <input name="direccion" placeholder="Dirección" value={nuevaDireccion.direccion} onChange={handleNuevaDireccion} className="input" />
        <input name="ciudad" placeholder="Ciudad" value={nuevaDireccion.ciudad} onChange={handleNuevaDireccion} className="input" />
        <input name="provincia" placeholder="Provincia" value={nuevaDireccion.provincia} onChange={handleNuevaDireccion} className="input" />
        <input name="pais" placeholder="País" value={nuevaDireccion.pais} onChange={handleNuevaDireccion} className="input" />
        <input name="codigo_postal" placeholder="Código Postal" value={nuevaDireccion.codigo_postal} onChange={handleNuevaDireccion} className="input" />
        <input name="telefono" placeholder="Teléfono" value={nuevaDireccion.telefono} onChange={handleNuevaDireccion} className="input" />
        <button onClick={handleCrearDireccion} className="bg-green-600 text-white px-4 py-1 rounded">
          Guardar dirección
        </button>
      </div>
    </div>
  );
};

export default Perfil;
