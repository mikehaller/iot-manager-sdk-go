/*
 * Bosch IoT Manager API
 *
 * The Bosch IoT Manager REST API provides the full functionality of the service, including:  &middot;    Retrieving devices and features; adding and modifying device-specific properties and attributes  &middot;    A simple device registration mechanism  &middot;    All grouping capabilities - listing, creating, modifying and deleting directories and tags  &middot;    The complete mass management experience, which allows full control over tasks and rules   Find out more details in our [documentation](https://docs.bosch-iot-suite.com/manager/).
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotmgr
// TaskStatusInfo struct for TaskStatusInfo
type TaskStatusInfo struct {
	// Task identifier
	TaskId string `json:"taskId,omitempty"`
	// Rule identifier
	RuleId string `json:"ruleId,omitempty"`
	// State of this Task
	State int32 `json:"state"`
	// Description of the Task state.
	StateDescription string `json:"stateDescription"`
	// Task launch time
	LaunchTime int64 `json:"launchTime"`
	// Task finish time
	FinishTime int64 `json:"finishTime"`
	// Count of involved devices with given execution status in this Task.
	InvolvedCount int32 `json:"involvedCount"`
	// Count of involved devices with FINISHED_SUCCESS(0) status
	SuccessCount int32 `json:"successCount,omitempty"`
	// Count of involved devices with FINISHED_WARNING(1) status
	WarningCount int32 `json:"warningCount,omitempty"`
	// Count of involved devices with FINISHED_ERROR(2) status
	ErrorCount int32 `json:"errorCount,omitempty"`
	// Count of involved devices with FINISHED_CANCELED(3) status
	CancelCount int32 `json:"cancelCount,omitempty"`
	IsPaused bool `json:"isPaused,omitempty"`
	// Time at which the execution will be resumed.
	NextResume int64 `json:"nextResume,omitempty"`
	// Time at which the execution will be paused.
	NextPause int64 `json:"nextPause,omitempty"`
	// User
	User string `json:"user,omitempty"`
	// Custom properties assigned to this Task.
	Props []string `json:"props,omitempty"`
	ScopeInfo ScopeInfo `json:"scopeInfo"`
	OptInfo ExecOptInfo `json:"optInfo"`
	Action ActionInfo `json:"action"`
	// A User-friendly display name of this Task.
	DisplayName string `json:"displayName,omitempty"`
	RuleInfo RuleInfo `json:"ruleInfo,omitempty"`
	//  Task execution is currently paused due to time-constraint restrictions.
	Paused bool `json:"paused,omitempty"`
}