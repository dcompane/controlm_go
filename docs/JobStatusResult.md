# JobStatusResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | [**[]JobRunStatus**](JobRunStatus.md) | The list of statuses tracked by the given runId. | [optional] [default to null]
**StartIndex** | **int64** | The index of the first item in the list. | [optional] [default to null]
**ItemsPerPage** | **int64** | The maximum number of items returned by each status request. | [optional] [default to null]
**Returned** | **int64** | The number of the return items by the search. | [optional] [default to null]
**Total** | **int64** | The total number of items. | [optional] [default to null]
**Sort** | **string** | The field the list is sorted by. | [optional] [default to null]
**NextURI** | **string** | URI to get the next items in the list, if any. | [optional] [default to null]
**PrevURI** | **string** | URI to get the previous items in the list, if any. | [optional] [default to null]
**MonitorPageURI** | **string** | A URI to a page displaying the workflow run live. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

