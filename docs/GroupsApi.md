# \GroupsApi

All URIs are relative to *https://manager.eu-1.bosch-iot-suite.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Attribute1**](GroupsApi.md#Attribute1) | **Get** /di/groups/directories/{path}/attributes/{attributePath} | Retrieve a specific Attribute of a specific device
[**CreateAttributes1**](GroupsApi.md#CreateAttributes1) | **Post** /di/groups/directories/{path}/attributes | Add Attributes of a specific Directory
[**CreateDirectory**](GroupsApi.md#CreateDirectory) | **Post** /di/groups/directories | Add a Directory to the directory hierarchy
[**CreateTag**](GroupsApi.md#CreateTag) | **Post** /di/groups/tags | Create a Tag
[**CreateTagAttributes**](GroupsApi.md#CreateTagAttributes) | **Post** /di/groups/tags/{name}/attributes | Add Attributes of a specific Tag
[**DeleteAttribute1**](GroupsApi.md#DeleteAttribute1) | **Delete** /di/groups/directories/{path}/attributes/{attributePath} | Delete Attribute of a specific Directory
[**DeleteAttributes1**](GroupsApi.md#DeleteAttributes1) | **Delete** /di/groups/directories/{path}/attributes | Delete Attributes of a specific Directory
[**DeleteDirectory**](GroupsApi.md#DeleteDirectory) | **Delete** /di/groups/directories/{path} | Delete an existing directory from the directory hierarchy
[**DeleteTag**](GroupsApi.md#DeleteTag) | **Delete** /di/groups/tags/{name} | Delete a Tag
[**DeleteTagAttribute**](GroupsApi.md#DeleteTagAttribute) | **Delete** /di/groups/tags/{name}/attributes/{attributePath} | Add Attribute of a specific Tag
[**DeleteTagAttributes**](GroupsApi.md#DeleteTagAttributes) | **Delete** /di/groups/tags/{name}/attributes | Add Attributes of a specific Tag
[**Directories**](GroupsApi.md#Directories) | **Get** /di/groups/directories | List all Directories
[**Directory**](GroupsApi.md#Directory) | **Get** /di/groups/directories/{path} | Retrieve an existing Directory from the groups hierarchy
[**DirectoryMembers**](GroupsApi.md#DirectoryMembers) | **Get** /di/groups/directories/{path}/members | Retrieve a list of Members
[**ListAttributes1**](GroupsApi.md#ListAttributes1) | **Get** /di/groups/directories/{path}/attributes | List all Attributes of a specific Directory
[**ListTagAttributes**](GroupsApi.md#ListTagAttributes) | **Get** /di/groups/tags/{name}/attributes | List all Attributes of a specific Tag
[**MembersTag**](GroupsApi.md#MembersTag) | **Get** /di/groups/tags/{name}/members | Retrieve a list of Members
[**MoveDirectory**](GroupsApi.md#MoveDirectory) | **Put** /di/groups/directories/{path}/directory/{targetParentPath} | Move the Directory
[**MoveMember**](GroupsApi.md#MoveMember) | **Put** /di/groups/directories/{path}/members/{memberId} | Move the Device or Gateway to the target directory
[**Parent**](GroupsApi.md#Parent) | **Get** /di/groups/directories/{path}/parent | Retrieve the parent directory of this directory
[**Root**](GroupsApi.md#Root) | **Get** /di/groups/directories/root | Retrieve the Root Directory of the directory hierarchy
[**SetAttribute1**](GroupsApi.md#SetAttribute1) | **Put** /di/groups/directories/{path}/attributes | Add Attribute of a specific Directory
[**SetTagAttribute**](GroupsApi.md#SetTagAttribute) | **Put** /di/groups/tags/{name}/attributes | Add Attribute of a specific Tag
[**Tag**](GroupsApi.md#Tag) | **Get** /di/groups/tags/{name} | Retrieve an existing Tag with a specified name
[**TagAttribute**](GroupsApi.md#TagAttribute) | **Get** /di/groups/tags/{name}/attributes/{attributePath} | Retrieve a specific Attribute of a specific Tag
[**TagMember**](GroupsApi.md#TagMember) | **Put** /di/groups/tags/{name}/{memberId} | Tag/Untag a Member of a Tag
[**Tags**](GroupsApi.md#Tags) | **Get** /di/groups/tags | Retrieve a list of Tags matching the specified filter



## Attribute1

> Attribute Attribute1(ctx, path, attributePath)

Retrieve a specific Attribute of a specific device

Retrieve a specific Attribute of the Directory, identified by the <code>path</code> parameter. The Attribute (JSON) can be referenced hierarchically by applying a JSON Pointer notation (RFC-6901), e.g.: /directories/{path}/attributes/house/room in order to retrieve the house field of an room object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**|  | 
**attributePath** | **string**|  | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAttributes1

> Directory CreateAttributes1(ctx, path, optional)

Add Attributes of a specific Directory

Add Attributes of the specific Directory. All previously existing attributes will be overwritten.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**|  | 
 **optional** | ***CreateAttributes1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAttributes1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributeObject** | [**optional.Interface of []AttributeObject**](AttributeObject.md)| List of Attributes | 

### Return type

[**Directory**](Directory.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDirectory

> Directory CreateDirectory(ctx, optional)

Add a Directory to the directory hierarchy

Add a directory to the directory hierarchy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDirectoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDirectoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directoryParameters** | [**optional.Interface of DirectoryParameters**](DirectoryParameters.md)| Directory | 

### Return type

[**Directory**](Directory.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTag

> Tag CreateTag(ctx, optional)

Create a Tag

Create a tag according to the specified parameters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parameters** | [**optional.Interface of Parameters**](Parameters.md)| Create new tag | 

### Return type

[**Tag**](Tag.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTagAttributes

> Directory CreateTagAttributes(ctx, name, optional)

Add Attributes of a specific Tag

Add Attributes of the specific Tag. All previously existing attributes will be overwritten.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 
 **optional** | ***CreateTagAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTagAttributesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributeObject** | [**optional.Interface of []AttributeObject**](AttributeObject.md)| List of Attributes | 

### Return type

[**Directory**](Directory.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttribute1

> DeleteAttribute1(ctx, path, attributePath)

Delete Attribute of a specific Directory

Delete Attribute of a specific Directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**|  | 
**attributePath** | **string**|  | 

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


## DeleteAttributes1

> DeleteAttributes1(ctx, path)

Delete Attributes of a specific Directory

Delete Attributes of a specific Directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**|  | 

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


## DeleteDirectory

> DeleteDirectory(ctx, path, optional)

Delete an existing directory from the directory hierarchy

Delete an existing directory from the directory tree. If the devices parameter's value is <code>true</code> the devices under the directory will be deleted as well, if the value is <code>false</code>, these devices will still be listed under the root, even after the directory is deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**| The path of a directory in the directory tree | 
 **optional** | ***DeleteDirectoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteDirectoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devices** | **optional.Bool**| If the value of the members&#39; parameter is &lt;code&gt;true&lt;/code&gt; the members under the directory should be deleted as  well. | [default to false]

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


## DeleteTag

> DeleteTag(ctx, name)

Delete a Tag

Delete a tag with the specified name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the tag | 

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


## DeleteTagAttribute

> DeleteTagAttribute(ctx, name, attributePath)

Add Attribute of a specific Tag

Add Attribute of a specific Tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 
**attributePath** | **string**|  | 

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


## DeleteTagAttributes

> DeleteTagAttributes(ctx, name)

Add Attributes of a specific Tag

Add Attributes of a specific Tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

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


## Directories

> []Directory Directories(ctx, optional)

List all Directories

Retrieve a list of directories matching the given filter. Optionally, you can also use field selectors (see parameter fields) to get only the specified fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| &lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only directories with the specified properties are returned. For example, the filter ne(name, \&quot;specifiedName\&quot;) will only return directories that do have the specifiedName name.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({id},\&quot;value\&quot;)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({path}, value)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq({path},\&quot;value\&quot;),eq({name},\&quot;value\&quot;))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;b&gt;Filters containing a wildcard * symbol at the beginning can slow down your search request.&lt;/b&gt; | 
 **option** | **optional.String**| &lt;p&gt;Possible parameter values:&lt;/p&gt;&lt;h5&gt;Sort operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;sort([+|-]{property})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;sort([+|-]{property},[+|-]{property},...)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25.&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;sort(+id)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;sort(-attributes/{attributePropertyValue})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM). | 

### Return type

[**[]Directory**](Directory.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Directory

> Directory Directory(ctx, path)

Retrieve an existing Directory from the groups hierarchy

Retrieve an existing directory from the groups hierarchy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**| The path of a directory in the directory tree | 

### Return type

[**Directory**](Directory.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryMembers

> []Device DirectoryMembers(ctx, path, optional)

Retrieve a list of Members

Retrieve a list of members(Gateways or Devices) that match the given filter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**| The path of a directory in the directory tree | 
 **optional** | ***DirectoryMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DirectoryMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devices** | **optional.Bool**| If the value is &lt;b&gt;true&lt;/b&gt;, the members are devices, if it is &lt;b&gt;false&lt;/b&gt;, the members are gateways. | [default to true]
 **recursive** | **optional.Bool**| If the value is &lt;b&gt;true&lt;/b&gt;, all members will be listed (including sub-directories&#39; gateways or devices), if &lt;b&gt;false&lt;/b&gt;, only members directly under this directory will be listed. | [default to false]
 **filter** | **optional.String**| &lt;h4&gt;Filter operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ne({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;gt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;ge({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;lt({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;le({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;in({property},{value},{value},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;like({property},{value})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists({property})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;Note: When using filter operations, only devices with the specified properties are returned. For example, the filter ne(attributes/owner, \&quot;SID123\&quot;) will only return devices that do have the owner attribute.&lt;h4&gt;Logical operations:&lt;/h4&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;and({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;or({query},{query},...)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;not({query})&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;h5&gt;Examples:&lt;/h5&gt;&lt;ui&gt;&lt;li&gt;&lt;tt&gt;eq(attributes/{attributeId},\&quot;value\&quot;)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;exists(features/{featureId}/properties/{category}/{propertyId}, value)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;and(eq(attributes/{attributeId1},\&quot;value\&quot;),eq(attributes/{attributeId2},\&quot;value\&quot;))&lt;/tt&gt;&lt;/li&gt;&lt;/ui&gt;&lt;b&gt;Filters containing a wildcard * symbol at the beginning can slow down your search request.&lt;/b&gt; | 
 **namespaces** | **optional.String**| A comma separated list of namespaces. This list is used to limit the query to things in the given namespaces only. If this parameter is omitted, all registered namespaces of your solution will be queried.&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?namespaces&#x3D;com.example.namespace&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?namespaces&#x3D;com.example.namespace1,com.example.namespace2&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;deviceId,attributes/location/longitude,attributes/address(city,street)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt; | 
 **option** | **optional.String**| &lt;p&gt;Possible parameter values:&lt;/p&gt;&lt;h5&gt;Sort operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;sort([+|-]{property})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;sort([+|-]{property},[+|-]{property},...)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;Paging operations&lt;/h5&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;size({page-size}) Maximum allowed page size is 200. Default page size is 25.&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor({cursor-id}) Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;sort(+id)&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;sort(-attributes/{attributePropertyValue})&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;size(10) return 10 results&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;cursor(LOREMIPSUM) return results after the position of the cursor LOREMIPSUM.&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;&lt;h5&gt;&lt;b&gt;Combine:&lt;/b&gt;&lt;/h5&gt;&lt;p&gt;If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character. size(200),cursor(LOREMIPSUM). | 
 **fields** | **optional.String**| Contains a comma separated list of fields to be included in the returned JSON. Attributes can be selected in the same manner.&lt;br&gt;&lt;p&gt;&lt;b&gt;Selectable fields&lt;/b&gt;&lt;/p&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;ID&lt;/tt&gt;&lt;/li&gt; &lt;li&gt;&lt;tt&gt;policyId&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;attributes&lt;/tt&gt;&lt;p&gt;Supports selecting arbitrary sub-fields by using a comma separated list:&lt;/p&gt;&lt;ul&gt;&lt;li&gt;several attribute paths can be passed as a comma separated list of JSON pointers (RFC-6901)&lt;br&gt;&lt;br&gt;&lt;b&gt;For example:&lt;/b&gt;&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId}/{attributeValue}&lt;/tt&gt; would select only attribute value (if present)&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}/{attributeValue},attributes/{attributeId2}/{attributeValue}&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}/{attributeValue1},attributes/{attributeId1}/{attributeValue2}&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;li&gt;a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901) separated with /) to select &lt;/li&gt;&lt;li&gt;sub-selectors can be used to request only specific sub-fields by placing expressions in parentheses ( ) after a selected subfield&lt;br&gt;&lt;b&gt;For example:&lt;/b&gt;&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId}/(attributeValue1,attributeValue2)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;features&lt;/tt&gt;&lt;p&gt;Supports selecting arbitrary fields in features similar to attributes&lt;/p&gt;&lt;b&gt;features/{featureId}/properties/{category}/{propertyId}&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;id,attributes,features&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}(attributeValue1,attributeValue1),features&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;id,attributes/{attributeId1},features/{featureId}/properties/{category}/{propertyId}&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**[]Device**](Device.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttributes1

> []Attribute ListAttributes1(ctx, path)

List all Attributes of a specific Directory

The Attributes of the specific Directory were successfully retrieved.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**|  | 

### Return type

[**[]Attribute**](Attribute.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTagAttributes

> []Attribute ListTagAttributes(ctx, name)

List all Attributes of a specific Tag

The Attributes of the specific Tag were successfully retrieved.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

### Return type

[**[]Attribute**](Attribute.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MembersTag

> []Device MembersTag(ctx, name, optional)

Retrieve a list of Members

Retrieve a list of members(Gateways or Devices) that match the given filter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the tag | 
 **optional** | ***MembersTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MembersTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devices** | **optional.Bool**| If the value is &lt;b&gt;true&lt;/b&gt;, the members are devices, if it is &lt;b&gt;false&lt;/b&gt;, the members are gateways. | [default to true]
 **filter** | **optional.String**|  | 
 **namespaces** | **optional.String**|  | 
 **option** | **optional.String**|  | 
 **fields** | **optional.String**| Contains a comma separated list of fields to be included in the returned JSON. Attributes can be selected in the same manner.&lt;br&gt;&lt;p&gt;&lt;b&gt;Selectable fields&lt;/b&gt;&lt;/p&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;ID&lt;/tt&gt;&lt;/li&gt; &lt;li&gt;&lt;tt&gt;policyId&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;attributes&lt;/tt&gt;&lt;p&gt;Supports selecting arbitrary sub-fields by using a comma separated list:&lt;/p&gt;&lt;ul&gt;&lt;li&gt;several attribute paths can be passed as a comma separated list of JSON pointers (RFC-6901)&lt;br&gt;&lt;br&gt;&lt;b&gt;For example:&lt;/b&gt;&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId}/{attributeValue}&lt;/tt&gt; would select only attribute value (if present)&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}/{attributeValue},attributes/{attributeId2}/{attributeValue}&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}/{attributeValue1},attributes/{attributeId1}/{attributeValue2}&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;li&gt;a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901) separated with /) to select &lt;/li&gt;&lt;li&gt;sub-selectors can be used to request only specific sub-fields by placing expressions in parentheses ( ) after a selected subfield&lt;br&gt;&lt;b&gt;For example:&lt;/b&gt;&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId}/(attributeValue1,attributeValue2)&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;features&lt;/tt&gt;&lt;p&gt;Supports selecting arbitrary fields in features similar to attributes&lt;/p&gt;&lt;b&gt;features/{featureId}/properties/{category}/{propertyId}&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;&lt;h4&gt;Examples:&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;id,attributes,features&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;attributes/{attributeId1}(attributeValue1,attributeValue1),features&lt;/tt&gt;&lt;/li&gt;&lt;li&gt;&lt;tt&gt;?fields&#x3D;id,attributes/{attributeId1},features/{featureId}/properties/{category}/{propertyId}&lt;/tt&gt;&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**[]Device**](Device.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveDirectory

> MoveDirectory(ctx, path, targetParentPath)

Move the Directory

Moves the directory with a specific identifier to a target directory with the specified identifier.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**| The path of a directory in the directory tree | 
**targetParentPath** | **string**| The path of the parent directory, the directory will be moved to | 

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


## MoveMember

> MoveMember(ctx, path, memberId)

Move the Device or Gateway to the target directory

Move a Device or Gateway with a specific ID to the target directory with the specified ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**| The path of the parent directory, the device will be moved to | 
**memberId** | **string**|  | 

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


## Parent

> Directory Parent(ctx, path)

Retrieve the parent directory of this directory

Retrieve the parent directory of this directory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**| The path of a directory in the directory tree | 

### Return type

[**Directory**](Directory.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Root

> Directory Root(ctx, )

Retrieve the Root Directory of the directory hierarchy

Retrieve the root directory of the directory hierarchy. All directories, that have no parent directory assigned, are subdirectories of this directory.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Directory**](Directory.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAttribute1

> Directory SetAttribute1(ctx, path, optional)

Add Attribute of a specific Directory

Add Attribute of the specific Directory. If an attribute with the same path already exists it will be overwritten.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**|  | 
 **optional** | ***SetAttribute1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetAttribute1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attribute** | [**optional.Interface of Attribute**](Attribute.md)| Attribute | 

### Return type

[**Directory**](Directory.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTagAttribute

> Directory SetTagAttribute(ctx, name, optional)

Add Attribute of a specific Tag

Add Attribute of the specific Tag. If an attribute with the same path already exists it will be overwritten. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 
 **optional** | ***SetTagAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetTagAttributeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attribute** | [**optional.Interface of Attribute**](Attribute.md)| Attribute | 

### Return type

[**Directory**](Directory.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Tag

> Tag Tag(ctx, name)

Retrieve an existing Tag with a specified name

Retrieve an existing Tag with a specified name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the tag | 

### Return type

[**Tag**](Tag.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagAttribute

> Attribute TagAttribute(ctx, name, attributePath)

Retrieve a specific Attribute of a specific Tag

Retrieve a specific Attribute of the Tag, identified by name. The Attribute (JSON) can be referenced hierarchically by applying a JSON Pointer notation (RFC-6901), e.g.: /tags/{name}/attributes/house/room in order to retrieve the house field of an room object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 
**attributePath** | **string**|  | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagMember

> TagMember(ctx, name, memberId, optional)

Tag/Untag a Member of a Tag

Tag/Untag a member with a specific ID i.e. attach/detach the tag with the specified ID from the member(gateway or device).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the tag | 
**memberId** | **string**| The memberId(of a device or gateway) | 
 **optional** | ***TagMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TagMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **optional.Bool**| If the value is &lt;b&gt;true&lt;/b&gt;, attaches the tag with the specified ID to the device, if &lt;b&gt;false&lt;/b&gt;, detaches the tag with the specified ID from the device. | [default to true]

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


## Tags

> []Tag Tags(ctx, optional)

Retrieve a list of Tags matching the specified filter

Retrieve a list of Tags matching the specified filter

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**|  | 
 **option** | **optional.String**|  | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth), [refreshToken](../README.md#refreshToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

