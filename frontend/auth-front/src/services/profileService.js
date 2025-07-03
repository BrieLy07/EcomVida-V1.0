import axios from "axios";

const API_URL = import.meta.env.VITE_USER_PROFILE_API + "/users";


export const obtenerPerfil = async (id) => {
  const res = await axios.get(`${API_URL}/${id}`);
  return res.data;
};

export const actualizarPerfil = async (id, datos) => {
  const res = await axios.put(`${API_URL}/${id}`, datos);
  return res.data;
};
