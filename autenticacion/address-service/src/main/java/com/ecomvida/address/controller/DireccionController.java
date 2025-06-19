package com.ecomvida.address.controller;

import com.ecomvida.address.model.Direccion;
import com.ecomvida.address.repository.DireccionRepository;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/direcciones")
public class DireccionController {

    @Autowired
    private DireccionRepository direccionRepository;

    @GetMapping
    public List<Direccion> obtenerDirecciones(HttpServletRequest request) {
        String usuarioId = (String) request.getAttribute("usuarioId");
        return direccionRepository.findByUsuarioId(usuarioId);
    }

    @PostMapping
    public Direccion crearDireccion(@RequestBody Direccion direccion, HttpServletRequest request) {
        String usuarioId = (String) request.getAttribute("usuarioId");
        System.out.println("usuarioId recibido = " + usuarioId);
        direccion.setUsuarioId(usuarioId);
        return direccionRepository.save(direccion);
    }

    @PutMapping("/{id}")
    public Direccion actualizarDireccion(@PathVariable Long id,
                                         @RequestBody Direccion datos,
                                         HttpServletRequest request) {
        String usuarioId = (String) request.getAttribute("usuarioId");

        Direccion direccion = direccionRepository.findById(id)
                .filter(d -> d.getUsuarioId().equals(usuarioId))
                .orElseThrow();

        direccion.setCalle(datos.getCalle());
        direccion.setCiudad(datos.getCiudad());
        direccion.setProvincia(datos.getProvincia());
        direccion.setPais(datos.getPais());
        direccion.setCodigoPostal(datos.getCodigoPostal());

        return direccionRepository.save(direccion);
    }

    @DeleteMapping("/{id}")
    public void eliminarDireccion(@PathVariable Long id, HttpServletRequest request) {
        String usuarioId = (String) request.getAttribute("usuarioId");

        Direccion direccion = direccionRepository.findById(id)
                .filter(d -> d.getUsuarioId().equals(usuarioId))
                .orElseThrow();

        direccionRepository.delete(direccion);
    }
}
