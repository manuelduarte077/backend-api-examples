package com.tasks.architectureAPI.domain.ports.in;

import com.tasks.architectureAPI.domain.models.Task;

import java.util.Optional;

public interface IUpdateTaskUseCase {
    Optional<Task> updateTask(Long id, Task updatedTask);
}
