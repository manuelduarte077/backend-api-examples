package com.tasks.architectureAPI.application.usecases;

import com.tasks.architectureAPI.domain.models.Task;
import com.tasks.architectureAPI.domain.ports.in.ICreateTaskUseCase;
import com.tasks.architectureAPI.domain.ports.out.ITaskRepositoryPort;

import java.util.List;

public class CreateTaskUseCaseImpl implements ICreateTaskUseCase {

    private final ITaskRepositoryPort taskRepositoryPort;

    public CreateTaskUseCaseImpl(ITaskRepositoryPort taskRepositoryPort) {
        this.taskRepositoryPort = taskRepositoryPort;
    }


    @Override
    public Task createTask(Task task) {
        return taskRepositoryPort.save(task);
    }

}
