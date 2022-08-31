# IssuesAddAssigneesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to **[]string** | Usernames of people to assign this issue to. _NOTE: Only users with push access can add assignees to an issue. Assignees are silently ignored otherwise._ | [optional] 

## Methods

### NewIssuesAddAssigneesRequest

`func NewIssuesAddAssigneesRequest() *IssuesAddAssigneesRequest`

NewIssuesAddAssigneesRequest instantiates a new IssuesAddAssigneesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesAddAssigneesRequestWithDefaults

`func NewIssuesAddAssigneesRequestWithDefaults() *IssuesAddAssigneesRequest`

NewIssuesAddAssigneesRequestWithDefaults instantiates a new IssuesAddAssigneesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *IssuesAddAssigneesRequest) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *IssuesAddAssigneesRequest) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *IssuesAddAssigneesRequest) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *IssuesAddAssigneesRequest) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


