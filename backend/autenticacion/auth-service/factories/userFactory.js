const bcrypt = require('bcryptjs');

class UserFactory {
  static async createUser(data) {
    const hashedPassword = await bcrypt.hash(data.password, 10);
    return {
      nombre: data.nombre,
      apellido: data.apellido,
      usuario: data.usuario,
      correo: data.correo,
      numero: data.numero,
      rol: data.rol,
      password: hashedPassword
    };
  }
}

module.exports = UserFactory;
