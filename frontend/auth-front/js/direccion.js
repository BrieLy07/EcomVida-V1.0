const direccionURL = "http://localhost:3003";

document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  const usuarioID = parseJwt(token).id;

  fetch(`${direccionURL}/address/me`, {
    headers: { Authorization: `Bearer ${token}` }
  })
    .then(res => res.json())
    .then(data => {
      if (data.calle) {
        const texto = `${data.calle}, ${data.ciudad}, ${data.provincia}, CP ${data.codigo_postal}`;
        document.getElementById("direccionTexto").innerText = texto;

        document.getElementById("calle").value = data.calle;
        document.getElementById("ciudad").value = data.ciudad;
        document.getElementById("provincia").value = data.provincia;
        document.getElementById("codigo_postal").value = data.codigo_postal;
      } else {
        document.getElementById("direccionTexto").innerText = "No hay dirección registrada.";
      }
    });
});

function mostrarFormularioDireccion() {
  document.getElementById("direccion-info").style.display = "none";
  document.getElementById("direccion-formulario").style.display = "block";
}

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
  }).then(() => location.reload());
}

function eliminarDireccion() {
  const token = localStorage.getItem("token");
  const usuarioID = parseJwt(token).id;

  fetch(`${direccionURL}/address/me`, {
    method: "DELETE",
    headers: { Authorization: `Bearer ${token}` }
  }).then(() => location.reload());
}

function parseJwt(token) {
  return JSON.parse(atob(token.split(".")[1]));
}
