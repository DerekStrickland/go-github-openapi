# ChecksCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the check. For example, \&quot;code-coverage\&quot;. | 
**HeadSha** | **string** | The SHA of the commit. | 
**DetailsUrl** | Pointer to **string** | The URL of the integrator&#39;s site that has the full details of the check. If the integrator does not provide this, then the homepage of the GitHub app is used. | [optional] 
**ExternalId** | Pointer to **string** | A reference for the run on the integrator&#39;s system. | [optional] 
**Status** | Pointer to **string** | The current status. | [optional] [default to "queued"]
**StartedAt** | Pointer to **time.Time** | The time that the check run began. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 
**Conclusion** | Pointer to **string** | **Required if you provide &#x60;completed_at&#x60; or a &#x60;status&#x60; of &#x60;completed&#x60;**. The final conclusion of the check.  **Note:** Providing &#x60;conclusion&#x60; will automatically set the &#x60;status&#x60; parameter to &#x60;completed&#x60;. You cannot change a check run conclusion to &#x60;stale&#x60;, only GitHub can set this. | [optional] 
**CompletedAt** | Pointer to **time.Time** | The time the check completed. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 
**Output** | Pointer to [**ChecksCreateRequestOutput**](ChecksCreateRequestOutput.md) |  | [optional] 
**Actions** | Pointer to [**[]ChecksCreateRequestActionsInner**](ChecksCreateRequestActionsInner.md) | Displays a button on GitHub that can be clicked to alert your app to do additional tasks. For example, a code linting app can display a button that automatically fixes detected errors. The button created in this object is displayed after the check run completes. When a user clicks the button, GitHub sends the [&#x60;check_run.requested_action&#x60; webhook](https://docs.github.com/webhooks/event-payloads/#check_run) to your app. Each action includes a &#x60;label&#x60;, &#x60;identifier&#x60; and &#x60;description&#x60;. A maximum of three actions are accepted. See the [&#x60;actions&#x60; object](https://docs.github.com/rest/reference/checks#actions-object) description. To learn more about check runs and requested actions, see \&quot;[Check runs and requested actions](https://docs.github.com/rest/reference/checks#check-runs-and-requested-actions).\&quot; | [optional] 

## Methods

### NewChecksCreateRequest

`func NewChecksCreateRequest(name string, headSha string, ) *ChecksCreateRequest`

NewChecksCreateRequest instantiates a new ChecksCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksCreateRequestWithDefaults

`func NewChecksCreateRequestWithDefaults() *ChecksCreateRequest`

NewChecksCreateRequestWithDefaults instantiates a new ChecksCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChecksCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChecksCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChecksCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetHeadSha

`func (o *ChecksCreateRequest) GetHeadSha() string`

GetHeadSha returns the HeadSha field if non-nil, zero value otherwise.

### GetHeadShaOk

`func (o *ChecksCreateRequest) GetHeadShaOk() (*string, bool)`

GetHeadShaOk returns a tuple with the HeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSha

`func (o *ChecksCreateRequest) SetHeadSha(v string)`

SetHeadSha sets HeadSha field to given value.


### GetDetailsUrl

`func (o *ChecksCreateRequest) GetDetailsUrl() string`

GetDetailsUrl returns the DetailsUrl field if non-nil, zero value otherwise.

### GetDetailsUrlOk

`func (o *ChecksCreateRequest) GetDetailsUrlOk() (*string, bool)`

GetDetailsUrlOk returns a tuple with the DetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsUrl

`func (o *ChecksCreateRequest) SetDetailsUrl(v string)`

SetDetailsUrl sets DetailsUrl field to given value.

### HasDetailsUrl

`func (o *ChecksCreateRequest) HasDetailsUrl() bool`

HasDetailsUrl returns a boolean if a field has been set.

### GetExternalId

`func (o *ChecksCreateRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ChecksCreateRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ChecksCreateRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ChecksCreateRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetStatus

`func (o *ChecksCreateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChecksCreateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChecksCreateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChecksCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *ChecksCreateRequest) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ChecksCreateRequest) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ChecksCreateRequest) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ChecksCreateRequest) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetConclusion

`func (o *ChecksCreateRequest) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *ChecksCreateRequest) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *ChecksCreateRequest) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *ChecksCreateRequest) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ChecksCreateRequest) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ChecksCreateRequest) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ChecksCreateRequest) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ChecksCreateRequest) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetOutput

`func (o *ChecksCreateRequest) GetOutput() ChecksCreateRequestOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ChecksCreateRequest) GetOutputOk() (*ChecksCreateRequestOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ChecksCreateRequest) SetOutput(v ChecksCreateRequestOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ChecksCreateRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetActions

`func (o *ChecksCreateRequest) GetActions() []ChecksCreateRequestActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ChecksCreateRequest) GetActionsOk() (*[]ChecksCreateRequestActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ChecksCreateRequest) SetActions(v []ChecksCreateRequestActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ChecksCreateRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


