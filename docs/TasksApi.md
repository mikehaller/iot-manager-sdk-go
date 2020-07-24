# \TasksApi

All URIs are relative to *https://manager.eu-1.bosch-iot-suite.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Action1**](TasksApi.md#Action1) | **Get** /mme/tasks/{taskId}/action | Retrieve the Task action
[**Cancel**](TasksApi.md#Cancel) | **Put** /mme/tasks/{taskId}/state | Cancel the execution of this Task
[**CancelDeviceInTask**](TasksApi.md#CancelDeviceInTask) | **Put** /mme/tasks/{taskId}/devices/{deviceId}/state | Cancel the execution of this Task over a Device
[**DeleteTask**](TasksApi.md#DeleteTask) | **Delete** /mme/tasks/{taskId} | Delete a Task
[**LaunchTask**](TasksApi.md#LaunchTask) | **Post** /mme/tasks | Launch a Task
[**ListDevicesForTask**](TasksApi.md#ListDevicesForTask) | **Get** /mme/tasks/{taskId}/devices | Retrieve all Devices involved in this Task
[**ListExecutionsForDevice**](TasksApi.md#ListExecutionsForDevice) | **Get** /mme/tasks/devices/{deviceId}/items | Retrieve all Execution Items for a device
[**ListTasksForDevice**](TasksApi.md#ListTasksForDevice) | **Get** /mme/tasks/devices/{deviceId} | Retrieve all Tasks and respective execution status for a particular Device
[**Options1**](TasksApi.md#Options1) | **Get** /mme/tasks/{taskId}/options | Retrieve the execution options of the Task
[**Properties1**](TasksApi.md#Properties1) | **Get** /mme/tasks/{taskId}/properties | Get the custom properties assigned to this Task
[**RetryDevice1**](TasksApi.md#RetryDevice1) | **Put** /mme/tasks/{taskId}/devices/{deviceId} | Retry the Task over a Device
[**RetryDevices1**](TasksApi.md#RetryDevices1) | **Put** /mme/tasks/{taskId}/devices | Retry a Task over a particular Devices
[**Rule1**](TasksApi.md#Rule1) | **Get** /mme/tasks/{taskId}/rule | Retrieve the Rule which launched this Task
[**Scope1**](TasksApi.md#Scope1) | **Get** /mme/tasks/{taskId}/scope | Retrieve the Task scope
[**SetProperties1**](TasksApi.md#SetProperties1) | **Post** /mme/tasks/{taskId}/properties | Set custom Task properties
[**State1**](TasksApi.md#State1) | **Get** /mme/tasks/{taskId}/state | Retrieve the state of the Task
[**Status**](TasksApi.md#Status) | **Get** /mme/tasks/{tasksId}/status | Retrieve a Task status
[**TaskDevice**](TasksApi.md#TaskDevice) | **Get** /mme/tasks/{taskId}/devices/{deviceId} | Retrieve the Device task status in the scope of a Task
[**Tasks1**](TasksApi.md#Tasks1) | **Get** /mme/tasks | Retrieve all Tasks



## Action1

> Action1(ctx, taskId)

Retrieve the Task action

Retrieve the Groovy script defining the action for this Task.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Cancel

> Cancel(ctx, taskId)

Cancel the execution of this Task

Cancel the execution of the Task by the specified <code>taskId</code>. If the Task could not be found or requester has insufficient permissions to access it, а <code>204 No Content</code> will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelDeviceInTask

> CancelDeviceInTask(ctx, taskId, deviceId)

Cancel the execution of this Task over a Device

Cancel the execution of this Task over a Device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 
**deviceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTask

> DeleteTask(ctx, taskId)

Delete a Task

Delete the Task with the specified <code>taskId</code>. If the Task could not be found or requester has insufficient permissions to access it, а <code>204 No Content</code> will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchTask

> TaskStatusInfo LaunchTask(ctx, optional)

Launch a Task

Launch a Task with the specified attributes - <code>displayName</code>, <code>scope</code>, <code>options</code> and <code>action</code>. The displayName, scope, and action are required.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaunchTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LaunchTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restTaskInfo** | [**optional.Interface of RestTaskInfo**](RestTaskInfo.md)| Launch new task | 

### Return type

[**TaskStatusInfo**](TaskStatusInfo.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevicesForTask

> ListDevicesForTask(ctx, taskId, optional)

Retrieve all Devices involved in this Task

Retrieve all Devices which are involved in this Task and have the specified execution state. The <code>stateFilter</code> is a filter on the execution state of the retrieved involved devices.<p><h4>Available stateFilter values:</h4> <ui><li><tt> FINISHED_SUCCESS(0) : Retrieve only Devices that have finished the Task execution with a sucess state</tt></li>  <li><tt> FINISHED_WARNING(1) : Retrieve only Devices that have finished the Task execution with a warning state</tt></li> <li><tt> FINISHED_ERROR(2) : Retrieve only Devices that have finished the Task execution with an error state</tt></li> <li><tt> FINISHED_CANCELED(3) : Retrieve only Devices that have finished the Task execution with a canceled state</tt></li> <li><tt> RUNNING(4) : Retrieve only Devices that have finished the Task execution with a running state</tt></li>  <li><tt>FINISHED(64): This value unites all FINISHED_xxx values and could be used only as parameter for filtering</tt></li><li><tt>ANY(128): Retrieve all involved Devices, regardless of the completion sate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 
 **optional** | ***ListDevicesForTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDevicesForTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stateFilter** | **optional.String**|  | 
 **option** | **optional.String**| &lt;p&gt;Possible values for the parameter:&lt;/p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM) | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecutionsForDevice

> ListExecutionsForDevice(ctx, deviceId, optional)

Retrieve all Execution Items for a device

Retrieve all Execution Items for a specific Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
 **optional** | ***ListExecutionsForDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExecutionsForDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stateFilter** | **optional.String**|  | 
 **option** | **optional.String**| &lt;p&gt;Possible values for the parameter:&lt;/p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM) | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasksForDevice

> ListTasksForDevice(ctx, deviceId, optional)

Retrieve all Tasks and respective execution status for a particular Device

Retrieve all Tasks and respective execution status for a particular Device, which  satisfies the applied state filter. The <code>freeTasks</code> value of <tt>true</tt> specifies that only Task executions not launched by any Rule should be retrieved, whereas the value of <tt>false</tt> specifies that all Task executions will be retrieved (including those launched by Rules).<p><h4>Available stateFilter values:</h4> <ui><li><tt> FINISHED_SUCCESS(0) : Retrieve only Devices that have finished the Task execution with a sucess state</tt></li>  <li><tt> FINISHED_WARNING(1) : Retrieve only Devices that have finished the Task execution with a warning state</tt></li> <li><tt> FINISHED_ERROR(2) : Retrieve only Devices that have finished the Task execution with an error state</tt></li> <li><tt> FINISHED_CANCELED(3) : Retrieve only Devices that have finished the Task execution with a canceled state</tt></li> <li><tt> RUNNING(4) : Retrieve only Devices that have finished the Task execution with a running state</tt></li>  <li><tt>FINISHED(64): This value unites all FINISHED_xxx values and could be used only as parameter for filtering</tt></li><li><tt>ANY(128): Retrieve all involved Devices, regardless of the completion sate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
 **optional** | ***ListTasksForDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTasksForDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stateFilter** | **optional.String**|  | 
 **freeTasks** | **optional.Bool**|  | 
 **option** | **optional.String**| &lt;p&gt;Possible values for the parameter:&lt;/p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM) | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Options1

> Options1(ctx, taskId)

Retrieve the execution options of the Task

Retrieve the execution options of the Task by the specified <code>taskId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Properties1

> Properties1(ctx, taskId)

Get the custom properties assigned to this Task

Get the custom properties assigned to the Task with the specified <code>taskId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryDevice1

> RetryDevice1(ctx, taskId, deviceId)

Retry the Task over a Device

Retry the Task over the Device with the specified <code>deviceId</code>. Retrying forces: <ui><li><tt>partial executions that are not finished to be canceled</tt></li><li><tt>all previous execution info for this Device within the Task to be deleted</tt></li><li><tt>the action script of the Task to be run again for this Device</tt></li>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 
**deviceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryDevices1

> RetryDevices1(ctx, taskId, stateFilter)

Retry a Task over a particular Devices

Retry the Task over the Devices whose status satisfies the applied state filter. Retrying forces: <ui><li><tt>partial executions that are not finished to be canceled</tt></li><li><tt>all previous execution info for these Devices within the Task to be deleted</tt></li><li><tt>the action script of the Task to be run again for these Devices</tt></li>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 
**stateFilter** | [**[]string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rule1

> Rule1(ctx, taskId)

Retrieve the Rule which launched this Task

Retrieve the Rule which launched this Task or <code>204(NO_CONTENT)</code> if Task was manually launched.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Scope1

> Scope1(ctx, taskId)

Retrieve the Task scope

Retrieve the scope of the Task with the specified <code>taskId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetProperties1

> SetProperties1(ctx, taskId, optional)

Set custom Task properties

Set custom properties to the Task with the specified <code>taskId</code>. If <tt> append </tt> is <tt> true</tt>, the supplied properties will be appended/added to any properties currently existing for  this Task, otherwise any existing properties will be fully replaced by the supplied ones.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 
 **optional** | ***SetProperties1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetProperties1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **properties** | [**optional.Interface of Properties**](Properties.md)| Set or add Task properties | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## State1

> string State1(ctx, taskId)

Retrieve the state of the Task

Retrieve the state of the Task with the specified <code>taskId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

**string**

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status

> TaskStatusInfo Status(ctx, tasksId)

Retrieve a Task status

Retrieve a Task status by a specified <code>taskId</code>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tasksId** | **string**|  | 

### Return type

[**TaskStatusInfo**](TaskStatusInfo.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskDevice

> TaskDevice(ctx, taskId, deviceId)

Retrieve the Device task status in the scope of a Task

Retrieve the Device task status of a particular Device in the scope of a Task.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 
**deviceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Tasks1

> Tasks1(ctx, optional)

Retrieve all Tasks

 Retrieve all available Tasks that satisfy the filtering criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Tasks1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Tasks1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| An RQL filter over the Task attributes.&lt;h5&gt;&lt;/tt&gt;&lt;br/&gt;The attribute names in the filtering conditions can be some of:&lt;/h5&gt; &lt;ui&gt;&lt;tt&gt;&lt;li/&gt;&lt;b&gt;state&lt;/b&gt; - integer value corresponding to the Task state as follows: &lt;/tt&gt; &lt;br/&gt;&lt;br/&gt;&lt;tt&gt;1 - RUNNING &lt;/tt&gt;&lt;br/&gt;&lt;tt&gt; 2 - FINISHED&lt;/tt&gt; &lt;br/&gt;&lt;tt&gt;3 - FAILED_TO_LAUNCH&lt;/tt&gt; &lt;li/&gt;&lt;tt&gt;&lt;b&gt;stateDescription&lt;/b&gt; - string value corresponding to the state description &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt; involvedCount&lt;/b&gt; - integer value corresponding to the the number of involved Devices in the Task.&lt;/tt&gt; &lt;li/&gt;&lt;tt&gt;&lt;b&gt; successCount&lt;/b&gt; - integer value corresponding to the the number of successfully finished Devices in the Task&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;warningCount&lt;/b&gt; - integer value corresponding the the number of Devices finished with warning in the Task &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt; errorCount&lt;/b&gt; - integer value corresponding to the the number of Devices finished with error in the Task.&lt;/tt&gt; &lt;li/&gt;&lt;tt&gt;&lt;b&gt;cancelCount&lt;/b&gt; - integer value corresponding to the the number of Devices finished with warning in the Task.&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;launchTime&lt;/b&gt; - long value corresponding to the Task launch time given as milliseconds since 1 Jan 1970&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;finishTime&lt;/b&gt; - long value corresponding to the Task finish time given as milliseconds since 1 Jan 1970&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;isPaused&lt;/b&gt; - boolean value that is &lt;code&gt;true&lt;/code&gt; when the execution is paused due to time constraint restrictions &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;nextResume&lt;/b&gt; - corresponds to resume time (in milliseconds since 1 Jan 1970) if the Task is currently paused &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;nextPause&lt;/b&gt; - corresponds to the next pause time (in milliseconds since 1 Jan 1970) if the Task is currently not paused &lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;taskId&lt;/b&gt; - string value corresponding to the Task id&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;ruleId&lt;/b&gt; - string value corresponding to the Rule id if the Task is launched by a Rule trigger&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;user&lt;/b&gt; -string value corresponding to the user name of the user that has launched the Task&lt;li/&gt;&lt;tt&gt;&lt;b&gt;action&lt;/b&gt; - string value corresponding to a textual representation of the searched action&lt;/tt&gt;&lt;li/&gt;&lt;tt&gt;&lt;b&gt;displayName&lt;/b&gt; - string value corresponding the Task display name&lt;/tt&gt;&lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only Tasks with the specified properties are returned.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq(state,\&quot;0\&quot;)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq(involvedCount,2),eq(state,0))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt; | 
 **option** | **optional.String**| &lt;p&gt;Possible values for the parameter:&lt;/p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM) | 

### Return type

 (empty response body)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

