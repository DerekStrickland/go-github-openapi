# PullRequestReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the review | 
**NodeId** | **string** |  | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Body** | **string** | The text of the review. | 
**State** | **string** |  | 
**HtmlUrl** | **string** |  | 
**PullRequestUrl** | **string** |  | 
**Links** | [**TimelineReviewedEventLinks**](TimelineReviewedEventLinks.md) |  | 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**CommitId** | **string** | A commit SHA for the review. | 
**BodyHtml** | Pointer to **string** |  | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 

## Methods

### NewPullRequestReview

`func NewPullRequestReview(id int32, nodeId string, user NullableNullableSimpleUser, body string, state string, htmlUrl string, pullRequestUrl string, links TimelineReviewedEventLinks, commitId string, authorAssociation AuthorAssociation, ) *PullRequestReview`

NewPullRequestReview instantiates a new PullRequestReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestReviewWithDefaults

`func NewPullRequestReviewWithDefaults() *PullRequestReview`

NewPullRequestReviewWithDefaults instantiates a new PullRequestReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PullRequestReview) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequestReview) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequestReview) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *PullRequestReview) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PullRequestReview) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PullRequestReview) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUser

`func (o *PullRequestReview) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PullRequestReview) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PullRequestReview) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *PullRequestReview) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PullRequestReview) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetBody

`func (o *PullRequestReview) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullRequestReview) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullRequestReview) SetBody(v string)`

SetBody sets Body field to given value.


### GetState

`func (o *PullRequestReview) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PullRequestReview) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PullRequestReview) SetState(v string)`

SetState sets State field to given value.


### GetHtmlUrl

`func (o *PullRequestReview) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *PullRequestReview) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *PullRequestReview) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetPullRequestUrl

`func (o *PullRequestReview) GetPullRequestUrl() string`

GetPullRequestUrl returns the PullRequestUrl field if non-nil, zero value otherwise.

### GetPullRequestUrlOk

`func (o *PullRequestReview) GetPullRequestUrlOk() (*string, bool)`

GetPullRequestUrlOk returns a tuple with the PullRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestUrl

`func (o *PullRequestReview) SetPullRequestUrl(v string)`

SetPullRequestUrl sets PullRequestUrl field to given value.


### GetLinks

`func (o *PullRequestReview) GetLinks() TimelineReviewedEventLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PullRequestReview) GetLinksOk() (*TimelineReviewedEventLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PullRequestReview) SetLinks(v TimelineReviewedEventLinks)`

SetLinks sets Links field to given value.


### GetSubmittedAt

`func (o *PullRequestReview) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *PullRequestReview) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *PullRequestReview) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *PullRequestReview) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetCommitId

`func (o *PullRequestReview) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *PullRequestReview) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *PullRequestReview) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetBodyHtml

`func (o *PullRequestReview) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *PullRequestReview) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *PullRequestReview) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *PullRequestReview) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetBodyText

`func (o *PullRequestReview) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *PullRequestReview) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *PullRequestReview) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *PullRequestReview) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetAuthorAssociation

`func (o *PullRequestReview) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *PullRequestReview) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *PullRequestReview) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


