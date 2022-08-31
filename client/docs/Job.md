# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the job. | 
**RunId** | **int32** | The id of the associated workflow run. | 
**RunUrl** | **string** |  | 
**RunAttempt** | Pointer to **int32** | Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run. | [optional] 
**NodeId** | **string** |  | 
**HeadSha** | **string** | The SHA of the commit that is being run. | 
**Url** | **string** |  | 
**HtmlUrl** | **NullableString** |  | 
**Status** | **string** | The phase of the lifecycle that the job is currently in. | 
**Conclusion** | **NullableString** | The outcome of the job. | 
**StartedAt** | **time.Time** | The time that the job started, in ISO 8601 format. | 
**CompletedAt** | **NullableTime** | The time that the job finished, in ISO 8601 format. | 
**Name** | **string** | The name of the job. | 
**Steps** | Pointer to [**[]JobStepsInner**](JobStepsInner.md) | Steps in this job. | [optional] 
**CheckRunUrl** | **string** |  | 
**Labels** | **[]string** | Labels for the workflow job. Specified by the \&quot;runs_on\&quot; attribute in the action&#39;s workflow file. | 
**RunnerId** | **NullableInt32** | The ID of the runner to which this job has been assigned. (If a runner hasn&#39;t yet been assigned, this will be null.) | 
**RunnerName** | **NullableString** | The name of the runner to which this job has been assigned. (If a runner hasn&#39;t yet been assigned, this will be null.) | 
**RunnerGroupId** | **NullableInt32** | The ID of the runner group to which this job has been assigned. (If a runner hasn&#39;t yet been assigned, this will be null.) | 
**RunnerGroupName** | **NullableString** | The name of the runner group to which this job has been assigned. (If a runner hasn&#39;t yet been assigned, this will be null.) | 

## Methods

### NewJob

`func NewJob(id int32, runId int32, runUrl string, nodeId string, headSha string, url string, htmlUrl NullableString, status string, conclusion NullableString, startedAt time.Time, completedAt NullableTime, name string, checkRunUrl string, labels []string, runnerId NullableInt32, runnerName NullableString, runnerGroupId NullableInt32, runnerGroupName NullableString, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v int32)`

SetId sets Id field to given value.


### GetRunId

`func (o *Job) GetRunId() int32`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *Job) GetRunIdOk() (*int32, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *Job) SetRunId(v int32)`

SetRunId sets RunId field to given value.


### GetRunUrl

`func (o *Job) GetRunUrl() string`

GetRunUrl returns the RunUrl field if non-nil, zero value otherwise.

### GetRunUrlOk

`func (o *Job) GetRunUrlOk() (*string, bool)`

GetRunUrlOk returns a tuple with the RunUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunUrl

`func (o *Job) SetRunUrl(v string)`

SetRunUrl sets RunUrl field to given value.


### GetRunAttempt

`func (o *Job) GetRunAttempt() int32`

GetRunAttempt returns the RunAttempt field if non-nil, zero value otherwise.

### GetRunAttemptOk

`func (o *Job) GetRunAttemptOk() (*int32, bool)`

GetRunAttemptOk returns a tuple with the RunAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAttempt

`func (o *Job) SetRunAttempt(v int32)`

SetRunAttempt sets RunAttempt field to given value.

### HasRunAttempt

`func (o *Job) HasRunAttempt() bool`

HasRunAttempt returns a boolean if a field has been set.

### GetNodeId

`func (o *Job) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Job) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Job) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetHeadSha

`func (o *Job) GetHeadSha() string`

GetHeadSha returns the HeadSha field if non-nil, zero value otherwise.

### GetHeadShaOk

`func (o *Job) GetHeadShaOk() (*string, bool)`

GetHeadShaOk returns a tuple with the HeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSha

`func (o *Job) SetHeadSha(v string)`

SetHeadSha sets HeadSha field to given value.


### GetUrl

`func (o *Job) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Job) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Job) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *Job) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Job) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Job) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### SetHtmlUrlNil

`func (o *Job) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *Job) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
### GetStatus

`func (o *Job) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetConclusion

`func (o *Job) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *Job) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *Job) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.


### SetConclusionNil

`func (o *Job) SetConclusionNil(b bool)`

 SetConclusionNil sets the value for Conclusion to be an explicit nil

### UnsetConclusion
`func (o *Job) UnsetConclusion()`

UnsetConclusion ensures that no value is present for Conclusion, not even an explicit nil
### GetStartedAt

`func (o *Job) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Job) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Job) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *Job) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Job) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Job) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *Job) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Job) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetName

`func (o *Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Job) SetName(v string)`

SetName sets Name field to given value.


### GetSteps

`func (o *Job) GetSteps() []JobStepsInner`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Job) GetStepsOk() (*[]JobStepsInner, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Job) SetSteps(v []JobStepsInner)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Job) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetCheckRunUrl

`func (o *Job) GetCheckRunUrl() string`

GetCheckRunUrl returns the CheckRunUrl field if non-nil, zero value otherwise.

### GetCheckRunUrlOk

`func (o *Job) GetCheckRunUrlOk() (*string, bool)`

GetCheckRunUrlOk returns a tuple with the CheckRunUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckRunUrl

`func (o *Job) SetCheckRunUrl(v string)`

SetCheckRunUrl sets CheckRunUrl field to given value.


### GetLabels

`func (o *Job) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Job) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Job) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetRunnerId

`func (o *Job) GetRunnerId() int32`

GetRunnerId returns the RunnerId field if non-nil, zero value otherwise.

### GetRunnerIdOk

`func (o *Job) GetRunnerIdOk() (*int32, bool)`

GetRunnerIdOk returns a tuple with the RunnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerId

`func (o *Job) SetRunnerId(v int32)`

SetRunnerId sets RunnerId field to given value.


### SetRunnerIdNil

`func (o *Job) SetRunnerIdNil(b bool)`

 SetRunnerIdNil sets the value for RunnerId to be an explicit nil

### UnsetRunnerId
`func (o *Job) UnsetRunnerId()`

UnsetRunnerId ensures that no value is present for RunnerId, not even an explicit nil
### GetRunnerName

`func (o *Job) GetRunnerName() string`

GetRunnerName returns the RunnerName field if non-nil, zero value otherwise.

### GetRunnerNameOk

`func (o *Job) GetRunnerNameOk() (*string, bool)`

GetRunnerNameOk returns a tuple with the RunnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerName

`func (o *Job) SetRunnerName(v string)`

SetRunnerName sets RunnerName field to given value.


### SetRunnerNameNil

`func (o *Job) SetRunnerNameNil(b bool)`

 SetRunnerNameNil sets the value for RunnerName to be an explicit nil

### UnsetRunnerName
`func (o *Job) UnsetRunnerName()`

UnsetRunnerName ensures that no value is present for RunnerName, not even an explicit nil
### GetRunnerGroupId

`func (o *Job) GetRunnerGroupId() int32`

GetRunnerGroupId returns the RunnerGroupId field if non-nil, zero value otherwise.

### GetRunnerGroupIdOk

`func (o *Job) GetRunnerGroupIdOk() (*int32, bool)`

GetRunnerGroupIdOk returns a tuple with the RunnerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerGroupId

`func (o *Job) SetRunnerGroupId(v int32)`

SetRunnerGroupId sets RunnerGroupId field to given value.


### SetRunnerGroupIdNil

`func (o *Job) SetRunnerGroupIdNil(b bool)`

 SetRunnerGroupIdNil sets the value for RunnerGroupId to be an explicit nil

### UnsetRunnerGroupId
`func (o *Job) UnsetRunnerGroupId()`

UnsetRunnerGroupId ensures that no value is present for RunnerGroupId, not even an explicit nil
### GetRunnerGroupName

`func (o *Job) GetRunnerGroupName() string`

GetRunnerGroupName returns the RunnerGroupName field if non-nil, zero value otherwise.

### GetRunnerGroupNameOk

`func (o *Job) GetRunnerGroupNameOk() (*string, bool)`

GetRunnerGroupNameOk returns a tuple with the RunnerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerGroupName

`func (o *Job) SetRunnerGroupName(v string)`

SetRunnerGroupName sets RunnerGroupName field to given value.


### SetRunnerGroupNameNil

`func (o *Job) SetRunnerGroupNameNil(b bool)`

 SetRunnerGroupNameNil sets the value for RunnerGroupName to be an explicit nil

### UnsetRunnerGroupName
`func (o *Job) UnsetRunnerGroupName()`

UnsetRunnerGroupName ensures that no value is present for RunnerGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


