package com.ecomvida.brand.ws;

import jakarta.jws.WebMethod;
import jakarta.jws.WebParam;
import jakarta.jws.WebService;
import com.ecomvida.brand.service.MarcaService;
import com.ecomvida.brand.model.Marca;

import java.util.List;

@WebService
public class MarcaEndpoint {

    private final MarcaService marcaService;

    public MarcaEndpoint(MarcaService marcaService) {
        this.marcaService = marcaService;
    }

    @WebMethod
    public Marca crearMarca(@WebParam(name = "nombre") String nombre) {
        return marcaService.crearMarca(nombre);
    }

    @WebMethod
    public List<Marca> listarMarcas() {
        return marcaService.listarMarcas();
    }

    @WebMethod
    public Marca obtenerMarca(@WebParam(name = "id") Long id) {
        return marcaService.obtenerMarca(id).orElse(null);
    }
}
