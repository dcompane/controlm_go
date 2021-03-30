# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildFile**](BuildApi.md#BuildFile) | **Post** /build | Compile definitions file to check its validity

# **BuildFile**
> []DeploymentFileResults BuildFile(ctx, definitionsFile, deployDescriptorFile)
Compile definitions file to check its validity

Compile the provided definition file (JSON or zip) to verify it is valid for Control-M.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definitionsFile** | ***os.File*****os.File**|  | 
  **deployDescriptorFile** | ***os.File*****os.File**|  | 

### Return type

[**[]DeploymentFileResults**](DeploymentFileResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

