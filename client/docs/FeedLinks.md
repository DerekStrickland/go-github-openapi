# FeedLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeline** | [**LinkWithType**](LinkWithType.md) |  | 
**User** | [**LinkWithType**](LinkWithType.md) |  | 
**SecurityAdvisories** | Pointer to [**LinkWithType**](LinkWithType.md) |  | [optional] 
**CurrentUser** | Pointer to [**LinkWithType**](LinkWithType.md) |  | [optional] 
**CurrentUserPublic** | Pointer to [**LinkWithType**](LinkWithType.md) |  | [optional] 
**CurrentUserActor** | Pointer to [**LinkWithType**](LinkWithType.md) |  | [optional] 
**CurrentUserOrganization** | Pointer to [**LinkWithType**](LinkWithType.md) |  | [optional] 
**CurrentUserOrganizations** | Pointer to [**[]LinkWithType**](LinkWithType.md) |  | [optional] 

## Methods

### NewFeedLinks

`func NewFeedLinks(timeline LinkWithType, user LinkWithType, ) *FeedLinks`

NewFeedLinks instantiates a new FeedLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedLinksWithDefaults

`func NewFeedLinksWithDefaults() *FeedLinks`

NewFeedLinksWithDefaults instantiates a new FeedLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeline

`func (o *FeedLinks) GetTimeline() LinkWithType`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *FeedLinks) GetTimelineOk() (*LinkWithType, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *FeedLinks) SetTimeline(v LinkWithType)`

SetTimeline sets Timeline field to given value.


### GetUser

`func (o *FeedLinks) GetUser() LinkWithType`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FeedLinks) GetUserOk() (*LinkWithType, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FeedLinks) SetUser(v LinkWithType)`

SetUser sets User field to given value.


### GetSecurityAdvisories

`func (o *FeedLinks) GetSecurityAdvisories() LinkWithType`

GetSecurityAdvisories returns the SecurityAdvisories field if non-nil, zero value otherwise.

### GetSecurityAdvisoriesOk

`func (o *FeedLinks) GetSecurityAdvisoriesOk() (*LinkWithType, bool)`

GetSecurityAdvisoriesOk returns a tuple with the SecurityAdvisories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAdvisories

`func (o *FeedLinks) SetSecurityAdvisories(v LinkWithType)`

SetSecurityAdvisories sets SecurityAdvisories field to given value.

### HasSecurityAdvisories

`func (o *FeedLinks) HasSecurityAdvisories() bool`

HasSecurityAdvisories returns a boolean if a field has been set.

### GetCurrentUser

`func (o *FeedLinks) GetCurrentUser() LinkWithType`

GetCurrentUser returns the CurrentUser field if non-nil, zero value otherwise.

### GetCurrentUserOk

`func (o *FeedLinks) GetCurrentUserOk() (*LinkWithType, bool)`

GetCurrentUserOk returns a tuple with the CurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUser

`func (o *FeedLinks) SetCurrentUser(v LinkWithType)`

SetCurrentUser sets CurrentUser field to given value.

### HasCurrentUser

`func (o *FeedLinks) HasCurrentUser() bool`

HasCurrentUser returns a boolean if a field has been set.

### GetCurrentUserPublic

`func (o *FeedLinks) GetCurrentUserPublic() LinkWithType`

GetCurrentUserPublic returns the CurrentUserPublic field if non-nil, zero value otherwise.

### GetCurrentUserPublicOk

`func (o *FeedLinks) GetCurrentUserPublicOk() (*LinkWithType, bool)`

GetCurrentUserPublicOk returns a tuple with the CurrentUserPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserPublic

`func (o *FeedLinks) SetCurrentUserPublic(v LinkWithType)`

SetCurrentUserPublic sets CurrentUserPublic field to given value.

### HasCurrentUserPublic

`func (o *FeedLinks) HasCurrentUserPublic() bool`

HasCurrentUserPublic returns a boolean if a field has been set.

### GetCurrentUserActor

`func (o *FeedLinks) GetCurrentUserActor() LinkWithType`

GetCurrentUserActor returns the CurrentUserActor field if non-nil, zero value otherwise.

### GetCurrentUserActorOk

`func (o *FeedLinks) GetCurrentUserActorOk() (*LinkWithType, bool)`

GetCurrentUserActorOk returns a tuple with the CurrentUserActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserActor

`func (o *FeedLinks) SetCurrentUserActor(v LinkWithType)`

SetCurrentUserActor sets CurrentUserActor field to given value.

### HasCurrentUserActor

`func (o *FeedLinks) HasCurrentUserActor() bool`

HasCurrentUserActor returns a boolean if a field has been set.

### GetCurrentUserOrganization

`func (o *FeedLinks) GetCurrentUserOrganization() LinkWithType`

GetCurrentUserOrganization returns the CurrentUserOrganization field if non-nil, zero value otherwise.

### GetCurrentUserOrganizationOk

`func (o *FeedLinks) GetCurrentUserOrganizationOk() (*LinkWithType, bool)`

GetCurrentUserOrganizationOk returns a tuple with the CurrentUserOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserOrganization

`func (o *FeedLinks) SetCurrentUserOrganization(v LinkWithType)`

SetCurrentUserOrganization sets CurrentUserOrganization field to given value.

### HasCurrentUserOrganization

`func (o *FeedLinks) HasCurrentUserOrganization() bool`

HasCurrentUserOrganization returns a boolean if a field has been set.

### GetCurrentUserOrganizations

`func (o *FeedLinks) GetCurrentUserOrganizations() []LinkWithType`

GetCurrentUserOrganizations returns the CurrentUserOrganizations field if non-nil, zero value otherwise.

### GetCurrentUserOrganizationsOk

`func (o *FeedLinks) GetCurrentUserOrganizationsOk() (*[]LinkWithType, bool)`

GetCurrentUserOrganizationsOk returns a tuple with the CurrentUserOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserOrganizations

`func (o *FeedLinks) SetCurrentUserOrganizations(v []LinkWithType)`

SetCurrentUserOrganizations sets CurrentUserOrganizations field to given value.

### HasCurrentUserOrganizations

`func (o *FeedLinks) HasCurrentUserOrganizations() bool`

HasCurrentUserOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


