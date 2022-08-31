# TimelineCommentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Actor** | [**SimpleUser**](SimpleUser.md) |  | 
**Id** | **int32** | Unique identifier of the issue comment | 
**NodeId** | **string** |  | 
**Url** | **string** | URL for the issue comment | 
**Body** | Pointer to **string** | Contents of the issue comment | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 
**BodyHtml** | Pointer to **string** |  | [optional] 
**HtmlUrl** | **string** |  | 
**User** | [**SimpleUser**](SimpleUser.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**IssueUrl** | **string** |  | 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 
**PerformedViaGithubApp** | Pointer to [**NullableNullableIntegration**](NullableIntegration.md) |  | [optional] 
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 

## Methods

### NewTimelineCommentEvent

`func NewTimelineCommentEvent(event string, actor SimpleUser, id int32, nodeId string, url string, htmlUrl string, user SimpleUser, createdAt time.Time, updatedAt time.Time, issueUrl string, authorAssociation AuthorAssociation, ) *TimelineCommentEvent`

NewTimelineCommentEvent instantiates a new TimelineCommentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineCommentEventWithDefaults

`func NewTimelineCommentEventWithDefaults() *TimelineCommentEvent`

NewTimelineCommentEventWithDefaults instantiates a new TimelineCommentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *TimelineCommentEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TimelineCommentEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TimelineCommentEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetActor

`func (o *TimelineCommentEvent) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *TimelineCommentEvent) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *TimelineCommentEvent) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.


### GetId

`func (o *TimelineCommentEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimelineCommentEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimelineCommentEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *TimelineCommentEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TimelineCommentEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TimelineCommentEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *TimelineCommentEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TimelineCommentEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TimelineCommentEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBody

`func (o *TimelineCommentEvent) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TimelineCommentEvent) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TimelineCommentEvent) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TimelineCommentEvent) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodyText

`func (o *TimelineCommentEvent) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *TimelineCommentEvent) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *TimelineCommentEvent) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *TimelineCommentEvent) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetBodyHtml

`func (o *TimelineCommentEvent) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *TimelineCommentEvent) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *TimelineCommentEvent) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *TimelineCommentEvent) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *TimelineCommentEvent) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TimelineCommentEvent) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TimelineCommentEvent) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetUser

`func (o *TimelineCommentEvent) GetUser() SimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TimelineCommentEvent) GetUserOk() (*SimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TimelineCommentEvent) SetUser(v SimpleUser)`

SetUser sets User field to given value.


### GetCreatedAt

`func (o *TimelineCommentEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimelineCommentEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimelineCommentEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TimelineCommentEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TimelineCommentEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TimelineCommentEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIssueUrl

`func (o *TimelineCommentEvent) GetIssueUrl() string`

GetIssueUrl returns the IssueUrl field if non-nil, zero value otherwise.

### GetIssueUrlOk

`func (o *TimelineCommentEvent) GetIssueUrlOk() (*string, bool)`

GetIssueUrlOk returns a tuple with the IssueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueUrl

`func (o *TimelineCommentEvent) SetIssueUrl(v string)`

SetIssueUrl sets IssueUrl field to given value.


### GetAuthorAssociation

`func (o *TimelineCommentEvent) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *TimelineCommentEvent) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *TimelineCommentEvent) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetPerformedViaGithubApp

`func (o *TimelineCommentEvent) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *TimelineCommentEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *TimelineCommentEvent) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.

### HasPerformedViaGithubApp

`func (o *TimelineCommentEvent) HasPerformedViaGithubApp() bool`

HasPerformedViaGithubApp returns a boolean if a field has been set.

### SetPerformedViaGithubAppNil

`func (o *TimelineCommentEvent) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *TimelineCommentEvent) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetReactions

`func (o *TimelineCommentEvent) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *TimelineCommentEvent) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *TimelineCommentEvent) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *TimelineCommentEvent) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


