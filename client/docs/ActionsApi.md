# \ActionsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionsAddCustomLabelsToSelfHostedRunnerForOrg**](ActionsApi.md#ActionsAddCustomLabelsToSelfHostedRunnerForOrg) | **Post** /orgs/{org}/actions/runners/{runner_id}/labels | Add custom labels to a self-hosted runner for an organization
[**ActionsAddCustomLabelsToSelfHostedRunnerForRepo**](ActionsApi.md#ActionsAddCustomLabelsToSelfHostedRunnerForRepo) | **Post** /repos/{owner}/{repo}/actions/runners/{runner_id}/labels | Add custom labels to a self-hosted runner for a repository
[**ActionsAddRepoAccessToSelfHostedRunnerGroupInOrg**](ActionsApi.md#ActionsAddRepoAccessToSelfHostedRunnerGroupInOrg) | **Put** /orgs/{org}/actions/runner-groups/{runner_group_id}/repositories/{repository_id} | Add repository access to a self-hosted runner group in an organization
[**ActionsAddSelectedRepoToOrgSecret**](ActionsApi.md#ActionsAddSelectedRepoToOrgSecret) | **Put** /orgs/{org}/actions/secrets/{secret_name}/repositories/{repository_id} | Add selected repository to an organization secret
[**ActionsAddSelfHostedRunnerToGroupForOrg**](ActionsApi.md#ActionsAddSelfHostedRunnerToGroupForOrg) | **Put** /orgs/{org}/actions/runner-groups/{runner_group_id}/runners/{runner_id} | Add a self-hosted runner to a group for an organization
[**ActionsApproveWorkflowRun**](ActionsApi.md#ActionsApproveWorkflowRun) | **Post** /repos/{owner}/{repo}/actions/runs/{run_id}/approve | Approve a workflow run for a fork pull request
[**ActionsCancelWorkflowRun**](ActionsApi.md#ActionsCancelWorkflowRun) | **Post** /repos/{owner}/{repo}/actions/runs/{run_id}/cancel | Cancel a workflow run
[**ActionsCreateOrUpdateEnvironmentSecret**](ActionsApi.md#ActionsCreateOrUpdateEnvironmentSecret) | **Put** /repositories/{repository_id}/environments/{environment_name}/secrets/{secret_name} | Create or update an environment secret
[**ActionsCreateOrUpdateOrgSecret**](ActionsApi.md#ActionsCreateOrUpdateOrgSecret) | **Put** /orgs/{org}/actions/secrets/{secret_name} | Create or update an organization secret
[**ActionsCreateOrUpdateRepoSecret**](ActionsApi.md#ActionsCreateOrUpdateRepoSecret) | **Put** /repos/{owner}/{repo}/actions/secrets/{secret_name} | Create or update a repository secret
[**ActionsCreateRegistrationTokenForOrg**](ActionsApi.md#ActionsCreateRegistrationTokenForOrg) | **Post** /orgs/{org}/actions/runners/registration-token | Create a registration token for an organization
[**ActionsCreateRegistrationTokenForRepo**](ActionsApi.md#ActionsCreateRegistrationTokenForRepo) | **Post** /repos/{owner}/{repo}/actions/runners/registration-token | Create a registration token for a repository
[**ActionsCreateRemoveTokenForOrg**](ActionsApi.md#ActionsCreateRemoveTokenForOrg) | **Post** /orgs/{org}/actions/runners/remove-token | Create a remove token for an organization
[**ActionsCreateRemoveTokenForRepo**](ActionsApi.md#ActionsCreateRemoveTokenForRepo) | **Post** /repos/{owner}/{repo}/actions/runners/remove-token | Create a remove token for a repository
[**ActionsCreateSelfHostedRunnerGroupForOrg**](ActionsApi.md#ActionsCreateSelfHostedRunnerGroupForOrg) | **Post** /orgs/{org}/actions/runner-groups | Create a self-hosted runner group for an organization
[**ActionsCreateWorkflowDispatch**](ActionsApi.md#ActionsCreateWorkflowDispatch) | **Post** /repos/{owner}/{repo}/actions/workflows/{workflow_id}/dispatches | Create a workflow dispatch event
[**ActionsDeleteActionsCacheById**](ActionsApi.md#ActionsDeleteActionsCacheById) | **Delete** /repos/{owner}/{repo}/actions/caches/{cache_id} | Delete a GitHub Actions cache for a repository (using a cache ID)
[**ActionsDeleteActionsCacheByKey**](ActionsApi.md#ActionsDeleteActionsCacheByKey) | **Delete** /repos/{owner}/{repo}/actions/caches | Delete GitHub Actions caches for a repository (using a cache key)
[**ActionsDeleteArtifact**](ActionsApi.md#ActionsDeleteArtifact) | **Delete** /repos/{owner}/{repo}/actions/artifacts/{artifact_id} | Delete an artifact
[**ActionsDeleteEnvironmentSecret**](ActionsApi.md#ActionsDeleteEnvironmentSecret) | **Delete** /repositories/{repository_id}/environments/{environment_name}/secrets/{secret_name} | Delete an environment secret
[**ActionsDeleteOrgSecret**](ActionsApi.md#ActionsDeleteOrgSecret) | **Delete** /orgs/{org}/actions/secrets/{secret_name} | Delete an organization secret
[**ActionsDeleteRepoSecret**](ActionsApi.md#ActionsDeleteRepoSecret) | **Delete** /repos/{owner}/{repo}/actions/secrets/{secret_name} | Delete a repository secret
[**ActionsDeleteSelfHostedRunnerFromOrg**](ActionsApi.md#ActionsDeleteSelfHostedRunnerFromOrg) | **Delete** /orgs/{org}/actions/runners/{runner_id} | Delete a self-hosted runner from an organization
[**ActionsDeleteSelfHostedRunnerFromRepo**](ActionsApi.md#ActionsDeleteSelfHostedRunnerFromRepo) | **Delete** /repos/{owner}/{repo}/actions/runners/{runner_id} | Delete a self-hosted runner from a repository
[**ActionsDeleteSelfHostedRunnerGroupFromOrg**](ActionsApi.md#ActionsDeleteSelfHostedRunnerGroupFromOrg) | **Delete** /orgs/{org}/actions/runner-groups/{runner_group_id} | Delete a self-hosted runner group from an organization
[**ActionsDeleteWorkflowRun**](ActionsApi.md#ActionsDeleteWorkflowRun) | **Delete** /repos/{owner}/{repo}/actions/runs/{run_id} | Delete a workflow run
[**ActionsDeleteWorkflowRunLogs**](ActionsApi.md#ActionsDeleteWorkflowRunLogs) | **Delete** /repos/{owner}/{repo}/actions/runs/{run_id}/logs | Delete workflow run logs
[**ActionsDisableSelectedRepositoryGithubActionsOrganization**](ActionsApi.md#ActionsDisableSelectedRepositoryGithubActionsOrganization) | **Delete** /orgs/{org}/actions/permissions/repositories/{repository_id} | Disable a selected repository for GitHub Actions in an organization
[**ActionsDisableWorkflow**](ActionsApi.md#ActionsDisableWorkflow) | **Put** /repos/{owner}/{repo}/actions/workflows/{workflow_id}/disable | Disable a workflow
[**ActionsDownloadArtifact**](ActionsApi.md#ActionsDownloadArtifact) | **Get** /repos/{owner}/{repo}/actions/artifacts/{artifact_id}/{archive_format} | Download an artifact
[**ActionsDownloadJobLogsForWorkflowRun**](ActionsApi.md#ActionsDownloadJobLogsForWorkflowRun) | **Get** /repos/{owner}/{repo}/actions/jobs/{job_id}/logs | Download job logs for a workflow run
[**ActionsDownloadWorkflowRunAttemptLogs**](ActionsApi.md#ActionsDownloadWorkflowRunAttemptLogs) | **Get** /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/logs | Download workflow run attempt logs
[**ActionsDownloadWorkflowRunLogs**](ActionsApi.md#ActionsDownloadWorkflowRunLogs) | **Get** /repos/{owner}/{repo}/actions/runs/{run_id}/logs | Download workflow run logs
[**ActionsEnableSelectedRepositoryGithubActionsOrganization**](ActionsApi.md#ActionsEnableSelectedRepositoryGithubActionsOrganization) | **Put** /orgs/{org}/actions/permissions/repositories/{repository_id} | Enable a selected repository for GitHub Actions in an organization
[**ActionsEnableWorkflow**](ActionsApi.md#ActionsEnableWorkflow) | **Put** /repos/{owner}/{repo}/actions/workflows/{workflow_id}/enable | Enable a workflow
[**ActionsGetActionsCacheList**](ActionsApi.md#ActionsGetActionsCacheList) | **Get** /repos/{owner}/{repo}/actions/caches | List GitHub Actions caches for a repository
[**ActionsGetActionsCacheUsage**](ActionsApi.md#ActionsGetActionsCacheUsage) | **Get** /repos/{owner}/{repo}/actions/cache/usage | Get GitHub Actions cache usage for a repository
[**ActionsGetActionsCacheUsageByRepoForOrg**](ActionsApi.md#ActionsGetActionsCacheUsageByRepoForOrg) | **Get** /orgs/{org}/actions/cache/usage-by-repository | List repositories with GitHub Actions cache usage for an organization
[**ActionsGetActionsCacheUsageForEnterprise**](ActionsApi.md#ActionsGetActionsCacheUsageForEnterprise) | **Get** /enterprises/{enterprise}/actions/cache/usage | Get GitHub Actions cache usage for an enterprise
[**ActionsGetActionsCacheUsageForOrg**](ActionsApi.md#ActionsGetActionsCacheUsageForOrg) | **Get** /orgs/{org}/actions/cache/usage | Get GitHub Actions cache usage for an organization
[**ActionsGetAllowedActionsOrganization**](ActionsApi.md#ActionsGetAllowedActionsOrganization) | **Get** /orgs/{org}/actions/permissions/selected-actions | Get allowed actions and reusable workflows for an organization
[**ActionsGetAllowedActionsRepository**](ActionsApi.md#ActionsGetAllowedActionsRepository) | **Get** /repos/{owner}/{repo}/actions/permissions/selected-actions | Get allowed actions and reusable workflows for a repository
[**ActionsGetArtifact**](ActionsApi.md#ActionsGetArtifact) | **Get** /repos/{owner}/{repo}/actions/artifacts/{artifact_id} | Get an artifact
[**ActionsGetCustomOidcSubClaimForRepo**](ActionsApi.md#ActionsGetCustomOidcSubClaimForRepo) | **Get** /repos/{owner}/{repo}/actions/oidc/customization/sub | Get the opt-out flag of an OIDC subject claim customization for a repository
[**ActionsGetEnvironmentPublicKey**](ActionsApi.md#ActionsGetEnvironmentPublicKey) | **Get** /repositories/{repository_id}/environments/{environment_name}/secrets/public-key | Get an environment public key
[**ActionsGetEnvironmentSecret**](ActionsApi.md#ActionsGetEnvironmentSecret) | **Get** /repositories/{repository_id}/environments/{environment_name}/secrets/{secret_name} | Get an environment secret
[**ActionsGetGithubActionsDefaultWorkflowPermissionsEnterprise**](ActionsApi.md#ActionsGetGithubActionsDefaultWorkflowPermissionsEnterprise) | **Get** /enterprises/{enterprise}/actions/permissions/workflow | Get default workflow permissions for an enterprise
[**ActionsGetGithubActionsDefaultWorkflowPermissionsOrganization**](ActionsApi.md#ActionsGetGithubActionsDefaultWorkflowPermissionsOrganization) | **Get** /orgs/{org}/actions/permissions/workflow | Get default workflow permissions for an organization
[**ActionsGetGithubActionsDefaultWorkflowPermissionsRepository**](ActionsApi.md#ActionsGetGithubActionsDefaultWorkflowPermissionsRepository) | **Get** /repos/{owner}/{repo}/actions/permissions/workflow | Get default workflow permissions for a repository
[**ActionsGetGithubActionsPermissionsOrganization**](ActionsApi.md#ActionsGetGithubActionsPermissionsOrganization) | **Get** /orgs/{org}/actions/permissions | Get GitHub Actions permissions for an organization
[**ActionsGetGithubActionsPermissionsRepository**](ActionsApi.md#ActionsGetGithubActionsPermissionsRepository) | **Get** /repos/{owner}/{repo}/actions/permissions | Get GitHub Actions permissions for a repository
[**ActionsGetJobForWorkflowRun**](ActionsApi.md#ActionsGetJobForWorkflowRun) | **Get** /repos/{owner}/{repo}/actions/jobs/{job_id} | Get a job for a workflow run
[**ActionsGetOrgPublicKey**](ActionsApi.md#ActionsGetOrgPublicKey) | **Get** /orgs/{org}/actions/secrets/public-key | Get an organization public key
[**ActionsGetOrgSecret**](ActionsApi.md#ActionsGetOrgSecret) | **Get** /orgs/{org}/actions/secrets/{secret_name} | Get an organization secret
[**ActionsGetPendingDeploymentsForRun**](ActionsApi.md#ActionsGetPendingDeploymentsForRun) | **Get** /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments | Get pending deployments for a workflow run
[**ActionsGetRepoPublicKey**](ActionsApi.md#ActionsGetRepoPublicKey) | **Get** /repos/{owner}/{repo}/actions/secrets/public-key | Get a repository public key
[**ActionsGetRepoSecret**](ActionsApi.md#ActionsGetRepoSecret) | **Get** /repos/{owner}/{repo}/actions/secrets/{secret_name} | Get a repository secret
[**ActionsGetReviewsForRun**](ActionsApi.md#ActionsGetReviewsForRun) | **Get** /repos/{owner}/{repo}/actions/runs/{run_id}/approvals | Get the review history for a workflow run
[**ActionsGetSelfHostedRunnerForOrg**](ActionsApi.md#ActionsGetSelfHostedRunnerForOrg) | **Get** /orgs/{org}/actions/runners/{runner_id} | Get a self-hosted runner for an organization
[**ActionsGetSelfHostedRunnerForRepo**](ActionsApi.md#ActionsGetSelfHostedRunnerForRepo) | **Get** /repos/{owner}/{repo}/actions/runners/{runner_id} | Get a self-hosted runner for a repository
[**ActionsGetSelfHostedRunnerGroupForOrg**](ActionsApi.md#ActionsGetSelfHostedRunnerGroupForOrg) | **Get** /orgs/{org}/actions/runner-groups/{runner_group_id} | Get a self-hosted runner group for an organization
[**ActionsGetWorkflow**](ActionsApi.md#ActionsGetWorkflow) | **Get** /repos/{owner}/{repo}/actions/workflows/{workflow_id} | Get a workflow
[**ActionsGetWorkflowAccessToRepository**](ActionsApi.md#ActionsGetWorkflowAccessToRepository) | **Get** /repos/{owner}/{repo}/actions/permissions/access | Get the level of access for workflows outside of the repository
[**ActionsGetWorkflowRun**](ActionsApi.md#ActionsGetWorkflowRun) | **Get** /repos/{owner}/{repo}/actions/runs/{run_id} | Get a workflow run
[**ActionsGetWorkflowRunAttempt**](ActionsApi.md#ActionsGetWorkflowRunAttempt) | **Get** /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number} | Get a workflow run attempt
[**ActionsGetWorkflowRunUsage**](ActionsApi.md#ActionsGetWorkflowRunUsage) | **Get** /repos/{owner}/{repo}/actions/runs/{run_id}/timing | Get workflow run usage
[**ActionsGetWorkflowUsage**](ActionsApi.md#ActionsGetWorkflowUsage) | **Get** /repos/{owner}/{repo}/actions/workflows/{workflow_id}/timing | Get workflow usage
[**ActionsListArtifactsForRepo**](ActionsApi.md#ActionsListArtifactsForRepo) | **Get** /repos/{owner}/{repo}/actions/artifacts | List artifacts for a repository
[**ActionsListEnvironmentSecrets**](ActionsApi.md#ActionsListEnvironmentSecrets) | **Get** /repositories/{repository_id}/environments/{environment_name}/secrets | List environment secrets
[**ActionsListJobsForWorkflowRun**](ActionsApi.md#ActionsListJobsForWorkflowRun) | **Get** /repos/{owner}/{repo}/actions/runs/{run_id}/jobs | List jobs for a workflow run
[**ActionsListJobsForWorkflowRunAttempt**](ActionsApi.md#ActionsListJobsForWorkflowRunAttempt) | **Get** /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/jobs | List jobs for a workflow run attempt
[**ActionsListLabelsForSelfHostedRunnerForOrg**](ActionsApi.md#ActionsListLabelsForSelfHostedRunnerForOrg) | **Get** /orgs/{org}/actions/runners/{runner_id}/labels | List labels for a self-hosted runner for an organization
[**ActionsListLabelsForSelfHostedRunnerForRepo**](ActionsApi.md#ActionsListLabelsForSelfHostedRunnerForRepo) | **Get** /repos/{owner}/{repo}/actions/runners/{runner_id}/labels | List labels for a self-hosted runner for a repository
[**ActionsListOrgSecrets**](ActionsApi.md#ActionsListOrgSecrets) | **Get** /orgs/{org}/actions/secrets | List organization secrets
[**ActionsListRepoAccessToSelfHostedRunnerGroupInOrg**](ActionsApi.md#ActionsListRepoAccessToSelfHostedRunnerGroupInOrg) | **Get** /orgs/{org}/actions/runner-groups/{runner_group_id}/repositories | List repository access to a self-hosted runner group in an organization
[**ActionsListRepoSecrets**](ActionsApi.md#ActionsListRepoSecrets) | **Get** /repos/{owner}/{repo}/actions/secrets | List repository secrets
[**ActionsListRepoWorkflows**](ActionsApi.md#ActionsListRepoWorkflows) | **Get** /repos/{owner}/{repo}/actions/workflows | List repository workflows
[**ActionsListRunnerApplicationsForOrg**](ActionsApi.md#ActionsListRunnerApplicationsForOrg) | **Get** /orgs/{org}/actions/runners/downloads | List runner applications for an organization
[**ActionsListRunnerApplicationsForRepo**](ActionsApi.md#ActionsListRunnerApplicationsForRepo) | **Get** /repos/{owner}/{repo}/actions/runners/downloads | List runner applications for a repository
[**ActionsListSelectedReposForOrgSecret**](ActionsApi.md#ActionsListSelectedReposForOrgSecret) | **Get** /orgs/{org}/actions/secrets/{secret_name}/repositories | List selected repositories for an organization secret
[**ActionsListSelectedRepositoriesEnabledGithubActionsOrganization**](ActionsApi.md#ActionsListSelectedRepositoriesEnabledGithubActionsOrganization) | **Get** /orgs/{org}/actions/permissions/repositories | List selected repositories enabled for GitHub Actions in an organization
[**ActionsListSelfHostedRunnerGroupsForOrg**](ActionsApi.md#ActionsListSelfHostedRunnerGroupsForOrg) | **Get** /orgs/{org}/actions/runner-groups | List self-hosted runner groups for an organization
[**ActionsListSelfHostedRunnersForOrg**](ActionsApi.md#ActionsListSelfHostedRunnersForOrg) | **Get** /orgs/{org}/actions/runners | List self-hosted runners for an organization
[**ActionsListSelfHostedRunnersForRepo**](ActionsApi.md#ActionsListSelfHostedRunnersForRepo) | **Get** /repos/{owner}/{repo}/actions/runners | List self-hosted runners for a repository
[**ActionsListSelfHostedRunnersInGroupForOrg**](ActionsApi.md#ActionsListSelfHostedRunnersInGroupForOrg) | **Get** /orgs/{org}/actions/runner-groups/{runner_group_id}/runners | List self-hosted runners in a group for an organization
[**ActionsListWorkflowRunArtifacts**](ActionsApi.md#ActionsListWorkflowRunArtifacts) | **Get** /repos/{owner}/{repo}/actions/runs/{run_id}/artifacts | List workflow run artifacts
[**ActionsListWorkflowRuns**](ActionsApi.md#ActionsListWorkflowRuns) | **Get** /repos/{owner}/{repo}/actions/workflows/{workflow_id}/runs | List workflow runs
[**ActionsListWorkflowRunsForRepo**](ActionsApi.md#ActionsListWorkflowRunsForRepo) | **Get** /repos/{owner}/{repo}/actions/runs | List workflow runs for a repository
[**ActionsReRunJobForWorkflowRun**](ActionsApi.md#ActionsReRunJobForWorkflowRun) | **Post** /repos/{owner}/{repo}/actions/jobs/{job_id}/rerun | Re-run a job from a workflow run
[**ActionsReRunWorkflow**](ActionsApi.md#ActionsReRunWorkflow) | **Post** /repos/{owner}/{repo}/actions/runs/{run_id}/rerun | Re-run a workflow
[**ActionsReRunWorkflowFailedJobs**](ActionsApi.md#ActionsReRunWorkflowFailedJobs) | **Post** /repos/{owner}/{repo}/actions/runs/{run_id}/rerun-failed-jobs | Re-run failed jobs from a workflow run
[**ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForOrg**](ActionsApi.md#ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForOrg) | **Delete** /orgs/{org}/actions/runners/{runner_id}/labels | Remove all custom labels from a self-hosted runner for an organization
[**ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForRepo**](ActionsApi.md#ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForRepo) | **Delete** /repos/{owner}/{repo}/actions/runners/{runner_id}/labels | Remove all custom labels from a self-hosted runner for a repository
[**ActionsRemoveCustomLabelFromSelfHostedRunnerForOrg**](ActionsApi.md#ActionsRemoveCustomLabelFromSelfHostedRunnerForOrg) | **Delete** /orgs/{org}/actions/runners/{runner_id}/labels/{name} | Remove a custom label from a self-hosted runner for an organization
[**ActionsRemoveCustomLabelFromSelfHostedRunnerForRepo**](ActionsApi.md#ActionsRemoveCustomLabelFromSelfHostedRunnerForRepo) | **Delete** /repos/{owner}/{repo}/actions/runners/{runner_id}/labels/{name} | Remove a custom label from a self-hosted runner for a repository
[**ActionsRemoveRepoAccessToSelfHostedRunnerGroupInOrg**](ActionsApi.md#ActionsRemoveRepoAccessToSelfHostedRunnerGroupInOrg) | **Delete** /orgs/{org}/actions/runner-groups/{runner_group_id}/repositories/{repository_id} | Remove repository access to a self-hosted runner group in an organization
[**ActionsRemoveSelectedRepoFromOrgSecret**](ActionsApi.md#ActionsRemoveSelectedRepoFromOrgSecret) | **Delete** /orgs/{org}/actions/secrets/{secret_name}/repositories/{repository_id} | Remove selected repository from an organization secret
[**ActionsRemoveSelfHostedRunnerFromGroupForOrg**](ActionsApi.md#ActionsRemoveSelfHostedRunnerFromGroupForOrg) | **Delete** /orgs/{org}/actions/runner-groups/{runner_group_id}/runners/{runner_id} | Remove a self-hosted runner from a group for an organization
[**ActionsReviewPendingDeploymentsForRun**](ActionsApi.md#ActionsReviewPendingDeploymentsForRun) | **Post** /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments | Review pending deployments for a workflow run
[**ActionsSetActionsOidcCustomIssuerPolicyForEnterprise**](ActionsApi.md#ActionsSetActionsOidcCustomIssuerPolicyForEnterprise) | **Put** /enterprises/{enterprise}/actions/oidc/customization/issuer | Set the GitHub Actions OIDC custom issuer policy for an enterprise
[**ActionsSetAllowedActionsOrganization**](ActionsApi.md#ActionsSetAllowedActionsOrganization) | **Put** /orgs/{org}/actions/permissions/selected-actions | Set allowed actions and reusable workflows for an organization
[**ActionsSetAllowedActionsRepository**](ActionsApi.md#ActionsSetAllowedActionsRepository) | **Put** /repos/{owner}/{repo}/actions/permissions/selected-actions | Set allowed actions and reusable workflows for a repository
[**ActionsSetCustomLabelsForSelfHostedRunnerForOrg**](ActionsApi.md#ActionsSetCustomLabelsForSelfHostedRunnerForOrg) | **Put** /orgs/{org}/actions/runners/{runner_id}/labels | Set custom labels for a self-hosted runner for an organization
[**ActionsSetCustomLabelsForSelfHostedRunnerForRepo**](ActionsApi.md#ActionsSetCustomLabelsForSelfHostedRunnerForRepo) | **Put** /repos/{owner}/{repo}/actions/runners/{runner_id}/labels | Set custom labels for a self-hosted runner for a repository
[**ActionsSetCustomOidcSubClaimForRepo**](ActionsApi.md#ActionsSetCustomOidcSubClaimForRepo) | **Put** /repos/{owner}/{repo}/actions/oidc/customization/sub | Set the opt-in flag of an OIDC subject claim customization for a repository
[**ActionsSetGithubActionsDefaultWorkflowPermissionsEnterprise**](ActionsApi.md#ActionsSetGithubActionsDefaultWorkflowPermissionsEnterprise) | **Put** /enterprises/{enterprise}/actions/permissions/workflow | Set default workflow permissions for an enterprise
[**ActionsSetGithubActionsDefaultWorkflowPermissionsOrganization**](ActionsApi.md#ActionsSetGithubActionsDefaultWorkflowPermissionsOrganization) | **Put** /orgs/{org}/actions/permissions/workflow | Set default workflow permissions for an organization
[**ActionsSetGithubActionsDefaultWorkflowPermissionsRepository**](ActionsApi.md#ActionsSetGithubActionsDefaultWorkflowPermissionsRepository) | **Put** /repos/{owner}/{repo}/actions/permissions/workflow | Set default workflow permissions for a repository
[**ActionsSetGithubActionsPermissionsOrganization**](ActionsApi.md#ActionsSetGithubActionsPermissionsOrganization) | **Put** /orgs/{org}/actions/permissions | Set GitHub Actions permissions for an organization
[**ActionsSetGithubActionsPermissionsRepository**](ActionsApi.md#ActionsSetGithubActionsPermissionsRepository) | **Put** /repos/{owner}/{repo}/actions/permissions | Set GitHub Actions permissions for a repository
[**ActionsSetRepoAccessToSelfHostedRunnerGroupInOrg**](ActionsApi.md#ActionsSetRepoAccessToSelfHostedRunnerGroupInOrg) | **Put** /orgs/{org}/actions/runner-groups/{runner_group_id}/repositories | Set repository access for a self-hosted runner group in an organization
[**ActionsSetSelectedReposForOrgSecret**](ActionsApi.md#ActionsSetSelectedReposForOrgSecret) | **Put** /orgs/{org}/actions/secrets/{secret_name}/repositories | Set selected repositories for an organization secret
[**ActionsSetSelectedRepositoriesEnabledGithubActionsOrganization**](ActionsApi.md#ActionsSetSelectedRepositoriesEnabledGithubActionsOrganization) | **Put** /orgs/{org}/actions/permissions/repositories | Set selected repositories enabled for GitHub Actions in an organization
[**ActionsSetSelfHostedRunnersInGroupForOrg**](ActionsApi.md#ActionsSetSelfHostedRunnersInGroupForOrg) | **Put** /orgs/{org}/actions/runner-groups/{runner_group_id}/runners | Set self-hosted runners in a group for an organization
[**ActionsSetWorkflowAccessToRepository**](ActionsApi.md#ActionsSetWorkflowAccessToRepository) | **Put** /repos/{owner}/{repo}/actions/permissions/access | Set the level of access for workflows outside of the repository
[**ActionsUpdateSelfHostedRunnerGroupForOrg**](ActionsApi.md#ActionsUpdateSelfHostedRunnerGroupForOrg) | **Patch** /orgs/{org}/actions/runner-groups/{runner_group_id} | Update a self-hosted runner group for an organization



## ActionsAddCustomLabelsToSelfHostedRunnerForOrg

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response ActionsAddCustomLabelsToSelfHostedRunnerForOrg(ctx, org, runnerId).EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest(enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest).Execute()

Add custom labels to a self-hosted runner for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.
    enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest := *openapiclient.NewEnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest([]string{"Labels_example"}) // EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsAddCustomLabelsToSelfHostedRunnerForOrg(context.Background(), org, runnerId).EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest(enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsAddCustomLabelsToSelfHostedRunnerForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsAddCustomLabelsToSelfHostedRunnerForOrg`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsAddCustomLabelsToSelfHostedRunnerForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsAddCustomLabelsToSelfHostedRunnerForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest** | [**EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest**](EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest.md) |  | 

### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsAddCustomLabelsToSelfHostedRunnerForRepo

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response ActionsAddCustomLabelsToSelfHostedRunnerForRepo(ctx, owner, repo, runnerId).EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest(enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest).Execute()

Add custom labels to a self-hosted runner for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.
    enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest := *openapiclient.NewEnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest([]string{"Labels_example"}) // EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsAddCustomLabelsToSelfHostedRunnerForRepo(context.Background(), owner, repo, runnerId).EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest(enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsAddCustomLabelsToSelfHostedRunnerForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsAddCustomLabelsToSelfHostedRunnerForRepo`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsAddCustomLabelsToSelfHostedRunnerForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsAddCustomLabelsToSelfHostedRunnerForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest** | [**EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest**](EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest.md) |  | 

### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsAddRepoAccessToSelfHostedRunnerGroupInOrg

> ActionsAddRepoAccessToSelfHostedRunnerGroupInOrg(ctx, org, runnerGroupId, repositoryId).Execute()

Add repository access to a self-hosted runner group in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    repositoryId := int32(56) // int32 | The unique identifier of the repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsAddRepoAccessToSelfHostedRunnerGroupInOrg(context.Background(), org, runnerGroupId, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsAddRepoAccessToSelfHostedRunnerGroupInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 
**repositoryId** | **int32** | The unique identifier of the repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsAddRepoAccessToSelfHostedRunnerGroupInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsAddSelectedRepoToOrgSecret

> ActionsAddSelectedRepoToOrgSecret(ctx, org, secretName, repositoryId).Execute()

Add selected repository to an organization secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    secretName := "secretName_example" // string | The name of the secret.
    repositoryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsAddSelectedRepoToOrgSecret(context.Background(), org, secretName, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsAddSelectedRepoToOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 
**repositoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsAddSelectedRepoToOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsAddSelfHostedRunnerToGroupForOrg

> ActionsAddSelfHostedRunnerToGroupForOrg(ctx, org, runnerGroupId, runnerId).Execute()

Add a self-hosted runner to a group for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsAddSelfHostedRunnerToGroupForOrg(context.Background(), org, runnerGroupId, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsAddSelfHostedRunnerToGroupForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsAddSelfHostedRunnerToGroupForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsApproveWorkflowRun

> map[string]interface{} ActionsApproveWorkflowRun(ctx, owner, repo, runId).Execute()

Approve a workflow run for a fork pull request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsApproveWorkflowRun(context.Background(), owner, repo, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsApproveWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsApproveWorkflowRun`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsApproveWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsApproveWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsCancelWorkflowRun

> map[string]interface{} ActionsCancelWorkflowRun(ctx, owner, repo, runId).Execute()

Cancel a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsCancelWorkflowRun(context.Background(), owner, repo, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsCancelWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsCancelWorkflowRun`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsCancelWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCancelWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsCreateOrUpdateEnvironmentSecret

> map[string]interface{} ActionsCreateOrUpdateEnvironmentSecret(ctx, repositoryId, environmentName, secretName).ActionsCreateOrUpdateEnvironmentSecretRequest(actionsCreateOrUpdateEnvironmentSecretRequest).Execute()

Create or update an environment secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repositoryId := int32(56) // int32 | The unique identifier of the repository.
    environmentName := "environmentName_example" // string | The name of the environment.
    secretName := "secretName_example" // string | The name of the secret.
    actionsCreateOrUpdateEnvironmentSecretRequest := *openapiclient.NewActionsCreateOrUpdateEnvironmentSecretRequest("EncryptedValue_example", "KeyId_example") // ActionsCreateOrUpdateEnvironmentSecretRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsCreateOrUpdateEnvironmentSecret(context.Background(), repositoryId, environmentName, secretName).ActionsCreateOrUpdateEnvironmentSecretRequest(actionsCreateOrUpdateEnvironmentSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsCreateOrUpdateEnvironmentSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsCreateOrUpdateEnvironmentSecret`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsCreateOrUpdateEnvironmentSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | **int32** | The unique identifier of the repository. | 
**environmentName** | **string** | The name of the environment. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCreateOrUpdateEnvironmentSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **actionsCreateOrUpdateEnvironmentSecretRequest** | [**ActionsCreateOrUpdateEnvironmentSecretRequest**](ActionsCreateOrUpdateEnvironmentSecretRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsCreateOrUpdateOrgSecret

> map[string]interface{} ActionsCreateOrUpdateOrgSecret(ctx, org, secretName).ActionsCreateOrUpdateOrgSecretRequest(actionsCreateOrUpdateOrgSecretRequest).Execute()

Create or update an organization secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    secretName := "secretName_example" // string | The name of the secret.
    actionsCreateOrUpdateOrgSecretRequest := *openapiclient.NewActionsCreateOrUpdateOrgSecretRequest("Visibility_example") // ActionsCreateOrUpdateOrgSecretRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsCreateOrUpdateOrgSecret(context.Background(), org, secretName).ActionsCreateOrUpdateOrgSecretRequest(actionsCreateOrUpdateOrgSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsCreateOrUpdateOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsCreateOrUpdateOrgSecret`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsCreateOrUpdateOrgSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCreateOrUpdateOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actionsCreateOrUpdateOrgSecretRequest** | [**ActionsCreateOrUpdateOrgSecretRequest**](ActionsCreateOrUpdateOrgSecretRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsCreateOrUpdateRepoSecret

> map[string]interface{} ActionsCreateOrUpdateRepoSecret(ctx, owner, repo, secretName).ActionsCreateOrUpdateRepoSecretRequest(actionsCreateOrUpdateRepoSecretRequest).Execute()

Create or update a repository secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    secretName := "secretName_example" // string | The name of the secret.
    actionsCreateOrUpdateRepoSecretRequest := *openapiclient.NewActionsCreateOrUpdateRepoSecretRequest() // ActionsCreateOrUpdateRepoSecretRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsCreateOrUpdateRepoSecret(context.Background(), owner, repo, secretName).ActionsCreateOrUpdateRepoSecretRequest(actionsCreateOrUpdateRepoSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsCreateOrUpdateRepoSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsCreateOrUpdateRepoSecret`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsCreateOrUpdateRepoSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCreateOrUpdateRepoSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **actionsCreateOrUpdateRepoSecretRequest** | [**ActionsCreateOrUpdateRepoSecretRequest**](ActionsCreateOrUpdateRepoSecretRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsCreateRegistrationTokenForOrg

> AuthenticationToken ActionsCreateRegistrationTokenForOrg(ctx, org).Execute()

Create a registration token for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsCreateRegistrationTokenForOrg(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsCreateRegistrationTokenForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsCreateRegistrationTokenForOrg`: AuthenticationToken
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsCreateRegistrationTokenForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCreateRegistrationTokenForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationToken**](AuthenticationToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsCreateRegistrationTokenForRepo

> AuthenticationToken ActionsCreateRegistrationTokenForRepo(ctx, owner, repo).Execute()

Create a registration token for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsCreateRegistrationTokenForRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsCreateRegistrationTokenForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsCreateRegistrationTokenForRepo`: AuthenticationToken
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsCreateRegistrationTokenForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCreateRegistrationTokenForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthenticationToken**](AuthenticationToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsCreateRemoveTokenForOrg

> AuthenticationToken ActionsCreateRemoveTokenForOrg(ctx, org).Execute()

Create a remove token for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsCreateRemoveTokenForOrg(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsCreateRemoveTokenForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsCreateRemoveTokenForOrg`: AuthenticationToken
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsCreateRemoveTokenForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCreateRemoveTokenForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationToken**](AuthenticationToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsCreateRemoveTokenForRepo

> AuthenticationToken ActionsCreateRemoveTokenForRepo(ctx, owner, repo).Execute()

Create a remove token for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsCreateRemoveTokenForRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsCreateRemoveTokenForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsCreateRemoveTokenForRepo`: AuthenticationToken
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsCreateRemoveTokenForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCreateRemoveTokenForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthenticationToken**](AuthenticationToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsCreateSelfHostedRunnerGroupForOrg

> RunnerGroupsOrg ActionsCreateSelfHostedRunnerGroupForOrg(ctx, org).ActionsCreateSelfHostedRunnerGroupForOrgRequest(actionsCreateSelfHostedRunnerGroupForOrgRequest).Execute()

Create a self-hosted runner group for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    actionsCreateSelfHostedRunnerGroupForOrgRequest := *openapiclient.NewActionsCreateSelfHostedRunnerGroupForOrgRequest("Name_example") // ActionsCreateSelfHostedRunnerGroupForOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsCreateSelfHostedRunnerGroupForOrg(context.Background(), org).ActionsCreateSelfHostedRunnerGroupForOrgRequest(actionsCreateSelfHostedRunnerGroupForOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsCreateSelfHostedRunnerGroupForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsCreateSelfHostedRunnerGroupForOrg`: RunnerGroupsOrg
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsCreateSelfHostedRunnerGroupForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCreateSelfHostedRunnerGroupForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsCreateSelfHostedRunnerGroupForOrgRequest** | [**ActionsCreateSelfHostedRunnerGroupForOrgRequest**](ActionsCreateSelfHostedRunnerGroupForOrgRequest.md) |  | 

### Return type

[**RunnerGroupsOrg**](RunnerGroupsOrg.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsCreateWorkflowDispatch

> ActionsCreateWorkflowDispatch(ctx, owner, repo, workflowId).ActionsCreateWorkflowDispatchRequest(actionsCreateWorkflowDispatchRequest).Execute()

Create a workflow dispatch event



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    workflowId := openapiclient.actions_get_workflow_workflow_id_parameter{Int32: new(int32)} // ActionsGetWorkflowWorkflowIdParameter | The ID of the workflow. You can also pass the workflow file name as a string.
    actionsCreateWorkflowDispatchRequest := *openapiclient.NewActionsCreateWorkflowDispatchRequest("Ref_example") // ActionsCreateWorkflowDispatchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsCreateWorkflowDispatch(context.Background(), owner, repo, workflowId).ActionsCreateWorkflowDispatchRequest(actionsCreateWorkflowDispatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsCreateWorkflowDispatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**workflowId** | [**ActionsGetWorkflowWorkflowIdParameter**](.md) | The ID of the workflow. You can also pass the workflow file name as a string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCreateWorkflowDispatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **actionsCreateWorkflowDispatchRequest** | [**ActionsCreateWorkflowDispatchRequest**](ActionsCreateWorkflowDispatchRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteActionsCacheById

> ActionsDeleteActionsCacheById(ctx, owner, repo, cacheId).Execute()

Delete a GitHub Actions cache for a repository (using a cache ID)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    cacheId := int32(56) // int32 | The unique identifier of the GitHub Actions cache.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteActionsCacheById(context.Background(), owner, repo, cacheId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteActionsCacheById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**cacheId** | **int32** | The unique identifier of the GitHub Actions cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteActionsCacheByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteActionsCacheByKey

> ActionsCacheList ActionsDeleteActionsCacheByKey(ctx, owner, repo).Key(key).Ref(ref).Execute()

Delete GitHub Actions caches for a repository (using a cache key)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    key := "key_example" // string | A key for identifying the cache.
    ref := "ref_example" // string | The Git reference for the results you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteActionsCacheByKey(context.Background(), owner, repo).Key(key).Ref(ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteActionsCacheByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsDeleteActionsCacheByKey`: ActionsCacheList
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsDeleteActionsCacheByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteActionsCacheByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | **string** | A key for identifying the cache. | 
 **ref** | **string** | The Git reference for the results you want to list. The &#x60;ref&#x60; for a branch can be formatted either as &#x60;refs/heads/&lt;branch name&gt;&#x60; or simply &#x60;&lt;branch name&gt;&#x60;. To reference a pull request use &#x60;refs/pull/&lt;number&gt;/merge&#x60;. | 

### Return type

[**ActionsCacheList**](ActionsCacheList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteArtifact

> ActionsDeleteArtifact(ctx, owner, repo, artifactId).Execute()

Delete an artifact



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    artifactId := int32(56) // int32 | The unique identifier of the artifact.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteArtifact(context.Background(), owner, repo, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**artifactId** | **int32** | The unique identifier of the artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteEnvironmentSecret

> ActionsDeleteEnvironmentSecret(ctx, repositoryId, environmentName, secretName).Execute()

Delete an environment secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repositoryId := int32(56) // int32 | The unique identifier of the repository.
    environmentName := "environmentName_example" // string | The name of the environment.
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteEnvironmentSecret(context.Background(), repositoryId, environmentName, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteEnvironmentSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | **int32** | The unique identifier of the repository. | 
**environmentName** | **string** | The name of the environment. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteEnvironmentSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteOrgSecret

> ActionsDeleteOrgSecret(ctx, org, secretName).Execute()

Delete an organization secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteOrgSecret(context.Background(), org, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteRepoSecret

> ActionsDeleteRepoSecret(ctx, owner, repo, secretName).Execute()

Delete a repository secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteRepoSecret(context.Background(), owner, repo, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteRepoSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteRepoSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteSelfHostedRunnerFromOrg

> ActionsDeleteSelfHostedRunnerFromOrg(ctx, org, runnerId).Execute()

Delete a self-hosted runner from an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteSelfHostedRunnerFromOrg(context.Background(), org, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteSelfHostedRunnerFromOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteSelfHostedRunnerFromOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteSelfHostedRunnerFromRepo

> ActionsDeleteSelfHostedRunnerFromRepo(ctx, owner, repo, runnerId).Execute()

Delete a self-hosted runner from a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteSelfHostedRunnerFromRepo(context.Background(), owner, repo, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteSelfHostedRunnerFromRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteSelfHostedRunnerFromRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteSelfHostedRunnerGroupFromOrg

> ActionsDeleteSelfHostedRunnerGroupFromOrg(ctx, org, runnerGroupId).Execute()

Delete a self-hosted runner group from an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteSelfHostedRunnerGroupFromOrg(context.Background(), org, runnerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteSelfHostedRunnerGroupFromOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteSelfHostedRunnerGroupFromOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteWorkflowRun

> ActionsDeleteWorkflowRun(ctx, owner, repo, runId).Execute()

Delete a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteWorkflowRun(context.Background(), owner, repo, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDeleteWorkflowRunLogs

> ActionsDeleteWorkflowRunLogs(ctx, owner, repo, runId).Execute()

Delete workflow run logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDeleteWorkflowRunLogs(context.Background(), owner, repo, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDeleteWorkflowRunLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDeleteWorkflowRunLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDisableSelectedRepositoryGithubActionsOrganization

> ActionsDisableSelectedRepositoryGithubActionsOrganization(ctx, org, repositoryId).Execute()

Disable a selected repository for GitHub Actions in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    repositoryId := int32(56) // int32 | The unique identifier of the repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDisableSelectedRepositoryGithubActionsOrganization(context.Background(), org, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDisableSelectedRepositoryGithubActionsOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**repositoryId** | **int32** | The unique identifier of the repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDisableSelectedRepositoryGithubActionsOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDisableWorkflow

> ActionsDisableWorkflow(ctx, owner, repo, workflowId).Execute()

Disable a workflow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    workflowId := openapiclient.actions_get_workflow_workflow_id_parameter{Int32: new(int32)} // ActionsGetWorkflowWorkflowIdParameter | The ID of the workflow. You can also pass the workflow file name as a string.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDisableWorkflow(context.Background(), owner, repo, workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDisableWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**workflowId** | [**ActionsGetWorkflowWorkflowIdParameter**](.md) | The ID of the workflow. You can also pass the workflow file name as a string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDisableWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDownloadArtifact

> ActionsDownloadArtifact(ctx, owner, repo, artifactId, archiveFormat).Execute()

Download an artifact



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    artifactId := int32(56) // int32 | The unique identifier of the artifact.
    archiveFormat := "archiveFormat_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDownloadArtifact(context.Background(), owner, repo, artifactId, archiveFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDownloadArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**artifactId** | **int32** | The unique identifier of the artifact. | 
**archiveFormat** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDownloadArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDownloadJobLogsForWorkflowRun

> ActionsDownloadJobLogsForWorkflowRun(ctx, owner, repo, jobId).Execute()

Download job logs for a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    jobId := int32(56) // int32 | The unique identifier of the job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDownloadJobLogsForWorkflowRun(context.Background(), owner, repo, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDownloadJobLogsForWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**jobId** | **int32** | The unique identifier of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDownloadJobLogsForWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDownloadWorkflowRunAttemptLogs

> ActionsDownloadWorkflowRunAttemptLogs(ctx, owner, repo, runId, attemptNumber).Execute()

Download workflow run attempt logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.
    attemptNumber := int32(56) // int32 | The attempt number of the workflow run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDownloadWorkflowRunAttemptLogs(context.Background(), owner, repo, runId, attemptNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDownloadWorkflowRunAttemptLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 
**attemptNumber** | **int32** | The attempt number of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDownloadWorkflowRunAttemptLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDownloadWorkflowRunLogs

> ActionsDownloadWorkflowRunLogs(ctx, owner, repo, runId).Execute()

Download workflow run logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsDownloadWorkflowRunLogs(context.Background(), owner, repo, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsDownloadWorkflowRunLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDownloadWorkflowRunLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsEnableSelectedRepositoryGithubActionsOrganization

> ActionsEnableSelectedRepositoryGithubActionsOrganization(ctx, org, repositoryId).Execute()

Enable a selected repository for GitHub Actions in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    repositoryId := int32(56) // int32 | The unique identifier of the repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsEnableSelectedRepositoryGithubActionsOrganization(context.Background(), org, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsEnableSelectedRepositoryGithubActionsOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**repositoryId** | **int32** | The unique identifier of the repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsEnableSelectedRepositoryGithubActionsOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsEnableWorkflow

> ActionsEnableWorkflow(ctx, owner, repo, workflowId).Execute()

Enable a workflow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    workflowId := openapiclient.actions_get_workflow_workflow_id_parameter{Int32: new(int32)} // ActionsGetWorkflowWorkflowIdParameter | The ID of the workflow. You can also pass the workflow file name as a string.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsEnableWorkflow(context.Background(), owner, repo, workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsEnableWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**workflowId** | [**ActionsGetWorkflowWorkflowIdParameter**](.md) | The ID of the workflow. You can also pass the workflow file name as a string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsEnableWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetActionsCacheList

> ActionsCacheList ActionsGetActionsCacheList(ctx, owner, repo).PerPage(perPage).Page(page).Ref(ref).Key(key).Sort(sort).Direction(direction).Execute()

List GitHub Actions caches for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    ref := "ref_example" // string | The Git reference for the results you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`. (optional)
    key := "key_example" // string | An explicit key or prefix for identifying the cache (optional)
    sort := "sort_example" // string | The property to sort the results by. `created_at` means when the cache was created. `last_accessed_at` means when the cache was last accessed. `size_in_bytes` is the size of the cache in bytes. (optional) (default to "last_accessed_at")
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetActionsCacheList(context.Background(), owner, repo).PerPage(perPage).Page(page).Ref(ref).Key(key).Sort(sort).Direction(direction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetActionsCacheList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetActionsCacheList`: ActionsCacheList
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetActionsCacheList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetActionsCacheListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **ref** | **string** | The Git reference for the results you want to list. The &#x60;ref&#x60; for a branch can be formatted either as &#x60;refs/heads/&lt;branch name&gt;&#x60; or simply &#x60;&lt;branch name&gt;&#x60;. To reference a pull request use &#x60;refs/pull/&lt;number&gt;/merge&#x60;. | 
 **key** | **string** | An explicit key or prefix for identifying the cache | 
 **sort** | **string** | The property to sort the results by. &#x60;created_at&#x60; means when the cache was created. &#x60;last_accessed_at&#x60; means when the cache was last accessed. &#x60;size_in_bytes&#x60; is the size of the cache in bytes. | [default to &quot;last_accessed_at&quot;]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]

### Return type

[**ActionsCacheList**](ActionsCacheList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetActionsCacheUsage

> ActionsCacheUsageByRepository ActionsGetActionsCacheUsage(ctx, owner, repo).Execute()

Get GitHub Actions cache usage for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetActionsCacheUsage(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetActionsCacheUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetActionsCacheUsage`: ActionsCacheUsageByRepository
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetActionsCacheUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetActionsCacheUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActionsCacheUsageByRepository**](ActionsCacheUsageByRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetActionsCacheUsageByRepoForOrg

> ActionsGetActionsCacheUsageByRepoForOrg200Response ActionsGetActionsCacheUsageByRepoForOrg(ctx, org).PerPage(perPage).Page(page).Execute()

List repositories with GitHub Actions cache usage for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetActionsCacheUsageByRepoForOrg(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetActionsCacheUsageByRepoForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetActionsCacheUsageByRepoForOrg`: ActionsGetActionsCacheUsageByRepoForOrg200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetActionsCacheUsageByRepoForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetActionsCacheUsageByRepoForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsGetActionsCacheUsageByRepoForOrg200Response**](ActionsGetActionsCacheUsageByRepoForOrg200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetActionsCacheUsageForEnterprise

> ActionsCacheUsageOrgEnterprise ActionsGetActionsCacheUsageForEnterprise(ctx, enterprise).Execute()

Get GitHub Actions cache usage for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetActionsCacheUsageForEnterprise(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetActionsCacheUsageForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetActionsCacheUsageForEnterprise`: ActionsCacheUsageOrgEnterprise
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetActionsCacheUsageForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetActionsCacheUsageForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsCacheUsageOrgEnterprise**](ActionsCacheUsageOrgEnterprise.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetActionsCacheUsageForOrg

> ActionsCacheUsageOrgEnterprise ActionsGetActionsCacheUsageForOrg(ctx, org).Execute()

Get GitHub Actions cache usage for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetActionsCacheUsageForOrg(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetActionsCacheUsageForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetActionsCacheUsageForOrg`: ActionsCacheUsageOrgEnterprise
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetActionsCacheUsageForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetActionsCacheUsageForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsCacheUsageOrgEnterprise**](ActionsCacheUsageOrgEnterprise.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetAllowedActionsOrganization

> SelectedActions ActionsGetAllowedActionsOrganization(ctx, org).Execute()

Get allowed actions and reusable workflows for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetAllowedActionsOrganization(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetAllowedActionsOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetAllowedActionsOrganization`: SelectedActions
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetAllowedActionsOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetAllowedActionsOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SelectedActions**](SelectedActions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetAllowedActionsRepository

> SelectedActions ActionsGetAllowedActionsRepository(ctx, owner, repo).Execute()

Get allowed actions and reusable workflows for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetAllowedActionsRepository(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetAllowedActionsRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetAllowedActionsRepository`: SelectedActions
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetAllowedActionsRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetAllowedActionsRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SelectedActions**](SelectedActions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetArtifact

> Artifact ActionsGetArtifact(ctx, owner, repo, artifactId).Execute()

Get an artifact



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    artifactId := int32(56) // int32 | The unique identifier of the artifact.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetArtifact(context.Background(), owner, repo, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetArtifact`: Artifact
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**artifactId** | **int32** | The unique identifier of the artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Artifact**](Artifact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetCustomOidcSubClaimForRepo

> OptOutOidcCustomSub ActionsGetCustomOidcSubClaimForRepo(ctx, owner, repo).Execute()

Get the opt-out flag of an OIDC subject claim customization for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetCustomOidcSubClaimForRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetCustomOidcSubClaimForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetCustomOidcSubClaimForRepo`: OptOutOidcCustomSub
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetCustomOidcSubClaimForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetCustomOidcSubClaimForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OptOutOidcCustomSub**](OptOutOidcCustomSub.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetEnvironmentPublicKey

> ActionsPublicKey ActionsGetEnvironmentPublicKey(ctx, repositoryId, environmentName).Execute()

Get an environment public key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repositoryId := int32(56) // int32 | The unique identifier of the repository.
    environmentName := "environmentName_example" // string | The name of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetEnvironmentPublicKey(context.Background(), repositoryId, environmentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetEnvironmentPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetEnvironmentPublicKey`: ActionsPublicKey
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetEnvironmentPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | **int32** | The unique identifier of the repository. | 
**environmentName** | **string** | The name of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetEnvironmentPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActionsPublicKey**](ActionsPublicKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetEnvironmentSecret

> ActionsSecret ActionsGetEnvironmentSecret(ctx, repositoryId, environmentName, secretName).Execute()

Get an environment secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repositoryId := int32(56) // int32 | The unique identifier of the repository.
    environmentName := "environmentName_example" // string | The name of the environment.
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetEnvironmentSecret(context.Background(), repositoryId, environmentName, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetEnvironmentSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetEnvironmentSecret`: ActionsSecret
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetEnvironmentSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | **int32** | The unique identifier of the repository. | 
**environmentName** | **string** | The name of the environment. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetEnvironmentSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ActionsSecret**](ActionsSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetGithubActionsDefaultWorkflowPermissionsEnterprise

> ActionsGetDefaultWorkflowPermissions ActionsGetGithubActionsDefaultWorkflowPermissionsEnterprise(ctx, enterprise).Execute()

Get default workflow permissions for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetGithubActionsDefaultWorkflowPermissionsEnterprise(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetGithubActionsDefaultWorkflowPermissionsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetGithubActionsDefaultWorkflowPermissionsEnterprise`: ActionsGetDefaultWorkflowPermissions
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetGithubActionsDefaultWorkflowPermissionsEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetGithubActionsDefaultWorkflowPermissionsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsGetDefaultWorkflowPermissions**](ActionsGetDefaultWorkflowPermissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetGithubActionsDefaultWorkflowPermissionsOrganization

> ActionsGetDefaultWorkflowPermissions ActionsGetGithubActionsDefaultWorkflowPermissionsOrganization(ctx, org).Execute()

Get default workflow permissions for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetGithubActionsDefaultWorkflowPermissionsOrganization(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetGithubActionsDefaultWorkflowPermissionsOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetGithubActionsDefaultWorkflowPermissionsOrganization`: ActionsGetDefaultWorkflowPermissions
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetGithubActionsDefaultWorkflowPermissionsOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetGithubActionsDefaultWorkflowPermissionsOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsGetDefaultWorkflowPermissions**](ActionsGetDefaultWorkflowPermissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetGithubActionsDefaultWorkflowPermissionsRepository

> ActionsGetDefaultWorkflowPermissions ActionsGetGithubActionsDefaultWorkflowPermissionsRepository(ctx, owner, repo).Execute()

Get default workflow permissions for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetGithubActionsDefaultWorkflowPermissionsRepository(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetGithubActionsDefaultWorkflowPermissionsRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetGithubActionsDefaultWorkflowPermissionsRepository`: ActionsGetDefaultWorkflowPermissions
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetGithubActionsDefaultWorkflowPermissionsRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetGithubActionsDefaultWorkflowPermissionsRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActionsGetDefaultWorkflowPermissions**](ActionsGetDefaultWorkflowPermissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetGithubActionsPermissionsOrganization

> ActionsOrganizationPermissions ActionsGetGithubActionsPermissionsOrganization(ctx, org).Execute()

Get GitHub Actions permissions for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetGithubActionsPermissionsOrganization(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetGithubActionsPermissionsOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetGithubActionsPermissionsOrganization`: ActionsOrganizationPermissions
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetGithubActionsPermissionsOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetGithubActionsPermissionsOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsOrganizationPermissions**](ActionsOrganizationPermissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetGithubActionsPermissionsRepository

> ActionsRepositoryPermissions ActionsGetGithubActionsPermissionsRepository(ctx, owner, repo).Execute()

Get GitHub Actions permissions for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetGithubActionsPermissionsRepository(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetGithubActionsPermissionsRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetGithubActionsPermissionsRepository`: ActionsRepositoryPermissions
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetGithubActionsPermissionsRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetGithubActionsPermissionsRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActionsRepositoryPermissions**](ActionsRepositoryPermissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetJobForWorkflowRun

> Job ActionsGetJobForWorkflowRun(ctx, owner, repo, jobId).Execute()

Get a job for a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    jobId := int32(56) // int32 | The unique identifier of the job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetJobForWorkflowRun(context.Background(), owner, repo, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetJobForWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetJobForWorkflowRun`: Job
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetJobForWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**jobId** | **int32** | The unique identifier of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetJobForWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Job**](Job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetOrgPublicKey

> ActionsPublicKey ActionsGetOrgPublicKey(ctx, org).Execute()

Get an organization public key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetOrgPublicKey(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetOrgPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetOrgPublicKey`: ActionsPublicKey
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetOrgPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetOrgPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsPublicKey**](ActionsPublicKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetOrgSecret

> OrganizationActionsSecret ActionsGetOrgSecret(ctx, org, secretName).Execute()

Get an organization secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetOrgSecret(context.Background(), org, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetOrgSecret`: OrganizationActionsSecret
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetOrgSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationActionsSecret**](OrganizationActionsSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetPendingDeploymentsForRun

> []PendingDeployment ActionsGetPendingDeploymentsForRun(ctx, owner, repo, runId).Execute()

Get pending deployments for a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetPendingDeploymentsForRun(context.Background(), owner, repo, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetPendingDeploymentsForRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetPendingDeploymentsForRun`: []PendingDeployment
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetPendingDeploymentsForRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetPendingDeploymentsForRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]PendingDeployment**](PendingDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetRepoPublicKey

> ActionsPublicKey ActionsGetRepoPublicKey(ctx, owner, repo).Execute()

Get a repository public key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetRepoPublicKey(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetRepoPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetRepoPublicKey`: ActionsPublicKey
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetRepoPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetRepoPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActionsPublicKey**](ActionsPublicKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetRepoSecret

> ActionsSecret ActionsGetRepoSecret(ctx, owner, repo, secretName).Execute()

Get a repository secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetRepoSecret(context.Background(), owner, repo, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetRepoSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetRepoSecret`: ActionsSecret
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetRepoSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetRepoSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ActionsSecret**](ActionsSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetReviewsForRun

> []EnvironmentApprovals ActionsGetReviewsForRun(ctx, owner, repo, runId).Execute()

Get the review history for a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetReviewsForRun(context.Background(), owner, repo, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetReviewsForRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetReviewsForRun`: []EnvironmentApprovals
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetReviewsForRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetReviewsForRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]EnvironmentApprovals**](EnvironmentApprovals.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetSelfHostedRunnerForOrg

> Runner ActionsGetSelfHostedRunnerForOrg(ctx, org, runnerId).Execute()

Get a self-hosted runner for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetSelfHostedRunnerForOrg(context.Background(), org, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetSelfHostedRunnerForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetSelfHostedRunnerForOrg`: Runner
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetSelfHostedRunnerForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetSelfHostedRunnerForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Runner**](Runner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetSelfHostedRunnerForRepo

> Runner ActionsGetSelfHostedRunnerForRepo(ctx, owner, repo, runnerId).Execute()

Get a self-hosted runner for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetSelfHostedRunnerForRepo(context.Background(), owner, repo, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetSelfHostedRunnerForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetSelfHostedRunnerForRepo`: Runner
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetSelfHostedRunnerForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetSelfHostedRunnerForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Runner**](Runner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetSelfHostedRunnerGroupForOrg

> RunnerGroupsOrg ActionsGetSelfHostedRunnerGroupForOrg(ctx, org, runnerGroupId).Execute()

Get a self-hosted runner group for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetSelfHostedRunnerGroupForOrg(context.Background(), org, runnerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetSelfHostedRunnerGroupForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetSelfHostedRunnerGroupForOrg`: RunnerGroupsOrg
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetSelfHostedRunnerGroupForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetSelfHostedRunnerGroupForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RunnerGroupsOrg**](RunnerGroupsOrg.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetWorkflow

> Workflow ActionsGetWorkflow(ctx, owner, repo, workflowId).Execute()

Get a workflow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    workflowId := openapiclient.actions_get_workflow_workflow_id_parameter{Int32: new(int32)} // ActionsGetWorkflowWorkflowIdParameter | The ID of the workflow. You can also pass the workflow file name as a string.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetWorkflow(context.Background(), owner, repo, workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetWorkflow`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**workflowId** | [**ActionsGetWorkflowWorkflowIdParameter**](.md) | The ID of the workflow. You can also pass the workflow file name as a string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Workflow**](Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetWorkflowAccessToRepository

> ActionsWorkflowAccessToRepository ActionsGetWorkflowAccessToRepository(ctx, owner, repo).Execute()

Get the level of access for workflows outside of the repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetWorkflowAccessToRepository(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetWorkflowAccessToRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetWorkflowAccessToRepository`: ActionsWorkflowAccessToRepository
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetWorkflowAccessToRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetWorkflowAccessToRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActionsWorkflowAccessToRepository**](ActionsWorkflowAccessToRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetWorkflowRun

> WorkflowRun ActionsGetWorkflowRun(ctx, owner, repo, runId).ExcludePullRequests(excludePullRequests).Execute()

Get a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.
    excludePullRequests := true // bool | If `true` pull requests are omitted from the response (empty array). (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetWorkflowRun(context.Background(), owner, repo, runId).ExcludePullRequests(excludePullRequests).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetWorkflowRun`: WorkflowRun
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **excludePullRequests** | **bool** | If &#x60;true&#x60; pull requests are omitted from the response (empty array). | [default to false]

### Return type

[**WorkflowRun**](WorkflowRun.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetWorkflowRunAttempt

> WorkflowRun ActionsGetWorkflowRunAttempt(ctx, owner, repo, runId, attemptNumber).ExcludePullRequests(excludePullRequests).Execute()

Get a workflow run attempt



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.
    attemptNumber := int32(56) // int32 | The attempt number of the workflow run.
    excludePullRequests := true // bool | If `true` pull requests are omitted from the response (empty array). (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetWorkflowRunAttempt(context.Background(), owner, repo, runId, attemptNumber).ExcludePullRequests(excludePullRequests).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetWorkflowRunAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetWorkflowRunAttempt`: WorkflowRun
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetWorkflowRunAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 
**attemptNumber** | **int32** | The attempt number of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetWorkflowRunAttemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **excludePullRequests** | **bool** | If &#x60;true&#x60; pull requests are omitted from the response (empty array). | [default to false]

### Return type

[**WorkflowRun**](WorkflowRun.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetWorkflowRunUsage

> WorkflowRunUsage ActionsGetWorkflowRunUsage(ctx, owner, repo, runId).Execute()

Get workflow run usage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetWorkflowRunUsage(context.Background(), owner, repo, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetWorkflowRunUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetWorkflowRunUsage`: WorkflowRunUsage
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetWorkflowRunUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetWorkflowRunUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkflowRunUsage**](WorkflowRunUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsGetWorkflowUsage

> WorkflowUsage ActionsGetWorkflowUsage(ctx, owner, repo, workflowId).Execute()

Get workflow usage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    workflowId := openapiclient.actions_get_workflow_workflow_id_parameter{Int32: new(int32)} // ActionsGetWorkflowWorkflowIdParameter | The ID of the workflow. You can also pass the workflow file name as a string.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsGetWorkflowUsage(context.Background(), owner, repo, workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsGetWorkflowUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsGetWorkflowUsage`: WorkflowUsage
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsGetWorkflowUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**workflowId** | [**ActionsGetWorkflowWorkflowIdParameter**](.md) | The ID of the workflow. You can also pass the workflow file name as a string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetWorkflowUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkflowUsage**](WorkflowUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListArtifactsForRepo

> ActionsListArtifactsForRepo200Response ActionsListArtifactsForRepo(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List artifacts for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListArtifactsForRepo(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListArtifactsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListArtifactsForRepo`: ActionsListArtifactsForRepo200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListArtifactsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListArtifactsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListArtifactsForRepo200Response**](ActionsListArtifactsForRepo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListEnvironmentSecrets

> ActionsListRepoSecrets200Response ActionsListEnvironmentSecrets(ctx, repositoryId, environmentName).PerPage(perPage).Page(page).Execute()

List environment secrets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repositoryId := int32(56) // int32 | The unique identifier of the repository.
    environmentName := "environmentName_example" // string | The name of the environment.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListEnvironmentSecrets(context.Background(), repositoryId, environmentName).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListEnvironmentSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListEnvironmentSecrets`: ActionsListRepoSecrets200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListEnvironmentSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | **int32** | The unique identifier of the repository. | 
**environmentName** | **string** | The name of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListEnvironmentSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListRepoSecrets200Response**](ActionsListRepoSecrets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListJobsForWorkflowRun

> ActionsListJobsForWorkflowRunAttempt200Response ActionsListJobsForWorkflowRun(ctx, owner, repo, runId).Filter(filter).PerPage(perPage).Page(page).Execute()

List jobs for a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.
    filter := "filter_example" // string | Filters jobs by their `completed_at` timestamp. `latest` returns jobs from the most recent execution of the workflow run. `all` returns all jobs for a workflow run, including from old executions of the workflow run. (optional) (default to "latest")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListJobsForWorkflowRun(context.Background(), owner, repo, runId).Filter(filter).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListJobsForWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListJobsForWorkflowRun`: ActionsListJobsForWorkflowRunAttempt200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListJobsForWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListJobsForWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **string** | Filters jobs by their &#x60;completed_at&#x60; timestamp. &#x60;latest&#x60; returns jobs from the most recent execution of the workflow run. &#x60;all&#x60; returns all jobs for a workflow run, including from old executions of the workflow run. | [default to &quot;latest&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListJobsForWorkflowRunAttempt200Response**](ActionsListJobsForWorkflowRunAttempt200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListJobsForWorkflowRunAttempt

> ActionsListJobsForWorkflowRunAttempt200Response ActionsListJobsForWorkflowRunAttempt(ctx, owner, repo, runId, attemptNumber).PerPage(perPage).Page(page).Execute()

List jobs for a workflow run attempt



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.
    attemptNumber := int32(56) // int32 | The attempt number of the workflow run.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListJobsForWorkflowRunAttempt(context.Background(), owner, repo, runId, attemptNumber).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListJobsForWorkflowRunAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListJobsForWorkflowRunAttempt`: ActionsListJobsForWorkflowRunAttempt200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListJobsForWorkflowRunAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 
**attemptNumber** | **int32** | The attempt number of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListJobsForWorkflowRunAttemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListJobsForWorkflowRunAttempt200Response**](ActionsListJobsForWorkflowRunAttempt200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListLabelsForSelfHostedRunnerForOrg

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response ActionsListLabelsForSelfHostedRunnerForOrg(ctx, org, runnerId).Execute()

List labels for a self-hosted runner for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListLabelsForSelfHostedRunnerForOrg(context.Background(), org, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListLabelsForSelfHostedRunnerForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListLabelsForSelfHostedRunnerForOrg`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListLabelsForSelfHostedRunnerForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListLabelsForSelfHostedRunnerForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListLabelsForSelfHostedRunnerForRepo

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response ActionsListLabelsForSelfHostedRunnerForRepo(ctx, owner, repo, runnerId).Execute()

List labels for a self-hosted runner for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListLabelsForSelfHostedRunnerForRepo(context.Background(), owner, repo, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListLabelsForSelfHostedRunnerForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListLabelsForSelfHostedRunnerForRepo`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListLabelsForSelfHostedRunnerForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListLabelsForSelfHostedRunnerForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListOrgSecrets

> ActionsListOrgSecrets200Response ActionsListOrgSecrets(ctx, org).PerPage(perPage).Page(page).Execute()

List organization secrets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListOrgSecrets(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListOrgSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListOrgSecrets`: ActionsListOrgSecrets200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListOrgSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListOrgSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListOrgSecrets200Response**](ActionsListOrgSecrets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListRepoAccessToSelfHostedRunnerGroupInOrg

> ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response ActionsListRepoAccessToSelfHostedRunnerGroupInOrg(ctx, org, runnerGroupId).Page(page).PerPage(perPage).Execute()

List repository access to a self-hosted runner group in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListRepoAccessToSelfHostedRunnerGroupInOrg(context.Background(), org, runnerGroupId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListRepoAccessToSelfHostedRunnerGroupInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListRepoAccessToSelfHostedRunnerGroupInOrg`: ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListRepoAccessToSelfHostedRunnerGroupInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListRepoAccessToSelfHostedRunnerGroupInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

### Return type

[**ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response**](ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListRepoSecrets

> ActionsListRepoSecrets200Response ActionsListRepoSecrets(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List repository secrets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListRepoSecrets(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListRepoSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListRepoSecrets`: ActionsListRepoSecrets200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListRepoSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListRepoSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListRepoSecrets200Response**](ActionsListRepoSecrets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListRepoWorkflows

> ActionsListRepoWorkflows200Response ActionsListRepoWorkflows(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List repository workflows



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListRepoWorkflows(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListRepoWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListRepoWorkflows`: ActionsListRepoWorkflows200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListRepoWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListRepoWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListRepoWorkflows200Response**](ActionsListRepoWorkflows200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListRunnerApplicationsForOrg

> []RunnerApplication ActionsListRunnerApplicationsForOrg(ctx, org).Execute()

List runner applications for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListRunnerApplicationsForOrg(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListRunnerApplicationsForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListRunnerApplicationsForOrg`: []RunnerApplication
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListRunnerApplicationsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListRunnerApplicationsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RunnerApplication**](RunnerApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListRunnerApplicationsForRepo

> []RunnerApplication ActionsListRunnerApplicationsForRepo(ctx, owner, repo).Execute()

List runner applications for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListRunnerApplicationsForRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListRunnerApplicationsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListRunnerApplicationsForRepo`: []RunnerApplication
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListRunnerApplicationsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListRunnerApplicationsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RunnerApplication**](RunnerApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListSelectedReposForOrgSecret

> ActionsListSelectedReposForOrgSecret200Response ActionsListSelectedReposForOrgSecret(ctx, org, secretName).Page(page).PerPage(perPage).Execute()

List selected repositories for an organization secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    secretName := "secretName_example" // string | The name of the secret.
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListSelectedReposForOrgSecret(context.Background(), org, secretName).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListSelectedReposForOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListSelectedReposForOrgSecret`: ActionsListSelectedReposForOrgSecret200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListSelectedReposForOrgSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListSelectedReposForOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

### Return type

[**ActionsListSelectedReposForOrgSecret200Response**](ActionsListSelectedReposForOrgSecret200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListSelectedRepositoriesEnabledGithubActionsOrganization

> ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response ActionsListSelectedRepositoriesEnabledGithubActionsOrganization(ctx, org).PerPage(perPage).Page(page).Execute()

List selected repositories enabled for GitHub Actions in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListSelectedRepositoriesEnabledGithubActionsOrganization(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListSelectedRepositoriesEnabledGithubActionsOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListSelectedRepositoriesEnabledGithubActionsOrganization`: ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListSelectedRepositoriesEnabledGithubActionsOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListSelectedRepositoriesEnabledGithubActionsOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response**](ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListSelfHostedRunnerGroupsForOrg

> ActionsListSelfHostedRunnerGroupsForOrg200Response ActionsListSelfHostedRunnerGroupsForOrg(ctx, org).PerPage(perPage).Page(page).VisibleToRepository(visibleToRepository).Execute()

List self-hosted runner groups for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    visibleToRepository := "visibleToRepository_example" // string | Only return runner groups that are allowed to be used by this repository. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListSelfHostedRunnerGroupsForOrg(context.Background(), org).PerPage(perPage).Page(page).VisibleToRepository(visibleToRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListSelfHostedRunnerGroupsForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListSelfHostedRunnerGroupsForOrg`: ActionsListSelfHostedRunnerGroupsForOrg200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListSelfHostedRunnerGroupsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListSelfHostedRunnerGroupsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **visibleToRepository** | **string** | Only return runner groups that are allowed to be used by this repository. | 

### Return type

[**ActionsListSelfHostedRunnerGroupsForOrg200Response**](ActionsListSelfHostedRunnerGroupsForOrg200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListSelfHostedRunnersForOrg

> ActionsListSelfHostedRunnersForOrg200Response ActionsListSelfHostedRunnersForOrg(ctx, org).PerPage(perPage).Page(page).Execute()

List self-hosted runners for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListSelfHostedRunnersForOrg(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListSelfHostedRunnersForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListSelfHostedRunnersForOrg`: ActionsListSelfHostedRunnersForOrg200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListSelfHostedRunnersForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListSelfHostedRunnersForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListSelfHostedRunnersForOrg200Response**](ActionsListSelfHostedRunnersForOrg200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListSelfHostedRunnersForRepo

> ActionsListSelfHostedRunnersForOrg200Response ActionsListSelfHostedRunnersForRepo(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List self-hosted runners for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListSelfHostedRunnersForRepo(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListSelfHostedRunnersForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListSelfHostedRunnersForRepo`: ActionsListSelfHostedRunnersForOrg200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListSelfHostedRunnersForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListSelfHostedRunnersForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListSelfHostedRunnersForOrg200Response**](ActionsListSelfHostedRunnersForOrg200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListSelfHostedRunnersInGroupForOrg

> EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response ActionsListSelfHostedRunnersInGroupForOrg(ctx, org, runnerGroupId).PerPage(perPage).Page(page).Execute()

List self-hosted runners in a group for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListSelfHostedRunnersInGroupForOrg(context.Background(), org, runnerGroupId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListSelfHostedRunnersInGroupForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListSelfHostedRunnersInGroupForOrg`: EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListSelfHostedRunnersInGroupForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListSelfHostedRunnersInGroupForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response**](EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListWorkflowRunArtifacts

> ActionsListArtifactsForRepo200Response ActionsListWorkflowRunArtifacts(ctx, owner, repo, runId).PerPage(perPage).Page(page).Execute()

List workflow run artifacts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListWorkflowRunArtifacts(context.Background(), owner, repo, runId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListWorkflowRunArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListWorkflowRunArtifacts`: ActionsListArtifactsForRepo200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListWorkflowRunArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListWorkflowRunArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActionsListArtifactsForRepo200Response**](ActionsListArtifactsForRepo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListWorkflowRuns

> ActionsListWorkflowRunsForRepo200Response ActionsListWorkflowRuns(ctx, owner, repo, workflowId).Actor(actor).Branch(branch).Event(event).Status(status).PerPage(perPage).Page(page).Created(created).ExcludePullRequests(excludePullRequests).CheckSuiteId(checkSuiteId).Execute()

List workflow runs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    workflowId := openapiclient.actions_get_workflow_workflow_id_parameter{Int32: new(int32)} // ActionsGetWorkflowWorkflowIdParameter | The ID of the workflow. You can also pass the workflow file name as a string.
    actor := "actor_example" // string | Returns someone's workflow runs. Use the login for the user who created the `push` associated with the check suite or workflow run. (optional)
    branch := "branch_example" // string | Returns workflow runs associated with a branch. Use the name of the branch of the `push`. (optional)
    event := "event_example" // string | Returns workflow run triggered by the event you specify. For example, `push`, `pull_request` or `issue`. For more information, see \"[Events that trigger workflows](https://docs.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows).\" (optional)
    status := "status_example" // string | Returns workflow runs with the check run `status` or `conclusion` that you specify. For example, a conclusion can be `success` or a status can be `in_progress`. Only GitHub can set a status of `waiting` or `requested`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    created := time.Now() // time.Time | Returns workflow runs created within the given date-time range. For more information on the syntax, see \"[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates).\" (optional)
    excludePullRequests := true // bool | If `true` pull requests are omitted from the response (empty array). (optional) (default to false)
    checkSuiteId := int32(56) // int32 | Returns workflow runs with the `check_suite_id` that you specify. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListWorkflowRuns(context.Background(), owner, repo, workflowId).Actor(actor).Branch(branch).Event(event).Status(status).PerPage(perPage).Page(page).Created(created).ExcludePullRequests(excludePullRequests).CheckSuiteId(checkSuiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListWorkflowRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListWorkflowRuns`: ActionsListWorkflowRunsForRepo200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListWorkflowRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**workflowId** | [**ActionsGetWorkflowWorkflowIdParameter**](.md) | The ID of the workflow. You can also pass the workflow file name as a string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListWorkflowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **actor** | **string** | Returns someone&#39;s workflow runs. Use the login for the user who created the &#x60;push&#x60; associated with the check suite or workflow run. | 
 **branch** | **string** | Returns workflow runs associated with a branch. Use the name of the branch of the &#x60;push&#x60;. | 
 **event** | **string** | Returns workflow run triggered by the event you specify. For example, &#x60;push&#x60;, &#x60;pull_request&#x60; or &#x60;issue&#x60;. For more information, see \&quot;[Events that trigger workflows](https://docs.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows).\&quot; | 
 **status** | **string** | Returns workflow runs with the check run &#x60;status&#x60; or &#x60;conclusion&#x60; that you specify. For example, a conclusion can be &#x60;success&#x60; or a status can be &#x60;in_progress&#x60;. Only GitHub can set a status of &#x60;waiting&#x60; or &#x60;requested&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **created** | **time.Time** | Returns workflow runs created within the given date-time range. For more information on the syntax, see \&quot;[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates).\&quot; | 
 **excludePullRequests** | **bool** | If &#x60;true&#x60; pull requests are omitted from the response (empty array). | [default to false]
 **checkSuiteId** | **int32** | Returns workflow runs with the &#x60;check_suite_id&#x60; that you specify. | 

### Return type

[**ActionsListWorkflowRunsForRepo200Response**](ActionsListWorkflowRunsForRepo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsListWorkflowRunsForRepo

> ActionsListWorkflowRunsForRepo200Response ActionsListWorkflowRunsForRepo(ctx, owner, repo).Actor(actor).Branch(branch).Event(event).Status(status).PerPage(perPage).Page(page).Created(created).ExcludePullRequests(excludePullRequests).CheckSuiteId(checkSuiteId).Execute()

List workflow runs for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    actor := "actor_example" // string | Returns someone's workflow runs. Use the login for the user who created the `push` associated with the check suite or workflow run. (optional)
    branch := "branch_example" // string | Returns workflow runs associated with a branch. Use the name of the branch of the `push`. (optional)
    event := "event_example" // string | Returns workflow run triggered by the event you specify. For example, `push`, `pull_request` or `issue`. For more information, see \"[Events that trigger workflows](https://docs.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows).\" (optional)
    status := "status_example" // string | Returns workflow runs with the check run `status` or `conclusion` that you specify. For example, a conclusion can be `success` or a status can be `in_progress`. Only GitHub can set a status of `waiting` or `requested`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    created := time.Now() // time.Time | Returns workflow runs created within the given date-time range. For more information on the syntax, see \"[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates).\" (optional)
    excludePullRequests := true // bool | If `true` pull requests are omitted from the response (empty array). (optional) (default to false)
    checkSuiteId := int32(56) // int32 | Returns workflow runs with the `check_suite_id` that you specify. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsListWorkflowRunsForRepo(context.Background(), owner, repo).Actor(actor).Branch(branch).Event(event).Status(status).PerPage(perPage).Page(page).Created(created).ExcludePullRequests(excludePullRequests).CheckSuiteId(checkSuiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsListWorkflowRunsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsListWorkflowRunsForRepo`: ActionsListWorkflowRunsForRepo200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsListWorkflowRunsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListWorkflowRunsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actor** | **string** | Returns someone&#39;s workflow runs. Use the login for the user who created the &#x60;push&#x60; associated with the check suite or workflow run. | 
 **branch** | **string** | Returns workflow runs associated with a branch. Use the name of the branch of the &#x60;push&#x60;. | 
 **event** | **string** | Returns workflow run triggered by the event you specify. For example, &#x60;push&#x60;, &#x60;pull_request&#x60; or &#x60;issue&#x60;. For more information, see \&quot;[Events that trigger workflows](https://docs.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows).\&quot; | 
 **status** | **string** | Returns workflow runs with the check run &#x60;status&#x60; or &#x60;conclusion&#x60; that you specify. For example, a conclusion can be &#x60;success&#x60; or a status can be &#x60;in_progress&#x60;. Only GitHub can set a status of &#x60;waiting&#x60; or &#x60;requested&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **created** | **time.Time** | Returns workflow runs created within the given date-time range. For more information on the syntax, see \&quot;[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates).\&quot; | 
 **excludePullRequests** | **bool** | If &#x60;true&#x60; pull requests are omitted from the response (empty array). | [default to false]
 **checkSuiteId** | **int32** | Returns workflow runs with the &#x60;check_suite_id&#x60; that you specify. | 

### Return type

[**ActionsListWorkflowRunsForRepo200Response**](ActionsListWorkflowRunsForRepo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsReRunJobForWorkflowRun

> map[string]interface{} ActionsReRunJobForWorkflowRun(ctx, owner, repo, jobId).ActionsReRunJobForWorkflowRunRequest(actionsReRunJobForWorkflowRunRequest).Execute()

Re-run a job from a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    jobId := int32(56) // int32 | The unique identifier of the job.
    actionsReRunJobForWorkflowRunRequest := *openapiclient.NewActionsReRunJobForWorkflowRunRequest() // ActionsReRunJobForWorkflowRunRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsReRunJobForWorkflowRun(context.Background(), owner, repo, jobId).ActionsReRunJobForWorkflowRunRequest(actionsReRunJobForWorkflowRunRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsReRunJobForWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsReRunJobForWorkflowRun`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsReRunJobForWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**jobId** | **int32** | The unique identifier of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsReRunJobForWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **actionsReRunJobForWorkflowRunRequest** | [**ActionsReRunJobForWorkflowRunRequest**](ActionsReRunJobForWorkflowRunRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsReRunWorkflow

> map[string]interface{} ActionsReRunWorkflow(ctx, owner, repo, runId).ActionsReRunJobForWorkflowRunRequest(actionsReRunJobForWorkflowRunRequest).Execute()

Re-run a workflow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.
    actionsReRunJobForWorkflowRunRequest := *openapiclient.NewActionsReRunJobForWorkflowRunRequest() // ActionsReRunJobForWorkflowRunRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsReRunWorkflow(context.Background(), owner, repo, runId).ActionsReRunJobForWorkflowRunRequest(actionsReRunJobForWorkflowRunRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsReRunWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsReRunWorkflow`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsReRunWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsReRunWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **actionsReRunJobForWorkflowRunRequest** | [**ActionsReRunJobForWorkflowRunRequest**](ActionsReRunJobForWorkflowRunRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsReRunWorkflowFailedJobs

> map[string]interface{} ActionsReRunWorkflowFailedJobs(ctx, owner, repo, runId).ActionsReRunJobForWorkflowRunRequest(actionsReRunJobForWorkflowRunRequest).Execute()

Re-run failed jobs from a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.
    actionsReRunJobForWorkflowRunRequest := *openapiclient.NewActionsReRunJobForWorkflowRunRequest() // ActionsReRunJobForWorkflowRunRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsReRunWorkflowFailedJobs(context.Background(), owner, repo, runId).ActionsReRunJobForWorkflowRunRequest(actionsReRunJobForWorkflowRunRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsReRunWorkflowFailedJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsReRunWorkflowFailedJobs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsReRunWorkflowFailedJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsReRunWorkflowFailedJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **actionsReRunJobForWorkflowRunRequest** | [**ActionsReRunJobForWorkflowRunRequest**](ActionsReRunJobForWorkflowRunRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForOrg

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForOrg(ctx, org, runnerId).Execute()

Remove all custom labels from a self-hosted runner for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForOrg(context.Background(), org, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForOrg`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRemoveAllCustomLabelsFromSelfHostedRunnerForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForRepo

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForRepo(ctx, owner, repo, runnerId).Execute()

Remove all custom labels from a self-hosted runner for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForRepo(context.Background(), owner, repo, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForRepo`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsRemoveAllCustomLabelsFromSelfHostedRunnerForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRemoveAllCustomLabelsFromSelfHostedRunnerForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRemoveCustomLabelFromSelfHostedRunnerForOrg

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response ActionsRemoveCustomLabelFromSelfHostedRunnerForOrg(ctx, org, runnerId, name).Execute()

Remove a custom label from a self-hosted runner for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.
    name := "name_example" // string | The name of a self-hosted runner's custom label.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsRemoveCustomLabelFromSelfHostedRunnerForOrg(context.Background(), org, runnerId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsRemoveCustomLabelFromSelfHostedRunnerForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsRemoveCustomLabelFromSelfHostedRunnerForOrg`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsRemoveCustomLabelFromSelfHostedRunnerForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 
**name** | **string** | The name of a self-hosted runner&#39;s custom label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRemoveCustomLabelFromSelfHostedRunnerForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRemoveCustomLabelFromSelfHostedRunnerForRepo

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response ActionsRemoveCustomLabelFromSelfHostedRunnerForRepo(ctx, owner, repo, runnerId, name).Execute()

Remove a custom label from a self-hosted runner for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.
    name := "name_example" // string | The name of a self-hosted runner's custom label.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsRemoveCustomLabelFromSelfHostedRunnerForRepo(context.Background(), owner, repo, runnerId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsRemoveCustomLabelFromSelfHostedRunnerForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsRemoveCustomLabelFromSelfHostedRunnerForRepo`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsRemoveCustomLabelFromSelfHostedRunnerForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 
**name** | **string** | The name of a self-hosted runner&#39;s custom label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRemoveCustomLabelFromSelfHostedRunnerForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRemoveRepoAccessToSelfHostedRunnerGroupInOrg

> ActionsRemoveRepoAccessToSelfHostedRunnerGroupInOrg(ctx, org, runnerGroupId, repositoryId).Execute()

Remove repository access to a self-hosted runner group in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    repositoryId := int32(56) // int32 | The unique identifier of the repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsRemoveRepoAccessToSelfHostedRunnerGroupInOrg(context.Background(), org, runnerGroupId, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsRemoveRepoAccessToSelfHostedRunnerGroupInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 
**repositoryId** | **int32** | The unique identifier of the repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRemoveRepoAccessToSelfHostedRunnerGroupInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRemoveSelectedRepoFromOrgSecret

> ActionsRemoveSelectedRepoFromOrgSecret(ctx, org, secretName, repositoryId).Execute()

Remove selected repository from an organization secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    secretName := "secretName_example" // string | The name of the secret.
    repositoryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsRemoveSelectedRepoFromOrgSecret(context.Background(), org, secretName, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsRemoveSelectedRepoFromOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 
**repositoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRemoveSelectedRepoFromOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRemoveSelfHostedRunnerFromGroupForOrg

> ActionsRemoveSelfHostedRunnerFromGroupForOrg(ctx, org, runnerGroupId, runnerId).Execute()

Remove a self-hosted runner from a group for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsRemoveSelfHostedRunnerFromGroupForOrg(context.Background(), org, runnerGroupId, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsRemoveSelfHostedRunnerFromGroupForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRemoveSelfHostedRunnerFromGroupForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsReviewPendingDeploymentsForRun

> []Deployment ActionsReviewPendingDeploymentsForRun(ctx, owner, repo, runId).ActionsReviewPendingDeploymentsForRunRequest(actionsReviewPendingDeploymentsForRunRequest).Execute()

Review pending deployments for a workflow run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runId := int32(56) // int32 | The unique identifier of the workflow run.
    actionsReviewPendingDeploymentsForRunRequest := *openapiclient.NewActionsReviewPendingDeploymentsForRunRequest([]int32{int32(161171787)}, "approved", "Ship it!") // ActionsReviewPendingDeploymentsForRunRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsReviewPendingDeploymentsForRun(context.Background(), owner, repo, runId).ActionsReviewPendingDeploymentsForRunRequest(actionsReviewPendingDeploymentsForRunRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsReviewPendingDeploymentsForRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsReviewPendingDeploymentsForRun`: []Deployment
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsReviewPendingDeploymentsForRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runId** | **int32** | The unique identifier of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsReviewPendingDeploymentsForRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **actionsReviewPendingDeploymentsForRunRequest** | [**ActionsReviewPendingDeploymentsForRunRequest**](ActionsReviewPendingDeploymentsForRunRequest.md) |  | 

### Return type

[**[]Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetActionsOidcCustomIssuerPolicyForEnterprise

> ActionsSetActionsOidcCustomIssuerPolicyForEnterprise(ctx, enterprise).ActionsOidcCustomIssuerPolicyForEnterprise(actionsOidcCustomIssuerPolicyForEnterprise).Execute()

Set the GitHub Actions OIDC custom issuer policy for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    actionsOidcCustomIssuerPolicyForEnterprise := *openapiclient.NewActionsOidcCustomIssuerPolicyForEnterprise() // ActionsOidcCustomIssuerPolicyForEnterprise | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetActionsOidcCustomIssuerPolicyForEnterprise(context.Background(), enterprise).ActionsOidcCustomIssuerPolicyForEnterprise(actionsOidcCustomIssuerPolicyForEnterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetActionsOidcCustomIssuerPolicyForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetActionsOidcCustomIssuerPolicyForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsOidcCustomIssuerPolicyForEnterprise** | [**ActionsOidcCustomIssuerPolicyForEnterprise**](ActionsOidcCustomIssuerPolicyForEnterprise.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetAllowedActionsOrganization

> ActionsSetAllowedActionsOrganization(ctx, org).SelectedActions(selectedActions).Execute()

Set allowed actions and reusable workflows for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    selectedActions := *openapiclient.NewSelectedActions() // SelectedActions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetAllowedActionsOrganization(context.Background(), org).SelectedActions(selectedActions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetAllowedActionsOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetAllowedActionsOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selectedActions** | [**SelectedActions**](SelectedActions.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetAllowedActionsRepository

> ActionsSetAllowedActionsRepository(ctx, owner, repo).SelectedActions(selectedActions).Execute()

Set allowed actions and reusable workflows for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    selectedActions := *openapiclient.NewSelectedActions() // SelectedActions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetAllowedActionsRepository(context.Background(), owner, repo).SelectedActions(selectedActions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetAllowedActionsRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetAllowedActionsRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectedActions** | [**SelectedActions**](SelectedActions.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetCustomLabelsForSelfHostedRunnerForOrg

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response ActionsSetCustomLabelsForSelfHostedRunnerForOrg(ctx, org, runnerId).EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest(enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest).Execute()

Set custom labels for a self-hosted runner for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.
    enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest := *openapiclient.NewEnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest([]string{"Labels_example"}) // EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetCustomLabelsForSelfHostedRunnerForOrg(context.Background(), org, runnerId).EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest(enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetCustomLabelsForSelfHostedRunnerForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsSetCustomLabelsForSelfHostedRunnerForOrg`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsSetCustomLabelsForSelfHostedRunnerForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetCustomLabelsForSelfHostedRunnerForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest** | [**EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest**](EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest.md) |  | 

### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetCustomLabelsForSelfHostedRunnerForRepo

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response ActionsSetCustomLabelsForSelfHostedRunnerForRepo(ctx, owner, repo, runnerId).EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest(enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest).Execute()

Set custom labels for a self-hosted runner for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.
    enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest := *openapiclient.NewEnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest([]string{"Labels_example"}) // EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetCustomLabelsForSelfHostedRunnerForRepo(context.Background(), owner, repo, runnerId).EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest(enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetCustomLabelsForSelfHostedRunnerForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsSetCustomLabelsForSelfHostedRunnerForRepo`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsSetCustomLabelsForSelfHostedRunnerForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetCustomLabelsForSelfHostedRunnerForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest** | [**EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest**](EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest.md) |  | 

### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetCustomOidcSubClaimForRepo

> map[string]interface{} ActionsSetCustomOidcSubClaimForRepo(ctx, owner, repo).OptOutOidcCustomSub(optOutOidcCustomSub).Execute()

Set the opt-in flag of an OIDC subject claim customization for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    optOutOidcCustomSub := *openapiclient.NewOptOutOidcCustomSub(false) // OptOutOidcCustomSub | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetCustomOidcSubClaimForRepo(context.Background(), owner, repo).OptOutOidcCustomSub(optOutOidcCustomSub).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetCustomOidcSubClaimForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsSetCustomOidcSubClaimForRepo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsSetCustomOidcSubClaimForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetCustomOidcSubClaimForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **optOutOidcCustomSub** | [**OptOutOidcCustomSub**](OptOutOidcCustomSub.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetGithubActionsDefaultWorkflowPermissionsEnterprise

> ActionsSetGithubActionsDefaultWorkflowPermissionsEnterprise(ctx, enterprise).ActionsSetDefaultWorkflowPermissions(actionsSetDefaultWorkflowPermissions).Execute()

Set default workflow permissions for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    actionsSetDefaultWorkflowPermissions := *openapiclient.NewActionsSetDefaultWorkflowPermissions() // ActionsSetDefaultWorkflowPermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetGithubActionsDefaultWorkflowPermissionsEnterprise(context.Background(), enterprise).ActionsSetDefaultWorkflowPermissions(actionsSetDefaultWorkflowPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetGithubActionsDefaultWorkflowPermissionsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetGithubActionsDefaultWorkflowPermissionsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsSetDefaultWorkflowPermissions** | [**ActionsSetDefaultWorkflowPermissions**](ActionsSetDefaultWorkflowPermissions.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetGithubActionsDefaultWorkflowPermissionsOrganization

> ActionsSetGithubActionsDefaultWorkflowPermissionsOrganization(ctx, org).ActionsSetDefaultWorkflowPermissions(actionsSetDefaultWorkflowPermissions).Execute()

Set default workflow permissions for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    actionsSetDefaultWorkflowPermissions := *openapiclient.NewActionsSetDefaultWorkflowPermissions() // ActionsSetDefaultWorkflowPermissions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetGithubActionsDefaultWorkflowPermissionsOrganization(context.Background(), org).ActionsSetDefaultWorkflowPermissions(actionsSetDefaultWorkflowPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetGithubActionsDefaultWorkflowPermissionsOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetGithubActionsDefaultWorkflowPermissionsOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsSetDefaultWorkflowPermissions** | [**ActionsSetDefaultWorkflowPermissions**](ActionsSetDefaultWorkflowPermissions.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetGithubActionsDefaultWorkflowPermissionsRepository

> ActionsSetGithubActionsDefaultWorkflowPermissionsRepository(ctx, owner, repo).ActionsSetDefaultWorkflowPermissions(actionsSetDefaultWorkflowPermissions).Execute()

Set default workflow permissions for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    actionsSetDefaultWorkflowPermissions := *openapiclient.NewActionsSetDefaultWorkflowPermissions() // ActionsSetDefaultWorkflowPermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetGithubActionsDefaultWorkflowPermissionsRepository(context.Background(), owner, repo).ActionsSetDefaultWorkflowPermissions(actionsSetDefaultWorkflowPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetGithubActionsDefaultWorkflowPermissionsRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetGithubActionsDefaultWorkflowPermissionsRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actionsSetDefaultWorkflowPermissions** | [**ActionsSetDefaultWorkflowPermissions**](ActionsSetDefaultWorkflowPermissions.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetGithubActionsPermissionsOrganization

> ActionsSetGithubActionsPermissionsOrganization(ctx, org).ActionsSetGithubActionsPermissionsOrganizationRequest(actionsSetGithubActionsPermissionsOrganizationRequest).Execute()

Set GitHub Actions permissions for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    actionsSetGithubActionsPermissionsOrganizationRequest := *openapiclient.NewActionsSetGithubActionsPermissionsOrganizationRequest(openapiclient.enabled-repositories("all")) // ActionsSetGithubActionsPermissionsOrganizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetGithubActionsPermissionsOrganization(context.Background(), org).ActionsSetGithubActionsPermissionsOrganizationRequest(actionsSetGithubActionsPermissionsOrganizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetGithubActionsPermissionsOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetGithubActionsPermissionsOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsSetGithubActionsPermissionsOrganizationRequest** | [**ActionsSetGithubActionsPermissionsOrganizationRequest**](ActionsSetGithubActionsPermissionsOrganizationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetGithubActionsPermissionsRepository

> ActionsSetGithubActionsPermissionsRepository(ctx, owner, repo).ActionsSetGithubActionsPermissionsRepositoryRequest(actionsSetGithubActionsPermissionsRepositoryRequest).Execute()

Set GitHub Actions permissions for a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    actionsSetGithubActionsPermissionsRepositoryRequest := *openapiclient.NewActionsSetGithubActionsPermissionsRepositoryRequest(false) // ActionsSetGithubActionsPermissionsRepositoryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetGithubActionsPermissionsRepository(context.Background(), owner, repo).ActionsSetGithubActionsPermissionsRepositoryRequest(actionsSetGithubActionsPermissionsRepositoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetGithubActionsPermissionsRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetGithubActionsPermissionsRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actionsSetGithubActionsPermissionsRepositoryRequest** | [**ActionsSetGithubActionsPermissionsRepositoryRequest**](ActionsSetGithubActionsPermissionsRepositoryRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetRepoAccessToSelfHostedRunnerGroupInOrg

> ActionsSetRepoAccessToSelfHostedRunnerGroupInOrg(ctx, org, runnerGroupId).ActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest(actionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest).Execute()

Set repository access for a self-hosted runner group in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    actionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest := *openapiclient.NewActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest([]int32{int32(123)}) // ActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetRepoAccessToSelfHostedRunnerGroupInOrg(context.Background(), org, runnerGroupId).ActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest(actionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetRepoAccessToSelfHostedRunnerGroupInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest** | [**ActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest**](ActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetSelectedReposForOrgSecret

> ActionsSetSelectedReposForOrgSecret(ctx, org, secretName).ActionsSetSelectedReposForOrgSecretRequest(actionsSetSelectedReposForOrgSecretRequest).Execute()

Set selected repositories for an organization secret



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    secretName := "secretName_example" // string | The name of the secret.
    actionsSetSelectedReposForOrgSecretRequest := *openapiclient.NewActionsSetSelectedReposForOrgSecretRequest([]int32{int32(123)}) // ActionsSetSelectedReposForOrgSecretRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetSelectedReposForOrgSecret(context.Background(), org, secretName).ActionsSetSelectedReposForOrgSecretRequest(actionsSetSelectedReposForOrgSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetSelectedReposForOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetSelectedReposForOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actionsSetSelectedReposForOrgSecretRequest** | [**ActionsSetSelectedReposForOrgSecretRequest**](ActionsSetSelectedReposForOrgSecretRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetSelectedRepositoriesEnabledGithubActionsOrganization

> ActionsSetSelectedRepositoriesEnabledGithubActionsOrganization(ctx, org).ActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest(actionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest).Execute()

Set selected repositories enabled for GitHub Actions in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    actionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest := *openapiclient.NewActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest([]int32{int32(123)}) // ActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetSelectedRepositoriesEnabledGithubActionsOrganization(context.Background(), org).ActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest(actionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetSelectedRepositoriesEnabledGithubActionsOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest** | [**ActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest**](ActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetSelfHostedRunnersInGroupForOrg

> ActionsSetSelfHostedRunnersInGroupForOrg(ctx, org, runnerGroupId).EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest(enterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest).Execute()

Set self-hosted runners in a group for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    enterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest := *openapiclient.NewEnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest([]int32{int32(123)}) // EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetSelfHostedRunnersInGroupForOrg(context.Background(), org, runnerGroupId).EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest(enterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetSelfHostedRunnersInGroupForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetSelfHostedRunnersInGroupForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest** | [**EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest**](EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsSetWorkflowAccessToRepository

> ActionsSetWorkflowAccessToRepository(ctx, owner, repo).ActionsWorkflowAccessToRepository(actionsWorkflowAccessToRepository).Execute()

Set the level of access for workflows outside of the repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    actionsWorkflowAccessToRepository := *openapiclient.NewActionsWorkflowAccessToRepository("AccessLevel_example") // ActionsWorkflowAccessToRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsSetWorkflowAccessToRepository(context.Background(), owner, repo).ActionsWorkflowAccessToRepository(actionsWorkflowAccessToRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsSetWorkflowAccessToRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsSetWorkflowAccessToRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actionsWorkflowAccessToRepository** | [**ActionsWorkflowAccessToRepository**](ActionsWorkflowAccessToRepository.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsUpdateSelfHostedRunnerGroupForOrg

> RunnerGroupsOrg ActionsUpdateSelfHostedRunnerGroupForOrg(ctx, org, runnerGroupId).ActionsUpdateSelfHostedRunnerGroupForOrgRequest(actionsUpdateSelfHostedRunnerGroupForOrgRequest).Execute()

Update a self-hosted runner group for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    org := "org_example" // string | The organization name. The name is not case sensitive.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    actionsUpdateSelfHostedRunnerGroupForOrgRequest := *openapiclient.NewActionsUpdateSelfHostedRunnerGroupForOrgRequest("Name_example") // ActionsUpdateSelfHostedRunnerGroupForOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionsApi.ActionsUpdateSelfHostedRunnerGroupForOrg(context.Background(), org, runnerGroupId).ActionsUpdateSelfHostedRunnerGroupForOrgRequest(actionsUpdateSelfHostedRunnerGroupForOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.ActionsUpdateSelfHostedRunnerGroupForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionsUpdateSelfHostedRunnerGroupForOrg`: RunnerGroupsOrg
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.ActionsUpdateSelfHostedRunnerGroupForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsUpdateSelfHostedRunnerGroupForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actionsUpdateSelfHostedRunnerGroupForOrgRequest** | [**ActionsUpdateSelfHostedRunnerGroupForOrgRequest**](ActionsUpdateSelfHostedRunnerGroupForOrgRequest.md) |  | 

### Return type

[**RunnerGroupsOrg**](RunnerGroupsOrg.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

