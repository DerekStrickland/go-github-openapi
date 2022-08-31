# ReposCreateOrUpdateFileContentsRequestAuthor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the author or committer of the commit. You&#39;ll receive a &#x60;422&#x60; status code if &#x60;name&#x60; is omitted. | 
**Email** | **string** | The email of the author or committer of the commit. You&#39;ll receive a &#x60;422&#x60; status code if &#x60;email&#x60; is omitted. | 
**Date** | Pointer to **string** |  | [optional] 

## Methods

### NewReposCreateOrUpdateFileContentsRequestAuthor

`func NewReposCreateOrUpdateFileContentsRequestAuthor(name string, email string, ) *ReposCreateOrUpdateFileContentsRequestAuthor`

NewReposCreateOrUpdateFileContentsRequestAuthor instantiates a new ReposCreateOrUpdateFileContentsRequestAuthor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateOrUpdateFileContentsRequestAuthorWithDefaults

`func NewReposCreateOrUpdateFileContentsRequestAuthorWithDefaults() *ReposCreateOrUpdateFileContentsRequestAuthor`

NewReposCreateOrUpdateFileContentsRequestAuthorWithDefaults instantiates a new ReposCreateOrUpdateFileContentsRequestAuthor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReposCreateOrUpdateFileContentsRequestAuthor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposCreateOrUpdateFileContentsRequestAuthor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposCreateOrUpdateFileContentsRequestAuthor) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *ReposCreateOrUpdateFileContentsRequestAuthor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ReposCreateOrUpdateFileContentsRequestAuthor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ReposCreateOrUpdateFileContentsRequestAuthor) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDate

`func (o *ReposCreateOrUpdateFileContentsRequestAuthor) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ReposCreateOrUpdateFileContentsRequestAuthor) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ReposCreateOrUpdateFileContentsRequestAuthor) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ReposCreateOrUpdateFileContentsRequestAuthor) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


