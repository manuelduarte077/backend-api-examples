package com.tasks.architectureAPI.domain.ports.in;

import com.tasks.architectureAPI.domain.models.Task;

import java.util.List;
import java.util.Optional;

public interface IRetrieveTaskUseCase {
    Optional<Task> getTaskById(Long id);

    List<Task> getAllTasks();

}
