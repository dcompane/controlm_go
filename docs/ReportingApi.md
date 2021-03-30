# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReportByName**](ReportingApi.md#GetReportByName) | **Get** /reporting/report/{name} | Retrieves a report by name.
[**GetReportFilters**](ReportingApi.md#GetReportFilters) | **Get** /reporting/reportFilters/{name} | Retrieves report filters
[**GetReportStatus**](ReportingApi.md#GetReportStatus) | **Get** /reporting/status/{reportId} | Retrieves status information for a report generation request based on the report ID
[**RunReport**](ReportingApi.md#RunReport) | **Post** /reporting/report | Run a report

# **GetReportByName**
> ReportResult GetReportByName(ctx, name, optional)
Retrieves a report by name.

Retrieves a report by name in the desired format (CSV,PDF ,EXCEL). If the report is shared, add [shared:] before the name. This REST API command will be deprecated soon.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The report name. | 
 **optional** | ***ReportingApiGetReportByNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportByNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 

### Return type

[**ReportResult**](ReportResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportFilters**
> RunReport GetReportFilters(ctx, name)
Retrieves report filters

Retrieves report filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The report name | 

### Return type

[**RunReport**](RunReport.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportStatus**
> RunReportInfo GetReportStatus(ctx, reportId)
Retrieves status information for a report generation request based on the report ID

Retrieves status information for a report generation request based on the report ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **string**| The ID of the report | 

### Return type

[**RunReportInfo**](RunReportInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunReport**
> RunReportInfo RunReport(ctx, body)
Run a report

Sends a request to generate a report asynchronously and returns the request status. If the report is shared, add [shared:] before the name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RunReport**](RunReport.md)| The report generation parameters | 

### Return type

[**RunReportInfo**](RunReportInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

