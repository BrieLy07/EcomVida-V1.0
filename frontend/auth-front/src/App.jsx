import { useState } from "react";
import LoginForm from "./components/LoginForm";
import RegisterForm from "./components/RegisterForm";
import Principal from "./pages/Principal";
import Perfil from "./pages/Perfil";

function App() {
  const [view, setView] = useState("home");
  const [usuarioId, setUsuarioId] = useState(null); // 🆕 para guardar el ID

  const handleLoginSuccess = (id) => {
    setUsuarioId(id); // guardar el ID del usuario
    setView("principal");
  };

  const handleLogout = () => {
    localStorage.removeItem("token"); // Limpia el token
    setUsuarioId(null);               // Limpia el ID
    setView("home");
  };

  const handleIrPerfil = () => {
    setView("perfil");
  };

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-start p-10">
      <h1 className="text-4xl font-bold mb-8">Bienvenido a ECOMVIDA</h1>

      {view === "home" && (
        <div className="space-x-4">
          <button
            onClick={() => setView("login")}
            className="bg-blue-600 text-white px-6 py-2 rounded-lg shadow hover:bg-blue-700"
          >
            Iniciar sesión
          </button>
          <button
            onClick={() => setView("register")}
            className="bg-green-600 text-white px-6 py-2 rounded-lg shadow hover:bg-green-700"
          >
            Registrarse
          </button>
        </div>
      )}

      {view === "login" && (
        <>
          <LoginForm onSuccess={handleLoginSuccess} />
          <button
            onClick={() => setView("home")}
            className="mt-4 text-sm text-gray-600 underline"
          >
            Volver
          </button>
        </>
      )}

      {view === "register" && (
        <>
          <RegisterForm />
          <button
            onClick={() => setView("home")}
            className="mt-4 text-sm text-gray-600 underline"
          >
            Volver
          </button>
        </>
      )}

      {view === "principal" && (
        <Principal onLogout={handleLogout} onPerfil={handleIrPerfil} />
      )}

      {view === "perfil" && (
        <Perfil onVolver={() => setView("principal")} userId={usuarioId} />
      )}
    </div>
  );
}

export default App;
