package com.tasks.architectureAPI.application.usecases;

import com.tasks.architectureAPI.domain.models.Task;
import com.tasks.architectureAPI.domain.ports.in.IRetrieveTaskUseCase;
import com.tasks.architectureAPI.domain.ports.out.ITaskRepositoryPort;

import java.util.List;
import java.util.Optional;

public class RetrieveTaskUseCaseImpl implements IRetrieveTaskUseCase {

    private final ITaskRepositoryPort taskRepositoryPort;

    public RetrieveTaskUseCaseImpl(ITaskRepositoryPort taskRepositoryPort) {
        this.taskRepositoryPort = taskRepositoryPort;
    }

    @Override
    public Optional<Task> getTaskById(Long id) {
        return taskRepositoryPort.findById(id);
    }

    @Override
    public List<Task> getAllTasks() {
        return taskRepositoryPort.findAll();
    }
}
