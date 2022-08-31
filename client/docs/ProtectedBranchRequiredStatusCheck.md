# ProtectedBranchRequiredStatusCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**EnforcementLevel** | Pointer to **string** |  | [optional] 
**Contexts** | **[]string** |  | 
**Checks** | [**[]ProtectedBranchRequiredStatusCheckChecksInner**](ProtectedBranchRequiredStatusCheckChecksInner.md) |  | 
**ContextsUrl** | Pointer to **string** |  | [optional] 
**Strict** | Pointer to **bool** |  | [optional] 

## Methods

### NewProtectedBranchRequiredStatusCheck

`func NewProtectedBranchRequiredStatusCheck(contexts []string, checks []ProtectedBranchRequiredStatusCheckChecksInner, ) *ProtectedBranchRequiredStatusCheck`

NewProtectedBranchRequiredStatusCheck instantiates a new ProtectedBranchRequiredStatusCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedBranchRequiredStatusCheckWithDefaults

`func NewProtectedBranchRequiredStatusCheckWithDefaults() *ProtectedBranchRequiredStatusCheck`

NewProtectedBranchRequiredStatusCheckWithDefaults instantiates a new ProtectedBranchRequiredStatusCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ProtectedBranchRequiredStatusCheck) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProtectedBranchRequiredStatusCheck) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProtectedBranchRequiredStatusCheck) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProtectedBranchRequiredStatusCheck) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnforcementLevel

`func (o *ProtectedBranchRequiredStatusCheck) GetEnforcementLevel() string`

GetEnforcementLevel returns the EnforcementLevel field if non-nil, zero value otherwise.

### GetEnforcementLevelOk

`func (o *ProtectedBranchRequiredStatusCheck) GetEnforcementLevelOk() (*string, bool)`

GetEnforcementLevelOk returns a tuple with the EnforcementLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementLevel

`func (o *ProtectedBranchRequiredStatusCheck) SetEnforcementLevel(v string)`

SetEnforcementLevel sets EnforcementLevel field to given value.

### HasEnforcementLevel

`func (o *ProtectedBranchRequiredStatusCheck) HasEnforcementLevel() bool`

HasEnforcementLevel returns a boolean if a field has been set.

### GetContexts

`func (o *ProtectedBranchRequiredStatusCheck) GetContexts() []string`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *ProtectedBranchRequiredStatusCheck) GetContextsOk() (*[]string, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *ProtectedBranchRequiredStatusCheck) SetContexts(v []string)`

SetContexts sets Contexts field to given value.


### GetChecks

`func (o *ProtectedBranchRequiredStatusCheck) GetChecks() []ProtectedBranchRequiredStatusCheckChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ProtectedBranchRequiredStatusCheck) GetChecksOk() (*[]ProtectedBranchRequiredStatusCheckChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ProtectedBranchRequiredStatusCheck) SetChecks(v []ProtectedBranchRequiredStatusCheckChecksInner)`

SetChecks sets Checks field to given value.


### GetContextsUrl

`func (o *ProtectedBranchRequiredStatusCheck) GetContextsUrl() string`

GetContextsUrl returns the ContextsUrl field if non-nil, zero value otherwise.

### GetContextsUrlOk

`func (o *ProtectedBranchRequiredStatusCheck) GetContextsUrlOk() (*string, bool)`

GetContextsUrlOk returns a tuple with the ContextsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextsUrl

`func (o *ProtectedBranchRequiredStatusCheck) SetContextsUrl(v string)`

SetContextsUrl sets ContextsUrl field to given value.

### HasContextsUrl

`func (o *ProtectedBranchRequiredStatusCheck) HasContextsUrl() bool`

HasContextsUrl returns a boolean if a field has been set.

### GetStrict

`func (o *ProtectedBranchRequiredStatusCheck) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *ProtectedBranchRequiredStatusCheck) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *ProtectedBranchRequiredStatusCheck) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *ProtectedBranchRequiredStatusCheck) HasStrict() bool`

HasStrict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


