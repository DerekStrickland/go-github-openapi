# CodeScanningAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | A unique identifier for the rule used to detect the alert. | [optional] 
**Name** | Pointer to **string** | The name of the rule used to detect the alert. | [optional] 
**Severity** | Pointer to **NullableString** | The severity of the alert. | [optional] 
**SecuritySeverityLevel** | Pointer to **NullableString** | The security severity of the alert. | [optional] 
**Description** | Pointer to **string** | A short description of the rule used to detect the alert. | [optional] 
**FullDescription** | Pointer to **string** | description of the rule used to detect the alert. | [optional] 
**Tags** | Pointer to **[]string** | A set of tags applicable for the rule. | [optional] 
**Help** | Pointer to **NullableString** | Detailed documentation for the rule as GitHub Flavored Markdown. | [optional] 

## Methods

### NewCodeScanningAlertRule

`func NewCodeScanningAlertRule() *CodeScanningAlertRule`

NewCodeScanningAlertRule instantiates a new CodeScanningAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningAlertRuleWithDefaults

`func NewCodeScanningAlertRuleWithDefaults() *CodeScanningAlertRule`

NewCodeScanningAlertRuleWithDefaults instantiates a new CodeScanningAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CodeScanningAlertRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodeScanningAlertRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodeScanningAlertRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CodeScanningAlertRule) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CodeScanningAlertRule) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CodeScanningAlertRule) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CodeScanningAlertRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodeScanningAlertRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodeScanningAlertRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CodeScanningAlertRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *CodeScanningAlertRule) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CodeScanningAlertRule) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CodeScanningAlertRule) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CodeScanningAlertRule) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *CodeScanningAlertRule) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *CodeScanningAlertRule) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetSecuritySeverityLevel

`func (o *CodeScanningAlertRule) GetSecuritySeverityLevel() string`

GetSecuritySeverityLevel returns the SecuritySeverityLevel field if non-nil, zero value otherwise.

### GetSecuritySeverityLevelOk

`func (o *CodeScanningAlertRule) GetSecuritySeverityLevelOk() (*string, bool)`

GetSecuritySeverityLevelOk returns a tuple with the SecuritySeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySeverityLevel

`func (o *CodeScanningAlertRule) SetSecuritySeverityLevel(v string)`

SetSecuritySeverityLevel sets SecuritySeverityLevel field to given value.

### HasSecuritySeverityLevel

`func (o *CodeScanningAlertRule) HasSecuritySeverityLevel() bool`

HasSecuritySeverityLevel returns a boolean if a field has been set.

### SetSecuritySeverityLevelNil

`func (o *CodeScanningAlertRule) SetSecuritySeverityLevelNil(b bool)`

 SetSecuritySeverityLevelNil sets the value for SecuritySeverityLevel to be an explicit nil

### UnsetSecuritySeverityLevel
`func (o *CodeScanningAlertRule) UnsetSecuritySeverityLevel()`

UnsetSecuritySeverityLevel ensures that no value is present for SecuritySeverityLevel, not even an explicit nil
### GetDescription

`func (o *CodeScanningAlertRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CodeScanningAlertRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CodeScanningAlertRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CodeScanningAlertRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFullDescription

`func (o *CodeScanningAlertRule) GetFullDescription() string`

GetFullDescription returns the FullDescription field if non-nil, zero value otherwise.

### GetFullDescriptionOk

`func (o *CodeScanningAlertRule) GetFullDescriptionOk() (*string, bool)`

GetFullDescriptionOk returns a tuple with the FullDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDescription

`func (o *CodeScanningAlertRule) SetFullDescription(v string)`

SetFullDescription sets FullDescription field to given value.

### HasFullDescription

`func (o *CodeScanningAlertRule) HasFullDescription() bool`

HasFullDescription returns a boolean if a field has been set.

### GetTags

`func (o *CodeScanningAlertRule) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CodeScanningAlertRule) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CodeScanningAlertRule) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CodeScanningAlertRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CodeScanningAlertRule) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CodeScanningAlertRule) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetHelp

`func (o *CodeScanningAlertRule) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *CodeScanningAlertRule) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *CodeScanningAlertRule) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *CodeScanningAlertRule) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### SetHelpNil

`func (o *CodeScanningAlertRule) SetHelpNil(b bool)`

 SetHelpNil sets the value for Help to be an explicit nil

### UnsetHelp
`func (o *CodeScanningAlertRule) UnsetHelp()`

UnsetHelp ensures that no value is present for Help, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


