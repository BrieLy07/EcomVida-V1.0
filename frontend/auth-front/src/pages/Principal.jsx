const Principal = ({ onLogout }) => {
  const cerrarSesion = () => {
    localStorage.removeItem("token");
    onLogout();
  };

  return (
    <div className="flex flex-col items-center mt-10">
      <h2 className="text-2xl font-bold mb-4">Bienvenido a tu panel principal</h2>
      <p className="text-gray-600 mb-6">Aquí se integrarán los microservicios del eCommerce.</p>
      <button
        onClick={cerrarSesion}
        className="bg-red-600 text-white px-5 py-2 rounded hover:bg-red-700"
      >
        Cerrar sesión
      </button>
    </div>
  );
};

export default Principal;
