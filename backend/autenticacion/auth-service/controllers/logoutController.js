const logout = (req, res) => {
  // Aquí podrías invalidar el token con una blacklist si fuera necesario
  res.json({ mensaje: 'Sesión cerrada correctamente' });
};

module.exports = logout;
