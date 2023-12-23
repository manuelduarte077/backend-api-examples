package com.tasks.architectureAPI.domain.ports.out;

import com.tasks.architectureAPI.domain.models.AdditionalTaskInfo;

public interface IExternalServicePort {
    AdditionalTaskInfo getAdditionalTaskInfo(Long taskId);
}
