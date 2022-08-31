# Feed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimelineUrl** | **string** |  | 
**UserUrl** | **string** |  | 
**CurrentUserPublicUrl** | Pointer to **string** |  | [optional] 
**CurrentUserUrl** | Pointer to **string** |  | [optional] 
**CurrentUserActorUrl** | Pointer to **string** |  | [optional] 
**CurrentUserOrganizationUrl** | Pointer to **string** |  | [optional] 
**CurrentUserOrganizationUrls** | Pointer to **[]string** |  | [optional] 
**SecurityAdvisoriesUrl** | Pointer to **string** |  | [optional] 
**Links** | [**FeedLinks**](FeedLinks.md) |  | 

## Methods

### NewFeed

`func NewFeed(timelineUrl string, userUrl string, links FeedLinks, ) *Feed`

NewFeed instantiates a new Feed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedWithDefaults

`func NewFeedWithDefaults() *Feed`

NewFeedWithDefaults instantiates a new Feed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimelineUrl

`func (o *Feed) GetTimelineUrl() string`

GetTimelineUrl returns the TimelineUrl field if non-nil, zero value otherwise.

### GetTimelineUrlOk

`func (o *Feed) GetTimelineUrlOk() (*string, bool)`

GetTimelineUrlOk returns a tuple with the TimelineUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineUrl

`func (o *Feed) SetTimelineUrl(v string)`

SetTimelineUrl sets TimelineUrl field to given value.


### GetUserUrl

`func (o *Feed) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *Feed) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *Feed) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.


### GetCurrentUserPublicUrl

`func (o *Feed) GetCurrentUserPublicUrl() string`

GetCurrentUserPublicUrl returns the CurrentUserPublicUrl field if non-nil, zero value otherwise.

### GetCurrentUserPublicUrlOk

`func (o *Feed) GetCurrentUserPublicUrlOk() (*string, bool)`

GetCurrentUserPublicUrlOk returns a tuple with the CurrentUserPublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserPublicUrl

`func (o *Feed) SetCurrentUserPublicUrl(v string)`

SetCurrentUserPublicUrl sets CurrentUserPublicUrl field to given value.

### HasCurrentUserPublicUrl

`func (o *Feed) HasCurrentUserPublicUrl() bool`

HasCurrentUserPublicUrl returns a boolean if a field has been set.

### GetCurrentUserUrl

`func (o *Feed) GetCurrentUserUrl() string`

GetCurrentUserUrl returns the CurrentUserUrl field if non-nil, zero value otherwise.

### GetCurrentUserUrlOk

`func (o *Feed) GetCurrentUserUrlOk() (*string, bool)`

GetCurrentUserUrlOk returns a tuple with the CurrentUserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserUrl

`func (o *Feed) SetCurrentUserUrl(v string)`

SetCurrentUserUrl sets CurrentUserUrl field to given value.

### HasCurrentUserUrl

`func (o *Feed) HasCurrentUserUrl() bool`

HasCurrentUserUrl returns a boolean if a field has been set.

### GetCurrentUserActorUrl

`func (o *Feed) GetCurrentUserActorUrl() string`

GetCurrentUserActorUrl returns the CurrentUserActorUrl field if non-nil, zero value otherwise.

### GetCurrentUserActorUrlOk

`func (o *Feed) GetCurrentUserActorUrlOk() (*string, bool)`

GetCurrentUserActorUrlOk returns a tuple with the CurrentUserActorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserActorUrl

`func (o *Feed) SetCurrentUserActorUrl(v string)`

SetCurrentUserActorUrl sets CurrentUserActorUrl field to given value.

### HasCurrentUserActorUrl

`func (o *Feed) HasCurrentUserActorUrl() bool`

HasCurrentUserActorUrl returns a boolean if a field has been set.

### GetCurrentUserOrganizationUrl

`func (o *Feed) GetCurrentUserOrganizationUrl() string`

GetCurrentUserOrganizationUrl returns the CurrentUserOrganizationUrl field if non-nil, zero value otherwise.

### GetCurrentUserOrganizationUrlOk

`func (o *Feed) GetCurrentUserOrganizationUrlOk() (*string, bool)`

GetCurrentUserOrganizationUrlOk returns a tuple with the CurrentUserOrganizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserOrganizationUrl

`func (o *Feed) SetCurrentUserOrganizationUrl(v string)`

SetCurrentUserOrganizationUrl sets CurrentUserOrganizationUrl field to given value.

### HasCurrentUserOrganizationUrl

`func (o *Feed) HasCurrentUserOrganizationUrl() bool`

HasCurrentUserOrganizationUrl returns a boolean if a field has been set.

### GetCurrentUserOrganizationUrls

`func (o *Feed) GetCurrentUserOrganizationUrls() []string`

GetCurrentUserOrganizationUrls returns the CurrentUserOrganizationUrls field if non-nil, zero value otherwise.

### GetCurrentUserOrganizationUrlsOk

`func (o *Feed) GetCurrentUserOrganizationUrlsOk() (*[]string, bool)`

GetCurrentUserOrganizationUrlsOk returns a tuple with the CurrentUserOrganizationUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserOrganizationUrls

`func (o *Feed) SetCurrentUserOrganizationUrls(v []string)`

SetCurrentUserOrganizationUrls sets CurrentUserOrganizationUrls field to given value.

### HasCurrentUserOrganizationUrls

`func (o *Feed) HasCurrentUserOrganizationUrls() bool`

HasCurrentUserOrganizationUrls returns a boolean if a field has been set.

### GetSecurityAdvisoriesUrl

`func (o *Feed) GetSecurityAdvisoriesUrl() string`

GetSecurityAdvisoriesUrl returns the SecurityAdvisoriesUrl field if non-nil, zero value otherwise.

### GetSecurityAdvisoriesUrlOk

`func (o *Feed) GetSecurityAdvisoriesUrlOk() (*string, bool)`

GetSecurityAdvisoriesUrlOk returns a tuple with the SecurityAdvisoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAdvisoriesUrl

`func (o *Feed) SetSecurityAdvisoriesUrl(v string)`

SetSecurityAdvisoriesUrl sets SecurityAdvisoriesUrl field to given value.

### HasSecurityAdvisoriesUrl

`func (o *Feed) HasSecurityAdvisoriesUrl() bool`

HasSecurityAdvisoriesUrl returns a boolean if a field has been set.

### GetLinks

`func (o *Feed) GetLinks() FeedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Feed) GetLinksOk() (*FeedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Feed) SetLinks(v FeedLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


