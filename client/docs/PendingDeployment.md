# PendingDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**PendingDeploymentEnvironment**](PendingDeploymentEnvironment.md) |  | 
**WaitTimer** | **int32** | The set duration of the wait timer | 
**WaitTimerStartedAt** | **NullableTime** | The time that the wait timer began. | 
**CurrentUserCanApprove** | **bool** | Whether the currently authenticated user can approve the deployment | 
**Reviewers** | [**[]PendingDeploymentReviewersInner**](PendingDeploymentReviewersInner.md) | The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed. | 

## Methods

### NewPendingDeployment

`func NewPendingDeployment(environment PendingDeploymentEnvironment, waitTimer int32, waitTimerStartedAt NullableTime, currentUserCanApprove bool, reviewers []PendingDeploymentReviewersInner, ) *PendingDeployment`

NewPendingDeployment instantiates a new PendingDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingDeploymentWithDefaults

`func NewPendingDeploymentWithDefaults() *PendingDeployment`

NewPendingDeploymentWithDefaults instantiates a new PendingDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *PendingDeployment) GetEnvironment() PendingDeploymentEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PendingDeployment) GetEnvironmentOk() (*PendingDeploymentEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PendingDeployment) SetEnvironment(v PendingDeploymentEnvironment)`

SetEnvironment sets Environment field to given value.


### GetWaitTimer

`func (o *PendingDeployment) GetWaitTimer() int32`

GetWaitTimer returns the WaitTimer field if non-nil, zero value otherwise.

### GetWaitTimerOk

`func (o *PendingDeployment) GetWaitTimerOk() (*int32, bool)`

GetWaitTimerOk returns a tuple with the WaitTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimer

`func (o *PendingDeployment) SetWaitTimer(v int32)`

SetWaitTimer sets WaitTimer field to given value.


### GetWaitTimerStartedAt

`func (o *PendingDeployment) GetWaitTimerStartedAt() time.Time`

GetWaitTimerStartedAt returns the WaitTimerStartedAt field if non-nil, zero value otherwise.

### GetWaitTimerStartedAtOk

`func (o *PendingDeployment) GetWaitTimerStartedAtOk() (*time.Time, bool)`

GetWaitTimerStartedAtOk returns a tuple with the WaitTimerStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimerStartedAt

`func (o *PendingDeployment) SetWaitTimerStartedAt(v time.Time)`

SetWaitTimerStartedAt sets WaitTimerStartedAt field to given value.


### SetWaitTimerStartedAtNil

`func (o *PendingDeployment) SetWaitTimerStartedAtNil(b bool)`

 SetWaitTimerStartedAtNil sets the value for WaitTimerStartedAt to be an explicit nil

### UnsetWaitTimerStartedAt
`func (o *PendingDeployment) UnsetWaitTimerStartedAt()`

UnsetWaitTimerStartedAt ensures that no value is present for WaitTimerStartedAt, not even an explicit nil
### GetCurrentUserCanApprove

`func (o *PendingDeployment) GetCurrentUserCanApprove() bool`

GetCurrentUserCanApprove returns the CurrentUserCanApprove field if non-nil, zero value otherwise.

### GetCurrentUserCanApproveOk

`func (o *PendingDeployment) GetCurrentUserCanApproveOk() (*bool, bool)`

GetCurrentUserCanApproveOk returns a tuple with the CurrentUserCanApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserCanApprove

`func (o *PendingDeployment) SetCurrentUserCanApprove(v bool)`

SetCurrentUserCanApprove sets CurrentUserCanApprove field to given value.


### GetReviewers

`func (o *PendingDeployment) GetReviewers() []PendingDeploymentReviewersInner`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *PendingDeployment) GetReviewersOk() (*[]PendingDeploymentReviewersInner, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *PendingDeployment) SetReviewers(v []PendingDeploymentReviewersInner)`

SetReviewers sets Reviewers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


