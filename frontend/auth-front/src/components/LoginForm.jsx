import { useState } from "react";
import { login } from "../services/authService";

const LoginForm = ({ onSuccess }) => {
  const [form, setForm] = useState({ correo: "", contraseña: "" });

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await login(form);
      const { token, id } = res.data;

      localStorage.setItem("token", token);  // almacena token
      onSuccess(id);                         // pasa el ID al App.jsx
    } catch {
      alert("Credenciales inválidas");
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="flex flex-col items-center gap-4 bg-transparent p-4"
    >
      <input
        name="correo"
        type="email"
        placeholder="Correo"
        onChange={handleChange}
        required
        className="px-3 py-1.5 text-sm border border-gray-300 rounded w-72"
      />
      <input
        name="contraseña"
        type="password"
        placeholder="Contraseña"
        onChange={handleChange}
        required
        className="px-3 py-1.5 text-sm border border-gray-300 rounded w-72"
      />
      <button
        type="submit"
        className="bg-blue-600 text-white px-6 py-1.5 rounded hover:bg-blue-700"
      >
        Iniciar sesión
      </button>
    </form>
  );
};

export default LoginForm;
