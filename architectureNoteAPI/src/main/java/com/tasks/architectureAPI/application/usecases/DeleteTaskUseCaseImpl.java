package com.tasks.architectureAPI.application.usecases;

import com.tasks.architectureAPI.domain.ports.in.IDeleteTaskUseCase;
import com.tasks.architectureAPI.domain.ports.out.ITaskRepositoryPort;

public class DeleteTaskUseCaseImpl implements IDeleteTaskUseCase {

    private final ITaskRepositoryPort taskRepositoryPort;

    public DeleteTaskUseCaseImpl(ITaskRepositoryPort taskRepositoryPort) {
        this.taskRepositoryPort = taskRepositoryPort;
    }

    @Override
    public boolean deleteTask(Long id) {
        return taskRepositoryPort.deleteById(id);
    }
}
