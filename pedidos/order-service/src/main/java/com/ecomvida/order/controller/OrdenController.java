package com.ecomvida.order.controller;

import com.ecomvida.order.model.Orden;
import com.ecomvida.order.repository.OrdenRepository;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/ordenes")
public class OrdenController {

    private final OrdenRepository repositorio;

    public OrdenController(OrdenRepository repositorio) {
        this.repositorio = repositorio;
    }

    @GetMapping
    public List<Orden> listar() {
        return repositorio.findAll();
    }

    @PostMapping
    public Orden crear(@RequestBody Orden orden) {
        return repositorio.save(orden);
    }
}
