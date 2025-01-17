from datetime import datetime

import pytest
import yaml
from data import ObservationServiceDataSinkType
from kubernetes import client, config

from tests.e2e.dataset_service_client import (
    TEST_PROJECT_ID,
    TEST_PROJECT_NAME,
    TIMEOUT_SECONDS,
    DatasetServiceClient,
)


@pytest.fixture()
def kubernetes_api_instance() -> client.CoreV1Api:
    # Generated by docker-compose kubeconfig helper container
    config_file = "/tmp/kubeconfig"
    with open(config_file, "r") as stream:
        try:
            yaml.safe_load(stream)
        except yaml.YAMLError:
            raise
    config.load_kube_config(config_file)
    api_instance = client.CoreV1Api()

    return api_instance


@pytest.mark.order(1)
def test_simple_observation_service_creation(
    dataset_service_client: DatasetServiceClient,
    kubernetes_api_instance: client.CoreV1Api,
):
    # Create Observation Service
    service_name = "test-service"
    req_body = {"observation_service": {"service_name": service_name}}
    observation_service = dataset_service_client.create_observation_service(
        TEST_PROJECT_ID, req_body
    )
    assert "id" in observation_service

    # Check pods for deployed Observation Service
    pod_name_prefix = f"observation-service-{service_name}"
    pod_ready_status = False
    pod_status = ""
    start = datetime.now().replace(microsecond=0)
    end = datetime.now().replace(microsecond=0)
    elapsed_duration = end - start
    while not pod_ready_status and elapsed_duration.total_seconds() < TIMEOUT_SECONDS:
        end = datetime.now().replace(microsecond=0)
        elapsed_duration = end - start

        pods: client.models.v1_pod_list.V1PodList = (
            kubernetes_api_instance.list_namespaced_pod(TEST_PROJECT_NAME)
        )
        for pod in pods.items:
            if pod_name_prefix in pod.metadata.name:
                if pod.status.conditions:
                    for condition in pod.status.conditions:
                        if condition.type == "Ready" and condition.status == "True":
                            pod_ready_status = True
                pod_status = pod.status.phase

    assert pod_status == "Running"
    assert pod_ready_status is True


@pytest.mark.order(2)
def test_simple_observation_service_updation(
    dataset_service_client: DatasetServiceClient,
    kubernetes_api_instance: client.CoreV1Api,
):
    # Update Observation Service
    service_name = "test-service"
    req_body = {
        "observation_service": {
            "service_name": service_name,
            "sink": {
                "type": ObservationServiceDataSinkType.OBSERVATION_SERVICE_DATA_SINK_TYPE_STDOUT.value
            },
        }
    }
    observation_service = dataset_service_client.update_observation_service(
        TEST_PROJECT_ID, 1, req_body
    )
    assert "id" in observation_service

    # Check secrets for updated Observation Service
    new_secret_present = False
    secrets: client.models.v1_secret_list.V1SecretList = (
        kubernetes_api_instance.list_namespaced_secret(TEST_PROJECT_NAME)
    )
    for secret in secrets.items:
        if "v2" in secret.metadata.name:
            new_secret_present = True

    assert new_secret_present is True
