# OrderFolderParameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctm** | **string** | The Control-M Server to order from. REQUIRED. | [optional] [default to null]
**Folder** | **string** | The folder to order. REQUIRED. | [optional] [default to null]
**Jobs** | **string** | Filter the jobs to order. | [optional] [default to null]
**Library** | **string** | The z/os library that contains the job (only for MF). | [optional] [default to null]
**CreateDuplicate** | **bool** | Is it allowed to order the same jobs more than once to the same SMART folder. HIDDEN. | [optional] [default to null]
**Hold** | **bool** | Are jobs ordered in a HOLD state. HIDDEN. | [optional] [default to null]
**IgnoreCriteria** | **bool** | Is scheduling criteria to be ignored. HIDDEN. | [optional] [default to null]
**IndependentFlow** | **bool** | Whether to generate new flow in this order. HIDDEN. | [optional] [default to null]
**OrderDate** | **string** | The order date that is attached to this order command. HIDDEN. | [optional] [default to null]
**OrderIntoFolder** | **string** | Policy for placing the jobs in a SMART folder. HIDDEN. | [optional] [default to null]
**WaitForOrderDate** | **bool** | Whether to wait for the order date when running the jobs. HIDDEN. | [optional] [default to null]
**Variables** | [**[]map[string]string**](map.md) | Job Variables for this run. HIDDEN. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

