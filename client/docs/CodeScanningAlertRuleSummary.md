# CodeScanningAlertRuleSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | A unique identifier for the rule used to detect the alert. | [optional] 
**Name** | Pointer to **string** | The name of the rule used to detect the alert. | [optional] 
**Tags** | Pointer to **[]string** | A set of tags applicable for the rule. | [optional] 
**Severity** | Pointer to **NullableString** | The severity of the alert. | [optional] 
**Description** | Pointer to **string** | A short description of the rule used to detect the alert. | [optional] 

## Methods

### NewCodeScanningAlertRuleSummary

`func NewCodeScanningAlertRuleSummary() *CodeScanningAlertRuleSummary`

NewCodeScanningAlertRuleSummary instantiates a new CodeScanningAlertRuleSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningAlertRuleSummaryWithDefaults

`func NewCodeScanningAlertRuleSummaryWithDefaults() *CodeScanningAlertRuleSummary`

NewCodeScanningAlertRuleSummaryWithDefaults instantiates a new CodeScanningAlertRuleSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CodeScanningAlertRuleSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodeScanningAlertRuleSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodeScanningAlertRuleSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CodeScanningAlertRuleSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CodeScanningAlertRuleSummary) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CodeScanningAlertRuleSummary) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CodeScanningAlertRuleSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodeScanningAlertRuleSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodeScanningAlertRuleSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CodeScanningAlertRuleSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *CodeScanningAlertRuleSummary) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CodeScanningAlertRuleSummary) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CodeScanningAlertRuleSummary) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CodeScanningAlertRuleSummary) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CodeScanningAlertRuleSummary) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CodeScanningAlertRuleSummary) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetSeverity

`func (o *CodeScanningAlertRuleSummary) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CodeScanningAlertRuleSummary) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CodeScanningAlertRuleSummary) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CodeScanningAlertRuleSummary) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *CodeScanningAlertRuleSummary) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *CodeScanningAlertRuleSummary) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetDescription

`func (o *CodeScanningAlertRuleSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CodeScanningAlertRuleSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CodeScanningAlertRuleSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CodeScanningAlertRuleSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


