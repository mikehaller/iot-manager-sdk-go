/*
 * Bosch IoT Manager API
 *
 * The Bosch IoT Manager REST API provides the full functionality of the service, including:  &middot;    Retrieving devices and features; adding and modifying device-specific properties and attributes  &middot;    A simple device registration mechanism  &middot;    All grouping capabilities - listing, creating, modifying and deleting directories and tags  &middot;    The complete mass management experience, which allows full control over tasks and rules   Find out more details in our [documentation](https://docs.bosch-iot-suite.com/manager/).
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotmgr
// Credentials  Credentials are used to authenticate Devices connecting to the adapter. They could be of a certain type which indicates which authentication mechanism the credentials can be used with
type Credentials struct {
	// The type of credentials. Values could be on of:  <ul>   <li>HASHED_PASSWORD - A credential type for storing a password for a device</li>   <li>PRE_SHARED_SECRET - A credential type for storing a Pre-shared Key as used in TLS handshakes</li>   <li>CERTIFICATE - A credential type for storing the formatted subject DN of a client certificate that is used to authenticate the device as part of a TLS handshake</li>  </ul>
	Type string `json:"type"`
	// The identity that the device should be authenticated as
	AuthId string `json:"authId,omitempty"`
	// Secret to authenticate the device against(valid for HASHED_PASSWORD and PRE_SHARED_SECRET type)
	Secret string `json:"secret,omitempty"`
	// If set to false the credentials are not supposed to be used to authenticate devices any longer.
	Enabled bool `json:"enabled,omitempty"`
}
