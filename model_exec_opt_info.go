/*
 * Bosch IoT Manager API
 *
 * The Bosch IoT Manager REST API provides the full functionality of the service, including:  &middot;    Retrieving devices and features; adding and modifying device-specific properties and attributes  &middot;    A simple device registration mechanism  &middot;    All grouping capabilities - listing, creating, modifying and deleting directories and tags  &middot;    The complete mass management experience, which allows full control over tasks and rules   Find out more details in our [documentation](https://docs.bosch-iot-suite.com/manager/).
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotmgr
// ExecOptInfo Execution options of this Rule
type ExecOptInfo struct {
	//   Defines whether multiple Tasks over the same device can overlap or should be replaced. This option is applicable only for Rules. Overlapping: <li> ALLOW_OVERLAPPING_TASKS(0): No relation constraint between Tasks. New Tasks can be triggered, although old Tasks are still active. This option should be used very carefully due to possible accumulation (explosion) of unfinished Tasks. </li> <li>NO_OVERLAPPING_CANCEL_OLD(1) : In case of Rule firing for a particular device, any previous executions still running for this device in the scope of the same Rule will be canceled. </li><li> NO_OVERLAPPING_KEEP_OLD_SKIP_NEW(2): Skip triggering a new Task for devices that already have unfinished Tasks in the scope of the Rule.;
	OverlapOption int32 `json:"overlapOption,omitempty"`
	// Restricts the maximum number of devices concurrently executing a Task or Rule.
	ConcurrencyLimit int32 `json:"concurrencyLimit"`
	// Timeout in seconds after which a non-confirmed execution expires and stops occupying room in the concurrency limit. The default value is 30 seconds.
	ConcurrencyTimeout int32 `json:"concurrencyTimeout,omitempty"`
	// If the Rule should be triggered no more than once for the same device, false otherwise
	OncePerDevice bool `json:"oncePerDevice,omitempty"`
	// If the Rule should be automatically disabled once all targets from the Scope pass through the triggering conditions, false otherwise
	AutoDisable bool `json:"autoDisable,omitempty"`
	// A time schedule in which the Task execution is permitted. The time schedule constraint must be provided as a cron expression. <code>Example:</code> <li/>* * 9-17 ? * MON-FRI: defines that the execution time will be on working days (Monday to Friday) between 9.00-17.00h<li/>* * 00-01 * * ?: defines that execution time will be every day between 00-01 in the morning.<li/>* * * 01 * ?: defines that execution time will be on the 1st day of each month .
	TimeConstraint string `json:"timeConstraint"`
}
