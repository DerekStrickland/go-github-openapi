/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
)

// ReposCreateForAuthenticatedUserRequest struct for ReposCreateForAuthenticatedUserRequest
type ReposCreateForAuthenticatedUserRequest struct {
	// The name of the repository.
	Name string `json:"name"`
	// A short description of the repository.
	Description *string `json:"description,omitempty"`
	// A URL with more information about the repository.
	Homepage *string `json:"homepage,omitempty"`
	// Whether the repository is private.
	Private *bool `json:"private,omitempty"`
	// Whether issues are enabled.
	HasIssues *bool `json:"has_issues,omitempty"`
	// Whether projects are enabled.
	HasProjects *bool `json:"has_projects,omitempty"`
	// Whether the wiki is enabled.
	HasWiki *bool `json:"has_wiki,omitempty"`
	// The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization.
	TeamId *int32 `json:"team_id,omitempty"`
	// Whether the repository is initialized with a minimal README.
	AutoInit *bool `json:"auto_init,omitempty"`
	// The desired language or platform to apply to the .gitignore.
	GitignoreTemplate *string `json:"gitignore_template,omitempty"`
	// The license keyword of the open source license for this repository.
	LicenseTemplate *string `json:"license_template,omitempty"`
	// Whether to allow squash merges for pull requests.
	AllowSquashMerge *bool `json:"allow_squash_merge,omitempty"`
	// Whether to allow merge commits for pull requests.
	AllowMergeCommit *bool `json:"allow_merge_commit,omitempty"`
	// Whether to allow rebase merges for pull requests.
	AllowRebaseMerge *bool `json:"allow_rebase_merge,omitempty"`
	// Whether to allow Auto-merge to be used on pull requests.
	AllowAutoMerge *bool `json:"allow_auto_merge,omitempty"`
	// Whether to delete head branches when pull requests are merged
	DeleteBranchOnMerge *bool `json:"delete_branch_on_merge,omitempty"`
	// The default value for a squash merge commit title:  - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	SquashMergeCommitTitle *string `json:"squash_merge_commit_title,omitempty"`
	// The default value for a squash merge commit message:  - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	SquashMergeCommitMessage *string `json:"squash_merge_commit_message,omitempty"`
	// The default value for a merge commit title.  - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	MergeCommitTitle *string `json:"merge_commit_title,omitempty"`
	// The default value for a merge commit message.  - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	MergeCommitMessage *string `json:"merge_commit_message,omitempty"`
	// Whether downloads are enabled.
	HasDownloads *bool `json:"has_downloads,omitempty"`
	// Whether this repository acts as a template that can be used to generate new repositories.
	IsTemplate *bool `json:"is_template,omitempty"`
}

// NewReposCreateForAuthenticatedUserRequest instantiates a new ReposCreateForAuthenticatedUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposCreateForAuthenticatedUserRequest(name string) *ReposCreateForAuthenticatedUserRequest {
	this := ReposCreateForAuthenticatedUserRequest{}
	this.Name = name
	var private bool = false
	this.Private = &private
	var hasIssues bool = true
	this.HasIssues = &hasIssues
	var hasProjects bool = true
	this.HasProjects = &hasProjects
	var hasWiki bool = true
	this.HasWiki = &hasWiki
	var autoInit bool = false
	this.AutoInit = &autoInit
	var allowSquashMerge bool = true
	this.AllowSquashMerge = &allowSquashMerge
	var allowMergeCommit bool = true
	this.AllowMergeCommit = &allowMergeCommit
	var allowRebaseMerge bool = true
	this.AllowRebaseMerge = &allowRebaseMerge
	var allowAutoMerge bool = false
	this.AllowAutoMerge = &allowAutoMerge
	var deleteBranchOnMerge bool = false
	this.DeleteBranchOnMerge = &deleteBranchOnMerge
	var hasDownloads bool = true
	this.HasDownloads = &hasDownloads
	var isTemplate bool = false
	this.IsTemplate = &isTemplate
	return &this
}

// NewReposCreateForAuthenticatedUserRequestWithDefaults instantiates a new ReposCreateForAuthenticatedUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposCreateForAuthenticatedUserRequestWithDefaults() *ReposCreateForAuthenticatedUserRequest {
	this := ReposCreateForAuthenticatedUserRequest{}
	var private bool = false
	this.Private = &private
	var hasIssues bool = true
	this.HasIssues = &hasIssues
	var hasProjects bool = true
	this.HasProjects = &hasProjects
	var hasWiki bool = true
	this.HasWiki = &hasWiki
	var autoInit bool = false
	this.AutoInit = &autoInit
	var allowSquashMerge bool = true
	this.AllowSquashMerge = &allowSquashMerge
	var allowMergeCommit bool = true
	this.AllowMergeCommit = &allowMergeCommit
	var allowRebaseMerge bool = true
	this.AllowRebaseMerge = &allowRebaseMerge
	var allowAutoMerge bool = false
	this.AllowAutoMerge = &allowAutoMerge
	var deleteBranchOnMerge bool = false
	this.DeleteBranchOnMerge = &deleteBranchOnMerge
	var hasDownloads bool = true
	this.HasDownloads = &hasDownloads
	var isTemplate bool = false
	this.IsTemplate = &isTemplate
	return &this
}

// GetName returns the Name field value
func (o *ReposCreateForAuthenticatedUserRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReposCreateForAuthenticatedUserRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReposCreateForAuthenticatedUserRequest) SetDescription(v string) {
	o.Description = &v
}

// GetHomepage returns the Homepage field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetHomepage() string {
	if o == nil || o.Homepage == nil {
		var ret string
		return ret
	}
	return *o.Homepage
}

// GetHomepageOk returns a tuple with the Homepage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetHomepageOk() (*string, bool) {
	if o == nil || o.Homepage == nil {
		return nil, false
	}
	return o.Homepage, true
}

// HasHomepage returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasHomepage() bool {
	if o != nil && o.Homepage != nil {
		return true
	}

	return false
}

// SetHomepage gets a reference to the given string and assigns it to the Homepage field.
func (o *ReposCreateForAuthenticatedUserRequest) SetHomepage(v string) {
	o.Homepage = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ReposCreateForAuthenticatedUserRequest) SetPrivate(v bool) {
	o.Private = &v
}

// GetHasIssues returns the HasIssues field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetHasIssues() bool {
	if o == nil || o.HasIssues == nil {
		var ret bool
		return ret
	}
	return *o.HasIssues
}

// GetHasIssuesOk returns a tuple with the HasIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetHasIssuesOk() (*bool, bool) {
	if o == nil || o.HasIssues == nil {
		return nil, false
	}
	return o.HasIssues, true
}

// HasHasIssues returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasHasIssues() bool {
	if o != nil && o.HasIssues != nil {
		return true
	}

	return false
}

// SetHasIssues gets a reference to the given bool and assigns it to the HasIssues field.
func (o *ReposCreateForAuthenticatedUserRequest) SetHasIssues(v bool) {
	o.HasIssues = &v
}

// GetHasProjects returns the HasProjects field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetHasProjects() bool {
	if o == nil || o.HasProjects == nil {
		var ret bool
		return ret
	}
	return *o.HasProjects
}

// GetHasProjectsOk returns a tuple with the HasProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetHasProjectsOk() (*bool, bool) {
	if o == nil || o.HasProjects == nil {
		return nil, false
	}
	return o.HasProjects, true
}

// HasHasProjects returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasHasProjects() bool {
	if o != nil && o.HasProjects != nil {
		return true
	}

	return false
}

// SetHasProjects gets a reference to the given bool and assigns it to the HasProjects field.
func (o *ReposCreateForAuthenticatedUserRequest) SetHasProjects(v bool) {
	o.HasProjects = &v
}

// GetHasWiki returns the HasWiki field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetHasWiki() bool {
	if o == nil || o.HasWiki == nil {
		var ret bool
		return ret
	}
	return *o.HasWiki
}

// GetHasWikiOk returns a tuple with the HasWiki field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetHasWikiOk() (*bool, bool) {
	if o == nil || o.HasWiki == nil {
		return nil, false
	}
	return o.HasWiki, true
}

// HasHasWiki returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasHasWiki() bool {
	if o != nil && o.HasWiki != nil {
		return true
	}

	return false
}

// SetHasWiki gets a reference to the given bool and assigns it to the HasWiki field.
func (o *ReposCreateForAuthenticatedUserRequest) SetHasWiki(v bool) {
	o.HasWiki = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetTeamId() int32 {
	if o == nil || o.TeamId == nil {
		var ret int32
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetTeamIdOk() (*int32, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasTeamId() bool {
	if o != nil && o.TeamId != nil {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given int32 and assigns it to the TeamId field.
func (o *ReposCreateForAuthenticatedUserRequest) SetTeamId(v int32) {
	o.TeamId = &v
}

// GetAutoInit returns the AutoInit field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetAutoInit() bool {
	if o == nil || o.AutoInit == nil {
		var ret bool
		return ret
	}
	return *o.AutoInit
}

// GetAutoInitOk returns a tuple with the AutoInit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetAutoInitOk() (*bool, bool) {
	if o == nil || o.AutoInit == nil {
		return nil, false
	}
	return o.AutoInit, true
}

// HasAutoInit returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasAutoInit() bool {
	if o != nil && o.AutoInit != nil {
		return true
	}

	return false
}

// SetAutoInit gets a reference to the given bool and assigns it to the AutoInit field.
func (o *ReposCreateForAuthenticatedUserRequest) SetAutoInit(v bool) {
	o.AutoInit = &v
}

// GetGitignoreTemplate returns the GitignoreTemplate field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetGitignoreTemplate() string {
	if o == nil || o.GitignoreTemplate == nil {
		var ret string
		return ret
	}
	return *o.GitignoreTemplate
}

// GetGitignoreTemplateOk returns a tuple with the GitignoreTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetGitignoreTemplateOk() (*string, bool) {
	if o == nil || o.GitignoreTemplate == nil {
		return nil, false
	}
	return o.GitignoreTemplate, true
}

// HasGitignoreTemplate returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasGitignoreTemplate() bool {
	if o != nil && o.GitignoreTemplate != nil {
		return true
	}

	return false
}

// SetGitignoreTemplate gets a reference to the given string and assigns it to the GitignoreTemplate field.
func (o *ReposCreateForAuthenticatedUserRequest) SetGitignoreTemplate(v string) {
	o.GitignoreTemplate = &v
}

// GetLicenseTemplate returns the LicenseTemplate field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetLicenseTemplate() string {
	if o == nil || o.LicenseTemplate == nil {
		var ret string
		return ret
	}
	return *o.LicenseTemplate
}

// GetLicenseTemplateOk returns a tuple with the LicenseTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetLicenseTemplateOk() (*string, bool) {
	if o == nil || o.LicenseTemplate == nil {
		return nil, false
	}
	return o.LicenseTemplate, true
}

// HasLicenseTemplate returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasLicenseTemplate() bool {
	if o != nil && o.LicenseTemplate != nil {
		return true
	}

	return false
}

// SetLicenseTemplate gets a reference to the given string and assigns it to the LicenseTemplate field.
func (o *ReposCreateForAuthenticatedUserRequest) SetLicenseTemplate(v string) {
	o.LicenseTemplate = &v
}

// GetAllowSquashMerge returns the AllowSquashMerge field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetAllowSquashMerge() bool {
	if o == nil || o.AllowSquashMerge == nil {
		var ret bool
		return ret
	}
	return *o.AllowSquashMerge
}

// GetAllowSquashMergeOk returns a tuple with the AllowSquashMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetAllowSquashMergeOk() (*bool, bool) {
	if o == nil || o.AllowSquashMerge == nil {
		return nil, false
	}
	return o.AllowSquashMerge, true
}

// HasAllowSquashMerge returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasAllowSquashMerge() bool {
	if o != nil && o.AllowSquashMerge != nil {
		return true
	}

	return false
}

// SetAllowSquashMerge gets a reference to the given bool and assigns it to the AllowSquashMerge field.
func (o *ReposCreateForAuthenticatedUserRequest) SetAllowSquashMerge(v bool) {
	o.AllowSquashMerge = &v
}

// GetAllowMergeCommit returns the AllowMergeCommit field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetAllowMergeCommit() bool {
	if o == nil || o.AllowMergeCommit == nil {
		var ret bool
		return ret
	}
	return *o.AllowMergeCommit
}

// GetAllowMergeCommitOk returns a tuple with the AllowMergeCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetAllowMergeCommitOk() (*bool, bool) {
	if o == nil || o.AllowMergeCommit == nil {
		return nil, false
	}
	return o.AllowMergeCommit, true
}

// HasAllowMergeCommit returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasAllowMergeCommit() bool {
	if o != nil && o.AllowMergeCommit != nil {
		return true
	}

	return false
}

// SetAllowMergeCommit gets a reference to the given bool and assigns it to the AllowMergeCommit field.
func (o *ReposCreateForAuthenticatedUserRequest) SetAllowMergeCommit(v bool) {
	o.AllowMergeCommit = &v
}

// GetAllowRebaseMerge returns the AllowRebaseMerge field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetAllowRebaseMerge() bool {
	if o == nil || o.AllowRebaseMerge == nil {
		var ret bool
		return ret
	}
	return *o.AllowRebaseMerge
}

// GetAllowRebaseMergeOk returns a tuple with the AllowRebaseMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetAllowRebaseMergeOk() (*bool, bool) {
	if o == nil || o.AllowRebaseMerge == nil {
		return nil, false
	}
	return o.AllowRebaseMerge, true
}

// HasAllowRebaseMerge returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasAllowRebaseMerge() bool {
	if o != nil && o.AllowRebaseMerge != nil {
		return true
	}

	return false
}

// SetAllowRebaseMerge gets a reference to the given bool and assigns it to the AllowRebaseMerge field.
func (o *ReposCreateForAuthenticatedUserRequest) SetAllowRebaseMerge(v bool) {
	o.AllowRebaseMerge = &v
}

// GetAllowAutoMerge returns the AllowAutoMerge field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetAllowAutoMerge() bool {
	if o == nil || o.AllowAutoMerge == nil {
		var ret bool
		return ret
	}
	return *o.AllowAutoMerge
}

// GetAllowAutoMergeOk returns a tuple with the AllowAutoMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetAllowAutoMergeOk() (*bool, bool) {
	if o == nil || o.AllowAutoMerge == nil {
		return nil, false
	}
	return o.AllowAutoMerge, true
}

// HasAllowAutoMerge returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasAllowAutoMerge() bool {
	if o != nil && o.AllowAutoMerge != nil {
		return true
	}

	return false
}

// SetAllowAutoMerge gets a reference to the given bool and assigns it to the AllowAutoMerge field.
func (o *ReposCreateForAuthenticatedUserRequest) SetAllowAutoMerge(v bool) {
	o.AllowAutoMerge = &v
}

// GetDeleteBranchOnMerge returns the DeleteBranchOnMerge field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetDeleteBranchOnMerge() bool {
	if o == nil || o.DeleteBranchOnMerge == nil {
		var ret bool
		return ret
	}
	return *o.DeleteBranchOnMerge
}

// GetDeleteBranchOnMergeOk returns a tuple with the DeleteBranchOnMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetDeleteBranchOnMergeOk() (*bool, bool) {
	if o == nil || o.DeleteBranchOnMerge == nil {
		return nil, false
	}
	return o.DeleteBranchOnMerge, true
}

// HasDeleteBranchOnMerge returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasDeleteBranchOnMerge() bool {
	if o != nil && o.DeleteBranchOnMerge != nil {
		return true
	}

	return false
}

// SetDeleteBranchOnMerge gets a reference to the given bool and assigns it to the DeleteBranchOnMerge field.
func (o *ReposCreateForAuthenticatedUserRequest) SetDeleteBranchOnMerge(v bool) {
	o.DeleteBranchOnMerge = &v
}

// GetSquashMergeCommitTitle returns the SquashMergeCommitTitle field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetSquashMergeCommitTitle() string {
	if o == nil || o.SquashMergeCommitTitle == nil {
		var ret string
		return ret
	}
	return *o.SquashMergeCommitTitle
}

// GetSquashMergeCommitTitleOk returns a tuple with the SquashMergeCommitTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetSquashMergeCommitTitleOk() (*string, bool) {
	if o == nil || o.SquashMergeCommitTitle == nil {
		return nil, false
	}
	return o.SquashMergeCommitTitle, true
}

// HasSquashMergeCommitTitle returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasSquashMergeCommitTitle() bool {
	if o != nil && o.SquashMergeCommitTitle != nil {
		return true
	}

	return false
}

// SetSquashMergeCommitTitle gets a reference to the given string and assigns it to the SquashMergeCommitTitle field.
func (o *ReposCreateForAuthenticatedUserRequest) SetSquashMergeCommitTitle(v string) {
	o.SquashMergeCommitTitle = &v
}

// GetSquashMergeCommitMessage returns the SquashMergeCommitMessage field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetSquashMergeCommitMessage() string {
	if o == nil || o.SquashMergeCommitMessage == nil {
		var ret string
		return ret
	}
	return *o.SquashMergeCommitMessage
}

// GetSquashMergeCommitMessageOk returns a tuple with the SquashMergeCommitMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetSquashMergeCommitMessageOk() (*string, bool) {
	if o == nil || o.SquashMergeCommitMessage == nil {
		return nil, false
	}
	return o.SquashMergeCommitMessage, true
}

// HasSquashMergeCommitMessage returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasSquashMergeCommitMessage() bool {
	if o != nil && o.SquashMergeCommitMessage != nil {
		return true
	}

	return false
}

// SetSquashMergeCommitMessage gets a reference to the given string and assigns it to the SquashMergeCommitMessage field.
func (o *ReposCreateForAuthenticatedUserRequest) SetSquashMergeCommitMessage(v string) {
	o.SquashMergeCommitMessage = &v
}

// GetMergeCommitTitle returns the MergeCommitTitle field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetMergeCommitTitle() string {
	if o == nil || o.MergeCommitTitle == nil {
		var ret string
		return ret
	}
	return *o.MergeCommitTitle
}

// GetMergeCommitTitleOk returns a tuple with the MergeCommitTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetMergeCommitTitleOk() (*string, bool) {
	if o == nil || o.MergeCommitTitle == nil {
		return nil, false
	}
	return o.MergeCommitTitle, true
}

// HasMergeCommitTitle returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasMergeCommitTitle() bool {
	if o != nil && o.MergeCommitTitle != nil {
		return true
	}

	return false
}

// SetMergeCommitTitle gets a reference to the given string and assigns it to the MergeCommitTitle field.
func (o *ReposCreateForAuthenticatedUserRequest) SetMergeCommitTitle(v string) {
	o.MergeCommitTitle = &v
}

// GetMergeCommitMessage returns the MergeCommitMessage field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetMergeCommitMessage() string {
	if o == nil || o.MergeCommitMessage == nil {
		var ret string
		return ret
	}
	return *o.MergeCommitMessage
}

// GetMergeCommitMessageOk returns a tuple with the MergeCommitMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetMergeCommitMessageOk() (*string, bool) {
	if o == nil || o.MergeCommitMessage == nil {
		return nil, false
	}
	return o.MergeCommitMessage, true
}

// HasMergeCommitMessage returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasMergeCommitMessage() bool {
	if o != nil && o.MergeCommitMessage != nil {
		return true
	}

	return false
}

// SetMergeCommitMessage gets a reference to the given string and assigns it to the MergeCommitMessage field.
func (o *ReposCreateForAuthenticatedUserRequest) SetMergeCommitMessage(v string) {
	o.MergeCommitMessage = &v
}

// GetHasDownloads returns the HasDownloads field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetHasDownloads() bool {
	if o == nil || o.HasDownloads == nil {
		var ret bool
		return ret
	}
	return *o.HasDownloads
}

// GetHasDownloadsOk returns a tuple with the HasDownloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetHasDownloadsOk() (*bool, bool) {
	if o == nil || o.HasDownloads == nil {
		return nil, false
	}
	return o.HasDownloads, true
}

// HasHasDownloads returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasHasDownloads() bool {
	if o != nil && o.HasDownloads != nil {
		return true
	}

	return false
}

// SetHasDownloads gets a reference to the given bool and assigns it to the HasDownloads field.
func (o *ReposCreateForAuthenticatedUserRequest) SetHasDownloads(v bool) {
	o.HasDownloads = &v
}

// GetIsTemplate returns the IsTemplate field value if set, zero value otherwise.
func (o *ReposCreateForAuthenticatedUserRequest) GetIsTemplate() bool {
	if o == nil || o.IsTemplate == nil {
		var ret bool
		return ret
	}
	return *o.IsTemplate
}

// GetIsTemplateOk returns a tuple with the IsTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateForAuthenticatedUserRequest) GetIsTemplateOk() (*bool, bool) {
	if o == nil || o.IsTemplate == nil {
		return nil, false
	}
	return o.IsTemplate, true
}

// HasIsTemplate returns a boolean if a field has been set.
func (o *ReposCreateForAuthenticatedUserRequest) HasIsTemplate() bool {
	if o != nil && o.IsTemplate != nil {
		return true
	}

	return false
}

// SetIsTemplate gets a reference to the given bool and assigns it to the IsTemplate field.
func (o *ReposCreateForAuthenticatedUserRequest) SetIsTemplate(v bool) {
	o.IsTemplate = &v
}

func (o ReposCreateForAuthenticatedUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Homepage != nil {
		toSerialize["homepage"] = o.Homepage
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.HasIssues != nil {
		toSerialize["has_issues"] = o.HasIssues
	}
	if o.HasProjects != nil {
		toSerialize["has_projects"] = o.HasProjects
	}
	if o.HasWiki != nil {
		toSerialize["has_wiki"] = o.HasWiki
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.AutoInit != nil {
		toSerialize["auto_init"] = o.AutoInit
	}
	if o.GitignoreTemplate != nil {
		toSerialize["gitignore_template"] = o.GitignoreTemplate
	}
	if o.LicenseTemplate != nil {
		toSerialize["license_template"] = o.LicenseTemplate
	}
	if o.AllowSquashMerge != nil {
		toSerialize["allow_squash_merge"] = o.AllowSquashMerge
	}
	if o.AllowMergeCommit != nil {
		toSerialize["allow_merge_commit"] = o.AllowMergeCommit
	}
	if o.AllowRebaseMerge != nil {
		toSerialize["allow_rebase_merge"] = o.AllowRebaseMerge
	}
	if o.AllowAutoMerge != nil {
		toSerialize["allow_auto_merge"] = o.AllowAutoMerge
	}
	if o.DeleteBranchOnMerge != nil {
		toSerialize["delete_branch_on_merge"] = o.DeleteBranchOnMerge
	}
	if o.SquashMergeCommitTitle != nil {
		toSerialize["squash_merge_commit_title"] = o.SquashMergeCommitTitle
	}
	if o.SquashMergeCommitMessage != nil {
		toSerialize["squash_merge_commit_message"] = o.SquashMergeCommitMessage
	}
	if o.MergeCommitTitle != nil {
		toSerialize["merge_commit_title"] = o.MergeCommitTitle
	}
	if o.MergeCommitMessage != nil {
		toSerialize["merge_commit_message"] = o.MergeCommitMessage
	}
	if o.HasDownloads != nil {
		toSerialize["has_downloads"] = o.HasDownloads
	}
	if o.IsTemplate != nil {
		toSerialize["is_template"] = o.IsTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableReposCreateForAuthenticatedUserRequest struct {
	value *ReposCreateForAuthenticatedUserRequest
	isSet bool
}

func (v NullableReposCreateForAuthenticatedUserRequest) Get() *ReposCreateForAuthenticatedUserRequest {
	return v.value
}

func (v *NullableReposCreateForAuthenticatedUserRequest) Set(val *ReposCreateForAuthenticatedUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReposCreateForAuthenticatedUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReposCreateForAuthenticatedUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposCreateForAuthenticatedUserRequest(val *ReposCreateForAuthenticatedUserRequest) *NullableReposCreateForAuthenticatedUserRequest {
	return &NullableReposCreateForAuthenticatedUserRequest{value: val, isSet: true}
}

func (v NullableReposCreateForAuthenticatedUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposCreateForAuthenticatedUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


