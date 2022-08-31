# ReposListDeploymentBranchPolicies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The number of deployment branch policies for the environment. | 
**BranchPolicies** | [**[]DeploymentBranchPolicy**](DeploymentBranchPolicy.md) |  | 

## Methods

### NewReposListDeploymentBranchPolicies200Response

`func NewReposListDeploymentBranchPolicies200Response(totalCount int32, branchPolicies []DeploymentBranchPolicy, ) *ReposListDeploymentBranchPolicies200Response`

NewReposListDeploymentBranchPolicies200Response instantiates a new ReposListDeploymentBranchPolicies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposListDeploymentBranchPolicies200ResponseWithDefaults

`func NewReposListDeploymentBranchPolicies200ResponseWithDefaults() *ReposListDeploymentBranchPolicies200Response`

NewReposListDeploymentBranchPolicies200ResponseWithDefaults instantiates a new ReposListDeploymentBranchPolicies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ReposListDeploymentBranchPolicies200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ReposListDeploymentBranchPolicies200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ReposListDeploymentBranchPolicies200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetBranchPolicies

`func (o *ReposListDeploymentBranchPolicies200Response) GetBranchPolicies() []DeploymentBranchPolicy`

GetBranchPolicies returns the BranchPolicies field if non-nil, zero value otherwise.

### GetBranchPoliciesOk

`func (o *ReposListDeploymentBranchPolicies200Response) GetBranchPoliciesOk() (*[]DeploymentBranchPolicy, bool)`

GetBranchPoliciesOk returns a tuple with the BranchPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchPolicies

`func (o *ReposListDeploymentBranchPolicies200Response) SetBranchPolicies(v []DeploymentBranchPolicy)`

SetBranchPolicies sets BranchPolicies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


