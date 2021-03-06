/*
 * Bosch IoT Manager API
 *
 * The Bosch IoT Manager REST API provides the full functionality of the service, including:  &middot;    Retrieving devices and features; adding and modifying device-specific properties and attributes  &middot;    A simple device registration mechanism  &middot;    All grouping capabilities - listing, creating, modifying and deleting directories and tags  &middot;    The complete mass management experience, which allows full control over tasks and rules   Find out more details in our [documentation](https://docs.bosch-iot-suite.com/manager/).
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotmgr
// Operation struct for Operation
type Operation struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Result ReturnType `json:"result,omitempty"`
	Params []Param `json:"params,omitempty"`
	Breakable bool `json:"breakable,omitempty"`
}
