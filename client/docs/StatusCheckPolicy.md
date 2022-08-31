# StatusCheckPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Strict** | **bool** |  | 
**Contexts** | **[]string** |  | 
**Checks** | [**[]StatusCheckPolicyChecksInner**](StatusCheckPolicyChecksInner.md) |  | 
**ContextsUrl** | **string** |  | 

## Methods

### NewStatusCheckPolicy

`func NewStatusCheckPolicy(url string, strict bool, contexts []string, checks []StatusCheckPolicyChecksInner, contextsUrl string, ) *StatusCheckPolicy`

NewStatusCheckPolicy instantiates a new StatusCheckPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCheckPolicyWithDefaults

`func NewStatusCheckPolicyWithDefaults() *StatusCheckPolicy`

NewStatusCheckPolicyWithDefaults instantiates a new StatusCheckPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *StatusCheckPolicy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StatusCheckPolicy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StatusCheckPolicy) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStrict

`func (o *StatusCheckPolicy) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *StatusCheckPolicy) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *StatusCheckPolicy) SetStrict(v bool)`

SetStrict sets Strict field to given value.


### GetContexts

`func (o *StatusCheckPolicy) GetContexts() []string`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *StatusCheckPolicy) GetContextsOk() (*[]string, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *StatusCheckPolicy) SetContexts(v []string)`

SetContexts sets Contexts field to given value.


### GetChecks

`func (o *StatusCheckPolicy) GetChecks() []StatusCheckPolicyChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *StatusCheckPolicy) GetChecksOk() (*[]StatusCheckPolicyChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *StatusCheckPolicy) SetChecks(v []StatusCheckPolicyChecksInner)`

SetChecks sets Checks field to given value.


### GetContextsUrl

`func (o *StatusCheckPolicy) GetContextsUrl() string`

GetContextsUrl returns the ContextsUrl field if non-nil, zero value otherwise.

### GetContextsUrlOk

`func (o *StatusCheckPolicy) GetContextsUrlOk() (*string, bool)`

GetContextsUrlOk returns a tuple with the ContextsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextsUrl

`func (o *StatusCheckPolicy) SetContextsUrl(v string)`

SetContextsUrl sets ContextsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


