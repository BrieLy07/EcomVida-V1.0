import { useState } from "react";
import { login } from "../services/authService";
import jwt_decode from "jwt-decode"; // Asegúrate de instalar: npm i jwt-decode

const LoginForm = ({ onSuccess }) => {
  const [form, setForm] = useState({ correo: "", contraseña: "" });

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await login(form);
      const { token } = res.data;

      // Guardar token en localStorage
      localStorage.setItem("token", token);

      // Decodificar token para obtener el ID
      const decoded = jwt_decode(token);
      const id = decoded.id;

      // Pasar el ID hacia el componente principal
      onSuccess(id);
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
