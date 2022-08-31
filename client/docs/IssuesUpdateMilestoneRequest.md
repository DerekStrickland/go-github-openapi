# IssuesUpdateMilestoneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The title of the milestone. | [optional] 
**State** | Pointer to **string** | The state of the milestone. Either &#x60;open&#x60; or &#x60;closed&#x60;. | [optional] [default to "open"]
**Description** | Pointer to **string** | A description of the milestone. | [optional] 
**DueOn** | Pointer to **time.Time** | The milestone due date. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 

## Methods

### NewIssuesUpdateMilestoneRequest

`func NewIssuesUpdateMilestoneRequest() *IssuesUpdateMilestoneRequest`

NewIssuesUpdateMilestoneRequest instantiates a new IssuesUpdateMilestoneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesUpdateMilestoneRequestWithDefaults

`func NewIssuesUpdateMilestoneRequestWithDefaults() *IssuesUpdateMilestoneRequest`

NewIssuesUpdateMilestoneRequestWithDefaults instantiates a new IssuesUpdateMilestoneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *IssuesUpdateMilestoneRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssuesUpdateMilestoneRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssuesUpdateMilestoneRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IssuesUpdateMilestoneRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetState

`func (o *IssuesUpdateMilestoneRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IssuesUpdateMilestoneRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IssuesUpdateMilestoneRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IssuesUpdateMilestoneRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDescription

`func (o *IssuesUpdateMilestoneRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssuesUpdateMilestoneRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssuesUpdateMilestoneRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssuesUpdateMilestoneRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueOn

`func (o *IssuesUpdateMilestoneRequest) GetDueOn() time.Time`

GetDueOn returns the DueOn field if non-nil, zero value otherwise.

### GetDueOnOk

`func (o *IssuesUpdateMilestoneRequest) GetDueOnOk() (*time.Time, bool)`

GetDueOnOk returns a tuple with the DueOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueOn

`func (o *IssuesUpdateMilestoneRequest) SetDueOn(v time.Time)`

SetDueOn sets DueOn field to given value.

### HasDueOn

`func (o *IssuesUpdateMilestoneRequest) HasDueOn() bool`

HasDueOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


