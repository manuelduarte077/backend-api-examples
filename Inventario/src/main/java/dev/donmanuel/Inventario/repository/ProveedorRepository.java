package dev.donmanuel.Inventario.repository;

import dev.donmanuel.Inventario.model.Proveedor;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ProveedorRepository extends JpaRepository<Proveedor, Long> {
}
