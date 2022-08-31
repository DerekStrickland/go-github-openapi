# AuditLogEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int32** | The time the audit log event occurred, given as a [Unix timestamp](http://en.wikipedia.org/wiki/Unix_time). | [optional] 
**Action** | Pointer to **string** | The name of the action that was performed, for example &#x60;user.login&#x60; or &#x60;repo.create&#x60;. | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ActiveWas** | Pointer to **bool** |  | [optional] 
**Actor** | Pointer to **string** | The actor who performed the action. | [optional] 
**ActorId** | Pointer to **int32** | The id of the actor who performed the action. | [optional] 
**ActorLocation** | Pointer to [**AuditLogEventActorLocation**](AuditLogEventActorLocation.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**BlockedUser** | Pointer to **string** | The username of the account being blocked. | [optional] 
**Business** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ConfigWas** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** | The time the audit log event was recorded, given as a [Unix timestamp](http://en.wikipedia.org/wiki/Unix_time). | [optional] 
**DeployKeyFingerprint** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** | A unique identifier for an audit event. | [optional] 
**Emoji** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EventsWere** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Explanation** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**HookId** | Pointer to **int32** |  | [optional] 
**LimitedAvailability** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OldUser** | Pointer to **string** |  | [optional] 
**OpensshPublicKey** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**PreviousVisibility** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Repo** | Pointer to **string** | The name of the repository. | [optional] 
**Repository** | Pointer to **string** | The name of the repository. | [optional] 
**RepositoryPublic** | Pointer to **bool** |  | [optional] 
**TargetLogin** | Pointer to **string** |  | [optional] 
**Team** | Pointer to **string** |  | [optional] 
**TransportProtocol** | Pointer to **int32** | The type of protocol (for example, HTTP or SSH) used to transfer Git data. | [optional] 
**TransportProtocolName** | Pointer to **string** | A human readable name for the protocol (for example, HTTP or SSH) used to transfer Git data. | [optional] 
**User** | Pointer to **string** | The user that was affected by the action performed (if available). | [optional] 
**Visibility** | Pointer to **string** | The repository visibility, for example &#x60;public&#x60; or &#x60;private&#x60;. | [optional] 

## Methods

### NewAuditLogEvent

`func NewAuditLogEvent() *AuditLogEvent`

NewAuditLogEvent instantiates a new AuditLogEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEventWithDefaults

`func NewAuditLogEventWithDefaults() *AuditLogEvent`

NewAuditLogEventWithDefaults instantiates a new AuditLogEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *AuditLogEvent) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditLogEvent) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditLogEvent) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AuditLogEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAction

`func (o *AuditLogEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditLogEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditLogEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuditLogEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActive

`func (o *AuditLogEvent) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AuditLogEvent) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AuditLogEvent) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AuditLogEvent) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActiveWas

`func (o *AuditLogEvent) GetActiveWas() bool`

GetActiveWas returns the ActiveWas field if non-nil, zero value otherwise.

### GetActiveWasOk

`func (o *AuditLogEvent) GetActiveWasOk() (*bool, bool)`

GetActiveWasOk returns a tuple with the ActiveWas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWas

`func (o *AuditLogEvent) SetActiveWas(v bool)`

SetActiveWas sets ActiveWas field to given value.

### HasActiveWas

`func (o *AuditLogEvent) HasActiveWas() bool`

HasActiveWas returns a boolean if a field has been set.

### GetActor

`func (o *AuditLogEvent) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AuditLogEvent) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AuditLogEvent) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AuditLogEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetActorId

`func (o *AuditLogEvent) GetActorId() int32`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *AuditLogEvent) GetActorIdOk() (*int32, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *AuditLogEvent) SetActorId(v int32)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *AuditLogEvent) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### GetActorLocation

`func (o *AuditLogEvent) GetActorLocation() AuditLogEventActorLocation`

GetActorLocation returns the ActorLocation field if non-nil, zero value otherwise.

### GetActorLocationOk

`func (o *AuditLogEvent) GetActorLocationOk() (*AuditLogEventActorLocation, bool)`

GetActorLocationOk returns a tuple with the ActorLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorLocation

`func (o *AuditLogEvent) SetActorLocation(v AuditLogEventActorLocation)`

SetActorLocation sets ActorLocation field to given value.

### HasActorLocation

`func (o *AuditLogEvent) HasActorLocation() bool`

HasActorLocation returns a boolean if a field has been set.

### GetData

`func (o *AuditLogEvent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuditLogEvent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuditLogEvent) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AuditLogEvent) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOrgId

`func (o *AuditLogEvent) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuditLogEvent) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuditLogEvent) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AuditLogEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetBlockedUser

`func (o *AuditLogEvent) GetBlockedUser() string`

GetBlockedUser returns the BlockedUser field if non-nil, zero value otherwise.

### GetBlockedUserOk

`func (o *AuditLogEvent) GetBlockedUserOk() (*string, bool)`

GetBlockedUserOk returns a tuple with the BlockedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUser

`func (o *AuditLogEvent) SetBlockedUser(v string)`

SetBlockedUser sets BlockedUser field to given value.

### HasBlockedUser

`func (o *AuditLogEvent) HasBlockedUser() bool`

HasBlockedUser returns a boolean if a field has been set.

### GetBusiness

`func (o *AuditLogEvent) GetBusiness() string`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *AuditLogEvent) GetBusinessOk() (*string, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *AuditLogEvent) SetBusiness(v string)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *AuditLogEvent) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetConfig

`func (o *AuditLogEvent) GetConfig() []map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AuditLogEvent) GetConfigOk() (*[]map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AuditLogEvent) SetConfig(v []map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AuditLogEvent) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConfigWas

`func (o *AuditLogEvent) GetConfigWas() []map[string]interface{}`

GetConfigWas returns the ConfigWas field if non-nil, zero value otherwise.

### GetConfigWasOk

`func (o *AuditLogEvent) GetConfigWasOk() (*[]map[string]interface{}, bool)`

GetConfigWasOk returns a tuple with the ConfigWas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigWas

`func (o *AuditLogEvent) SetConfigWas(v []map[string]interface{})`

SetConfigWas sets ConfigWas field to given value.

### HasConfigWas

`func (o *AuditLogEvent) HasConfigWas() bool`

HasConfigWas returns a boolean if a field has been set.

### GetContentType

`func (o *AuditLogEvent) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AuditLogEvent) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AuditLogEvent) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *AuditLogEvent) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditLogEvent) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditLogEvent) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditLogEvent) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditLogEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeployKeyFingerprint

`func (o *AuditLogEvent) GetDeployKeyFingerprint() string`

GetDeployKeyFingerprint returns the DeployKeyFingerprint field if non-nil, zero value otherwise.

### GetDeployKeyFingerprintOk

`func (o *AuditLogEvent) GetDeployKeyFingerprintOk() (*string, bool)`

GetDeployKeyFingerprintOk returns a tuple with the DeployKeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployKeyFingerprint

`func (o *AuditLogEvent) SetDeployKeyFingerprint(v string)`

SetDeployKeyFingerprint sets DeployKeyFingerprint field to given value.

### HasDeployKeyFingerprint

`func (o *AuditLogEvent) HasDeployKeyFingerprint() bool`

HasDeployKeyFingerprint returns a boolean if a field has been set.

### GetDocumentId

`func (o *AuditLogEvent) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *AuditLogEvent) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *AuditLogEvent) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *AuditLogEvent) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetEmoji

`func (o *AuditLogEvent) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *AuditLogEvent) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *AuditLogEvent) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *AuditLogEvent) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetEvents

`func (o *AuditLogEvent) GetEvents() []map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AuditLogEvent) GetEventsOk() (*[]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AuditLogEvent) SetEvents(v []map[string]interface{})`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AuditLogEvent) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetEventsWere

`func (o *AuditLogEvent) GetEventsWere() []map[string]interface{}`

GetEventsWere returns the EventsWere field if non-nil, zero value otherwise.

### GetEventsWereOk

`func (o *AuditLogEvent) GetEventsWereOk() (*[]map[string]interface{}, bool)`

GetEventsWereOk returns a tuple with the EventsWere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsWere

`func (o *AuditLogEvent) SetEventsWere(v []map[string]interface{})`

SetEventsWere sets EventsWere field to given value.

### HasEventsWere

`func (o *AuditLogEvent) HasEventsWere() bool`

HasEventsWere returns a boolean if a field has been set.

### GetExplanation

`func (o *AuditLogEvent) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *AuditLogEvent) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *AuditLogEvent) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *AuditLogEvent) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetFingerprint

`func (o *AuditLogEvent) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *AuditLogEvent) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *AuditLogEvent) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *AuditLogEvent) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHookId

`func (o *AuditLogEvent) GetHookId() int32`

GetHookId returns the HookId field if non-nil, zero value otherwise.

### GetHookIdOk

`func (o *AuditLogEvent) GetHookIdOk() (*int32, bool)`

GetHookIdOk returns a tuple with the HookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookId

`func (o *AuditLogEvent) SetHookId(v int32)`

SetHookId sets HookId field to given value.

### HasHookId

`func (o *AuditLogEvent) HasHookId() bool`

HasHookId returns a boolean if a field has been set.

### GetLimitedAvailability

`func (o *AuditLogEvent) GetLimitedAvailability() bool`

GetLimitedAvailability returns the LimitedAvailability field if non-nil, zero value otherwise.

### GetLimitedAvailabilityOk

`func (o *AuditLogEvent) GetLimitedAvailabilityOk() (*bool, bool)`

GetLimitedAvailabilityOk returns a tuple with the LimitedAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedAvailability

`func (o *AuditLogEvent) SetLimitedAvailability(v bool)`

SetLimitedAvailability sets LimitedAvailability field to given value.

### HasLimitedAvailability

`func (o *AuditLogEvent) HasLimitedAvailability() bool`

HasLimitedAvailability returns a boolean if a field has been set.

### GetMessage

`func (o *AuditLogEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuditLogEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuditLogEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuditLogEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *AuditLogEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditLogEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditLogEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditLogEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOldUser

`func (o *AuditLogEvent) GetOldUser() string`

GetOldUser returns the OldUser field if non-nil, zero value otherwise.

### GetOldUserOk

`func (o *AuditLogEvent) GetOldUserOk() (*string, bool)`

GetOldUserOk returns a tuple with the OldUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldUser

`func (o *AuditLogEvent) SetOldUser(v string)`

SetOldUser sets OldUser field to given value.

### HasOldUser

`func (o *AuditLogEvent) HasOldUser() bool`

HasOldUser returns a boolean if a field has been set.

### GetOpensshPublicKey

`func (o *AuditLogEvent) GetOpensshPublicKey() string`

GetOpensshPublicKey returns the OpensshPublicKey field if non-nil, zero value otherwise.

### GetOpensshPublicKeyOk

`func (o *AuditLogEvent) GetOpensshPublicKeyOk() (*string, bool)`

GetOpensshPublicKeyOk returns a tuple with the OpensshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpensshPublicKey

`func (o *AuditLogEvent) SetOpensshPublicKey(v string)`

SetOpensshPublicKey sets OpensshPublicKey field to given value.

### HasOpensshPublicKey

`func (o *AuditLogEvent) HasOpensshPublicKey() bool`

HasOpensshPublicKey returns a boolean if a field has been set.

### GetOrg

`func (o *AuditLogEvent) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AuditLogEvent) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AuditLogEvent) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AuditLogEvent) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPreviousVisibility

`func (o *AuditLogEvent) GetPreviousVisibility() string`

GetPreviousVisibility returns the PreviousVisibility field if non-nil, zero value otherwise.

### GetPreviousVisibilityOk

`func (o *AuditLogEvent) GetPreviousVisibilityOk() (*string, bool)`

GetPreviousVisibilityOk returns a tuple with the PreviousVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVisibility

`func (o *AuditLogEvent) SetPreviousVisibility(v string)`

SetPreviousVisibility sets PreviousVisibility field to given value.

### HasPreviousVisibility

`func (o *AuditLogEvent) HasPreviousVisibility() bool`

HasPreviousVisibility returns a boolean if a field has been set.

### GetReadOnly

`func (o *AuditLogEvent) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AuditLogEvent) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AuditLogEvent) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *AuditLogEvent) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRepo

`func (o *AuditLogEvent) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AuditLogEvent) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AuditLogEvent) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AuditLogEvent) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRepository

`func (o *AuditLogEvent) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AuditLogEvent) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AuditLogEvent) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AuditLogEvent) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRepositoryPublic

`func (o *AuditLogEvent) GetRepositoryPublic() bool`

GetRepositoryPublic returns the RepositoryPublic field if non-nil, zero value otherwise.

### GetRepositoryPublicOk

`func (o *AuditLogEvent) GetRepositoryPublicOk() (*bool, bool)`

GetRepositoryPublicOk returns a tuple with the RepositoryPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryPublic

`func (o *AuditLogEvent) SetRepositoryPublic(v bool)`

SetRepositoryPublic sets RepositoryPublic field to given value.

### HasRepositoryPublic

`func (o *AuditLogEvent) HasRepositoryPublic() bool`

HasRepositoryPublic returns a boolean if a field has been set.

### GetTargetLogin

`func (o *AuditLogEvent) GetTargetLogin() string`

GetTargetLogin returns the TargetLogin field if non-nil, zero value otherwise.

### GetTargetLoginOk

`func (o *AuditLogEvent) GetTargetLoginOk() (*string, bool)`

GetTargetLoginOk returns a tuple with the TargetLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLogin

`func (o *AuditLogEvent) SetTargetLogin(v string)`

SetTargetLogin sets TargetLogin field to given value.

### HasTargetLogin

`func (o *AuditLogEvent) HasTargetLogin() bool`

HasTargetLogin returns a boolean if a field has been set.

### GetTeam

`func (o *AuditLogEvent) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *AuditLogEvent) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *AuditLogEvent) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *AuditLogEvent) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetTransportProtocol

`func (o *AuditLogEvent) GetTransportProtocol() int32`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *AuditLogEvent) GetTransportProtocolOk() (*int32, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *AuditLogEvent) SetTransportProtocol(v int32)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *AuditLogEvent) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetTransportProtocolName

`func (o *AuditLogEvent) GetTransportProtocolName() string`

GetTransportProtocolName returns the TransportProtocolName field if non-nil, zero value otherwise.

### GetTransportProtocolNameOk

`func (o *AuditLogEvent) GetTransportProtocolNameOk() (*string, bool)`

GetTransportProtocolNameOk returns a tuple with the TransportProtocolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocolName

`func (o *AuditLogEvent) SetTransportProtocolName(v string)`

SetTransportProtocolName sets TransportProtocolName field to given value.

### HasTransportProtocolName

`func (o *AuditLogEvent) HasTransportProtocolName() bool`

HasTransportProtocolName returns a boolean if a field has been set.

### GetUser

`func (o *AuditLogEvent) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuditLogEvent) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuditLogEvent) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AuditLogEvent) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVisibility

`func (o *AuditLogEvent) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AuditLogEvent) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AuditLogEvent) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AuditLogEvent) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


