# \DevicesApi

All URIs are relative to *https://manager.eu-1.bosch-iot-suite.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Attribute**](DevicesApi.md#Attribute) | **Get** /di/devices/{deviceId}/attributes/{attributePath} | Retrieve a specific Attribute of a specific device
[**CreateAttributes**](DevicesApi.md#CreateAttributes) | **Post** /di/devices/{deviceId}/attributes | Add Attributes of a specific Device
[**Definition**](DevicesApi.md#Definition) | **Get** /di/devices/{deviceId}/features/{featureId}/definitions/{fullyQualifiedIdentifier} | Retrieve the Definition of a Feature
[**Definitions**](DevicesApi.md#Definitions) | **Get** /di/devices/{deviceId}/features/{featureId}/definitions | List all Definitions of a Feature
[**DeleteAttribute**](DevicesApi.md#DeleteAttribute) | **Delete** /di/devices/{deviceId}/attributes/{attributePath} | Delete Attribute of a specific Device
[**DeleteAttributes**](DevicesApi.md#DeleteAttributes) | **Delete** /di/devices/{deviceId}/attributes | Delete Attributes of a specific Device
[**Device**](DevicesApi.md#Device) | **Get** /di/devices/{deviceId} | Retrieve an existing Device identified by the given deviceId
[**Devices**](DevicesApi.md#Devices) | **Get** /di/gateways/{gatewayId}/devices | Retrieve a list of Edge Devices attached to the gateway
[**Devices1**](DevicesApi.md#Devices1) | **Get** /di/devices | Retrieve all Devices
[**DevicesCount**](DevicesApi.md#DevicesCount) | **Get** /di/devices/count | Retrieve the number of Devices matching the given filter
[**DevicesCount1**](DevicesApi.md#DevicesCount1) | **Get** /di/gateways/{gatewayId}/devices/count | Retrieve the number of Devices matching the given filter
[**Feature**](DevicesApi.md#Feature) | **Get** /di/devices/{deviceId}/features/{featureId} | Retrieve a Feature of a Device
[**FeatureProperties**](DevicesApi.md#FeatureProperties) | **Get** /di/devices/{deviceId}/features/{featureId}/properties | List the Properties of a Feature
[**Features**](DevicesApi.md#Features) | **Get** /di/devices/{deviceId}/features | List all Features of a specific device
[**Gateway**](DevicesApi.md#Gateway) | **Get** /di/gateways/{gatewayId} | Retrieve an existing Gateway identified by the given gatewayId
[**Gateways**](DevicesApi.md#Gateways) | **Get** /di/gateways | Retrieve all Gateways
[**GatewaysCount**](DevicesApi.md#GatewaysCount) | **Get** /di/gateways/count | Retrieve the number of gateways matching the given filter
[**ListAttributes**](DevicesApi.md#ListAttributes) | **Get** /di/devices/{deviceId}/attributes | List all Attributes of a specific Device
[**Property**](DevicesApi.md#Property) | **Get** /di/devices/{deviceId}/features/{featureId}/properties/{propertyName} | Retrieve the Property of a Feature
[**Register**](DevicesApi.md#Register) | **Post** /di/devices/{deviceId} | Register a specific Device
[**SetAttribute**](DevicesApi.md#SetAttribute) | **Put** /di/devices/{deviceId}/attributes | Add Attribute of a specific Device
[**Unregister**](DevicesApi.md#Unregister) | **Delete** /di/devices/{deviceId} | Unregister a specific Device



## Attribute

> Attribute Attribute(ctx, deviceId, attributePath)

Retrieve a specific Attribute of a specific device

Retrieve a specific Attribute of the device, identified by the <code>deviceId</code> path parameter. The Attribute (JSON) can be referenced hierarchically by applying a JSON Pointer notation (RFC-6901), e.g.: /devices/{deviceId}/attributes/house/room in order to retrieve the house field of an room object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
**attributePath** | **string**|  | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAttributes

> Attribute CreateAttributes(ctx, deviceId, optional)

Add Attributes of a specific Device

Add Attributes of the specific Device. All previously existing attributes will be overwritten.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
 **optional** | ***CreateAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAttributesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributeObject** | [**optional.Interface of []AttributeObject**](AttributeObject.md)| List of Attributes | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Definition

> FeatureDefinition Definition(ctx, deviceId, featureId, fullyQualifiedIdentifier)

Retrieve the Definition of a Feature

Retrieve the complete Definition of the Feature identified by the <code>deviceId</code>, <code>featureId</code> and <code>fullyQualifiedIdentifier<code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
**featureId** | **string**|  | 
**fullyQualifiedIdentifier** | **string**|  Pattern: ([_a-zA-Z0-9\\-.]+):([_a-zA-Z0-9\\-.]+):([_a-zA-Z0-9\\-.]+)    A single fully qualified identifier of a Feature Definition in the form &#39;namespace:name:version&#39; | 

### Return type

[**FeatureDefinition**](FeatureDefinition.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Definitions

> []FeatureDefinition Definitions(ctx, deviceId, featureId)

List all Definitions of a Feature

List all Definitions of the Feature identified by the <code>deviceId</code> and <code>featureId</code> path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
**featureId** | **string**|  | 

### Return type

[**[]FeatureDefinition**](FeatureDefinition.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttribute

> DeleteAttribute(ctx, deviceId, attributePath)

Delete Attribute of a specific Device

Delete Attribute of a specific Device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
**attributePath** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttributes

> DeleteAttributes(ctx, deviceId)

Delete Attributes of a specific Device

Delete Attributes of a specific Device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Device

> Device Device(ctx, deviceId)

Retrieve an existing Device identified by the given deviceId

Retrieve an existing Device identified by the given <code>deviceId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Devices

> []Device Devices(ctx, gatewayId, optional)

Retrieve a list of Edge Devices attached to the gateway

Retrieve a list of edge devices attached to the gateway identified by <code>gatewayId</code> and matching the given filter. Optionally, you can also use field selectors (see parameter fields) to get only the specified fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 
 **optional** | ***DevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| &lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only devices with the specified properties are returned. For example, the filter ne(attributes/owner, \&quot;SID123\&quot;) will only return devices that do have the owner attribute.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq(attributes/{attributeId},\&quot;value\&quot;)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists(features/{featureId}/properties/{category}/{propertyId}, value)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq(attributes/{attributeId1},\&quot;value\&quot;),eq(attributes/{attributeId2},\&quot;value\&quot;))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;b&gt;Filters containing a wildcard * symbol at the beginning can slow down your search request.&lt;/b&gt; | 
 **namespaces** | **optional.String**| A comma separated list of namespaces. This list is used to limit the query to things in the given namespaces only. If this parameter is omitted, all registered namespaces of your solution will be queried.&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?namespaces&#x3D;com.example.namespace&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?namespaces&#x3D;com.example.namespace1,com.example.namespace2&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;deviceId,attributes/location/longitude,attributes/address(city,street)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt; | 
 **option** | **optional.String**| &lt;p&gt;Possible parameter values:&lt;/p&gt;&lt;h5&gt;Sort operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;sort([+|-]{property})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;sort([+|-]{property},[+|-]{property},...)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25.&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;sort(+id)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;sort(-attributes/{attributePropertyValue})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM). | 
 **fields** | **optional.String**| Contains a comma separated list of fields to be included in the returned JSON. Attributes can be selected in the same manner.&lt;br&gt;&lt;p&gt;&lt;b&gt;Selectable fields&lt;/b&gt;&lt;/p&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;ID&lt;/tt&gt;&lt;/li&gt; &lt;li&gt;&lt;tt&gt;policyId&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;attributes&lt;/tt&gt;&lt;p&gt;Supports selecting arbitrary sub-fields by using a comma separated list:&lt;/p&gt;&lt;ul&gt;&lt;li&gt;several attribute paths can be passed as a comma separated list of JSON pointers (RFC-6901)&lt;br&gt;&lt;br&gt;&lt;b&gt;For example:&lt;/b&gt;&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId}/{attributeValue}&lt;/tt&gt; would select only attribute value (if present)&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}/{attributeValue},attributes/{attributeId2}/{attributeValue}&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}/{attributeValue1},attributes/{attributeId1}/{attributeValue2}&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;li&gt;a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901) separated with /) to select &lt;/li&gt;&lt;li&gt;sub-selectors can be used to request only specific sub-fields by placing expressions in parentheses ( ) after a selected subfield&lt;br&gt;&lt;b&gt;For example:&lt;/b&gt;&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId}/(attributeValue1,attributeValue2)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;features&lt;/tt&gt;&lt;p&gt;Supports selecting arbitrary fields in features similar to attributes&lt;/p&gt;&lt;b&gt;features/{featureId}/properties/{category}/{propertyId}&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;id,attributes,features&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}(attributeValue1,attributeValue1),features&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;id,attributes/{attributeId1},features/{featureId}/properties/{category}/{propertyId}&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**[]Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Devices1

> []Device Devices1(ctx, optional)

Retrieve all Devices

Retrieve a list of devices matching the given filter. Optionally, you can also use field.selectors (see parameter fields) to get only the specified fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Devices1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Devices1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| &lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only devices with the specified properties are returned. For example, the filter ne(attributes/owner, \&quot;SID123\&quot;) will only return devices that do have the owner attribute.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq(attributes/{attributeId},\&quot;value\&quot;)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists(features/{featureId}/properties/{category}/{propertyId}, value)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq(attributes/{attributeId1},\&quot;value\&quot;),eq(attributes/{attributeId2},\&quot;value\&quot;))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;b&gt;Filters containing a wildcard * symbol at the beginning can slow down your search request.&lt;/b&gt; | 
 **namespaces** | **optional.String**| A comma separated list of namespaces. This list is used to limit the query to things in the given namespaces only. If this parameter is omitted, all registered namespaces of your solution will be queried.&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?namespaces&#x3D;com.example.namespace&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?namespaces&#x3D;com.example.namespace1,com.example.namespace2&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;deviceId,attributes/location/longitude,attributes/address(city,street)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt; | 
 **option** | **optional.String**| &lt;p&gt;Possible parameter values:&lt;/p&gt;&lt;h5&gt;Sort operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;sort([+|-]{property})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;sort([+|-]{property},[+|-]{property},...)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25.&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;sort(+id)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;sort(-attributes/{attributePropertyValue})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM). | 
 **fields** | **optional.String**| Contains a comma separated list of fields to be included in the returned JSON. Attributes can be selected in the same manner.&lt;br&gt;&lt;p&gt;&lt;b&gt;Selectable fields&lt;/b&gt;&lt;/p&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;ID&lt;/tt&gt;&lt;/li&gt; &lt;li&gt;&lt;tt&gt;policyId&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;attributes&lt;/tt&gt;&lt;p&gt;Supports selecting arbitrary sub-fields by using a comma separated list:&lt;/p&gt;&lt;ul&gt;&lt;li&gt;several attribute paths can be passed as a comma separated list of JSON pointers (RFC-6901)&lt;br&gt;&lt;br&gt;&lt;b&gt;For example:&lt;/b&gt;&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId}/{attributeValue}&lt;/tt&gt; would select only attribute value (if present)&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}/{attributeValue},attributes/{attributeId2}/{attributeValue}&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}/{attributeValue1},attributes/{attributeId1}/{attributeValue2}&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;li&gt;a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901) separated with /) to select &lt;/li&gt;&lt;li&gt;sub-selectors can be used to request only specific sub-fields by placing expressions in parentheses ( ) after a selected subfield&lt;br&gt;&lt;b&gt;For example:&lt;/b&gt;&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId}/(attributeValue1,attributeValue2)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;features&lt;/tt&gt;&lt;p&gt;Supports selecting arbitrary fields in features similar to attributes&lt;/p&gt;&lt;b&gt;features/{featureId}/properties/{category}/{propertyId}&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;id,attributes,features&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}(attributeValue1,attributeValue1),features&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;id,attributes/{attributeId1},features/{featureId}/properties/{category}/{propertyId}&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**[]Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesCount

> DevicesCount(ctx, optional)

Retrieve the number of Devices matching the given filter

Retrieve the number of devices matching the given filter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| &lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only devices with the specified properties are returned. For example, the filter ne(attributes/owner, \&quot;SID123\&quot;) will only return devices that do have the owner attribute.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq(attributes/{attributeId},\&quot;value\&quot;)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists(features/{featureId}/properties/{category}/{propertyId}, value)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq(attributes/{attributeId1},\&quot;value\&quot;),eq(attributes/{attributeId2},\&quot;value\&quot;))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;b&gt;Filters containing a wildcard * symbol at the beginning can slow down your search request.&lt;/b&gt; | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesCount1

> DevicesCount1(ctx, gatewayId, optional)

Retrieve the number of Devices matching the given filter

Retrieve the number of devices matching the given filter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 
 **optional** | ***DevicesCount1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesCount1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| &lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only devices with the specified properties are returned. For example, the filter ne(attributes/owner, \&quot;SID123\&quot;) will only return devices that do have the owner attribute.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq(attributes/{attributeId},\&quot;value\&quot;)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists(features/{featureId}/properties/{category}/{propertyId}, value)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq(attributes/{attributeId1},\&quot;value\&quot;),eq(attributes/{attributeId2},\&quot;value\&quot;))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;b&gt;Filters containing a wildcard * symbol at the beginning can slow down your search request.&lt;/b&gt; | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Feature

> Feature Feature(ctx, deviceId, featureId)

Retrieve a Feature of a Device

Retrieve a Feature of the Device identified by the deviceId path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
**featureId** | **string**|  | 

### Return type

[**Feature**](Feature.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FeatureProperties

> []Property FeatureProperties(ctx, deviceId, featureId)

List the Properties of a Feature

Retrieve the complete Properties of the Feature identified by the <code>deviceId</code> and <code>featureId</code> path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
**featureId** | **string**|  | 

### Return type

[**[]Property**](Property.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Features

> []Feature Features(ctx, deviceId)

List all Features of a specific device

Retrieve all Features of the Device identified by the <code>deviceId</code> path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 

### Return type

[**[]Feature**](Feature.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Gateway

> []Gateway Gateway(ctx, gatewayId)

Retrieve an existing Gateway identified by the given gatewayId

Retrieve an existing Gateway identified by the given gatewayId.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 

### Return type

[**[]Gateway**](Gateway.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Gateways

> []Gateway Gateways(ctx, optional)

Retrieve all Gateways

Retrieve a list of gateways matching the given filter. Optionally, you can also use field selectors (see parameter fields) to get only specified fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GatewaysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GatewaysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| &lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only devices with the specified properties are returned. For example, the filter ne(attributes/owner, \&quot;SID123\&quot;) will only return devices that do have the owner attribute.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq(attributes/{attributeId},\&quot;value\&quot;)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists(features/{featureId}/properties/{category}/{propertyId}, value)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq(attributes/{attributeId1},\&quot;value\&quot;),eq(attributes/{attributeId2},\&quot;value\&quot;))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;b&gt;Filters containing a wildcard * symbol at the beginning can slow down your search request.&lt;/b&gt; | 
 **namespaces** | **optional.String**| A comma separated list of namespaces. This list is used to limit the query to things in the given namespaces only. If this parameter is omitted, all registered namespaces of your solution will be queried.&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?namespaces&#x3D;com.example.namespace&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?namespaces&#x3D;com.example.namespace1,com.example.namespace2&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;deviceId,attributes/location/longitude,attributes/address(city,street)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt; | 
 **option** | **optional.String**| &lt;p&gt;Possible parameter values:&lt;/p&gt;&lt;h5&gt;Sort operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;sort([+|-]{property})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;sort([+|-]{property},[+|-]{property},...)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25.&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;sort(+id)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;sort(-attributes/{attributePropertyValue})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM). | 
 **fields** | **optional.String**| Contains a comma separated list of fields to be included in the returned JSON. Attributes can be selected in the same manner.&lt;br&gt;&lt;p&gt;&lt;b&gt;Selectable fields&lt;/b&gt;&lt;/p&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;ID&lt;/tt&gt;&lt;/li&gt; &lt;li&gt;&lt;tt&gt;policyId&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;attributes&lt;/tt&gt;&lt;p&gt;Supports selecting arbitrary sub-fields by using a comma separated list:&lt;/p&gt;&lt;ul&gt;&lt;li&gt;several attribute paths can be passed as a comma separated list of JSON pointers (RFC-6901)&lt;br&gt;&lt;br&gt;&lt;b&gt;For example:&lt;/b&gt;&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId}/{attributeValue}&lt;/tt&gt; would select only attribute value (if present)&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}/{attributeValue},attributes/{attributeId2}/{attributeValue}&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}/{attributeValue1},attributes/{attributeId1}/{attributeValue2}&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;li&gt;a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901) separated with /) to select &lt;/li&gt;&lt;li&gt;sub-selectors can be used to request only specific sub-fields by placing expressions in parentheses ( ) after a selected subfield&lt;br&gt;&lt;b&gt;For example:&lt;/b&gt;&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId}/(attributeValue1,attributeValue2)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;features&lt;/tt&gt;&lt;p&gt;Supports selecting arbitrary fields in features similar to attributes&lt;/p&gt;&lt;b&gt;features/{featureId}/properties/{category}/{propertyId}&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;id,attributes,features&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}(attributeValue1,attributeValue1),features&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;id,attributes/{attributeId1},features/{featureId}/properties/{category}/{propertyId}&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**[]Gateway**](Gateway.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewaysCount

> GatewaysCount(ctx, optional)

Retrieve the number of gateways matching the given filter

Retrieve the number of gateways matching the given filter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GatewaysCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GatewaysCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| &lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only devices with the specified properties are returned. For example, the filter ne(attributes/owner, \&quot;SID123\&quot;) will only return devices that do have the owner attribute.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq(attributes/{attributeId},\&quot;value\&quot;)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists(features/{featureId}/properties/{category}/{propertyId}, value)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq(attributes/{attributeId1},\&quot;value\&quot;),eq(attributes/{attributeId2},\&quot;value\&quot;))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;b&gt;Filters containing a wildcard * symbol at the beginning can slow down your search request.&lt;/b&gt; | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttributes

> []Attribute ListAttributes(ctx, deviceId)

List all Attributes of a specific Device

List all Attributes of a specific Device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 

### Return type

[**[]Attribute**](Attribute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Property

> Property Property(ctx, deviceId, featureId, propertyName)

Retrieve the Property of a Feature

Retrieve the complete Property of the Feature identified by the <code>deviceId</code> and <code>featureId</code> path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
**featureId** | **string**|  | 
**propertyName** | **string**|  | 

### Return type

[**Property**](Property.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> Property Register(ctx, deviceId, optional)

Register a specific Device

Register a specific Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| Device ID | 
 **optional** | ***RegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerInfo** | [**optional.Interface of RegisterInfo**](RegisterInfo.md)| Registration info | 

### Return type

[**Property**](Property.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAttribute

> Attribute SetAttribute(ctx, deviceId, optional)

Add Attribute of a specific Device

Add Attribute of the specific Device. If an attribute with the same path already exists it will be overwritten.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
 **optional** | ***SetAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetAttributeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attribute** | [**optional.Interface of Attribute**](Attribute.md)| Directory | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unregister

> Property Unregister(ctx, deviceId, optional)

Unregister a specific Device

Unregister a specific Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| Device ID | 
 **optional** | ***UnregisterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnregisterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keepCredentials** | **optional.Bool**|  | [default to false]
 **keepPolicy** | **optional.Bool**|  | [default to false]

### Return type

[**Property**](Property.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

