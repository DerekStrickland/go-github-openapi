# GitCommitAuthor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** | Timestamp of the commit | 
**Email** | **string** | Git email address of the user | 
**Name** | **string** | Name of the git user | 

## Methods

### NewGitCommitAuthor

`func NewGitCommitAuthor(date time.Time, email string, name string, ) *GitCommitAuthor`

NewGitCommitAuthor instantiates a new GitCommitAuthor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCommitAuthorWithDefaults

`func NewGitCommitAuthorWithDefaults() *GitCommitAuthor`

NewGitCommitAuthorWithDefaults instantiates a new GitCommitAuthor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GitCommitAuthor) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GitCommitAuthor) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GitCommitAuthor) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetEmail

`func (o *GitCommitAuthor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GitCommitAuthor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GitCommitAuthor) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *GitCommitAuthor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitCommitAuthor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitCommitAuthor) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


