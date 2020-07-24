# TaskStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Task identifier | [optional] 
**RuleId** | **string** | Rule identifier | [optional] 
**State** | **int32** | State of this Task | 
**StateDescription** | **string** | Description of the Task state. | 
**LaunchTime** | **int64** | Task launch time | 
**FinishTime** | **int64** | Task finish time | 
**InvolvedCount** | **int32** | Count of involved devices with given execution status in this Task. | 
**SuccessCount** | **int32** | Count of involved devices with FINISHED_SUCCESS(0) status | [optional] 
**WarningCount** | **int32** | Count of involved devices with FINISHED_WARNING(1) status | [optional] 
**ErrorCount** | **int32** | Count of involved devices with FINISHED_ERROR(2) status | [optional] 
**CancelCount** | **int32** | Count of involved devices with FINISHED_CANCELED(3) status | [optional] 
**IsPaused** | **bool** |  | [optional] 
**NextResume** | **int64** | Time at which the execution will be resumed. | [optional] 
**NextPause** | **int64** | Time at which the execution will be paused. | [optional] 
**User** | **string** | User | [optional] 
**Props** | **[]string** | Custom properties assigned to this Task. | [optional] 
**ScopeInfo** | [**ScopeInfo**](ScopeInfo.md) |  | 
**OptInfo** | [**ExecOptInfo**](ExecOptInfo.md) |  | 
**Action** | [**ActionInfo**](ActionInfo.md) |  | 
**DisplayName** | **string** | A User-friendly display name of this Task. | [optional] 
**RuleInfo** | [**RuleInfo**](RuleInfo.md) |  | [optional] 
**Paused** | **bool** |  Task execution is currently paused due to time-constraint restrictions. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


