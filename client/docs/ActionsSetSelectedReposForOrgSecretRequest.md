# ActionsSetSelectedReposForOrgSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedRepositoryIds** | **[]int32** | An array of repository ids that can access the organization secret. You can only provide a list of repository ids when the &#x60;visibility&#x60; is set to &#x60;selected&#x60;. You can add and remove individual repositories using the [Set selected repositories for an organization secret](https://docs.github.com/rest/reference/actions#set-selected-repositories-for-an-organization-secret) and [Remove selected repository from an organization secret](https://docs.github.com/rest/reference/actions#remove-selected-repository-from-an-organization-secret) endpoints. | 

## Methods

### NewActionsSetSelectedReposForOrgSecretRequest

`func NewActionsSetSelectedReposForOrgSecretRequest(selectedRepositoryIds []int32, ) *ActionsSetSelectedReposForOrgSecretRequest`

NewActionsSetSelectedReposForOrgSecretRequest instantiates a new ActionsSetSelectedReposForOrgSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsSetSelectedReposForOrgSecretRequestWithDefaults

`func NewActionsSetSelectedReposForOrgSecretRequestWithDefaults() *ActionsSetSelectedReposForOrgSecretRequest`

NewActionsSetSelectedReposForOrgSecretRequestWithDefaults instantiates a new ActionsSetSelectedReposForOrgSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedRepositoryIds

`func (o *ActionsSetSelectedReposForOrgSecretRequest) GetSelectedRepositoryIds() []int32`

GetSelectedRepositoryIds returns the SelectedRepositoryIds field if non-nil, zero value otherwise.

### GetSelectedRepositoryIdsOk

`func (o *ActionsSetSelectedReposForOrgSecretRequest) GetSelectedRepositoryIdsOk() (*[]int32, bool)`

GetSelectedRepositoryIdsOk returns a tuple with the SelectedRepositoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositoryIds

`func (o *ActionsSetSelectedReposForOrgSecretRequest) SetSelectedRepositoryIds(v []int32)`

SetSelectedRepositoryIds sets SelectedRepositoryIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


