# AppPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **string** | The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts. | [optional] 
**Administration** | Pointer to **string** | The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation. | [optional] 
**Checks** | Pointer to **string** | The level of permission to grant the access token for checks on code. | [optional] 
**Contents** | Pointer to **string** | The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges. | [optional] 
**Deployments** | Pointer to **string** | The level of permission to grant the access token for deployments and deployment statuses. | [optional] 
**Environments** | Pointer to **string** | The level of permission to grant the access token for managing repository environments. | [optional] 
**Issues** | Pointer to **string** | The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones. | [optional] 
**Metadata** | Pointer to **string** | The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata. | [optional] 
**Packages** | Pointer to **string** | The level of permission to grant the access token for packages published to GitHub Packages. | [optional] 
**Pages** | Pointer to **string** | The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds. | [optional] 
**PullRequests** | Pointer to **string** | The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges. | [optional] 
**RepositoryHooks** | Pointer to **string** | The level of permission to grant the access token to manage the post-receive hooks for a repository. | [optional] 
**RepositoryProjects** | Pointer to **string** | The level of permission to grant the access token to manage repository projects, columns, and cards. | [optional] 
**SecretScanningAlerts** | Pointer to **string** | The level of permission to grant the access token to view and manage secret scanning alerts. | [optional] 
**Secrets** | Pointer to **string** | The level of permission to grant the access token to manage repository secrets. | [optional] 
**SecurityEvents** | Pointer to **string** | The level of permission to grant the access token to view and manage security events like code scanning alerts. | [optional] 
**SingleFile** | Pointer to **string** | The level of permission to grant the access token to manage just a single file. | [optional] 
**Statuses** | Pointer to **string** | The level of permission to grant the access token for commit statuses. | [optional] 
**VulnerabilityAlerts** | Pointer to **string** | The level of permission to grant the access token to manage Dependabot alerts. | [optional] 
**Workflows** | Pointer to **string** | The level of permission to grant the access token to update GitHub Actions workflow files. | [optional] 
**Members** | Pointer to **string** | The level of permission to grant the access token for organization teams and members. | [optional] 
**OrganizationAdministration** | Pointer to **string** | The level of permission to grant the access token to manage access to an organization. | [optional] 
**OrganizationCustomRoles** | Pointer to **string** | The level of permission to grant the access token for custom roles management. This property is in beta and is subject to change. | [optional] 
**OrganizationHooks** | Pointer to **string** | The level of permission to grant the access token to manage the post-receive hooks for an organization. | [optional] 
**OrganizationPlan** | Pointer to **string** | The level of permission to grant the access token for viewing an organization&#39;s plan. | [optional] 
**OrganizationProjects** | Pointer to **string** | The level of permission to grant the access token to manage organization projects and projects beta (where available). | [optional] 
**OrganizationPackages** | Pointer to **string** | The level of permission to grant the access token for organization packages published to GitHub Packages. | [optional] 
**OrganizationSecrets** | Pointer to **string** | The level of permission to grant the access token to manage organization secrets. | [optional] 
**OrganizationSelfHostedRunners** | Pointer to **string** | The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization. | [optional] 
**OrganizationUserBlocking** | Pointer to **string** | The level of permission to grant the access token to view and manage users blocked by the organization. | [optional] 
**TeamDiscussions** | Pointer to **string** | The level of permission to grant the access token to manage team discussions and related comments. | [optional] 

## Methods

### NewAppPermissions

`func NewAppPermissions() *AppPermissions`

NewAppPermissions instantiates a new AppPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPermissionsWithDefaults

`func NewAppPermissionsWithDefaults() *AppPermissions`

NewAppPermissionsWithDefaults instantiates a new AppPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AppPermissions) GetActions() string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AppPermissions) GetActionsOk() (*string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AppPermissions) SetActions(v string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AppPermissions) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAdministration

`func (o *AppPermissions) GetAdministration() string`

GetAdministration returns the Administration field if non-nil, zero value otherwise.

### GetAdministrationOk

`func (o *AppPermissions) GetAdministrationOk() (*string, bool)`

GetAdministrationOk returns a tuple with the Administration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministration

`func (o *AppPermissions) SetAdministration(v string)`

SetAdministration sets Administration field to given value.

### HasAdministration

`func (o *AppPermissions) HasAdministration() bool`

HasAdministration returns a boolean if a field has been set.

### GetChecks

`func (o *AppPermissions) GetChecks() string`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *AppPermissions) GetChecksOk() (*string, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *AppPermissions) SetChecks(v string)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *AppPermissions) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetContents

`func (o *AppPermissions) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *AppPermissions) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *AppPermissions) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *AppPermissions) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetDeployments

`func (o *AppPermissions) GetDeployments() string`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *AppPermissions) GetDeploymentsOk() (*string, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *AppPermissions) SetDeployments(v string)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *AppPermissions) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetEnvironments

`func (o *AppPermissions) GetEnvironments() string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *AppPermissions) GetEnvironmentsOk() (*string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *AppPermissions) SetEnvironments(v string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *AppPermissions) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetIssues

`func (o *AppPermissions) GetIssues() string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *AppPermissions) GetIssuesOk() (*string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *AppPermissions) SetIssues(v string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *AppPermissions) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetMetadata

`func (o *AppPermissions) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppPermissions) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppPermissions) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AppPermissions) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPackages

`func (o *AppPermissions) GetPackages() string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AppPermissions) GetPackagesOk() (*string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AppPermissions) SetPackages(v string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AppPermissions) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetPages

`func (o *AppPermissions) GetPages() string`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *AppPermissions) GetPagesOk() (*string, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *AppPermissions) SetPages(v string)`

SetPages sets Pages field to given value.

### HasPages

`func (o *AppPermissions) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetPullRequests

`func (o *AppPermissions) GetPullRequests() string`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *AppPermissions) GetPullRequestsOk() (*string, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *AppPermissions) SetPullRequests(v string)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *AppPermissions) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.

### GetRepositoryHooks

`func (o *AppPermissions) GetRepositoryHooks() string`

GetRepositoryHooks returns the RepositoryHooks field if non-nil, zero value otherwise.

### GetRepositoryHooksOk

`func (o *AppPermissions) GetRepositoryHooksOk() (*string, bool)`

GetRepositoryHooksOk returns a tuple with the RepositoryHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryHooks

`func (o *AppPermissions) SetRepositoryHooks(v string)`

SetRepositoryHooks sets RepositoryHooks field to given value.

### HasRepositoryHooks

`func (o *AppPermissions) HasRepositoryHooks() bool`

HasRepositoryHooks returns a boolean if a field has been set.

### GetRepositoryProjects

`func (o *AppPermissions) GetRepositoryProjects() string`

GetRepositoryProjects returns the RepositoryProjects field if non-nil, zero value otherwise.

### GetRepositoryProjectsOk

`func (o *AppPermissions) GetRepositoryProjectsOk() (*string, bool)`

GetRepositoryProjectsOk returns a tuple with the RepositoryProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryProjects

`func (o *AppPermissions) SetRepositoryProjects(v string)`

SetRepositoryProjects sets RepositoryProjects field to given value.

### HasRepositoryProjects

`func (o *AppPermissions) HasRepositoryProjects() bool`

HasRepositoryProjects returns a boolean if a field has been set.

### GetSecretScanningAlerts

`func (o *AppPermissions) GetSecretScanningAlerts() string`

GetSecretScanningAlerts returns the SecretScanningAlerts field if non-nil, zero value otherwise.

### GetSecretScanningAlertsOk

`func (o *AppPermissions) GetSecretScanningAlertsOk() (*string, bool)`

GetSecretScanningAlertsOk returns a tuple with the SecretScanningAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScanningAlerts

`func (o *AppPermissions) SetSecretScanningAlerts(v string)`

SetSecretScanningAlerts sets SecretScanningAlerts field to given value.

### HasSecretScanningAlerts

`func (o *AppPermissions) HasSecretScanningAlerts() bool`

HasSecretScanningAlerts returns a boolean if a field has been set.

### GetSecrets

`func (o *AppPermissions) GetSecrets() string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *AppPermissions) GetSecretsOk() (*string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *AppPermissions) SetSecrets(v string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *AppPermissions) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetSecurityEvents

`func (o *AppPermissions) GetSecurityEvents() string`

GetSecurityEvents returns the SecurityEvents field if non-nil, zero value otherwise.

### GetSecurityEventsOk

`func (o *AppPermissions) GetSecurityEventsOk() (*string, bool)`

GetSecurityEventsOk returns a tuple with the SecurityEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEvents

`func (o *AppPermissions) SetSecurityEvents(v string)`

SetSecurityEvents sets SecurityEvents field to given value.

### HasSecurityEvents

`func (o *AppPermissions) HasSecurityEvents() bool`

HasSecurityEvents returns a boolean if a field has been set.

### GetSingleFile

`func (o *AppPermissions) GetSingleFile() string`

GetSingleFile returns the SingleFile field if non-nil, zero value otherwise.

### GetSingleFileOk

`func (o *AppPermissions) GetSingleFileOk() (*string, bool)`

GetSingleFileOk returns a tuple with the SingleFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleFile

`func (o *AppPermissions) SetSingleFile(v string)`

SetSingleFile sets SingleFile field to given value.

### HasSingleFile

`func (o *AppPermissions) HasSingleFile() bool`

HasSingleFile returns a boolean if a field has been set.

### GetStatuses

`func (o *AppPermissions) GetStatuses() string`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *AppPermissions) GetStatusesOk() (*string, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *AppPermissions) SetStatuses(v string)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *AppPermissions) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetVulnerabilityAlerts

`func (o *AppPermissions) GetVulnerabilityAlerts() string`

GetVulnerabilityAlerts returns the VulnerabilityAlerts field if non-nil, zero value otherwise.

### GetVulnerabilityAlertsOk

`func (o *AppPermissions) GetVulnerabilityAlertsOk() (*string, bool)`

GetVulnerabilityAlertsOk returns a tuple with the VulnerabilityAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityAlerts

`func (o *AppPermissions) SetVulnerabilityAlerts(v string)`

SetVulnerabilityAlerts sets VulnerabilityAlerts field to given value.

### HasVulnerabilityAlerts

`func (o *AppPermissions) HasVulnerabilityAlerts() bool`

HasVulnerabilityAlerts returns a boolean if a field has been set.

### GetWorkflows

`func (o *AppPermissions) GetWorkflows() string`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *AppPermissions) GetWorkflowsOk() (*string, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *AppPermissions) SetWorkflows(v string)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *AppPermissions) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.

### GetMembers

`func (o *AppPermissions) GetMembers() string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AppPermissions) GetMembersOk() (*string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AppPermissions) SetMembers(v string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *AppPermissions) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetOrganizationAdministration

`func (o *AppPermissions) GetOrganizationAdministration() string`

GetOrganizationAdministration returns the OrganizationAdministration field if non-nil, zero value otherwise.

### GetOrganizationAdministrationOk

`func (o *AppPermissions) GetOrganizationAdministrationOk() (*string, bool)`

GetOrganizationAdministrationOk returns a tuple with the OrganizationAdministration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationAdministration

`func (o *AppPermissions) SetOrganizationAdministration(v string)`

SetOrganizationAdministration sets OrganizationAdministration field to given value.

### HasOrganizationAdministration

`func (o *AppPermissions) HasOrganizationAdministration() bool`

HasOrganizationAdministration returns a boolean if a field has been set.

### GetOrganizationCustomRoles

`func (o *AppPermissions) GetOrganizationCustomRoles() string`

GetOrganizationCustomRoles returns the OrganizationCustomRoles field if non-nil, zero value otherwise.

### GetOrganizationCustomRolesOk

`func (o *AppPermissions) GetOrganizationCustomRolesOk() (*string, bool)`

GetOrganizationCustomRolesOk returns a tuple with the OrganizationCustomRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCustomRoles

`func (o *AppPermissions) SetOrganizationCustomRoles(v string)`

SetOrganizationCustomRoles sets OrganizationCustomRoles field to given value.

### HasOrganizationCustomRoles

`func (o *AppPermissions) HasOrganizationCustomRoles() bool`

HasOrganizationCustomRoles returns a boolean if a field has been set.

### GetOrganizationHooks

`func (o *AppPermissions) GetOrganizationHooks() string`

GetOrganizationHooks returns the OrganizationHooks field if non-nil, zero value otherwise.

### GetOrganizationHooksOk

`func (o *AppPermissions) GetOrganizationHooksOk() (*string, bool)`

GetOrganizationHooksOk returns a tuple with the OrganizationHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationHooks

`func (o *AppPermissions) SetOrganizationHooks(v string)`

SetOrganizationHooks sets OrganizationHooks field to given value.

### HasOrganizationHooks

`func (o *AppPermissions) HasOrganizationHooks() bool`

HasOrganizationHooks returns a boolean if a field has been set.

### GetOrganizationPlan

`func (o *AppPermissions) GetOrganizationPlan() string`

GetOrganizationPlan returns the OrganizationPlan field if non-nil, zero value otherwise.

### GetOrganizationPlanOk

`func (o *AppPermissions) GetOrganizationPlanOk() (*string, bool)`

GetOrganizationPlanOk returns a tuple with the OrganizationPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationPlan

`func (o *AppPermissions) SetOrganizationPlan(v string)`

SetOrganizationPlan sets OrganizationPlan field to given value.

### HasOrganizationPlan

`func (o *AppPermissions) HasOrganizationPlan() bool`

HasOrganizationPlan returns a boolean if a field has been set.

### GetOrganizationProjects

`func (o *AppPermissions) GetOrganizationProjects() string`

GetOrganizationProjects returns the OrganizationProjects field if non-nil, zero value otherwise.

### GetOrganizationProjectsOk

`func (o *AppPermissions) GetOrganizationProjectsOk() (*string, bool)`

GetOrganizationProjectsOk returns a tuple with the OrganizationProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProjects

`func (o *AppPermissions) SetOrganizationProjects(v string)`

SetOrganizationProjects sets OrganizationProjects field to given value.

### HasOrganizationProjects

`func (o *AppPermissions) HasOrganizationProjects() bool`

HasOrganizationProjects returns a boolean if a field has been set.

### GetOrganizationPackages

`func (o *AppPermissions) GetOrganizationPackages() string`

GetOrganizationPackages returns the OrganizationPackages field if non-nil, zero value otherwise.

### GetOrganizationPackagesOk

`func (o *AppPermissions) GetOrganizationPackagesOk() (*string, bool)`

GetOrganizationPackagesOk returns a tuple with the OrganizationPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationPackages

`func (o *AppPermissions) SetOrganizationPackages(v string)`

SetOrganizationPackages sets OrganizationPackages field to given value.

### HasOrganizationPackages

`func (o *AppPermissions) HasOrganizationPackages() bool`

HasOrganizationPackages returns a boolean if a field has been set.

### GetOrganizationSecrets

`func (o *AppPermissions) GetOrganizationSecrets() string`

GetOrganizationSecrets returns the OrganizationSecrets field if non-nil, zero value otherwise.

### GetOrganizationSecretsOk

`func (o *AppPermissions) GetOrganizationSecretsOk() (*string, bool)`

GetOrganizationSecretsOk returns a tuple with the OrganizationSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationSecrets

`func (o *AppPermissions) SetOrganizationSecrets(v string)`

SetOrganizationSecrets sets OrganizationSecrets field to given value.

### HasOrganizationSecrets

`func (o *AppPermissions) HasOrganizationSecrets() bool`

HasOrganizationSecrets returns a boolean if a field has been set.

### GetOrganizationSelfHostedRunners

`func (o *AppPermissions) GetOrganizationSelfHostedRunners() string`

GetOrganizationSelfHostedRunners returns the OrganizationSelfHostedRunners field if non-nil, zero value otherwise.

### GetOrganizationSelfHostedRunnersOk

`func (o *AppPermissions) GetOrganizationSelfHostedRunnersOk() (*string, bool)`

GetOrganizationSelfHostedRunnersOk returns a tuple with the OrganizationSelfHostedRunners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationSelfHostedRunners

`func (o *AppPermissions) SetOrganizationSelfHostedRunners(v string)`

SetOrganizationSelfHostedRunners sets OrganizationSelfHostedRunners field to given value.

### HasOrganizationSelfHostedRunners

`func (o *AppPermissions) HasOrganizationSelfHostedRunners() bool`

HasOrganizationSelfHostedRunners returns a boolean if a field has been set.

### GetOrganizationUserBlocking

`func (o *AppPermissions) GetOrganizationUserBlocking() string`

GetOrganizationUserBlocking returns the OrganizationUserBlocking field if non-nil, zero value otherwise.

### GetOrganizationUserBlockingOk

`func (o *AppPermissions) GetOrganizationUserBlockingOk() (*string, bool)`

GetOrganizationUserBlockingOk returns a tuple with the OrganizationUserBlocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUserBlocking

`func (o *AppPermissions) SetOrganizationUserBlocking(v string)`

SetOrganizationUserBlocking sets OrganizationUserBlocking field to given value.

### HasOrganizationUserBlocking

`func (o *AppPermissions) HasOrganizationUserBlocking() bool`

HasOrganizationUserBlocking returns a boolean if a field has been set.

### GetTeamDiscussions

`func (o *AppPermissions) GetTeamDiscussions() string`

GetTeamDiscussions returns the TeamDiscussions field if non-nil, zero value otherwise.

### GetTeamDiscussionsOk

`func (o *AppPermissions) GetTeamDiscussionsOk() (*string, bool)`

GetTeamDiscussionsOk returns a tuple with the TeamDiscussions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamDiscussions

`func (o *AppPermissions) SetTeamDiscussions(v string)`

SetTeamDiscussions sets TeamDiscussions field to given value.

### HasTeamDiscussions

`func (o *AppPermissions) HasTeamDiscussions() bool`

HasTeamDiscussions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


