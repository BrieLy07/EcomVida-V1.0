const direccionURL = "http://localhost:3003";

document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  const usuarioID = parseJwt(token).id;

  fetch(`${direccionURL}/address/${usuarioID}`, {
    headers: { Authorization: `Bearer ${token}` }
  })
    .then(res => res.json())
    .then(data => {
      if (data.calle) {
        document.getElementById("calle").value = data.calle;
        document.getElementById("ciudad").value = data.ciudad;
        document.getElementById("provincia").value = data.provincia;
        document.getElementById("codigo_postal").value = data.codigo_postal;
      }
    });
});

function guardarDireccion() {
  const token = localStorage.getItem("token");
  const usuarioID = parseJwt(token).id;

  const data = {
    usuario_id: usuarioID,
    calle: document.getElementById("calle").value,
    ciudad: document.getElementById("ciudad").value,
    provincia: document.getElementById("provincia").value,
    codigo_postal: document.getElementById("codigo_postal").value
  };

  fetch(`${direccionURL}/address`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify(data)
  }).then(() => alert("Dirección guardada"));
}

function parseJwt(token) {
  return JSON.parse(atob(token.split(".")[1]));
}
