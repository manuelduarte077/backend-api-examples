package com.tasks.architectureAPI.application.usecases;

import com.tasks.architectureAPI.domain.models.AdditionalTaskInfo;
import com.tasks.architectureAPI.domain.ports.in.IGetAdditionalTaskInfoUseCase;
import com.tasks.architectureAPI.domain.ports.out.IExternalServicePort;

public class GetAdditionalTaskInfoUseCaseImpl implements IGetAdditionalTaskInfoUseCase {
    private final IExternalServicePort externalServicePort;

    public GetAdditionalTaskInfoUseCaseImpl(IExternalServicePort externalServicePort) {
        this.externalServicePort = externalServicePort;
    }

    @Override
    public AdditionalTaskInfo getAdditionalTaskInfo(Long id) {
        return externalServicePort.getAdditionalTaskInfo(id);
    }
}
