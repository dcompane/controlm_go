# ArchiveRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The Control-M Workload Archiving rule name. REQUIRED. HIDDEN. | [optional] [default to null]
**Description** | **string** | The description of Control-M Workload Archiving rule. HIDDEN. | [optional] [default to null]
**MaxOutputSize** | **string** | The maximum job&#x27;s output size to collect. HIDDEN. | [optional] [default to null]
**MaxOutputSizeType** | **string** | The maximum job&#x27;s output size type to collect - KB or MB. The default is MB. HIDDEN. | [optional] [default to null]
**TrimType** | **string** | Trim in case the output exceed the maximum job&#x27;s output - Omit , Beginning, End. The default is to Omit. HIDDEN. | [optional] [default to null]
**Retention** | **string** | The retention period to keep archive job by rule. The default is 1. HIDDEN. | [optional] [default to null]
**RetentionType** | **string** | The retention period type to keep archive job by rule- Years, Months and Days are available. The default is Years. HIDDEN. | [optional] [default to null]
**IsActive** | **string** | Is Control-M Workload Archiving rule is active. HIDDEN. | [optional] [default to null]
**ArchivedType** | **string** | The rule archived data - logs, output or both. The default is both. HIDDEN. | [optional] [default to null]
**RuleParameters** | [**[]RuleCriteria**](RuleCriteria.md) | Rule parameters - ctm, type, jobName, jobType, application, subApplication, jobStatus, folder and library. HIDDEN. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

