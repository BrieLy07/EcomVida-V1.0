const BASE_URL = import.meta.env.VITE_ADDRESS_SERVICE_URL;

export const obtenerDirecciones = async (usuarioId, token) => {
  const res = await fetch(`${BASE_URL}/users/${usuarioId}/addresses`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return await res.json();
};

export const crearDireccion = async (usuarioId, direccion, token) => {
  const res = await fetch(`${BASE_URL}/users/${usuarioId}/addresses`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(direccion),
  });
  return res.ok;
};

export const actualizarDireccion = async (direccionId, direccion, token) => {
  const res = await fetch(`${BASE_URL}/addresses/${direccionId}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(direccion),
  });
  return res.ok;
};

export const eliminarDireccion = async (direccionId, token) => {
  const res = await fetch(`${BASE_URL}/addresses/${direccionId}`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.ok;
};
