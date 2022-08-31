# PullsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The title of the pull request. | [optional] 
**Body** | Pointer to **string** | The contents of the pull request. | [optional] 
**State** | Pointer to **string** | State of this Pull Request. Either &#x60;open&#x60; or &#x60;closed&#x60;. | [optional] 
**Base** | Pointer to **string** | The name of the branch you want your changes pulled into. This should be an existing branch on the current repository. You cannot update the base branch on a pull request to point to another repository. | [optional] 
**MaintainerCanModify** | Pointer to **bool** | Indicates whether [maintainers can modify](https://docs.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request. | [optional] 

## Methods

### NewPullsUpdateRequest

`func NewPullsUpdateRequest() *PullsUpdateRequest`

NewPullsUpdateRequest instantiates a new PullsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsUpdateRequestWithDefaults

`func NewPullsUpdateRequestWithDefaults() *PullsUpdateRequest`

NewPullsUpdateRequestWithDefaults instantiates a new PullsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PullsUpdateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullsUpdateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullsUpdateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PullsUpdateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBody

`func (o *PullsUpdateRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullsUpdateRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullsUpdateRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PullsUpdateRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetState

`func (o *PullsUpdateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PullsUpdateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PullsUpdateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PullsUpdateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetBase

`func (o *PullsUpdateRequest) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PullsUpdateRequest) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PullsUpdateRequest) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *PullsUpdateRequest) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetMaintainerCanModify

`func (o *PullsUpdateRequest) GetMaintainerCanModify() bool`

GetMaintainerCanModify returns the MaintainerCanModify field if non-nil, zero value otherwise.

### GetMaintainerCanModifyOk

`func (o *PullsUpdateRequest) GetMaintainerCanModifyOk() (*bool, bool)`

GetMaintainerCanModifyOk returns a tuple with the MaintainerCanModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerCanModify

`func (o *PullsUpdateRequest) SetMaintainerCanModify(v bool)`

SetMaintainerCanModify sets MaintainerCanModify field to given value.

### HasMaintainerCanModify

`func (o *PullsUpdateRequest) HasMaintainerCanModify() bool`

HasMaintainerCanModify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


