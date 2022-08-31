# ReposReplaceAllTopicsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | **[]string** | An array of topics to add to the repository. Pass one or more topics to _replace_ the set of existing topics. Send an empty array (&#x60;[]&#x60;) to clear all topics from the repository. **Note:** Topic &#x60;names&#x60; cannot contain uppercase letters. | 

## Methods

### NewReposReplaceAllTopicsRequest

`func NewReposReplaceAllTopicsRequest(names []string, ) *ReposReplaceAllTopicsRequest`

NewReposReplaceAllTopicsRequest instantiates a new ReposReplaceAllTopicsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposReplaceAllTopicsRequestWithDefaults

`func NewReposReplaceAllTopicsRequestWithDefaults() *ReposReplaceAllTopicsRequest`

NewReposReplaceAllTopicsRequestWithDefaults instantiates a new ReposReplaceAllTopicsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *ReposReplaceAllTopicsRequest) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *ReposReplaceAllTopicsRequest) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *ReposReplaceAllTopicsRequest) SetNames(v []string)`

SetNames sets Names field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


