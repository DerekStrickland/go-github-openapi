# CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedRepositoryIds** | **[]int32** | An array of repository ids for which a codespace can access the secret. You can manage the list of selected repositories using the [List selected repositories for a user secret](https://docs.github.com/rest/reference/codespaces#list-selected-repositories-for-a-user-secret), [Add a selected repository to a user secret](https://docs.github.com/rest/reference/codespaces#add-a-selected-repository-to-a-user-secret), and [Remove a selected repository from a user secret](https://docs.github.com/rest/reference/codespaces#remove-a-selected-repository-from-a-user-secret) endpoints. | 

## Methods

### NewCodespacesSetRepositoriesForSecretForAuthenticatedUserRequest

`func NewCodespacesSetRepositoriesForSecretForAuthenticatedUserRequest(selectedRepositoryIds []int32, ) *CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest`

NewCodespacesSetRepositoriesForSecretForAuthenticatedUserRequest instantiates a new CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesSetRepositoriesForSecretForAuthenticatedUserRequestWithDefaults

`func NewCodespacesSetRepositoriesForSecretForAuthenticatedUserRequestWithDefaults() *CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest`

NewCodespacesSetRepositoriesForSecretForAuthenticatedUserRequestWithDefaults instantiates a new CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedRepositoryIds

`func (o *CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest) GetSelectedRepositoryIds() []int32`

GetSelectedRepositoryIds returns the SelectedRepositoryIds field if non-nil, zero value otherwise.

### GetSelectedRepositoryIdsOk

`func (o *CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest) GetSelectedRepositoryIdsOk() (*[]int32, bool)`

GetSelectedRepositoryIdsOk returns a tuple with the SelectedRepositoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositoryIds

`func (o *CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest) SetSelectedRepositoryIds(v []int32)`

SetSelectedRepositoryIds sets SelectedRepositoryIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


