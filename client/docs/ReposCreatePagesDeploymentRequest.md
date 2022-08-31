# ReposCreatePagesDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactUrl** | **string** | The URL of an artifact that contains the .zip or .tar of static assets to deploy. The artifact belongs to the repository. | 
**Environment** | Pointer to **string** | The target environment for this GitHub Pages deployment. | [optional] [default to "github-pages"]
**PagesBuildVersion** | **string** | A unique string that represents the version of the build for this deployment. | [default to "GITHUB_SHA"]
**OidcToken** | **string** | The OIDC token issued by GitHub Actions certifying the origin of the deployment. | 

## Methods

### NewReposCreatePagesDeploymentRequest

`func NewReposCreatePagesDeploymentRequest(artifactUrl string, pagesBuildVersion string, oidcToken string, ) *ReposCreatePagesDeploymentRequest`

NewReposCreatePagesDeploymentRequest instantiates a new ReposCreatePagesDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreatePagesDeploymentRequestWithDefaults

`func NewReposCreatePagesDeploymentRequestWithDefaults() *ReposCreatePagesDeploymentRequest`

NewReposCreatePagesDeploymentRequestWithDefaults instantiates a new ReposCreatePagesDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactUrl

`func (o *ReposCreatePagesDeploymentRequest) GetArtifactUrl() string`

GetArtifactUrl returns the ArtifactUrl field if non-nil, zero value otherwise.

### GetArtifactUrlOk

`func (o *ReposCreatePagesDeploymentRequest) GetArtifactUrlOk() (*string, bool)`

GetArtifactUrlOk returns a tuple with the ArtifactUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactUrl

`func (o *ReposCreatePagesDeploymentRequest) SetArtifactUrl(v string)`

SetArtifactUrl sets ArtifactUrl field to given value.


### GetEnvironment

`func (o *ReposCreatePagesDeploymentRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ReposCreatePagesDeploymentRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ReposCreatePagesDeploymentRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ReposCreatePagesDeploymentRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetPagesBuildVersion

`func (o *ReposCreatePagesDeploymentRequest) GetPagesBuildVersion() string`

GetPagesBuildVersion returns the PagesBuildVersion field if non-nil, zero value otherwise.

### GetPagesBuildVersionOk

`func (o *ReposCreatePagesDeploymentRequest) GetPagesBuildVersionOk() (*string, bool)`

GetPagesBuildVersionOk returns a tuple with the PagesBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesBuildVersion

`func (o *ReposCreatePagesDeploymentRequest) SetPagesBuildVersion(v string)`

SetPagesBuildVersion sets PagesBuildVersion field to given value.


### GetOidcToken

`func (o *ReposCreatePagesDeploymentRequest) GetOidcToken() string`

GetOidcToken returns the OidcToken field if non-nil, zero value otherwise.

### GetOidcTokenOk

`func (o *ReposCreatePagesDeploymentRequest) GetOidcTokenOk() (*string, bool)`

GetOidcTokenOk returns a tuple with the OidcToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcToken

`func (o *ReposCreatePagesDeploymentRequest) SetOidcToken(v string)`

SetOidcToken sets OidcToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


