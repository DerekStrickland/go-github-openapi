# PendingDeploymentReviewersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DeploymentReviewerType**](DeploymentReviewerType.md) |  | [optional] 
**Reviewer** | Pointer to [**PendingDeploymentReviewersInnerReviewer**](PendingDeploymentReviewersInnerReviewer.md) |  | [optional] 

## Methods

### NewPendingDeploymentReviewersInner

`func NewPendingDeploymentReviewersInner() *PendingDeploymentReviewersInner`

NewPendingDeploymentReviewersInner instantiates a new PendingDeploymentReviewersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingDeploymentReviewersInnerWithDefaults

`func NewPendingDeploymentReviewersInnerWithDefaults() *PendingDeploymentReviewersInner`

NewPendingDeploymentReviewersInnerWithDefaults instantiates a new PendingDeploymentReviewersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PendingDeploymentReviewersInner) GetType() DeploymentReviewerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PendingDeploymentReviewersInner) GetTypeOk() (*DeploymentReviewerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PendingDeploymentReviewersInner) SetType(v DeploymentReviewerType)`

SetType sets Type field to given value.

### HasType

`func (o *PendingDeploymentReviewersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReviewer

`func (o *PendingDeploymentReviewersInner) GetReviewer() PendingDeploymentReviewersInnerReviewer`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *PendingDeploymentReviewersInner) GetReviewerOk() (*PendingDeploymentReviewersInnerReviewer, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *PendingDeploymentReviewersInner) SetReviewer(v PendingDeploymentReviewersInnerReviewer)`

SetReviewer sets Reviewer field to given value.

### HasReviewer

`func (o *PendingDeploymentReviewersInner) HasReviewer() bool`

HasReviewer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


