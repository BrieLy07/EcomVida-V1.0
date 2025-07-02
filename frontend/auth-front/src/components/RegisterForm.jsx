import { useState } from "react";
import { register } from "../services/authService";

const RegisterForm = () => {
  const [form, setForm] = useState({
    nombre: "",
    apellido: "",
    usuario: "",
    correo: "",
    numero: "",
    contraseña: "",
  });

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await register(form);
      alert("Registro exitoso");
    } catch {
      alert("Error al registrar");
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="flex flex-col items-center gap-4 bg-transparent p-4"
    >
      {["nombre", "apellido", "usuario", "correo", "numero", "contraseña"].map(
        (field) => (
          <input
            key={field}
            name={field}
            type={field === "contraseña" ? "password" : "text"}
            placeholder={field.charAt(0).toUpperCase() + field.slice(1)}
            onChange={handleChange}
            required
            className="px-3 py-1.5 text-sm border border-gray-300 rounded w-72"
          />
        )
      )}
      <button
        type="submit"
        className="bg-green-600 text-white px-6 py-1.5 rounded hover:bg-green-700"
      >
        Registrarse
      </button>
    </form>
  );
};

export default RegisterForm;
