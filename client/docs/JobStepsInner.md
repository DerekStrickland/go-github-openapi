# JobStepsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The phase of the lifecycle that the job is currently in. | 
**Conclusion** | **NullableString** | The outcome of the job. | 
**Name** | **string** | The name of the job. | 
**Number** | **int32** |  | 
**StartedAt** | Pointer to **NullableTime** | The time that the step started, in ISO 8601 format. | [optional] 
**CompletedAt** | Pointer to **NullableTime** | The time that the job finished, in ISO 8601 format. | [optional] 

## Methods

### NewJobStepsInner

`func NewJobStepsInner(status string, conclusion NullableString, name string, number int32, ) *JobStepsInner`

NewJobStepsInner instantiates a new JobStepsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStepsInnerWithDefaults

`func NewJobStepsInnerWithDefaults() *JobStepsInner`

NewJobStepsInnerWithDefaults instantiates a new JobStepsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *JobStepsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobStepsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobStepsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetConclusion

`func (o *JobStepsInner) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *JobStepsInner) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *JobStepsInner) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.


### SetConclusionNil

`func (o *JobStepsInner) SetConclusionNil(b bool)`

 SetConclusionNil sets the value for Conclusion to be an explicit nil

### UnsetConclusion
`func (o *JobStepsInner) UnsetConclusion()`

UnsetConclusion ensures that no value is present for Conclusion, not even an explicit nil
### GetName

`func (o *JobStepsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobStepsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobStepsInner) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *JobStepsInner) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *JobStepsInner) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *JobStepsInner) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetStartedAt

`func (o *JobStepsInner) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *JobStepsInner) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *JobStepsInner) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *JobStepsInner) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *JobStepsInner) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *JobStepsInner) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *JobStepsInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *JobStepsInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *JobStepsInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *JobStepsInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *JobStepsInner) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *JobStepsInner) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


