# ChecksUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the check. For example, \&quot;code-coverage\&quot;. | [optional] 
**DetailsUrl** | Pointer to **string** | The URL of the integrator&#39;s site that has the full details of the check. | [optional] 
**ExternalId** | Pointer to **string** | A reference for the run on the integrator&#39;s system. | [optional] 
**StartedAt** | Pointer to **time.Time** | This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 
**Status** | Pointer to **string** | The current status. | [optional] 
**Conclusion** | Pointer to **string** | **Required if you provide &#x60;completed_at&#x60; or a &#x60;status&#x60; of &#x60;completed&#x60;**. The final conclusion of the check.  **Note:** Providing &#x60;conclusion&#x60; will automatically set the &#x60;status&#x60; parameter to &#x60;completed&#x60;. You cannot change a check run conclusion to &#x60;stale&#x60;, only GitHub can set this. | [optional] 
**CompletedAt** | Pointer to **time.Time** | The time the check completed. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 
**Output** | Pointer to [**ChecksUpdateRequestOutput**](ChecksUpdateRequestOutput.md) |  | [optional] 
**Actions** | Pointer to [**[]ChecksCreateRequestActionsInner**](ChecksCreateRequestActionsInner.md) | Possible further actions the integrator can perform, which a user may trigger. Each action includes a &#x60;label&#x60;, &#x60;identifier&#x60; and &#x60;description&#x60;. A maximum of three actions are accepted. See the [&#x60;actions&#x60; object](https://docs.github.com/rest/reference/checks#actions-object) description. To learn more about check runs and requested actions, see \&quot;[Check runs and requested actions](https://docs.github.com/rest/reference/checks#check-runs-and-requested-actions).\&quot; | [optional] 

## Methods

### NewChecksUpdateRequest

`func NewChecksUpdateRequest() *ChecksUpdateRequest`

NewChecksUpdateRequest instantiates a new ChecksUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksUpdateRequestWithDefaults

`func NewChecksUpdateRequestWithDefaults() *ChecksUpdateRequest`

NewChecksUpdateRequestWithDefaults instantiates a new ChecksUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChecksUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChecksUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChecksUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChecksUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDetailsUrl

`func (o *ChecksUpdateRequest) GetDetailsUrl() string`

GetDetailsUrl returns the DetailsUrl field if non-nil, zero value otherwise.

### GetDetailsUrlOk

`func (o *ChecksUpdateRequest) GetDetailsUrlOk() (*string, bool)`

GetDetailsUrlOk returns a tuple with the DetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsUrl

`func (o *ChecksUpdateRequest) SetDetailsUrl(v string)`

SetDetailsUrl sets DetailsUrl field to given value.

### HasDetailsUrl

`func (o *ChecksUpdateRequest) HasDetailsUrl() bool`

HasDetailsUrl returns a boolean if a field has been set.

### GetExternalId

`func (o *ChecksUpdateRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ChecksUpdateRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ChecksUpdateRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ChecksUpdateRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetStartedAt

`func (o *ChecksUpdateRequest) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ChecksUpdateRequest) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ChecksUpdateRequest) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ChecksUpdateRequest) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ChecksUpdateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChecksUpdateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChecksUpdateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChecksUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConclusion

`func (o *ChecksUpdateRequest) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *ChecksUpdateRequest) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *ChecksUpdateRequest) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *ChecksUpdateRequest) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ChecksUpdateRequest) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ChecksUpdateRequest) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ChecksUpdateRequest) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ChecksUpdateRequest) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetOutput

`func (o *ChecksUpdateRequest) GetOutput() ChecksUpdateRequestOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ChecksUpdateRequest) GetOutputOk() (*ChecksUpdateRequestOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ChecksUpdateRequest) SetOutput(v ChecksUpdateRequestOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ChecksUpdateRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetActions

`func (o *ChecksUpdateRequest) GetActions() []ChecksCreateRequestActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ChecksUpdateRequest) GetActionsOk() (*[]ChecksCreateRequestActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ChecksUpdateRequest) SetActions(v []ChecksCreateRequestActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ChecksUpdateRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


