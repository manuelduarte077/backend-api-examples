package dev.donmanuel.Inventario.repository;


import dev.donmanuel.Inventario.model.Producto;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ProductoRepository extends JpaRepository<Producto, Long> {
}
