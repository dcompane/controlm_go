# FtsFtpSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerEnabled** | **bool** | Enable/Disable listening for FTP/S connection | [optional] [default to null]
**Port** | **int32** | FTP server port | [optional] [default to null]
**AuthenticationMethod** | **string** | Authentication method being used to connect FTP server | [optional] [default to null]
**Secured** | **bool** | Use FTP secure connection (SSL/TLS) | [optional] [default to null]
**KeystoreFilePath** | **string** | FTPS keystore file location | [optional] [default to null]
**KeystoreFilePassword** | **string** | Password being used to access the FTPS keystore | [optional] [default to null]
**Ciphers** | **string** | Ftps server allowed cipher suites (comma-separated). Leave empty to allow all supported cipher suites. | [optional] [default to null]
**ListenForImplicitConnection** | **bool** | Implicit negotiation mode - requires that the entire FTP session must be encrypted | [optional] [default to null]
**PassivePorts** | **string** | Passive FTP ports range - for PASV connections, the server will open a random port in this range for client to connect to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

