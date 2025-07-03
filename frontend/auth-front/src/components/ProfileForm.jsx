import { useEffect, useState } from "react";
import { getProfile, updateProfile } from "../services/perfilService";

const ProfileForm = ({ userId }) => {
  const [form, setForm] = useState({
    nombre: "",
    apellido: "",
    usuario: "",
    correo: "",
    numero: ""
  });

  const [mensaje, setMensaje] = useState("");

  useEffect(() => {
    const cargarPerfil = async () => {
      try {
        const datos = await getProfile(userId);
        setForm(datos);
      } catch {
        setMensaje("Error al cargar el perfil.");
      }
    };
    cargarPerfil();
  }, [userId]);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await updateProfile(userId, form);
      setMensaje("✅ Perfil actualizado correctamente.");
    } catch {
      setMensaje("Error al actualizar el perfil.");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="bg-white p-6 rounded shadow w-full max-w-md">
      <h2 className="text-xl font-semibold mb-4">Mi Perfil</h2>

      <input
        name="nombre"
        value={form.nombre}
        onChange={handleChange}
        placeholder="Nombre"
        className="w-full border px-3 py-2 mb-3 rounded text-sm"
        required
      />
      <input
        name="apellido"
        value={form.apellido}
        onChange={handleChange}
        placeholder="Apellido"
        className="w-full border px-3 py-2 mb-3 rounded text-sm"
        required
      />
      <input
        name="usuario"
        value={form.usuario}
        onChange={handleChange}
        placeholder="Usuario"
        className="w-full border px-3 py-2 mb-3 rounded text-sm"
        required
      />
      <input
        name="correo"
        value={form.correo}
        onChange={handleChange}
        placeholder="Correo"
        className="w-full border px-3 py-2 mb-3 rounded text-sm"
        type="email"
        required
      />
      <input
        name="numero"
        value={form.numero}
        onChange={handleChange}
        placeholder="Número"
        className="w-full border px-3 py-2 mb-4 rounded text-sm"
        required
      />

      <button
        type="submit"
        className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 w-full"
      >
        Guardar cambios
      </button>

      {mensaje && <p className="mt-4 text-sm text-center">{mensaje}</p>}
    </form>
  );
};

export default ProfileForm;
