# ExecOptInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverlapOption** | **int32** |   Defines whether multiple Tasks over the same device can overlap or should be replaced. This option is applicable only for Rules. Overlapping: &lt;li&gt; ALLOW_OVERLAPPING_TASKS(0): No relation constraint between Tasks. New Tasks can be triggered, although old Tasks are still active. This option should be used very carefully due to possible accumulation (explosion) of unfinished Tasks. &lt;/li&gt; &lt;li&gt;NO_OVERLAPPING_CANCEL_OLD(1) : In case of Rule firing for a particular device, any previous executions still running for this device in the scope of the same Rule will be canceled. &lt;/li&gt;&lt;li&gt; NO_OVERLAPPING_KEEP_OLD_SKIP_NEW(2): Skip triggering a new Task for devices that already have unfinished Tasks in the scope of the Rule.; | [optional] [default to OVERLAP_OPTION__1]
**ConcurrencyLimit** | **int32** | Restricts the maximum number of devices concurrently executing a Task or Rule. | [default to -1]
**ConcurrencyTimeout** | **int32** | Timeout in seconds after which a non-confirmed execution expires and stops occupying room in the concurrency limit. The default value is 30 seconds. | [optional] 
**OncePerDevice** | **bool** | If the Rule should be triggered no more than once for the same device, false otherwise | [optional] [default to true]
**AutoDisable** | **bool** | If the Rule should be automatically disabled once all targets from the Scope pass through the triggering conditions, false otherwise | [optional] [default to true]
**TimeConstraint** | **string** | A time schedule in which the Task execution is permitted. The time schedule constraint must be provided as a cron expression. &lt;code&gt;Example:&lt;/code&gt; &lt;li/&gt;* * 9-17 ? * MON-FRI: defines that the execution time will be on working days (Monday to Friday) between 9.00-17.00h&lt;li/&gt;* * 00-01 * * ?: defines that execution time will be every day between 00-01 in the morning.&lt;li/&gt;* * * 01 * ?: defines that execution time will be on the 1st day of each month . | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


