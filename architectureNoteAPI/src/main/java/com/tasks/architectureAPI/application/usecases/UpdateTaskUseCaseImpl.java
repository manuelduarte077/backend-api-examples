package com.tasks.architectureAPI.application.usecases;

import com.tasks.architectureAPI.domain.models.Task;
import com.tasks.architectureAPI.domain.ports.in.IUpdateTaskUseCase;
import com.tasks.architectureAPI.domain.ports.out.ITaskRepositoryPort;

import java.util.Optional;

public class UpdateTaskUseCaseImpl implements IUpdateTaskUseCase {
    private final ITaskRepositoryPort taskRepositoryPort;

    public UpdateTaskUseCaseImpl(ITaskRepositoryPort taskRepositoryPort) {
        this.taskRepositoryPort = taskRepositoryPort;
    }

    @Override
    public Optional<Task> updateTask(Long id, Task updatedTask) {
        return taskRepositoryPort.update(updatedTask);
    }
}
