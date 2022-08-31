# IssueComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the issue comment | 
**NodeId** | **string** |  | 
**Url** | **string** | URL for the issue comment | 
**Body** | Pointer to **string** | Contents of the issue comment | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 
**BodyHtml** | Pointer to **string** |  | [optional] 
**HtmlUrl** | **string** |  | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**IssueUrl** | **string** |  | 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 
**PerformedViaGithubApp** | Pointer to [**NullableNullableIntegration**](NullableIntegration.md) |  | [optional] 
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 

## Methods

### NewIssueComment

`func NewIssueComment(id int32, nodeId string, url string, htmlUrl string, user NullableNullableSimpleUser, createdAt time.Time, updatedAt time.Time, issueUrl string, authorAssociation AuthorAssociation, ) *IssueComment`

NewIssueComment instantiates a new IssueComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCommentWithDefaults

`func NewIssueCommentWithDefaults() *IssueComment`

NewIssueCommentWithDefaults instantiates a new IssueComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueComment) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *IssueComment) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *IssueComment) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *IssueComment) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *IssueComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IssueComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IssueComment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBody

`func (o *IssueComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *IssueComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *IssueComment) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *IssueComment) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodyText

`func (o *IssueComment) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *IssueComment) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *IssueComment) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *IssueComment) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetBodyHtml

`func (o *IssueComment) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *IssueComment) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *IssueComment) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *IssueComment) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *IssueComment) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *IssueComment) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *IssueComment) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetUser

`func (o *IssueComment) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IssueComment) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IssueComment) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *IssueComment) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *IssueComment) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetCreatedAt

`func (o *IssueComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IssueComment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueComment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueComment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIssueUrl

`func (o *IssueComment) GetIssueUrl() string`

GetIssueUrl returns the IssueUrl field if non-nil, zero value otherwise.

### GetIssueUrlOk

`func (o *IssueComment) GetIssueUrlOk() (*string, bool)`

GetIssueUrlOk returns a tuple with the IssueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueUrl

`func (o *IssueComment) SetIssueUrl(v string)`

SetIssueUrl sets IssueUrl field to given value.


### GetAuthorAssociation

`func (o *IssueComment) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *IssueComment) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *IssueComment) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetPerformedViaGithubApp

`func (o *IssueComment) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *IssueComment) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *IssueComment) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.

### HasPerformedViaGithubApp

`func (o *IssueComment) HasPerformedViaGithubApp() bool`

HasPerformedViaGithubApp returns a boolean if a field has been set.

### SetPerformedViaGithubAppNil

`func (o *IssueComment) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *IssueComment) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetReactions

`func (o *IssueComment) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *IssueComment) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *IssueComment) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *IssueComment) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


