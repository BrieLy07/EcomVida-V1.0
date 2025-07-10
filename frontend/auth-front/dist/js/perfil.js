const perfilURL = "http://localhost:3002";

document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  const usuarioID = parseJwt(token).id;

  fetch(`${perfilURL}/users/${usuarioID}`, {
    headers: { Authorization: `Bearer ${token}` }
  })
    .then(res => res.json())
    .then(data => {
      document.getElementById("nombre").value = data.nombre;
      document.getElementById("apellido").value = data.apellido;
      document.getElementById("correo").value = data.correo;
      document.getElementById("numero").value = data.numero;
    });
});

function guardarPerfil() {
  const token = localStorage.getItem("token");
  const usuarioID = parseJwt(token).id;

  const data = {
    nombre: document.getElementById("nombre").value,
    apellido: document.getElementById("apellido").value,
    correo: document.getElementById("correo").value,
    numero: document.getElementById("numero").value
  };

  fetch(`http://localhost:3002/users/${usuarioID}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify(data)
  }).then(() => alert("Perfil actualizado"));
}

function parseJwt(token) {
  return JSON.parse(atob(token.split(".")[1]));
}
