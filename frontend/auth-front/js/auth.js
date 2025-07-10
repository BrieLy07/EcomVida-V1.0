const baseURL = "http://localhost:3001";

function login() {
  const usuario = document.getElementById("usuario").value;
  const password = document.getElementById("password").value;

  fetch(`${baseURL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ usuario, password })
  })
    .then(res => res.json())
    .then(data => {
      if (data.token) {
        localStorage.setItem("token", data.token);
        location.href = "dashboard.html";
      } else {
        alert(data.error || "Error de autenticación");
      }
    });
}

function register() {
  const data = {
    nombre: document.getElementById("nombre").value,
    apellido: document.getElementById("apellido").value,
    usuario: document.getElementById("usuario").value,
    correo: document.getElementById("correo").value,
    numero: document.getElementById("numero").value,
    rol: document.getElementById("rol").value,
    password: document.getElementById("password").value
  };

  fetch(`${baseURL}/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data)
  })
    .then(res => res.json())
    .then(data => {
      alert("Registro exitoso, ahora inicia sesión");
      location.href = "login.html";
    });
}
