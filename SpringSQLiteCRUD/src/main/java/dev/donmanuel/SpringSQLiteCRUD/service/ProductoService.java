package dev.donmanuel.SpringSQLiteCRUD.service;

import dev.donmanuel.SpringSQLiteCRUD.model.Producto;
import dev.donmanuel.SpringSQLiteCRUD.repository.ProductoRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class ProductoService {

    @Autowired
    private ProductoRepository productoRepository;

    public List<Producto> obtenerTodos() {
        return productoRepository.findAll();
    }

    public Optional<Producto> obtenerPorId(Long id) {
        return productoRepository.findById(id);
    }

    public Producto guardar(Producto producto) {
        return productoRepository.save(producto);
    }

    public void eliminarPorId(Long id) {
        productoRepository.deleteById(id);
    }

}