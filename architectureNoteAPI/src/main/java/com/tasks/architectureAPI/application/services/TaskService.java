package com.tasks.architectureAPI.application.services;


import com.tasks.architectureAPI.domain.models.AdditionalTaskInfo;
import com.tasks.architectureAPI.domain.models.Task;
import com.tasks.architectureAPI.domain.ports.in.*;

import java.util.List;
import java.util.Optional;

public class TaskService implements ICreateTaskUseCase, IDeleteTaskUseCase, IGetAdditionalTaskInfoUseCase, IRetrieveTaskUseCase, IUpdateTaskUseCase {

    private final ICreateTaskUseCase createTaskUseCase;
    private final IRetrieveTaskUseCase retrieveTaskUseCase;
    private final IUpdateTaskUseCase updateTaskUseCase;
    private final IDeleteTaskUseCase deleteTaskUseCase;
    private final IGetAdditionalTaskInfoUseCase getAdditionalTaskInfoUseCase;

    public TaskService(ICreateTaskUseCase createTaskUseCase, IRetrieveTaskUseCase retrieveTaskUseCase, IUpdateTaskUseCase updateTaskUseCase, IDeleteTaskUseCase deleteTaskUseCase, IGetAdditionalTaskInfoUseCase getAdditionalTaskInfoUseCase) {
        this.createTaskUseCase = createTaskUseCase;
        this.retrieveTaskUseCase = retrieveTaskUseCase;
        this.updateTaskUseCase = updateTaskUseCase;
        this.deleteTaskUseCase = deleteTaskUseCase;
        this.getAdditionalTaskInfoUseCase = getAdditionalTaskInfoUseCase;
    }

    @Override
    public Task createTask(Task task) {
        return createTaskUseCase.createTask(task);
    }

    @Override
    public Optional<Task> getTaskById(Long id) {
        return retrieveTaskUseCase.getTaskById(id);
    }

    @Override
    public List<Task> getAllTasks() {
        return retrieveTaskUseCase.getAllTasks();
    }

    @Override
    public Optional<Task> updateTask(Long id, Task updatedTask) {
        return updateTaskUseCase.updateTask(id, updatedTask);
    }

    @Override
    public boolean deleteTask(Long id) {
        return deleteTaskUseCase.deleteTask(id);
    }

    @Override
    public AdditionalTaskInfo getAdditionalTaskInfo(Long taskId) {
        return getAdditionalTaskInfoUseCase.getAdditionalTaskInfo(taskId);
    }
}
