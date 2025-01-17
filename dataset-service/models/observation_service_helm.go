package models

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/resource"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	os "github.com/caraml-dev/timber/observation-service/config"
)

// NewFluentdConfig converts FluentdConfig proto to Observation Service's FluentdConfig
func NewFluentdConfig(config *timberv1.FluentdConfig, projectName string) *os.FluentdConfig {
	return &os.FluentdConfig{
		Tag: config.GetTag(),
		// Set default values
		Host: fmt.Sprintf("observation-service-fluentd.%s.svc.cluster.local", projectName),
		Port: 24224,
		Kind: os.LoggerBQSinkFluentdProducer,
	}
}

// NewKafkaConfig converts KafkaConfig proto to Observation Service's KafkaConfig
func NewKafkaConfig(config *timberv1.KafkaConfig) *os.KafkaConfig {
	return &os.KafkaConfig{
		Brokers: config.GetBrokers(),
		Topic:   config.GetTopic(),
		// Set default values
		AutoOffsetReset:  "latest",
		CompressionType:  "none",
		ConnectTimeoutMS: 1000,
		MaxMessageBytes:  1048588,
		PollInterval:     1000,
	}
}

// Env represents configurable environment parameters for Observation Service deployment
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ConfigurableResources represents configurable resource parameters for Observation Service deployment
type ConfigurableResources struct {
	CPU    resource.Quantity `json:"cpu"`
	Memory resource.Quantity `json:"memory"`
}

// Resources represents multiple configurable resource parameters for Observation Service deployment
type Resources struct {
	Requests ConfigurableResources `json:"requests"`
	Limits   ConfigurableResources `json:"limits"`
}

// Autoscaling represents configurable autoscaling parameters for Observation Service deployment
type Autoscaling struct {
	Enabled                           bool `json:"enabled"`
	MinReplicas                       int  `json:"minReplicas"`
	MaxReplicas                       int  `json:"maxReplicas"`
	TargetCPUUtilizationPercentage    int  `json:"targetCPUUtilizationPercentage"`
	TargetMemoryUtilizationPercentage int  `json:"targetMemoryUtilizationPercentage"`
}

// ObservationServiceConfig is required in helm chart - observationService field
type ObservationServiceConfig struct {
	Image       Image       `json:"image"`
	ApiConfig   os.Config   `json:"apiConfig"`
	ExtraEnvs   []Env       `json:"extraEnvs"`
	Resources   Resources   `json:"resources"`
	Autoscaling Autoscaling `json:"autoscaling"`
}

// Image represents configurable image parameters for Observation Service deployment
type Image struct {
	Tag string `json:"tag"`
}

// FluentdConfig is required in helm chart - fluentd field
type FluentdConfig struct {
	Enabled     bool        `json:"enabled"`
	Image       Image       `json:"image"`
	ExtraEnvs   []Env       `json:"extraEnvs"`
	Resources   Resources   `json:"resources"`
	Autoscaling Autoscaling `json:"autoscaling"`
}

// ObservationServiceHelmValues is required in helm chart - observationService apiConfig field
type ObservationServiceHelmValues struct {
	ObservationServiceConfig ObservationServiceConfig `json:"observationService"`
	FluentdConfig            FluentdConfig            `json:"fluentd"`
}
