# StarredRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StarredAt** | **time.Time** |  | 
**Repo** | [**Repository**](Repository.md) |  | 

## Methods

### NewStarredRepository

`func NewStarredRepository(starredAt time.Time, repo Repository, ) *StarredRepository`

NewStarredRepository instantiates a new StarredRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStarredRepositoryWithDefaults

`func NewStarredRepositoryWithDefaults() *StarredRepository`

NewStarredRepositoryWithDefaults instantiates a new StarredRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStarredAt

`func (o *StarredRepository) GetStarredAt() time.Time`

GetStarredAt returns the StarredAt field if non-nil, zero value otherwise.

### GetStarredAtOk

`func (o *StarredRepository) GetStarredAtOk() (*time.Time, bool)`

GetStarredAtOk returns a tuple with the StarredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredAt

`func (o *StarredRepository) SetStarredAt(v time.Time)`

SetStarredAt sets StarredAt field to given value.


### GetRepo

`func (o *StarredRepository) GetRepo() Repository`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *StarredRepository) GetRepoOk() (*Repository, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *StarredRepository) SetRepo(v Repository)`

SetRepo sets Repo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


