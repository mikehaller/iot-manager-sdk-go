# GroupSelectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryPath** | **string** | Path of the directory whose devices belong to this selection | 
**Recursive** | **bool** | Indicates whether devices from the entire subtree (true) or only direct members (false) of the given directory are included in this selection | [optional] 
**GatewaysOnly** | **bool** | Indicates whether only gateways are included in this selection. | [optional] 
**Tags** | **[]string** | List of tag names to further narrow down the selection - only devices belonging to the given directory and tagged with the given tags are included in this selection. | [optional] 
**RqlQuery** | **string** | RQL filter query which is used to narrow down the devices belonging to this selection. If present, only devices matching this query will be present in the selection. | [optional] 
**FilteringScript** | **string** | A custom Groovy filtering script which is used to narrow down the devices belonging to this selection. If present, only devices satisfying this script will be included in the selection. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


