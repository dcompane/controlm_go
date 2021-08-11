# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddArchiveRule**](ConfigApi.md#AddArchiveRule) | **Post** /config/archive/rule | Add Workload Archiving rule
[**AddExternalUser**](ConfigApi.md#AddExternalUser) | **Post** /config/mft/externaluser | Add and external user
[**AddExternalUserOrUserGroupToMFTFolder**](ConfigApi.md#AddExternalUserOrUserGroupToMFTFolder) | **Post** /config/mft/virtualfolder/{folderName}/user/{userOrGroup} | Add external user or user groups to virtual folder external users list.
[**AddGateway**](ConfigApi.md#AddGateway) | **Post** /config/mft/gateway | add gateway.
[**AddHostToHostgroup**](ConfigApi.md#AddHostToHostgroup) | **Post** /config/server/{server}/hostgroup/{hostgroup}/agent | add agent to hostgroup
[**AddHubToCluster**](ConfigApi.md#AddHubToCluster) | **Post** /config/mft/cluster/hub/{agentname} | add hub to cluster.
[**AddMFTFolder**](ConfigApi.md#AddMFTFolder) | **Post** /config/mft/virtualfolder | Add virtual folder
[**AddMftUserGroup**](ConfigApi.md#AddMftUserGroup) | **Post** /config/mft/usergroup | Add user group.
[**AddPgpTemplate**](ConfigApi.md#AddPgpTemplate) | **Post** /config/server/{server}/agent/{agent}/mft/pgptemplate/{templateName} | Add PGP Template
[**AddRemoteHost**](ConfigApi.md#AddRemoteHost) | **Post** /config/server/{server}/remotehost | add remote host to Server
[**AddRole**](ConfigApi.md#AddRole) | **Post** /config/authorization/role | Add Authorization Role
[**AddRoleToLdapGroup**](ConfigApi.md#AddRoleToLdapGroup) | **Post** /config/authorization/ldap/{ldapgroup}/role/{role} | Add a role to LDAP group
[**AddRoleToUser**](ConfigApi.md#AddRoleToUser) | **Post** /config/authorization/user/{user}/role/{role} | Add a role to user
[**AddSecret**](ConfigApi.md#AddSecret) | **Post** /config/secret | Add a new secret
[**AddServer**](ConfigApi.md#AddServer) | **Post** /config/server | add server to the system
[**AddUser**](ConfigApi.md#AddUser) | **Post** /config/authorization/user | Add user
[**AddZosTemplate**](ConfigApi.md#AddZosTemplate) | **Post** /config/server/{server}/agent/{agent}/mft/zostemplate/{templateName} | Add z/OS Template
[**AuthorizeMftSshCluster**](ConfigApi.md#AuthorizeMftSshCluster) | **Post** /config/server/{server}/agent/{agent}/mft/ssh/cluster/{clusterName} | Authorize SSH Cluster
[**AuthorizeMftSshHost**](ConfigApi.md#AuthorizeMftSshHost) | **Post** /config/server/{server}/agent/{agent}/mft/ssh/host/{hostname} | Authorize SSH Host
[**AuthorizeSshKnownRemotehost**](ConfigApi.md#AuthorizeSshKnownRemotehost) | **Post** /config/server/{server}/remotehost/{remotehost}/authorize | Authorize
[**ChangeUserPassword**](ConfigApi.md#ChangeUserPassword) | **Post** /config/user/{user}/password/adminUpdate | Change user password
[**CreateAgentCertificateSigningRequest**](ConfigApi.md#CreateAgentCertificateSigningRequest) | **Post** /config/server/{server}/agent/{agent}/csr | Create certificate signing request (CSR).
[**CreateRunAsUser**](ConfigApi.md#CreateRunAsUser) | **Post** /config/server/{server}/runasuser | Add a new Run-as user
[**DeleteAgent**](ConfigApi.md#DeleteAgent) | **Delete** /config/server/{server}/agent/{agent} | delete an agent from Server
[**DeleteArchiveRule**](ConfigApi.md#DeleteArchiveRule) | **Delete** /config/archive/rule/{ruleName} | Delete Workload Archiving rule
[**DeleteAuthorizationRole**](ConfigApi.md#DeleteAuthorizationRole) | **Delete** /config/authorization/role/{role} | Delete Authorization Role
[**DeleteExternalUser**](ConfigApi.md#DeleteExternalUser) | **Delete** /config/mft/externaluser/{username} | Delete an external user
[**DeleteExternalUserOrUserGroupFromMFTFolder**](ConfigApi.md#DeleteExternalUserOrUserGroupFromMFTFolder) | **Delete** /config/mft/virtualfolder/{folderName}/user/{userOrGroup} | Remove an external user or user group from an existing virtual folder in MFT.
[**DeleteHostFromGroup**](ConfigApi.md#DeleteHostFromGroup) | **Delete** /config/server/{server}/hostgroup/{hostgroup}/agent/{host} | delete an agent from a hostgroup
[**DeleteHostGroup**](ConfigApi.md#DeleteHostGroup) | **Delete** /config/server/{server}/hostgroup/{hostgroup} | delete host group
[**DeleteMFTFolder**](ConfigApi.md#DeleteMFTFolder) | **Delete** /config/mft/virtualfolder/{folderName} | Delete a virtual folder.
[**DeleteMftUserGroup**](ConfigApi.md#DeleteMftUserGroup) | **Delete** /config/mft/usergroup/{name} | Delete user group.
[**DeletePgpTemplate**](ConfigApi.md#DeletePgpTemplate) | **Delete** /config/server/{server}/agent/{agent}/mft/pgptemplate/{templateName} | Delete PGP Template
[**DeleteRemoteHost**](ConfigApi.md#DeleteRemoteHost) | **Delete** /config/server/{server}/remotehost/{remotehost} | delete a remote host from Server
[**DeleteRoleFromLdapGroup**](ConfigApi.md#DeleteRoleFromLdapGroup) | **Delete** /config/authorization/ldap/{ldapgroup}/role/{role} | Delete a role from LDAP group
[**DeleteRunAsUser**](ConfigApi.md#DeleteRunAsUser) | **Delete** /config/server/{server}/runasuser/{agent}/{user} | delete Run-as user
[**DeleteSecret**](ConfigApi.md#DeleteSecret) | **Delete** /config/secret/{name} | Delete an existing secret
[**DeleteUser**](ConfigApi.md#DeleteUser) | **Delete** /config/authorization/user/{user} | Delete user
[**DeleteZosTemplate**](ConfigApi.md#DeleteZosTemplate) | **Delete** /config/server/{server}/agent/{agent}/mft/zostemplate/{templateName} | Delete z/OS Template
[**DeployAgentCertificate**](ConfigApi.md#DeployAgentCertificate) | **Post** /config/server/{server}/agent/{agent}/crt | Deploy certificate (CRT).
[**DisableAgent**](ConfigApi.md#DisableAgent) | **Post** /config/server/{server}/agent/{agent}/disable | disable agent from the Server
[**EnableAgent**](ConfigApi.md#EnableAgent) | **Post** /config/server/{server}/agent/{agent}/enable | enable agent from the Server
[**Failover**](ConfigApi.md#Failover) | **Put** /config/server/{server}/failover | Perform Manual Failover on a specified Server
[**GenerateMftRsaSshKey**](ConfigApi.md#GenerateMftRsaSshKey) | **Post** /config/server/{server}/agent/{agent}/mft/ssh/key | Generate RSA SSH Key
[**GetAgentCertificateExpirationDate**](ConfigApi.md#GetAgentCertificateExpirationDate) | **Get** /config/server/{server}/agent/{agent}/crt/expiration | Get certificate expiration date.
[**GetAgentParameters**](ConfigApi.md#GetAgentParameters) | **Get** /config/server/{server}/agent/{agent}/params | get agent parameters
[**GetAgents**](ConfigApi.md#GetAgents) | **Get** /config/server/{server}/agents | get Server agents
[**GetAllArchiveRules**](ConfigApi.md#GetAllArchiveRules) | **Get** /config/archive/rules | Get all Workload Archiving rules
[**GetAllAuthorizationRoles**](ConfigApi.md#GetAllAuthorizationRoles) | **Get** /config/authorization/roles | Get Authorization Roles
[**GetAllLdapGroups**](ConfigApi.md#GetAllLdapGroups) | **Get** /config/authorization/ldaps | Get All LDAP groups
[**GetAllRolesAssociatedWithLdap**](ConfigApi.md#GetAllRolesAssociatedWithLdap) | **Get** /config/authorization/ldap/{ldapgroup}/roles | Get Authorization Roles associated with an LDAP group
[**GetAllUsers**](ConfigApi.md#GetAllUsers) | **Get** /config/authorization/users | Get users
[**GetArchiveStatistics**](ConfigApi.md#GetArchiveStatistics) | **Get** /config/archive/statistics | Get Workload Archiving statistics
[**GetExternalUserAuthorizedFolders**](ConfigApi.md#GetExternalUserAuthorizedFolders) | **Get** /config/mft/externaluser/{name}/virtualfolders | Get MFT external user authorized folders
[**GetExternalUsers**](ConfigApi.md#GetExternalUsers) | **Get** /config/mft/externalusers | Get MFT external users that match the search criteria.
[**GetFtsSettings**](ConfigApi.md#GetFtsSettings) | **Get** /config/server/{server}/agent/{agent}/mft/fts/settings | Get File Transfer Server (FTS) configuration data.
[**GetHostgroups**](ConfigApi.md#GetHostgroups) | **Get** /config/server/{server}/hostgroups | get Server hostgroups
[**GetHostsInGroup**](ConfigApi.md#GetHostsInGroup) | **Get** /config/server/{server}/hostgroup/{hostgroup}/agents | get hostgroup agents
[**GetHubStatusDetails**](ConfigApi.md#GetHubStatusDetails) | **Get** /config/mft/hub/{nodeId}/status | Get hub status.
[**GetMFTFolders**](ConfigApi.md#GetMFTFolders) | **Get** /config/mft/virtualfolders | Get MFT virtual folders that match the search criteria.
[**GetMftConfiguration**](ConfigApi.md#GetMftConfiguration) | **Get** /config/server/{server}/agent/{agent}/mft/configuration | Get MFT Configuration
[**GetMftGateways**](ConfigApi.md#GetMftGateways) | **Get** /config/mft/gateways | Get MFT gateways
[**GetMftUserGroups**](ConfigApi.md#GetMftUserGroups) | **Get** /config/mft/usergroups | Get all user groups that match the search criteria.
[**GetPgpTemplates**](ConfigApi.md#GetPgpTemplates) | **Get** /config/server/{server}/agent/{agent}/mft/pgptemplates | Get PGP Templates
[**GetRemoteHostProperties**](ConfigApi.md#GetRemoteHostProperties) | **Get** /config/server/{server}/remotehost/{remotehost} | get a remote host configuration from Server
[**GetRemoteHosts**](ConfigApi.md#GetRemoteHosts) | **Get** /config/server/{server}/remotehosts | get Server remote hosts
[**GetRole**](ConfigApi.md#GetRole) | **Get** /config/authorization/role/{role} | Get Authorization Role
[**GetRoleAssociates**](ConfigApi.md#GetRoleAssociates) | **Get** /config/authorization/role/{role}/associates | Get all authorization entities associated with role
[**GetRunAsUser**](ConfigApi.md#GetRunAsUser) | **Get** /config/server/{server}/runasuser/{agent}/{user} | Get Run-as user
[**GetRunAsUsersList**](ConfigApi.md#GetRunAsUsersList) | **Get** /config/server/{server}/runasusers | Get Run-as user list that match the requested search criteria.
[**GetServerParameters**](ConfigApi.md#GetServerParameters) | **Get** /config/server/{server}/params | get Server parameters
[**GetServers**](ConfigApi.md#GetServers) | **Get** /config/servers | get all the Servers name and hostname in the system
[**GetUser**](ConfigApi.md#GetUser) | **Get** /config/authorization/user/{user} | Get user
[**GetUserEffectiveRights**](ConfigApi.md#GetUserEffectiveRights) | **Get** /config/authorization/user/effectiveRights | Get user real effective authorizations
[**GetZosTemplates**](ConfigApi.md#GetZosTemplates) | **Get** /config/server/{server}/agent/{agent}/mft/zostemplates | Get z/OS Templates
[**ListSecrets**](ConfigApi.md#ListSecrets) | **Get** /config/secrets | Get list of secret names
[**PingAgent**](ConfigApi.md#PingAgent) | **Post** /config/server/{server}/agent/{agent}/ping | ping to the agent in the Server
[**RecycleItem**](ConfigApi.md#RecycleItem) | **Post** /config/item/{id}/recycle | recycle item
[**RemoveControlmServer**](ConfigApi.md#RemoveControlmServer) | **Delete** /config/server/{server} | Delete Server
[**RemoveGateway**](ConfigApi.md#RemoveGateway) | **Delete** /config/mft/gateway/{gatewayName} | remove gateway.
[**RemoveHubFromCluster**](ConfigApi.md#RemoveHubFromCluster) | **Delete** /config/mft/cluster/hub/{agentname} | remove hub from cluster.
[**RemoveRoleFromUser**](ConfigApi.md#RemoveRoleFromUser) | **Delete** /config/authorization/user/{user}/role/{role} | Remove a role from a user
[**SendArchiveCleanupRequest**](ConfigApi.md#SendArchiveCleanupRequest) | **Delete** /config/archive/cleanup | Deletes data (jobs including outputs and logs) from the Workload Archiving database.
[**SetAgentParameter**](ConfigApi.md#SetAgentParameter) | **Post** /config/server/{server}/agent/{agent}/param/{name} | set agent parameter
[**SetSystemParam**](ConfigApi.md#SetSystemParam) | **Post** /config/em/param/{name} | set value of a an em system parameter
[**Setasprimary**](ConfigApi.md#Setasprimary) | **Put** /config/server/{server}/setasprimary | Set secondary server as Primary on a specified Server
[**TestRunAsUser**](ConfigApi.md#TestRunAsUser) | **Post** /config/server/{server}/runasuser/{agent}/{user}/test | Test existed Run-as user
[**UpdateArchiveRule**](ConfigApi.md#UpdateArchiveRule) | **Post** /config/archive/rule/{ruleName} | Edit Workload Archiving rule
[**UpdateExternalUser**](ConfigApi.md#UpdateExternalUser) | **Post** /config/mft/externaluser/{username} | Update an external user
[**UpdateFtsSettings**](ConfigApi.md#UpdateFtsSettings) | **Post** /config/server/{server}/agent/{agent}/mft/fts/settings | Update File Transfer Server (FTS) configuration data.
[**UpdateHostsInHostgroup**](ConfigApi.md#UpdateHostsInHostgroup) | **Post** /config/server/{server}/hostgroup/{hostgroup} | update agents in hostgroup.
[**UpdateMFTFolder**](ConfigApi.md#UpdateMFTFolder) | **Post** /config/mft/virtualfolder/{folderName} | Update an existing virtual folder in MFT.
[**UpdateMftConfiguration**](ConfigApi.md#UpdateMftConfiguration) | **Post** /config/server/{server}/agent/{agent}/mft/configuration | Update MFT Configuration
[**UpdateMftUserGroup**](ConfigApi.md#UpdateMftUserGroup) | **Post** /config/mft/usergroup/{name} | Update user group.
[**UpdatePgpTemplate**](ConfigApi.md#UpdatePgpTemplate) | **Put** /config/server/{server}/agent/{agent}/mft/pgptemplate/{templateName} | Update PGP Template
[**UpdateRole**](ConfigApi.md#UpdateRole) | **Post** /config/authorization/role/{role} | Update Authorization Role
[**UpdateRunAsUser**](ConfigApi.md#UpdateRunAsUser) | **Post** /config/server/{server}/runasuser/{agent}/{user} | Update Run-as user
[**UpdateSecret**](ConfigApi.md#UpdateSecret) | **Post** /config/secret/{name} | Update an existing secret
[**UpdateUser**](ConfigApi.md#UpdateUser) | **Post** /config/authorization/user/{user} | Update user
[**UpdateZosTemplate**](ConfigApi.md#UpdateZosTemplate) | **Put** /config/server/{server}/agent/{agent}/mft/zostemplate/{templateName} | Update z/OS Template

# **AddArchiveRule**
> SuccessData AddArchiveRule(ctx, body)
Add Workload Archiving rule

Add a new Workload Archiving rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ArchiveRule**](ArchiveRule.md)| archive rule details to add | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddExternalUser**
> SuccessData AddExternalUser(ctx, body)
Add and external user

Add and external user for b2b

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExternalUserData**](ExternalUserData.md)| External user data | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddExternalUserOrUserGroupToMFTFolder**
> SuccessData AddExternalUserOrUserGroupToMFTFolder(ctx, folderName, userOrGroup)
Add external user or user groups to virtual folder external users list.

Add external user user groups to virtual folder external users list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderName** | **string**| Name of folder | 
  **userOrGroup** | **string**| The user name or group name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddGateway**
> SuccessData AddGateway(ctx, body)
add gateway.

add gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatewayData**](GatewayData.md)| gateway data | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddHostToHostgroup**
> AgentsInGroupSuccessData AddHostToHostgroup(ctx, body, server, hostgroup)
add agent to hostgroup

Add an agent to hostgroup. Create the the hostgroup if it does not exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AgentInHostgroup**](AgentInHostgroup.md)| The hostname of the new agent | 
  **server** | **string**| The Server the hostgroup belongs to. | 
  **hostgroup** | **string**| The hostgroup name | 

### Return type

[**AgentsInGroupSuccessData**](AgentsInGroupSuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddHubToCluster**
> SuccessData AddHubToCluster(ctx, agentname)
add hub to cluster.

add hub to cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentname** | **string**| Agent name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMFTFolder**
> SuccessData AddMFTFolder(ctx, body)
Add virtual folder

Add virtual folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FolderPropertiesData**](FolderPropertiesData.md)| virtual folder data | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMftUserGroup**
> SuccessData AddMftUserGroup(ctx, body)
Add user group.

Add user group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserGroupPropertiesData**](UserGroupPropertiesData.md)| User group object properites | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPgpTemplate**
> SuccessData AddPgpTemplate(ctx, body, server, agent, templateName)
Add PGP Template

Add PGP Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgpTemplateData**](PgpTemplateData.md)| PGP Template Data | 
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 
  **templateName** | **string**| The PGP Template Name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRemoteHost**
> SuccessData AddRemoteHost(ctx, server, optional)
add remote host to Server

Add a remote host to Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the remote host is going to be added to. | 
 **optional** | ***ConfigApiAddRemoteHostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiAddRemoteHostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AddRemoteHostParams**](AddRemoteHostParams.md)| The non default, advanced configuration data | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRole**
> SuccessData AddRole(ctx, roleFile)
Add Authorization Role

Add Authorization Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleFile** | ***os.File*****os.File**|  | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRoleToLdapGroup**
> SuccessData AddRoleToLdapGroup(ctx, ldapgroup, role)
Add a role to LDAP group

Add a role to LDAP group so any user belong to the LDAP group will get all permissions defined in the role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapgroup** | **string**| Name of LDAP group | 
  **role** | **string**| Name of role | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRoleToUser**
> SuccessData AddRoleToUser(ctx, user, role)
Add a role to user

Add a role to user so that user will inherit role authorization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **user** | **string**| Name of user | 
  **role** | **string**| Name of role | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSecret**
> SuccessData AddSecret(ctx, body)
Add a new secret

Add a new secret to the secrets vault.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretKeyValue**](SecretKeyValue.md)| The new secret value | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServer**
> SuccessData AddServer(ctx, body)
add server to the system

Add a Server. This command setting up new server in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddServerParams**](AddServerParams.md)|  | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUser**
> SuccessData AddUser(ctx, userFile)
Add user

Add user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userFile** | ***os.File*****os.File**|  | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddZosTemplate**
> SuccessData AddZosTemplate(ctx, body, server, agent, templateName)
Add z/OS Template

Add z/OS Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ZosTemplateData**](ZosTemplateData.md)| z/OS Template Data | 
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 
  **templateName** | **string**| The z/OS Template Name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizeMftSshCluster**
> SuccessData AuthorizeMftSshCluster(ctx, body, server, agent, clusterName)
Authorize SSH Cluster

Authorize SSH Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClusterAuthorizationData**](ClusterAuthorizationData.md)| File with content of hostnames and ports | 
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 
  **clusterName** | **string**| Ssh Cluster Name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizeMftSshHost**
> SuccessData AuthorizeMftSshHost(ctx, server, agent, hostname, optional)
Authorize SSH Host

Authorize SSH Host for SFTP account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 
  **hostname** | **string**| Ssh Hostname | 
 **optional** | ***ConfigApiAuthorizeMftSshHostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiAuthorizeMftSshHostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **port** | **optional.String**| Ssh port for the relevant host | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizeSshKnownRemotehost**
> SuccessData AuthorizeSshKnownRemotehost(ctx, server, remotehost)
Authorize

Authorized known ssh remote host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the remote host is connected to. | 
  **remotehost** | **string**| The name of the remote host. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeUserPassword**
> SuccessData ChangeUserPassword(ctx, user, optional)
Change user password

Change user password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **user** | **string**| user name | 
 **optional** | ***ConfigApiChangeUserPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiChangeUserPasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserPassword**](UserPassword.md)| The new password. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAgentCertificateSigningRequest**
> string CreateAgentCertificateSigningRequest(ctx, body, server, agent)
Create certificate signing request (CSR).

Create certificate signing request (CSR) on SSL configured Agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CertificateSigningRequestData**](CertificateSigningRequestData.md)| Certificate Signing Request (CSR) data | 
  **server** | **string**| The Server. | 
  **agent** | **string**| The Agent. | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRunAsUser**
> SuccessData CreateRunAsUser(ctx, body, server)
Add a new Run-as user

Add a new Run-as user to server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RunAsUserData**](RunAsUserData.md)| Run as user data | 
  **server** | **string**| The Server. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAgent**
> SuccessData DeleteAgent(ctx, server, agent)
delete an agent from Server

Delete an agent from a Server. This will not shut the agent down. It only disconnects and removes it from the list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the agent is connected to. | 
  **agent** | **string**| The name of the agent to delete. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArchiveRule**
> SuccessData DeleteArchiveRule(ctx, ruleName, deleteRuleDataFlag)
Delete Workload Archiving rule

Deletes Workload Archiving rule by name. It is required to send deleteRuleData flag to specify if rule need to be deleted with all the collected data or deleteRuleWithoutData otherwise.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleName** | **string**| Rule name to delete | 
  **deleteRuleDataFlag** | **string**| Remove rule with collected data or without. REQUIRED. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthorizationRole**
> SuccessData DeleteAuthorizationRole(ctx, role)
Delete Authorization Role

Delete Authorization Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **role** | **string**| The Role name. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExternalUser**
> SuccessData DeleteExternalUser(ctx, username)
Delete an external user

Delete an existing external user in MFT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The name of the external user to delete | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExternalUserOrUserGroupFromMFTFolder**
> SuccessData DeleteExternalUserOrUserGroupFromMFTFolder(ctx, folderName, userOrGroup)
Remove an external user or user group from an existing virtual folder in MFT.

Remove an external user or user group from an existing virtual folder in MFT.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderName** | **string**| Name of folder | 
  **userOrGroup** | **string**| The user name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHostFromGroup**
> AgentsInGroupSuccessData DeleteHostFromGroup(ctx, server, hostgroup, host)
delete an agent from a hostgroup

Delete an agent from the specified hostgroup. If the group is empty it will also be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the hostgroup belongs to. | 
  **hostgroup** | **string**| The hostgroup name | 
  **host** | **string**| The agent to be deleted | 

### Return type

[**AgentsInGroupSuccessData**](AgentsInGroupSuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHostGroup**
> SuccessData DeleteHostGroup(ctx, server, hostgroup)
delete host group

delete host group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the agent is connected to. | 
  **hostgroup** | **string**| The hostgroup name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMFTFolder**
> SuccessData DeleteMFTFolder(ctx, folderName)
Delete a virtual folder.

Delete an existing virtual folder in MFT.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderName** | **string**| Name of folder | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMftUserGroup**
> SuccessData DeleteMftUserGroup(ctx, name)
Delete user group.

Delete user group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| User group name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePgpTemplate**
> SuccessData DeletePgpTemplate(ctx, server, agent, templateName)
Delete PGP Template

Delete PGP Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 
  **templateName** | **string**| The PGP Template Name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRemoteHost**
> SuccessData DeleteRemoteHost(ctx, server, remotehost)
delete a remote host from Server

Delete a remote host from a Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the remote host is connected to. | 
  **remotehost** | **string**| The name of the remote host to delete. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleFromLdapGroup**
> SuccessData DeleteRoleFromLdapGroup(ctx, ldapgroup, role)
Delete a role from LDAP group

Delete a role from LDAP group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapgroup** | **string**| Name of LDAP group | 
  **role** | **string**| Name of role | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRunAsUser**
> SuccessData DeleteRunAsUser(ctx, server, agent, user)
delete Run-as user

Delete Run-as user from server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server. | 
  **agent** | **string**| The Agent | 
  **user** | **string**| The user name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecret**
> SuccessData DeleteSecret(ctx, name)
Delete an existing secret

Delete an existing secret from the secrets vault.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the secret to update | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> SuccessData DeleteUser(ctx, user)
Delete user

Delete user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **user** | **string**| The user name. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteZosTemplate**
> SuccessData DeleteZosTemplate(ctx, server, agent, templateName)
Delete z/OS Template

Delete z/OS Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 
  **templateName** | **string**| The z/OS Template Name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployAgentCertificate**
> SuccessData DeployAgentCertificate(ctx, crtFile, caChainFile, server, agent)
Deploy certificate (CRT).

Deploy certificate (CRT) on SSL configured Agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crtFile** | ***os.File*****os.File**|  | 
  **caChainFile** | ***os.File*****os.File**|  | 
  **server** | **string**| The Server. | 
  **agent** | **string**| The Agent. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableAgent**
> SuccessData DisableAgent(ctx, server, agent)
disable agent from the Server

Disable an Agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the agent is connected too. | 
  **agent** | **string**| The Agent to be disabled. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableAgent**
> SuccessData EnableAgent(ctx, server, agent)
enable agent from the Server

Enable an Agent. This command does not install or configure the agent. It only enable existing agent in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the agent is connected too. | 
  **agent** | **string**| The Agent to be enabled. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Failover**
> SuccessData Failover(ctx, server)
Perform Manual Failover on a specified Server

Perform Manual Failover on a specified Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**|  | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateMftRsaSshKey**
> SuccessData GenerateMftRsaSshKey(ctx, body, server, agent)
Generate RSA SSH Key

Generate RSA SSH Key pair for SFTP account authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SshKeyProperties**](SshKeyProperties.md)| Ssh Key pair properites | 
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgentCertificateExpirationDate**
> AgentCertificateExpirationData GetAgentCertificateExpirationDate(ctx, server, agent)
Get certificate expiration date.

Get the certificate expiration date of SSL configured Agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server. | 
  **agent** | **string**| The Agent. | 

### Return type

[**AgentCertificateExpirationData**](AgentCertificateExpirationData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgentParameters**
> []KeyValue GetAgentParameters(ctx, server, agent, optional)
get agent parameters

Get all the parameters of the specified Agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the agent is connected to. | 
  **agent** | **string**| The name of the agent to query. | 
 **optional** | ***ConfigApiGetAgentParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetAgentParametersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extendedData** | **optional.Bool**| True to return more agent parameters. HIDDEN | 

### Return type

[**[]KeyValue**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgents**
> AgentDetailsList GetAgents(ctx, server, optional)
get Server agents

Get all the agents of the specified Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server to query. Optionally you can filter agent name of host or alias of the Agent | 
 **optional** | ***ConfigApiGetAgentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetAgentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agent** | **optional.String**| Optionally case insensitive agent name filter of host or alias of the Agent. &#x60;ctm server:agents::get Server AgentName&#x60; returns all agents which names start with &#x60;agentname&#x60; | 

### Return type

[**AgentDetailsList**](AgentDetailsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllArchiveRules**
> ArchiveRulesList GetAllArchiveRules(ctx, )
Get all Workload Archiving rules

Get all the Archiving rules

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ArchiveRulesList**](ArchiveRulesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAuthorizationRoles**
> []RoleHeader GetAllAuthorizationRoles(ctx, optional)
Get Authorization Roles

Get Authorization Roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiGetAllAuthorizationRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetAllAuthorizationRolesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **optional.String**| The Role name. | 
 **description** | **optional.String**| The Role description. | 

### Return type

[**[]RoleHeader**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllLdapGroups**
> []string GetAllLdapGroups(ctx, optional)
Get All LDAP groups

Get All LDAP groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiGetAllLdapGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetAllLdapGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldap** | **optional.String**| The ldap name. | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRolesAssociatedWithLdap**
> []string GetAllRolesAssociatedWithLdap(ctx, ldapgroup, optional)
Get Authorization Roles associated with an LDAP group

Get Authorization Roles associated with an LDAP group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapgroup** | **string**| Name of Ldap group | 
 **optional** | ***ConfigApiGetAllRolesAssociatedWithLdapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetAllRolesAssociatedWithLdapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **optional.String**| The Role name. | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllUsers**
> []UserHeader GetAllUsers(ctx, optional)
Get users

Get users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiGetAllUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetAllUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The user name. | 
 **fullName** | **optional.String**| The user full name. | 
 **description** | **optional.String**| The user description. | 

### Return type

[**[]UserHeader**](UserHeader.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArchiveStatistics**
> RulesStatisticListSummary GetArchiveStatistics(ctx, )
Get Workload Archiving statistics

Get list of statistical information for each Archiving rule and total information about the number of jobs that have been archived, data size of all job logs and outputs that have been archived, size of the Workload Archiving database including all tables and indexes and percentage of disk space used on the Workload Archiving server

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RulesStatisticListSummary**](RulesStatisticListSummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalUserAuthorizedFolders**
> []string GetExternalUserAuthorizedFolders(ctx, name)
Get MFT external user authorized folders

Get MFT external user authorized folders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The external user name. | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalUsers**
> []ExternalUserData GetExternalUsers(ctx, optional)
Get MFT external users that match the search criteria.

Get MFT external users that match the search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiGetExternalUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetExternalUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The user name. | 
 **email** | **optional.String**| The user email. | 
 **description** | **optional.String**| The user description. | 
 **company** | **optional.String**| The user company. | 
 **phoneNumber** | **optional.String**| The user phoneNumber. | 

### Return type

[**[]ExternalUserData**](ExternalUserData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFtsSettings**
> FtsSettingsData GetFtsSettings(ctx, server, agent)
Get File Transfer Server (FTS) configuration data.

Get File Transfer Server (FTS) configuration data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 

### Return type

[**FtsSettingsData**](FtsSettingsData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostgroups**
> []string GetHostgroups(ctx, server)
get Server hostgroups

Get all the hostgroups of the specified Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the hostgroups belong to. | 

### Return type

[**[]string**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostsInGroup**
> []AgentInGroupParams GetHostsInGroup(ctx, server, hostgroup)
get hostgroup agents

Get the agents that compose the specified hostgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the hostgroup belongs to. | 
  **hostgroup** | **string**| The hostgroup name | 

### Return type

[**[]AgentInGroupParams**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHubStatusDetails**
> string GetHubStatusDetails(ctx, nodeId)
Get hub status.

Get hub status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| Node ID of the hub | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMFTFolders**
> []FolderPropertiesData GetMFTFolders(ctx, optional)
Get MFT virtual folders that match the search criteria.

Get MFT virtual folders that match the search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiGetMFTFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetMFTFoldersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The virtual folder name. | 

### Return type

[**[]FolderPropertiesData**](FolderPropertiesData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMftConfiguration**
> MftConfigurationData GetMftConfiguration(ctx, server, agent)
Get MFT Configuration

Get MFT Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 

### Return type

[**MftConfigurationData**](MftConfigurationData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMftGateways**
> []GatewayData GetMftGateways(ctx, )
Get MFT gateways

Get MFT gateways

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]GatewayData**](GatewayData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMftUserGroups**
> []UserGroupPropertiesData GetMftUserGroups(ctx, optional)
Get all user groups that match the search criteria.

Get all user groups that match the search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiGetMftUserGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetMftUserGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The user group name. | 
 **externalUsers** | **optional.String**| external users. | 
 **ldapGroups** | **optional.String**| ldap groups. | 

### Return type

[**[]UserGroupPropertiesData**](UserGroupPropertiesData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPgpTemplates**
> []PgpTemplateData GetPgpTemplates(ctx, server, agent, optional)
Get PGP Templates

Get PGP Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 
 **optional** | ***ConfigApiGetPgpTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetPgpTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| The PGP Template Name | 

### Return type

[**[]PgpTemplateData**](PgpTemplateData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRemoteHostProperties**
> AddRemoteHostParams GetRemoteHostProperties(ctx, server, remotehost)
get a remote host configuration from Server

Get the remote host configuration properties from the Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server the remote host  is connected to. | 
  **remotehost** | **string**| The name of the remote host. | 

### Return type

[**AddRemoteHostParams**](AddRemoteHostParams.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRemoteHosts**
> []string GetRemoteHosts(ctx, server)
get Server remote hosts

Get all the remote hosts of the specified Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server to query. | 

### Return type

[**[]string**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> RoleData GetRole(ctx, role)
Get Authorization Role

Get Authorization Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **role** | **string**| The Role name. | 

### Return type

[**RoleData**](RoleData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleAssociates**
> []AssociateData GetRoleAssociates(ctx, role)
Get all authorization entities associated with role

Get all authorization entities associated with role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **role** | **string**| role name. | 

### Return type

[**[]AssociateData**](AssociateData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRunAsUser**
> RunAsUserData GetRunAsUser(ctx, server, agent, user)
Get Run-as user

Get Run-as user details from server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server. | 
  **agent** | **string**| The Agent | 
  **user** | **string**| The user name | 

### Return type

[**RunAsUserData**](RunAsUserData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRunAsUsersList**
> []RunAsUserData GetRunAsUsersList(ctx, server, optional)
Get Run-as user list that match the requested search criteria.

Get Run-as user list that match the requested search criteria from server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server. | 
 **optional** | ***ConfigApiGetRunAsUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetRunAsUsersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | **optional.String**| The Run-as user. | 
 **agent** | **optional.String**| The agent. | 

### Return type

[**[]RunAsUserData**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerParameters**
> []KeyValue GetServerParameters(ctx, server)
get Server parameters

Get all the parameters of the specified Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server to query. | 

### Return type

[**[]KeyValue**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServers**
> []CtmDetails GetServers(ctx, )
get all the Servers name and hostname in the system

Get the names and hostnames of all Servers in the system.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CtmDetails**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> UserData GetUser(ctx, user)
Get user

Get user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **user** | **string**| The user name. | 

### Return type

[**UserData**](UserData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserEffectiveRights**
> RoleData GetUserEffectiveRights(ctx, )
Get user real effective authorizations

Get user real effective authorizations by all his roles

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RoleData**](RoleData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetZosTemplates**
> []ZosTemplateData GetZosTemplates(ctx, server, agent, optional)
Get z/OS Templates

Get z/OS Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 
 **optional** | ***ConfigApiGetZosTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiGetZosTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| The z/OS Template Name | 

### Return type

[**[]ZosTemplateData**](ZosTemplateData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecrets**
> []string ListSecrets(ctx, )
Get list of secret names

Get the list of names of all the secrets in the vault

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]string**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingAgent**
> SuccessData PingAgent(ctx, server, agent, optional)
ping to the agent in the Server

Ping an Agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server. | 
  **agent** | **string**| The Agent. | 
 **optional** | ***ConfigApiPingAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiPingAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PingAgentParams**](PingAgentParams.md)|  | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecycleItem**
> SuccessData RecycleItem(ctx, id)
recycle item

Recycle an item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| item data | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveControlmServer**
> SuccessData RemoveControlmServer(ctx, server)
Delete Server

Delete Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| Server host name. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveGateway**
> SuccessData RemoveGateway(ctx, gatewayName)
remove gateway.

remove gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayName** | **string**| gateway name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveHubFromCluster**
> SuccessData RemoveHubFromCluster(ctx, agentname)
remove hub from cluster.

remove hub from cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentname** | **string**| Agent name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoleFromUser**
> SuccessData RemoveRoleFromUser(ctx, user, role)
Remove a role from a user

Remove a role from a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **user** | **string**| Name of user | 
  **role** | **string**| Name of role | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendArchiveCleanupRequest**
> SuccessData SendArchiveCleanupRequest(ctx, optional)
Deletes data (jobs including outputs and logs) from the Workload Archiving database.

Deletes data (jobs including outputs and logs) by search criteria from the Workload Archiving database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiSendArchiveCleanupRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiSendArchiveCleanupRequestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **optional.String**| Job&#x27;s application. | 
 **applicationExceptions** | **optional.String**| Cleanup should skip job&#x27;s application that are mentioned in exceptions | 
 **subApplication** | **optional.String**| Job&#x27;s sub application | 
 **subApplicationExceptions** | **optional.String**| Job&#x27;s sub application exception | 
 **ctm** | **optional.String**| server name | 
 **server** | **optional.String**| Server name | 
 **ctmExceptions** | **optional.String**| server exceptions | 
 **serverExceptions** | **optional.String**| Server exceptions | 
 **folder** | **optional.String**| Job&#x27;s folder. | 
 **folderExceptions** | **optional.String**| Job&#x27;s folder exceptions | 
 **jobname** | **optional.String**| Job&#x27;s name | 
 **jobnameExceptions** | **optional.String**| Job&#x27;s name exceptions | 
 **library** | **optional.String**| Job&#x27;s library | 
 **libraryExceptions** | **optional.String**| Job&#x27;s library exceptions | 
 **ruleName** | **optional.String**| Job&#x27;s archive rule | 
 **jobStatus** | **optional.String**| The job&#x27;s end status. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAgentParameter**
> KeyValue SetAgentParameter(ctx, body, server, agent, name)
set agent parameter

Set the value of the specified parameter in the specified agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Value**](Value.md)| The new parameter value. | 
  **server** | **string**| The Server the agent is connected to. | 
  **agent** | **string**| The name of the agent to update. | 
  **name** | **string**| The parameter name. | 

### Return type

[**KeyValue**](KeyValue.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSystemParam**
> SuccessData SetSystemParam(ctx, body, name)
set value of a an em system parameter

Set value of an enterprise management system parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Value**](Value.md)| Param new value | 
  **name** | **string**| Parameter name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Setasprimary**
> SuccessData Setasprimary(ctx, server)
Set secondary server as Primary on a specified Server

Set secondary server as Primary on a specified Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**|  | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestRunAsUser**
> SuccessData TestRunAsUser(ctx, server, agent, user, optional)
Test existed Run-as user

Test existing Run-as user in server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**| The Server. | 
  **agent** | **string**| The Agent | 
  **user** | **string**| The user name | 
 **optional** | ***ConfigApiTestRunAsUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiTestRunAsUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of RunAsUserDetailsData**](RunAsUserDetailsData.md)| Run as user details data | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateArchiveRule**
> SuccessData UpdateArchiveRule(ctx, body, ruleName)
Edit Workload Archiving rule

Edit Workload Archiving rule details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ArchiveRule**](ArchiveRule.md)| Edit Workload Archiving rule details | 
  **ruleName** | **string**| Rule name to update | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExternalUser**
> SuccessData UpdateExternalUser(ctx, body, username)
Update an external user

Update an external user for b2b

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExternalUserData**](ExternalUserData.md)| External user data | 
  **username** | **string**| The external user name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFtsSettings**
> SuccessData UpdateFtsSettings(ctx, body, server, agent)
Update File Transfer Server (FTS) configuration data.

Update File Transfer Server (FTS) configuration data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FtsSettingsData**](FtsSettingsData.md)| File Transfer Server (FTS) configuration data | 
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHostsInHostgroup**
> SuccessData UpdateHostsInHostgroup(ctx, body, server, hostgroup)
update agents in hostgroup.

update agents in hostgroup.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HostgroupProperties**](HostgroupProperties.md)| Agent list to update in a hostgroup | 
  **server** | **string**| The Server the agent is connected to. | 
  **hostgroup** | **string**| The hostgroup name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMFTFolder**
> SuccessData UpdateMFTFolder(ctx, body, folderName)
Update an existing virtual folder in MFT.

Update an existing virtual folder in MFT.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FolderPropertiesData**](FolderPropertiesData.md)| virtual folder data | 
  **folderName** | **string**| Name of folder | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMftConfiguration**
> SuccessData UpdateMftConfiguration(ctx, body, server, agent)
Update MFT Configuration

Update MFT Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MftConfigurationData**](MftConfigurationData.md)| MFT Configuration Data | 
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMftUserGroup**
> SuccessData UpdateMftUserGroup(ctx, body, name)
Update user group.

Update user group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserGroupDetailsData**](UserGroupDetailsData.md)| User group details | 
  **name** | **string**| User group name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePgpTemplate**
> SuccessData UpdatePgpTemplate(ctx, body, server, agent, templateName)
Update PGP Template

Update PGP Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PgpTemplateData**](PgpTemplateData.md)| PGP Template Data | 
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 
  **templateName** | **string**| The PGP Template Name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRole**
> SuccessData UpdateRole(ctx, roleFile, role)
Update Authorization Role

Update Authorization Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleFile** | ***os.File*****os.File**|  | 
  **role** | **string**| The Role name. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRunAsUser**
> SuccessData UpdateRunAsUser(ctx, body, server, agent, user)
Update Run-as user

Update Run-as user details in server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RunAsUserDetailsData**](RunAsUserDetailsData.md)| Run as user details data | 
  **server** | **string**| The Server. | 
  **agent** | **string**| The Agent | 
  **user** | **string**| The user name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecret**
> SuccessData UpdateSecret(ctx, name, optional)
Update an existing secret

Update an existing secret in the secrets vault.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the secret to update | 
 **optional** | ***ConfigApiUpdateSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiUpdateSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SecretValue**](SecretValue.md)| The new value for the secret | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> SuccessData UpdateUser(ctx, userFile, user)
Update user

Update user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userFile** | ***os.File*****os.File**|  | 
  **user** | **string**| The user name. | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateZosTemplate**
> SuccessData UpdateZosTemplate(ctx, body, server, agent, templateName)
Update z/OS Template

Update z/OS Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ZosTemplateData**](ZosTemplateData.md)| z/OS Template Data | 
  **server** | **string**| The Server | 
  **agent** | **string**| The Agent | 
  **templateName** | **string**| The z/OS Template Name | 

### Return type

[**SuccessData**](SuccessData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

