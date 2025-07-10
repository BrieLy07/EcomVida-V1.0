const perfilURL = "http://localhost:3002";

document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  const usuarioID = parseJwt(token).id;

  fetch(`${perfilURL}/users/me`, {
    headers: { Authorization: `Bearer ${token}` }
  })
    .then(res => res.json())
    .then(data => {
      document.getElementById("nombre").innerText = data.nombre;
      document.getElementById("apellido").innerText = data.apellido;
      document.getElementById("correo").innerText = data.correo;
      document.getElementById("numero").innerText = data.numero;

      document.getElementById("nombreInput").value = data.nombre;
      document.getElementById("apellidoInput").value = data.apellido;
      document.getElementById("correoInput").value = data.correo;
      document.getElementById("numeroInput").value = data.numero;
    });
});

function mostrarFormularioPerfil() {
  document.getElementById("perfil-info").style.display = "none";
  document.getElementById("perfil-formulario").style.display = "block";
}

function guardarPerfil() {
  const token = localStorage.getItem("token");
  const usuarioID = parseJwt(token).id;

  const data = {
    nombre: document.getElementById("nombreInput").value,
    apellido: document.getElementById("apellidoInput").value,
    correo: document.getElementById("correoInput").value,
    numero: document.getElementById("numeroInput").value
  };

  fetch(`${perfilURL}/users/me`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify(data)
  }).then(() => location.reload());
}

function parseJwt(token) {
  return JSON.parse(atob(token.split(".")[1]));
}
