# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetArchiveJobLog**](ArchiveApi.md#GetArchiveJobLog) | **Get** /archive/{jobId}/log | Get job log
[**GetArchiveJobOutput**](ArchiveApi.md#GetArchiveJobOutput) | **Get** /archive/{jobId}/output | Get job output
[**SearchJobs**](ArchiveApi.md#SearchJobs) | **Get** /archive/search | Search jobs in Archive

# **GetArchiveJobLog**
> string GetArchiveJobLog(ctx, jobId, runNo)
Get job log

Get job log by unique job key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 
  **runNo** | **int64**| The execution number in case of multiple executions | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArchiveJobOutput**
> string GetArchiveJobOutput(ctx, jobId, runNo)
Get job output

Get job output by unique job key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 
  **runNo** | **int64**| The execution number in case of multiple executions | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchJobs**
> ArchiveJobsList SearchJobs(ctx, optional)
Search jobs in Archive

Get all the Control-M Archiving jobs that match the search criterias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ArchiveApiSearchJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArchiveApiSearchJobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| maximum jobs to fetch (default 500). | [default to 500]
 **jobname** | **optional.String**| The name of the job. | 
 **jobid** | **optional.String**|  | 
 **ctm** | **optional.String**| The name of the Control-M server in which the job was ordered from. | 
 **server** | **optional.String**| The name of the Control-M server in which the job was ordered from. | 
 **folder** | **optional.String**| The name of the parent folder. | 
 **fromTime** | **optional.String**| Job execution start date. Date format - YYYY-MM-DD. | 
 **toTime** | **optional.String**| Job execution end date. Date format - YYYY-MM-DD. | 
 **logContains** | **optional.String**| Job log must contain the given phrase. | 
 **outputContains** | **optional.String**| Job output must contain the given phrase. | 
 **application** | **optional.String**| The name of the application the jobs belong to. | 
 **subApplication** | **optional.String**| The name of the sub-application the jobs belong to. | 
 **library** | **optional.String**| The job&#x27;s library name. | 
 **memName** | **optional.String**| Member name. | 
 **memLibrary** | **optional.String**| Member&#x27;s library. | 
 **host** | **optional.String**|  | 
 **hostGroup** | **optional.String**| Job&#x27;s host group. | 
 **runAs** | **optional.String**| Runs as (username on agent machine). | 
 **orderId** | **optional.String**| Job&#x27;s order id. | 
 **status** | **optional.String**| The job&#x27;s end status. | [default to All]
 **orderDateFrom** | **optional.String**| Indicating a date by which will look for jobs that their order date started afterwards. Date format - YYYY-MM-DD. | 
 **orderDateTo** | **optional.String**| Indicating a date by which will look for jobs that their order date ended before. Date format - YYYY-MM-DD. | 
 **numberOfRuns** | **optional.Int64**|  | 

### Return type

[**ArchiveJobsList**](ArchiveJobsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

