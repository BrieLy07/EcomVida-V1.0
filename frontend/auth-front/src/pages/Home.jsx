import React from 'react';

const Home = ({ onLogout }) => {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-white p-4">
      <h2 className="text-2xl font-bold text-green-700 mb-4">¡Bienvenido a tu panel de ECOMVIDA!</h2>
      <p className="mb-6 text-gray-700 text-center">
        Aquí podrás acceder a todos los módulos del sistema.
      </p>

      {/* Aquí luego agregarás enlaces reales a los microservicios */}
      <div className="space-y-2">
        <button className="bg-blue-500 text-white px-4 py-2 rounded shadow">Productos</button>
        <button className="bg-blue-500 text-white px-4 py-2 rounded shadow">Pedidos</button>
        <button className="bg-blue-500 text-white px-4 py-2 rounded shadow">Perfil</button>
      </div>

      <button
        onClick={onLogout}
        className="mt-8 bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded"
      >
        Cerrar sesión
      </button>
    </div>
  );
};

export default Home;
