const Principal = ({ onLogout, onPerfil }) => {
  const cerrarSesion = () => {
    localStorage.removeItem("token");
    onLogout();
  };

  return (
    <div className="flex flex-col items-center mt-10 space-y-6">
      <h2 className="text-2xl font-bold">Bienvenido a tu panel principal</h2>
      <p className="text-gray-600">Aquí se integrarán los microservicios del eCommerce.</p>

      <div className="space-x-4">
        <button
          onClick={onPerfil}
          className="bg-green-600 text-white px-5 py-2 rounded hover:bg-green-700"
        >
          Ir al Perfil
        </button>

        <button
          onClick={cerrarSesion}
          className="bg-red-600 text-white px-5 py-2 rounded hover:bg-red-700"
        >
          Cerrar sesión
        </button>
      </div>
    </div>
  );
};

export default Principal;
