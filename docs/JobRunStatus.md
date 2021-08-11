# JobRunStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** | Order ID of the job. | [default to null]
**FolderId** | **string** | Order ID of the folder containing this job. | [optional] [default to null]
**NumberOfRuns** | **int64** | The run number (in case of cyclic jobs or reruns) | [optional] [default to null]
**Name** | **string** | The name of the run job. | [optional] [default to null]
**Folder** | **string** | The name of the run job. | [optional] [default to null]
**Type_** | **string** | The type of the run job. | [optional] [default to null]
**Status** | **string** | The status of the run job. | [optional] [default to null]
**Held** | **bool** | Is job held. | [optional] [default to null]
**Deleted** | **bool** | Is job held. | [optional] [default to null]
**Cyclic** | **bool** | Is it a cyclic job. | [optional] [default to null]
**StartTime** | **string** | The start time of the job run. | [optional] [default to null]
**EndTime** | **string** | The end time of the job run. | [optional] [default to null]
**EstimatedStartTime** | **[]string** | The estimated start time of the jobs. | [optional] [default to null]
**EstimatedEndTime** | **[]string** | The estimated end time of the jobs. | [optional] [default to null]
**OrderDate** | **string** | The order date. | [optional] [default to null]
**Ctm** | **string** | The controlm server. | [optional] [default to null]
**Description** | **string** | The job description. | [optional] [default to null]
**Host** | **string** | host machine where the job runs. | [optional] [default to null]
**Application** | **string** | job application. | [optional] [default to null]
**SubApplication** | **string** | job subApplication. | [optional] [default to null]
**JobJSON** | **string** | The JSON string that describes the job. | [optional] [default to null]
**OutputURI** | **string** | A URI that can be used to get the output of the run job | [optional] [default to null]
**LogURI** | **string** | A URI that can be used to get the log of the run job | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

