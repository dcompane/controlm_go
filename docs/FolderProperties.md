# FolderProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the folder. REQUIRED:addMFTFolder | HIDDEN | [optional] [default to null]
**AllowedInternalUserNames** | **[]string** | Authorized Internal Users. HIDDEN | [optional] [default to null]
**AllowedUserNames** | **[]string** | Authorized External Users And User Groups. HIDDEN | [optional] [default to null]
**AllowedGroupNames** | **[]string** | Array of allowed group names. HIDDEN | [optional] [default to null]
**DeleteFilesAfterProcessing** | **bool** | Delete file after downloaded from incoming folder. HIDDEN | [optional] [default to null]
**NotifyByEmailWhenFileArrived** | **bool** | Send email notification to external users when a new file arrives. HIDDEN | [optional] [default to null]
**RetentionHours** | **int64** | Retention Time in hours. HIDDEN | [optional] [default to null]
**SizeLimit** | **int64** | Size limit for folder (in Gigabyte). HIDDEN | [optional] [default to null]
**AllowedFilePattern** | **string** | allowed file pattern wildcard. HIDDEN | [optional] [default to null]
**ExcludeFilePattern** | **string** | blocked file pattern wildcard. HIDDEN | [optional] [default to null]
**AccessType** | **string** | Folder&#x27;s access type (Limited, Unlimited). HIDDEN | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

