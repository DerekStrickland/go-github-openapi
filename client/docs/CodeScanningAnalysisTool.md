# CodeScanningAnalysisTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the tool used to generate the code scanning analysis. | [optional] 
**Version** | Pointer to **NullableString** | The version of the tool used to generate the code scanning analysis. | [optional] 
**Guid** | Pointer to **NullableString** | The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data. | [optional] 

## Methods

### NewCodeScanningAnalysisTool

`func NewCodeScanningAnalysisTool() *CodeScanningAnalysisTool`

NewCodeScanningAnalysisTool instantiates a new CodeScanningAnalysisTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningAnalysisToolWithDefaults

`func NewCodeScanningAnalysisToolWithDefaults() *CodeScanningAnalysisTool`

NewCodeScanningAnalysisToolWithDefaults instantiates a new CodeScanningAnalysisTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CodeScanningAnalysisTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodeScanningAnalysisTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodeScanningAnalysisTool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CodeScanningAnalysisTool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *CodeScanningAnalysisTool) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CodeScanningAnalysisTool) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CodeScanningAnalysisTool) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CodeScanningAnalysisTool) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *CodeScanningAnalysisTool) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *CodeScanningAnalysisTool) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetGuid

`func (o *CodeScanningAnalysisTool) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *CodeScanningAnalysisTool) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *CodeScanningAnalysisTool) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *CodeScanningAnalysisTool) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *CodeScanningAnalysisTool) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *CodeScanningAnalysisTool) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


