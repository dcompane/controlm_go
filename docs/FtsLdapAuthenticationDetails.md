# FtsLdapAuthenticationDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchUserName** | **string** | Name of the user that runs the search action for users that log on | [optional] [default to null]
**SearchUserPassword** | **string** | Password of the user that runs the search action for users that log on | [optional] [default to null]
**ServerUrl** | **string** | URL of the LDAP Directory server | [optional] [default to null]
**BaseDn** | **string** | Base DN (the point from where the server will search for users) | [optional] [default to null]
**UsernameAttributeName** | **string** | Name of the LDAP attribute containing the username | [optional] [default to null]
**DnAttributeName** | **string** | Name of the LDAP attribute containing the distinguished name | [optional] [default to null]
**ConnectionTimeout** | **int32** | LDAP server connection timeout in milliseconds | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

