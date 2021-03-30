# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoLogin**](SessionApi.md#DoLogin) | **Post** /session/login | login user to Control-M
[**DoLogout**](SessionApi.md#DoLogout) | **Post** /session/logout | logout user from Control-M
[**UpdateMyPassword**](SessionApi.md#UpdateMyPassword) | **Post** /session/user/password/update | Change my password

# **DoLogin**
> LoginResult DoLogin(ctx, body)
login user to Control-M

Authenticate the user with the specified password and return a token that can be used later in subsequent requests to access Control-M.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoginCredentials**](LoginCredentials.md)| The credentials to use for login. | 

### Return type

[**LoginResult**](LoginResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoLogout**
> SuccessData DoLogout(ctx, )
logout user from Control-M

Disconnects the user session specified by the request authentication token, and removes it from the server memory.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMyPassword**
> SuccessData UpdateMyPassword(ctx, body)
Change my password

Change my password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PasswordsObject**](PasswordsObject.md)| The new password. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

