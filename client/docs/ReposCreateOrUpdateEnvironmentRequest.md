# ReposCreateOrUpdateEnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaitTimer** | Pointer to **int32** | The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days). | [optional] 
**Reviewers** | Pointer to [**[]ReposCreateOrUpdateEnvironmentRequestReviewersInner**](ReposCreateOrUpdateEnvironmentRequestReviewersInner.md) | The people or teams that may review jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed. | [optional] 
**DeploymentBranchPolicy** | Pointer to [**NullableDeploymentBranchPolicySettings**](DeploymentBranchPolicySettings.md) |  | [optional] 

## Methods

### NewReposCreateOrUpdateEnvironmentRequest

`func NewReposCreateOrUpdateEnvironmentRequest() *ReposCreateOrUpdateEnvironmentRequest`

NewReposCreateOrUpdateEnvironmentRequest instantiates a new ReposCreateOrUpdateEnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateOrUpdateEnvironmentRequestWithDefaults

`func NewReposCreateOrUpdateEnvironmentRequestWithDefaults() *ReposCreateOrUpdateEnvironmentRequest`

NewReposCreateOrUpdateEnvironmentRequestWithDefaults instantiates a new ReposCreateOrUpdateEnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaitTimer

`func (o *ReposCreateOrUpdateEnvironmentRequest) GetWaitTimer() int32`

GetWaitTimer returns the WaitTimer field if non-nil, zero value otherwise.

### GetWaitTimerOk

`func (o *ReposCreateOrUpdateEnvironmentRequest) GetWaitTimerOk() (*int32, bool)`

GetWaitTimerOk returns a tuple with the WaitTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimer

`func (o *ReposCreateOrUpdateEnvironmentRequest) SetWaitTimer(v int32)`

SetWaitTimer sets WaitTimer field to given value.

### HasWaitTimer

`func (o *ReposCreateOrUpdateEnvironmentRequest) HasWaitTimer() bool`

HasWaitTimer returns a boolean if a field has been set.

### GetReviewers

`func (o *ReposCreateOrUpdateEnvironmentRequest) GetReviewers() []ReposCreateOrUpdateEnvironmentRequestReviewersInner`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *ReposCreateOrUpdateEnvironmentRequest) GetReviewersOk() (*[]ReposCreateOrUpdateEnvironmentRequestReviewersInner, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *ReposCreateOrUpdateEnvironmentRequest) SetReviewers(v []ReposCreateOrUpdateEnvironmentRequestReviewersInner)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *ReposCreateOrUpdateEnvironmentRequest) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### SetReviewersNil

`func (o *ReposCreateOrUpdateEnvironmentRequest) SetReviewersNil(b bool)`

 SetReviewersNil sets the value for Reviewers to be an explicit nil

### UnsetReviewers
`func (o *ReposCreateOrUpdateEnvironmentRequest) UnsetReviewers()`

UnsetReviewers ensures that no value is present for Reviewers, not even an explicit nil
### GetDeploymentBranchPolicy

`func (o *ReposCreateOrUpdateEnvironmentRequest) GetDeploymentBranchPolicy() DeploymentBranchPolicySettings`

GetDeploymentBranchPolicy returns the DeploymentBranchPolicy field if non-nil, zero value otherwise.

### GetDeploymentBranchPolicyOk

`func (o *ReposCreateOrUpdateEnvironmentRequest) GetDeploymentBranchPolicyOk() (*DeploymentBranchPolicySettings, bool)`

GetDeploymentBranchPolicyOk returns a tuple with the DeploymentBranchPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentBranchPolicy

`func (o *ReposCreateOrUpdateEnvironmentRequest) SetDeploymentBranchPolicy(v DeploymentBranchPolicySettings)`

SetDeploymentBranchPolicy sets DeploymentBranchPolicy field to given value.

### HasDeploymentBranchPolicy

`func (o *ReposCreateOrUpdateEnvironmentRequest) HasDeploymentBranchPolicy() bool`

HasDeploymentBranchPolicy returns a boolean if a field has been set.

### SetDeploymentBranchPolicyNil

`func (o *ReposCreateOrUpdateEnvironmentRequest) SetDeploymentBranchPolicyNil(b bool)`

 SetDeploymentBranchPolicyNil sets the value for DeploymentBranchPolicy to be an explicit nil

### UnsetDeploymentBranchPolicy
`func (o *ReposCreateOrUpdateEnvironmentRequest) UnsetDeploymentBranchPolicy()`

UnsetDeploymentBranchPolicy ensures that no value is present for DeploymentBranchPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


