package com.tasks.architectureAPI.domain.ports.out;

import com.tasks.architectureAPI.domain.models.Task;

import java.util.List;
import java.util.Optional;

public interface ITaskRepositoryPort {
    Task save(Task task);

    Optional<Task> findById(Long id);

    List<Task> findAll();

    Optional<Task> update(Task task);

    boolean deleteById(Long id);
}


