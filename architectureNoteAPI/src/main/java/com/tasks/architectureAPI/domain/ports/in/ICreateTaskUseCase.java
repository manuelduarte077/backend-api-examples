package com.tasks.architectureAPI.domain.ports.in;

import com.tasks.architectureAPI.domain.models.Task;

public interface ICreateTaskUseCase {
    Task createTask(Task task);
}
