# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCalendar**](DeployApi.md#DeleteCalendar) | **Delete** /deploy/calendar/{calendarName} | delete a calendar
[**DeleteConnectionProfile**](DeployApi.md#DeleteConnectionProfile) | **Delete** /deploy/connectionprofile/{server}/{agent}/{type}/{name} | Delete Local Connection Profile
[**DeleteFolder**](DeployApi.md#DeleteFolder) | **Delete** /deploy/folder/{controlM}/{folderName} | delete a folder
[**DeleteLocalConnectionProfile**](DeployApi.md#DeleteLocalConnectionProfile) | **Delete** /deploy/connectionprofile/local/{server}/{agent}/{type}/{name} | Delete Local Connection Profile
[**DeleteSharedConnectionProfile**](DeployApi.md#DeleteSharedConnectionProfile) | **Delete** /deploy/connectionprofile/centralized/{type}/{name} | Delete centralized Connection Profile
[**DeployAiJobtype**](DeployApi.md#DeployAiJobtype) | **Post** /deploy/ai/jobtype | Deploy of Application Integrator job type.
[**DeployFile**](DeployApi.md#DeployFile) | **Post** /deploy | Deploy definitions file
[**DeployJobtypeFile**](DeployApi.md#DeployJobtypeFile) | **Post** /deploy/jobtype | Deploy jobtype
[**GetConnectionProfilesDeploymentStatus**](DeployApi.md#GetConnectionProfilesDeploymentStatus) | **Get** /deploy/connectionprofile/centralized/deploymentstatus/{type}/{name} | Get deployed connection profiles deployment status
[**GetDeployedAiJobtypes**](DeployApi.md#GetDeployedAiJobtypes) | **Get** /deploy/ai/jobtypes | Get Application Integrator job types
[**GetDeployedCalendars**](DeployApi.md#GetDeployedCalendars) | **Get** /deploy/calendars | Get deployed calendars that match the search criteria.
[**GetDeployedConnectionProfiles**](DeployApi.md#GetDeployedConnectionProfiles) | **Get** /deploy/connectionprofiles | Get local deployed connection profiles
[**GetDeployedConnectionProfilesStatus**](DeployApi.md#GetDeployedConnectionProfilesStatus) | **Get** /deploy/connectionprofiles/centralized/status | Get deployed connection profiles status
[**GetDeployedFoldersNew**](DeployApi.md#GetDeployedFoldersNew) | **Get** /deploy/jobs | Get deployed jobs that match the search criteria.
[**GetLocalConnectionProfiles**](DeployApi.md#GetLocalConnectionProfiles) | **Get** /deploy/connectionprofiles/local | Get local deployed connection profiles
[**GetSharedConnectionProfiles**](DeployApi.md#GetSharedConnectionProfiles) | **Get** /deploy/connectionprofiles/centralized | Get centralized deployed connection profile
[**GetSiteStandardFieldRestrictions**](DeployApi.md#GetSiteStandardFieldRestrictions) | **Get** /deploy/sitestandard/{standardName}/fieldRestriction/{fieldName} | Get the allowed values for the specified field in the specified site standard.
[**SetSiteStandardFieldRestrictions**](DeployApi.md#SetSiteStandardFieldRestrictions) | **Post** /deploy/sitestandard/{standardName}/fieldRestriction/{fieldName} | Replace the allowed values for the specified field in the specified site standard.
[**TestConnectionProfile**](DeployApi.md#TestConnectionProfile) | **Post** /deploy/connectionprofile/test | Test connection profile on agent
[**TransformFile**](DeployApi.md#TransformFile) | **Post** /deploy/transform | Transform a definitions file

# **DeleteCalendar**
> SuccessData DeleteCalendar(ctx, calendarName, optional)
delete a calendar

Delete a calendar

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **calendarName** | **string**| The name of the calendar to be deleted. | 
 **optional** | ***DeployApiDeleteCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiDeleteCalendarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **server** | **optional.String**| The name of the server in which the calendar deploy. | [default to Global]

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnectionProfile**
> SuccessData DeleteConnectionProfile(ctx, server, agent, type_, name)
Delete Local Connection Profile

Delete Local Connection Profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The name of the Control-M in which the connection profile is deployed. | 
  **agent** | **string**| The name of the agent the connection profile is deployed on. | 
  **type_** | **string**| The type of connection profile to delete. | 
  **name** | **string**| Name of the Connection Profile | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> SuccessData DeleteFolder(ctx, controlM, folderName)
delete a folder

Delete a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **controlM** | **string**| The name of the Control-M in which the folder(s) are deployed. | 
  **folderName** | **string**| The name of the required folder(s). | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLocalConnectionProfile**
> SuccessData DeleteLocalConnectionProfile(ctx, server, agent, type_, name)
Delete Local Connection Profile

Delete Local Connection Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The name of the Control-M in which the connection profile is deployed. | 
  **agent** | **string**| The name of the agent the connection profile is deployed on. | 
  **type_** | **string**| The type of connection profile to delete. | 
  **name** | **string**| Name of the Connection Profile | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSharedConnectionProfile**
> SuccessData DeleteSharedConnectionProfile(ctx, type_, name)
Delete centralized Connection Profile

Delete centralized Connection Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The type of connection profile to delete. | 
  **name** | **string**| Name of the Connection Profile | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployAiJobtype**
> AiDeployResponse DeployAiJobtype(ctx, ctm, agent, jobTypeId)
Deploy of Application Integrator job type.

Deploy an exsiting Application Integrator job type to agent in order to allow it to run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ctm** | **string**|  | 
  **agent** | **string**|  | 
  **jobTypeId** | **string**|  | 

### Return type

[**AiDeployResponse**](AiDeployResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployFile**
> []DeploymentFileResults DeployFile(ctx, definitionsFile, deployDescriptorFile, additionalConfiguration)
Deploy definitions file

Deploy the provided definition file (JSON, XML or zip) to Control-M

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionsFile** | ***os.File*****os.File**|  | 
  **deployDescriptorFile** | ***os.File*****os.File**|  | 
  **additionalConfiguration** | ***os.File*****os.File**|  | 

### Return type

[**[]DeploymentFileResults**](DeploymentFileResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployJobtypeFile**
> DeployJobtypeResponse DeployJobtypeFile(ctx, definitionsFile, payloadFile, optional)
Deploy jobtype

Deploy the provided jobtype to AI server, EM server, and Agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionsFile** | ***os.File*****os.File**|  | 
  **payloadFile** | ***os.File*****os.File**|  | 
 **optional** | ***DeployApiDeployJobtypeFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiDeployJobtypeFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agent** | **optional.**|  | 
 **server** | **optional.**|  | 

### Return type

[**DeployJobtypeResponse**](DeployJobtypeResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionProfilesDeploymentStatus**
> ConnectionProfilesDeploymentStatusResult GetConnectionProfilesDeploymentStatus(ctx, type_, name)
Get deployed connection profiles deployment status

Get currently deployed connection profiles deployment status according to the search query as JSON.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The type of connection profile such as Database, FileTransfer, Hadoop, Informatica, SAP. | 
  **name** | **string**| Name of the Connection Profile | 

### Return type

[**ConnectionProfilesDeploymentStatusResult**](ConnectionProfilesDeploymentStatusResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployedAiJobtypes**
> AiJobtypeList GetDeployedAiJobtypes(ctx, optional)
Get Application Integrator job types

Get deployed Application Integrator job types that match the requested search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeployApiGetDeployedAiJobtypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiGetDeployedAiJobtypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobTypeName** | **optional.String**| Job type display name ( or partial name ) for query. It accepts * as wildcard. | 
 **jobTypeId** | **optional.String**| Job type id ( or partial name ) for query. It accepts * as wildcard. | 

### Return type

[**AiJobtypeList**](AiJobtypeList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployedCalendars**
> string GetDeployedCalendars(ctx, optional)
Get deployed calendars that match the search criteria.

Get definition of calendars as json code that match the requested search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeployApiGetDeployedCalendarsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiGetDeployedCalendarsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **server** | **optional.String**|  | 
 **type_** | **optional.String**| Calendar type. | 
 **alias** | **optional.String**| Calendar alias name for z/OS servers. | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployedConnectionProfiles**
> string GetDeployedConnectionProfiles(ctx, agent, type_, optional)
Get local deployed connection profiles

Get currently local deployed connection profiles according to the search query as JSON.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agent** | **string**| The name of the agent the connection profile is deployed on | 
  **type_** | **string**| The type of connection profile such as Database, FileTransfer, Hadoop, Informatica, SAP. | 
 **optional** | ***DeployApiGetDeployedConnectionProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiGetDeployedConnectionProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ctm** | **optional.String**| The name of the Control-M in which the connection profile is deployed on | 
 **server** | **optional.String**| The name of the Control-M in which the connection profile is deployed on | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployedConnectionProfilesStatus**
> ConnectionProfilesStatusResult GetDeployedConnectionProfilesStatus(ctx, optional)
Get deployed connection profiles status

Get currently deployed connection profiles status according to the search query as JSON.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeployApiGetDeployedConnectionProfilesStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiGetDeployedConnectionProfilesStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| number to limit the returned connection profiles. If missed - get all | [default to 0]
 **name** | **optional.String**| conn profile name (support *, ?, and comma, default is * for all). | [default to *]
 **type_** | **optional.String**| conn profile type (default is * for accounts from all CMs). | [default to *]

### Return type

[**ConnectionProfilesStatusResult**](ConnectionProfilesStatusResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployedFoldersNew**
> string GetDeployedFoldersNew(ctx, optional)
Get deployed jobs that match the search criteria.

Get definition of jobs and folders (in the desired format - JSON or XML) that match the requested search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeployApiGetDeployedFoldersNewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiGetDeployedFoldersNewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**| Output format (json or xml) | [default to json]
 **folder** | **optional.String**|  | 
 **ctm** | **optional.String**|  | 
 **server** | **optional.String**|  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLocalConnectionProfiles**
> string GetLocalConnectionProfiles(ctx, agent, type_, optional)
Get local deployed connection profiles

Get currently local deployed connection profiles according to the search query as JSON.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agent** | **string**| The name of the agent the connection profile is deployed on | 
  **type_** | **string**| The type of connection profile such as Database, FileTransfer, Hadoop, Informatica, SAP. | 
 **optional** | ***DeployApiGetLocalConnectionProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiGetLocalConnectionProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ctm** | **optional.String**| The name of the Control-M in which the connection profile is deployed on | 
 **server** | **optional.String**| The name of the Control-M in which the connection profile is deployed on | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSharedConnectionProfiles**
> string GetSharedConnectionProfiles(ctx, type_, optional)
Get centralized deployed connection profile

Get currently centralized deployed connection profiles according to the search query as JSON.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The type of connection profile such as Database, FileTransfer, Hadoop, Informatica, SAP. Use * to get all types | [default to *]
 **optional** | ***DeployApiGetSharedConnectionProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiGetSharedConnectionProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| The name of centralized connection profile. Supports for *, ?, and comma. By default is * | [default to *]

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteStandardFieldRestrictions**
> string GetSiteStandardFieldRestrictions(ctx, standardName, fieldName)
Get the allowed values for the specified field in the specified site standard.

Get the allowed values for the specified field in the specified site standard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **standardName** | **string**|  | 
  **fieldName** | **string**|  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSiteStandardFieldRestrictions**
> SuccessData SetSiteStandardFieldRestrictions(ctx, body, standardName, fieldName)
Replace the allowed values for the specified field in the specified site standard.

Replace the allowed values for the specified field in the specified site standard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FieldValues**](FieldValues.md)| The JSON file with the allowed values | 
  **standardName** | **string**|  | 
  **fieldName** | **string**|  | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestConnectionProfile**
> SuccessData TestConnectionProfile(ctx, definitionsFile, optional)
Test connection profile on agent

Test connection profile on agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionsFile** | ***os.File*****os.File**|  | 
 **optional** | ***DeployApiTestConnectionProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiTestConnectionProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ctm** | **optional.**|  | 
 **agent** | **optional.**|  | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransformFile**
> string TransformFile(ctx, definitionsFile, deployDescriptorFile)
Transform a definitions file

Transform the provided definitions file (JSON) according to the provided Deploy Descriptor file (JSON).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionsFile** | ***os.File*****os.File**|  | 
  **deployDescriptorFile** | ***os.File*****os.File**|  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

