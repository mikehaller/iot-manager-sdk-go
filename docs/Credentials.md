# Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of credentials. Values could be on of:  &lt;ul&gt;   &lt;li&gt;HASHED_PASSWORD - A credential type for storing a password for a device&lt;/li&gt;   &lt;li&gt;PRE_SHARED_SECRET - A credential type for storing a Pre-shared Key as used in TLS handshakes&lt;/li&gt;   &lt;li&gt;CERTIFICATE - A credential type for storing the formatted subject DN of a client certificate that is used to authenticate the device as part of a TLS handshake&lt;/li&gt;  &lt;/ul&gt; | 
**AuthId** | **string** | The identity that the device should be authenticated as | [optional] 
**Secret** | **string** | Secret to authenticate the device against(valid for HASHED_PASSWORD and PRE_SHARED_SECRET type) | [optional] 
**Enabled** | **bool** | If set to false the credentials are not supposed to be used to authenticate devices any longer. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


