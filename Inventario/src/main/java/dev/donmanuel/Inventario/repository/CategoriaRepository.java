package dev.donmanuel.Inventario.repository;

import dev.donmanuel.Inventario.model.Categoria;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CategoriaRepository extends JpaRepository<Categoria, Long> {
}
