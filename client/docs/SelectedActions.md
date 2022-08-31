# SelectedActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GithubOwnedAllowed** | Pointer to **bool** | Whether GitHub-owned actions are allowed. For example, this includes the actions in the &#x60;actions&#x60; organization. | [optional] 
**VerifiedAllowed** | Pointer to **bool** | Whether actions from GitHub Marketplace verified creators are allowed. Set to &#x60;true&#x60; to allow all actions by GitHub Marketplace verified creators. | [optional] 
**PatternsAllowed** | Pointer to **[]string** | Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, &#x60;monalisa/octocat@*&#x60;, &#x60;monalisa/octocat@v2&#x60;, &#x60;monalisa/_*&#x60;.\&quot; | [optional] 

## Methods

### NewSelectedActions

`func NewSelectedActions() *SelectedActions`

NewSelectedActions instantiates a new SelectedActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedActionsWithDefaults

`func NewSelectedActionsWithDefaults() *SelectedActions`

NewSelectedActionsWithDefaults instantiates a new SelectedActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGithubOwnedAllowed

`func (o *SelectedActions) GetGithubOwnedAllowed() bool`

GetGithubOwnedAllowed returns the GithubOwnedAllowed field if non-nil, zero value otherwise.

### GetGithubOwnedAllowedOk

`func (o *SelectedActions) GetGithubOwnedAllowedOk() (*bool, bool)`

GetGithubOwnedAllowedOk returns a tuple with the GithubOwnedAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubOwnedAllowed

`func (o *SelectedActions) SetGithubOwnedAllowed(v bool)`

SetGithubOwnedAllowed sets GithubOwnedAllowed field to given value.

### HasGithubOwnedAllowed

`func (o *SelectedActions) HasGithubOwnedAllowed() bool`

HasGithubOwnedAllowed returns a boolean if a field has been set.

### GetVerifiedAllowed

`func (o *SelectedActions) GetVerifiedAllowed() bool`

GetVerifiedAllowed returns the VerifiedAllowed field if non-nil, zero value otherwise.

### GetVerifiedAllowedOk

`func (o *SelectedActions) GetVerifiedAllowedOk() (*bool, bool)`

GetVerifiedAllowedOk returns a tuple with the VerifiedAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAllowed

`func (o *SelectedActions) SetVerifiedAllowed(v bool)`

SetVerifiedAllowed sets VerifiedAllowed field to given value.

### HasVerifiedAllowed

`func (o *SelectedActions) HasVerifiedAllowed() bool`

HasVerifiedAllowed returns a boolean if a field has been set.

### GetPatternsAllowed

`func (o *SelectedActions) GetPatternsAllowed() []string`

GetPatternsAllowed returns the PatternsAllowed field if non-nil, zero value otherwise.

### GetPatternsAllowedOk

`func (o *SelectedActions) GetPatternsAllowedOk() (*[]string, bool)`

GetPatternsAllowedOk returns a tuple with the PatternsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternsAllowed

`func (o *SelectedActions) SetPatternsAllowed(v []string)`

SetPatternsAllowed sets PatternsAllowed field to given value.

### HasPatternsAllowed

`func (o *SelectedActions) HasPatternsAllowed() bool`

HasPatternsAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


