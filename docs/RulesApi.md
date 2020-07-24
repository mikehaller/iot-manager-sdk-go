# \RulesApi

All URIs are relative to *https://manager.eu-1.bosch-iot-suite.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Action**](RulesApi.md#Action) | **Get** /mme/rules/{ruleId}/action | Retrieve the Rule action
[**CreateRule**](RulesApi.md#CreateRule) | **Post** /mme/rules | Create a Rule
[**DeleteRule**](RulesApi.md#DeleteRule) | **Delete** /mme/rules/{ruleId} | Delete the Rule
[**EventMetadata**](RulesApi.md#EventMetadata) | **Get** /mme/rules/trigger/types/{triggerType}/eventsMetadata | Retrieve a JSON schema of all possible events
[**Fire**](RulesApi.md#Fire) | **Put** /mme/rules/{ruleId}/tasks | Fire the Rule manually
[**GetActionDefMetada**](RulesApi.md#GetActionDefMetada) | **Get** /mme/rules/actions/{actionType} | Retrieve metadata in JSON schema format 
[**GetActionTypes**](RulesApi.md#GetActionTypes) | **Get** /mme/rules/actions | Retrieve all available registered Action types
[**ListDevicesForRule**](RulesApi.md#ListDevicesForRule) | **Get** /mme/rules/{ruleId}/devices | Retrieve a list of all involved Devices
[**ListRulesForDevice**](RulesApi.md#ListRulesForDevice) | **Get** /mme/rules/devices/{deviceId} | Retrieve all Rules involving a particular Device
[**ListTasksForDeviceInRule**](RulesApi.md#ListTasksForDeviceInRule) | **Get** /mme/rules/{ruleId}/devices/{deviceId}/tasks | Retrieve all Tasks launched by the Rule which involve a specified Device
[**Modify**](RulesApi.md#Modify) | **Put** /mme/rules/{ruleId} | Modify an existing Rule
[**Options**](RulesApi.md#Options) | **Get** /mme/rules/{ruleId}/options | Retrieve the Rule execution options
[**Properties**](RulesApi.md#Properties) | **Get** /mme/rules/{ruleId}/properties | Get the custom Rule properties
[**RetryDevice**](RulesApi.md#RetryDevice) | **Put** /mme/rules/{ruleId}/devices/{deviceId} | Retry the Rule over the Device
[**RetryDevices**](RulesApi.md#RetryDevices) | **Put** /mme/rules/{ruleId}/devices | Retry the Rule over the Devices whose state satisfies the applied stateFilter
[**Rule**](RulesApi.md#Rule) | **Get** /mme/rules/{ruleId}/status | Retrieve a Rule
[**RuleDevice**](RulesApi.md#RuleDevice) | **Get** /mme/rules/{ruleId}/devices/{deviceId} | Retrieve the execution status of a specified Device within the Rule
[**Rules**](RulesApi.md#Rules) | **Get** /mme/rules | Retrieve all available Rules
[**Scope**](RulesApi.md#Scope) | **Get** /mme/rules/{ruleId}/scope | Retrieve the Rule scope
[**SetProperties**](RulesApi.md#SetProperties) | **Post** /mme/rules/{ruleId}/properties | Set custom Rule properties
[**SetState**](RulesApi.md#SetState) | **Put** /mme/rules/{ruleId}/state | Enable or disable the Rule
[**State**](RulesApi.md#State) | **Get** /mme/rules/{ruleId}/state | Retrieve the Rule state
[**Tasks**](RulesApi.md#Tasks) | **Get** /mme/rules/{ruleId}/tasks | Retrieve all Tasks launched by the Rule
[**TriggerMetadata**](RulesApi.md#TriggerMetadata) | **Get** /mme/rules/trigger/types/{triggerType}/inputMetadata | Retrieve a JSON schema of the required input
[**TriggerTypes**](RulesApi.md#TriggerTypes) | **Get** /mme/rules/trigger/types | Retrieve all available registered Trigger types



## Action

> Action(ctx, ruleId)

Retrieve the Rule action

Retrieve the Groovy script defining the action for this Rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 

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


## CreateRule

> RuleStatusInfo CreateRule(ctx, optional)

Create a Rule

Create a Rule with the specified attributes - <code>displayName</code>, <code>scope</code>, <code>trigger</code>, <code>options</code>, and <code>action</code>. The <code>displayName</code>, <code>scope</code>, <code>trigger</code>, and <code>action</code> are required.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleInfo** | [**optional.Interface of RuleInfo**](RuleInfo.md)| Creates new rule | 

### Return type

[**RuleStatusInfo**](RuleStatusInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRule(ctx, ruleId)

Delete the Rule

Delete the Rule with the specified <code>ruleId</code>. If the Rule could not be found or requester has insufficient permissions to access it, а <code>204 No Content </code>will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 

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


## EventMetadata

> EventMetadata(ctx, triggerType)

Retrieve a JSON schema of all possible events

Retrieve a JSON schema describing a tree-structured collection of all possible events delivered by this Trigger type.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerType** | **string**|  | 

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


## Fire

> Fire(ctx, ruleId)

Fire the Rule manually

Fire the Rule with the specified <code>ruleId</code> manually.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 

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


## GetActionDefMetada

> GetActionDefMetada(ctx, actionType)

Retrieve metadata in JSON schema format 

Retrieve JSON schema format metadata specified by the given <code>actionType</code>, describing the input needed on Rule activation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionType** | **string**|  | 

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


## GetActionTypes

> GetActionTypes(ctx, )

Retrieve all available registered Action types

Retrieve all available registered Action provider types. Groovy Script action type is available by default.

### Required Parameters

This endpoint does not need any parameter.

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


## ListDevicesForRule

> ListDevicesForRule(ctx, ruleId, optional)

Retrieve a list of all involved Devices

Retrieve a list of all Devices which are involved in this Rule and satisfy the applied <code>stateFilter</code>. A Device is considered involved in the Rule if it is involved in at least one Task launched by that Rule.<p> <h4>Available stateFilter values:</h4> <ui><li><tt> FINISHED_SUCCESS(0) : Retrieve only Devices that have finished with a success state</tt></li>  <li><tt> FINISHED_WARNING(1) : Retrieve only Devices that have finished the execution with a warning state</tt></li> <li><tt> FINISHED_ERROR(2) : Retrieve only Devices that have finished the execution with an error state</tt></li> <li><tt> FINISHED_CANCELED(3) : Retrieve only devices that have finished the execution with a canceled state</tt></li> <li><tt> RUNNING(4) : Retrieve only Devices that have a running state</tt></li>  <li><tt>FINISHED(64): This value unites all FINISHED_xxx values and could be used only as a parameter for filtering</tt></li><li><tt>ANY(128): Retrieve all involved devices, regardless of the completion state

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 
 **optional** | ***ListDevicesForRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDevicesForRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stateFilter** | **optional.String**|  | 
 **option** | **optional.String**| &lt;p&gt;Possible values for the parameter:&lt;/p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM) | 

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


## ListRulesForDevice

> ListRulesForDevice(ctx, deviceId, optional)

Retrieve all Rules involving a particular Device

Retrieve all Rules involving a particular Device in which the latter satisfies the applied state filter.<p> <h4>Available stateFilter values:</h4> <ui><li><tt> FINISHED_SUCCESS(0) : Retrieve only Devices that have finished with a success state</tt></li>  <li><tt> FINISHED_WARNING(1) : Retrieve only Devices that have finished the execution with a warning state</tt></li> <li><tt> FINISHED_ERROR(2) : Retrieve only Devices that have finished the execution with an error state</tt></li> <li><tt> FINISHED_CANCELED(3) : Retrieve only devices that have finished the execution with a canceled state</tt></li> <li><tt> RUNNING(4) : Retrieve only Devices that have a running state</tt></li>  <li><tt>FINISHED(64): This value unites all FINISHED_xxx values and could be used only as a parameter for filtering</tt></li><li><tt>ANY(128): Retrieve all involved devices, regardless of the completion state

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
 **optional** | ***ListRulesForDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRulesForDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stateFilter** | **optional.String**|  | 
 **option** | **optional.String**| &lt;p&gt;Possible values for the parameter:&lt;/p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM) | 

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


## ListTasksForDeviceInRule

> ListTasksForDeviceInRule(ctx, ruleId, deviceId, optional)

Retrieve all Tasks launched by the Rule which involve a specified Device

Retrieve all Tasks launched by the Rule which involve a specified Device. Limited amount of data is retrieved, by default on a page with 20 elements.<p> <h4>Available stateFilter values:</h4> <ui><li><tt> FINISHED_SUCCESS(0) : Retrieve only Devices that have finished with a success state</tt></li>  <li><tt> FINISHED_WARNING(1) : Retrieve only Devices that have finished the execution with a warning state</tt></li> <li><tt> FINISHED_ERROR(2) : Retrieve only Devices that have finished the execution with an error state</tt></li> <li><tt> FINISHED_CANCELED(3) : Retrieve only devices that have finished the execution with a canceled state</tt></li> <li><tt> RUNNING(4) : Retrieve only Devices that have a running state</tt></li>  <li><tt>FINISHED(64): This value unites all FINISHED_xxx values and could be used only as a parameter for filtering</tt></li><li><tt>ANY(128): Retrieve all involved devices, regardless of the completion state

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 
**deviceId** | **string**|  | 
 **optional** | ***ListTasksForDeviceInRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTasksForDeviceInRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stateFilter** | **optional.String**|  | 
 **option** | **optional.String**| &lt;p&gt;Possible values for the parameter:&lt;/p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM) | 

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


## Modify

> Modify(ctx, ruleId, optional)

Modify an existing Rule

Modify an existing Rule. Only Rules whose state is not enabled can be modified. After a Rule is modified, any information related to its previous execution (e.g. launched Tasks, number of involved, finished Devices, etc.) will be cleared.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 
 **optional** | ***ModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleInfo** | [**optional.Interface of RuleInfo**](RuleInfo.md)| Modifies the Rule | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Options

> Options(ctx, ruleId)

Retrieve the Rule execution options

Retrieve the execution options of the Rule with the specified <code>ruleId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 

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


## Properties

> Properties(ctx, ruleId)

Get the custom Rule properties

Get the custom properties of the Rule with the specified <code>ruleId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 

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


## RetryDevice

> RetryDevice(ctx, ruleId, deviceId)

Retry the Rule over the Device

Retry the Rule over the Device with the specified <code>deviceId</code>. 'Once Per Device' option must be selected for the Rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 
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


## RetryDevices

> RetryDevices(ctx, ruleId, stateFilter)

Retry the Rule over the Devices whose state satisfies the applied stateFilter

Retry the Rule over the Devices whose state satisfies the applied <code>stateFilter</code>. 'Once Per Device' option must be selected for the Rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 
**stateFilter** | [**[]string**](string.md)|  | 

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


## Rule

> RuleStatusInfo Rule(ctx, ruleId, statistic)

Retrieve a Rule

Retrieve a Rule by a specified <code>ruleId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 
**statistic** | **bool**| Load statistic | [default to true]

### Return type

[**RuleStatusInfo**](RuleStatusInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RuleDevice

> RuleDevice(ctx, ruleId, deviceId)

Retrieve the execution status of a specified Device within the Rule

Retrieve the execution status of a specified Device within the Rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 
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


## Rules

> Rules(ctx, optional)

Retrieve all available Rules

 Retrieve all available Rules that satisfy the specified filtering criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| An RQL filter over the Rule attributes.&lt;h5&gt;&lt;/tt&gt;&lt;br/&gt;The attribute names in the filtering conditions can be some of:&lt;/h5&gt; &lt;ui&gt;&lt;tt&gt;&lt;li/&gt;&lt;b&gt;state&lt;/b&gt; - integer value corresponding to the Rule state as follows: &lt;/tt&gt; &lt;br/&gt; &lt;tt&gt;0 - DEFINED &lt;/tt&gt; &lt;br/&gt;&lt;tt&gt;1 - ENABLED &lt;/tt&gt;&lt;br/&gt;&lt;tt&gt; 2 - DISABLED&lt;/tt&gt; &lt;br/&gt;&lt;tt&gt;3 - FAILED_TO_ENABLE&lt;/tt&gt; &lt;li/&gt;&lt;tt&gt;&lt;b&gt;stateDescription&lt;/b&gt; - string value corresponding to the state description. &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt; involvedCount&lt;/b&gt; - integer value corresponding to the the number of involved Devices in the Rule.&lt;/tt&gt; &lt;li/&gt;&lt;tt&gt;&lt;b&gt; successCount&lt;/b&gt; - integer value corresponding to the the number of successfully finished Devices in the Rule.&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;warningCount&lt;/b&gt; - integer value corresponding to the the number of Devices finished with warning in the Rule. &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt; errorCount&lt;/b&gt; - integer value corresponding the the number of Devices finished with error in the Rule.&lt;/tt&gt; &lt;li/&gt;&lt;tt&gt;&lt;b&gt;cancelCount&lt;/b&gt; - integer value corresponding to the the number of Devices finished with warning in the Rule.&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;nextTimeTrigger&lt;/b&gt; - corresponds to the next trigger time (in milliseconds since 1 Jan 1970) if the Rule has a timer trigger.&lt;/tt&gt; &lt;li/&gt;&lt;tt&gt;&lt;b&gt;ruleId&lt;/b&gt; - string value corresponding to the Rule id. &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;user&lt;/b&gt; - string value corresponding to the user name of the user that has created the Rule.&lt;li/&gt; &lt;tt&gt;&lt;b&gt;displayName&lt;/b&gt; - string value corresponding to the Rule display name. &lt;/tt&gt;&lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only Rules with the specified properties are returned.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq(state,0)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq(involvedCount,2),eq(state,0))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt; | 
 **option** | **optional.String**| &lt;p&gt;Possible values for the parameter:&lt;/p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM) | 

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


## Scope

> Scope(ctx, ruleId)

Retrieve the Rule scope

Retrieve the scope of an existing Rule identified by the specified <code>ruleId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 

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


## SetProperties

> SetProperties(ctx, ruleId, optional)

Set custom Rule properties

Set custom properties of the Rule with the specified <code>ruleId</code>. If <tt> append </tt> is <tt> true</tt> the supplied properties will be appended/added to any properties currently existing for  this Rule, otherwise any existing properties will be fully replaced by the newly set ones.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 
 **optional** | ***SetPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **properties** | [**optional.Interface of Properties**](Properties.md)| Set or add rule properties | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetState

> SetState(ctx, ruleId, optional)

Enable or disable the Rule

Enable or disable the Rule with the specified <code>ruleId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 
 **optional** | ***SetStateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetStateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**|  | 

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


## State

> string State(ctx, ruleId)

Retrieve the Rule state

Retrieve the state of the Rule with the specified <code>ruleId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Tasks

> Tasks(ctx, ruleId, optional)

Retrieve all Tasks launched by the Rule

Retrieve all Tasks launched by the Rule. By default, only a limited amount of data is retrieved on a page with 20 elements.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 
 **optional** | ***TasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TasksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| An RQL filter over the Task attributes.&lt;h5&gt;&lt;/tt&gt;&lt;br/&gt;The attribute names in the filtering conditions can be some of:&lt;/h5&gt; &lt;ui&gt;&lt;tt&gt;&lt;li/&gt;&lt;b&gt;state&lt;/b&gt; - integer value corresponding to the Task state as follows: &lt;/tt&gt; &lt;br/&gt;&lt;br/&gt;&lt;tt&gt;1 - RUNNING &lt;/tt&gt;&lt;br/&gt;&lt;tt&gt; 2 - FINISHED&lt;/tt&gt; &lt;br/&gt;&lt;tt&gt;3 - FAILED_TO_LAUNCH&lt;/tt&gt; &lt;li/&gt;&lt;tt&gt;&lt;b&gt;stateDescription&lt;/b&gt; - string value corresponding to the state description &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt; involvedCount&lt;/b&gt; - integer value corresponding to the the number of involved Devices in the Task.&lt;/tt&gt; &lt;li/&gt;&lt;tt&gt;&lt;b&gt; successCount&lt;/b&gt; - integer value corresponding to the the number of successfully finished Devices in the Task&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;warningCount&lt;/b&gt; - integer value corresponding the the number of Devices finished with warning in the Task &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt; errorCount&lt;/b&gt; - integer value corresponding to the the number of Devices finished with error in the Task.&lt;/tt&gt; &lt;li/&gt;&lt;tt&gt;&lt;b&gt;cancelCount&lt;/b&gt; - integer value corresponding to the the number of Devices finished with warning in the Task.&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;launchTime&lt;/b&gt; - long value corresponding to the Task launch time given as milliseconds since 1 Jan 1970&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;finishTime&lt;/b&gt; - long value corresponding to the Task finish time given as milliseconds since 1 Jan 1970&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;isPaused&lt;/b&gt; - boolean value that is &lt;code&gt;true&lt;/code&gt; when the execution is paused due to time constraint restrictions &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;nextResume&lt;/b&gt; - corresponds to resume time (in milliseconds since 1 Jan 1970) if the Task is currently paused &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;nextPause&lt;/b&gt; - corresponds to the next pause time (in milliseconds since 1 Jan 1970) if the Task is currently not paused &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;taskId&lt;/b&gt; - string value corresponding to the Task id&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;ruleId&lt;/b&gt; - string value corresponding to the Rule id if the Task is launched by a Rule trigger&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;user&lt;/b&gt; -string value corresponding to the user name of the user that has launched the Task&lt;li/&gt;&lt;tt&gt;&lt;b&gt;action&lt;/b&gt; - string value corresponding to a textual representation of the searched action&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;displayName&lt;/b&gt; - string value corresponding the Task display name&lt;/tt&gt;&lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only Tasks with the specified properties are returned.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq(state,\&quot;0\&quot;)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq(involvedCount,2),eq(state,0))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt; | 
 **option** | **optional.String**| &lt;p&gt;Possible values for the parameter:&lt;/p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM) | 

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


## TriggerMetadata

> TriggerMetadata(ctx, triggerType)

Retrieve a JSON schema of the required input

Retrieve JSON schema format metadata describing the input required on Rule activation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerType** | **string**|  | 

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


## TriggerTypes

> TriggerTypes(ctx, )

Retrieve all available registered Trigger types

Retrieve all available registered Trigger Provider types.

### Required Parameters

This endpoint does not need any parameter.

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

