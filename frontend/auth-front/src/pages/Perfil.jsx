import { useEffect, useState } from "react";
import { obtenerPerfil, actualizarPerfil } from "../services/profileService";

const Perfil = ({ onVolver }) => {
  const userId = "6866eac444d2a028f0877a21";
  const [modoEditar, setModoEditar] = useState(false);
  const [perfil, setPerfil] = useState(null);
  const [form, setForm] = useState({
    nombre: "",
    apellido: "",
    usuario: "",
    correo: "",
    numero: "",
  });

useEffect(() => {
  const cargarPerfil = async () => {
    try {
      console.log("Consultando perfil con ID:", userId);
      const datos = await obtenerPerfil(userId);
      console.log("Datos recibidos:", datos);
      setPerfil(datos);
      setForm(datos);
    } catch (error) {
      console.error("Error al obtener perfil:", error);
    }
  };
  cargarPerfil();
}, []);


  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleGuardar = async () => {
    await actualizarPerfil(userId, form);
    setPerfil(form);
    setModoEditar(false);
  };

  if (!perfil) return <p>Cargando perfil...</p>;

  return (
    <div className="bg-white shadow p-6 rounded-lg w-full max-w-md">
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
    </div>
  );
};

export default Perfil;
