# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateWorkloadPolicy**](RunApi.md#ActivateWorkloadPolicy) | **Post** /run/workloadpolicy/{policy}/activate | activate workload policy
[**AddEvent**](RunApi.md#AddEvent) | **Post** /run/event/{server} | Add a new  event.
[**AddResource**](RunApi.md#AddResource) | **Post** /run/resource/{server} | Add a new pool resource.
[**ConfirmJob**](RunApi.md#ConfirmJob) | **Post** /run/job/{jobId}/confirm | confirm a job
[**DeactivateWorkloadPolicy**](RunApi.md#DeactivateWorkloadPolicy) | **Post** /run/workloadpolicy/{policy}/deactivate | deactivate a workload policy
[**DeleteEvent**](RunApi.md#DeleteEvent) | **Delete** /run/event/{server}/{name}/{date} | Delete a  event.
[**DeleteJob**](RunApi.md#DeleteJob) | **Post** /run/job/{jobId}/delete | mark job as deleted
[**DeleteResource**](RunApi.md#DeleteResource) | **Delete** /run/resource/{server}/{name} | Delete a pool resource.
[**FreeJob**](RunApi.md#FreeJob) | **Post** /run/job/{jobId}/free | free an already held the job
[**GetActiveJob**](RunApi.md#GetActiveJob) | **Get** /run/job/{jobId}/get | get active job
[**GetActiveServices**](RunApi.md#GetActiveServices) | **Get** /run/services/sla | Get SLA active services
[**GetEvents**](RunApi.md#GetEvents) | **Get** /run/events | Get all events records for specific search.
[**GetJobLog**](RunApi.md#GetJobLog) | **Get** /run/job/{jobId}/log | Get job&#x27;s log
[**GetJobOutput**](RunApi.md#GetJobOutput) | **Get** /run/job/{jobId}/output | Get job output
[**GetJobStatistics**](RunApi.md#GetJobStatistics) | **Get** /run/job/{jobId}/statistics | Get job statistics
[**GetJobStatus**](RunApi.md#GetJobStatus) | **Get** /run/job/{jobId}/status | Get status of a job
[**GetJobsStatus**](RunApi.md#GetJobsStatus) | **Get** /run/status/{runId} | Get status of running jobs
[**GetJobsStatusByFilter**](RunApi.md#GetJobsStatusByFilter) | **Get** /run/jobs/status | Get jobs that match the search criteria.
[**GetResources**](RunApi.md#GetResources) | **Get** /run/resources | Get all resources records matching search.
[**GetWaitingInfo**](RunApi.md#GetWaitingInfo) | **Get** /run/job/{jobId}/waitingInfo | get job&#x27;s waiting information
[**GetWorkloadPolicies**](RunApi.md#GetWorkloadPolicies) | **Get** /run/workloadpolicies | get workload policies
[**HoldJob**](RunApi.md#HoldJob) | **Post** /run/job/{jobId}/hold | hold the job so it will not start untill it is freed
[**KillJob**](RunApi.md#KillJob) | **Post** /run/job/{jobId}/kill | Cancel running job
[**ModifyJob**](RunApi.md#ModifyJob) | **Post** /run/job/{jobId}/modify | Modify active job
[**OrderJobsInFolder**](RunApi.md#OrderJobsInFolder) | **Post** /run/order | Execute requested jobs in certain folder
[**RerunJob**](RunApi.md#RerunJob) | **Post** /run/job/{jobId}/rerun | Run job again
[**RunJobs**](RunApi.md#RunJobs) | **Post** /run | Run jobs
[**RunNow**](RunApi.md#RunNow) | **Post** /run/job/{jobId}/runNow | Bypass scheduling cretirias and start the job
[**SetToOK**](RunApi.md#SetToOK) | **Post** /run/job/{jobId}/setToOk | set job end status to OK
[**UndeleteJob**](RunApi.md#UndeleteJob) | **Post** /run/job/{jobId}/undelete | recover a mark for deletion job
[**UpdateAlert**](RunApi.md#UpdateAlert) | **Post** /run/alerts | Update alert.
[**UpdateAlertStatus**](RunApi.md#UpdateAlertStatus) | **Post** /run/alerts/status | Update alert status.
[**UpdateResource**](RunApi.md#UpdateResource) | **Post** /run/resource/{server}/{name} | Update a pool resource.

# **ActivateWorkloadPolicy**
> WorkloadPolicyStateList ActivateWorkloadPolicy(ctx, policy, optional)
activate workload policy

Activate a workload policy, supports wildcard in names

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policy** | **string**| The policy name to be activated. Case sensitive. Wildcards can be used. | 
 **optional** | ***RunApiActivateWorkloadPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunApiActivateWorkloadPolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ctm** | **optional.String**| Optional Control-M Server filter. | 
 **server** | **optional.String**| Optional Control-M Server filter. | 

### Return type

[**WorkloadPolicyStateList**](WorkloadPolicyStateList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddEvent**
> SuccessData AddEvent(ctx, body, server)
Add a new  event.

Add a new  event. date may be of format MMDD, ODAT to set current controlm date, STAT to set no date. default value is ODAT.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventParam**](EventParam.md)| The defined event name. | 
  **server** | **string**| The Control-M Server hosting the event. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddResource**
> SuccessData AddResource(ctx, body, server)
Add a new pool resource.

Add a new pool resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceParam**](ResourceParam.md)| The defined resource name. | 
  **server** | **string**| The Control-M Server hosting the resource. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfirmJob**
> SuccessData ConfirmJob(ctx, jobId)
confirm a job

confirm a job that waits for confirmation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateWorkloadPolicy**
> WorkloadPolicyStateList DeactivateWorkloadPolicy(ctx, policy, optional)
deactivate a workload policy

Deactivate a workload policy, supports wildcard in names

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policy** | **string**| The policy name to be deactivated. Case sensitive. Wildcards can be used. | 
 **optional** | ***RunApiDeactivateWorkloadPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunApiDeactivateWorkloadPolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ctm** | **optional.String**| Optional Control-M Server filter. | 
 **server** | **optional.String**| Optional Control-M Server filter. | 

### Return type

[**WorkloadPolicyStateList**](WorkloadPolicyStateList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEvent**
> SuccessData DeleteEvent(ctx, server, name, date)
Delete a  event.

Delete a  event.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Control-M Server hosting the event. | 
  **name** | **string**| event name | 
  **date** | **string**| event date | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJob**
> SuccessData DeleteJob(ctx, jobId)
mark job as deleted

mark delete as deleted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteResource**
> SuccessData DeleteResource(ctx, server, name)
Delete a pool resource.

Delete a pool resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Control-M Server hosting the resource. | 
  **name** | **string**| Resource name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FreeJob**
> SuccessData FreeJob(ctx, jobId)
free an already held the job

free the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveJob**
> string GetActiveJob(ctx, jobId)
get active job

get the active job's data by job's order ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveServices**
> ActiveServices GetActiveServices(ctx, )
Get SLA active services

Get all SLA active services

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ActiveServices**](ActiveServices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvents**
> []Event GetEvents(ctx, optional)
Get all events records for specific search.

Get all events records for specific search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunApiGetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunApiGetEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctm** | **optional.String**| Control-M Server filter. | 
 **server** | **optional.String**| Control-M Server filter. | 
 **name** | **optional.String**| The event name filter. | 
 **date** | **optional.String**| The event date filter. | 
 **limit** | **optional.Int64**| maximum events to fetch (default 1000). | [default to 1000]

### Return type

[**[]Event**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobLog**
> string GetJobLog(ctx, jobId)
Get job's log

Get the job execution log.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobOutput**
> string GetJobOutput(ctx, jobId, optional)
Get job output

Get the output returned from a job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 
 **optional** | ***RunApiGetJobOutputOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunApiGetJobOutputOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runNo** | **optional.Int64**| The execution number in case of multiple executions (0 will get the last execution&#x27;s output) | [default to 0]

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobStatistics**
> Statistics GetJobStatistics(ctx, jobId)
Get job statistics

Get the statistics from a job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

[**Statistics**](Statistics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobStatus**
> JobRunStatus GetJobStatus(ctx, jobId)
Get status of a job

Get the job status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| Job ID returned from the run status action. | 

### Return type

[**JobRunStatus**](JobRunStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobsStatus**
> JobStatusResult GetJobsStatus(ctx, runId, optional)
Get status of running jobs

Run status of jobs started with the Run service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runId** | **string**| Run ID returned from the run action. | 
 **optional** | ***RunApiGetJobsStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunApiGetJobsStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **optional.Int64**| The index of the job status from which to start. returning results | [default to 0]

### Return type

[**JobStatusResult**](JobStatusResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobsStatusByFilter**
> JobStatusResult GetJobsStatusByFilter(ctx, optional)
Get jobs that match the search criteria.

Get status of jobs that match the requested search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunApiGetJobsStatusByFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunApiGetJobsStatusByFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| maximum jobs status to fetch (default 1000). | [default to 1000]
 **jobname** | **optional.String**|  | 
 **ctm** | **optional.String**|  | 
 **server** | **optional.String**|  | 
 **application** | **optional.String**|  | 
 **subApplication** | **optional.String**|  | 
 **host** | **optional.String**|  | 
 **status** | **optional.String**|  | 
 **folder** | **optional.String**|  | 
 **description** | **optional.String**|  | 
 **jobid** | **optional.String**|  | 
 **neighborhood** | **optional.String**|  | 
 **depth** | **optional.Int32**|  | 
 **direction** | **optional.String**|  | 
 **orderDateFrom** | **optional.String**|  | 
 **orderDateTo** | **optional.String**|  | 
 **fromTime** | **optional.String**|  | 
 **toTime** | **optional.String**|  | 
 **folderLibrary** | **optional.String**|  | 
 **hostGroup** | **optional.String**|  | 
 **runAs** | **optional.String**|  | 
 **command** | **optional.String**|  | 
 **filePath** | **optional.String**|  | 
 **fileName** | **optional.String**|  | 
 **workloadPolicy** | **optional.String**|  | 
 **ruleBasedCalendar** | **optional.String**|  | 
 **resourceMutex** | **optional.String**|  | 
 **resourceSemaphore** | **optional.String**|  | 
 **resourceLock** | **optional.String**|  | 
 **resourcePool** | **optional.String**|  | 
 **held** | **optional.Bool**|  | 
 **folderHeld** | **optional.Bool**|  | 
 **cyclic** | **optional.Bool**|  | 
 **deleted** | **optional.Bool**|  | 

### Return type

[**JobStatusResult**](JobStatusResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResources**
> []ResourceObj GetResources(ctx, optional)
Get all resources records matching search.

Get all resources records matching search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunApiGetResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunApiGetResourcesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctm** | **optional.String**| Control-M Server filter. | 
 **server** | **optional.String**| Control-M Server filter. | 
 **name** | **optional.String**| The resource name filter. | 

### Return type

[**[]ResourceObj**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWaitingInfo**
> []string GetWaitingInfo(ctx, jobId)
get job's waiting information

get the reason why the job is in waiting status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

[**[]string**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkloadPolicies**
> WorkloadPolicyList GetWorkloadPolicies(ctx, optional)
get workload policies

Get all the workload policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunApiGetWorkloadPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunApiGetWorkloadPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **optional.String**| Optionally state filter. Available values Active, Inactive | 

### Return type

[**WorkloadPolicyList**](WorkloadPolicyList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HoldJob**
> SuccessData HoldJob(ctx, jobId)
hold the job so it will not start untill it is freed

hold the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KillJob**
> SuccessData KillJob(ctx, jobId)
Cancel running job

Abort job execution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyJob**
> SuccessData ModifyJob(ctx, jobDefinitionsFile, jobId)
Modify active job

Modify active job, specified by order id according to given definitions file (JSON).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobDefinitionsFile** | ***os.File*****os.File**|  | 
  **jobId** | **string**| The job ID | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderJobsInFolder**
> RunResult OrderJobsInFolder(ctx, optional)
Execute requested jobs in certain folder

Run jobs from selected folder according to given filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunApiOrderJobsInFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunApiOrderJobsInFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OrderFolderParameters**](OrderFolderParameters.md)| parameters to select the jobs to run | 

### Return type

[**RunResult**](RunResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RerunJob**
> JobRunStatus RerunJob(ctx, jobId, optional)
Run job again

Run an already executed job (again).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 
 **optional** | ***RunApiRerunJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunApiRerunJobOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RerunParameters**](RerunParameters.md)| The JSON file with the restart configuration and parameters | 

### Return type

[**JobRunStatus**](JobRunStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunJobs**
> RunResult RunJobs(ctx, jobDefinitionsFile, deployDescriptorFile, additionalConfiguration)
Run jobs

Run jobs according to given definitions file (JSON or zip).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobDefinitionsFile** | ***os.File*****os.File**|  | 
  **deployDescriptorFile** | ***os.File*****os.File**|  | 
  **additionalConfiguration** | ***os.File*****os.File**|  | 

### Return type

[**RunResult**](RunResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunNow**
> SuccessData RunNow(ctx, jobId)
Bypass scheduling cretirias and start the job

start a job immediately

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetToOK**
> SuccessData SetToOK(ctx, jobId)
set job end status to OK

set job status to OK, post processing action

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UndeleteJob**
> SuccessData UndeleteJob(ctx, jobId)
recover a mark for deletion job

recover a mark for deletion job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The job ID | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAlert**
> SuccessData UpdateAlert(ctx, body)
Update alert.

Update alert.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AlertParam**](AlertParam.md)| File that contains the alert propery that want to be update. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAlertStatus**
> SuccessData UpdateAlertStatus(ctx, body)
Update alert status.

Update alert status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AlertStatusParam**](AlertStatusParam.md)| File that contains the alert status propery that want to be update. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResource**
> SuccessData UpdateResource(ctx, body, server, name)
Update a pool resource.

Update a pool resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceMax**](ResourceMax.md)| The defined resource name. | 
  **server** | **string**| The Control-M Server hosting the resource. | 
  **name** | **string**| Resource name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

