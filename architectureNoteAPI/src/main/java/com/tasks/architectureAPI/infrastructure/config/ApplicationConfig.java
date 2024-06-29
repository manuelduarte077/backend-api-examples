package com.tasks.architectureAPI.infrastructure.config;

import com.tasks.architectureAPI.application.services.TaskService;
import com.tasks.architectureAPI.application.usecases.*;
import com.tasks.architectureAPI.domain.ports.in.IGetAdditionalTaskInfoUseCase;
import com.tasks.architectureAPI.domain.ports.out.IExternalServicePort;
import com.tasks.architectureAPI.domain.ports.out.ITaskRepositoryPort;
import com.tasks.architectureAPI.infrastructure.adapters.ExternalServiceAdapter;
import com.tasks.architectureAPI.infrastructure.repositories.JpaTaskRepositoryAdapter;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class ApplicationConfig {

    @Bean
    public TaskService taskService(ITaskRepositoryPort taskRepositoryPort, IGetAdditionalTaskInfoUseCase getAdditionalTaskInfoUseCase) {
        return new TaskService(
                new CreateTaskUseCaseImpl(taskRepositoryPort),
                new RetrieveTaskUseCaseImpl(taskRepositoryPort),
                new UpdateTaskUseCaseImpl(taskRepositoryPort),
                new DeleteTaskUseCaseImpl(taskRepositoryPort),

                getAdditionalTaskInfoUseCase
        );
    }


    @Bean
    public ITaskRepositoryPort taskRepositoryPort(JpaTaskRepositoryAdapter jpaTaskRepositoryAdapter) {
        return jpaTaskRepositoryAdapter;
    }


    @Bean
    public IGetAdditionalTaskInfoUseCase getAdditionalTaskInfoUseCase(IExternalServicePort externalServicePort) {
        return new GetAdditionalTaskInfoUseCaseImpl(externalServicePort);
    }

    @Bean
    public IExternalServicePort externalServicePort() {
        return new ExternalServiceAdapter();
    }
}
