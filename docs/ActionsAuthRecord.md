# ActionsAuthRecord

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** | component in which the action sent - Possible Values - EM, CTM_Server, CTM_Agent, CTM_CM | [optional] [default to null]
**Action** | **string** | action name unique list of actions&#x27; names as appear in the actions_auth em db table, along with the destination makes a unique action auth entry | [optional] [default to null]
**Category** | **string** | Privilege (Can be empty) Possible Values for categories fields  - OPER, DATABASE, CONFIG, CTMSEC, CPMAN, RAMAN, AGMAN, CMMAN, CCP, UNKNOWN | [optional] [default to null]
**AuthLevel** | **string** | required minimum authorization level - BROWSE, UPDATE, FULL | [optional] [default to null]
**ActionType** | **string** | R - request always pass, authorization is done on the response; C - Connection Profile name will remain empty and will be ignored; E - Not related to authorization | [optional] [default to null]
**AdditionalInfo** | **string** | currently not in use regarding authorization | [optional] [default to null]
**Category2** | **string** |  | [optional] [default to null]
**Policy** | **string** | relation between Category and category2 - 1. Only category, 2. Only Category2, 3. Category AND Category2, 4. Category OR Category2 | [optional] [default to null]
**AuthAttr** | **string** | In case that not all information for authorization exist in the header request, it contain path in the body for the missing field\\fields. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

