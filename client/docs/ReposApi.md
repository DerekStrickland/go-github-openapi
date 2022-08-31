# \ReposApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReposAcceptInvitationForAuthenticatedUser**](ReposApi.md#ReposAcceptInvitationForAuthenticatedUser) | **Patch** /user/repository_invitations/{invitation_id} | Accept a repository invitation
[**ReposAddAppAccessRestrictions**](ReposApi.md#ReposAddAppAccessRestrictions) | **Post** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps | Add app access restrictions
[**ReposAddCollaborator**](ReposApi.md#ReposAddCollaborator) | **Put** /repos/{owner}/{repo}/collaborators/{username} | Add a repository collaborator
[**ReposAddStatusCheckContexts**](ReposApi.md#ReposAddStatusCheckContexts) | **Post** /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts | Add status check contexts
[**ReposAddTeamAccessRestrictions**](ReposApi.md#ReposAddTeamAccessRestrictions) | **Post** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams | Add team access restrictions
[**ReposAddUserAccessRestrictions**](ReposApi.md#ReposAddUserAccessRestrictions) | **Post** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users | Add user access restrictions
[**ReposCheckCollaborator**](ReposApi.md#ReposCheckCollaborator) | **Get** /repos/{owner}/{repo}/collaborators/{username} | Check if a user is a repository collaborator
[**ReposCheckVulnerabilityAlerts**](ReposApi.md#ReposCheckVulnerabilityAlerts) | **Get** /repos/{owner}/{repo}/vulnerability-alerts | Check if vulnerability alerts are enabled for a repository
[**ReposCodeownersErrors**](ReposApi.md#ReposCodeownersErrors) | **Get** /repos/{owner}/{repo}/codeowners/errors | List CODEOWNERS errors
[**ReposCompareCommits**](ReposApi.md#ReposCompareCommits) | **Get** /repos/{owner}/{repo}/compare/{basehead} | Compare two commits
[**ReposCreateAutolink**](ReposApi.md#ReposCreateAutolink) | **Post** /repos/{owner}/{repo}/autolinks | Create an autolink reference for a repository
[**ReposCreateCommitComment**](ReposApi.md#ReposCreateCommitComment) | **Post** /repos/{owner}/{repo}/commits/{commit_sha}/comments | Create a commit comment
[**ReposCreateCommitSignatureProtection**](ReposApi.md#ReposCreateCommitSignatureProtection) | **Post** /repos/{owner}/{repo}/branches/{branch}/protection/required_signatures | Create commit signature protection
[**ReposCreateCommitStatus**](ReposApi.md#ReposCreateCommitStatus) | **Post** /repos/{owner}/{repo}/statuses/{sha} | Create a commit status
[**ReposCreateDeployKey**](ReposApi.md#ReposCreateDeployKey) | **Post** /repos/{owner}/{repo}/keys | Create a deploy key
[**ReposCreateDeployment**](ReposApi.md#ReposCreateDeployment) | **Post** /repos/{owner}/{repo}/deployments | Create a deployment
[**ReposCreateDeploymentBranchPolicy**](ReposApi.md#ReposCreateDeploymentBranchPolicy) | **Post** /repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies | Create a deployment branch policy
[**ReposCreateDeploymentStatus**](ReposApi.md#ReposCreateDeploymentStatus) | **Post** /repos/{owner}/{repo}/deployments/{deployment_id}/statuses | Create a deployment status
[**ReposCreateDispatchEvent**](ReposApi.md#ReposCreateDispatchEvent) | **Post** /repos/{owner}/{repo}/dispatches | Create a repository dispatch event
[**ReposCreateForAuthenticatedUser**](ReposApi.md#ReposCreateForAuthenticatedUser) | **Post** /user/repos | Create a repository for the authenticated user
[**ReposCreateFork**](ReposApi.md#ReposCreateFork) | **Post** /repos/{owner}/{repo}/forks | Create a fork
[**ReposCreateInOrg**](ReposApi.md#ReposCreateInOrg) | **Post** /orgs/{org}/repos | Create an organization repository
[**ReposCreateOrUpdateEnvironment**](ReposApi.md#ReposCreateOrUpdateEnvironment) | **Put** /repos/{owner}/{repo}/environments/{environment_name} | Create or update an environment
[**ReposCreateOrUpdateFileContents**](ReposApi.md#ReposCreateOrUpdateFileContents) | **Put** /repos/{owner}/{repo}/contents/{path} | Create or update file contents
[**ReposCreatePagesDeployment**](ReposApi.md#ReposCreatePagesDeployment) | **Post** /repos/{owner}/{repo}/pages/deployment | Create a GitHub Pages deployment
[**ReposCreatePagesSite**](ReposApi.md#ReposCreatePagesSite) | **Post** /repos/{owner}/{repo}/pages | Create a GitHub Pages site
[**ReposCreateRelease**](ReposApi.md#ReposCreateRelease) | **Post** /repos/{owner}/{repo}/releases | Create a release
[**ReposCreateTagProtection**](ReposApi.md#ReposCreateTagProtection) | **Post** /repos/{owner}/{repo}/tags/protection | Create a tag protection state for a repository
[**ReposCreateUsingTemplate**](ReposApi.md#ReposCreateUsingTemplate) | **Post** /repos/{template_owner}/{template_repo}/generate | Create a repository using a template
[**ReposCreateWebhook**](ReposApi.md#ReposCreateWebhook) | **Post** /repos/{owner}/{repo}/hooks | Create a repository webhook
[**ReposDeclineInvitationForAuthenticatedUser**](ReposApi.md#ReposDeclineInvitationForAuthenticatedUser) | **Delete** /user/repository_invitations/{invitation_id} | Decline a repository invitation
[**ReposDelete**](ReposApi.md#ReposDelete) | **Delete** /repos/{owner}/{repo} | Delete a repository
[**ReposDeleteAccessRestrictions**](ReposApi.md#ReposDeleteAccessRestrictions) | **Delete** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions | Delete access restrictions
[**ReposDeleteAdminBranchProtection**](ReposApi.md#ReposDeleteAdminBranchProtection) | **Delete** /repos/{owner}/{repo}/branches/{branch}/protection/enforce_admins | Delete admin branch protection
[**ReposDeleteAnEnvironment**](ReposApi.md#ReposDeleteAnEnvironment) | **Delete** /repos/{owner}/{repo}/environments/{environment_name} | Delete an environment
[**ReposDeleteAutolink**](ReposApi.md#ReposDeleteAutolink) | **Delete** /repos/{owner}/{repo}/autolinks/{autolink_id} | Delete an autolink reference from a repository
[**ReposDeleteBranchProtection**](ReposApi.md#ReposDeleteBranchProtection) | **Delete** /repos/{owner}/{repo}/branches/{branch}/protection | Delete branch protection
[**ReposDeleteCommitComment**](ReposApi.md#ReposDeleteCommitComment) | **Delete** /repos/{owner}/{repo}/comments/{comment_id} | Delete a commit comment
[**ReposDeleteCommitSignatureProtection**](ReposApi.md#ReposDeleteCommitSignatureProtection) | **Delete** /repos/{owner}/{repo}/branches/{branch}/protection/required_signatures | Delete commit signature protection
[**ReposDeleteDeployKey**](ReposApi.md#ReposDeleteDeployKey) | **Delete** /repos/{owner}/{repo}/keys/{key_id} | Delete a deploy key
[**ReposDeleteDeployment**](ReposApi.md#ReposDeleteDeployment) | **Delete** /repos/{owner}/{repo}/deployments/{deployment_id} | Delete a deployment
[**ReposDeleteDeploymentBranchPolicy**](ReposApi.md#ReposDeleteDeploymentBranchPolicy) | **Delete** /repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies/{branch_policy_id} | Delete a deployment branch policy
[**ReposDeleteFile**](ReposApi.md#ReposDeleteFile) | **Delete** /repos/{owner}/{repo}/contents/{path} | Delete a file
[**ReposDeleteInvitation**](ReposApi.md#ReposDeleteInvitation) | **Delete** /repos/{owner}/{repo}/invitations/{invitation_id} | Delete a repository invitation
[**ReposDeletePagesSite**](ReposApi.md#ReposDeletePagesSite) | **Delete** /repos/{owner}/{repo}/pages | Delete a GitHub Pages site
[**ReposDeletePullRequestReviewProtection**](ReposApi.md#ReposDeletePullRequestReviewProtection) | **Delete** /repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews | Delete pull request review protection
[**ReposDeleteRelease**](ReposApi.md#ReposDeleteRelease) | **Delete** /repos/{owner}/{repo}/releases/{release_id} | Delete a release
[**ReposDeleteReleaseAsset**](ReposApi.md#ReposDeleteReleaseAsset) | **Delete** /repos/{owner}/{repo}/releases/assets/{asset_id} | Delete a release asset
[**ReposDeleteTagProtection**](ReposApi.md#ReposDeleteTagProtection) | **Delete** /repos/{owner}/{repo}/tags/protection/{tag_protection_id} | Delete a tag protection state for a repository
[**ReposDeleteWebhook**](ReposApi.md#ReposDeleteWebhook) | **Delete** /repos/{owner}/{repo}/hooks/{hook_id} | Delete a repository webhook
[**ReposDisableAutomatedSecurityFixes**](ReposApi.md#ReposDisableAutomatedSecurityFixes) | **Delete** /repos/{owner}/{repo}/automated-security-fixes | Disable automated security fixes
[**ReposDisableLfsForRepo**](ReposApi.md#ReposDisableLfsForRepo) | **Delete** /repos/{owner}/{repo}/lfs | Disable Git LFS for a repository
[**ReposDisableVulnerabilityAlerts**](ReposApi.md#ReposDisableVulnerabilityAlerts) | **Delete** /repos/{owner}/{repo}/vulnerability-alerts | Disable vulnerability alerts
[**ReposDownloadTarballArchive**](ReposApi.md#ReposDownloadTarballArchive) | **Get** /repos/{owner}/{repo}/tarball/{ref} | Download a repository archive (tar)
[**ReposDownloadZipballArchive**](ReposApi.md#ReposDownloadZipballArchive) | **Get** /repos/{owner}/{repo}/zipball/{ref} | Download a repository archive (zip)
[**ReposEnableAutomatedSecurityFixes**](ReposApi.md#ReposEnableAutomatedSecurityFixes) | **Put** /repos/{owner}/{repo}/automated-security-fixes | Enable automated security fixes
[**ReposEnableLfsForRepo**](ReposApi.md#ReposEnableLfsForRepo) | **Put** /repos/{owner}/{repo}/lfs | Enable Git LFS for a repository
[**ReposEnableVulnerabilityAlerts**](ReposApi.md#ReposEnableVulnerabilityAlerts) | **Put** /repos/{owner}/{repo}/vulnerability-alerts | Enable vulnerability alerts
[**ReposGenerateReleaseNotes**](ReposApi.md#ReposGenerateReleaseNotes) | **Post** /repos/{owner}/{repo}/releases/generate-notes | Generate release notes content for a release
[**ReposGet**](ReposApi.md#ReposGet) | **Get** /repos/{owner}/{repo} | Get a repository
[**ReposGetAccessRestrictions**](ReposApi.md#ReposGetAccessRestrictions) | **Get** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions | Get access restrictions
[**ReposGetAdminBranchProtection**](ReposApi.md#ReposGetAdminBranchProtection) | **Get** /repos/{owner}/{repo}/branches/{branch}/protection/enforce_admins | Get admin branch protection
[**ReposGetAllEnvironments**](ReposApi.md#ReposGetAllEnvironments) | **Get** /repos/{owner}/{repo}/environments | List environments
[**ReposGetAllStatusCheckContexts**](ReposApi.md#ReposGetAllStatusCheckContexts) | **Get** /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts | Get all status check contexts
[**ReposGetAllTopics**](ReposApi.md#ReposGetAllTopics) | **Get** /repos/{owner}/{repo}/topics | Get all repository topics
[**ReposGetAppsWithAccessToProtectedBranch**](ReposApi.md#ReposGetAppsWithAccessToProtectedBranch) | **Get** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps | Get apps with access to the protected branch
[**ReposGetAutolink**](ReposApi.md#ReposGetAutolink) | **Get** /repos/{owner}/{repo}/autolinks/{autolink_id} | Get an autolink reference of a repository
[**ReposGetBranch**](ReposApi.md#ReposGetBranch) | **Get** /repos/{owner}/{repo}/branches/{branch} | Get a branch
[**ReposGetBranchProtection**](ReposApi.md#ReposGetBranchProtection) | **Get** /repos/{owner}/{repo}/branches/{branch}/protection | Get branch protection
[**ReposGetClones**](ReposApi.md#ReposGetClones) | **Get** /repos/{owner}/{repo}/traffic/clones | Get repository clones
[**ReposGetCodeFrequencyStats**](ReposApi.md#ReposGetCodeFrequencyStats) | **Get** /repos/{owner}/{repo}/stats/code_frequency | Get the weekly commit activity
[**ReposGetCollaboratorPermissionLevel**](ReposApi.md#ReposGetCollaboratorPermissionLevel) | **Get** /repos/{owner}/{repo}/collaborators/{username}/permission | Get repository permissions for a user
[**ReposGetCombinedStatusForRef**](ReposApi.md#ReposGetCombinedStatusForRef) | **Get** /repos/{owner}/{repo}/commits/{ref}/status | Get the combined status for a specific reference
[**ReposGetCommit**](ReposApi.md#ReposGetCommit) | **Get** /repos/{owner}/{repo}/commits/{ref} | Get a commit
[**ReposGetCommitActivityStats**](ReposApi.md#ReposGetCommitActivityStats) | **Get** /repos/{owner}/{repo}/stats/commit_activity | Get the last year of commit activity
[**ReposGetCommitComment**](ReposApi.md#ReposGetCommitComment) | **Get** /repos/{owner}/{repo}/comments/{comment_id} | Get a commit comment
[**ReposGetCommitSignatureProtection**](ReposApi.md#ReposGetCommitSignatureProtection) | **Get** /repos/{owner}/{repo}/branches/{branch}/protection/required_signatures | Get commit signature protection
[**ReposGetCommunityProfileMetrics**](ReposApi.md#ReposGetCommunityProfileMetrics) | **Get** /repos/{owner}/{repo}/community/profile | Get community profile metrics
[**ReposGetContent**](ReposApi.md#ReposGetContent) | **Get** /repos/{owner}/{repo}/contents/{path} | Get repository content
[**ReposGetContributorsStats**](ReposApi.md#ReposGetContributorsStats) | **Get** /repos/{owner}/{repo}/stats/contributors | Get all contributor commit activity
[**ReposGetDeployKey**](ReposApi.md#ReposGetDeployKey) | **Get** /repos/{owner}/{repo}/keys/{key_id} | Get a deploy key
[**ReposGetDeployment**](ReposApi.md#ReposGetDeployment) | **Get** /repos/{owner}/{repo}/deployments/{deployment_id} | Get a deployment
[**ReposGetDeploymentBranchPolicy**](ReposApi.md#ReposGetDeploymentBranchPolicy) | **Get** /repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies/{branch_policy_id} | Get a deployment branch policy
[**ReposGetDeploymentStatus**](ReposApi.md#ReposGetDeploymentStatus) | **Get** /repos/{owner}/{repo}/deployments/{deployment_id}/statuses/{status_id} | Get a deployment status
[**ReposGetEnvironment**](ReposApi.md#ReposGetEnvironment) | **Get** /repos/{owner}/{repo}/environments/{environment_name} | Get an environment
[**ReposGetLatestPagesBuild**](ReposApi.md#ReposGetLatestPagesBuild) | **Get** /repos/{owner}/{repo}/pages/builds/latest | Get latest Pages build
[**ReposGetLatestRelease**](ReposApi.md#ReposGetLatestRelease) | **Get** /repos/{owner}/{repo}/releases/latest | Get the latest release
[**ReposGetPages**](ReposApi.md#ReposGetPages) | **Get** /repos/{owner}/{repo}/pages | Get a GitHub Pages site
[**ReposGetPagesBuild**](ReposApi.md#ReposGetPagesBuild) | **Get** /repos/{owner}/{repo}/pages/builds/{build_id} | Get GitHub Pages build
[**ReposGetPagesHealthCheck**](ReposApi.md#ReposGetPagesHealthCheck) | **Get** /repos/{owner}/{repo}/pages/health | Get a DNS health check for GitHub Pages
[**ReposGetParticipationStats**](ReposApi.md#ReposGetParticipationStats) | **Get** /repos/{owner}/{repo}/stats/participation | Get the weekly commit count
[**ReposGetPullRequestReviewProtection**](ReposApi.md#ReposGetPullRequestReviewProtection) | **Get** /repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews | Get pull request review protection
[**ReposGetPunchCardStats**](ReposApi.md#ReposGetPunchCardStats) | **Get** /repos/{owner}/{repo}/stats/punch_card | Get the hourly commit count for each day
[**ReposGetReadme**](ReposApi.md#ReposGetReadme) | **Get** /repos/{owner}/{repo}/readme | Get a repository README
[**ReposGetReadmeInDirectory**](ReposApi.md#ReposGetReadmeInDirectory) | **Get** /repos/{owner}/{repo}/readme/{dir} | Get a repository README for a directory
[**ReposGetRelease**](ReposApi.md#ReposGetRelease) | **Get** /repos/{owner}/{repo}/releases/{release_id} | Get a release
[**ReposGetReleaseAsset**](ReposApi.md#ReposGetReleaseAsset) | **Get** /repos/{owner}/{repo}/releases/assets/{asset_id} | Get a release asset
[**ReposGetReleaseByTag**](ReposApi.md#ReposGetReleaseByTag) | **Get** /repos/{owner}/{repo}/releases/tags/{tag} | Get a release by tag name
[**ReposGetStatusChecksProtection**](ReposApi.md#ReposGetStatusChecksProtection) | **Get** /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks | Get status checks protection
[**ReposGetTeamsWithAccessToProtectedBranch**](ReposApi.md#ReposGetTeamsWithAccessToProtectedBranch) | **Get** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams | Get teams with access to the protected branch
[**ReposGetTopPaths**](ReposApi.md#ReposGetTopPaths) | **Get** /repos/{owner}/{repo}/traffic/popular/paths | Get top referral paths
[**ReposGetTopReferrers**](ReposApi.md#ReposGetTopReferrers) | **Get** /repos/{owner}/{repo}/traffic/popular/referrers | Get top referral sources
[**ReposGetUsersWithAccessToProtectedBranch**](ReposApi.md#ReposGetUsersWithAccessToProtectedBranch) | **Get** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users | Get users with access to the protected branch
[**ReposGetViews**](ReposApi.md#ReposGetViews) | **Get** /repos/{owner}/{repo}/traffic/views | Get page views
[**ReposGetWebhook**](ReposApi.md#ReposGetWebhook) | **Get** /repos/{owner}/{repo}/hooks/{hook_id} | Get a repository webhook
[**ReposGetWebhookConfigForRepo**](ReposApi.md#ReposGetWebhookConfigForRepo) | **Get** /repos/{owner}/{repo}/hooks/{hook_id}/config | Get a webhook configuration for a repository
[**ReposGetWebhookDelivery**](ReposApi.md#ReposGetWebhookDelivery) | **Get** /repos/{owner}/{repo}/hooks/{hook_id}/deliveries/{delivery_id} | Get a delivery for a repository webhook
[**ReposListAutolinks**](ReposApi.md#ReposListAutolinks) | **Get** /repos/{owner}/{repo}/autolinks | List all autolinks of a repository
[**ReposListBranches**](ReposApi.md#ReposListBranches) | **Get** /repos/{owner}/{repo}/branches | List branches
[**ReposListBranchesForHeadCommit**](ReposApi.md#ReposListBranchesForHeadCommit) | **Get** /repos/{owner}/{repo}/commits/{commit_sha}/branches-where-head | List branches for HEAD commit
[**ReposListCollaborators**](ReposApi.md#ReposListCollaborators) | **Get** /repos/{owner}/{repo}/collaborators | List repository collaborators
[**ReposListCommentsForCommit**](ReposApi.md#ReposListCommentsForCommit) | **Get** /repos/{owner}/{repo}/commits/{commit_sha}/comments | List commit comments
[**ReposListCommitCommentsForRepo**](ReposApi.md#ReposListCommitCommentsForRepo) | **Get** /repos/{owner}/{repo}/comments | List commit comments for a repository
[**ReposListCommitStatusesForRef**](ReposApi.md#ReposListCommitStatusesForRef) | **Get** /repos/{owner}/{repo}/commits/{ref}/statuses | List commit statuses for a reference
[**ReposListCommits**](ReposApi.md#ReposListCommits) | **Get** /repos/{owner}/{repo}/commits | List commits
[**ReposListContributors**](ReposApi.md#ReposListContributors) | **Get** /repos/{owner}/{repo}/contributors | List repository contributors
[**ReposListDeployKeys**](ReposApi.md#ReposListDeployKeys) | **Get** /repos/{owner}/{repo}/keys | List deploy keys
[**ReposListDeploymentBranchPolicies**](ReposApi.md#ReposListDeploymentBranchPolicies) | **Get** /repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies | List deployment branch policies
[**ReposListDeploymentStatuses**](ReposApi.md#ReposListDeploymentStatuses) | **Get** /repos/{owner}/{repo}/deployments/{deployment_id}/statuses | List deployment statuses
[**ReposListDeployments**](ReposApi.md#ReposListDeployments) | **Get** /repos/{owner}/{repo}/deployments | List deployments
[**ReposListForAuthenticatedUser**](ReposApi.md#ReposListForAuthenticatedUser) | **Get** /user/repos | List repositories for the authenticated user
[**ReposListForOrg**](ReposApi.md#ReposListForOrg) | **Get** /orgs/{org}/repos | List organization repositories
[**ReposListForUser**](ReposApi.md#ReposListForUser) | **Get** /users/{username}/repos | List repositories for a user
[**ReposListForks**](ReposApi.md#ReposListForks) | **Get** /repos/{owner}/{repo}/forks | List forks
[**ReposListInvitations**](ReposApi.md#ReposListInvitations) | **Get** /repos/{owner}/{repo}/invitations | List repository invitations
[**ReposListInvitationsForAuthenticatedUser**](ReposApi.md#ReposListInvitationsForAuthenticatedUser) | **Get** /user/repository_invitations | List repository invitations for the authenticated user
[**ReposListLanguages**](ReposApi.md#ReposListLanguages) | **Get** /repos/{owner}/{repo}/languages | List repository languages
[**ReposListPagesBuilds**](ReposApi.md#ReposListPagesBuilds) | **Get** /repos/{owner}/{repo}/pages/builds | List GitHub Pages builds
[**ReposListPublic**](ReposApi.md#ReposListPublic) | **Get** /repositories | List public repositories
[**ReposListPullRequestsAssociatedWithCommit**](ReposApi.md#ReposListPullRequestsAssociatedWithCommit) | **Get** /repos/{owner}/{repo}/commits/{commit_sha}/pulls | List pull requests associated with a commit
[**ReposListReleaseAssets**](ReposApi.md#ReposListReleaseAssets) | **Get** /repos/{owner}/{repo}/releases/{release_id}/assets | List release assets
[**ReposListReleases**](ReposApi.md#ReposListReleases) | **Get** /repos/{owner}/{repo}/releases | List releases
[**ReposListTagProtection**](ReposApi.md#ReposListTagProtection) | **Get** /repos/{owner}/{repo}/tags/protection | List tag protection states for a repository
[**ReposListTags**](ReposApi.md#ReposListTags) | **Get** /repos/{owner}/{repo}/tags | List repository tags
[**ReposListTeams**](ReposApi.md#ReposListTeams) | **Get** /repos/{owner}/{repo}/teams | List repository teams
[**ReposListWebhookDeliveries**](ReposApi.md#ReposListWebhookDeliveries) | **Get** /repos/{owner}/{repo}/hooks/{hook_id}/deliveries | List deliveries for a repository webhook
[**ReposListWebhooks**](ReposApi.md#ReposListWebhooks) | **Get** /repos/{owner}/{repo}/hooks | List repository webhooks
[**ReposMerge**](ReposApi.md#ReposMerge) | **Post** /repos/{owner}/{repo}/merges | Merge a branch
[**ReposMergeUpstream**](ReposApi.md#ReposMergeUpstream) | **Post** /repos/{owner}/{repo}/merge-upstream | Sync a fork branch with the upstream repository
[**ReposPingWebhook**](ReposApi.md#ReposPingWebhook) | **Post** /repos/{owner}/{repo}/hooks/{hook_id}/pings | Ping a repository webhook
[**ReposRedeliverWebhookDelivery**](ReposApi.md#ReposRedeliverWebhookDelivery) | **Post** /repos/{owner}/{repo}/hooks/{hook_id}/deliveries/{delivery_id}/attempts | Redeliver a delivery for a repository webhook
[**ReposRemoveAppAccessRestrictions**](ReposApi.md#ReposRemoveAppAccessRestrictions) | **Delete** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps | Remove app access restrictions
[**ReposRemoveCollaborator**](ReposApi.md#ReposRemoveCollaborator) | **Delete** /repos/{owner}/{repo}/collaborators/{username} | Remove a repository collaborator
[**ReposRemoveStatusCheckContexts**](ReposApi.md#ReposRemoveStatusCheckContexts) | **Delete** /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts | Remove status check contexts
[**ReposRemoveStatusCheckProtection**](ReposApi.md#ReposRemoveStatusCheckProtection) | **Delete** /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks | Remove status check protection
[**ReposRemoveTeamAccessRestrictions**](ReposApi.md#ReposRemoveTeamAccessRestrictions) | **Delete** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams | Remove team access restrictions
[**ReposRemoveUserAccessRestrictions**](ReposApi.md#ReposRemoveUserAccessRestrictions) | **Delete** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users | Remove user access restrictions
[**ReposRenameBranch**](ReposApi.md#ReposRenameBranch) | **Post** /repos/{owner}/{repo}/branches/{branch}/rename | Rename a branch
[**ReposReplaceAllTopics**](ReposApi.md#ReposReplaceAllTopics) | **Put** /repos/{owner}/{repo}/topics | Replace all repository topics
[**ReposRequestPagesBuild**](ReposApi.md#ReposRequestPagesBuild) | **Post** /repos/{owner}/{repo}/pages/builds | Request a GitHub Pages build
[**ReposSetAdminBranchProtection**](ReposApi.md#ReposSetAdminBranchProtection) | **Post** /repos/{owner}/{repo}/branches/{branch}/protection/enforce_admins | Set admin branch protection
[**ReposSetAppAccessRestrictions**](ReposApi.md#ReposSetAppAccessRestrictions) | **Put** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps | Set app access restrictions
[**ReposSetStatusCheckContexts**](ReposApi.md#ReposSetStatusCheckContexts) | **Put** /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts | Set status check contexts
[**ReposSetTeamAccessRestrictions**](ReposApi.md#ReposSetTeamAccessRestrictions) | **Put** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams | Set team access restrictions
[**ReposSetUserAccessRestrictions**](ReposApi.md#ReposSetUserAccessRestrictions) | **Put** /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users | Set user access restrictions
[**ReposTestPushWebhook**](ReposApi.md#ReposTestPushWebhook) | **Post** /repos/{owner}/{repo}/hooks/{hook_id}/tests | Test the push repository webhook
[**ReposTransfer**](ReposApi.md#ReposTransfer) | **Post** /repos/{owner}/{repo}/transfer | Transfer a repository
[**ReposUpdate**](ReposApi.md#ReposUpdate) | **Patch** /repos/{owner}/{repo} | Update a repository
[**ReposUpdateBranchProtection**](ReposApi.md#ReposUpdateBranchProtection) | **Put** /repos/{owner}/{repo}/branches/{branch}/protection | Update branch protection
[**ReposUpdateCommitComment**](ReposApi.md#ReposUpdateCommitComment) | **Patch** /repos/{owner}/{repo}/comments/{comment_id} | Update a commit comment
[**ReposUpdateDeploymentBranchPolicy**](ReposApi.md#ReposUpdateDeploymentBranchPolicy) | **Put** /repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies/{branch_policy_id} | Update a deployment branch policy
[**ReposUpdateInformationAboutPagesSite**](ReposApi.md#ReposUpdateInformationAboutPagesSite) | **Put** /repos/{owner}/{repo}/pages | Update information about a GitHub Pages site
[**ReposUpdateInvitation**](ReposApi.md#ReposUpdateInvitation) | **Patch** /repos/{owner}/{repo}/invitations/{invitation_id} | Update a repository invitation
[**ReposUpdatePullRequestReviewProtection**](ReposApi.md#ReposUpdatePullRequestReviewProtection) | **Patch** /repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews | Update pull request review protection
[**ReposUpdateRelease**](ReposApi.md#ReposUpdateRelease) | **Patch** /repos/{owner}/{repo}/releases/{release_id} | Update a release
[**ReposUpdateReleaseAsset**](ReposApi.md#ReposUpdateReleaseAsset) | **Patch** /repos/{owner}/{repo}/releases/assets/{asset_id} | Update a release asset
[**ReposUpdateStatusCheckProtection**](ReposApi.md#ReposUpdateStatusCheckProtection) | **Patch** /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks | Update status check protection
[**ReposUpdateWebhook**](ReposApi.md#ReposUpdateWebhook) | **Patch** /repos/{owner}/{repo}/hooks/{hook_id} | Update a repository webhook
[**ReposUpdateWebhookConfigForRepo**](ReposApi.md#ReposUpdateWebhookConfigForRepo) | **Patch** /repos/{owner}/{repo}/hooks/{hook_id}/config | Update a webhook configuration for a repository
[**ReposUploadReleaseAsset**](ReposApi.md#ReposUploadReleaseAsset) | **Post** /repos/{owner}/{repo}/releases/{release_id}/assets | Upload a release asset



## ReposAcceptInvitationForAuthenticatedUser

> ReposAcceptInvitationForAuthenticatedUser(ctx, invitationId).Execute()

Accept a repository invitation



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
    invitationId := int32(56) // int32 | The unique identifier of the invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposAcceptInvitationForAuthenticatedUser(context.Background(), invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposAcceptInvitationForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **int32** | The unique identifier of the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposAcceptInvitationForAuthenticatedUserRequest struct via the builder pattern


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


## ReposAddAppAccessRestrictions

> []Integration ReposAddAppAccessRestrictions(ctx, owner, repo, branch).ReposSetAppAccessRestrictionsRequest(reposSetAppAccessRestrictionsRequest).Execute()

Add app access restrictions



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
    branch := "branch_example" // string | The name of the branch.
    reposSetAppAccessRestrictionsRequest := openapiclient.repos_set_app_access_restrictions_request{ReposSetAppAccessRestrictionsRequestOneOf: openapiclient.NewReposSetAppAccessRestrictionsRequestOneOf([]string{"Apps_example"})} // ReposSetAppAccessRestrictionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposAddAppAccessRestrictions(context.Background(), owner, repo, branch).ReposSetAppAccessRestrictionsRequest(reposSetAppAccessRestrictionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposAddAppAccessRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposAddAppAccessRestrictions`: []Integration
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposAddAppAccessRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposAddAppAccessRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposSetAppAccessRestrictionsRequest** | [**ReposSetAppAccessRestrictionsRequest**](ReposSetAppAccessRestrictionsRequest.md) |  | 

### Return type

[**[]Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposAddCollaborator

> RepositoryInvitation ReposAddCollaborator(ctx, owner, repo, username).ReposAddCollaboratorRequest(reposAddCollaboratorRequest).Execute()

Add a repository collaborator



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
    username := "username_example" // string | The handle for the GitHub user account.
    reposAddCollaboratorRequest := *openapiclient.NewReposAddCollaboratorRequest() // ReposAddCollaboratorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposAddCollaborator(context.Background(), owner, repo, username).ReposAddCollaboratorRequest(reposAddCollaboratorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposAddCollaborator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposAddCollaborator`: RepositoryInvitation
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposAddCollaborator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposAddCollaboratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposAddCollaboratorRequest** | [**ReposAddCollaboratorRequest**](ReposAddCollaboratorRequest.md) |  | 

### Return type

[**RepositoryInvitation**](RepositoryInvitation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposAddStatusCheckContexts

> []string ReposAddStatusCheckContexts(ctx, owner, repo, branch).ReposSetStatusCheckContextsRequest(reposSetStatusCheckContextsRequest).Execute()

Add status check contexts



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
    branch := "branch_example" // string | The name of the branch.
    reposSetStatusCheckContextsRequest := openapiclient.repos_set_status_check_contexts_request{ReposSetStatusCheckContextsRequestOneOf: openapiclient.NewReposSetStatusCheckContextsRequestOneOf([]string{"Contexts_example"})} // ReposSetStatusCheckContextsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposAddStatusCheckContexts(context.Background(), owner, repo, branch).ReposSetStatusCheckContextsRequest(reposSetStatusCheckContextsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposAddStatusCheckContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposAddStatusCheckContexts`: []string
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposAddStatusCheckContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposAddStatusCheckContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposSetStatusCheckContextsRequest** | [**ReposSetStatusCheckContextsRequest**](ReposSetStatusCheckContextsRequest.md) |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposAddTeamAccessRestrictions

> []Team ReposAddTeamAccessRestrictions(ctx, owner, repo, branch).ReposAddTeamAccessRestrictionsRequest(reposAddTeamAccessRestrictionsRequest).Execute()

Add team access restrictions



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
    branch := "branch_example" // string | The name of the branch.
    reposAddTeamAccessRestrictionsRequest := openapiclient.repos_add_team_access_restrictions_request{ReposAddTeamAccessRestrictionsRequestOneOf: openapiclient.NewReposAddTeamAccessRestrictionsRequestOneOf([]string{"Teams_example"})} // ReposAddTeamAccessRestrictionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposAddTeamAccessRestrictions(context.Background(), owner, repo, branch).ReposAddTeamAccessRestrictionsRequest(reposAddTeamAccessRestrictionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposAddTeamAccessRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposAddTeamAccessRestrictions`: []Team
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposAddTeamAccessRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposAddTeamAccessRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposAddTeamAccessRestrictionsRequest** | [**ReposAddTeamAccessRestrictionsRequest**](ReposAddTeamAccessRestrictionsRequest.md) |  | 

### Return type

[**[]Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposAddUserAccessRestrictions

> []SimpleUser ReposAddUserAccessRestrictions(ctx, owner, repo, branch).ReposSetUserAccessRestrictionsRequest(reposSetUserAccessRestrictionsRequest).Execute()

Add user access restrictions



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
    branch := "branch_example" // string | The name of the branch.
    reposSetUserAccessRestrictionsRequest := openapiclient.repos_set_user_access_restrictions_request{ReposSetUserAccessRestrictionsRequestOneOf: openapiclient.NewReposSetUserAccessRestrictionsRequestOneOf([]string{"Users_example"})} // ReposSetUserAccessRestrictionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposAddUserAccessRestrictions(context.Background(), owner, repo, branch).ReposSetUserAccessRestrictionsRequest(reposSetUserAccessRestrictionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposAddUserAccessRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposAddUserAccessRestrictions`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposAddUserAccessRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposAddUserAccessRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposSetUserAccessRestrictionsRequest** | [**ReposSetUserAccessRestrictionsRequest**](ReposSetUserAccessRestrictionsRequest.md) |  | 

### Return type

[**[]SimpleUser**](SimpleUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCheckCollaborator

> ReposCheckCollaborator(ctx, owner, repo, username).Execute()

Check if a user is a repository collaborator



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
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCheckCollaborator(context.Background(), owner, repo, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCheckCollaborator``: %v\n", err)
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
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCheckCollaboratorRequest struct via the builder pattern


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


## ReposCheckVulnerabilityAlerts

> ReposCheckVulnerabilityAlerts(ctx, owner, repo).Execute()

Check if vulnerability alerts are enabled for a repository



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
    resp, r, err := apiClient.ReposApi.ReposCheckVulnerabilityAlerts(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCheckVulnerabilityAlerts``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposCheckVulnerabilityAlertsRequest struct via the builder pattern


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


## ReposCodeownersErrors

> CodeownersErrors ReposCodeownersErrors(ctx, owner, repo).Ref(ref).Execute()

List CODEOWNERS errors



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
    ref := "ref_example" // string | A branch, tag or commit name used to determine which version of the CODEOWNERS file to use. Default: the repository's default branch (e.g. `main`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCodeownersErrors(context.Background(), owner, repo).Ref(ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCodeownersErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCodeownersErrors`: CodeownersErrors
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCodeownersErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCodeownersErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ref** | **string** | A branch, tag or commit name used to determine which version of the CODEOWNERS file to use. Default: the repository&#39;s default branch (e.g. &#x60;main&#x60;) | 

### Return type

[**CodeownersErrors**](CodeownersErrors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCompareCommits

> CommitComparison ReposCompareCommits(ctx, owner, repo, basehead).Page(page).PerPage(perPage).Execute()

Compare two commits



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
    basehead := "basehead_example" // string | The base branch and head branch to compare. This parameter expects the format `{base}...{head}`.
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCompareCommits(context.Background(), owner, repo, basehead).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCompareCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCompareCommits`: CommitComparison
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCompareCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**basehead** | **string** | The base branch and head branch to compare. This parameter expects the format &#x60;{base}...{head}&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCompareCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

### Return type

[**CommitComparison**](CommitComparison.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateAutolink

> Autolink ReposCreateAutolink(ctx, owner, repo).ReposCreateAutolinkRequest(reposCreateAutolinkRequest).Execute()

Create an autolink reference for a repository



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
    reposCreateAutolinkRequest := *openapiclient.NewReposCreateAutolinkRequest("KeyPrefix_example", "UrlTemplate_example") // ReposCreateAutolinkRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateAutolink(context.Background(), owner, repo).ReposCreateAutolinkRequest(reposCreateAutolinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateAutolink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateAutolink`: Autolink
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateAutolink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateAutolinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreateAutolinkRequest** | [**ReposCreateAutolinkRequest**](ReposCreateAutolinkRequest.md) |  | 

### Return type

[**Autolink**](Autolink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateCommitComment

> CommitComment ReposCreateCommitComment(ctx, owner, repo, commitSha).ReposCreateCommitCommentRequest(reposCreateCommitCommentRequest).Execute()

Create a commit comment



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
    commitSha := "commitSha_example" // string | The SHA of the commit.
    reposCreateCommitCommentRequest := *openapiclient.NewReposCreateCommitCommentRequest("Body_example") // ReposCreateCommitCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateCommitComment(context.Background(), owner, repo, commitSha).ReposCreateCommitCommentRequest(reposCreateCommitCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateCommitComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateCommitComment`: CommitComment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateCommitComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**commitSha** | **string** | The SHA of the commit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateCommitCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposCreateCommitCommentRequest** | [**ReposCreateCommitCommentRequest**](ReposCreateCommitCommentRequest.md) |  | 

### Return type

[**CommitComment**](CommitComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateCommitSignatureProtection

> ProtectedBranchAdminEnforced ReposCreateCommitSignatureProtection(ctx, owner, repo, branch).Execute()

Create commit signature protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateCommitSignatureProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateCommitSignatureProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateCommitSignatureProtection`: ProtectedBranchAdminEnforced
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateCommitSignatureProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateCommitSignatureProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProtectedBranchAdminEnforced**](ProtectedBranchAdminEnforced.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateCommitStatus

> Status ReposCreateCommitStatus(ctx, owner, repo, sha).ReposCreateCommitStatusRequest(reposCreateCommitStatusRequest).Execute()

Create a commit status



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
    sha := "sha_example" // string | 
    reposCreateCommitStatusRequest := *openapiclient.NewReposCreateCommitStatusRequest("State_example") // ReposCreateCommitStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateCommitStatus(context.Background(), owner, repo, sha).ReposCreateCommitStatusRequest(reposCreateCommitStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateCommitStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateCommitStatus`: Status
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateCommitStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**sha** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateCommitStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposCreateCommitStatusRequest** | [**ReposCreateCommitStatusRequest**](ReposCreateCommitStatusRequest.md) |  | 

### Return type

[**Status**](Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateDeployKey

> DeployKey ReposCreateDeployKey(ctx, owner, repo).ReposCreateDeployKeyRequest(reposCreateDeployKeyRequest).Execute()

Create a deploy key



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
    reposCreateDeployKeyRequest := *openapiclient.NewReposCreateDeployKeyRequest("Key_example") // ReposCreateDeployKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateDeployKey(context.Background(), owner, repo).ReposCreateDeployKeyRequest(reposCreateDeployKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateDeployKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateDeployKey`: DeployKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateDeployKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateDeployKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreateDeployKeyRequest** | [**ReposCreateDeployKeyRequest**](ReposCreateDeployKeyRequest.md) |  | 

### Return type

[**DeployKey**](DeployKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateDeployment

> Deployment ReposCreateDeployment(ctx, owner, repo).ReposCreateDeploymentRequest(reposCreateDeploymentRequest).Execute()

Create a deployment



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
    reposCreateDeploymentRequest := *openapiclient.NewReposCreateDeploymentRequest("Ref_example") // ReposCreateDeploymentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateDeployment(context.Background(), owner, repo).ReposCreateDeploymentRequest(reposCreateDeploymentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateDeployment`: Deployment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreateDeploymentRequest** | [**ReposCreateDeploymentRequest**](ReposCreateDeploymentRequest.md) |  | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateDeploymentBranchPolicy

> DeploymentBranchPolicy ReposCreateDeploymentBranchPolicy(ctx, owner, repo, environmentName).DeploymentBranchPolicyNamePattern(deploymentBranchPolicyNamePattern).Execute()

Create a deployment branch policy



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
    environmentName := "environmentName_example" // string | The name of the environment.
    deploymentBranchPolicyNamePattern := *openapiclient.NewDeploymentBranchPolicyNamePattern("release/*") // DeploymentBranchPolicyNamePattern | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateDeploymentBranchPolicy(context.Background(), owner, repo, environmentName).DeploymentBranchPolicyNamePattern(deploymentBranchPolicyNamePattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateDeploymentBranchPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateDeploymentBranchPolicy`: DeploymentBranchPolicy
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateDeploymentBranchPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**environmentName** | **string** | The name of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateDeploymentBranchPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deploymentBranchPolicyNamePattern** | [**DeploymentBranchPolicyNamePattern**](DeploymentBranchPolicyNamePattern.md) |  | 

### Return type

[**DeploymentBranchPolicy**](DeploymentBranchPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateDeploymentStatus

> DeploymentStatus ReposCreateDeploymentStatus(ctx, owner, repo, deploymentId).ReposCreateDeploymentStatusRequest(reposCreateDeploymentStatusRequest).Execute()

Create a deployment status



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
    deploymentId := int32(56) // int32 | deployment_id parameter
    reposCreateDeploymentStatusRequest := *openapiclient.NewReposCreateDeploymentStatusRequest("State_example") // ReposCreateDeploymentStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateDeploymentStatus(context.Background(), owner, repo, deploymentId).ReposCreateDeploymentStatusRequest(reposCreateDeploymentStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateDeploymentStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateDeploymentStatus`: DeploymentStatus
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateDeploymentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**deploymentId** | **int32** | deployment_id parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateDeploymentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposCreateDeploymentStatusRequest** | [**ReposCreateDeploymentStatusRequest**](ReposCreateDeploymentStatusRequest.md) |  | 

### Return type

[**DeploymentStatus**](DeploymentStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateDispatchEvent

> ReposCreateDispatchEvent(ctx, owner, repo).ReposCreateDispatchEventRequest(reposCreateDispatchEventRequest).Execute()

Create a repository dispatch event



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
    reposCreateDispatchEventRequest := *openapiclient.NewReposCreateDispatchEventRequest("EventType_example") // ReposCreateDispatchEventRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateDispatchEvent(context.Background(), owner, repo).ReposCreateDispatchEventRequest(reposCreateDispatchEventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateDispatchEvent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposCreateDispatchEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreateDispatchEventRequest** | [**ReposCreateDispatchEventRequest**](ReposCreateDispatchEventRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateForAuthenticatedUser

> Repository ReposCreateForAuthenticatedUser(ctx).ReposCreateForAuthenticatedUserRequest(reposCreateForAuthenticatedUserRequest).Execute()

Create a repository for the authenticated user



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
    reposCreateForAuthenticatedUserRequest := *openapiclient.NewReposCreateForAuthenticatedUserRequest("Team Environment") // ReposCreateForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateForAuthenticatedUser(context.Background()).ReposCreateForAuthenticatedUserRequest(reposCreateForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateForAuthenticatedUser`: Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reposCreateForAuthenticatedUserRequest** | [**ReposCreateForAuthenticatedUserRequest**](ReposCreateForAuthenticatedUserRequest.md) |  | 

### Return type

[**Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateFork

> FullRepository ReposCreateFork(ctx, owner, repo).ReposCreateForkRequest(reposCreateForkRequest).Execute()

Create a fork



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
    reposCreateForkRequest := *openapiclient.NewReposCreateForkRequest() // ReposCreateForkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateFork(context.Background(), owner, repo).ReposCreateForkRequest(reposCreateForkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateFork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateFork`: FullRepository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateFork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreateForkRequest** | [**ReposCreateForkRequest**](ReposCreateForkRequest.md) |  | 

### Return type

[**FullRepository**](FullRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateInOrg

> Repository ReposCreateInOrg(ctx, org).ReposCreateInOrgRequest(reposCreateInOrgRequest).Execute()

Create an organization repository



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
    reposCreateInOrgRequest := *openapiclient.NewReposCreateInOrgRequest("Name_example") // ReposCreateInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateInOrg(context.Background(), org).ReposCreateInOrgRequest(reposCreateInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateInOrg`: Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reposCreateInOrgRequest** | [**ReposCreateInOrgRequest**](ReposCreateInOrgRequest.md) |  | 

### Return type

[**Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateOrUpdateEnvironment

> Environment ReposCreateOrUpdateEnvironment(ctx, owner, repo, environmentName).ReposCreateOrUpdateEnvironmentRequest(reposCreateOrUpdateEnvironmentRequest).Execute()

Create or update an environment



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
    environmentName := "environmentName_example" // string | The name of the environment.
    reposCreateOrUpdateEnvironmentRequest := *openapiclient.NewReposCreateOrUpdateEnvironmentRequest() // ReposCreateOrUpdateEnvironmentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateOrUpdateEnvironment(context.Background(), owner, repo, environmentName).ReposCreateOrUpdateEnvironmentRequest(reposCreateOrUpdateEnvironmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateOrUpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateOrUpdateEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateOrUpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**environmentName** | **string** | The name of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateOrUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposCreateOrUpdateEnvironmentRequest** | [**ReposCreateOrUpdateEnvironmentRequest**](ReposCreateOrUpdateEnvironmentRequest.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateOrUpdateFileContents

> FileCommit ReposCreateOrUpdateFileContents(ctx, owner, repo, path).ReposCreateOrUpdateFileContentsRequest(reposCreateOrUpdateFileContentsRequest).Execute()

Create or update file contents



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
    path := "path_example" // string | path parameter
    reposCreateOrUpdateFileContentsRequest := *openapiclient.NewReposCreateOrUpdateFileContentsRequest("Message_example", "Content_example") // ReposCreateOrUpdateFileContentsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateOrUpdateFileContents(context.Background(), owner, repo, path).ReposCreateOrUpdateFileContentsRequest(reposCreateOrUpdateFileContentsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateOrUpdateFileContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateOrUpdateFileContents`: FileCommit
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateOrUpdateFileContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**path** | **string** | path parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateOrUpdateFileContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposCreateOrUpdateFileContentsRequest** | [**ReposCreateOrUpdateFileContentsRequest**](ReposCreateOrUpdateFileContentsRequest.md) |  | 

### Return type

[**FileCommit**](FileCommit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreatePagesDeployment

> PageDeployment ReposCreatePagesDeployment(ctx, owner, repo).ReposCreatePagesDeploymentRequest(reposCreatePagesDeploymentRequest).Execute()

Create a GitHub Pages deployment



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
    reposCreatePagesDeploymentRequest := *openapiclient.NewReposCreatePagesDeploymentRequest("ArtifactUrl_example", "PagesBuildVersion_example", "OidcToken_example") // ReposCreatePagesDeploymentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreatePagesDeployment(context.Background(), owner, repo).ReposCreatePagesDeploymentRequest(reposCreatePagesDeploymentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreatePagesDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreatePagesDeployment`: PageDeployment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreatePagesDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreatePagesDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreatePagesDeploymentRequest** | [**ReposCreatePagesDeploymentRequest**](ReposCreatePagesDeploymentRequest.md) |  | 

### Return type

[**PageDeployment**](PageDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreatePagesSite

> Page ReposCreatePagesSite(ctx, owner, repo).ReposCreatePagesSiteRequest(reposCreatePagesSiteRequest).Execute()

Create a GitHub Pages site



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
    reposCreatePagesSiteRequest := *openapiclient.NewReposCreatePagesSiteRequest() // ReposCreatePagesSiteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreatePagesSite(context.Background(), owner, repo).ReposCreatePagesSiteRequest(reposCreatePagesSiteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreatePagesSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreatePagesSite`: Page
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreatePagesSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreatePagesSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreatePagesSiteRequest** | [**ReposCreatePagesSiteRequest**](ReposCreatePagesSiteRequest.md) |  | 

### Return type

[**Page**](Page.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateRelease

> Release ReposCreateRelease(ctx, owner, repo).ReposCreateReleaseRequest(reposCreateReleaseRequest).Execute()

Create a release



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
    reposCreateReleaseRequest := *openapiclient.NewReposCreateReleaseRequest("TagName_example") // ReposCreateReleaseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateRelease(context.Background(), owner, repo).ReposCreateReleaseRequest(reposCreateReleaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateRelease`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreateReleaseRequest** | [**ReposCreateReleaseRequest**](ReposCreateReleaseRequest.md) |  | 

### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateTagProtection

> TagProtection ReposCreateTagProtection(ctx, owner, repo).ReposCreateTagProtectionRequest(reposCreateTagProtectionRequest).Execute()

Create a tag protection state for a repository



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
    reposCreateTagProtectionRequest := *openapiclient.NewReposCreateTagProtectionRequest("Pattern_example") // ReposCreateTagProtectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateTagProtection(context.Background(), owner, repo).ReposCreateTagProtectionRequest(reposCreateTagProtectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateTagProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateTagProtection`: TagProtection
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateTagProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateTagProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreateTagProtectionRequest** | [**ReposCreateTagProtectionRequest**](ReposCreateTagProtectionRequest.md) |  | 

### Return type

[**TagProtection**](TagProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateUsingTemplate

> Repository ReposCreateUsingTemplate(ctx, templateOwner, templateRepo).ReposCreateUsingTemplateRequest(reposCreateUsingTemplateRequest).Execute()

Create a repository using a template



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
    templateOwner := "templateOwner_example" // string | 
    templateRepo := "templateRepo_example" // string | 
    reposCreateUsingTemplateRequest := *openapiclient.NewReposCreateUsingTemplateRequest("Name_example") // ReposCreateUsingTemplateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateUsingTemplate(context.Background(), templateOwner, templateRepo).ReposCreateUsingTemplateRequest(reposCreateUsingTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateUsingTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateUsingTemplate`: Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateUsingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateOwner** | **string** |  | 
**templateRepo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateUsingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreateUsingTemplateRequest** | [**ReposCreateUsingTemplateRequest**](ReposCreateUsingTemplateRequest.md) |  | 

### Return type

[**Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreateWebhook

> Hook ReposCreateWebhook(ctx, owner, repo).ReposCreateWebhookRequest(reposCreateWebhookRequest).Execute()

Create a repository webhook



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
    reposCreateWebhookRequest := *openapiclient.NewReposCreateWebhookRequest() // ReposCreateWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreateWebhook(context.Background(), owner, repo).ReposCreateWebhookRequest(reposCreateWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreateWebhook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposCreateWebhookRequest** | [**ReposCreateWebhookRequest**](ReposCreateWebhookRequest.md) |  | 

### Return type

[**Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposDeclineInvitationForAuthenticatedUser

> ReposDeclineInvitationForAuthenticatedUser(ctx, invitationId).Execute()

Decline a repository invitation



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
    invitationId := int32(56) // int32 | The unique identifier of the invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeclineInvitationForAuthenticatedUser(context.Background(), invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeclineInvitationForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **int32** | The unique identifier of the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeclineInvitationForAuthenticatedUserRequest struct via the builder pattern


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


## ReposDelete

> ReposDelete(ctx, owner, repo).Execute()

Delete a repository



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
    resp, r, err := apiClient.ReposApi.ReposDelete(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposDeleteRequest struct via the builder pattern


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


## ReposDeleteAccessRestrictions

> ReposDeleteAccessRestrictions(ctx, owner, repo, branch).Execute()

Delete access restrictions



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteAccessRestrictions(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteAccessRestrictions``: %v\n", err)
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
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteAccessRestrictionsRequest struct via the builder pattern


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


## ReposDeleteAdminBranchProtection

> ReposDeleteAdminBranchProtection(ctx, owner, repo, branch).Execute()

Delete admin branch protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteAdminBranchProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteAdminBranchProtection``: %v\n", err)
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
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteAdminBranchProtectionRequest struct via the builder pattern


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


## ReposDeleteAnEnvironment

> ReposDeleteAnEnvironment(ctx, owner, repo, environmentName).Execute()

Delete an environment



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
    environmentName := "environmentName_example" // string | The name of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteAnEnvironment(context.Background(), owner, repo, environmentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteAnEnvironment``: %v\n", err)
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
**environmentName** | **string** | The name of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteAnEnvironmentRequest struct via the builder pattern


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


## ReposDeleteAutolink

> ReposDeleteAutolink(ctx, owner, repo, autolinkId).Execute()

Delete an autolink reference from a repository



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
    autolinkId := int32(56) // int32 | The unique identifier of the autolink.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteAutolink(context.Background(), owner, repo, autolinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteAutolink``: %v\n", err)
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
**autolinkId** | **int32** | The unique identifier of the autolink. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteAutolinkRequest struct via the builder pattern


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


## ReposDeleteBranchProtection

> ReposDeleteBranchProtection(ctx, owner, repo, branch).Execute()

Delete branch protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteBranchProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteBranchProtection``: %v\n", err)
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
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteBranchProtectionRequest struct via the builder pattern


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


## ReposDeleteCommitComment

> ReposDeleteCommitComment(ctx, owner, repo, commentId).Execute()

Delete a commit comment



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
    commentId := int32(56) // int32 | The unique identifier of the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteCommitComment(context.Background(), owner, repo, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteCommitComment``: %v\n", err)
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
**commentId** | **int32** | The unique identifier of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteCommitCommentRequest struct via the builder pattern


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


## ReposDeleteCommitSignatureProtection

> ReposDeleteCommitSignatureProtection(ctx, owner, repo, branch).Execute()

Delete commit signature protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteCommitSignatureProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteCommitSignatureProtection``: %v\n", err)
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
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteCommitSignatureProtectionRequest struct via the builder pattern


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


## ReposDeleteDeployKey

> ReposDeleteDeployKey(ctx, owner, repo, keyId).Execute()

Delete a deploy key



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
    keyId := int32(56) // int32 | The unique identifier of the key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteDeployKey(context.Background(), owner, repo, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteDeployKey``: %v\n", err)
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
**keyId** | **int32** | The unique identifier of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteDeployKeyRequest struct via the builder pattern


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


## ReposDeleteDeployment

> ReposDeleteDeployment(ctx, owner, repo, deploymentId).Execute()

Delete a deployment



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
    deploymentId := int32(56) // int32 | deployment_id parameter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteDeployment(context.Background(), owner, repo, deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteDeployment``: %v\n", err)
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
**deploymentId** | **int32** | deployment_id parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteDeploymentRequest struct via the builder pattern


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


## ReposDeleteDeploymentBranchPolicy

> ReposDeleteDeploymentBranchPolicy(ctx, owner, repo, environmentName, branchPolicyId).Execute()

Delete a deployment branch policy



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
    environmentName := "environmentName_example" // string | The name of the environment.
    branchPolicyId := int32(56) // int32 | The unique identifier of the branch policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteDeploymentBranchPolicy(context.Background(), owner, repo, environmentName, branchPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteDeploymentBranchPolicy``: %v\n", err)
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
**environmentName** | **string** | The name of the environment. | 
**branchPolicyId** | **int32** | The unique identifier of the branch policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteDeploymentBranchPolicyRequest struct via the builder pattern


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


## ReposDeleteFile

> FileCommit ReposDeleteFile(ctx, owner, repo, path).ReposDeleteFileRequest(reposDeleteFileRequest).Execute()

Delete a file



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
    path := "path_example" // string | path parameter
    reposDeleteFileRequest := *openapiclient.NewReposDeleteFileRequest("Message_example", "Sha_example") // ReposDeleteFileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteFile(context.Background(), owner, repo, path).ReposDeleteFileRequest(reposDeleteFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposDeleteFile`: FileCommit
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposDeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**path** | **string** | path parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposDeleteFileRequest** | [**ReposDeleteFileRequest**](ReposDeleteFileRequest.md) |  | 

### Return type

[**FileCommit**](FileCommit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposDeleteInvitation

> ReposDeleteInvitation(ctx, owner, repo, invitationId).Execute()

Delete a repository invitation



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
    invitationId := int32(56) // int32 | The unique identifier of the invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteInvitation(context.Background(), owner, repo, invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteInvitation``: %v\n", err)
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
**invitationId** | **int32** | The unique identifier of the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteInvitationRequest struct via the builder pattern


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


## ReposDeletePagesSite

> ReposDeletePagesSite(ctx, owner, repo).Execute()

Delete a GitHub Pages site



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
    resp, r, err := apiClient.ReposApi.ReposDeletePagesSite(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeletePagesSite``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposDeletePagesSiteRequest struct via the builder pattern


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


## ReposDeletePullRequestReviewProtection

> ReposDeletePullRequestReviewProtection(ctx, owner, repo, branch).Execute()

Delete pull request review protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeletePullRequestReviewProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeletePullRequestReviewProtection``: %v\n", err)
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
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeletePullRequestReviewProtectionRequest struct via the builder pattern


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


## ReposDeleteRelease

> ReposDeleteRelease(ctx, owner, repo, releaseId).Execute()

Delete a release



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
    releaseId := int32(56) // int32 | The unique identifier of the release.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteRelease(context.Background(), owner, repo, releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteRelease``: %v\n", err)
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
**releaseId** | **int32** | The unique identifier of the release. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteReleaseRequest struct via the builder pattern


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


## ReposDeleteReleaseAsset

> ReposDeleteReleaseAsset(ctx, owner, repo, assetId).Execute()

Delete a release asset



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
    assetId := int32(56) // int32 | The unique identifier of the asset.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteReleaseAsset(context.Background(), owner, repo, assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteReleaseAsset``: %v\n", err)
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
**assetId** | **int32** | The unique identifier of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteReleaseAssetRequest struct via the builder pattern


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


## ReposDeleteTagProtection

> ReposDeleteTagProtection(ctx, owner, repo, tagProtectionId).Execute()

Delete a tag protection state for a repository



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
    tagProtectionId := int32(56) // int32 | The unique identifier of the tag protection.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteTagProtection(context.Background(), owner, repo, tagProtectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteTagProtection``: %v\n", err)
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
**tagProtectionId** | **int32** | The unique identifier of the tag protection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteTagProtectionRequest struct via the builder pattern


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


## ReposDeleteWebhook

> ReposDeleteWebhook(ctx, owner, repo, hookId).Execute()

Delete a repository webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDeleteWebhook(context.Background(), owner, repo, hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDeleteWebhook``: %v\n", err)
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
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteWebhookRequest struct via the builder pattern


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


## ReposDisableAutomatedSecurityFixes

> ReposDisableAutomatedSecurityFixes(ctx, owner, repo).Execute()

Disable automated security fixes



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
    resp, r, err := apiClient.ReposApi.ReposDisableAutomatedSecurityFixes(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDisableAutomatedSecurityFixes``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposDisableAutomatedSecurityFixesRequest struct via the builder pattern


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


## ReposDisableLfsForRepo

> ReposDisableLfsForRepo(ctx, owner, repo).Execute()

Disable Git LFS for a repository



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
    resp, r, err := apiClient.ReposApi.ReposDisableLfsForRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDisableLfsForRepo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposDisableLfsForRepoRequest struct via the builder pattern


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


## ReposDisableVulnerabilityAlerts

> ReposDisableVulnerabilityAlerts(ctx, owner, repo).Execute()

Disable vulnerability alerts



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
    resp, r, err := apiClient.ReposApi.ReposDisableVulnerabilityAlerts(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDisableVulnerabilityAlerts``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposDisableVulnerabilityAlertsRequest struct via the builder pattern


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


## ReposDownloadTarballArchive

> ReposDownloadTarballArchive(ctx, owner, repo, ref).Execute()

Download a repository archive (tar)



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
    ref := "ref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDownloadTarballArchive(context.Background(), owner, repo, ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDownloadTarballArchive``: %v\n", err)
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
**ref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDownloadTarballArchiveRequest struct via the builder pattern


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


## ReposDownloadZipballArchive

> ReposDownloadZipballArchive(ctx, owner, repo, ref).Execute()

Download a repository archive (zip)



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
    ref := "ref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDownloadZipballArchive(context.Background(), owner, repo, ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDownloadZipballArchive``: %v\n", err)
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
**ref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDownloadZipballArchiveRequest struct via the builder pattern


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


## ReposEnableAutomatedSecurityFixes

> ReposEnableAutomatedSecurityFixes(ctx, owner, repo).Execute()

Enable automated security fixes



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
    resp, r, err := apiClient.ReposApi.ReposEnableAutomatedSecurityFixes(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposEnableAutomatedSecurityFixes``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposEnableAutomatedSecurityFixesRequest struct via the builder pattern


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


## ReposEnableLfsForRepo

> map[string]interface{} ReposEnableLfsForRepo(ctx, owner, repo).Execute()

Enable Git LFS for a repository



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
    resp, r, err := apiClient.ReposApi.ReposEnableLfsForRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposEnableLfsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposEnableLfsForRepo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposEnableLfsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposEnableLfsForRepoRequest struct via the builder pattern


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


## ReposEnableVulnerabilityAlerts

> ReposEnableVulnerabilityAlerts(ctx, owner, repo).Execute()

Enable vulnerability alerts



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
    resp, r, err := apiClient.ReposApi.ReposEnableVulnerabilityAlerts(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposEnableVulnerabilityAlerts``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposEnableVulnerabilityAlertsRequest struct via the builder pattern


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


## ReposGenerateReleaseNotes

> ReleaseNotesContent ReposGenerateReleaseNotes(ctx, owner, repo).ReposGenerateReleaseNotesRequest(reposGenerateReleaseNotesRequest).Execute()

Generate release notes content for a release



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
    reposGenerateReleaseNotesRequest := *openapiclient.NewReposGenerateReleaseNotesRequest("TagName_example") // ReposGenerateReleaseNotesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGenerateReleaseNotes(context.Background(), owner, repo).ReposGenerateReleaseNotesRequest(reposGenerateReleaseNotesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGenerateReleaseNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGenerateReleaseNotes`: ReleaseNotesContent
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGenerateReleaseNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGenerateReleaseNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposGenerateReleaseNotesRequest** | [**ReposGenerateReleaseNotesRequest**](ReposGenerateReleaseNotesRequest.md) |  | 

### Return type

[**ReleaseNotesContent**](ReleaseNotesContent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGet

> FullRepository ReposGet(ctx, owner, repo).Execute()

Get a repository



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
    resp, r, err := apiClient.ReposApi.ReposGet(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGet`: FullRepository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FullRepository**](FullRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetAccessRestrictions

> BranchRestrictionPolicy ReposGetAccessRestrictions(ctx, owner, repo, branch).Execute()

Get access restrictions



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetAccessRestrictions(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetAccessRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetAccessRestrictions`: BranchRestrictionPolicy
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetAccessRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetAccessRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BranchRestrictionPolicy**](BranchRestrictionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetAdminBranchProtection

> ProtectedBranchAdminEnforced ReposGetAdminBranchProtection(ctx, owner, repo, branch).Execute()

Get admin branch protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetAdminBranchProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetAdminBranchProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetAdminBranchProtection`: ProtectedBranchAdminEnforced
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetAdminBranchProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetAdminBranchProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProtectedBranchAdminEnforced**](ProtectedBranchAdminEnforced.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetAllEnvironments

> ReposGetAllEnvironments200Response ReposGetAllEnvironments(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List environments



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
    resp, r, err := apiClient.ReposApi.ReposGetAllEnvironments(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetAllEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetAllEnvironments`: ReposGetAllEnvironments200Response
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetAllEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetAllEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ReposGetAllEnvironments200Response**](ReposGetAllEnvironments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetAllStatusCheckContexts

> []string ReposGetAllStatusCheckContexts(ctx, owner, repo, branch).Execute()

Get all status check contexts



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetAllStatusCheckContexts(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetAllStatusCheckContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetAllStatusCheckContexts`: []string
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetAllStatusCheckContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetAllStatusCheckContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetAllTopics

> Topic ReposGetAllTopics(ctx, owner, repo).Page(page).PerPage(perPage).Execute()

Get all repository topics



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
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetAllTopics(context.Background(), owner, repo).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetAllTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetAllTopics`: Topic
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetAllTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetAllTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

### Return type

[**Topic**](Topic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetAppsWithAccessToProtectedBranch

> []Integration ReposGetAppsWithAccessToProtectedBranch(ctx, owner, repo, branch).Execute()

Get apps with access to the protected branch



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetAppsWithAccessToProtectedBranch(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetAppsWithAccessToProtectedBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetAppsWithAccessToProtectedBranch`: []Integration
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetAppsWithAccessToProtectedBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetAppsWithAccessToProtectedBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetAutolink

> Autolink ReposGetAutolink(ctx, owner, repo, autolinkId).Execute()

Get an autolink reference of a repository



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
    autolinkId := int32(56) // int32 | The unique identifier of the autolink.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetAutolink(context.Background(), owner, repo, autolinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetAutolink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetAutolink`: Autolink
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetAutolink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**autolinkId** | **int32** | The unique identifier of the autolink. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetAutolinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Autolink**](Autolink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetBranch

> BranchWithProtection ReposGetBranch(ctx, owner, repo, branch).Execute()

Get a branch



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetBranch(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetBranch`: BranchWithProtection
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BranchWithProtection**](BranchWithProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetBranchProtection

> BranchProtection ReposGetBranchProtection(ctx, owner, repo, branch).Execute()

Get branch protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetBranchProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetBranchProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetBranchProtection`: BranchProtection
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetBranchProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetBranchProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BranchProtection**](BranchProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetClones

> CloneTraffic ReposGetClones(ctx, owner, repo).Per(per).Execute()

Get repository clones



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
    per := "per_example" // string | The time frame to display results for. (optional) (default to "day")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetClones(context.Background(), owner, repo).Per(per).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetClones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetClones`: CloneTraffic
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetClones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetClonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **per** | **string** | The time frame to display results for. | [default to &quot;day&quot;]

### Return type

[**CloneTraffic**](CloneTraffic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetCodeFrequencyStats

> [][]int32 ReposGetCodeFrequencyStats(ctx, owner, repo).Execute()

Get the weekly commit activity



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
    resp, r, err := apiClient.ReposApi.ReposGetCodeFrequencyStats(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetCodeFrequencyStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetCodeFrequencyStats`: [][]int32
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetCodeFrequencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetCodeFrequencyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[][]int32**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetCollaboratorPermissionLevel

> RepositoryCollaboratorPermission ReposGetCollaboratorPermissionLevel(ctx, owner, repo, username).Execute()

Get repository permissions for a user



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
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetCollaboratorPermissionLevel(context.Background(), owner, repo, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetCollaboratorPermissionLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetCollaboratorPermissionLevel`: RepositoryCollaboratorPermission
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetCollaboratorPermissionLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetCollaboratorPermissionLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RepositoryCollaboratorPermission**](RepositoryCollaboratorPermission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetCombinedStatusForRef

> CombinedCommitStatus ReposGetCombinedStatusForRef(ctx, owner, repo, ref).PerPage(perPage).Page(page).Execute()

Get the combined status for a specific reference



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
    ref := "ref_example" // string | ref parameter
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetCombinedStatusForRef(context.Background(), owner, repo, ref).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetCombinedStatusForRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetCombinedStatusForRef`: CombinedCommitStatus
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetCombinedStatusForRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**ref** | **string** | ref parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetCombinedStatusForRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**CombinedCommitStatus**](CombinedCommitStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetCommit

> Commit ReposGetCommit(ctx, owner, repo, ref).Page(page).PerPage(perPage).Execute()

Get a commit



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
    ref := "ref_example" // string | ref parameter
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetCommit(context.Background(), owner, repo, ref).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetCommit`: Commit
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**ref** | **string** | ref parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

### Return type

[**Commit**](Commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetCommitActivityStats

> []CommitActivity ReposGetCommitActivityStats(ctx, owner, repo).Execute()

Get the last year of commit activity



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
    resp, r, err := apiClient.ReposApi.ReposGetCommitActivityStats(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetCommitActivityStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetCommitActivityStats`: []CommitActivity
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetCommitActivityStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetCommitActivityStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CommitActivity**](CommitActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetCommitComment

> CommitComment ReposGetCommitComment(ctx, owner, repo, commentId).Execute()

Get a commit comment



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
    commentId := int32(56) // int32 | The unique identifier of the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetCommitComment(context.Background(), owner, repo, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetCommitComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetCommitComment`: CommitComment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetCommitComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**commentId** | **int32** | The unique identifier of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetCommitCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CommitComment**](CommitComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetCommitSignatureProtection

> ProtectedBranchAdminEnforced ReposGetCommitSignatureProtection(ctx, owner, repo, branch).Execute()

Get commit signature protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetCommitSignatureProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetCommitSignatureProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetCommitSignatureProtection`: ProtectedBranchAdminEnforced
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetCommitSignatureProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetCommitSignatureProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProtectedBranchAdminEnforced**](ProtectedBranchAdminEnforced.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetCommunityProfileMetrics

> CommunityProfile ReposGetCommunityProfileMetrics(ctx, owner, repo).Execute()

Get community profile metrics



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
    resp, r, err := apiClient.ReposApi.ReposGetCommunityProfileMetrics(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetCommunityProfileMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetCommunityProfileMetrics`: CommunityProfile
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetCommunityProfileMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetCommunityProfileMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CommunityProfile**](CommunityProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetContent

> ContentTree ReposGetContent(ctx, owner, repo, path).Ref(ref).Execute()

Get repository content



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
    path := "path_example" // string | path parameter
    ref := "ref_example" // string | The name of the commit/branch/tag. Default: the repository’s default branch (usually `master`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetContent(context.Background(), owner, repo, path).Ref(ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetContent`: ContentTree
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**path** | **string** | path parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ref** | **string** | The name of the commit/branch/tag. Default: the repository’s default branch (usually &#x60;master&#x60;) | 

### Return type

[**ContentTree**](ContentTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.github.object, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetContributorsStats

> []ContributorActivity ReposGetContributorsStats(ctx, owner, repo).Execute()

Get all contributor commit activity



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
    resp, r, err := apiClient.ReposApi.ReposGetContributorsStats(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetContributorsStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetContributorsStats`: []ContributorActivity
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetContributorsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetContributorsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ContributorActivity**](ContributorActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetDeployKey

> DeployKey ReposGetDeployKey(ctx, owner, repo, keyId).Execute()

Get a deploy key



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
    keyId := int32(56) // int32 | The unique identifier of the key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetDeployKey(context.Background(), owner, repo, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetDeployKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetDeployKey`: DeployKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetDeployKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**keyId** | **int32** | The unique identifier of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetDeployKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeployKey**](DeployKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetDeployment

> Deployment ReposGetDeployment(ctx, owner, repo, deploymentId).Execute()

Get a deployment



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
    deploymentId := int32(56) // int32 | deployment_id parameter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetDeployment(context.Background(), owner, repo, deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetDeployment`: Deployment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**deploymentId** | **int32** | deployment_id parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetDeploymentBranchPolicy

> DeploymentBranchPolicy ReposGetDeploymentBranchPolicy(ctx, owner, repo, environmentName, branchPolicyId).Execute()

Get a deployment branch policy



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
    environmentName := "environmentName_example" // string | The name of the environment.
    branchPolicyId := int32(56) // int32 | The unique identifier of the branch policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetDeploymentBranchPolicy(context.Background(), owner, repo, environmentName, branchPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetDeploymentBranchPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetDeploymentBranchPolicy`: DeploymentBranchPolicy
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetDeploymentBranchPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**environmentName** | **string** | The name of the environment. | 
**branchPolicyId** | **int32** | The unique identifier of the branch policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetDeploymentBranchPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DeploymentBranchPolicy**](DeploymentBranchPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetDeploymentStatus

> DeploymentStatus ReposGetDeploymentStatus(ctx, owner, repo, deploymentId, statusId).Execute()

Get a deployment status



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
    deploymentId := int32(56) // int32 | deployment_id parameter
    statusId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetDeploymentStatus(context.Background(), owner, repo, deploymentId, statusId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetDeploymentStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetDeploymentStatus`: DeploymentStatus
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetDeploymentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**deploymentId** | **int32** | deployment_id parameter | 
**statusId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetDeploymentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DeploymentStatus**](DeploymentStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetEnvironment

> Environment ReposGetEnvironment(ctx, owner, repo, environmentName).Execute()

Get an environment



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
    environmentName := "environmentName_example" // string | The name of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetEnvironment(context.Background(), owner, repo, environmentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**environmentName** | **string** | The name of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Environment**](Environment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetLatestPagesBuild

> PageBuild ReposGetLatestPagesBuild(ctx, owner, repo).Execute()

Get latest Pages build



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
    resp, r, err := apiClient.ReposApi.ReposGetLatestPagesBuild(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetLatestPagesBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetLatestPagesBuild`: PageBuild
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetLatestPagesBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetLatestPagesBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PageBuild**](PageBuild.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetLatestRelease

> Release ReposGetLatestRelease(ctx, owner, repo).Execute()

Get the latest release



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
    resp, r, err := apiClient.ReposApi.ReposGetLatestRelease(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetLatestRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetLatestRelease`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetLatestRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetLatestReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetPages

> Page ReposGetPages(ctx, owner, repo).Execute()

Get a GitHub Pages site



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
    resp, r, err := apiClient.ReposApi.ReposGetPages(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetPages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetPages`: Page
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Page**](Page.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetPagesBuild

> PageBuild ReposGetPagesBuild(ctx, owner, repo, buildId).Execute()

Get GitHub Pages build



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
    buildId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetPagesBuild(context.Background(), owner, repo, buildId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetPagesBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetPagesBuild`: PageBuild
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetPagesBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**buildId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetPagesBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PageBuild**](PageBuild.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetPagesHealthCheck

> PagesHealthCheck ReposGetPagesHealthCheck(ctx, owner, repo).Execute()

Get a DNS health check for GitHub Pages



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
    resp, r, err := apiClient.ReposApi.ReposGetPagesHealthCheck(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetPagesHealthCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetPagesHealthCheck`: PagesHealthCheck
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetPagesHealthCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetPagesHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PagesHealthCheck**](PagesHealthCheck.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetParticipationStats

> ParticipationStats ReposGetParticipationStats(ctx, owner, repo).Execute()

Get the weekly commit count



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
    resp, r, err := apiClient.ReposApi.ReposGetParticipationStats(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetParticipationStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetParticipationStats`: ParticipationStats
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetParticipationStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetParticipationStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ParticipationStats**](ParticipationStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetPullRequestReviewProtection

> ProtectedBranchPullRequestReview ReposGetPullRequestReviewProtection(ctx, owner, repo, branch).Execute()

Get pull request review protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetPullRequestReviewProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetPullRequestReviewProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetPullRequestReviewProtection`: ProtectedBranchPullRequestReview
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetPullRequestReviewProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetPullRequestReviewProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProtectedBranchPullRequestReview**](ProtectedBranchPullRequestReview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetPunchCardStats

> [][]int32 ReposGetPunchCardStats(ctx, owner, repo).Execute()

Get the hourly commit count for each day



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
    resp, r, err := apiClient.ReposApi.ReposGetPunchCardStats(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetPunchCardStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetPunchCardStats`: [][]int32
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetPunchCardStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetPunchCardStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[][]int32**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetReadme

> ContentFile ReposGetReadme(ctx, owner, repo).Ref(ref).Execute()

Get a repository README



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
    ref := "ref_example" // string | The name of the commit/branch/tag. Default: the repository’s default branch (usually `master`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetReadme(context.Background(), owner, repo).Ref(ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetReadme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetReadme`: ContentFile
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetReadme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetReadmeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ref** | **string** | The name of the commit/branch/tag. Default: the repository’s default branch (usually &#x60;master&#x60;) | 

### Return type

[**ContentFile**](ContentFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetReadmeInDirectory

> ContentFile ReposGetReadmeInDirectory(ctx, owner, repo, dir).Ref(ref).Execute()

Get a repository README for a directory



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
    dir := "dir_example" // string | The alternate path to look for a README file
    ref := "ref_example" // string | The name of the commit/branch/tag. Default: the repository’s default branch (usually `master`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetReadmeInDirectory(context.Background(), owner, repo, dir).Ref(ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetReadmeInDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetReadmeInDirectory`: ContentFile
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetReadmeInDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**dir** | **string** | The alternate path to look for a README file | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetReadmeInDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ref** | **string** | The name of the commit/branch/tag. Default: the repository’s default branch (usually &#x60;master&#x60;) | 

### Return type

[**ContentFile**](ContentFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetRelease

> Release ReposGetRelease(ctx, owner, repo, releaseId).Execute()

Get a release



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
    releaseId := int32(56) // int32 | The unique identifier of the release.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetRelease(context.Background(), owner, repo, releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetRelease`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**releaseId** | **int32** | The unique identifier of the release. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetReleaseAsset

> ReleaseAsset ReposGetReleaseAsset(ctx, owner, repo, assetId).Execute()

Get a release asset



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
    assetId := int32(56) // int32 | The unique identifier of the asset.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetReleaseAsset(context.Background(), owner, repo, assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetReleaseAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetReleaseAsset`: ReleaseAsset
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetReleaseAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**assetId** | **int32** | The unique identifier of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetReleaseAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ReleaseAsset**](ReleaseAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetReleaseByTag

> Release ReposGetReleaseByTag(ctx, owner, repo, tag).Execute()

Get a release by tag name



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
    tag := "tag_example" // string | tag parameter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetReleaseByTag(context.Background(), owner, repo, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetReleaseByTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetReleaseByTag`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetReleaseByTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**tag** | **string** | tag parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetReleaseByTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetStatusChecksProtection

> StatusCheckPolicy ReposGetStatusChecksProtection(ctx, owner, repo, branch).Execute()

Get status checks protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetStatusChecksProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetStatusChecksProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetStatusChecksProtection`: StatusCheckPolicy
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetStatusChecksProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetStatusChecksProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StatusCheckPolicy**](StatusCheckPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetTeamsWithAccessToProtectedBranch

> []Team ReposGetTeamsWithAccessToProtectedBranch(ctx, owner, repo, branch).Execute()

Get teams with access to the protected branch



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetTeamsWithAccessToProtectedBranch(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetTeamsWithAccessToProtectedBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetTeamsWithAccessToProtectedBranch`: []Team
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetTeamsWithAccessToProtectedBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetTeamsWithAccessToProtectedBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetTopPaths

> []ContentTraffic ReposGetTopPaths(ctx, owner, repo).Execute()

Get top referral paths



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
    resp, r, err := apiClient.ReposApi.ReposGetTopPaths(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetTopPaths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetTopPaths`: []ContentTraffic
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetTopPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetTopPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ContentTraffic**](ContentTraffic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetTopReferrers

> []ReferrerTraffic ReposGetTopReferrers(ctx, owner, repo).Execute()

Get top referral sources



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
    resp, r, err := apiClient.ReposApi.ReposGetTopReferrers(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetTopReferrers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetTopReferrers`: []ReferrerTraffic
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetTopReferrers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetTopReferrersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ReferrerTraffic**](ReferrerTraffic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetUsersWithAccessToProtectedBranch

> []SimpleUser ReposGetUsersWithAccessToProtectedBranch(ctx, owner, repo, branch).Execute()

Get users with access to the protected branch



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetUsersWithAccessToProtectedBranch(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetUsersWithAccessToProtectedBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetUsersWithAccessToProtectedBranch`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetUsersWithAccessToProtectedBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetUsersWithAccessToProtectedBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]SimpleUser**](SimpleUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetViews

> ViewTraffic ReposGetViews(ctx, owner, repo).Per(per).Execute()

Get page views



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
    per := "per_example" // string | The time frame to display results for. (optional) (default to "day")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetViews(context.Background(), owner, repo).Per(per).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetViews`: ViewTraffic
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **per** | **string** | The time frame to display results for. | [default to &quot;day&quot;]

### Return type

[**ViewTraffic**](ViewTraffic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetWebhook

> Hook ReposGetWebhook(ctx, owner, repo, hookId).Execute()

Get a repository webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetWebhook(context.Background(), owner, repo, hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetWebhook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetWebhookConfigForRepo

> WebhookConfig ReposGetWebhookConfigForRepo(ctx, owner, repo, hookId).Execute()

Get a webhook configuration for a repository



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
    hookId := int32(56) // int32 | The unique identifier of the hook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetWebhookConfigForRepo(context.Background(), owner, repo, hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetWebhookConfigForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetWebhookConfigForRepo`: WebhookConfig
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetWebhookConfigForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetWebhookConfigForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WebhookConfig**](WebhookConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGetWebhookDelivery

> HookDelivery ReposGetWebhookDelivery(ctx, owner, repo, hookId, deliveryId).Execute()

Get a delivery for a repository webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.
    deliveryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGetWebhookDelivery(context.Background(), owner, repo, hookId, deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGetWebhookDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGetWebhookDelivery`: HookDelivery
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGetWebhookDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 
**deliveryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGetWebhookDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**HookDelivery**](HookDelivery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListAutolinks

> []Autolink ReposListAutolinks(ctx, owner, repo).Page(page).Execute()

List all autolinks of a repository



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
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListAutolinks(context.Background(), owner, repo).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListAutolinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListAutolinks`: []Autolink
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListAutolinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListAutolinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Autolink**](Autolink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListBranches

> []ShortBranch ReposListBranches(ctx, owner, repo).Protected(protected).PerPage(perPage).Page(page).Execute()

List branches



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
    protected := true // bool | Setting to `true` returns only protected branches. When set to `false`, only unprotected branches are returned. Omitting this parameter returns all branches. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListBranches(context.Background(), owner, repo).Protected(protected).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListBranches`: []ShortBranch
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **protected** | **bool** | Setting to &#x60;true&#x60; returns only protected branches. When set to &#x60;false&#x60;, only unprotected branches are returned. Omitting this parameter returns all branches. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]ShortBranch**](ShortBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListBranchesForHeadCommit

> []BranchShort ReposListBranchesForHeadCommit(ctx, owner, repo, commitSha).Execute()

List branches for HEAD commit



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
    commitSha := "commitSha_example" // string | The SHA of the commit.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListBranchesForHeadCommit(context.Background(), owner, repo, commitSha).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListBranchesForHeadCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListBranchesForHeadCommit`: []BranchShort
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListBranchesForHeadCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**commitSha** | **string** | The SHA of the commit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListBranchesForHeadCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]BranchShort**](BranchShort.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListCollaborators

> []Collaborator ReposListCollaborators(ctx, owner, repo).Affiliation(affiliation).PerPage(perPage).Page(page).Execute()

List repository collaborators



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
    affiliation := "affiliation_example" // string | Filter collaborators returned by their affiliation. `outside` means all outside collaborators of an organization-owned repository. `direct` means all collaborators with permissions to an organization-owned repository, regardless of organization membership status. `all` means all collaborators the authenticated user can see. (optional) (default to "all")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListCollaborators(context.Background(), owner, repo).Affiliation(affiliation).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListCollaborators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListCollaborators`: []Collaborator
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListCollaborators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListCollaboratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **affiliation** | **string** | Filter collaborators returned by their affiliation. &#x60;outside&#x60; means all outside collaborators of an organization-owned repository. &#x60;direct&#x60; means all collaborators with permissions to an organization-owned repository, regardless of organization membership status. &#x60;all&#x60; means all collaborators the authenticated user can see. | [default to &quot;all&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Collaborator**](Collaborator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListCommentsForCommit

> []CommitComment ReposListCommentsForCommit(ctx, owner, repo, commitSha).PerPage(perPage).Page(page).Execute()

List commit comments



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
    commitSha := "commitSha_example" // string | The SHA of the commit.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListCommentsForCommit(context.Background(), owner, repo, commitSha).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListCommentsForCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListCommentsForCommit`: []CommitComment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListCommentsForCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**commitSha** | **string** | The SHA of the commit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListCommentsForCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]CommitComment**](CommitComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListCommitCommentsForRepo

> []CommitComment ReposListCommitCommentsForRepo(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List commit comments for a repository



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
    resp, r, err := apiClient.ReposApi.ReposListCommitCommentsForRepo(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListCommitCommentsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListCommitCommentsForRepo`: []CommitComment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListCommitCommentsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListCommitCommentsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]CommitComment**](CommitComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListCommitStatusesForRef

> []Status ReposListCommitStatusesForRef(ctx, owner, repo, ref).PerPage(perPage).Page(page).Execute()

List commit statuses for a reference



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
    ref := "ref_example" // string | ref parameter
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListCommitStatusesForRef(context.Background(), owner, repo, ref).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListCommitStatusesForRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListCommitStatusesForRef`: []Status
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListCommitStatusesForRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**ref** | **string** | ref parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListCommitStatusesForRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Status**](Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListCommits

> []Commit ReposListCommits(ctx, owner, repo).Sha(sha).Path(path).Author(author).Since(since).Until(until).PerPage(perPage).Page(page).Execute()

List commits



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
    sha := "sha_example" // string | SHA or branch to start listing commits from. Default: the repository’s default branch (usually `master`). (optional)
    path := "path_example" // string | Only commits containing this file path will be returned. (optional)
    author := "author_example" // string | GitHub login or email address by which to filter by commit author. (optional)
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    until := time.Now() // time.Time | Only commits before this date will be returned. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListCommits(context.Background(), owner, repo).Sha(sha).Path(path).Author(author).Since(since).Until(until).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListCommits`: []Commit
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sha** | **string** | SHA or branch to start listing commits from. Default: the repository’s default branch (usually &#x60;master&#x60;). | 
 **path** | **string** | Only commits containing this file path will be returned. | 
 **author** | **string** | GitHub login or email address by which to filter by commit author. | 
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **until** | **time.Time** | Only commits before this date will be returned. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Commit**](Commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListContributors

> []Contributor ReposListContributors(ctx, owner, repo).Anon(anon).PerPage(perPage).Page(page).Execute()

List repository contributors



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
    anon := "anon_example" // string | Set to `1` or `true` to include anonymous contributors in results. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListContributors(context.Background(), owner, repo).Anon(anon).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListContributors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListContributors`: []Contributor
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListContributors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **anon** | **string** | Set to &#x60;1&#x60; or &#x60;true&#x60; to include anonymous contributors in results. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Contributor**](Contributor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListDeployKeys

> []DeployKey ReposListDeployKeys(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List deploy keys



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
    resp, r, err := apiClient.ReposApi.ReposListDeployKeys(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListDeployKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListDeployKeys`: []DeployKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListDeployKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListDeployKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]DeployKey**](DeployKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListDeploymentBranchPolicies

> ReposListDeploymentBranchPolicies200Response ReposListDeploymentBranchPolicies(ctx, owner, repo, environmentName).PerPage(perPage).Page(page).Execute()

List deployment branch policies



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
    environmentName := "environmentName_example" // string | The name of the environment.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListDeploymentBranchPolicies(context.Background(), owner, repo, environmentName).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListDeploymentBranchPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListDeploymentBranchPolicies`: ReposListDeploymentBranchPolicies200Response
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListDeploymentBranchPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**environmentName** | **string** | The name of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListDeploymentBranchPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ReposListDeploymentBranchPolicies200Response**](ReposListDeploymentBranchPolicies200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListDeploymentStatuses

> []DeploymentStatus ReposListDeploymentStatuses(ctx, owner, repo, deploymentId).PerPage(perPage).Page(page).Execute()

List deployment statuses



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
    deploymentId := int32(56) // int32 | deployment_id parameter
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListDeploymentStatuses(context.Background(), owner, repo, deploymentId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListDeploymentStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListDeploymentStatuses`: []DeploymentStatus
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListDeploymentStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**deploymentId** | **int32** | deployment_id parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListDeploymentStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]DeploymentStatus**](DeploymentStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListDeployments

> []Deployment ReposListDeployments(ctx, owner, repo).Sha(sha).Ref(ref).Task(task).Environment(environment).PerPage(perPage).Page(page).Execute()

List deployments



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
    sha := "sha_example" // string | The SHA recorded at creation time. (optional) (default to "none")
    ref := "ref_example" // string | The name of the ref. This can be a branch, tag, or SHA. (optional) (default to "none")
    task := "task_example" // string | The name of the task for the deployment (e.g., `deploy` or `deploy:migrations`). (optional) (default to "none")
    environment := "environment_example" // string | The name of the environment that was deployed to (e.g., `staging` or `production`). (optional) (default to "none")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListDeployments(context.Background(), owner, repo).Sha(sha).Ref(ref).Task(task).Environment(environment).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListDeployments`: []Deployment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sha** | **string** | The SHA recorded at creation time. | [default to &quot;none&quot;]
 **ref** | **string** | The name of the ref. This can be a branch, tag, or SHA. | [default to &quot;none&quot;]
 **task** | **string** | The name of the task for the deployment (e.g., &#x60;deploy&#x60; or &#x60;deploy:migrations&#x60;). | [default to &quot;none&quot;]
 **environment** | **string** | The name of the environment that was deployed to (e.g., &#x60;staging&#x60; or &#x60;production&#x60;). | [default to &quot;none&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListForAuthenticatedUser

> []Repository ReposListForAuthenticatedUser(ctx).Visibility(visibility).Affiliation(affiliation).Type_(type_).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Since(since).Before(before).Execute()

List repositories for the authenticated user



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
    visibility := "visibility_example" // string | Limit results to repositories with the specified visibility. (optional) (default to "all")
    affiliation := "affiliation_example" // string | Comma-separated list of values. Can include:   \\* `owner`: Repositories that are owned by the authenticated user.   \\* `collaborator`: Repositories that the user has been added to as a collaborator.   \\* `organization_member`: Repositories that the user has access to through being a member of an organization. This includes every repository on every team that the user is on. (optional) (default to "owner,collaborator,organization_member")
    type_ := "type__example" // string | Limit results to repositories of the specified type. Will cause a `422` error if used in the same request as **visibility** or **affiliation**. (optional) (default to "all")
    sort := "sort_example" // string | The property to sort the results by. (optional) (default to "full_name")
    direction := "direction_example" // string | The order to sort by. Default: `asc` when using `full_name`, otherwise `desc`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    before := time.Now() // time.Time | Only show notifications updated before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListForAuthenticatedUser(context.Background()).Visibility(visibility).Affiliation(affiliation).Type_(type_).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Since(since).Before(before).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListForAuthenticatedUser`: []Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReposListForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visibility** | **string** | Limit results to repositories with the specified visibility. | [default to &quot;all&quot;]
 **affiliation** | **string** | Comma-separated list of values. Can include:   \\* &#x60;owner&#x60;: Repositories that are owned by the authenticated user.   \\* &#x60;collaborator&#x60;: Repositories that the user has been added to as a collaborator.   \\* &#x60;organization_member&#x60;: Repositories that the user has access to through being a member of an organization. This includes every repository on every team that the user is on. | [default to &quot;owner,collaborator,organization_member&quot;]
 **type_** | **string** | Limit results to repositories of the specified type. Will cause a &#x60;422&#x60; error if used in the same request as **visibility** or **affiliation**. | [default to &quot;all&quot;]
 **sort** | **string** | The property to sort the results by. | [default to &quot;full_name&quot;]
 **direction** | **string** | The order to sort by. Default: &#x60;asc&#x60; when using &#x60;full_name&#x60;, otherwise &#x60;desc&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **before** | **time.Time** | Only show notifications updated before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListForOrg

> []MinimalRepository ReposListForOrg(ctx, org).Type_(type_).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()

List organization repositories



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
    type_ := "type__example" // string | Specifies the types of repositories you want returned. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, `type` can also be `internal`. However, the `internal` value is not yet supported when a GitHub App calls this API with an installation access token. (optional)
    sort := "sort_example" // string | The property to sort the results by. (optional) (default to "created")
    direction := "direction_example" // string | The order to sort by. Default: `asc` when using `full_name`, otherwise `desc`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListForOrg(context.Background(), org).Type_(type_).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListForOrg`: []MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Specifies the types of repositories you want returned. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, &#x60;type&#x60; can also be &#x60;internal&#x60;. However, the &#x60;internal&#x60; value is not yet supported when a GitHub App calls this API with an installation access token. | 
 **sort** | **string** | The property to sort the results by. | [default to &quot;created&quot;]
 **direction** | **string** | The order to sort by. Default: &#x60;asc&#x60; when using &#x60;full_name&#x60;, otherwise &#x60;desc&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]MinimalRepository**](MinimalRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListForUser

> []MinimalRepository ReposListForUser(ctx, username).Type_(type_).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()

List repositories for a user



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
    username := "username_example" // string | The handle for the GitHub user account.
    type_ := "type__example" // string | Limit results to repositories of the specified type. (optional) (default to "owner")
    sort := "sort_example" // string | The property to sort the results by. (optional) (default to "full_name")
    direction := "direction_example" // string | The order to sort by. Default: `asc` when using `full_name`, otherwise `desc`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListForUser(context.Background(), username).Type_(type_).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListForUser`: []MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Limit results to repositories of the specified type. | [default to &quot;owner&quot;]
 **sort** | **string** | The property to sort the results by. | [default to &quot;full_name&quot;]
 **direction** | **string** | The order to sort by. Default: &#x60;asc&#x60; when using &#x60;full_name&#x60;, otherwise &#x60;desc&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]MinimalRepository**](MinimalRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListForks

> []MinimalRepository ReposListForks(ctx, owner, repo).Sort(sort).PerPage(perPage).Page(page).Execute()

List forks



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
    sort := "sort_example" // string | The sort order. Can be either `newest`, `oldest`, or `stargazers`. (optional) (default to "newest")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListForks(context.Background(), owner, repo).Sort(sort).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListForks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListForks`: []MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListForks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListForksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **string** | The sort order. Can be either &#x60;newest&#x60;, &#x60;oldest&#x60;, or &#x60;stargazers&#x60;. | [default to &quot;newest&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]MinimalRepository**](MinimalRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListInvitations

> []RepositoryInvitation ReposListInvitations(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List repository invitations



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
    resp, r, err := apiClient.ReposApi.ReposListInvitations(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListInvitations`: []RepositoryInvitation
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]RepositoryInvitation**](RepositoryInvitation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListInvitationsForAuthenticatedUser

> []RepositoryInvitation ReposListInvitationsForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List repository invitations for the authenticated user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListInvitationsForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListInvitationsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListInvitationsForAuthenticatedUser`: []RepositoryInvitation
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListInvitationsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReposListInvitationsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]RepositoryInvitation**](RepositoryInvitation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListLanguages

> map[string]int32 ReposListLanguages(ctx, owner, repo).Execute()

List repository languages



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
    resp, r, err := apiClient.ReposApi.ReposListLanguages(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListLanguages`: map[string]int32
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListLanguages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListPagesBuilds

> []PageBuild ReposListPagesBuilds(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List GitHub Pages builds



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
    resp, r, err := apiClient.ReposApi.ReposListPagesBuilds(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListPagesBuilds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListPagesBuilds`: []PageBuild
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListPagesBuilds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListPagesBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]PageBuild**](PageBuild.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListPublic

> []MinimalRepository ReposListPublic(ctx).Since(since).Execute()

List public repositories



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
    since := int32(56) // int32 | A repository ID. Only return repositories with an ID greater than this ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListPublic(context.Background()).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListPublic`: []MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReposListPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **int32** | A repository ID. Only return repositories with an ID greater than this ID. | 

### Return type

[**[]MinimalRepository**](MinimalRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListPullRequestsAssociatedWithCommit

> []PullRequestSimple ReposListPullRequestsAssociatedWithCommit(ctx, owner, repo, commitSha).PerPage(perPage).Page(page).Execute()

List pull requests associated with a commit



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
    commitSha := "commitSha_example" // string | The SHA of the commit.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListPullRequestsAssociatedWithCommit(context.Background(), owner, repo, commitSha).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListPullRequestsAssociatedWithCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListPullRequestsAssociatedWithCommit`: []PullRequestSimple
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListPullRequestsAssociatedWithCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**commitSha** | **string** | The SHA of the commit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListPullRequestsAssociatedWithCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]PullRequestSimple**](PullRequestSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListReleaseAssets

> []ReleaseAsset ReposListReleaseAssets(ctx, owner, repo, releaseId).PerPage(perPage).Page(page).Execute()

List release assets



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
    releaseId := int32(56) // int32 | The unique identifier of the release.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListReleaseAssets(context.Background(), owner, repo, releaseId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListReleaseAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListReleaseAssets`: []ReleaseAsset
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListReleaseAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**releaseId** | **int32** | The unique identifier of the release. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListReleaseAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]ReleaseAsset**](ReleaseAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListReleases

> []Release ReposListReleases(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List releases



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
    resp, r, err := apiClient.ReposApi.ReposListReleases(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListReleases`: []Release
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListTagProtection

> []TagProtection ReposListTagProtection(ctx, owner, repo).Execute()

List tag protection states for a repository



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
    resp, r, err := apiClient.ReposApi.ReposListTagProtection(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListTagProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListTagProtection`: []TagProtection
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListTagProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListTagProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TagProtection**](TagProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListTags

> []Tag ReposListTags(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List repository tags



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
    resp, r, err := apiClient.ReposApi.ReposListTags(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListTags`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListTeams

> []Team ReposListTeams(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List repository teams



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
    resp, r, err := apiClient.ReposApi.ReposListTeams(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListTeams`: []Team
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListWebhookDeliveries

> []HookDeliveryItem ReposListWebhookDeliveries(ctx, owner, repo, hookId).PerPage(perPage).Cursor(cursor).Execute()

List deliveries for a repository webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    cursor := "cursor_example" // string | Used for pagination: the starting delivery from which the page of deliveries is fetched. Refer to the `link` header for the next and previous page cursors. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposListWebhookDeliveries(context.Background(), owner, repo, hookId).PerPage(perPage).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListWebhookDeliveries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListWebhookDeliveries`: []HookDeliveryItem
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListWebhookDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListWebhookDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **cursor** | **string** | Used for pagination: the starting delivery from which the page of deliveries is fetched. Refer to the &#x60;link&#x60; header for the next and previous page cursors. | 

### Return type

[**[]HookDeliveryItem**](HookDeliveryItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposListWebhooks

> []Hook ReposListWebhooks(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List repository webhooks



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
    resp, r, err := apiClient.ReposApi.ReposListWebhooks(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposListWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposListWebhooks`: []Hook
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposListWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposMerge

> Commit ReposMerge(ctx, owner, repo).ReposMergeRequest(reposMergeRequest).Execute()

Merge a branch



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
    reposMergeRequest := *openapiclient.NewReposMergeRequest("Base_example", "Head_example") // ReposMergeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposMerge(context.Background(), owner, repo).ReposMergeRequest(reposMergeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposMerge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposMerge`: Commit
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposMerge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposMergeRequest** | [**ReposMergeRequest**](ReposMergeRequest.md) |  | 

### Return type

[**Commit**](Commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposMergeUpstream

> MergedUpstream ReposMergeUpstream(ctx, owner, repo).ReposMergeUpstreamRequest(reposMergeUpstreamRequest).Execute()

Sync a fork branch with the upstream repository



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
    reposMergeUpstreamRequest := *openapiclient.NewReposMergeUpstreamRequest("Branch_example") // ReposMergeUpstreamRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposMergeUpstream(context.Background(), owner, repo).ReposMergeUpstreamRequest(reposMergeUpstreamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposMergeUpstream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposMergeUpstream`: MergedUpstream
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposMergeUpstream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposMergeUpstreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposMergeUpstreamRequest** | [**ReposMergeUpstreamRequest**](ReposMergeUpstreamRequest.md) |  | 

### Return type

[**MergedUpstream**](MergedUpstream.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPingWebhook

> ReposPingWebhook(ctx, owner, repo, hookId).Execute()

Ping a repository webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposPingWebhook(context.Background(), owner, repo, hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPingWebhook``: %v\n", err)
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
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposPingWebhookRequest struct via the builder pattern


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


## ReposRedeliverWebhookDelivery

> map[string]interface{} ReposRedeliverWebhookDelivery(ctx, owner, repo, hookId, deliveryId).Execute()

Redeliver a delivery for a repository webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.
    deliveryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRedeliverWebhookDelivery(context.Background(), owner, repo, hookId, deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRedeliverWebhookDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRedeliverWebhookDelivery`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRedeliverWebhookDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 
**deliveryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRedeliverWebhookDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRemoveAppAccessRestrictions

> []Integration ReposRemoveAppAccessRestrictions(ctx, owner, repo, branch).ReposSetAppAccessRestrictionsRequest(reposSetAppAccessRestrictionsRequest).Execute()

Remove app access restrictions



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
    branch := "branch_example" // string | The name of the branch.
    reposSetAppAccessRestrictionsRequest := openapiclient.repos_set_app_access_restrictions_request{ReposSetAppAccessRestrictionsRequestOneOf: openapiclient.NewReposSetAppAccessRestrictionsRequestOneOf([]string{"Apps_example"})} // ReposSetAppAccessRestrictionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRemoveAppAccessRestrictions(context.Background(), owner, repo, branch).ReposSetAppAccessRestrictionsRequest(reposSetAppAccessRestrictionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRemoveAppAccessRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRemoveAppAccessRestrictions`: []Integration
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRemoveAppAccessRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRemoveAppAccessRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposSetAppAccessRestrictionsRequest** | [**ReposSetAppAccessRestrictionsRequest**](ReposSetAppAccessRestrictionsRequest.md) |  | 

### Return type

[**[]Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRemoveCollaborator

> ReposRemoveCollaborator(ctx, owner, repo, username).Execute()

Remove a repository collaborator



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
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRemoveCollaborator(context.Background(), owner, repo, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRemoveCollaborator``: %v\n", err)
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
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRemoveCollaboratorRequest struct via the builder pattern


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


## ReposRemoveStatusCheckContexts

> []string ReposRemoveStatusCheckContexts(ctx, owner, repo, branch).ReposSetStatusCheckContextsRequest(reposSetStatusCheckContextsRequest).Execute()

Remove status check contexts



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
    branch := "branch_example" // string | The name of the branch.
    reposSetStatusCheckContextsRequest := openapiclient.repos_set_status_check_contexts_request{ReposSetStatusCheckContextsRequestOneOf: openapiclient.NewReposSetStatusCheckContextsRequestOneOf([]string{"Contexts_example"})} // ReposSetStatusCheckContextsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRemoveStatusCheckContexts(context.Background(), owner, repo, branch).ReposSetStatusCheckContextsRequest(reposSetStatusCheckContextsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRemoveStatusCheckContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRemoveStatusCheckContexts`: []string
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRemoveStatusCheckContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRemoveStatusCheckContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposSetStatusCheckContextsRequest** | [**ReposSetStatusCheckContextsRequest**](ReposSetStatusCheckContextsRequest.md) |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRemoveStatusCheckProtection

> ReposRemoveStatusCheckProtection(ctx, owner, repo, branch).Execute()

Remove status check protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRemoveStatusCheckProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRemoveStatusCheckProtection``: %v\n", err)
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
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRemoveStatusCheckProtectionRequest struct via the builder pattern


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


## ReposRemoveTeamAccessRestrictions

> []Team ReposRemoveTeamAccessRestrictions(ctx, owner, repo, branch).ReposAddTeamAccessRestrictionsRequest(reposAddTeamAccessRestrictionsRequest).Execute()

Remove team access restrictions



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
    branch := "branch_example" // string | The name of the branch.
    reposAddTeamAccessRestrictionsRequest := openapiclient.repos_add_team_access_restrictions_request{ReposAddTeamAccessRestrictionsRequestOneOf: openapiclient.NewReposAddTeamAccessRestrictionsRequestOneOf([]string{"Teams_example"})} // ReposAddTeamAccessRestrictionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRemoveTeamAccessRestrictions(context.Background(), owner, repo, branch).ReposAddTeamAccessRestrictionsRequest(reposAddTeamAccessRestrictionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRemoveTeamAccessRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRemoveTeamAccessRestrictions`: []Team
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRemoveTeamAccessRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRemoveTeamAccessRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposAddTeamAccessRestrictionsRequest** | [**ReposAddTeamAccessRestrictionsRequest**](ReposAddTeamAccessRestrictionsRequest.md) |  | 

### Return type

[**[]Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRemoveUserAccessRestrictions

> []SimpleUser ReposRemoveUserAccessRestrictions(ctx, owner, repo, branch).ReposSetUserAccessRestrictionsRequest(reposSetUserAccessRestrictionsRequest).Execute()

Remove user access restrictions



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
    branch := "branch_example" // string | The name of the branch.
    reposSetUserAccessRestrictionsRequest := openapiclient.repos_set_user_access_restrictions_request{ReposSetUserAccessRestrictionsRequestOneOf: openapiclient.NewReposSetUserAccessRestrictionsRequestOneOf([]string{"Users_example"})} // ReposSetUserAccessRestrictionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRemoveUserAccessRestrictions(context.Background(), owner, repo, branch).ReposSetUserAccessRestrictionsRequest(reposSetUserAccessRestrictionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRemoveUserAccessRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRemoveUserAccessRestrictions`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRemoveUserAccessRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRemoveUserAccessRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposSetUserAccessRestrictionsRequest** | [**ReposSetUserAccessRestrictionsRequest**](ReposSetUserAccessRestrictionsRequest.md) |  | 

### Return type

[**[]SimpleUser**](SimpleUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRenameBranch

> BranchWithProtection ReposRenameBranch(ctx, owner, repo, branch).ReposRenameBranchRequest(reposRenameBranchRequest).Execute()

Rename a branch



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
    branch := "branch_example" // string | The name of the branch.
    reposRenameBranchRequest := *openapiclient.NewReposRenameBranchRequest("NewName_example") // ReposRenameBranchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRenameBranch(context.Background(), owner, repo, branch).ReposRenameBranchRequest(reposRenameBranchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRenameBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRenameBranch`: BranchWithProtection
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRenameBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRenameBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposRenameBranchRequest** | [**ReposRenameBranchRequest**](ReposRenameBranchRequest.md) |  | 

### Return type

[**BranchWithProtection**](BranchWithProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposReplaceAllTopics

> Topic ReposReplaceAllTopics(ctx, owner, repo).ReposReplaceAllTopicsRequest(reposReplaceAllTopicsRequest).Execute()

Replace all repository topics



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
    reposReplaceAllTopicsRequest := *openapiclient.NewReposReplaceAllTopicsRequest([]string{"Names_example"}) // ReposReplaceAllTopicsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposReplaceAllTopics(context.Background(), owner, repo).ReposReplaceAllTopicsRequest(reposReplaceAllTopicsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposReplaceAllTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposReplaceAllTopics`: Topic
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposReplaceAllTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposReplaceAllTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposReplaceAllTopicsRequest** | [**ReposReplaceAllTopicsRequest**](ReposReplaceAllTopicsRequest.md) |  | 

### Return type

[**Topic**](Topic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRequestPagesBuild

> PageBuildStatus ReposRequestPagesBuild(ctx, owner, repo).Execute()

Request a GitHub Pages build



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
    resp, r, err := apiClient.ReposApi.ReposRequestPagesBuild(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRequestPagesBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRequestPagesBuild`: PageBuildStatus
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRequestPagesBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRequestPagesBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PageBuildStatus**](PageBuildStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposSetAdminBranchProtection

> ProtectedBranchAdminEnforced ReposSetAdminBranchProtection(ctx, owner, repo, branch).Execute()

Set admin branch protection



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
    branch := "branch_example" // string | The name of the branch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposSetAdminBranchProtection(context.Background(), owner, repo, branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposSetAdminBranchProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposSetAdminBranchProtection`: ProtectedBranchAdminEnforced
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposSetAdminBranchProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposSetAdminBranchProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProtectedBranchAdminEnforced**](ProtectedBranchAdminEnforced.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposSetAppAccessRestrictions

> []Integration ReposSetAppAccessRestrictions(ctx, owner, repo, branch).ReposSetAppAccessRestrictionsRequest(reposSetAppAccessRestrictionsRequest).Execute()

Set app access restrictions



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
    branch := "branch_example" // string | The name of the branch.
    reposSetAppAccessRestrictionsRequest := openapiclient.repos_set_app_access_restrictions_request{ReposSetAppAccessRestrictionsRequestOneOf: openapiclient.NewReposSetAppAccessRestrictionsRequestOneOf([]string{"Apps_example"})} // ReposSetAppAccessRestrictionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposSetAppAccessRestrictions(context.Background(), owner, repo, branch).ReposSetAppAccessRestrictionsRequest(reposSetAppAccessRestrictionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposSetAppAccessRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposSetAppAccessRestrictions`: []Integration
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposSetAppAccessRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposSetAppAccessRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposSetAppAccessRestrictionsRequest** | [**ReposSetAppAccessRestrictionsRequest**](ReposSetAppAccessRestrictionsRequest.md) |  | 

### Return type

[**[]Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposSetStatusCheckContexts

> []string ReposSetStatusCheckContexts(ctx, owner, repo, branch).ReposSetStatusCheckContextsRequest(reposSetStatusCheckContextsRequest).Execute()

Set status check contexts



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
    branch := "branch_example" // string | The name of the branch.
    reposSetStatusCheckContextsRequest := openapiclient.repos_set_status_check_contexts_request{ReposSetStatusCheckContextsRequestOneOf: openapiclient.NewReposSetStatusCheckContextsRequestOneOf([]string{"Contexts_example"})} // ReposSetStatusCheckContextsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposSetStatusCheckContexts(context.Background(), owner, repo, branch).ReposSetStatusCheckContextsRequest(reposSetStatusCheckContextsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposSetStatusCheckContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposSetStatusCheckContexts`: []string
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposSetStatusCheckContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposSetStatusCheckContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposSetStatusCheckContextsRequest** | [**ReposSetStatusCheckContextsRequest**](ReposSetStatusCheckContextsRequest.md) |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposSetTeamAccessRestrictions

> []Team ReposSetTeamAccessRestrictions(ctx, owner, repo, branch).ReposSetTeamAccessRestrictionsRequest(reposSetTeamAccessRestrictionsRequest).Execute()

Set team access restrictions



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
    branch := "branch_example" // string | The name of the branch.
    reposSetTeamAccessRestrictionsRequest := openapiclient.repos_set_team_access_restrictions_request{ReposSetTeamAccessRestrictionsRequestOneOf: openapiclient.NewReposSetTeamAccessRestrictionsRequestOneOf([]string{"Teams_example"})} // ReposSetTeamAccessRestrictionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposSetTeamAccessRestrictions(context.Background(), owner, repo, branch).ReposSetTeamAccessRestrictionsRequest(reposSetTeamAccessRestrictionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposSetTeamAccessRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposSetTeamAccessRestrictions`: []Team
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposSetTeamAccessRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposSetTeamAccessRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposSetTeamAccessRestrictionsRequest** | [**ReposSetTeamAccessRestrictionsRequest**](ReposSetTeamAccessRestrictionsRequest.md) |  | 

### Return type

[**[]Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposSetUserAccessRestrictions

> []SimpleUser ReposSetUserAccessRestrictions(ctx, owner, repo, branch).ReposSetUserAccessRestrictionsRequest(reposSetUserAccessRestrictionsRequest).Execute()

Set user access restrictions



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
    branch := "branch_example" // string | The name of the branch.
    reposSetUserAccessRestrictionsRequest := openapiclient.repos_set_user_access_restrictions_request{ReposSetUserAccessRestrictionsRequestOneOf: openapiclient.NewReposSetUserAccessRestrictionsRequestOneOf([]string{"Users_example"})} // ReposSetUserAccessRestrictionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposSetUserAccessRestrictions(context.Background(), owner, repo, branch).ReposSetUserAccessRestrictionsRequest(reposSetUserAccessRestrictionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposSetUserAccessRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposSetUserAccessRestrictions`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposSetUserAccessRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposSetUserAccessRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposSetUserAccessRestrictionsRequest** | [**ReposSetUserAccessRestrictionsRequest**](ReposSetUserAccessRestrictionsRequest.md) |  | 

### Return type

[**[]SimpleUser**](SimpleUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposTestPushWebhook

> ReposTestPushWebhook(ctx, owner, repo, hookId).Execute()

Test the push repository webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposTestPushWebhook(context.Background(), owner, repo, hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposTestPushWebhook``: %v\n", err)
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
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposTestPushWebhookRequest struct via the builder pattern


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


## ReposTransfer

> MinimalRepository ReposTransfer(ctx, owner, repo).ReposTransferRequest(reposTransferRequest).Execute()

Transfer a repository



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
    reposTransferRequest := *openapiclient.NewReposTransferRequest("NewOwner_example") // ReposTransferRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposTransfer(context.Background(), owner, repo).ReposTransferRequest(reposTransferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposTransfer`: MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposTransferRequest** | [**ReposTransferRequest**](ReposTransferRequest.md) |  | 

### Return type

[**MinimalRepository**](MinimalRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdate

> FullRepository ReposUpdate(ctx, owner, repo).ReposUpdateRequest(reposUpdateRequest).Execute()

Update a repository



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
    reposUpdateRequest := *openapiclient.NewReposUpdateRequest() // ReposUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdate(context.Background(), owner, repo).ReposUpdateRequest(reposUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdate`: FullRepository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposUpdateRequest** | [**ReposUpdateRequest**](ReposUpdateRequest.md) |  | 

### Return type

[**FullRepository**](FullRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdateBranchProtection

> ProtectedBranch ReposUpdateBranchProtection(ctx, owner, repo, branch).ReposUpdateBranchProtectionRequest(reposUpdateBranchProtectionRequest).Execute()

Update branch protection



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
    branch := "branch_example" // string | The name of the branch.
    reposUpdateBranchProtectionRequest := *openapiclient.NewReposUpdateBranchProtectionRequest("TODO", false, "TODO", "TODO") // ReposUpdateBranchProtectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdateBranchProtection(context.Background(), owner, repo, branch).ReposUpdateBranchProtectionRequest(reposUpdateBranchProtectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdateBranchProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdateBranchProtection`: ProtectedBranch
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdateBranchProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdateBranchProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposUpdateBranchProtectionRequest** | [**ReposUpdateBranchProtectionRequest**](ReposUpdateBranchProtectionRequest.md) |  | 

### Return type

[**ProtectedBranch**](ProtectedBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdateCommitComment

> CommitComment ReposUpdateCommitComment(ctx, owner, repo, commentId).ReposUpdateCommitCommentRequest(reposUpdateCommitCommentRequest).Execute()

Update a commit comment



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
    commentId := int32(56) // int32 | The unique identifier of the comment.
    reposUpdateCommitCommentRequest := *openapiclient.NewReposUpdateCommitCommentRequest("Body_example") // ReposUpdateCommitCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdateCommitComment(context.Background(), owner, repo, commentId).ReposUpdateCommitCommentRequest(reposUpdateCommitCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdateCommitComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdateCommitComment`: CommitComment
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdateCommitComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**commentId** | **int32** | The unique identifier of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdateCommitCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposUpdateCommitCommentRequest** | [**ReposUpdateCommitCommentRequest**](ReposUpdateCommitCommentRequest.md) |  | 

### Return type

[**CommitComment**](CommitComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdateDeploymentBranchPolicy

> DeploymentBranchPolicy ReposUpdateDeploymentBranchPolicy(ctx, owner, repo, environmentName, branchPolicyId).DeploymentBranchPolicyNamePattern(deploymentBranchPolicyNamePattern).Execute()

Update a deployment branch policy



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
    environmentName := "environmentName_example" // string | The name of the environment.
    branchPolicyId := int32(56) // int32 | The unique identifier of the branch policy.
    deploymentBranchPolicyNamePattern := *openapiclient.NewDeploymentBranchPolicyNamePattern("release/*") // DeploymentBranchPolicyNamePattern | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdateDeploymentBranchPolicy(context.Background(), owner, repo, environmentName, branchPolicyId).DeploymentBranchPolicyNamePattern(deploymentBranchPolicyNamePattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdateDeploymentBranchPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdateDeploymentBranchPolicy`: DeploymentBranchPolicy
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdateDeploymentBranchPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**environmentName** | **string** | The name of the environment. | 
**branchPolicyId** | **int32** | The unique identifier of the branch policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdateDeploymentBranchPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **deploymentBranchPolicyNamePattern** | [**DeploymentBranchPolicyNamePattern**](DeploymentBranchPolicyNamePattern.md) |  | 

### Return type

[**DeploymentBranchPolicy**](DeploymentBranchPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdateInformationAboutPagesSite

> ReposUpdateInformationAboutPagesSite(ctx, owner, repo).ReposUpdateInformationAboutPagesSiteRequest(reposUpdateInformationAboutPagesSiteRequest).Execute()

Update information about a GitHub Pages site



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
    reposUpdateInformationAboutPagesSiteRequest := *openapiclient.NewReposUpdateInformationAboutPagesSiteRequest() // ReposUpdateInformationAboutPagesSiteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdateInformationAboutPagesSite(context.Background(), owner, repo).ReposUpdateInformationAboutPagesSiteRequest(reposUpdateInformationAboutPagesSiteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdateInformationAboutPagesSite``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposUpdateInformationAboutPagesSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reposUpdateInformationAboutPagesSiteRequest** | [**ReposUpdateInformationAboutPagesSiteRequest**](ReposUpdateInformationAboutPagesSiteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdateInvitation

> RepositoryInvitation ReposUpdateInvitation(ctx, owner, repo, invitationId).ReposUpdateInvitationRequest(reposUpdateInvitationRequest).Execute()

Update a repository invitation



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
    invitationId := int32(56) // int32 | The unique identifier of the invitation.
    reposUpdateInvitationRequest := *openapiclient.NewReposUpdateInvitationRequest() // ReposUpdateInvitationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdateInvitation(context.Background(), owner, repo, invitationId).ReposUpdateInvitationRequest(reposUpdateInvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdateInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdateInvitation`: RepositoryInvitation
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdateInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**invitationId** | **int32** | The unique identifier of the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposUpdateInvitationRequest** | [**ReposUpdateInvitationRequest**](ReposUpdateInvitationRequest.md) |  | 

### Return type

[**RepositoryInvitation**](RepositoryInvitation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdatePullRequestReviewProtection

> ProtectedBranchPullRequestReview ReposUpdatePullRequestReviewProtection(ctx, owner, repo, branch).ReposUpdatePullRequestReviewProtectionRequest(reposUpdatePullRequestReviewProtectionRequest).Execute()

Update pull request review protection



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
    branch := "branch_example" // string | The name of the branch.
    reposUpdatePullRequestReviewProtectionRequest := *openapiclient.NewReposUpdatePullRequestReviewProtectionRequest() // ReposUpdatePullRequestReviewProtectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdatePullRequestReviewProtection(context.Background(), owner, repo, branch).ReposUpdatePullRequestReviewProtectionRequest(reposUpdatePullRequestReviewProtectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdatePullRequestReviewProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdatePullRequestReviewProtection`: ProtectedBranchPullRequestReview
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdatePullRequestReviewProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdatePullRequestReviewProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposUpdatePullRequestReviewProtectionRequest** | [**ReposUpdatePullRequestReviewProtectionRequest**](ReposUpdatePullRequestReviewProtectionRequest.md) |  | 

### Return type

[**ProtectedBranchPullRequestReview**](ProtectedBranchPullRequestReview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdateRelease

> Release ReposUpdateRelease(ctx, owner, repo, releaseId).ReposUpdateReleaseRequest(reposUpdateReleaseRequest).Execute()

Update a release



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
    releaseId := int32(56) // int32 | The unique identifier of the release.
    reposUpdateReleaseRequest := *openapiclient.NewReposUpdateReleaseRequest() // ReposUpdateReleaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdateRelease(context.Background(), owner, repo, releaseId).ReposUpdateReleaseRequest(reposUpdateReleaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdateRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdateRelease`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdateRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**releaseId** | **int32** | The unique identifier of the release. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposUpdateReleaseRequest** | [**ReposUpdateReleaseRequest**](ReposUpdateReleaseRequest.md) |  | 

### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdateReleaseAsset

> ReleaseAsset ReposUpdateReleaseAsset(ctx, owner, repo, assetId).ReposUpdateReleaseAssetRequest(reposUpdateReleaseAssetRequest).Execute()

Update a release asset



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
    assetId := int32(56) // int32 | The unique identifier of the asset.
    reposUpdateReleaseAssetRequest := *openapiclient.NewReposUpdateReleaseAssetRequest() // ReposUpdateReleaseAssetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdateReleaseAsset(context.Background(), owner, repo, assetId).ReposUpdateReleaseAssetRequest(reposUpdateReleaseAssetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdateReleaseAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdateReleaseAsset`: ReleaseAsset
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdateReleaseAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**assetId** | **int32** | The unique identifier of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdateReleaseAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposUpdateReleaseAssetRequest** | [**ReposUpdateReleaseAssetRequest**](ReposUpdateReleaseAssetRequest.md) |  | 

### Return type

[**ReleaseAsset**](ReleaseAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdateStatusCheckProtection

> StatusCheckPolicy ReposUpdateStatusCheckProtection(ctx, owner, repo, branch).ReposUpdateStatusCheckProtectionRequest(reposUpdateStatusCheckProtectionRequest).Execute()

Update status check protection



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
    branch := "branch_example" // string | The name of the branch.
    reposUpdateStatusCheckProtectionRequest := *openapiclient.NewReposUpdateStatusCheckProtectionRequest() // ReposUpdateStatusCheckProtectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdateStatusCheckProtection(context.Background(), owner, repo, branch).ReposUpdateStatusCheckProtectionRequest(reposUpdateStatusCheckProtectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdateStatusCheckProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdateStatusCheckProtection`: StatusCheckPolicy
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdateStatusCheckProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**branch** | **string** | The name of the branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdateStatusCheckProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposUpdateStatusCheckProtectionRequest** | [**ReposUpdateStatusCheckProtectionRequest**](ReposUpdateStatusCheckProtectionRequest.md) |  | 

### Return type

[**StatusCheckPolicy**](StatusCheckPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdateWebhook

> Hook ReposUpdateWebhook(ctx, owner, repo, hookId).ReposUpdateWebhookRequest(reposUpdateWebhookRequest).Execute()

Update a repository webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.
    reposUpdateWebhookRequest := *openapiclient.NewReposUpdateWebhookRequest() // ReposUpdateWebhookRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdateWebhook(context.Background(), owner, repo, hookId).ReposUpdateWebhookRequest(reposUpdateWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdateWebhook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposUpdateWebhookRequest** | [**ReposUpdateWebhookRequest**](ReposUpdateWebhookRequest.md) |  | 

### Return type

[**Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUpdateWebhookConfigForRepo

> WebhookConfig ReposUpdateWebhookConfigForRepo(ctx, owner, repo, hookId).ReposUpdateWebhookConfigForRepoRequest(reposUpdateWebhookConfigForRepoRequest).Execute()

Update a webhook configuration for a repository



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
    hookId := int32(56) // int32 | The unique identifier of the hook.
    reposUpdateWebhookConfigForRepoRequest := *openapiclient.NewReposUpdateWebhookConfigForRepoRequest() // ReposUpdateWebhookConfigForRepoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUpdateWebhookConfigForRepo(context.Background(), owner, repo, hookId).ReposUpdateWebhookConfigForRepoRequest(reposUpdateWebhookConfigForRepoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUpdateWebhookConfigForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUpdateWebhookConfigForRepo`: WebhookConfig
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUpdateWebhookConfigForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUpdateWebhookConfigForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reposUpdateWebhookConfigForRepoRequest** | [**ReposUpdateWebhookConfigForRepoRequest**](ReposUpdateWebhookConfigForRepoRequest.md) |  | 

### Return type

[**WebhookConfig**](WebhookConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUploadReleaseAsset

> ReleaseAsset ReposUploadReleaseAsset(ctx, owner, repo, releaseId).Name(name).Label(label).Body(body).Execute()

Upload a release asset



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
    releaseId := int32(56) // int32 | The unique identifier of the release.
    name := "name_example" // string | 
    label := "label_example" // string |  (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposUploadReleaseAsset(context.Background(), owner, repo, releaseId).Name(name).Label(label).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUploadReleaseAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUploadReleaseAsset`: ReleaseAsset
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUploadReleaseAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**releaseId** | **int32** | The unique identifier of the release. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposUploadReleaseAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **name** | **string** |  | 
 **label** | **string** |  | 
 **body** | **string** |  | 

### Return type

[**ReleaseAsset**](ReleaseAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

