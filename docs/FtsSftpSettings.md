# FtsSftpSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerEnabled** | **bool** | Enable/Disable listening for SFTP connection | [optional] [default to null]
**Port** | **int32** | SFTP server port | [optional] [default to null]
**AuthenticationMethod** | **string** | Authentication method being used to connect FTP server | [optional] [default to null]
**KeystoreFilePath** | **string** | SFTP keystore file location | [optional] [default to null]
**KeystoreFilePassword** | **string** | Password being used to access the SFTP keystore | [optional] [default to null]
**Ciphers** | **string** | Ftps server allowed cipher suites (comma-separated). Leave empty to allow all supported cipher suites. | [optional] [default to null]
**KnownUsersFilePath** | **string** | Known users file location | [optional] [default to null]
**OverriddenUsersHomeDirectories** | [**[]FtsUserHomeDirectoryData**](FtsUserHomeDirectoryData.md) | Overridden home directories for specific internal users | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

