# OrganizationSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Url** | **string** |  | 
**ReposUrl** | **string** |  | 
**EventsUrl** | **string** |  | 
**HooksUrl** | **string** |  | 
**IssuesUrl** | **string** |  | 
**MembersUrl** | **string** |  | 
**PublicMembersUrl** | **string** |  | 
**AvatarUrl** | **string** |  | 
**Description** | **NullableString** |  | 

## Methods

### NewOrganizationSimple

`func NewOrganizationSimple(login string, id int32, nodeId string, url string, reposUrl string, eventsUrl string, hooksUrl string, issuesUrl string, membersUrl string, publicMembersUrl string, avatarUrl string, description NullableString, ) *OrganizationSimple`

NewOrganizationSimple instantiates a new OrganizationSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSimpleWithDefaults

`func NewOrganizationSimpleWithDefaults() *OrganizationSimple`

NewOrganizationSimpleWithDefaults instantiates a new OrganizationSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *OrganizationSimple) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *OrganizationSimple) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *OrganizationSimple) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetId

`func (o *OrganizationSimple) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationSimple) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationSimple) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *OrganizationSimple) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *OrganizationSimple) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *OrganizationSimple) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *OrganizationSimple) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrganizationSimple) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrganizationSimple) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetReposUrl

`func (o *OrganizationSimple) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *OrganizationSimple) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *OrganizationSimple) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.


### GetEventsUrl

`func (o *OrganizationSimple) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *OrganizationSimple) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *OrganizationSimple) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetHooksUrl

`func (o *OrganizationSimple) GetHooksUrl() string`

GetHooksUrl returns the HooksUrl field if non-nil, zero value otherwise.

### GetHooksUrlOk

`func (o *OrganizationSimple) GetHooksUrlOk() (*string, bool)`

GetHooksUrlOk returns a tuple with the HooksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooksUrl

`func (o *OrganizationSimple) SetHooksUrl(v string)`

SetHooksUrl sets HooksUrl field to given value.


### GetIssuesUrl

`func (o *OrganizationSimple) GetIssuesUrl() string`

GetIssuesUrl returns the IssuesUrl field if non-nil, zero value otherwise.

### GetIssuesUrlOk

`func (o *OrganizationSimple) GetIssuesUrlOk() (*string, bool)`

GetIssuesUrlOk returns a tuple with the IssuesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesUrl

`func (o *OrganizationSimple) SetIssuesUrl(v string)`

SetIssuesUrl sets IssuesUrl field to given value.


### GetMembersUrl

`func (o *OrganizationSimple) GetMembersUrl() string`

GetMembersUrl returns the MembersUrl field if non-nil, zero value otherwise.

### GetMembersUrlOk

`func (o *OrganizationSimple) GetMembersUrlOk() (*string, bool)`

GetMembersUrlOk returns a tuple with the MembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersUrl

`func (o *OrganizationSimple) SetMembersUrl(v string)`

SetMembersUrl sets MembersUrl field to given value.


### GetPublicMembersUrl

`func (o *OrganizationSimple) GetPublicMembersUrl() string`

GetPublicMembersUrl returns the PublicMembersUrl field if non-nil, zero value otherwise.

### GetPublicMembersUrlOk

`func (o *OrganizationSimple) GetPublicMembersUrlOk() (*string, bool)`

GetPublicMembersUrlOk returns a tuple with the PublicMembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicMembersUrl

`func (o *OrganizationSimple) SetPublicMembersUrl(v string)`

SetPublicMembersUrl sets PublicMembersUrl field to given value.


### GetAvatarUrl

`func (o *OrganizationSimple) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *OrganizationSimple) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *OrganizationSimple) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetDescription

`func (o *OrganizationSimple) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationSimple) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationSimple) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *OrganizationSimple) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrganizationSimple) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


