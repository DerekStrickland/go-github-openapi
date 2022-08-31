# TeamsCreateDiscussionInOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The discussion post&#39;s title. | 
**Body** | **string** | The discussion post&#39;s body text. | 
**Private** | Pointer to **bool** | Private posts are only visible to team members, organization owners, and team maintainers. Public posts are visible to all members of the organization. Set to &#x60;true&#x60; to create a private post. | [optional] [default to false]

## Methods

### NewTeamsCreateDiscussionInOrgRequest

`func NewTeamsCreateDiscussionInOrgRequest(title string, body string, ) *TeamsCreateDiscussionInOrgRequest`

NewTeamsCreateDiscussionInOrgRequest instantiates a new TeamsCreateDiscussionInOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsCreateDiscussionInOrgRequestWithDefaults

`func NewTeamsCreateDiscussionInOrgRequestWithDefaults() *TeamsCreateDiscussionInOrgRequest`

NewTeamsCreateDiscussionInOrgRequestWithDefaults instantiates a new TeamsCreateDiscussionInOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TeamsCreateDiscussionInOrgRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TeamsCreateDiscussionInOrgRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TeamsCreateDiscussionInOrgRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *TeamsCreateDiscussionInOrgRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TeamsCreateDiscussionInOrgRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TeamsCreateDiscussionInOrgRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetPrivate

`func (o *TeamsCreateDiscussionInOrgRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TeamsCreateDiscussionInOrgRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TeamsCreateDiscussionInOrgRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *TeamsCreateDiscussionInOrgRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


