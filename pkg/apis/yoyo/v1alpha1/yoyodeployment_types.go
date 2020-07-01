package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// YoyoDeploymentSpec defines the desired state of YoyoDeployment
// +k8s:openapi-gen=true
type YoyoDeploymentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Containers is a list of application containers being part of the deployment
	Containers []Container `json:"containers"`

	// RollbackPolicy defines how a rollback of the deployment should be executed
	// +optional
	RollbackPolicy RollbackPolicy `json:"rollbackPolicy,omitempty"`
}

// RollbackPolicy defines rollback policy for the deployment
// +k8s:openapi-gen=true
type RollbackPolicy string

const (
	// Automatic defines automatic rollback policy; the rollback will be executed automatically after startup failure is detected
	Automatic RollbackPolicy = "Automatic"
	// Disabled defines no rollback policy; even if startup fails, no action will be undertaken
	Disabled RollbackPolicy = "Disabled"
)

// Container defines  application container that will be run as part of the deployment
// +k8s:openapi-gen=true
type Container struct {
	// Name specifies the name of the container
	Name string `json:"name"`

	// Image defines the container image that will be started
	Image string `json:"image"`

	// Ports specifies list of ports to expose from the container
	// +optional
	Ports []ContainerPort `json:"ports,omitempty"`

	// HealtCheck defines how health of the container is verified
	// +optional
	HealthCheck corev1.Probe `json:"healthCheck,omitempty"`
}

// ContainerPort defines network port
// +k8s:openapi-gen=true
type ContainerPort struct {
	// PortNumber is the number of the network port
	PortNumber int32 `json:"portNumber"`
	// Protocol specifies protocol for which the port is defined
	Protocol corev1.Protocol `json:"protocol"`
	// ServicePort specifies port number that should be used to map to this container port
	ServicePort int32 `json:"servicePort"`
}

// YoyoDeploymentStatus defines the observed state of a Yoyo Deployment
type YoyoDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// +optional
	// DeploymentName reports the name of dependant k8s Deployment
	DeploymentName string `json:"deploymentName"`

	// +optional
	// Current service state of the deployment.
	Conditions []YoyoDeploymentCondition `json:"conditions"`
}

// YoyoDeploymentCondition defines the observed state of YoyoDeploymentCondition conditions
// +k8s:openapi-gen=true
type YoyoDeploymentCondition struct {
	// Type of virtual machine import condition
	Type YoyoDeploymentConditionType `json:"type"`

	// Status of the condition, one of True, False, Unknown
	Status corev1.ConditionStatus `json:"status"`

	// A brief CamelCase string that describes why the Yoyo Deployment is in current condition status
	// +optional
	Reason string `json:"reason,omitempty"`

	// A human-readable message indicating details about last transition
	// +optional
	Message string `json:"message,omitempty"`

	// The last time we got an update on a given condition
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`

	// The last time the condition transit from one status to another
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

// YoyoDeploymentConditionType represents a yoyo deployment condition type
// +k8s:openapi-gen=true
type YoyoDeploymentConditionType string

const (
	// Deploying represents the status of the deployment being deployed
	Deploying YoyoDeploymentConditionType = "Deploying"

	// Deployed represents the status of a deployed deployment
	Deployed YoyoDeploymentConditionType = "Deployed"

	// Degraded represents the status of deployment that is not fully operational; some components are unhealthy
	Degraded YoyoDeploymentConditionType = "Degraded"

	// RollingBack represents the status of a deployment being rolled-back
	RollingBack YoyoDeploymentConditionType = "RollingBack"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// YoyoDeployment is the Schema for the yoyodeployments API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=yoyodeployments,scope=Namespaced
type YoyoDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   YoyoDeploymentSpec   `json:"spec,omitempty"`
	Status YoyoDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// YoyoDeploymentList contains a list of YoyoDeployment
type YoyoDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []YoyoDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&YoyoDeployment{}, &YoyoDeploymentList{})
}
