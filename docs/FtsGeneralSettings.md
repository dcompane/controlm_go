# FtsGeneralSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HomeDirectory** | **string** | Root path where transferred files are stored. If you want to use a different directory for each logged in user, you must add /${userName} to the path. | [optional] [default to null]
**MultipleLoginAllowed** | **bool** | Allow multiple open sessions | [optional] [default to null]
**MaxOpenSessions** | **int32** | Maximum concurrent open sessions | [optional] [default to null]
**MaxLoginFailures** | **int32** | Number of retries before the server closes the connection | [optional] [default to null]
**DelayAfterLoginFailure** | **int32** | Time of waiting before letting the user to do another login in seconds | [optional] [default to null]
**ThrottlingActivated** | **bool** | Allow bandwidth throttling | [optional] [default to null]
**MaxSimultaneousUploads** | **int32** | Maximum simultaneos uploads | [optional] [default to null]
**MaxSimultaneousDownloads** | **int32** | Maximum simultaneos downloads | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

