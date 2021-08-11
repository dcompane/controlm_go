# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelUpgradeActivity**](ProvisionApi.md#CancelUpgradeActivity) | **Post** /provision/upgrade/{upgradeId}/cancel | Cancel upgrade activity
[**DeleteUpgradeActivity**](ProvisionApi.md#DeleteUpgradeActivity) | **Delete** /provision/upgrade/{upgradeId} | Delete upgrade activity status for specific upgrade id.
[**GetAllUpgradeActivitiesStatus**](ProvisionApi.md#GetAllUpgradeActivitiesStatus) | **Get** /provision/upgrades | Get all upgrade activities status.
[**GetDeployVersions**](ProvisionApi.md#GetDeployVersions) | **Get** /provision/upgrades/versions | Get available versions for upgrade.
[**GetEligibleAgentsForUpgrade**](ProvisionApi.md#GetEligibleAgentsForUpgrade) | **Get** /provision/upgrades/agents | Get eligible agents for upgrade that match the requested search criteria.
[**GetImages**](ProvisionApi.md#GetImages) | **Get** /provision/images/{os} | get list of available images for the requested OS
[**GetUpgradeActivityLog**](ProvisionApi.md#GetUpgradeActivityLog) | **Get** /provision/upgrade/{upgradeId}/output | Returns log of upgrade activity.
[**GetUpgradeActivityStatusPerUpgradeId**](ProvisionApi.md#GetUpgradeActivityStatusPerUpgradeId) | **Get** /provision/upgrade/{upgradeId} | Get upgrade activity status for specific upgrade id.
[**RetryUpgradeActivity**](ProvisionApi.md#RetryUpgradeActivity) | **Post** /provision/upgrade/{upgradeId}/retry | Retry upgrade activity
[**TransferAndInstallProduct**](ProvisionApi.md#TransferAndInstallProduct) | **Post** /provision/upgrade/install | Transfer and install a product on an agent
[**UninstallProduct**](ProvisionApi.md#UninstallProduct) | **Post** /provision/upgrade/uninstall | Uninstall a product from an agent

# **CancelUpgradeActivity**
> SuccessData CancelUpgradeActivity(ctx, upgradeId)
Cancel upgrade activity

Cancel upgrade activity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **upgradeId** | **string**| Id of upgrade to cancel | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUpgradeActivity**
> SuccessData DeleteUpgradeActivity(ctx, upgradeId)
Delete upgrade activity status for specific upgrade id.

Delete upgrade activity status for specific upgrade id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **upgradeId** | **string**| The upgrade id. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllUpgradeActivitiesStatus**
> UpgradeRecordList GetAllUpgradeActivitiesStatus(ctx, optional)
Get all upgrade activities status.

Get all upgrade activities status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisionApiGetAllUpgradeActivitiesStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisionApiGetAllUpgradeActivitiesStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctm** | **optional.String**| The Control-M server name | 
 **server** | **optional.String**| The Control-M server name | 
 **agent** | **optional.String**| The Control-M Agent name | 
 **fromVersion** | **optional.String**| Current product version | 
 **toVersion** | **optional.String**| Upgrade to version | 
 **activity** | **optional.String**| Activity type (Install, Uninstall) | 
 **status** | **optional.String**| Upgrade activity status (Cancel, Running, Completed, TransferCompleted, Failed, Unavailable) | 
 **activityName** | **optional.String**| Name of the upgrade activity | 

### Return type

[**UpgradeRecordList**](UpgradeRecordList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployVersions**
> []UpgradeInfo GetDeployVersions(ctx, )
Get available versions for upgrade.

Get available versions for upgrade

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UpgradeInfo**](UpgradeInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEligibleAgentsForUpgrade**
> []UpgradeAgentInfo GetEligibleAgentsForUpgrade(ctx, optional)
Get eligible agents for upgrade that match the requested search criteria.

Get eligible agents for upgrade that match the requested search criteria from Control-M server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisionApiGetEligibleAgentsForUpgradeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisionApiGetEligibleAgentsForUpgradeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| The type (Agent, MFT, AppPack). | 
 **version** | **optional.String**| The version. | 

### Return type

[**[]UpgradeAgentInfo**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImages**
> []string GetImages(ctx, os, optional)
get list of available images for the requested OS

Get a list of the images in the system for the requested OS.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **os** | **string**| The OS name of the requested images. | 
 **optional** | ***ProvisionApiGetImagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisionApiGetImagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| filter according to specific version. | 

### Return type

[**[]string**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeActivityLog**
> string GetUpgradeActivityLog(ctx, upgradeId)
Returns log of upgrade activity.

Returns log of upgrade activity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **upgradeId** | **string**| The upgrade id. | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeActivityStatusPerUpgradeId**
> UpgradeRecord GetUpgradeActivityStatusPerUpgradeId(ctx, upgradeId)
Get upgrade activity status for specific upgrade id.

Get upgrade activity status for specific upgrade id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **upgradeId** | **string**| The upgrade id. | 

### Return type

[**UpgradeRecord**](UpgradeRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetryUpgradeActivity**
> SuccessData RetryUpgradeActivity(ctx, upgradeId)
Retry upgrade activity

Retry upgrade activity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **upgradeId** | **string**| Id of upgrade to retry | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferAndInstallProduct**
> UpgradeResponse TransferAndInstallProduct(ctx, body)
Transfer and install a product on an agent

Transfer and install a product on an agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpgradeRequest**](UpgradeRequest.md)| Upgrade request details | 

### Return type

[**UpgradeResponse**](UpgradeResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UninstallProduct**
> UpgradeResponse UninstallProduct(ctx, body)
Uninstall a product from an agent

Uninstall a product from an agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpgradeRequest**](UpgradeRequest.md)| Rollback request details | 

### Return type

[**UpgradeResponse**](UpgradeResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

