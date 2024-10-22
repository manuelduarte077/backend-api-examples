package dev.donmanuel.Inventario.repository;

import dev.donmanuel.Inventario.model.Marca;
import org.springframework.data.jpa.repository.JpaRepository;

public interface MarcaRepository extends JpaRepository<Marca, Long> {
}
