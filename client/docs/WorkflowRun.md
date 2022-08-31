# WorkflowRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the workflow run. | 
**Name** | Pointer to **NullableString** | The name of the workflow run. | [optional] 
**NodeId** | **string** |  | 
**CheckSuiteId** | Pointer to **int32** | The ID of the associated check suite. | [optional] 
**CheckSuiteNodeId** | Pointer to **string** | The node ID of the associated check suite. | [optional] 
**HeadBranch** | **NullableString** |  | 
**HeadSha** | **string** | The SHA of the head commit that points to the version of the workflow being run. | 
**Path** | **string** | The full path of the workflow | 
**RunNumber** | **int32** | The auto incrementing run number for the workflow run. | 
**RunAttempt** | Pointer to **int32** | Attempt number of the run, 1 for first attempt and higher if the workflow was re-run. | [optional] 
**ReferencedWorkflows** | Pointer to [**[]ReferencedWorkflow**](ReferencedWorkflow.md) |  | [optional] 
**Event** | **string** |  | 
**Status** | **NullableString** |  | 
**Conclusion** | **NullableString** |  | 
**WorkflowId** | **int32** | The ID of the parent workflow. | 
**Url** | **string** | The URL to the workflow run. | 
**HtmlUrl** | **string** |  | 
**PullRequests** | [**[]PullRequestMinimal**](PullRequestMinimal.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Actor** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 
**TriggeringActor** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 
**RunStartedAt** | Pointer to **time.Time** | The start time of the latest run. Resets on re-run. | [optional] 
**JobsUrl** | **string** | The URL to the jobs for the workflow run. | 
**LogsUrl** | **string** | The URL to download the logs for the workflow run. | 
**CheckSuiteUrl** | **string** | The URL to the associated check suite. | 
**ArtifactsUrl** | **string** | The URL to the artifacts for the workflow run. | 
**CancelUrl** | **string** | The URL to cancel the workflow run. | 
**RerunUrl** | **string** | The URL to rerun the workflow run. | 
**PreviousAttemptUrl** | Pointer to **NullableString** | The URL to the previous attempted run of this workflow, if one exists. | [optional] 
**WorkflowUrl** | **string** | The URL to the workflow. | 
**HeadCommit** | [**NullableNullableSimpleCommit**](NullableSimpleCommit.md) |  | 
**Repository** | [**MinimalRepository**](MinimalRepository.md) |  | 
**HeadRepository** | [**MinimalRepository**](MinimalRepository.md) |  | 
**HeadRepositoryId** | Pointer to **int32** |  | [optional] 

## Methods

### NewWorkflowRun

`func NewWorkflowRun(id int32, nodeId string, headBranch NullableString, headSha string, path string, runNumber int32, event string, status NullableString, conclusion NullableString, workflowId int32, url string, htmlUrl string, pullRequests []PullRequestMinimal, createdAt time.Time, updatedAt time.Time, jobsUrl string, logsUrl string, checkSuiteUrl string, artifactsUrl string, cancelUrl string, rerunUrl string, workflowUrl string, headCommit NullableNullableSimpleCommit, repository MinimalRepository, headRepository MinimalRepository, ) *WorkflowRun`

NewWorkflowRun instantiates a new WorkflowRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunWithDefaults

`func NewWorkflowRunWithDefaults() *WorkflowRun`

NewWorkflowRunWithDefaults instantiates a new WorkflowRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowRun) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowRun) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowRun) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowRun) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowRun) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkflowRun) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkflowRun) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNodeId

`func (o *WorkflowRun) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *WorkflowRun) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *WorkflowRun) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetCheckSuiteId

`func (o *WorkflowRun) GetCheckSuiteId() int32`

GetCheckSuiteId returns the CheckSuiteId field if non-nil, zero value otherwise.

### GetCheckSuiteIdOk

`func (o *WorkflowRun) GetCheckSuiteIdOk() (*int32, bool)`

GetCheckSuiteIdOk returns a tuple with the CheckSuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSuiteId

`func (o *WorkflowRun) SetCheckSuiteId(v int32)`

SetCheckSuiteId sets CheckSuiteId field to given value.

### HasCheckSuiteId

`func (o *WorkflowRun) HasCheckSuiteId() bool`

HasCheckSuiteId returns a boolean if a field has been set.

### GetCheckSuiteNodeId

`func (o *WorkflowRun) GetCheckSuiteNodeId() string`

GetCheckSuiteNodeId returns the CheckSuiteNodeId field if non-nil, zero value otherwise.

### GetCheckSuiteNodeIdOk

`func (o *WorkflowRun) GetCheckSuiteNodeIdOk() (*string, bool)`

GetCheckSuiteNodeIdOk returns a tuple with the CheckSuiteNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSuiteNodeId

`func (o *WorkflowRun) SetCheckSuiteNodeId(v string)`

SetCheckSuiteNodeId sets CheckSuiteNodeId field to given value.

### HasCheckSuiteNodeId

`func (o *WorkflowRun) HasCheckSuiteNodeId() bool`

HasCheckSuiteNodeId returns a boolean if a field has been set.

### GetHeadBranch

`func (o *WorkflowRun) GetHeadBranch() string`

GetHeadBranch returns the HeadBranch field if non-nil, zero value otherwise.

### GetHeadBranchOk

`func (o *WorkflowRun) GetHeadBranchOk() (*string, bool)`

GetHeadBranchOk returns a tuple with the HeadBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadBranch

`func (o *WorkflowRun) SetHeadBranch(v string)`

SetHeadBranch sets HeadBranch field to given value.


### SetHeadBranchNil

`func (o *WorkflowRun) SetHeadBranchNil(b bool)`

 SetHeadBranchNil sets the value for HeadBranch to be an explicit nil

### UnsetHeadBranch
`func (o *WorkflowRun) UnsetHeadBranch()`

UnsetHeadBranch ensures that no value is present for HeadBranch, not even an explicit nil
### GetHeadSha

`func (o *WorkflowRun) GetHeadSha() string`

GetHeadSha returns the HeadSha field if non-nil, zero value otherwise.

### GetHeadShaOk

`func (o *WorkflowRun) GetHeadShaOk() (*string, bool)`

GetHeadShaOk returns a tuple with the HeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSha

`func (o *WorkflowRun) SetHeadSha(v string)`

SetHeadSha sets HeadSha field to given value.


### GetPath

`func (o *WorkflowRun) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WorkflowRun) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WorkflowRun) SetPath(v string)`

SetPath sets Path field to given value.


### GetRunNumber

`func (o *WorkflowRun) GetRunNumber() int32`

GetRunNumber returns the RunNumber field if non-nil, zero value otherwise.

### GetRunNumberOk

`func (o *WorkflowRun) GetRunNumberOk() (*int32, bool)`

GetRunNumberOk returns a tuple with the RunNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNumber

`func (o *WorkflowRun) SetRunNumber(v int32)`

SetRunNumber sets RunNumber field to given value.


### GetRunAttempt

`func (o *WorkflowRun) GetRunAttempt() int32`

GetRunAttempt returns the RunAttempt field if non-nil, zero value otherwise.

### GetRunAttemptOk

`func (o *WorkflowRun) GetRunAttemptOk() (*int32, bool)`

GetRunAttemptOk returns a tuple with the RunAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAttempt

`func (o *WorkflowRun) SetRunAttempt(v int32)`

SetRunAttempt sets RunAttempt field to given value.

### HasRunAttempt

`func (o *WorkflowRun) HasRunAttempt() bool`

HasRunAttempt returns a boolean if a field has been set.

### GetReferencedWorkflows

`func (o *WorkflowRun) GetReferencedWorkflows() []ReferencedWorkflow`

GetReferencedWorkflows returns the ReferencedWorkflows field if non-nil, zero value otherwise.

### GetReferencedWorkflowsOk

`func (o *WorkflowRun) GetReferencedWorkflowsOk() (*[]ReferencedWorkflow, bool)`

GetReferencedWorkflowsOk returns a tuple with the ReferencedWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedWorkflows

`func (o *WorkflowRun) SetReferencedWorkflows(v []ReferencedWorkflow)`

SetReferencedWorkflows sets ReferencedWorkflows field to given value.

### HasReferencedWorkflows

`func (o *WorkflowRun) HasReferencedWorkflows() bool`

HasReferencedWorkflows returns a boolean if a field has been set.

### SetReferencedWorkflowsNil

`func (o *WorkflowRun) SetReferencedWorkflowsNil(b bool)`

 SetReferencedWorkflowsNil sets the value for ReferencedWorkflows to be an explicit nil

### UnsetReferencedWorkflows
`func (o *WorkflowRun) UnsetReferencedWorkflows()`

UnsetReferencedWorkflows ensures that no value is present for ReferencedWorkflows, not even an explicit nil
### GetEvent

`func (o *WorkflowRun) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WorkflowRun) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WorkflowRun) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetStatus

`func (o *WorkflowRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowRun) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *WorkflowRun) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *WorkflowRun) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetConclusion

`func (o *WorkflowRun) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *WorkflowRun) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *WorkflowRun) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.


### SetConclusionNil

`func (o *WorkflowRun) SetConclusionNil(b bool)`

 SetConclusionNil sets the value for Conclusion to be an explicit nil

### UnsetConclusion
`func (o *WorkflowRun) UnsetConclusion()`

UnsetConclusion ensures that no value is present for Conclusion, not even an explicit nil
### GetWorkflowId

`func (o *WorkflowRun) GetWorkflowId() int32`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowRun) GetWorkflowIdOk() (*int32, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowRun) SetWorkflowId(v int32)`

SetWorkflowId sets WorkflowId field to given value.


### GetUrl

`func (o *WorkflowRun) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WorkflowRun) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WorkflowRun) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *WorkflowRun) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *WorkflowRun) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *WorkflowRun) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetPullRequests

`func (o *WorkflowRun) GetPullRequests() []PullRequestMinimal`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *WorkflowRun) GetPullRequestsOk() (*[]PullRequestMinimal, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *WorkflowRun) SetPullRequests(v []PullRequestMinimal)`

SetPullRequests sets PullRequests field to given value.


### SetPullRequestsNil

`func (o *WorkflowRun) SetPullRequestsNil(b bool)`

 SetPullRequestsNil sets the value for PullRequests to be an explicit nil

### UnsetPullRequests
`func (o *WorkflowRun) UnsetPullRequests()`

UnsetPullRequests ensures that no value is present for PullRequests, not even an explicit nil
### GetCreatedAt

`func (o *WorkflowRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WorkflowRun) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkflowRun) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkflowRun) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetActor

`func (o *WorkflowRun) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *WorkflowRun) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *WorkflowRun) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.

### HasActor

`func (o *WorkflowRun) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetTriggeringActor

`func (o *WorkflowRun) GetTriggeringActor() SimpleUser`

GetTriggeringActor returns the TriggeringActor field if non-nil, zero value otherwise.

### GetTriggeringActorOk

`func (o *WorkflowRun) GetTriggeringActorOk() (*SimpleUser, bool)`

GetTriggeringActorOk returns a tuple with the TriggeringActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeringActor

`func (o *WorkflowRun) SetTriggeringActor(v SimpleUser)`

SetTriggeringActor sets TriggeringActor field to given value.

### HasTriggeringActor

`func (o *WorkflowRun) HasTriggeringActor() bool`

HasTriggeringActor returns a boolean if a field has been set.

### GetRunStartedAt

`func (o *WorkflowRun) GetRunStartedAt() time.Time`

GetRunStartedAt returns the RunStartedAt field if non-nil, zero value otherwise.

### GetRunStartedAtOk

`func (o *WorkflowRun) GetRunStartedAtOk() (*time.Time, bool)`

GetRunStartedAtOk returns a tuple with the RunStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStartedAt

`func (o *WorkflowRun) SetRunStartedAt(v time.Time)`

SetRunStartedAt sets RunStartedAt field to given value.

### HasRunStartedAt

`func (o *WorkflowRun) HasRunStartedAt() bool`

HasRunStartedAt returns a boolean if a field has been set.

### GetJobsUrl

`func (o *WorkflowRun) GetJobsUrl() string`

GetJobsUrl returns the JobsUrl field if non-nil, zero value otherwise.

### GetJobsUrlOk

`func (o *WorkflowRun) GetJobsUrlOk() (*string, bool)`

GetJobsUrlOk returns a tuple with the JobsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsUrl

`func (o *WorkflowRun) SetJobsUrl(v string)`

SetJobsUrl sets JobsUrl field to given value.


### GetLogsUrl

`func (o *WorkflowRun) GetLogsUrl() string`

GetLogsUrl returns the LogsUrl field if non-nil, zero value otherwise.

### GetLogsUrlOk

`func (o *WorkflowRun) GetLogsUrlOk() (*string, bool)`

GetLogsUrlOk returns a tuple with the LogsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsUrl

`func (o *WorkflowRun) SetLogsUrl(v string)`

SetLogsUrl sets LogsUrl field to given value.


### GetCheckSuiteUrl

`func (o *WorkflowRun) GetCheckSuiteUrl() string`

GetCheckSuiteUrl returns the CheckSuiteUrl field if non-nil, zero value otherwise.

### GetCheckSuiteUrlOk

`func (o *WorkflowRun) GetCheckSuiteUrlOk() (*string, bool)`

GetCheckSuiteUrlOk returns a tuple with the CheckSuiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSuiteUrl

`func (o *WorkflowRun) SetCheckSuiteUrl(v string)`

SetCheckSuiteUrl sets CheckSuiteUrl field to given value.


### GetArtifactsUrl

`func (o *WorkflowRun) GetArtifactsUrl() string`

GetArtifactsUrl returns the ArtifactsUrl field if non-nil, zero value otherwise.

### GetArtifactsUrlOk

`func (o *WorkflowRun) GetArtifactsUrlOk() (*string, bool)`

GetArtifactsUrlOk returns a tuple with the ArtifactsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactsUrl

`func (o *WorkflowRun) SetArtifactsUrl(v string)`

SetArtifactsUrl sets ArtifactsUrl field to given value.


### GetCancelUrl

`func (o *WorkflowRun) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *WorkflowRun) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *WorkflowRun) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### GetRerunUrl

`func (o *WorkflowRun) GetRerunUrl() string`

GetRerunUrl returns the RerunUrl field if non-nil, zero value otherwise.

### GetRerunUrlOk

`func (o *WorkflowRun) GetRerunUrlOk() (*string, bool)`

GetRerunUrlOk returns a tuple with the RerunUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunUrl

`func (o *WorkflowRun) SetRerunUrl(v string)`

SetRerunUrl sets RerunUrl field to given value.


### GetPreviousAttemptUrl

`func (o *WorkflowRun) GetPreviousAttemptUrl() string`

GetPreviousAttemptUrl returns the PreviousAttemptUrl field if non-nil, zero value otherwise.

### GetPreviousAttemptUrlOk

`func (o *WorkflowRun) GetPreviousAttemptUrlOk() (*string, bool)`

GetPreviousAttemptUrlOk returns a tuple with the PreviousAttemptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAttemptUrl

`func (o *WorkflowRun) SetPreviousAttemptUrl(v string)`

SetPreviousAttemptUrl sets PreviousAttemptUrl field to given value.

### HasPreviousAttemptUrl

`func (o *WorkflowRun) HasPreviousAttemptUrl() bool`

HasPreviousAttemptUrl returns a boolean if a field has been set.

### SetPreviousAttemptUrlNil

`func (o *WorkflowRun) SetPreviousAttemptUrlNil(b bool)`

 SetPreviousAttemptUrlNil sets the value for PreviousAttemptUrl to be an explicit nil

### UnsetPreviousAttemptUrl
`func (o *WorkflowRun) UnsetPreviousAttemptUrl()`

UnsetPreviousAttemptUrl ensures that no value is present for PreviousAttemptUrl, not even an explicit nil
### GetWorkflowUrl

`func (o *WorkflowRun) GetWorkflowUrl() string`

GetWorkflowUrl returns the WorkflowUrl field if non-nil, zero value otherwise.

### GetWorkflowUrlOk

`func (o *WorkflowRun) GetWorkflowUrlOk() (*string, bool)`

GetWorkflowUrlOk returns a tuple with the WorkflowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowUrl

`func (o *WorkflowRun) SetWorkflowUrl(v string)`

SetWorkflowUrl sets WorkflowUrl field to given value.


### GetHeadCommit

`func (o *WorkflowRun) GetHeadCommit() NullableSimpleCommit`

GetHeadCommit returns the HeadCommit field if non-nil, zero value otherwise.

### GetHeadCommitOk

`func (o *WorkflowRun) GetHeadCommitOk() (*NullableSimpleCommit, bool)`

GetHeadCommitOk returns a tuple with the HeadCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadCommit

`func (o *WorkflowRun) SetHeadCommit(v NullableSimpleCommit)`

SetHeadCommit sets HeadCommit field to given value.


### SetHeadCommitNil

`func (o *WorkflowRun) SetHeadCommitNil(b bool)`

 SetHeadCommitNil sets the value for HeadCommit to be an explicit nil

### UnsetHeadCommit
`func (o *WorkflowRun) UnsetHeadCommit()`

UnsetHeadCommit ensures that no value is present for HeadCommit, not even an explicit nil
### GetRepository

`func (o *WorkflowRun) GetRepository() MinimalRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *WorkflowRun) GetRepositoryOk() (*MinimalRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *WorkflowRun) SetRepository(v MinimalRepository)`

SetRepository sets Repository field to given value.


### GetHeadRepository

`func (o *WorkflowRun) GetHeadRepository() MinimalRepository`

GetHeadRepository returns the HeadRepository field if non-nil, zero value otherwise.

### GetHeadRepositoryOk

`func (o *WorkflowRun) GetHeadRepositoryOk() (*MinimalRepository, bool)`

GetHeadRepositoryOk returns a tuple with the HeadRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadRepository

`func (o *WorkflowRun) SetHeadRepository(v MinimalRepository)`

SetHeadRepository sets HeadRepository field to given value.


### GetHeadRepositoryId

`func (o *WorkflowRun) GetHeadRepositoryId() int32`

GetHeadRepositoryId returns the HeadRepositoryId field if non-nil, zero value otherwise.

### GetHeadRepositoryIdOk

`func (o *WorkflowRun) GetHeadRepositoryIdOk() (*int32, bool)`

GetHeadRepositoryIdOk returns a tuple with the HeadRepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadRepositoryId

`func (o *WorkflowRun) SetHeadRepositoryId(v int32)`

SetHeadRepositoryId sets HeadRepositoryId field to given value.

### HasHeadRepositoryId

`func (o *WorkflowRun) HasHeadRepositoryId() bool`

HasHeadRepositoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


