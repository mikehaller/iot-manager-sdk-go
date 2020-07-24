# RuleStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | User-friendly display name of this Rule | 
**RuleId** | **string** | Rule identifier | 
**State** | **int32** |  | [optional] 
**StateDescription** | **string** |  | [optional] 
**InvolvedCount** | **int32** | The count of involved devices with the given execution status. | 
**SuccessCount** | **int32** | The count of involved devices with FINISHED_SUCCESS(0) status. | [optional] 
**WarningCount** | **int32** | The count of involved devices with  FINISHED_WARNING(1) status. | [optional] 
**ErrorCount** | **int32** | The count of involved devices with  FINISHED_ERROR(2) status. | [optional] 
**CancelCount** | **int32** | The count of involved devices with  FINISHED_CANCELED(3) status. | [optional] 
**TriggersCount** | **int32** |  | [optional] 
**NextTimeTrigger** | **int64** |  | [optional] 
**User** | **string** | User | [optional] 
**Props** | **[]string** | Custom properties assigned to this Rule. | [optional] 
**Scope** | [**ScopeInfo**](ScopeInfo.md) |  | 
**Trigger** | [**TriggerInfo**](TriggerInfo.md) |  | [optional] 
**Options** | [**ExecOptInfo**](ExecOptInfo.md) |  | [optional] 
**Action** | [**ActionInfo**](ActionInfo.md) |  | 
**FullStatus** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


