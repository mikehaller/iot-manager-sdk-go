# TaskExecOptInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConcurrencyLimit** | **int32** | Restricts the maximum number of devices concurrently executing a Task or Rule. | [default to -1]
**TimeConstraint** | **string** | Time-schedule is permitted for the Task to act. Time schedule constraint supplied as Cron-expression. &lt;code&gt;Example:&lt;/code&gt; &lt;li/&gt;* * 9-17 ? * MON-FRI: defines work time to be on working days (Monday to Friday) between 9.00-17.00h&lt;li/&gt;* * 00-01 * * ?: defines work to be time every day between 00-01 in the morning.&lt;li/&gt;* * * 01 * ?: defines work time to be on the 1st day of each month . | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


