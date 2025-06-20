package com.ecomvida.brand.service;

import com.ecomvida.brand.model.Marca;
import com.ecomvida.brand.repository.MarcaRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class MarcaService {
    private final MarcaRepository repository;

    public MarcaService(MarcaRepository repository) {
        this.repository = repository;
    }

    public Marca crearMarca(String nombre) {
        Marca marca = new Marca();
        marca.setNombre(nombre);
        return repository.save(marca);
    }

    public List<Marca> listarMarcas() {
        return repository.findAll();
    }

    public Optional<Marca> obtenerMarca(Long id) {
        return repository.findById(id);
    }
}
