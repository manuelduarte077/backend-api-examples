package com.tasks.architectureAPI.domain.ports.in;

import com.tasks.architectureAPI.domain.models.AdditionalTaskInfo;

public interface IGetAdditionalTaskInfoUseCase {
    AdditionalTaskInfo getAdditionalTaskInfo(Long taskId);
}
