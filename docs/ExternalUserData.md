# ExternalUserData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | external user name REQUIRED:addExternalUser | HIDDEN | [optional] [default to null]
**Email** | **string** | user email REQUIRED:addExternalUser | HIDDEN | [optional] [default to null]
**Company** | **string** | user&#x27;s company REQUIRED:addExternalUser | HIDDEN | [optional] [default to null]
**Password** | **string** | user password HIDDEN:updateExternalUser | [optional] [default to null]
**Description** | **string** | user description HIDDEN | [optional] [default to null]
**PhoneNumber** | **string** | external user phone number HIDDEN | [optional] [default to null]
**SshKey** | **string** | SSH key string HIDDEN | [optional] [default to null]
**As2Key** | [***As2KeyData**](As2KeyData.md) |  | [optional] [default to null]
**IsLocked** | **bool** | indicates whether the user account is locked HIDDEN | [optional] [default to null]
**LockReason** | **string** | provides the reason for locking the user account HIDDEN | [optional] [default to null]
**ChangePasswordAtNextLogin** | **bool** | indicates whether a password change is required at next login HIDDEN | [optional] [default to null]
**PasswordNeverExpires** | **bool** | indicates whether the user&#x27;s password is set to never expire HIDDEN | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

