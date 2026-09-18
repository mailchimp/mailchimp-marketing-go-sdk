# Reference
## root
<details><summary><code>client.Root.List() -> *mailchimpmarketinggosdk.ListRootResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get links to all other resources available in the API.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListRootRequest{}
client.Root.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AccountExports
<details><summary><code>client.AccountExports.List() -> *mailchimpmarketinggosdk.ListAccountExportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of account exports for a given account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListAccountExportsRequest{}
client.AccountExports.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AccountExports.Create(request) -> *mailchimpmarketinggosdk.CreateAccountExportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new account export in your Mailchimp account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateAccountExportsRequest{
    IncludeStages: []mailchimpmarketinggosdk.CreateAccountExportsRequestIncludeStagesItem{
        mailchimpmarketinggosdk.CreateAccountExportsRequestIncludeStagesItemAudiences,
        mailchimpmarketinggosdk.CreateAccountExportsRequestIncludeStagesItemGalleryFiles,
    },
}
client.AccountExports.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**includeStages:** `[]*mailchimpmarketinggosdk.CreateAccountExportsRequestIncludeStagesItem` — The stages of an account export to include.
    
</dd>
</dl>

<dl>
<dd>

**sinceTimestamp:** `*time.Time` — An ISO 8601 date that will limit the export to only records created after a given time. For instance, the reports stage will contain any campaign sent after the given timestamp. Audiences, however, are excluded from this limit.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AccountExports.Get(ExportID) -> *mailchimpmarketinggosdk.GetAccountExportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific account export.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetAccountExportsRequest{
    ExportID: "export_id",
}
client.AccountExports.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**exportID:** `string` — The unique id for the account export.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ActivityFeed
<details><summary><code>client.ActivityFeed.List() -> []*mailchimpmarketinggosdk.ListActivityFeedResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about the activity feed endpoint's resources.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ActivityFeed.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ActivityFeed.ListChimpChatter() -> *mailchimpmarketinggosdk.ListChimpChatterActivityFeedResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return the Chimp Chatter for this account ordered by most recent.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListChimpChatterActivityFeedRequest{}
client.ActivityFeed.ListChimpChatter(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Audiences
<details><summary><code>client.Audiences.GetAudienceContactList(AudienceID) -> *mailchimpmarketinggosdk.GetAudienceContactListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of omni-channel contacts for a given audience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetAudienceContactListRequest{
    AudienceID: "audience_id",
}
client.Audiences.GetAudienceContactList(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**audienceID:** `string` — The unique ID for the audience.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Paginate through a collection of records by setting the `cursor` parameter to a `next_cursor` attribute returned by a previous request. Default value fetches the first "page" of results.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Restricts the response to contacts created at or before the specified time (inclusive). Uses ISO 8601 format: 2025-04-23T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**createdSince:** `*time.Time` — Restricts the response to contacts created after the specified time (exclusive). Uses ISO 8601 format: 2025-04-23T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**updatedBefore:** `*time.Time` — Restricts the response to contacts updated at or before the specified time (inclusive). Uses ISO 8601 format: 2025-04-23T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**updatedSince:** `*time.Time` — Restricts the response to contacts updated after the specified time (exclusive). Uses ISO 8601 format: 2025-04-23T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.GetAudienceContactListRequestSortField` — Specifies the field to sort the returned contacts by.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.GetAudienceContactListRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.CreateAudienceContact(AudienceID, request) -> *mailchimpmarketinggosdk.AudiencesContact</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new omni-channel contact for an audience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateAudienceContactRequest{
    AudienceID: "audience_id",
}
client.Audiences.CreateAudienceContact(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**audienceID:** `string` — The unique ID for the audience.
    
</dd>
</dl>

<dl>
<dd>

**mergeFieldValidationMode:** `*mailchimpmarketinggosdk.CreateAudienceContactRequestMergeFieldValidationMode` — Defines how merge field validation is handled. When set to `ignore_required_checks`, the API does not raise an error if required merge fields are missing from the request. When set to `strict`, the API enforces validation and returns an error if any required merge field is not provided. If this setting is omitted, `strict` is applied by default.
    
</dd>
</dl>

<dl>
<dd>

**dataMode:** `*mailchimpmarketinggosdk.CreateAudienceContactRequestDataMode` — Indicates the data processing mode. In `historical` mode, contact data changes do not trigger automations or webhooks. In `live mode`, such changes do trigger them.
    
</dd>
</dl>

<dl>
<dd>

**emailChannel:** `*mailchimpmarketinggosdk.CreateAudienceContactRequestEmailChannel` 
    
</dd>
</dl>

<dl>
<dd>

**language:** `*string` — The contact's detected language.
    
</dd>
</dl>

<dl>
<dd>

**mergeFields:** `map[string]*mailchimpmarketinggosdk.CreateAudienceContactRequestMergeFieldsValue` — A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
    
</dd>
</dl>

<dl>
<dd>

**smsChannel:** `*mailchimpmarketinggosdk.CreateAudienceContactRequestSmsChannel` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]*mailchimpmarketinggosdk.CreateAudienceContactRequestTagsItem` — An array of tags to add to the contact. Accepts tag name strings or objects with name and status. This operation is append-only; existing tags will be preserved, and only new tags from this array will be added.
    
</dd>
</dl>

<dl>
<dd>

**updateExisting:** `*bool` — If a contact already exists, update them instead of returning a conflict error. When `true` and a matching contact is found (by email or phone), the existing contact is updated with the provided channel data. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.GetAudienceContact(AudienceID, ContactID) -> *mailchimpmarketinggosdk.AudiencesContact</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific omni-channel contact in an audience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetAudienceContactRequest{
    AudienceID: "audience_id",
    ContactID: "contact_id",
}
client.Audiences.GetAudienceContact(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**audienceID:** `string` — The unique ID for the audience.
    
</dd>
</dl>

<dl>
<dd>

**contactID:** `string` — A unique identifier for the contact, which can be a Mailchimp contact ID or a channel hash. A channel hash must follow the format email:[md5_hash] (where the hash is the MD5 of the lowercased email address) or sms:[sha256_hash] (where the hash is the SHA256 of the E.164-formatted phone number).
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.PatchAudienceContact(AudienceID, ContactID, request) -> *mailchimpmarketinggosdk.AudiencesContact</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an existing omni-channel contact.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.PatchAudienceContactRequest{
    AudienceID: "audience_id",
    ContactID: "contact_id",
}
client.Audiences.PatchAudienceContact(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**audienceID:** `string` — The unique ID for the audience.
    
</dd>
</dl>

<dl>
<dd>

**contactID:** `string` — The unique id for the contact.
    
</dd>
</dl>

<dl>
<dd>

**mergeFieldValidationMode:** `*mailchimpmarketinggosdk.PatchAudienceContactRequestMergeFieldValidationMode` — Defines how merge field validation is handled. When set to `ignore_required_checks`, the API does not raise an error if required merge fields are missing from the request. When set to `strict`, the API enforces validation and returns an error if any required merge field is not provided. If this setting is omitted, `strict` is applied by default.
    
</dd>
</dl>

<dl>
<dd>

**dataMode:** `*mailchimpmarketinggosdk.PatchAudienceContactRequestDataMode` — Indicates the data processing mode. In `historical` mode, contact data changes do not trigger automations or webhooks. In `live mode`, such changes do trigger them.
    
</dd>
</dl>

<dl>
<dd>

**emailChannel:** `*mailchimpmarketinggosdk.PatchAudienceContactRequestEmailChannel` 
    
</dd>
</dl>

<dl>
<dd>

**language:** `*string` — The contact's detected language.
    
</dd>
</dl>

<dl>
<dd>

**mergeFields:** `map[string]*mailchimpmarketinggosdk.PatchAudienceContactRequestMergeFieldsValue` — A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
    
</dd>
</dl>

<dl>
<dd>

**smsChannel:** `*mailchimpmarketinggosdk.PatchAudienceContactRequestSmsChannel` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]*mailchimpmarketinggosdk.PatchAudienceContactRequestTagsItem` — An array of tags to add to the contact. Accepts tag name strings or objects with name and status. This operation is append-only; existing tags will be preserved, and only new tags from this array will be added.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.PostAudiencesContactsActionsArchive(AudienceID, ContactID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Archives a Contact.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.PostAudiencesContactsActionsArchiveRequest{
    AudienceID: "audience_id",
    ContactID: "contact_id",
}
client.Audiences.PostAudiencesContactsActionsArchive(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**audienceID:** `string` — The unique ID for the audience.
    
</dd>
</dl>

<dl>
<dd>

**contactID:** `string` — The unique id for the contact.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.PostAudiencesContactsActionsForget(AudienceID, ContactID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Forgets a Contact.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.PostAudiencesContactsActionsForgetRequest{
    AudienceID: "audience_id",
    ContactID: "contact_id",
}
client.Audiences.PostAudiencesContactsActionsForget(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**audienceID:** `string` — The unique ID for the audience.
    
</dd>
</dl>

<dl>
<dd>

**contactID:** `string` — The unique id for the contact.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AuthorizedApps
<details><summary><code>client.AuthorizedApps.List() -> *mailchimpmarketinggosdk.ListAuthorizedAppsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of an account's registered, connected applications.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListAuthorizedAppsRequest{}
client.AuthorizedApps.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AuthorizedApps.Get(AppID) -> *mailchimpmarketinggosdk.GetAuthorizedAppsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific authorized application.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetAuthorizedAppsRequest{
    AppID: "app_id",
}
client.AuthorizedApps.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**appID:** `string` — The unique id for the connected authorized application.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## automations
<details><summary><code>client.Automations.List() -> *mailchimpmarketinggosdk.ListAutomationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a summary of an account's classic automations.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListAutomationsRequest{}
client.Automations.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**beforeCreateTime:** `*time.Time` — Restrict the response to automations created before this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceCreateTime:** `*time.Time` — Restrict the response to automations created after this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**beforeStartTime:** `*time.Time` — Restrict the response to automations started before this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceStartTime:** `*time.Time` — Restrict the response to automations started after this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*mailchimpmarketinggosdk.ListAutomationsRequestStatus` — Restrict the results to automations with the specified status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.Create(request) -> *mailchimpmarketinggosdk.AutomationWorkflow</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new classic automation in your Mailchimp account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateAutomationsRequest{
    Recipients: &mailchimpmarketinggosdk.CreateAutomationsRequestRecipients{},
    TriggerSettings: &mailchimpmarketinggosdk.CreateAutomationsRequestTriggerSettings{
        WorkflowType: mailchimpmarketinggosdk.CreateAutomationsRequestTriggerSettingsWorkflowTypeAbandonedBrowse,
    },
}
client.Automations.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**recipients:** `*mailchimpmarketinggosdk.CreateAutomationsRequestRecipients` — List settings for the Automation.
    
</dd>
</dl>

<dl>
<dd>

**settings:** `*mailchimpmarketinggosdk.CreateAutomationsRequestSettings` — The settings for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**triggerSettings:** `*mailchimpmarketinggosdk.CreateAutomationsRequestTriggerSettings` — Trigger settings for the Automation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.Get(WorkflowID) -> *mailchimpmarketinggosdk.AutomationWorkflow</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a summary of an individual classic automation workflow's settings and content. The `trigger_settings` object returns information for the first email in the workflow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetAutomationsRequest{
    WorkflowID: "workflow_id",
}
client.Automations.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.CreateActionArchive(WorkflowID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Archiving will permanently end your automation and keep the report data. You’ll be able to replicate your archived automation, but you can’t restart it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionArchiveAutomationsRequest{
    WorkflowID: "workflow_id",
}
client.Automations.CreateActionArchive(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.CreateActionPauseAllEmail(WorkflowID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pause all emails in a specific classic automation workflow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionPauseAllEmailAutomationsRequest{
    WorkflowID: "workflow_id",
}
client.Automations.CreateActionPauseAllEmail(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.CreateActionStartAllEmail(WorkflowID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Start all emails in a classic automation workflow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionStartAllEmailAutomationsRequest{
    WorkflowID: "workflow_id",
}
client.Automations.CreateActionStartAllEmail(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.ListEmails(WorkflowID) -> *mailchimpmarketinggosdk.ListEmailsAutomationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a summary of the emails in a classic automation workflow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListEmailsAutomationsRequest{
    WorkflowID: "workflow_id",
}
client.Automations.ListEmails(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.GetEmail(WorkflowID, WorkflowEmailID) -> *mailchimpmarketinggosdk.AutomationWorkflowEmail</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about an individual classic automation workflow email.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetEmailAutomationsRequest{
    WorkflowID: "workflow_id",
    WorkflowEmailID: "workflow_email_id",
}
client.Automations.GetEmail(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**workflowEmailID:** `string` — The unique id for the Automation workflow email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.DeleteEmail(WorkflowID, WorkflowEmailID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes an individual classic automation workflow email. Emails from certain workflow types, including the Abandoned Cart Email (abandonedCart) and Product Retargeting Email (abandonedBrowse) Workflows, cannot be deleted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteEmailAutomationsRequest{
    WorkflowID: "workflow_id",
    WorkflowEmailID: "workflow_email_id",
}
client.Automations.DeleteEmail(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**workflowEmailID:** `string` — The unique id for the Automation workflow email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.UpdateEmail(WorkflowID, WorkflowEmailID, request) -> *mailchimpmarketinggosdk.AutomationWorkflowEmail</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update settings for a classic automation workflow email.  Only works with workflows of type: abandonedBrowse, abandonedCart, emailFollowup, or singleWelcome.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateEmailAutomationsRequest{
    WorkflowID: "workflow_id",
    WorkflowEmailID: "workflow_email_id",
}
client.Automations.UpdateEmail(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**workflowEmailID:** `string` — The unique id for the Automation workflow email.
    
</dd>
</dl>

<dl>
<dd>

**delay:** `*mailchimpmarketinggosdk.UpdateEmailAutomationsRequestDelay` — The delay settings for an automation email.
    
</dd>
</dl>

<dl>
<dd>

**settings:** `*mailchimpmarketinggosdk.UpdateEmailAutomationsRequestSettings` — Settings for the campaign including the email subject, from name, and from email address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.CreateEmailActionPause(WorkflowID, WorkflowEmailID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pause an automated email.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateEmailActionPauseAutomationsRequest{
    WorkflowID: "workflow_id",
    WorkflowEmailID: "workflow_email_id",
}
client.Automations.CreateEmailActionPause(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**workflowEmailID:** `string` — The unique id for the Automation workflow email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.CreateEmailActionStart(WorkflowID, WorkflowEmailID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Start an automated email.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateEmailActionStartAutomationsRequest{
    WorkflowID: "workflow_id",
    WorkflowEmailID: "workflow_email_id",
}
client.Automations.CreateEmailActionStart(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**workflowEmailID:** `string` — The unique id for the Automation workflow email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.ListEmailQueue(WorkflowID, WorkflowEmailID) -> *mailchimpmarketinggosdk.ListEmailQueueAutomationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a classic automation email queue.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListEmailQueueAutomationsRequest{
    WorkflowID: "workflow_id",
    WorkflowEmailID: "workflow_email_id",
}
client.Automations.ListEmailQueue(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**workflowEmailID:** `string` — The unique id for the Automation workflow email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.CreateEmailQueue(WorkflowID, WorkflowEmailID, request) -> *mailchimpmarketinggosdk.SubscriberInAutomationQueue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Manually add a subscriber to a workflow, bypassing the default trigger settings. You can also use this endpoint to trigger a series of automated emails in an API 3.0 workflow type.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateEmailQueueAutomationsRequest{
    WorkflowID: "workflow_id",
    WorkflowEmailID: "workflow_email_id",
    EmailAddress: "email_address",
}
client.Automations.CreateEmailQueue(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**workflowEmailID:** `string` — The unique id for the Automation workflow email.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `string` — The list member's email address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.GetEmailQueue(WorkflowID, WorkflowEmailID, SubscriberHash) -> *mailchimpmarketinggosdk.SubscriberInAutomationQueue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific subscriber in a classic automation email queue.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetEmailQueueAutomationsRequest{
    WorkflowID: "workflow_id",
    WorkflowEmailID: "workflow_email_id",
    SubscriberHash: "subscriber_hash",
}
client.Automations.GetEmailQueue(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**workflowEmailID:** `string` — The unique id for the Automation workflow email.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.ListRemovedSubscribers(WorkflowID) -> *mailchimpmarketinggosdk.ListRemovedSubscribersAutomationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about subscribers who were removed from a classic automation workflow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListRemovedSubscribersAutomationsRequest{
    WorkflowID: "workflow_id",
}
client.Automations.ListRemovedSubscribers(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.CreateRemovedSubscriber(WorkflowID, request) -> *mailchimpmarketinggosdk.SubscriberRemovedFromAutomationWorkflow</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a subscriber from a specific classic automation workflow. You can remove a subscriber at any point in an automation workflow, regardless of how many emails they've been sent from that workflow. Once they're removed, they can never be added back to the same workflow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateRemovedSubscriberAutomationsRequest{
    WorkflowID: "workflow_id",
    EmailAddress: "email_address",
}
client.Automations.CreateRemovedSubscriber(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `string` — The list member's email address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Automations.GetRemovedSubscriber(WorkflowID, SubscriberHash) -> *mailchimpmarketinggosdk.SubscriberRemovedFromAutomationWorkflow</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific subscriber who was removed from a classic automation workflow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetRemovedSubscriberAutomationsRequest{
    WorkflowID: "workflow_id",
    SubscriberHash: "subscriber_hash",
}
client.Automations.GetRemovedSubscriber(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workflowID:** `string` — The unique id for the Automation workflow.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BatchWebhooks
<details><summary><code>client.BatchWebhooks.List() -> *mailchimpmarketinggosdk.ListBatchWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all webhooks that have been configured for batches.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListBatchWebhooksRequest{}
client.BatchWebhooks.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BatchWebhooks.Create(request) -> *mailchimpmarketinggosdk.CreateBatchWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Configure a webhook that will fire whenever any batch request completes processing.  You may only have a maximum of 20 batch webhooks.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateBatchWebhooksRequest{
    URL: "http://yourdomain.com/webhook",
}
client.BatchWebhooks.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enabled:** `*bool` — Whether the webhook receives requests or not.
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` — A valid URL for the Webhook.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BatchWebhooks.Get(BatchWebhookID) -> *mailchimpmarketinggosdk.BatchWebhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific batch webhook.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetBatchWebhooksRequest{
    BatchWebhookID: "batch_webhook_id",
}
client.BatchWebhooks.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchWebhookID:** `string` — The unique id for the batch webhook.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BatchWebhooks.Delete(BatchWebhookID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a batch webhook. Webhooks will no longer be sent to the given URL.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteBatchWebhooksRequest{
    BatchWebhookID: "batch_webhook_id",
}
client.BatchWebhooks.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchWebhookID:** `string` — The unique id for the batch webhook.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BatchWebhooks.Update(BatchWebhookID, request) -> *mailchimpmarketinggosdk.BatchWebhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a webhook that will fire whenever any batch request completes processing.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateBatchWebhooksRequest{
    BatchWebhookID: "batch_webhook_id",
}
client.BatchWebhooks.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchWebhookID:** `string` — The unique id for the batch webhook.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the webhook receives requests or not.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — A valid URL for the Webhook.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## batches
<details><summary><code>client.Batches.List() -> *mailchimpmarketinggosdk.ListBatchesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a summary of batch requests that have been made.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListBatchesRequest{}
client.Batches.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batches.Create(request) -> *mailchimpmarketinggosdk.Batch</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Begin processing a batch operations request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateBatchesRequest{
    Operations: []*mailchimpmarketinggosdk.CreateBatchesRequestOperationsItem{
        &mailchimpmarketinggosdk.CreateBatchesRequestOperationsItem{
            Method: mailchimpmarketinggosdk.CreateBatchesRequestOperationsItemMethodGet,
            Path: "/lists",
        },
    },
}
client.Batches.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**operations:** `[]*mailchimpmarketinggosdk.CreateBatchesRequestOperationsItem` — An array of objects that describes operations to perform.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batches.Get(BatchID) -> *mailchimpmarketinggosdk.Batch</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the status of a batch request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetBatchesRequest{
    BatchID: "batch_id",
}
client.Batches.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The unique id for the batch operation.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batches.Delete(BatchID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Stops a batch request from running. Since only one batch request is run at a time, this can be used to cancel a long running request. The results of any completed operations will not be available after this call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteBatchesRequest{
    BatchID: "batch_id",
}
client.Batches.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The unique id for the batch operation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CampaignFolders
<details><summary><code>client.CampaignFolders.List() -> *mailchimpmarketinggosdk.CampaignFolders</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all folders used to organize campaigns.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListCampaignFoldersRequest{}
client.CampaignFolders.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CampaignFolders.Create(request) -> *mailchimpmarketinggosdk.CampaignFolders</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new campaign folder.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateCampaignFoldersRequest{
    Name: "name",
}
client.CampaignFolders.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Name to associate with the folder.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CampaignFolders.Get(FolderID) -> *mailchimpmarketinggosdk.GetCampaignFoldersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific folder used to organize campaigns.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetCampaignFoldersRequest{
    FolderID: "folder_id",
}
client.CampaignFolders.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `string` — The unique id for the campaign folder.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CampaignFolders.Delete(FolderID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific campaign folder, and mark all the campaigns in the folder as 'unfiled'.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteCampaignFoldersRequest{
    FolderID: "folder_id",
}
client.CampaignFolders.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `string` — The unique id for the campaign folder.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CampaignFolders.Update(FolderID, request) -> *mailchimpmarketinggosdk.UpdateCampaignFoldersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific folder used to organize campaigns.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateCampaignFoldersRequest{
    FolderID: "folder_id",
    Name: "name",
}
client.CampaignFolders.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `string` — The unique id for the campaign folder.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Name to associate with the folder.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## campaigns
<details><summary><code>client.Campaigns.List() -> *mailchimpmarketinggosdk.ListCampaignsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all campaigns in an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListCampaignsRequest{}
client.Campaigns.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*mailchimpmarketinggosdk.ListCampaignsRequestType` — The campaign type.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*mailchimpmarketinggosdk.ListCampaignsRequestStatus` — The status of the campaign.
    
</dd>
</dl>

<dl>
<dd>

**beforeSendTime:** `*time.Time` — Restrict the response to campaigns sent before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceSendTime:** `*time.Time` — Restrict the response to campaigns sent after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**beforeCreateTime:** `*time.Time` — Restrict the response to campaigns created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceCreateTime:** `*time.Time` — Restrict the response to campaigns created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**listID:** `*string` — The unique id for the list.
    
</dd>
</dl>

<dl>
<dd>

**folderID:** `*string` — The unique folder id.
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `*string` — Retrieve campaigns sent to a particular list member. Member ID is The MD5 hash of the lowercase version of the list member’s email address.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListCampaignsRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListCampaignsRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>

<dl>
<dd>

**includeResendShortcutEligibility:** `*bool` — Return the `resend_shortcut_eligibility` field in the response, which tells you if the campaign is eligible for the various Campaign Resend Shortcuts offered.
    
</dd>
</dl>

<dl>
<dd>

**includeResendShortcutUsage:** `*bool` — Return the `resend_shortcut_usage` field in the response.  This includes information about campaigns related by a shortcut.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.Create(request) -> *mailchimpmarketinggosdk.Campaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new Mailchimp campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateCampaignsRequest{
    Type: mailchimpmarketinggosdk.CreateCampaignsRequestTypeRegular,
}
client.Campaigns.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**contentType:** `*mailchimpmarketinggosdk.CreateCampaignsRequestContentType` — How the campaign's content is put together. The old drag and drop editor uses 'template' while the new editor uses 'multichannel'. Defaults to template.
    
</dd>
</dl>

<dl>
<dd>

**recipients:** `*mailchimpmarketinggosdk.CreateCampaignsRequestRecipients` — List settings for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**rssOpts:** `*mailchimpmarketinggosdk.CreateCampaignsRequestRssOpts` — [RSS](https://mailchimp.com/help/share-your-blog-posts-with-mailchimp/) options, specific to an RSS campaign.
    
</dd>
</dl>

<dl>
<dd>

**settings:** `*mailchimpmarketinggosdk.CreateCampaignsRequestSettings` — The settings for your campaign, including subject, from name, reply-to address, and more.
    
</dd>
</dl>

<dl>
<dd>

**socialCard:** `*mailchimpmarketinggosdk.CreateCampaignsRequestSocialCard` — The preview for the campaign, rendered by social networks like Facebook and Twitter. [Learn more](https://mailchimp.com/help/enable-and-customize-social-cards/).
    
</dd>
</dl>

<dl>
<dd>

**tracking:** `*mailchimpmarketinggosdk.CampaignTrackingOptions` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*mailchimpmarketinggosdk.CreateCampaignsRequestType` — There are four types of [campaigns](https://mailchimp.com/help/getting-started-with-campaigns/) you can create in Mailchimp. A/B Split campaigns have been deprecated and variate campaigns should be used instead.
    
</dd>
</dl>

<dl>
<dd>

**variateSettings:** `*mailchimpmarketinggosdk.CreateCampaignsRequestVariateSettings` — The settings specific to A/B test campaigns.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.Get(CampaignID) -> *mailchimpmarketinggosdk.Campaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**includeResendShortcutEligibility:** `*bool` — Return the `resend_shortcut_eligibility` field in the response, which tells you if the campaign is eligible for the various Campaign Resend Shortcuts offered.
    
</dd>
</dl>

<dl>
<dd>

**includeResendShortcutUsage:** `*bool` — Return the `resend_shortcut_usage` field in the response.  This includes information about campaigns related by a shortcut.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.Delete(CampaignID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a campaign from your Mailchimp account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.Update(CampaignID, request) -> *mailchimpmarketinggosdk.Campaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update some or all of the settings for a specific campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**recipients:** `*mailchimpmarketinggosdk.UpdateCampaignsRequestRecipients` — List settings for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**rssOpts:** `*mailchimpmarketinggosdk.UpdateCampaignsRequestRssOpts` — [RSS](https://mailchimp.com/help/share-your-blog-posts-with-mailchimp/) options for a campaign.
    
</dd>
</dl>

<dl>
<dd>

**settings:** `*mailchimpmarketinggosdk.UpdateCampaignsRequestSettings` — The settings for your campaign, including subject, from name, reply-to address, and more.
    
</dd>
</dl>

<dl>
<dd>

**socialCard:** `*mailchimpmarketinggosdk.UpdateCampaignsRequestSocialCard` — The preview for the campaign, rendered by social networks like Facebook and Twitter. [Learn more](https://mailchimp.com/help/enable-and-customize-social-cards/).
    
</dd>
</dl>

<dl>
<dd>

**tracking:** `*mailchimpmarketinggosdk.CampaignTrackingOptions` 
    
</dd>
</dl>

<dl>
<dd>

**variateSettings:** `*mailchimpmarketinggosdk.UpdateCampaignsRequestVariateSettings` — The settings specific to A/B test campaigns.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.CreateActionCancelSend(CampaignID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancel a Regular or Plain-Text Campaign after you send, before all of your recipients receive it. This feature is included with Mailchimp Pro.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionCancelSendCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.CreateActionCancelSend(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.CreateActionCreateResend(CampaignID, request) -> *mailchimpmarketinggosdk.Campaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove the guesswork for resending a campaign to certain segments. You can use this endpoint as a shortcut to replicate a campaign and resend it to common segments, such as those who didn't open the campaign, or any new subscribers since it was sent.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionCreateResendCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.CreateActionCreateResend(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**shortcutType:** `*mailchimpmarketinggosdk.CreateActionCreateResendCampaignsRequestShortcutType` — Which campaign resend shortcut to use. Default is `to_non_openers`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.CreateActionPause(CampaignID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pause an RSS-Driven campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionPauseCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.CreateActionPause(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.CreateActionReplicate(CampaignID) -> *mailchimpmarketinggosdk.Campaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replicate a campaign in saved or send status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionReplicateCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.CreateActionReplicate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.CreateActionResume(CampaignID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resume an RSS-Driven campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionResumeCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.CreateActionResume(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.CreateActionSchedule(CampaignID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Schedule a campaign for delivery. If you're using Multivariate Campaigns to test send times or sending RSS Campaigns, use the send action instead.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionScheduleCampaignsRequest{
    CampaignID: "campaign_id",
    ScheduleTime: mailchimpmarketinggosdk.MustParseDateTime(
        "2024-01-15T09:30:00Z",
    ),
}
client.Campaigns.CreateActionSchedule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**batchDelivery:** `*mailchimpmarketinggosdk.CreateActionScheduleCampaignsRequestBatchDelivery` — Choose whether the campaign should use [Batch Delivery](https://mailchimp.com/help/schedule-batch-delivery/). Cannot be set to `true` for campaigns using [Timewarp](https://mailchimp.com/help/use-timewarp/).
    
</dd>
</dl>

<dl>
<dd>

**scheduleTime:** `time.Time` — The UTC date and time to schedule the campaign for delivery in ISO 8601 format. Campaigns may only be scheduled to send on the quarter-hour (:00, :15, :30, :45).
    
</dd>
</dl>

<dl>
<dd>

**timewarp:** `*bool` — Choose whether the campaign should use [Timewarp](https://mailchimp.com/help/use-timewarp/) when sending. Campaigns scheduled with Timewarp are localized based on the recipients' time zones. For example, a Timewarp campaign with a `schedule_time` of 13:00 will be sent to each recipient at 1:00pm in their local time. Cannot be set to `true` for campaigns using [Batch Delivery](https://mailchimp.com/help/schedule-batch-delivery/).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.CreateActionSend(CampaignID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Send a Mailchimp campaign. For RSS Campaigns, the campaign will send according to its schedule. All other campaigns will send immediately.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionSendCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.CreateActionSend(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.CreateActionTest(CampaignID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Send a test email.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionTestCampaignsRequest{
    CampaignID: "campaign_id",
    SendType: mailchimpmarketinggosdk.CreateActionTestCampaignsRequestSendTypeHTML,
    TestEmails: []string{
        "test_emails",
    },
}
client.Campaigns.CreateActionTest(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**sendType:** `*mailchimpmarketinggosdk.CreateActionTestCampaignsRequestSendType` — Choose the type of test email to send.
    
</dd>
</dl>

<dl>
<dd>

**testEmails:** `[]string` — An array of email addresses to send the test email to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.CreateActionUnschedule(CampaignID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unschedule a scheduled campaign that hasn't started sending.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionUnscheduleCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.CreateActionUnschedule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.GetContent(CampaignID) -> *mailchimpmarketinggosdk.CampaignContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the the HTML and plain-text content for a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetContentCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.GetContent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.UpsertContent(CampaignID, request) -> *mailchimpmarketinggosdk.CampaignContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set the content for a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpsertContentCampaignsRequest{
    CampaignID: "campaign_id",
    Body: &mailchimpmarketinggosdk.CampaignContent{},
}
client.Campaigns.UpsertContent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*mailchimpmarketinggosdk.CampaignContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.ListFeedback(CampaignID) -> *mailchimpmarketinggosdk.ListFeedbackCampaignsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get team feedback while you're working together on a Mailchimp campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListFeedbackCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.ListFeedback(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.CreateFeedback(CampaignID, request) -> *mailchimpmarketinggosdk.CreateFeedbackCampaignsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add feedback on a specific campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateFeedbackCampaignsRequest{
    CampaignID: "campaign_id",
    Message: "message",
}
client.Campaigns.CreateFeedback(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**blockID:** `*int` — The block id for the editable block that the feedback addresses.
    
</dd>
</dl>

<dl>
<dd>

**isComplete:** `*bool` — The status of feedback.
    
</dd>
</dl>

<dl>
<dd>

**message:** `string` — The content of the feedback.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.GetFeedback(CampaignID, FeedbackID) -> *mailchimpmarketinggosdk.CampaignFeedback</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific feedback message from a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetFeedbackCampaignsRequest{
    CampaignID: "campaign_id",
    FeedbackID: "feedback_id",
}
client.Campaigns.GetFeedback(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**feedbackID:** `string` — The unique id for the feedback message.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.DeleteFeedback(CampaignID, FeedbackID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a specific feedback message for a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteFeedbackCampaignsRequest{
    CampaignID: "campaign_id",
    FeedbackID: "feedback_id",
}
client.Campaigns.DeleteFeedback(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**feedbackID:** `string` — The unique id for the feedback message.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.UpdateFeedback(CampaignID, FeedbackID, request) -> *mailchimpmarketinggosdk.CampaignFeedback</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific feedback message for a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateFeedbackCampaignsRequest{
    CampaignID: "campaign_id",
    FeedbackID: "feedback_id",
}
client.Campaigns.UpdateFeedback(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**feedbackID:** `string` — The unique id for the feedback message.
    
</dd>
</dl>

<dl>
<dd>

**blockID:** `*int` — The block id for the editable block that the feedback addresses.
    
</dd>
</dl>

<dl>
<dd>

**isComplete:** `*bool` — The status of feedback.
    
</dd>
</dl>

<dl>
<dd>

**message:** `*string` — The content of the feedback.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Campaigns.ListSendChecklist(CampaignID) -> *mailchimpmarketinggosdk.ListSendChecklistCampaignsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Review the send checklist for a campaign, and resolve any issues before sending.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSendChecklistCampaignsRequest{
    CampaignID: "campaign_id",
}
client.Campaigns.ListSendChecklist(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ConnectedSites
<details><summary><code>client.ConnectedSites.List() -> *mailchimpmarketinggosdk.ListConnectedSitesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all connected sites in an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListConnectedSitesRequest{}
client.ConnectedSites.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConnectedSites.Create(request) -> *mailchimpmarketinggosdk.ConnectedSite</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new Mailchimp connected site.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateConnectedSitesRequest{
    Domain: "example.com",
    ForeignID: "MC001",
}
client.ConnectedSites.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domain:** `string` — The connected site domain.
    
</dd>
</dl>

<dl>
<dd>

**foreignID:** `string` — The unique identifier for the site.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConnectedSites.Get(ConnectedSiteID) -> *mailchimpmarketinggosdk.ConnectedSite</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific connected site.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetConnectedSitesRequest{
    ConnectedSiteID: "connected_site_id",
}
client.ConnectedSites.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connectedSiteID:** `string` — The unique identifier for the site.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConnectedSites.Delete(ConnectedSiteID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a connected site from your Mailchimp account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteConnectedSitesRequest{
    ConnectedSiteID: "connected_site_id",
}
client.ConnectedSites.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connectedSiteID:** `string` — The unique identifier for the site.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConnectedSites.CreateActionVerifyScriptInstallation(ConnectedSiteID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Verify that the connected sites script has been installed, either via the script URL or fragment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionVerifyScriptInstallationConnectedSitesRequest{
    ConnectedSiteID: "connected_site_id",
}
client.ConnectedSites.CreateActionVerifyScriptInstallation(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connectedSiteID:** `string` — The unique identifier for the site.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## conversations
<details><summary><code>client.Conversations.List() -> *mailchimpmarketinggosdk.ListConversationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of conversations for the account. Conversations has been deprecated in favor of Inbox and these endpoints don't include Inbox data. Past Conversations are still available via this endpoint, but new campaign replies and other Inbox messages aren’t available using this endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListConversationsRequest{}
client.Conversations.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**hasUnreadMessages:** `*mailchimpmarketinggosdk.ListConversationsRequestHasUnreadMessages` — Whether the conversation has any unread messages.
    
</dd>
</dl>

<dl>
<dd>

**listID:** `*string` — The unique id for the list.
    
</dd>
</dl>

<dl>
<dd>

**campaignID:** `*string` — The unique id for the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.Get(ConversationID) -> *mailchimpmarketinggosdk.Conversation</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details about an individual conversation. Conversations has been deprecated in favor of Inbox and these endpoints don't include Inbox data. Past Conversations are still available via this endpoint, but new campaign replies and other Inbox messages aren’t available using this endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetConversationsRequest{
    ConversationID: "conversation_id",
}
client.Conversations.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationID:** `string` — The unique id for the conversation.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.ListMessages(ConversationID) -> *mailchimpmarketinggosdk.ListMessagesConversationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get messages from a specific conversation. Conversations has been deprecated in favor of Inbox and these endpoints don't include Inbox data. Past Conversations are still available via this endpoint, but new campaign replies and other Inbox messages aren’t available using this endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListMessagesConversationsRequest{
    ConversationID: "conversation_id",
}
client.Conversations.ListMessages(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationID:** `string` — The unique id for the conversation.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**isRead:** `*mailchimpmarketinggosdk.ListMessagesConversationsRequestIsRead` — Whether a conversation message has been marked as read.
    
</dd>
</dl>

<dl>
<dd>

**beforeTimestamp:** `*time.Time` — Restrict the response to messages created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceTimestamp:** `*time.Time` — Restrict the response to messages created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetMessage(ConversationID, MessageID) -> *mailchimpmarketinggosdk.ConversationMessage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get an individual message in a conversation. Conversations has been deprecated in favor of Inbox and these endpoints don't include Inbox data. Past Conversations are still available via this endpoint, but new campaign replies and other Inbox messages aren’t available using this endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetMessageConversationsRequest{
    ConversationID: "conversation_id",
    MessageID: "message_id",
}
client.Conversations.GetMessage(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationID:** `string` — The unique id for the conversation.
    
</dd>
</dl>

<dl>
<dd>

**messageID:** `string` — The unique id for the conversation message.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CustomerJourneys
<details><summary><code>client.CustomerJourneys.CreateJourneyStepActionTrigger(JourneyID, StepID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

A step trigger in an Automation flow. To use it, create a starting point or step from the Automation flow builder in the app using the Customer Journeys API condition. We’ll provide a url during the process that includes the {journey_id} and {step_id}. You’ll then be able to use this endpoint to trigger the condition for the posted contact.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateJourneyStepActionTriggerCustomerJourneysRequest{
    JourneyID: 1,
    StepID: 1,
    EmailAddress: "email_address",
}
client.CustomerJourneys.CreateJourneyStepActionTrigger(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**journeyID:** `int` — The id for the flow.
    
</dd>
</dl>

<dl>
<dd>

**stepID:** `int` — The id for the Step.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `string` — The list member's email address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ecommerce
<details><summary><code>client.Ecommerce.List() -> *mailchimpmarketinggosdk.ListEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about the e-commerce endpoint's resources.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ecommerce.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListOrders() -> *mailchimpmarketinggosdk.ListOrdersEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about an account's orders.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListOrdersEcommerceRequest{}
client.Ecommerce.ListOrders(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**campaignID:** `*string` — Restrict results to orders with a specific `campaign_id` value.
    
</dd>
</dl>

<dl>
<dd>

**outreachID:** `*string` — Restrict results to orders with a specific `outreach_id` value.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `*string` — Restrict results to orders made by a specific customer.
    
</dd>
</dl>

<dl>
<dd>

**hasOutreach:** `*bool` — Restrict results to orders that have an outreach attached. For example, an email campaign or Facebook ad.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStores() -> *mailchimpmarketinggosdk.ListStoresEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about all stores in the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStoresEcommerceRequest{}
client.Ecommerce.ListStores(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStore(request) -> *mailchimpmarketinggosdk.ECommerceStore</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new store to your Mailchimp account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStoreEcommerceRequest{
    CurrencyCode: "USD",
    ID: "example_store",
    ListID: "1a2df69511",
    Name: "Freddie's Cat Hat Emporium",
}
client.Ecommerce.CreateStore(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**address:** `*mailchimpmarketinggosdk.CreateStoreEcommerceRequestAddress` — The store address.
    
</dd>
</dl>

<dl>
<dd>

**currencyCode:** `string` — The three-letter ISO 4217 code for the currency that the store accepts.
    
</dd>
</dl>

<dl>
<dd>

**domain:** `*string` — The store domain. This parameter is required for Connected Sites and Google Ads.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` — The email address for the store.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — The unique identifier for the store.
    
</dd>
</dl>

<dl>
<dd>

**isSyncing:** `*bool` — Whether to disable automations because the store is currently [syncing](https://mailchimp.com/developer/marketing/docs/e-commerce/#pausing-store-automations).
    
</dd>
</dl>

<dl>
<dd>

**listID:** `string` — The unique identifier for the list associated with the store. The `list_id` for a specific store cannot change.
    
</dd>
</dl>

<dl>
<dd>

**moneyFormat:** `*string` — The currency format for the store. For example: `$`, `£`, etc.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the store.
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` — The store phone number.
    
</dd>
</dl>

<dl>
<dd>

**platform:** `*string` — The e-commerce platform of the store.
    
</dd>
</dl>

<dl>
<dd>

**primaryLocale:** `*string` — The primary locale for the store. For example: `en`, `de`, etc.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — The timezone for the store.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStore(StoreID) -> *mailchimpmarketinggosdk.ECommerceStore</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStoreEcommerceRequest{
    StoreID: "store_id",
}
client.Ecommerce.GetStore(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStore(StoreID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a store. Deleting a store will also delete any associated subresources, including Customers, Orders, Products, and Carts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStoreEcommerceRequest{
    StoreID: "store_id",
}
client.Ecommerce.DeleteStore(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStore(StoreID, request) -> *mailchimpmarketinggosdk.ECommerceStore</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStoreEcommerceRequest{
    StoreID: "store_id",
}
client.Ecommerce.UpdateStore(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*mailchimpmarketinggosdk.UpdateStoreEcommerceRequestAddress` — The store address.
    
</dd>
</dl>

<dl>
<dd>

**currencyCode:** `*string` — The three-letter ISO 4217 code for the currency that the store accepts.
    
</dd>
</dl>

<dl>
<dd>

**domain:** `*string` — The store domain.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` — The email address for the store.
    
</dd>
</dl>

<dl>
<dd>

**isSyncing:** `*bool` — Whether to disable automations because the store is currently [syncing](https://mailchimp.com/developer/marketing/docs/e-commerce/#pausing-store-automations).
    
</dd>
</dl>

<dl>
<dd>

**moneyFormat:** `*string` — The currency format for the store. For example: `$`, `£`, etc.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the store.
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` — The store phone number.
    
</dd>
</dl>

<dl>
<dd>

**platform:** `*string` — The e-commerce platform of the store.
    
</dd>
</dl>

<dl>
<dd>

**primaryLocale:** `*string` — The primary locale for the store. For example: `en`, `de`, etc.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — The timezone for the store.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStoreCarts(StoreID) -> *mailchimpmarketinggosdk.ListStoreCartsEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a store's carts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStoreCartsEcommerceRequest{
    StoreID: "store_id",
}
client.Ecommerce.ListStoreCarts(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStoreCart(StoreID, request) -> *mailchimpmarketinggosdk.ECommerceCart</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new cart to a store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStoreCartEcommerceRequest{
    StoreID: "store_id",
    CurrencyCode: "currency_code",
    Customer: &mailchimpmarketinggosdk.EcommerceStoresCartsPost{
        ID: "id",
    },
    ID: &mailchimpmarketinggosdk.CreateStoreCartEcommerceRequestID{
        String: "id",
    },
    Lines: []*mailchimpmarketinggosdk.CreateStoreCartEcommerceRequestLinesItem{
        &mailchimpmarketinggosdk.CreateStoreCartEcommerceRequestLinesItem{
            ID: "id",
            Price: &mailchimpmarketinggosdk.CreateStoreCartEcommerceRequestLinesItemPrice{
                Double: 1.1,
            },
            ProductID: "product_id",
            ProductVariantID: "product_variant_id",
            Quantity: 1,
        },
    },
    OrderTotal: &mailchimpmarketinggosdk.CreateStoreCartEcommerceRequestOrderTotal{
        Double: 1.1,
    },
}
client.Ecommerce.CreateStoreCart(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**campaignID:** `*string` — A string that uniquely identifies the campaign for a cart.
    
</dd>
</dl>

<dl>
<dd>

**checkoutURL:** `*string` — The URL for the cart. This parameter is required for [Abandoned Cart](https://mailchimp.com/help/create-a-classic-abandoned-cart-email/) automations.
    
</dd>
</dl>

<dl>
<dd>

**currencyCode:** `string` — The three-letter ISO 4217 code for the currency that the cart uses.
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*mailchimpmarketinggosdk.EcommerceStoresCartsPost` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `*mailchimpmarketinggosdk.CreateStoreCartEcommerceRequestID` — A unique identifier for the cart.
    
</dd>
</dl>

<dl>
<dd>

**lines:** `[]*mailchimpmarketinggosdk.CreateStoreCartEcommerceRequestLinesItem` — An array of the cart's line items.
    
</dd>
</dl>

<dl>
<dd>

**orderTotal:** `*mailchimpmarketinggosdk.CreateStoreCartEcommerceRequestOrderTotal` 
    
</dd>
</dl>

<dl>
<dd>

**taxTotal:** `*mailchimpmarketinggosdk.CreateStoreCartEcommerceRequestTaxTotal` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStoreCart(StoreID, CartID) -> *mailchimpmarketinggosdk.ECommerceCart</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific cart.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStoreCartEcommerceRequest{
    StoreID: "store_id",
    CartID: "cart_id",
}
client.Ecommerce.GetStoreCart(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**cartID:** `string` — The id for the cart.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStoreCart(StoreID, CartID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a cart.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStoreCartEcommerceRequest{
    StoreID: "store_id",
    CartID: "cart_id",
}
client.Ecommerce.DeleteStoreCart(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**cartID:** `string` — The id for the cart.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStoreCart(StoreID, CartID, request) -> *mailchimpmarketinggosdk.ECommerceCart</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific cart.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStoreCartEcommerceRequest{
    StoreID: "store_id",
    CartID: "cart_id",
}
client.Ecommerce.UpdateStoreCart(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**cartID:** `string` — The id for the cart.
    
</dd>
</dl>

<dl>
<dd>

**campaignID:** `*string` — A string that uniquely identifies the campaign associated with a cart.
    
</dd>
</dl>

<dl>
<dd>

**checkoutURL:** `*string` — The URL for the cart. This parameter is required for [Abandoned Cart](https://mailchimp.com/help/create-a-classic-abandoned-cart-email/) automations.
    
</dd>
</dl>

<dl>
<dd>

**currencyCode:** `*string` — The three-letter ISO 4217 code for the currency that the cart uses.
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*mailchimpmarketinggosdk.EcommerceStoresCartsPatch` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `*mailchimpmarketinggosdk.UpdateStoreCartEcommerceRequestID` — A unique identifier for the cart.
    
</dd>
</dl>

<dl>
<dd>

**lines:** `[]*mailchimpmarketinggosdk.UpdateStoreCartEcommerceRequestLinesItem` — An array of the cart's line items.
    
</dd>
</dl>

<dl>
<dd>

**orderTotal:** `*mailchimpmarketinggosdk.UpdateStoreCartEcommerceRequestOrderTotal` 
    
</dd>
</dl>

<dl>
<dd>

**taxTotal:** `*mailchimpmarketinggosdk.UpdateStoreCartEcommerceRequestTaxTotal` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStoreCartLines(StoreID, CartID) -> *mailchimpmarketinggosdk.ListStoreCartLinesEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a cart's line items.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStoreCartLinesEcommerceRequest{
    StoreID: "store_id",
    CartID: "cart_id",
}
client.Ecommerce.ListStoreCartLines(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**cartID:** `string` — The id for the cart.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStoreCartLine(StoreID, CartID, request) -> *mailchimpmarketinggosdk.ECommerceCartLineItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new line item to an existing cart.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStoreCartLineEcommerceRequest{
    StoreID: "store_id",
    CartID: "cart_id",
    ID: "id",
    Price: &mailchimpmarketinggosdk.CreateStoreCartLineEcommerceRequestPrice{
        Double: 1.1,
    },
    ProductID: "product_id",
    ProductVariantID: "product_variant_id",
    Quantity: 1,
}
client.Ecommerce.CreateStoreCartLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**cartID:** `string` — The id for the cart.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — A unique identifier for the cart line item.
    
</dd>
</dl>

<dl>
<dd>

**price:** `*mailchimpmarketinggosdk.CreateStoreCartLineEcommerceRequestPrice` 
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — A unique identifier for the product associated with the cart line item.
    
</dd>
</dl>

<dl>
<dd>

**productVariantID:** `string` — A unique identifier for the product variant associated with the cart line item.
    
</dd>
</dl>

<dl>
<dd>

**quantity:** `int` — The quantity of a cart line item.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStoreCartLine(StoreID, CartID, LineID) -> *mailchimpmarketinggosdk.ECommerceCartLineItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific cart line item.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStoreCartLineEcommerceRequest{
    StoreID: "store_id",
    CartID: "cart_id",
    LineID: "line_id",
}
client.Ecommerce.GetStoreCartLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**cartID:** `string` — The id for the cart.
    
</dd>
</dl>

<dl>
<dd>

**lineID:** `string` — The id for the line item of a cart.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStoreCartLine(StoreID, CartID, LineID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific cart line item.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStoreCartLineEcommerceRequest{
    StoreID: "store_id",
    CartID: "cart_id",
    LineID: "line_id",
}
client.Ecommerce.DeleteStoreCartLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**cartID:** `string` — The id for the cart.
    
</dd>
</dl>

<dl>
<dd>

**lineID:** `string` — The id for the line item of a cart.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStoreCartLine(StoreID, CartID, LineID, request) -> *mailchimpmarketinggosdk.ECommerceCartLineItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific cart line item.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStoreCartLineEcommerceRequest{
    StoreID: "store_id",
    CartID: "cart_id",
    LineID: "line_id",
}
client.Ecommerce.UpdateStoreCartLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**cartID:** `string` — The id for the cart.
    
</dd>
</dl>

<dl>
<dd>

**lineID:** `string` — The id for the line item of a cart.
    
</dd>
</dl>

<dl>
<dd>

**price:** `*mailchimpmarketinggosdk.UpdateStoreCartLineEcommerceRequestPrice` 
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` — A unique identifier for the product associated with the cart line item.
    
</dd>
</dl>

<dl>
<dd>

**productVariantID:** `*string` — A unique identifier for the product variant associated with the cart line item.
    
</dd>
</dl>

<dl>
<dd>

**quantity:** `*int` — The quantity of a cart line item.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStoreCustomers(StoreID) -> *mailchimpmarketinggosdk.ListStoreCustomersEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a store's customers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStoreCustomersEcommerceRequest{
    StoreID: "store_id",
}
client.Ecommerce.ListStoreCustomers(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` — Restrict the response to customers with the email address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStoreCustomer(StoreID, request) -> *mailchimpmarketinggosdk.ECommerceCustomer</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new customer to a store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStoreCustomerEcommerceRequest{
    StoreID: "store_id",
    ID: "id",
    OptInStatus: true,
}
client.Ecommerce.CreateStoreCustomer(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*mailchimpmarketinggosdk.CreateStoreCustomerEcommerceRequestAddress` — The customer's address.
    
</dd>
</dl>

<dl>
<dd>

**company:** `*string` — The customer's company.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` — The customer's email address.
    
</dd>
</dl>

<dl>
<dd>

**firstName:** `*string` — The customer's first name.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — A unique identifier for the customer. Limited to 50 characters.
    
</dd>
</dl>

<dl>
<dd>

**lastName:** `*string` — The customer's last name.
    
</dd>
</dl>

<dl>
<dd>

**optInStatus:** `bool` — The customer's opt-in status. This value will never overwrite the opt-in status of a pre-existing Mailchimp list member, but will apply to list members that are added through the e-commerce API endpoints. Customers who don't opt in to your Mailchimp list [will be added as `Transactional` members](https://mailchimp.com/developer/marketing/docs/e-commerce/#customers).
    
</dd>
</dl>

<dl>
<dd>

**smsPhoneNumber:** `*string` — A US phone number for SMS contact.
    
</dd>
</dl>

<dl>
<dd>

**totalSpent:** `*mailchimpmarketinggosdk.CreateStoreCustomerEcommerceRequestTotalSpent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStoreCustomer(StoreID, CustomerID) -> *mailchimpmarketinggosdk.ECommerceCustomer</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific customer.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStoreCustomerEcommerceRequest{
    StoreID: "store_id",
    CustomerID: "customer_id",
}
client.Ecommerce.GetStoreCustomer(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `string` — The id for the customer of a store.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpsertStoreCustomer(StoreID, CustomerID, request) -> *mailchimpmarketinggosdk.ECommerceCustomer</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add or update a customer.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpsertStoreCustomerEcommerceRequest{
    StoreID: "store_id",
    CustomerID: "customer_id",
}
client.Ecommerce.UpsertStoreCustomer(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `string` — The id for the customer of a store.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*mailchimpmarketinggosdk.UpsertStoreCustomerEcommerceRequestAddress` — The customer's address.
    
</dd>
</dl>

<dl>
<dd>

**company:** `*string` — The customer's company.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` — The customer's email address.
    
</dd>
</dl>

<dl>
<dd>

**firstName:** `*string` — The customer's first name.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` — A unique identifier for the customer. Limited to 50 characters.
    
</dd>
</dl>

<dl>
<dd>

**lastName:** `*string` — The customer's last name.
    
</dd>
</dl>

<dl>
<dd>

**optInStatus:** `*bool` — The customer's opt-in status. This value will never overwrite the opt-in status of a pre-existing Mailchimp list member, but will apply to list members that are added through the e-commerce API endpoints. Customers who don't opt in to your Mailchimp list [will be added as `Transactional` members](https://mailchimp.com/developer/marketing/docs/e-commerce/#customers).
    
</dd>
</dl>

<dl>
<dd>

**smsPhoneNumber:** `*string` — A US phone number for SMS contact.
    
</dd>
</dl>

<dl>
<dd>

**totalSpent:** `*mailchimpmarketinggosdk.UpsertStoreCustomerEcommerceRequestTotalSpent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStoreCustomer(StoreID, CustomerID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a customer from a store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStoreCustomerEcommerceRequest{
    StoreID: "store_id",
    CustomerID: "customer_id",
}
client.Ecommerce.DeleteStoreCustomer(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `string` — The id for the customer of a store.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStoreCustomer(StoreID, CustomerID, request) -> *mailchimpmarketinggosdk.ECommerceCustomer</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a customer.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStoreCustomerEcommerceRequest{
    StoreID: "store_id",
    CustomerID: "customer_id",
    Body: &mailchimpmarketinggosdk.EcommerceStoresCartsPatch{},
}
client.Ecommerce.UpdateStoreCustomer(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `string` — The id for the customer of a store.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*mailchimpmarketinggosdk.EcommerceStoresCartsPatch` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStoreOrders(StoreID) -> *mailchimpmarketinggosdk.ListStoreOrdersEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a store's orders.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStoreOrdersEcommerceRequest{
    StoreID: "store_id",
}
client.Ecommerce.ListStoreOrders(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**customerID:** `*string` — Restrict results to orders made by a specific customer.
    
</dd>
</dl>

<dl>
<dd>

**hasOutreach:** `*bool` — Restrict results to orders that have an outreach attached. For example, an email campaign or Facebook ad.
    
</dd>
</dl>

<dl>
<dd>

**campaignID:** `*string` — Restrict results to orders with a specific `campaign_id` value.
    
</dd>
</dl>

<dl>
<dd>

**outreachID:** `*string` — Restrict results to orders with a specific `outreach_id` value.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStoreOrder(StoreID, request) -> *mailchimpmarketinggosdk.ECommerceOrder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new order to a store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequest{
    StoreID: "store_id",
    CurrencyCode: "currency_code",
    Customer: &mailchimpmarketinggosdk.EcommerceStoresCartsPost{
        ID: "id",
    },
    ID: "id",
    Lines: []*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestLinesItem{
        &mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestLinesItem{
            ID: "id",
            Price: &mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestLinesItemPrice{
                Double: 1.1,
            },
            ProductID: "product_id",
            ProductVariantID: "product_variant_id",
            Quantity: 1,
        },
    },
    OrderTotal: &mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestOrderTotal{
        Double: 1.1,
    },
}
client.Ecommerce.CreateStoreOrder(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**billingAddress:** `*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestBillingAddress` — The billing address for the order.
    
</dd>
</dl>

<dl>
<dd>

**campaignID:** `*string` — A string that uniquely identifies the campaign for an order.
    
</dd>
</dl>

<dl>
<dd>

**cartID:** `*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestCartID` — A cart id that the order was placed for.
    
</dd>
</dl>

<dl>
<dd>

**cancelledAtForeign:** `*string` — The date and time the order was cancelled in ISO 8601 format. Note: passing a value for this parameter will cancel the order being created.
    
</dd>
</dl>

<dl>
<dd>

**currencyCode:** `string` — The three-letter ISO 4217 code for the currency that the store accepts.
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*mailchimpmarketinggosdk.EcommerceStoresCartsPost` 
    
</dd>
</dl>

<dl>
<dd>

**discountTotal:** `*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestDiscountTotal` 
    
</dd>
</dl>

<dl>
<dd>

**financialStatus:** `*string` — The order status. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
    
</dd>
</dl>

<dl>
<dd>

**fulfillmentStatus:** `*string` — The fulfillment status for the order. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — A unique identifier for the order.
    
</dd>
</dl>

<dl>
<dd>

**landingSite:** `*string` — The URL for the page where the buyer landed when entering the shop.
    
</dd>
</dl>

<dl>
<dd>

**lines:** `[]*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestLinesItem` — An array of the order's line items.
    
</dd>
</dl>

<dl>
<dd>

**orderTotal:** `*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestOrderTotal` 
    
</dd>
</dl>

<dl>
<dd>

**orderURL:** `*string` — The URL for the order.
    
</dd>
</dl>

<dl>
<dd>

**outreach:** `*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestOutreach` — The outreach associated with this order. For example, an email campaign or Facebook ad.
    
</dd>
</dl>

<dl>
<dd>

**processedAtForeign:** `*string` — The date and time the order was processed in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**promos:** `[]*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestPromosItem` — The promo codes applied on the order
    
</dd>
</dl>

<dl>
<dd>

**shippingAddress:** `*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestShippingAddress` — The shipping address for the order.
    
</dd>
</dl>

<dl>
<dd>

**shippingTotal:** `*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestShippingTotal` 
    
</dd>
</dl>

<dl>
<dd>

**taxTotal:** `*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestTaxTotal` 
    
</dd>
</dl>

<dl>
<dd>

**trackingCarrier:** `*string` — The tracking carrier associated with the order.
    
</dd>
</dl>

<dl>
<dd>

**trackingCode:** `*mailchimpmarketinggosdk.CreateStoreOrderEcommerceRequestTrackingCode` — The Mailchimp tracking code for the order. Uses the 'mc_tc' parameter in E-Commerce tracking URLs.
    
</dd>
</dl>

<dl>
<dd>

**trackingNumber:** `*string` — The tracking number associated with the order.
    
</dd>
</dl>

<dl>
<dd>

**trackingURL:** `*string` — The tracking URL associated with the order.
    
</dd>
</dl>

<dl>
<dd>

**updatedAtForeign:** `*string` — The date and time the order was updated in ISO 8601 format.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStoreOrder(StoreID, OrderID) -> *mailchimpmarketinggosdk.ECommerceOrder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStoreOrderEcommerceRequest{
    StoreID: "store_id",
    OrderID: "order_id",
}
client.Ecommerce.GetStoreOrder(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `string` — The id for the order in a store.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStoreOrder(StoreID, OrderID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete an order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStoreOrderEcommerceRequest{
    StoreID: "store_id",
    OrderID: "order_id",
}
client.Ecommerce.DeleteStoreOrder(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `string` — The id for the order in a store.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStoreOrder(StoreID, OrderID, request) -> *mailchimpmarketinggosdk.ECommerceOrder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequest{
    StoreID: "store_id",
    OrderID: "order_id",
}
client.Ecommerce.UpdateStoreOrder(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `string` — The id for the order in a store.
    
</dd>
</dl>

<dl>
<dd>

**billingAddress:** `*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestBillingAddress` — The billing address for the order.
    
</dd>
</dl>

<dl>
<dd>

**campaignID:** `*string` — A string that uniquely identifies the campaign associated with an order.
    
</dd>
</dl>

<dl>
<dd>

**cartID:** `*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestCartID` — A cart id that the order was placed for.
    
</dd>
</dl>

<dl>
<dd>

**cancelledAtForeign:** `*string` — The date and time the order was cancelled in ISO 8601 format. Note: passing a value for this parameter will cancel the order being edited.
    
</dd>
</dl>

<dl>
<dd>

**currencyCode:** `*string` — The three-letter ISO 4217 code for the currency that the store accepts.
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*mailchimpmarketinggosdk.EcommerceStoresCartsPatch` 
    
</dd>
</dl>

<dl>
<dd>

**discountTotal:** `*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestDiscountTotal` 
    
</dd>
</dl>

<dl>
<dd>

**financialStatus:** `*string` — The order status. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
    
</dd>
</dl>

<dl>
<dd>

**fulfillmentStatus:** `*string` — The fulfillment status for the order. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` — A unique identifier for the order.
    
</dd>
</dl>

<dl>
<dd>

**landingSite:** `*string` — The URL for the page where the buyer landed when entering the shop.
    
</dd>
</dl>

<dl>
<dd>

**lines:** `[]*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestLinesItem` — An array of the order's line items.
    
</dd>
</dl>

<dl>
<dd>

**orderTotal:** `*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestOrderTotal` 
    
</dd>
</dl>

<dl>
<dd>

**orderURL:** `*string` — The URL for the order.
    
</dd>
</dl>

<dl>
<dd>

**outreach:** `*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestOutreach` — The outreach associated with this order. For example, an email campaign or Facebook ad.
    
</dd>
</dl>

<dl>
<dd>

**processedAtForeign:** `*string` — The date and time the order was processed in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**promos:** `[]*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestPromosItem` — The promo codes applied on the order. Note: Patch will completely replace the value of promos with the new one provided.
    
</dd>
</dl>

<dl>
<dd>

**shippingAddress:** `*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestShippingAddress` — The shipping address for the order.
    
</dd>
</dl>

<dl>
<dd>

**shippingTotal:** `*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestShippingTotal` 
    
</dd>
</dl>

<dl>
<dd>

**taxTotal:** `*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestTaxTotal` 
    
</dd>
</dl>

<dl>
<dd>

**trackingCarrier:** `*string` — The tracking carrier associated with the order.
    
</dd>
</dl>

<dl>
<dd>

**trackingCode:** `*mailchimpmarketinggosdk.UpdateStoreOrderEcommerceRequestTrackingCode` — The Mailchimp tracking code for the order. Uses the 'mc_tc' parameter in E-Commerce tracking URLs.
    
</dd>
</dl>

<dl>
<dd>

**trackingNumber:** `*string` — The tracking number associated with the order.
    
</dd>
</dl>

<dl>
<dd>

**trackingURL:** `*string` — The tracking URL associated with the order.
    
</dd>
</dl>

<dl>
<dd>

**updatedAtForeign:** `*string` — The date and time the order was updated in ISO 8601 format.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStoreOrderLines(StoreID, OrderID) -> *mailchimpmarketinggosdk.ListStoreOrderLinesEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about an order's line items.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStoreOrderLinesEcommerceRequest{
    StoreID: "store_id",
    OrderID: "order_id",
}
client.Ecommerce.ListStoreOrderLines(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `string` — The id for the order in a store.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStoreOrderLine(StoreID, OrderID, request) -> *mailchimpmarketinggosdk.ECommerceOrderLineItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new line item to an existing order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStoreOrderLineEcommerceRequest{
    StoreID: "store_id",
    OrderID: "order_id",
    ID: "id",
    Price: &mailchimpmarketinggosdk.CreateStoreOrderLineEcommerceRequestPrice{
        Double: 1.1,
    },
    ProductID: "product_id",
    ProductVariantID: "product_variant_id",
    Quantity: 1,
}
client.Ecommerce.CreateStoreOrderLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `string` — The id for the order in a store.
    
</dd>
</dl>

<dl>
<dd>

**discount:** `*mailchimpmarketinggosdk.CreateStoreOrderLineEcommerceRequestDiscount` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — A unique identifier for the order line item.
    
</dd>
</dl>

<dl>
<dd>

**price:** `*mailchimpmarketinggosdk.CreateStoreOrderLineEcommerceRequestPrice` 
    
</dd>
</dl>

<dl>
<dd>

**product:** `*mailchimpmarketinggosdk.EcommerceStoresOrdersPost` 
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — A unique identifier for the product associated with the order line item.
    
</dd>
</dl>

<dl>
<dd>

**productVariantID:** `string` — A unique identifier for the product variant associated with the order line item.
    
</dd>
</dl>

<dl>
<dd>

**quantity:** `int` — The quantity of an order line item.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStoreOrderLine(StoreID, OrderID, LineID) -> *mailchimpmarketinggosdk.ECommerceOrderLineItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific order line item.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStoreOrderLineEcommerceRequest{
    StoreID: "store_id",
    OrderID: "order_id",
    LineID: "line_id",
}
client.Ecommerce.GetStoreOrderLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `string` — The id for the order in a store.
    
</dd>
</dl>

<dl>
<dd>

**lineID:** `string` — The id for the line item of an order.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStoreOrderLine(StoreID, OrderID, LineID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific order line item.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStoreOrderLineEcommerceRequest{
    StoreID: "store_id",
    OrderID: "order_id",
    LineID: "line_id",
}
client.Ecommerce.DeleteStoreOrderLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `string` — The id for the order in a store.
    
</dd>
</dl>

<dl>
<dd>

**lineID:** `string` — The id for the line item of an order.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStoreOrderLine(StoreID, OrderID, LineID, request) -> *mailchimpmarketinggosdk.ECommerceOrderLineItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific order line item.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStoreOrderLineEcommerceRequest{
    StoreID: "store_id",
    OrderID: "order_id",
    LineID: "line_id",
}
client.Ecommerce.UpdateStoreOrderLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `string` — The id for the order in a store.
    
</dd>
</dl>

<dl>
<dd>

**lineID:** `string` — The id for the line item of an order.
    
</dd>
</dl>

<dl>
<dd>

**discount:** `*mailchimpmarketinggosdk.UpdateStoreOrderLineEcommerceRequestDiscount` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` — A unique identifier for the order line item.
    
</dd>
</dl>

<dl>
<dd>

**price:** `*mailchimpmarketinggosdk.UpdateStoreOrderLineEcommerceRequestPrice` 
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` — A unique identifier for the product associated with the order line item.
    
</dd>
</dl>

<dl>
<dd>

**productVariantID:** `*string` — A unique identifier for the product variant associated with the order line item.
    
</dd>
</dl>

<dl>
<dd>

**quantity:** `*int` — The quantity of an order line item.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStoreProducts(StoreID) -> *mailchimpmarketinggosdk.ListStoreProductsEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a store's products.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStoreProductsEcommerceRequest{
    StoreID: "store_id",
}
client.Ecommerce.ListStoreProducts(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStoreProduct(StoreID, request) -> *mailchimpmarketinggosdk.ECommerceProduct</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new product to a store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStoreProductEcommerceRequest{
    StoreID: "store_id",
    Body: &mailchimpmarketinggosdk.EcommerceStoresOrdersPost{
        ID: &mailchimpmarketinggosdk.EcommerceStoresOrdersPostID{
            String: "id",
        },
        Title: "Cat Hat",
        Variants: []*mailchimpmarketinggosdk.EcommerceStoresOrdersPostVariantsItem{
            &mailchimpmarketinggosdk.EcommerceStoresOrdersPostVariantsItem{
                ID: &mailchimpmarketinggosdk.EcommerceStoresOrdersPostVariantsItemID{
                    String: "id",
                },
                Title: "Cat Hat",
            },
        },
    },
}
client.Ecommerce.CreateStoreProduct(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*mailchimpmarketinggosdk.EcommerceStoresOrdersPost` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStoreProduct(StoreID, ProductID) -> *mailchimpmarketinggosdk.ECommerceProduct</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific product.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStoreProductEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
}
client.Ecommerce.GetStoreProduct(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpsertStoreProduct(StoreID, ProductID, request) -> *mailchimpmarketinggosdk.ECommerceProduct</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific product.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpsertStoreProductEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
    ID: &mailchimpmarketinggosdk.UpsertStoreProductEcommerceRequestID{
        String: "id",
    },
}
client.Ecommerce.UpsertStoreProduct(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — The description of a product.
    
</dd>
</dl>

<dl>
<dd>

**handle:** `*string` — The handle of a product.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*mailchimpmarketinggosdk.UpsertStoreProductEcommerceRequestID` — A unique identifier for the product.
    
</dd>
</dl>

<dl>
<dd>

**imageURL:** `*string` — The image URL for a product.
    
</dd>
</dl>

<dl>
<dd>

**images:** `[]*mailchimpmarketinggosdk.UpsertStoreProductEcommerceRequestImagesItem` — An array of the product's images.
    
</dd>
</dl>

<dl>
<dd>

**publishedAtForeign:** `*string` — The date and time the product was published.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title of a product.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — The type of product.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL for a product.
    
</dd>
</dl>

<dl>
<dd>

**variants:** `[]*mailchimpmarketinggosdk.UpsertStoreProductEcommerceRequestVariantsItem` — An array of the product's variants. At least one variant is required for each product. A variant can use the same `id` and `title` as the parent product.
    
</dd>
</dl>

<dl>
<dd>

**vendor_:** `*string` — The vendor for a product.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStoreProduct(StoreID, ProductID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a product.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStoreProductEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
}
client.Ecommerce.DeleteStoreProduct(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStoreProduct(StoreID, ProductID, request) -> *mailchimpmarketinggosdk.ECommerceProduct</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific product.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStoreProductEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
}
client.Ecommerce.UpdateStoreProduct(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — The description of a product.
    
</dd>
</dl>

<dl>
<dd>

**handle:** `*string` — The handle of a product.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*mailchimpmarketinggosdk.UpdateStoreProductEcommerceRequestID` — A unique identifier for the product.
    
</dd>
</dl>

<dl>
<dd>

**imageURL:** `*string` — The image URL for a product.
    
</dd>
</dl>

<dl>
<dd>

**images:** `[]*mailchimpmarketinggosdk.UpdateStoreProductEcommerceRequestImagesItem` — An array of the product's images.
    
</dd>
</dl>

<dl>
<dd>

**publishedAtForeign:** `*string` — The date and time the product was published in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title of a product.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — The type of product.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL for a product.
    
</dd>
</dl>

<dl>
<dd>

**variants:** `[]*mailchimpmarketinggosdk.UpdateStoreProductEcommerceRequestVariantsItem` — An array of the product's variants. At least one variant is required for each product. A variant can use the same `id` and `title` as the parent product.
    
</dd>
</dl>

<dl>
<dd>

**vendor_:** `*string` — The vendor for a product.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStoreProductImages(StoreID, ProductID) -> *mailchimpmarketinggosdk.ListStoreProductImagesEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a product's images.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStoreProductImagesEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
}
client.Ecommerce.ListStoreProductImages(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStoreProductImage(StoreID, ProductID, request) -> *mailchimpmarketinggosdk.CreateStoreProductImageEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new image to the product.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStoreProductImageEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
    ID: "id",
    URL: "url",
}
client.Ecommerce.CreateStoreProductImage(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — A unique identifier for the product image.
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` — The URL for a product image.
    
</dd>
</dl>

<dl>
<dd>

**variantIDs:** `[]*mailchimpmarketinggosdk.CreateStoreProductImageEcommerceRequestVariantIDsItem` — The list of product variants using the image.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStoreProductImage(StoreID, ProductID, ImageID) -> *mailchimpmarketinggosdk.GetStoreProductImageEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific product image.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStoreProductImageEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
    ImageID: "image_id",
}
client.Ecommerce.GetStoreProductImage(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**imageID:** `string` — The id for the product image.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStoreProductImage(StoreID, ProductID, ImageID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a product image.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStoreProductImageEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
    ImageID: "image_id",
}
client.Ecommerce.DeleteStoreProductImage(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**imageID:** `string` — The id for the product image.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStoreProductImage(StoreID, ProductID, ImageID, request) -> *mailchimpmarketinggosdk.UpdateStoreProductImageEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a product image.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStoreProductImageEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
    ImageID: "image_id",
}
client.Ecommerce.UpdateStoreProductImage(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**imageID:** `string` — The id for the product image.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` — A unique identifier for the product image.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL for a product image.
    
</dd>
</dl>

<dl>
<dd>

**variantIDs:** `[]*mailchimpmarketinggosdk.UpdateStoreProductImageEcommerceRequestVariantIDsItem` — The list of product variants using the image.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStoreProductVariants(StoreID, ProductID) -> *mailchimpmarketinggosdk.ListStoreProductVariantsEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a product's variants.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStoreProductVariantsEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
}
client.Ecommerce.ListStoreProductVariants(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStoreProductVariant(StoreID, ProductID, request) -> *mailchimpmarketinggosdk.ECommerceProductVariant</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new variant to the product.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStoreProductVariantEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
    ID: &mailchimpmarketinggosdk.CreateStoreProductVariantEcommerceRequestID{
        String: "id",
    },
    Title: "Cat Hat",
}
client.Ecommerce.CreateStoreProductVariant(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**backorders:** `*string` — The backorders of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*mailchimpmarketinggosdk.CreateStoreProductVariantEcommerceRequestID` — A unique identifier for the product variant.
    
</dd>
</dl>

<dl>
<dd>

**imageURL:** `*string` — The image URL for a product variant.
    
</dd>
</dl>

<dl>
<dd>

**inventoryQuantity:** `*int` — The inventory quantity of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**price:** `*mailchimpmarketinggosdk.CreateStoreProductVariantEcommerceRequestPrice` 
    
</dd>
</dl>

<dl>
<dd>

**sku:** `*string` — The stock keeping unit (SKU) of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` — The title of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL for a product variant.
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*string` — The visibility of a product variant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStoreProductVariant(StoreID, ProductID, VariantID) -> *mailchimpmarketinggosdk.ECommerceProductVariant</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific product variant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStoreProductVariantEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
    VariantID: "variant_id",
}
client.Ecommerce.GetStoreProductVariant(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**variantID:** `string` — The id for the product variant.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpsertStoreProductVariant(StoreID, ProductID, VariantID, request) -> *mailchimpmarketinggosdk.ECommerceProductVariant</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add or update a product variant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpsertStoreProductVariantEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
    VariantID: "variant_id",
}
client.Ecommerce.UpsertStoreProductVariant(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**variantID:** `string` — The id for the product variant.
    
</dd>
</dl>

<dl>
<dd>

**backorders:** `*string` — The backorders of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` — A unique identifier for the product variant.
    
</dd>
</dl>

<dl>
<dd>

**imageURL:** `*string` — The image URL for a product variant.
    
</dd>
</dl>

<dl>
<dd>

**inventoryQuantity:** `*int` — The inventory quantity of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**price:** `*mailchimpmarketinggosdk.UpsertStoreProductVariantEcommerceRequestPrice` 
    
</dd>
</dl>

<dl>
<dd>

**sku:** `*string` — The stock keeping unit (SKU) of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL for a product variant.
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*string` — The visibility of a product variant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStoreProductVariant(StoreID, ProductID, VariantID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a product variant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStoreProductVariantEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
    VariantID: "variant_id",
}
client.Ecommerce.DeleteStoreProductVariant(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**variantID:** `string` — The id for the product variant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStoreProductVariant(StoreID, ProductID, VariantID, request) -> *mailchimpmarketinggosdk.ECommerceProductVariant</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a product variant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStoreProductVariantEcommerceRequest{
    StoreID: "store_id",
    ProductID: "product_id",
    VariantID: "variant_id",
}
client.Ecommerce.UpdateStoreProductVariant(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The id for the product of a store.
    
</dd>
</dl>

<dl>
<dd>

**variantID:** `string` — The id for the product variant.
    
</dd>
</dl>

<dl>
<dd>

**backorders:** `*string` — The backorders of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**imageURL:** `*string` — The image URL for a product variant.
    
</dd>
</dl>

<dl>
<dd>

**inventoryQuantity:** `*int` — The inventory quantity of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**price:** `*mailchimpmarketinggosdk.UpdateStoreProductVariantEcommerceRequestPrice` 
    
</dd>
</dl>

<dl>
<dd>

**sku:** `*string` — The stock keeping unit (SKU) of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title of a product variant.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL for a product variant.
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*string` — The visibility of a product variant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStorePromoRules(StoreID) -> *mailchimpmarketinggosdk.ListStorePromoRulesEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a store's promo rules.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStorePromoRulesEcommerceRequest{
    StoreID: "store_id",
}
client.Ecommerce.ListStorePromoRules(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStorePromoRule(StoreID, request) -> *mailchimpmarketinggosdk.ECommercePromoRule</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new promo rule to a store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStorePromoRuleEcommerceRequest{
    StoreID: "store_id",
    Amount: &mailchimpmarketinggosdk.CreateStorePromoRuleEcommerceRequestAmount{
        Double: 1.1,
    },
    Description: "Save BIG during our summer sale!",
    ID: "id",
    Target: mailchimpmarketinggosdk.CreateStorePromoRuleEcommerceRequestTargetPerItem,
    Type: mailchimpmarketinggosdk.CreateStorePromoRuleEcommerceRequestTypeFixed,
}
client.Ecommerce.CreateStorePromoRule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*mailchimpmarketinggosdk.CreateStorePromoRuleEcommerceRequestAmount` 
    
</dd>
</dl>

<dl>
<dd>

**createdAtForeign:** `*string` — The date and time the promotion was created in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**description:** `string` — The description of a promotion restricted to UTF-8 characters with max length 255.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the promo rule is currently enabled.
    
</dd>
</dl>

<dl>
<dd>

**endsAt:** `*mailchimpmarketinggosdk.CreateStorePromoRuleEcommerceRequestEndsAt` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — A unique identifier for the promo rule. If Ecommerce platform does not support promo rule, use promo code id as promo rule id. Restricted to UTF-8 characters with max length 50.
    
</dd>
</dl>

<dl>
<dd>

**startsAt:** `*mailchimpmarketinggosdk.CreateStorePromoRuleEcommerceRequestStartsAt` 
    
</dd>
</dl>

<dl>
<dd>

**target:** `*mailchimpmarketinggosdk.CreateStorePromoRuleEcommerceRequestTarget` — The target that the discount applies to.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title that will show up in promotion campaign. Restricted to UTF-8 characters with max length of 100 bytes.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*mailchimpmarketinggosdk.CreateStorePromoRuleEcommerceRequestType` — Type of discount. For free shipping set type to fixed.
    
</dd>
</dl>

<dl>
<dd>

**updatedAtForeign:** `*string` — The date and time the promotion was updated in ISO 8601 format.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStorePromoRule(StoreID, PromoRuleID) -> *mailchimpmarketinggosdk.ECommercePromoRule</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific promo rule.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStorePromoRuleEcommerceRequest{
    StoreID: "store_id",
    PromoRuleID: "promo_rule_id",
}
client.Ecommerce.GetStorePromoRule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**promoRuleID:** `string` — The id for the promo rule of a store.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStorePromoRule(StoreID, PromoRuleID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a promo rule from a store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStorePromoRuleEcommerceRequest{
    StoreID: "store_id",
    PromoRuleID: "promo_rule_id",
}
client.Ecommerce.DeleteStorePromoRule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**promoRuleID:** `string` — The id for the promo rule of a store.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStorePromoRule(StoreID, PromoRuleID, request) -> *mailchimpmarketinggosdk.ECommercePromoRule</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a promo rule.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStorePromoRuleEcommerceRequest{
    StoreID: "store_id",
    PromoRuleID: "promo_rule_id",
}
client.Ecommerce.UpdateStorePromoRule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**promoRuleID:** `string` — The id for the promo rule of a store.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*mailchimpmarketinggosdk.UpdateStorePromoRuleEcommerceRequestAmount` 
    
</dd>
</dl>

<dl>
<dd>

**createdAtForeign:** `*string` — The date and time the promotion was created in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — The description of a promotion restricted to UTF-8 characters with max length 255.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the promo rule is currently enabled.
    
</dd>
</dl>

<dl>
<dd>

**endsAt:** `*mailchimpmarketinggosdk.UpdateStorePromoRuleEcommerceRequestEndsAt` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` — A unique identifier for the promo rule. If Ecommerce platform does not support promo rule, use promo code id as promo rule id. Restricted to UTF-8 characters with max length 50.
    
</dd>
</dl>

<dl>
<dd>

**startsAt:** `*mailchimpmarketinggosdk.UpdateStorePromoRuleEcommerceRequestStartsAt` 
    
</dd>
</dl>

<dl>
<dd>

**target:** `*mailchimpmarketinggosdk.UpdateStorePromoRuleEcommerceRequestTarget` — The target that the discount applies to.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title that will show up in promotion campaign. Restricted to UTF-8 characters with max length of 100 bytes.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*mailchimpmarketinggosdk.UpdateStorePromoRuleEcommerceRequestType` — Type of discount. For free shipping set type to fixed.
    
</dd>
</dl>

<dl>
<dd>

**updatedAtForeign:** `*string` — The date and time the promotion was updated in ISO 8601 format.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.ListStorePromoRulePromoCodes(StoreID, PromoRuleID) -> *mailchimpmarketinggosdk.ListStorePromoRulePromoCodesEcommerceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a store's promo codes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListStorePromoRulePromoCodesEcommerceRequest{
    StoreID: "store_id",
    PromoRuleID: "promo_rule_id",
}
client.Ecommerce.ListStorePromoRulePromoCodes(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**promoRuleID:** `string` — The id for the promo rule of a store.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.CreateStorePromoRulePromoCode(StoreID, PromoRuleID, request) -> *mailchimpmarketinggosdk.ECommercePromoCode</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new promo code to a store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateStorePromoRulePromoCodeEcommerceRequest{
    StoreID: "store_id",
    PromoRuleID: "promo_rule_id",
    Code: "summersale",
    ID: "id",
    RedemptionURL: "A url that applies promo code directly at checkout or a url that points to sale page or store url",
}
client.Ecommerce.CreateStorePromoRulePromoCode(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**promoRuleID:** `string` — The id for the promo rule of a store.
    
</dd>
</dl>

<dl>
<dd>

**code:** `string` — The discount code. Restricted to UTF-8 characters with max length 50.
    
</dd>
</dl>

<dl>
<dd>

**createdAtForeign:** `*string` — The date and time the promotion was created in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the promo code is currently enabled.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — A unique identifier for the promo code. Restricted to UTF-8 characters with max length 50.
    
</dd>
</dl>

<dl>
<dd>

**redemptionURL:** `string` — The url that should be used in the promotion campaign restricted to UTF-8 characters with max length 2000.
    
</dd>
</dl>

<dl>
<dd>

**updatedAtForeign:** `*string` — The date and time the promotion was updated in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**usageCount:** `*int` — Number of times promo code has been used.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.GetStorePromoRulePromoCode(StoreID, PromoRuleID, PromoCodeID) -> *mailchimpmarketinggosdk.ECommercePromoCode</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific promo code.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetStorePromoRulePromoCodeEcommerceRequest{
    StoreID: "store_id",
    PromoRuleID: "promo_rule_id",
    PromoCodeID: "promo_code_id",
}
client.Ecommerce.GetStorePromoRulePromoCode(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**promoRuleID:** `string` — The id for the promo rule of a store.
    
</dd>
</dl>

<dl>
<dd>

**promoCodeID:** `string` — The id for the promo code of a store.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.DeleteStorePromoRulePromoCode(StoreID, PromoRuleID, PromoCodeID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a promo code from a store.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteStorePromoRulePromoCodeEcommerceRequest{
    StoreID: "store_id",
    PromoRuleID: "promo_rule_id",
    PromoCodeID: "promo_code_id",
}
client.Ecommerce.DeleteStorePromoRulePromoCode(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**promoRuleID:** `string` — The id for the promo rule of a store.
    
</dd>
</dl>

<dl>
<dd>

**promoCodeID:** `string` — The id for the promo code of a store.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ecommerce.UpdateStorePromoRulePromoCode(StoreID, PromoRuleID, PromoCodeID, request) -> *mailchimpmarketinggosdk.ECommercePromoCode</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a promo code.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateStorePromoRulePromoCodeEcommerceRequest{
    StoreID: "store_id",
    PromoRuleID: "promo_rule_id",
    PromoCodeID: "promo_code_id",
}
client.Ecommerce.UpdateStorePromoRulePromoCode(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**storeID:** `string` — The store id.
    
</dd>
</dl>

<dl>
<dd>

**promoRuleID:** `string` — The id for the promo rule of a store.
    
</dd>
</dl>

<dl>
<dd>

**promoCodeID:** `string` — The id for the promo code of a store.
    
</dd>
</dl>

<dl>
<dd>

**code:** `*string` — The discount code. Restricted to UTF-8 characters with max length 50.
    
</dd>
</dl>

<dl>
<dd>

**createdAtForeign:** `*string` — The date and time the promotion was created in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the promo code is currently enabled.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` — A unique identifier for the promo code. Restricted to UTF-8 characters with max length 50.
    
</dd>
</dl>

<dl>
<dd>

**redemptionURL:** `*string` — The url that should be used in the promotion campaign restricted to UTF-8 characters with max length 2000.
    
</dd>
</dl>

<dl>
<dd>

**updatedAtForeign:** `*string` — The date and time the promotion was updated in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**usageCount:** `*int` — Number of times promo code has been used.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FacebookAds
<details><summary><code>client.FacebookAds.List() -> *mailchimpmarketinggosdk.ListFacebookAdsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of Facebook ads.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListFacebookAdsRequest{}
client.FacebookAds.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListFacebookAdsRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListFacebookAdsRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FacebookAds.Get(OutreachID) -> *mailchimpmarketinggosdk.FacebookAds</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details of a Facebook ad.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetFacebookAdsRequest{
    OutreachID: "outreach_id",
}
client.FacebookAds.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**outreachID:** `string` — The outreach id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileManager
<details><summary><code>client.FileManager.List() -> []*mailchimpmarketinggosdk.ListFileManagerResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about the file-manager endpoint's resources
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileManager.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.ListFiles() -> *mailchimpmarketinggosdk.ListFilesFileManagerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of available images and files stored in the File Manager for the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListFilesFileManagerRequest{}
client.FileManager.ListFiles(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — The file type for the File Manager file.
    
</dd>
</dl>

<dl>
<dd>

**createdBy:** `*string` — The Mailchimp account user who created the File Manager file.
    
</dd>
</dl>

<dl>
<dd>

**beforeCreatedAt:** `*string` — Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceCreatedAt:** `*string` — Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListFilesFileManagerRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListFilesFileManagerRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.CreateFile(request) -> *mailchimpmarketinggosdk.GalleryFile</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upload a new image or file to the File Manager.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateFileFileManagerRequest{
    FileData: "file_data",
    Name: "name",
}
client.FileManager.CreateFile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fileData:** `string` — The base64-encoded contents of the file.
    
</dd>
</dl>

<dl>
<dd>

**folderID:** `*int` — The id of the folder.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the file.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.GetFile(FileID) -> *mailchimpmarketinggosdk.GalleryFile</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific file in the File Manager.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetFileFileManagerRequest{
    FileID: "file_id",
}
client.FileManager.GetFile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fileID:** `string` — The unique id for the File Manager file.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.DeleteFile(FileID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a specific file from the File Manager.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteFileFileManagerRequest{
    FileID: "file_id",
}
client.FileManager.DeleteFile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fileID:** `string` — The unique id for the File Manager file.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.UpdateFile(FileID, request) -> *mailchimpmarketinggosdk.GalleryFile</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a file in the File Manager.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateFileFileManagerRequest{
    FileID: "file_id",
}
client.FileManager.UpdateFile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fileID:** `string` — The unique id for the File Manager file.
    
</dd>
</dl>

<dl>
<dd>

**folderID:** `*int` — The id of the folder. Setting `folder_id` to `0` will remove a file from its current folder.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the file.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.ListFolders() -> *mailchimpmarketinggosdk.ListFoldersFileManagerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of all folders in the File Manager.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListFoldersFileManagerRequest{}
client.FileManager.ListFolders(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**createdBy:** `*string` — The Mailchimp account user who created the File Manager file.
    
</dd>
</dl>

<dl>
<dd>

**beforeCreatedAt:** `*string` — Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceCreatedAt:** `*string` — Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.CreateFolder(request) -> *mailchimpmarketinggosdk.CreateFolderFileManagerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new folder in the File Manager.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateFolderFileManagerRequest{
    Name: "name",
}
client.FileManager.CreateFolder(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the folder.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.GetFolder(FolderID) -> *mailchimpmarketinggosdk.GetFolderFileManagerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific folder in the File Manager.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetFolderFileManagerRequest{
    FolderID: "folder_id",
}
client.FileManager.GetFolder(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `string` — The unique id for the File Manager folder.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.DeleteFolder(FolderID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific folder in the File Manager.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteFolderFileManagerRequest{
    FolderID: "folder_id",
}
client.FileManager.DeleteFolder(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `string` — The unique id for the File Manager folder.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.UpdateFolder(FolderID, request) -> *mailchimpmarketinggosdk.UpdateFolderFileManagerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific File Manager folder.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateFolderFileManagerRequest{
    FolderID: "folder_id",
    Name: "name",
}
client.FileManager.UpdateFolder(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `string` — The unique id for the File Manager folder.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the folder.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileManager.ListFolderFiles(FolderID) -> *mailchimpmarketinggosdk.ListFolderFilesFileManagerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of available images and files stored in this folder.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListFolderFilesFileManagerRequest{
    FolderID: "folder_id",
}
client.FileManager.ListFolderFiles(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `string` — The unique id for the File Manager folder.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — The file type for the File Manager file.
    
</dd>
</dl>

<dl>
<dd>

**createdBy:** `*string` — The Mailchimp account user who created the File Manager file.
    
</dd>
</dl>

<dl>
<dd>

**beforeCreatedAt:** `*string` — Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceCreatedAt:** `*string` — Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListFolderFilesFileManagerRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListFolderFilesFileManagerRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## LandingPages
<details><summary><code>client.LandingPages.List() -> *mailchimpmarketinggosdk.ListLandingPagesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all landing pages.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListLandingPagesRequest{}
client.LandingPages.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListLandingPagesRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListLandingPagesRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LandingPages.Create(request) -> *mailchimpmarketinggosdk.LandingPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create an unpublished and contentless Mailchimp landing page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateLandingPagesRequest{}
client.LandingPages.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**useDefaultList:** `*bool` — Will create the Landing Page using the account's Default List instead of requiring a list_id.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — The description of this landing page.
    
</dd>
</dl>

<dl>
<dd>

**listID:** `*string` — The list's ID associated with this landing page.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of this landing page.
    
</dd>
</dl>

<dl>
<dd>

**storeID:** `*string` — The ID of the store associated with this landing page.
    
</dd>
</dl>

<dl>
<dd>

**templateID:** `*int` — The template_id of this landing page.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title of this landing page seen in the browser's title bar.
    
</dd>
</dl>

<dl>
<dd>

**tracking:** `*mailchimpmarketinggosdk.CreateLandingPagesRequestTracking` — The tracking settings applied to this landing page.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*mailchimpmarketinggosdk.CreateLandingPagesRequestType` — The type of template the landing page has.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LandingPages.Get(PageID) -> *mailchimpmarketinggosdk.LandingPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetLandingPagesRequest{
    PageID: "page_id",
}
client.LandingPages.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pageID:** `string` — The unique id for the page.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LandingPages.Delete(PageID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a landing page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteLandingPagesRequest{
    PageID: "page_id",
}
client.LandingPages.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pageID:** `string` — The unique id for the page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LandingPages.Update(PageID, request) -> *mailchimpmarketinggosdk.LandingPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a landing page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateLandingPagesRequest{
    PageID: "page_id",
}
client.LandingPages.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pageID:** `string` — The unique id for the page.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — The description of this landing page.
    
</dd>
</dl>

<dl>
<dd>

**listID:** `*string` — The list's ID associated with this landing page.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of this landing page.
    
</dd>
</dl>

<dl>
<dd>

**storeID:** `*string` — The ID of the store associated with this landing page.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title of this landing page seen in the browser's title bar.
    
</dd>
</dl>

<dl>
<dd>

**tracking:** `*mailchimpmarketinggosdk.UpdateLandingPagesRequestTracking` — The tracking settings applied to this landing page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LandingPages.CreateActionPublish(PageID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Publish a landing page that is in draft, unpublished, or has been previously published and edited.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionPublishLandingPagesRequest{
    PageID: "page_id",
}
client.LandingPages.CreateActionPublish(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pageID:** `string` — The unique id for the page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LandingPages.CreateActionUnpublish(PageID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unpublish a landing page that is in draft or has been published.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionUnpublishLandingPagesRequest{
    PageID: "page_id",
}
client.LandingPages.CreateActionUnpublish(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pageID:** `string` — The unique id for the page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LandingPages.ListContent(PageID) -> *mailchimpmarketinggosdk.ListContentLandingPagesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the the HTML for your landing page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListContentLandingPagesRequest{
    PageID: "page_id",
}
client.LandingPages.ListContent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pageID:** `string` — The unique id for the page.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## lists
<details><summary><code>client.Lists.List() -> *mailchimpmarketinggosdk.ListListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about all lists in the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListListsRequest{}
client.Lists.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**beforeDateCreated:** `*string` — Restrict response to lists created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceDateCreated:** `*string` — Restrict results to lists created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**beforeCampaignLastSent:** `*string` — Restrict results to lists created before the last campaign send date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceCampaignLastSent:** `*string` — Restrict results to lists created after the last campaign send date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — Restrict results to lists that include a specific subscriber's email address.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListListsRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListListsRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>

<dl>
<dd>

**hasEcommerceStore:** `*bool` — Restrict results to lists that contain an active, connected, undeleted ecommerce store.
    
</dd>
</dl>

<dl>
<dd>

**includeTotalContacts:** `*bool` — Deprecated. Return the total_contacts field in the stats response, which contains an approximate count of subscribed, unsubscribed, and transactional contacts. For a complete audience contact count, use the /audiences endpoint instead.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.Create(request) -> *mailchimpmarketinggosdk.SubscriberList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new list in your Mailchimp account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateListsRequest{
    CampaignDefaults: &mailchimpmarketinggosdk.CreateListsRequestCampaignDefaults{
        FromEmail: "from_email",
        FromName: "from_name",
        Language: "language",
        Subject: "subject",
    },
    Contact: &mailchimpmarketinggosdk.CreateListsRequestContact{
        Address1: "address1",
        City: "city",
        Company: "company",
        Country: "country",
    },
    EmailTypeOption: true,
    Name: "name",
    PermissionReminder: "permission_reminder",
}
client.Lists.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignDefaults:** `*mailchimpmarketinggosdk.CreateListsRequestCampaignDefaults` — [Default values for campaigns](https://mailchimp.com/help/edit-your-emails-subject-preview-text-from-name-or-from-email-address/) created for this list.
    
</dd>
</dl>

<dl>
<dd>

**contact:** `*mailchimpmarketinggosdk.CreateListsRequestContact` — [Contact information displayed in campaign footers](https://mailchimp.com/help/about-campaign-footers/) to comply with international spam laws.
    
</dd>
</dl>

<dl>
<dd>

**doubleOptin:** `*bool` — Whether or not to require the subscriber to confirm subscription via email.
    
</dd>
</dl>

<dl>
<dd>

**emailTypeOption:** `bool` — Whether the list supports [multiple formats for emails](https://mailchimp.com/help/audience-settings-and-defaults/). When set to `true`, subscribers can choose whether they want to receive HTML or plain-text emails. When set to `false`, subscribers will receive HTML emails, with a plain-text alternative backup.
    
</dd>
</dl>

<dl>
<dd>

**marketingPermissions:** `*bool` — Whether or not the list has marketing permissions (eg. GDPR) enabled.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the list.
    
</dd>
</dl>

<dl>
<dd>

**notifyOnSubscribe:** `*string` — The email address to send [subscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to.
    
</dd>
</dl>

<dl>
<dd>

**notifyOnUnsubscribe:** `*string` — The email address to send [unsubscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to.
    
</dd>
</dl>

<dl>
<dd>

**permissionReminder:** `string` — The [permission reminder](https://mailchimp.com/help/edit-the-permission-reminder/) for the list.
    
</dd>
</dl>

<dl>
<dd>

**useArchiveBar:** `*bool` — Whether campaigns for this list use the [Archive Bar](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) in archives by default.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.Get(ListID) -> *mailchimpmarketinggosdk.SubscriberList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific list in your Mailchimp account. Results include list members who have signed up but haven't confirmed their subscription yet and unsubscribed or cleaned.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetListsRequest{
    ListID: "list_id",
}
client.Lists.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**includeTotalContacts:** `*bool` — Deprecated. Return the total_contacts field in the stats response, which contains an approximate count of subscribed, unsubscribed, and transactional contacts. For a complete audience contact count, use the /audiences endpoint instead.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.BatchSubscribeOrUnsubscribe(ListID, request) -> *mailchimpmarketinggosdk.BatchSubscribeOrUnsubscribeListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Batch subscribe or unsubscribe list members.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.BatchSubscribeOrUnsubscribeListsRequest{
    ListID: "list_id",
    Members: []*mailchimpmarketinggosdk.BatchSubscribeOrUnsubscribeListsRequestMembersItem{},
}
client.Lists.BatchSubscribeOrUnsubscribe(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**skipMergeValidation:** `*bool` — If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false.
    
</dd>
</dl>

<dl>
<dd>

**skipDuplicateCheck:** `*bool` — If skip_duplicate_check is true, we will ignore duplicates sent in the request when using the batch sub/unsub on the lists endpoint. The status of the first appearance in the request will be saved. This defaults to false.
    
</dd>
</dl>

<dl>
<dd>

**members:** `[]*mailchimpmarketinggosdk.BatchSubscribeOrUnsubscribeListsRequestMembersItem` — An array of objects, each representing an email address and the subscription status for a specific list. Up to 500 members may be added or updated with each API call.
    
</dd>
</dl>

<dl>
<dd>

**syncTags:** `*bool` — Whether this batch operation will replace all existing tags with tags in request.
    
</dd>
</dl>

<dl>
<dd>

**updateExisting:** `*bool` — Whether this batch operation will change existing members' subscription status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.Delete(ListID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a list from your Mailchimp account. If you delete a list, you'll lose the list history—including subscriber activity, unsubscribes, complaints, and bounces. You’ll also lose subscribers’ email addresses, unless you exported and backed up your list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteListsRequest{
    ListID: "list_id",
}
client.Lists.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.Update(ListID, request) -> *mailchimpmarketinggosdk.SubscriberList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the settings for a specific list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateListsRequest{
    ListID: "list_id",
}
client.Lists.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**campaignDefaults:** `*mailchimpmarketinggosdk.UpdateListsRequestCampaignDefaults` — [Default values for campaigns](https://mailchimp.com/help/edit-your-emails-subject-preview-text-from-name-or-from-email-address/) created for this list.
    
</dd>
</dl>

<dl>
<dd>

**contact:** `*mailchimpmarketinggosdk.UpdateListsRequestContact` — [Contact information displayed in campaign footers](https://mailchimp.com/help/about-campaign-footers/) to comply with international spam laws.
    
</dd>
</dl>

<dl>
<dd>

**doubleOptin:** `*bool` — Whether or not to require the subscriber to confirm subscription via email.
    
</dd>
</dl>

<dl>
<dd>

**emailTypeOption:** `*bool` — Whether the list supports [multiple formats for emails](https://mailchimp.com/help/audience-settings-and-defaults/). When set to `true`, subscribers can choose whether they want to receive HTML or plain-text emails. When set to `false`, subscribers will receive HTML emails, with a plain-text alternative backup.
    
</dd>
</dl>

<dl>
<dd>

**marketingPermissions:** `*bool` — Whether or not the list has marketing permissions (eg. GDPR) enabled.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the list.
    
</dd>
</dl>

<dl>
<dd>

**notifyOnSubscribe:** `*string` — The email address to send [subscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to.
    
</dd>
</dl>

<dl>
<dd>

**notifyOnUnsubscribe:** `*string` — The email address to send [unsubscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to.
    
</dd>
</dl>

<dl>
<dd>

**permissionReminder:** `*string` — The [permission reminder](https://mailchimp.com/help/edit-the-permission-reminder/) for the list.
    
</dd>
</dl>

<dl>
<dd>

**useArchiveBar:** `*bool` — Whether campaigns for this list use the [Archive Bar](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) in archives by default.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListAbuseReports(ListID) -> *mailchimpmarketinggosdk.ListAbuseReportsListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all abuse reports for a specific list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListAbuseReportsListsRequest{
    ListID: "list_id",
}
client.Lists.ListAbuseReports(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetAbuseReport(ListID, ReportID) -> *mailchimpmarketinggosdk.ListsAbuseReports</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details about a specific abuse report.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetAbuseReportListsRequest{
    ListID: "list_id",
    ReportID: "report_id",
}
client.Lists.GetAbuseReport(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**reportID:** `string` — The id for the abuse report.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListActivity(ListID) -> *mailchimpmarketinggosdk.ListActivityListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get up to the previous 180 days of daily detailed aggregated activity stats for a list, not including Automation activity.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListActivityListsRequest{
    ListID: "list_id",
}
client.Lists.ListActivity(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListClients(ListID) -> *mailchimpmarketinggosdk.ListClientsListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of the top email clients based on user-agent strings.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListClientsListsRequest{
    ListID: "list_id",
}
client.Lists.ListClients(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListGrowthHistory(ListID) -> *mailchimpmarketinggosdk.ListGrowthHistoryListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a month-by-month summary of a specific list's growth activity.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListGrowthHistoryListsRequest{
    ListID: "list_id",
}
client.Lists.ListGrowthHistory(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListGrowthHistoryListsRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListGrowthHistoryListsRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetGrowthHistory(ListID, Month) -> *mailchimpmarketinggosdk.GrowthHistory</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a summary of a specific list's growth activity for a specific month and year.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetGrowthHistoryListsRequest{
    ListID: "list_id",
    Month: "month",
}
client.Lists.GetGrowthHistory(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**month:** `string` — A specific month of list growth history.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListInterestCategories(ListID) -> *mailchimpmarketinggosdk.ListInterestCategoriesListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a list's interest categories.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListInterestCategoriesListsRequest{
    ListID: "list_id",
}
client.Lists.ListInterestCategories(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Restrict results a type of interest group
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListInterestCategoriesListsRequestSortField` — Returns interest categories sorted by the specified field. Defaults to display_order.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListInterestCategoriesListsRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateInterestCategory(ListID, request) -> *mailchimpmarketinggosdk.InterestCategory</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new interest category.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateInterestCategoryListsRequest{
    ListID: "list_id",
    Title: "title",
    Type: mailchimpmarketinggosdk.CreateInterestCategoryListsRequestTypeCheckboxes,
}
client.Lists.CreateInterestCategory(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**displayOrder:** `*int` — The order that the categories are displayed in the list. Lower numbers display first.
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` — The text description of this category. This field appears on signup forms and is often phrased as a question.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*mailchimpmarketinggosdk.CreateInterestCategoryListsRequestType` — Determines how this category’s interests appear on signup forms.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetInterestCategory(ListID, InterestCategoryID) -> *mailchimpmarketinggosdk.InterestCategory</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific interest category.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetInterestCategoryListsRequest{
    ListID: "list_id",
    InterestCategoryID: "interest_category_id",
}
client.Lists.GetInterestCategory(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**interestCategoryID:** `string` — The unique ID for the interest category.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.DeleteInterestCategory(ListID, InterestCategoryID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific interest category.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteInterestCategoryListsRequest{
    ListID: "list_id",
    InterestCategoryID: "interest_category_id",
}
client.Lists.DeleteInterestCategory(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**interestCategoryID:** `string` — The unique ID for the interest category.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.UpdateInterestCategory(ListID, InterestCategoryID, request) -> *mailchimpmarketinggosdk.InterestCategory</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific interest category.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateInterestCategoryListsRequest{
    ListID: "list_id",
    InterestCategoryID: "interest_category_id",
}
client.Lists.UpdateInterestCategory(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**interestCategoryID:** `string` — The unique ID for the interest category.
    
</dd>
</dl>

<dl>
<dd>

**displayOrder:** `*int` — The order that the categories are displayed in the list. Lower numbers display first.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The text description of this category. This field appears on signup forms and is often phrased as a question.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*mailchimpmarketinggosdk.UpdateInterestCategoryListsRequestType` — Determines how this category’s interests appear on signup forms.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListInterestCategoryInterests(ListID, InterestCategoryID) -> *mailchimpmarketinggosdk.ListInterestCategoryInterestsListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of this category's interests.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListInterestCategoryInterestsListsRequest{
    ListID: "list_id",
    InterestCategoryID: "interest_category_id",
}
client.Lists.ListInterestCategoryInterests(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**interestCategoryID:** `string` — The unique ID for the interest category.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateInterestCategoryInterest(ListID, InterestCategoryID, request) -> *mailchimpmarketinggosdk.Interest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new interest or 'group name' for a specific category.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateInterestCategoryInterestListsRequest{
    ListID: "list_id",
    InterestCategoryID: "interest_category_id",
    Name: "name",
}
client.Lists.CreateInterestCategoryInterest(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**interestCategoryID:** `string` — The unique ID for the interest category.
    
</dd>
</dl>

<dl>
<dd>

**displayOrder:** `*int` — The display order for interests.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the interest. This can be shown publicly on a subscription form.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetInterestCategoryInterest(ListID, InterestCategoryID, InterestID) -> *mailchimpmarketinggosdk.Interest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get interests or 'group names' for a specific category.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetInterestCategoryInterestListsRequest{
    ListID: "list_id",
    InterestCategoryID: "interest_category_id",
    InterestID: "interest_id",
}
client.Lists.GetInterestCategoryInterest(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**interestCategoryID:** `string` — The unique ID for the interest category.
    
</dd>
</dl>

<dl>
<dd>

**interestID:** `string` — The specific interest or 'group name'.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.DeleteInterestCategoryInterest(ListID, InterestCategoryID, InterestID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete interests or group names in a specific category.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteInterestCategoryInterestListsRequest{
    ListID: "list_id",
    InterestCategoryID: "interest_category_id",
    InterestID: "interest_id",
}
client.Lists.DeleteInterestCategoryInterest(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**interestCategoryID:** `string` — The unique ID for the interest category.
    
</dd>
</dl>

<dl>
<dd>

**interestID:** `string` — The specific interest or 'group name'.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.UpdateInterestCategoryInterest(ListID, InterestCategoryID, InterestID, request) -> *mailchimpmarketinggosdk.Interest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update interests or 'group names' for a specific category.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateInterestCategoryInterestListsRequest{
    ListID: "list_id",
    InterestCategoryID: "interest_category_id",
    InterestID: "interest_id",
}
client.Lists.UpdateInterestCategoryInterest(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**interestCategoryID:** `string` — The unique ID for the interest category.
    
</dd>
</dl>

<dl>
<dd>

**interestID:** `string` — The specific interest or 'group name'.
    
</dd>
</dl>

<dl>
<dd>

**displayOrder:** `*int` — The display order for interests.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the interest. This can be shown publicly on a subscription form.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListLocations(ListID) -> *mailchimpmarketinggosdk.ListLocationsListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the locations (countries) that the list's subscribers have been tagged to based on geocoding their IP address.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListLocationsListsRequest{
    ListID: "list_id",
}
client.Lists.ListLocations(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListMembers(ListID) -> *mailchimpmarketinggosdk.ListMembersListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about members in a specific Mailchimp list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListMembersListsRequest{
    ListID: "list_id",
}
client.Lists.ListMembers(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**emailType:** `*string` — The email type.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*mailchimpmarketinggosdk.ListMembersListsRequestStatus` — The subscriber's status.
    
</dd>
</dl>

<dl>
<dd>

**sinceTimestampOpt:** `*string` — Restrict results to subscribers who opted-in after the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**beforeTimestampOpt:** `*string` — Restrict results to subscribers who opted-in before the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceLastChanged:** `*string` — Restrict results to subscribers whose information changed after the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**beforeLastChanged:** `*string` — Restrict results to subscribers whose information changed before the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**uniqueEmailID:** `*string` — A unique identifier for the email address across all Mailchimp lists.
    
</dd>
</dl>

<dl>
<dd>

**vipOnly:** `*bool` — A filter to return only the list's VIP members. Passing `true` will restrict results to VIP list members, passing `false` will return all list members.
    
</dd>
</dl>

<dl>
<dd>

**interestCategoryID:** `*string` — The unique id for the interest category.
    
</dd>
</dl>

<dl>
<dd>

**interestIDs:** `*string` — Used to filter list members by interests. Must be accompanied by interest_category_id and interest_match. The value must be a comma separated list of interest ids present for any supplied interest categories.
    
</dd>
</dl>

<dl>
<dd>

**interestMatch:** `*mailchimpmarketinggosdk.ListMembersListsRequestInterestMatch` — Used to filter list members by interests. Must be accompanied by interest_category_id and interest_ids. "any" will match a member with any of the interest supplied, "all" will only match members with every interest supplied, and "none" will match members without any of the interest supplied.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListMembersListsRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListMembersListsRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>

<dl>
<dd>

**sinceLastCampaign:** `*bool` — Filter subscribers by those subscribed/unsubscribed/pending/cleaned since last email campaign send. Member status is required to use this filter.
    
</dd>
</dl>

<dl>
<dd>

**unsubscribedSince:** `*string` — Filter subscribers by those unsubscribed since a specific date. Using any status other than unsubscribed with this filter will result in an error.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateMember(ListID, request) -> *mailchimpmarketinggosdk.ListMembers</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new member to the list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateMemberListsRequest{
    ListID: "list_id",
    EmailAddress: "email_address",
    Status: mailchimpmarketinggosdk.CreateMemberListsRequestStatusSubscribed,
}
client.Lists.CreateMember(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**skipMergeValidation:** `*bool` — If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `string` — Email address for a subscriber.
    
</dd>
</dl>

<dl>
<dd>

**emailType:** `*string` — Type of email this member asked to get ('html' or 'text').
    
</dd>
</dl>

<dl>
<dd>

**interests:** `map[string]bool` — The key of this object's properties is the ID of the interest in question.
    
</dd>
</dl>

<dl>
<dd>

**ipOpt:** `*string` — The IP address the subscriber used to confirm their opt-in status.
    
</dd>
</dl>

<dl>
<dd>

**ipSignup:** `*string` — IP address the subscriber signed up from.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*string` — If set/detected, the [subscriber's language](https://mailchimp.com/help/view-and-edit-contact-languages/).
    
</dd>
</dl>

<dl>
<dd>

**location:** `*mailchimpmarketinggosdk.CreateMemberListsRequestLocation` — Subscriber location information.
    
</dd>
</dl>

<dl>
<dd>

**marketingPermissions:** `[]*mailchimpmarketinggosdk.CreateMemberListsRequestMarketingPermissionsItem` — The marketing permissions for the subscriber.
    
</dd>
</dl>

<dl>
<dd>

**mergeFields:** `map[string]*mailchimpmarketinggosdk.CreateMemberListsRequestMergeFieldsValue` — A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*mailchimpmarketinggosdk.CreateMemberListsRequestStatus` — Subscriber's current status.
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` — The tags that are associated with a member.
    
</dd>
</dl>

<dl>
<dd>

**timestampOpt:** `*mailchimpmarketinggosdk.CreateMemberListsRequestTimestampOpt` 
    
</dd>
</dl>

<dl>
<dd>

**timestampSignup:** `*mailchimpmarketinggosdk.CreateMemberListsRequestTimestampSignup` 
    
</dd>
</dl>

<dl>
<dd>

**vip:** `*bool` — [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetMember(ListID, SubscriberHash) -> *mailchimpmarketinggosdk.ListMembers</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific list member, including a currently subscribed, unsubscribed, or bounced member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetMemberListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.GetMember(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.UpsertMember(ListID, SubscriberHash, request) -> *mailchimpmarketinggosdk.ListMembers</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add or update a list member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpsertMemberListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
    EmailAddress: "email_address",
}
client.Lists.UpsertMember(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**skipMergeValidation:** `*bool` — If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `string` — Email address for a subscriber. This value is required only if the email address is not already present on the list.
    
</dd>
</dl>

<dl>
<dd>

**emailType:** `*string` — Type of email this member asked to get ('html' or 'text').
    
</dd>
</dl>

<dl>
<dd>

**interests:** `map[string]bool` — The key of this object's properties is the ID of the interest in question.
    
</dd>
</dl>

<dl>
<dd>

**ipOpt:** `*string` — The IP address the subscriber used to confirm their opt-in status.
    
</dd>
</dl>

<dl>
<dd>

**ipSignup:** `*string` — IP address the subscriber signed up from.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*string` — If set/detected, the [subscriber's language](https://mailchimp.com/help/view-and-edit-contact-languages/).
    
</dd>
</dl>

<dl>
<dd>

**location:** `*mailchimpmarketinggosdk.UpsertMemberListsRequestLocation` — Subscriber location information.
    
</dd>
</dl>

<dl>
<dd>

**marketingPermissions:** `[]*mailchimpmarketinggosdk.UpsertMemberListsRequestMarketingPermissionsItem` — The marketing permissions for the subscriber.
    
</dd>
</dl>

<dl>
<dd>

**mergeFields:** `map[string]*mailchimpmarketinggosdk.UpsertMemberListsRequestMergeFieldsValue` — A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*mailchimpmarketinggosdk.UpsertMemberListsRequestStatus` — Subscriber's current status.
    
</dd>
</dl>

<dl>
<dd>

**statusIfNew:** `*mailchimpmarketinggosdk.UpsertMemberListsRequestStatusIfNew` — Subscriber's status. This value is required only if the email address is not already present on the list.
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` — The tags that are associated with a member.
    
</dd>
</dl>

<dl>
<dd>

**timestampOpt:** `*mailchimpmarketinggosdk.UpsertMemberListsRequestTimestampOpt` 
    
</dd>
</dl>

<dl>
<dd>

**timestampSignup:** `*mailchimpmarketinggosdk.UpsertMemberListsRequestTimestampSignup` 
    
</dd>
</dl>

<dl>
<dd>

**vip:** `*bool` — [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.DeleteMember(ListID, SubscriberHash) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Archive a list member. To permanently delete, use the delete-permanent action.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteMemberListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.DeleteMember(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.UpdateMember(ListID, SubscriberHash, request) -> *mailchimpmarketinggosdk.ListMembers</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update information for a specific list member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateMemberListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.UpdateMember(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**skipMergeValidation:** `*bool` — If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` — Email address for a subscriber.
    
</dd>
</dl>

<dl>
<dd>

**emailType:** `*string` — Type of email this member asked to get ('html' or 'text').
    
</dd>
</dl>

<dl>
<dd>

**interests:** `map[string]bool` — The key of this object's properties is the ID of the interest in question.
    
</dd>
</dl>

<dl>
<dd>

**ipOpt:** `*string` — The IP address the subscriber used to confirm their opt-in status.
    
</dd>
</dl>

<dl>
<dd>

**ipSignup:** `*string` — IP address the subscriber signed up from.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*string` — If set/detected, the [subscriber's language](https://mailchimp.com/help/view-and-edit-contact-languages/).
    
</dd>
</dl>

<dl>
<dd>

**location:** `*mailchimpmarketinggosdk.UpdateMemberListsRequestLocation` — Subscriber location information.
    
</dd>
</dl>

<dl>
<dd>

**marketingPermissions:** `[]*mailchimpmarketinggosdk.UpdateMemberListsRequestMarketingPermissionsItem` — The marketing permissions for the subscriber.
    
</dd>
</dl>

<dl>
<dd>

**mergeFields:** `map[string]*mailchimpmarketinggosdk.UpdateMemberListsRequestMergeFieldsValue` — A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*mailchimpmarketinggosdk.UpdateMemberListsRequestStatus` — Subscriber's current status.
    
</dd>
</dl>

<dl>
<dd>

**timestampOpt:** `*mailchimpmarketinggosdk.UpdateMemberListsRequestTimestampOpt` 
    
</dd>
</dl>

<dl>
<dd>

**timestampSignup:** `*mailchimpmarketinggosdk.UpdateMemberListsRequestTimestampSignup` 
    
</dd>
</dl>

<dl>
<dd>

**vip:** `*bool` — [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateMemberActionDeletePermanent(ListID, SubscriberHash) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete all personally identifiable information related to a list member, and remove them from a list. This will make it impossible to re-import the list member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateMemberActionDeletePermanentListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.CreateMemberActionDeletePermanent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListMemberActivity(ListID, SubscriberHash) -> *mailchimpmarketinggosdk.ListMemberActivityListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last 50 events of a member's activity on a specific list, including opens, clicks, and unsubscribes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListMemberActivityListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.ListMemberActivity(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**action:** `*mailchimpmarketinggosdk.ListMemberActivityListsRequestActionItem` — A comma seperated list of actions to return.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListMemberActivityFeed(ListID, SubscriberHash) -> *mailchimpmarketinggosdk.ListMemberActivityFeedListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a member's activity on a specific list, including opens, clicks, and unsubscribes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListMemberActivityFeedListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.ListMemberActivityFeed(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**activityFilters:** `*mailchimpmarketinggosdk.ListMemberActivityFeedListsRequestActivityFiltersItem` — A comma-separated list of activity filters that correspond to a set of activity types, e.g "?activity_filters=open,bounce,click".
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListMemberEvents(ListID, SubscriberHash) -> *mailchimpmarketinggosdk.ListMemberEventsListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get events for a contact.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListMemberEventsListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.ListMemberEvents(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateMemberEvent(ListID, SubscriberHash, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add an event for a list member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateMemberEventListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
    Name: "name",
}
client.Lists.CreateMemberEvent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**isSyncing:** `*bool` — Events created with the is_syncing value set to `true` will not trigger automations.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name for this type of event ('purchased', 'visited', etc). Must be 2-30 characters in length
    
</dd>
</dl>

<dl>
<dd>

**occurredAt:** `*time.Time` — The date and time the event occurred in ISO 8601 format.
    
</dd>
</dl>

<dl>
<dd>

**properties:** `map[string]string` — An optional list of properties
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListMemberGoals(ListID, SubscriberHash) -> *mailchimpmarketinggosdk.ListMemberGoalsListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last 50 Goal events for a member on a specific list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListMemberGoalsListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.ListMemberGoals(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListMemberNotes(ListID, SubscriberHash) -> *mailchimpmarketinggosdk.ListMemberNotesListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get recent notes for a specific list member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListMemberNotesListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.ListMemberNotes(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListMemberNotesListsRequestSortField` — Returns notes sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListMemberNotesListsRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateMemberNote(ListID, SubscriberHash, request) -> *mailchimpmarketinggosdk.MemberNotes</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new note for a specific subscriber.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateMemberNoteListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.CreateMemberNote(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` — The content of the note. Note length is limited to 1,000 characters.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetMemberNote(ListID, SubscriberHash, NoteID) -> *mailchimpmarketinggosdk.MemberNotes</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific note for a specific list member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetMemberNoteListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
    NoteID: "note_id",
}
client.Lists.GetMemberNote(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**noteID:** `string` — The id for the note.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.DeleteMemberNote(ListID, SubscriberHash, NoteID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific note for a specific list member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteMemberNoteListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
    NoteID: "note_id",
}
client.Lists.DeleteMemberNote(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**noteID:** `string` — The id for the note.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.UpdateMemberNote(ListID, SubscriberHash, NoteID, request) -> *mailchimpmarketinggosdk.MemberNotes</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific note for a specific list member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateMemberNoteListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
    NoteID: "note_id",
}
client.Lists.UpdateMemberNote(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**noteID:** `string` — The id for the note.
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` — The content of the note. Note length is limited to 1,000 characters.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListMemberTags(ListID, SubscriberHash) -> *mailchimpmarketinggosdk.ListMemberTagsListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the tags on a list member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListMemberTagsListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.ListMemberTags(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateMemberTag(ListID, SubscriberHash, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add or remove tags from a list member. If a tag that does not exist is passed in and set as 'active', a new tag will be created.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateMemberTagListsRequest{
    ListID: "list_id",
    SubscriberHash: "subscriber_hash",
    Tags: []*mailchimpmarketinggosdk.CreateMemberTagListsRequestTagsItem{
        &mailchimpmarketinggosdk.CreateMemberTagListsRequestTagsItem{
            Name: "name",
            Status: mailchimpmarketinggosdk.CreateMemberTagListsRequestTagsItemStatusInactive,
        },
    },
}
client.Lists.CreateMemberTag(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>

<dl>
<dd>

**isSyncing:** `*bool` — When is_syncing is true, automations based on the tags in the request will not fire
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]*mailchimpmarketinggosdk.CreateMemberTagListsRequestTagsItem` — A list of tags assigned to the list member.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListMergeFields(ListID) -> *mailchimpmarketinggosdk.ListMergeFieldsListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of all merge fields for an audience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListMergeFieldsListsRequest{
    ListID: "list_id",
}
client.Lists.ListMergeFields(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — The merge field type.
    
</dd>
</dl>

<dl>
<dd>

**required:** `*bool` — Whether it's a required merge field.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateMergeField(ListID, request) -> *mailchimpmarketinggosdk.MergeField</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new merge field for a specific audience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateMergeFieldListsRequest{
    ListID: "list_id",
    Name: "name",
    Type: mailchimpmarketinggosdk.CreateMergeFieldListsRequestTypeText,
}
client.Lists.CreateMergeField(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**defaultValue:** `*string` — The default value for the merge field if `null`.
    
</dd>
</dl>

<dl>
<dd>

**displayOrder:** `*int` — The order that the merge field displays on the list signup form.
    
</dd>
</dl>

<dl>
<dd>

**helpText:** `*string` — Extra text to help the subscriber fill out the form.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the merge field (audience field).
    
</dd>
</dl>

<dl>
<dd>

**options:** `*mailchimpmarketinggosdk.CreateMergeFieldListsRequestOptions` — Extra options for some merge field types.
    
</dd>
</dl>

<dl>
<dd>

**public:** `*bool` — Whether the merge field is displayed on the signup form.
    
</dd>
</dl>

<dl>
<dd>

**required:** `*bool` — Whether the merge field is required to import a contact.
    
</dd>
</dl>

<dl>
<dd>

**tag:** `*string` — The merge tag used for Mailchimp campaigns and [adding contact information](https://mailchimp.com/developer/marketing/docs/merge-fields/#add-merge-data-to-contacts).
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*mailchimpmarketinggosdk.CreateMergeFieldListsRequestType` — The [type](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for the merge field.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetMergeField(ListID, MergeID) -> *mailchimpmarketinggosdk.MergeField</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific merge field.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetMergeFieldListsRequest{
    ListID: "list_id",
    MergeID: "merge_id",
}
client.Lists.GetMergeField(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**mergeID:** `string` — The id for the merge field.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.DeleteMergeField(ListID, MergeID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific merge field.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteMergeFieldListsRequest{
    ListID: "list_id",
    MergeID: "merge_id",
}
client.Lists.DeleteMergeField(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**mergeID:** `string` — The id for the merge field.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.UpdateMergeField(ListID, MergeID, request) -> *mailchimpmarketinggosdk.MergeField</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific merge field.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateMergeFieldListsRequest{
    ListID: "list_id",
    MergeID: "merge_id",
}
client.Lists.UpdateMergeField(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**mergeID:** `string` — The id for the merge field.
    
</dd>
</dl>

<dl>
<dd>

**defaultValue:** `*string` — The default value for the merge field if `null`.
    
</dd>
</dl>

<dl>
<dd>

**displayOrder:** `*int` — The order that the merge field displays on the list signup form.
    
</dd>
</dl>

<dl>
<dd>

**helpText:** `*string` — Extra text to help the subscriber fill out the form.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the merge field (audience field).
    
</dd>
</dl>

<dl>
<dd>

**options:** `*mailchimpmarketinggosdk.UpdateMergeFieldListsRequestOptions` — Extra options for some merge field types.
    
</dd>
</dl>

<dl>
<dd>

**public:** `*bool` — Whether the merge field is displayed on the signup form.
    
</dd>
</dl>

<dl>
<dd>

**required:** `*bool` — Whether the merge field is required to import a contact.
    
</dd>
</dl>

<dl>
<dd>

**tag:** `*string` — The merge tag used for Mailchimp campaigns and [adding contact information](https://mailchimp.com/developer/marketing/docs/merge-fields/#add-merge-data-to-contacts).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListSegments(ListID) -> *mailchimpmarketinggosdk.ListSegmentsListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about all available segments for a specific list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSegmentsListsRequest{
    ListID: "list_id",
}
client.Lists.ListSegments(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Limit results based on segment type.
    
</dd>
</dl>

<dl>
<dd>

**sinceCreatedAt:** `*string` — Restrict results to segments created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**beforeCreatedAt:** `*string` — Restrict results to segments created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**includeCleaned:** `*bool` — Include cleaned members in response
    
</dd>
</dl>

<dl>
<dd>

**includeTransactional:** `*bool` — Include transactional members in response
    
</dd>
</dl>

<dl>
<dd>

**includeUnsubscribed:** `*bool` — Include unsubscribed members in response
    
</dd>
</dl>

<dl>
<dd>

**sinceUpdatedAt:** `*string` — Restrict results to segments update after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**beforeUpdatedAt:** `*string` — Restrict results to segments update before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**excludeType:** `*mailchimpmarketinggosdk.ListSegmentsListsRequestExcludeType` — Exclude results based on segment type. For example, use `exclude_type=static` to exclude tags from the response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateSegment(ListID, request) -> *mailchimpmarketinggosdk.List</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new segment in a specific list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateSegmentListsRequest{
    ListID: "list_id",
    Name: "name",
}
client.Lists.CreateSegment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the segment.
    
</dd>
</dl>

<dl>
<dd>

**options:** `*mailchimpmarketinggosdk.CreateSegmentListsRequestOptions` — The [conditions of the segment](https://mailchimp.com/help/save-and-manage-segments/). Static and fuzzy segments don't have conditions.
    
</dd>
</dl>

<dl>
<dd>

**staticSegment:** `[]string` — An array of emails to be used for a static segment. Any emails provided that are not present on the list will be ignored. Passing an empty array will create a static segment without any subscribers. This field cannot be provided with the options field.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetSegment(ListID, SegmentID) -> *mailchimpmarketinggosdk.List</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific segment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetSegmentListsRequest{
    ListID: "list_id",
    SegmentID: "segment_id",
}
client.Lists.GetSegment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**segmentID:** `string` — The unique id for the segment.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**includeCleaned:** `*bool` — Include cleaned members in response
    
</dd>
</dl>

<dl>
<dd>

**includeTransactional:** `*bool` — Include transactional members in response
    
</dd>
</dl>

<dl>
<dd>

**includeUnsubscribed:** `*bool` — Include unsubscribed members in response
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.BatchAddOrRemoveMembers(ListID, SegmentID, request) -> *mailchimpmarketinggosdk.BatchAddOrRemoveMembersListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Batch add/remove list members to static segment
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.BatchAddOrRemoveMembersListsRequest{
    ListID: "list_id",
    SegmentID: "segment_id",
}
client.Lists.BatchAddOrRemoveMembers(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**segmentID:** `string` — The unique id for the segment.
    
</dd>
</dl>

<dl>
<dd>

**membersToAdd:** `[]string` — An array of emails to be used for a static segment. Any emails provided that are not present on the list will be ignored. A maximum of 500 members can be sent.
    
</dd>
</dl>

<dl>
<dd>

**membersToRemove:** `[]string` — An array of emails to be used for a static segment. Any emails provided that are not present on the list will be ignored. A maximum of 500 members can be sent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.DeleteSegment(ListID, SegmentID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific segment in a list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteSegmentListsRequest{
    ListID: "list_id",
    SegmentID: "segment_id",
}
client.Lists.DeleteSegment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**segmentID:** `string` — The unique id for the segment.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.UpdateSegment(ListID, SegmentID, request) -> *mailchimpmarketinggosdk.List</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific segment in a list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateSegmentListsRequest{
    ListID: "list_id",
    SegmentID: "segment_id",
}
client.Lists.UpdateSegment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**segmentID:** `string` — The unique id for the segment.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the segment.
    
</dd>
</dl>

<dl>
<dd>

**options:** `*mailchimpmarketinggosdk.UpdateSegmentListsRequestOptions` — The [conditions of the segment](https://mailchimp.com/help/save-and-manage-segments/). Static and fuzzy segments don't have conditions.
    
</dd>
</dl>

<dl>
<dd>

**staticSegment:** `[]string` — An array of emails to be used for a static segment. Any emails provided that are not present on the list will be ignored. Passing an empty array for an existing static segment will reset that segment and remove all members. This field cannot be provided with the `options` field.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListSegmentMembers(ListID, SegmentID) -> *mailchimpmarketinggosdk.ListSegmentMembersListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about members in a saved segment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSegmentMembersListsRequest{
    ListID: "list_id",
    SegmentID: "segment_id",
}
client.Lists.ListSegmentMembers(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**segmentID:** `string` — The unique id for the segment.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**includeCleaned:** `*bool` — Include cleaned members in response
    
</dd>
</dl>

<dl>
<dd>

**includeTransactional:** `*bool` — Include transactional members in response
    
</dd>
</dl>

<dl>
<dd>

**includeUnsubscribed:** `*bool` — Include unsubscribed members in response
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateSegmentMember(ListID, SegmentID, request) -> *mailchimpmarketinggosdk.ListsSegmentsMembers</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a member to a static segment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateSegmentMemberListsRequest{
    ListID: "list_id",
    SegmentID: "segment_id",
    EmailAddress: "email_address",
}
client.Lists.CreateSegmentMember(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**segmentID:** `string` — The unique id for the segment.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `string` — Email address for a subscriber.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.DeleteSegmentMember(ListID, SegmentID, SubscriberHash) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a member from the specified static segment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteSegmentMemberListsRequest{
    ListID: "list_id",
    SegmentID: "segment_id",
    SubscriberHash: "subscriber_hash",
}
client.Lists.DeleteSegmentMember(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**segmentID:** `string` — The unique id for the segment.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListSignupForms(ListID) -> *mailchimpmarketinggosdk.ListSignupFormsListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get signup forms for a specific list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSignupFormsListsRequest{
    ListID: "list_id",
}
client.Lists.ListSignupForms(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateSignupForm(ListID, request) -> *mailchimpmarketinggosdk.SignupForm</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Customize a list's default signup form.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateSignupFormListsRequest{
    ListID: "list_id",
}
client.Lists.CreateSignupForm(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**contents:** `[]*mailchimpmarketinggosdk.CreateSignupFormListsRequestContentsItem` — The signup form body content.
    
</dd>
</dl>

<dl>
<dd>

**header:** `*mailchimpmarketinggosdk.CreateSignupFormListsRequestHeader` — Options for customizing your signup form header.
    
</dd>
</dl>

<dl>
<dd>

**styles:** `[]*mailchimpmarketinggosdk.CreateSignupFormListsRequestStylesItem` — An array of objects, each representing an element style for the signup form.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListSurveys(ListID) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about all available surveys for a specific list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSurveysListsRequest{
    ListID: "list_id",
}
client.Lists.ListSurveys(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateSurvey(ListID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a draft survey for an audience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateSurveyListsRequest{
    ListID: "list_id",
}
client.Lists.CreateSurvey(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title of the survey.
    
</dd>
</dl>

<dl>
<dd>

**sections:** `[]*mailchimpmarketinggosdk.SurveySectionRequest` — Initial survey sections.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetSurvey(ListID, SurveyID) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details about a specific survey.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetSurveyListsRequest{
    ListID: "list_id",
    SurveyID: "survey_id",
}
client.Lists.GetSurvey(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.DeleteSurvey(ListID, SurveyID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a survey.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteSurveyListsRequest{
    ListID: "list_id",
    SurveyID: "survey_id",
}
client.Lists.DeleteSurvey(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.UpdateSurvey(ListID, SurveyID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a survey. When sections is provided, send the complete section list in display order. Any existing section not included is deleted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateSurveyListsRequest{
    ListID: "list_id",
    SurveyID: "survey_id",
}
client.Lists.UpdateSurvey(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title of the survey.
    
</dd>
</dl>

<dl>
<dd>

**isPipedToInbox:** `*bool` — Whether responses are sent to Mailchimp Inbox.
    
</dd>
</dl>

<dl>
<dd>

**sections:** `[]*mailchimpmarketinggosdk.SurveySectionRequest` — The complete survey section list in display order. On update, sections omitted from this array are deleted. Include section id to update an existing section; omit section id to add a new section.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateListSurveyActionReplicate(ListIDPathParam, SurveyID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replicate a survey.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateListSurveyActionReplicateListsRequest{
    ListIDPathParam: "list_id",
    SurveyID: "survey_id",
}
client.Lists.CreateListSurveyActionReplicate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listIDPathParam:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title for the replicated survey.
    
</dd>
</dl>

<dl>
<dd>

**listID:** `*string` — The unique ID of the audience for the replicated survey. Defaults to the source survey audience.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListTagSearch(ListID) -> *mailchimpmarketinggosdk.ListTagSearchListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search for tags on a list by name. If no name is provided, will return all tags on the list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListTagSearchListsRequest{
    ListID: "list_id",
}
client.Lists.ListTagSearch(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The search query used to filter tags.  The search query will be compared to each tag as a prefix, so all tags that have a name starting with this field will be returned.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.ListWebhooks(ListID) -> *mailchimpmarketinggosdk.ListWebhooksListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about all webhooks for a specific list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListWebhooksListsRequest{
    ListID: "list_id",
}
client.Lists.ListWebhooks(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.CreateWebhook(ListID, request) -> *mailchimpmarketinggosdk.CreateWebhookListsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new webhook for a specific list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateWebhookListsRequest{
    ListID: "list_id",
    Body: &mailchimpmarketinggosdk.AddWebhook{},
}
client.Lists.CreateWebhook(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*mailchimpmarketinggosdk.AddWebhook` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.GetWebhook(ListID, WebhookID) -> *mailchimpmarketinggosdk.ListWebhooks</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific webhook.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetWebhookListsRequest{
    ListID: "list_id",
    WebhookID: "webhook_id",
}
client.Lists.GetWebhook(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**webhookID:** `string` — The webhook's id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.DeleteWebhook(ListID, WebhookID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific webhook in a list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteWebhookListsRequest{
    ListID: "list_id",
    WebhookID: "webhook_id",
}
client.Lists.DeleteWebhook(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**webhookID:** `string` — The webhook's id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Lists.UpdateWebhook(ListID, WebhookID, request) -> *mailchimpmarketinggosdk.ListWebhooks</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the settings for an existing webhook.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateWebhookListsRequest{
    ListID: "list_id",
    WebhookID: "webhook_id",
    Body: &mailchimpmarketinggosdk.AddWebhook{},
}
client.Lists.UpdateWebhook(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**webhookID:** `string` — The webhook's id.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*mailchimpmarketinggosdk.AddWebhook` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## surveys
<details><summary><code>client.Surveys.CreateListSurveyActionCreateEmail(ListID, SurveyID) -> *mailchimpmarketinggosdk.Campaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Utilize the List ID and Survey ID to generate a Campaign that links to your survey.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateListSurveyActionCreateEmailSurveysRequest{
    ListID: "list_id",
    SurveyID: "survey_id",
}
client.Surveys.CreateListSurveyActionCreateEmail(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Surveys.CreateListSurveyActionPublish(ListID, SurveyID) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Publish a survey that is in draft, unpublished, or has been previously published and edited.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateListSurveyActionPublishSurveysRequest{
    ListID: "list_id",
    SurveyID: "survey_id",
}
client.Surveys.CreateListSurveyActionPublish(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Surveys.CreateListSurveyActionUnpublish(ListID, SurveyID) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unpublish a survey that has been published.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateListSurveyActionUnpublishSurveysRequest{
    ListID: "list_id",
    SurveyID: "survey_id",
}
client.Surveys.CreateListSurveyActionUnpublish(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listID:** `string` — The unique ID for the list.
    
</dd>
</dl>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ping
<details><summary><code>client.Ping.List() -> *mailchimpmarketinggosdk.ListPingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

A health check for the API that won't return any account-specific information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ping.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## reporting
<details><summary><code>client.Reporting.List() -> []*mailchimpmarketinggosdk.ListReportingResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about the reporting endpoint's resources.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Reporting.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.ListFacebookAds() -> *mailchimpmarketinggosdk.ListFacebookAdsReportingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get reports of Facebook ads.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListFacebookAdsReportingRequest{}
client.Reporting.ListFacebookAds(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListFacebookAdsReportingRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListFacebookAdsReportingRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.GetFacebookAd(OutreachID) -> *mailchimpmarketinggosdk.ReportingFacebookAd</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get report of a Facebook ad.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetFacebookAdReportingRequest{
    OutreachID: "outreach_id",
}
client.Reporting.GetFacebookAd(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**outreachID:** `string` — The outreach id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.ListFacebookAdEcommerceProductActivity(OutreachID) -> *mailchimpmarketinggosdk.ListFacebookAdEcommerceProductActivityReportingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get breakdown of product activity for an outreach.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListFacebookAdEcommerceProductActivityReportingRequest{
    OutreachID: "outreach_id",
}
client.Reporting.ListFacebookAdEcommerceProductActivity(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**outreachID:** `string` — The outreach id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListFacebookAdEcommerceProductActivityReportingRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.ListLandingPages() -> *mailchimpmarketinggosdk.ListLandingPagesReportingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get reports of landing pages.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListLandingPagesReportingRequest{}
client.Reporting.ListLandingPages(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.GetLandingPage(OutreachID) -> *mailchimpmarketinggosdk.LandingPageReport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get report of a landing page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetLandingPageReportingRequest{
    OutreachID: "outreach_id",
}
client.Reporting.GetLandingPage(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**outreachID:** `string` — The outreach id.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.ListSurveys() -> *mailchimpmarketinggosdk.ListSurveysReportingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get reports for surveys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSurveysReportingRequest{}
client.Reporting.ListSurveys(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.GetSurvey(SurveyID) -> *mailchimpmarketinggosdk.GetSurveyReportingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get report for a survey.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetSurveyReportingRequest{
    SurveyID: "survey_id",
}
client.Reporting.GetSurvey(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.ListSurveyQuestions(SurveyID) -> *mailchimpmarketinggosdk.ListSurveyQuestionsReportingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get reports for survey questions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSurveyQuestionsReportingRequest{
    SurveyID: "survey_id",
}
client.Reporting.ListSurveyQuestions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.GetSurveyQuestion(SurveyID, QuestionID) -> *mailchimpmarketinggosdk.SurveyQuestionReport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get report for a survey question.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetSurveyQuestionReportingRequest{
    SurveyID: "survey_id",
    QuestionID: "question_id",
}
client.Reporting.GetSurveyQuestion(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>

<dl>
<dd>

**questionID:** `string` — The ID of the survey question.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.ListSurveyQuestionAnswers(SurveyID, QuestionID) -> *mailchimpmarketinggosdk.ListSurveyQuestionAnswersReportingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get answers for a survey question.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSurveyQuestionAnswersReportingRequest{
    SurveyID: "survey_id",
    QuestionID: "question_id",
}
client.Reporting.ListSurveyQuestionAnswers(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>

<dl>
<dd>

**questionID:** `string` — The ID of the survey question.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**respondentFamiliarityIs:** `*mailchimpmarketinggosdk.ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs` — Filter survey responses by familiarity of the respondents.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.ListSurveyResponses(SurveyID) -> *mailchimpmarketinggosdk.ListSurveyResponsesReportingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get responses to a survey.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSurveyResponsesReportingRequest{
    SurveyID: "survey_id",
}
client.Reporting.ListSurveyResponses(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**answeredQuestion:** `*int` — The ID of the question that was answered.
    
</dd>
</dl>

<dl>
<dd>

**choseAnswer:** `*string` — The ID of the option chosen to filter responses on.
    
</dd>
</dl>

<dl>
<dd>

**respondentFamiliarityIs:** `*mailchimpmarketinggosdk.ListSurveyResponsesReportingRequestRespondentFamiliarityIs` — Filter survey responses by familiarity of the respondents.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.GetSurveyRespons(SurveyID, ResponseID) -> *mailchimpmarketinggosdk.GetSurveyResponsReportingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a single survey response.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetSurveyResponsReportingRequest{
    SurveyID: "survey_id",
    ResponseID: "response_id",
}
client.Reporting.GetSurveyRespons(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**surveyID:** `string` — The ID of the survey.
    
</dd>
</dl>

<dl>
<dd>

**responseID:** `string` — The ID of the survey response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## reports
<details><summary><code>client.Reports.List() -> *mailchimpmarketinggosdk.ListReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get campaign reports.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListReportsRequest{}
client.Reports.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*mailchimpmarketinggosdk.ListReportsRequestType` — The campaign type.
    
</dd>
</dl>

<dl>
<dd>

**beforeSendTime:** `*time.Time` — Restrict the response to campaigns sent before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sinceSendTime:** `*time.Time` — Restrict the response to campaigns sent after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.Get(CampaignID) -> *mailchimpmarketinggosdk.CampaignReport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get report details for a specific sent campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListAbuseReports(CampaignID) -> *mailchimpmarketinggosdk.ListAbuseReportsReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of abuse complaints for a specific campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListAbuseReportsReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListAbuseReports(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.GetAbuseReport(CampaignID, ReportID) -> *mailchimpmarketinggosdk.AbuseComplaint</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific abuse report for a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetAbuseReportReportsRequest{
    CampaignID: "campaign_id",
    ReportID: "report_id",
}
client.Reports.GetAbuseReport(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**reportID:** `string` — The id for the abuse report.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListAdvice(CampaignID) -> *mailchimpmarketinggosdk.ListAdviceReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get feedback based on a campaign's statistics. Advice feedback is based on campaign stats like opens, clicks, unsubscribes, bounces, and more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListAdviceReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListAdvice(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListClickDetails(CampaignID) -> *mailchimpmarketinggosdk.ListClickDetailsReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about clicks on specific links in your Mailchimp campaigns.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListClickDetailsReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListClickDetails(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListClickDetailsReportsRequestSortField` — Returns click reports sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListClickDetailsReportsRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>

<dl>
<dd>

**filterBots:** `*bool` — When true, exclude automated bot clicks so the returned click counts reflect human clicks only, matching the in-app Recipient Activity view. Filtering changes a link's counts, but never removes a link from the response. Defaults to false (all clicks).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.GetClickDetail(CampaignID, LinkID) -> *mailchimpmarketinggosdk.ClickDetailReport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get click details for a specific link in a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetClickDetailReportsRequest{
    CampaignID: "campaign_id",
    LinkID: "link_id",
}
client.Reports.GetClickDetail(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**linkID:** `string` — The id for the link.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**filterBots:** `*bool` — When true, exclude automated bot clicks so the returned click counts reflect human clicks only, matching the in-app Recipient Activity view. Filtering changes a link's counts, but never removes a link from the response. Defaults to false (all clicks).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListClickDetailMembers(CampaignID, LinkID) -> *mailchimpmarketinggosdk.ListClickDetailMembersReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about list members who clicked on a specific link in a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListClickDetailMembersReportsRequest{
    CampaignID: "campaign_id",
    LinkID: "link_id",
}
client.Reports.ListClickDetailMembers(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**linkID:** `string` — The id for the link.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.GetClickDetailMember(CampaignID, LinkID, SubscriberHash) -> *mailchimpmarketinggosdk.ClickDetailMember</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific subscriber who clicked a link in a specific campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetClickDetailMemberReportsRequest{
    CampaignID: "campaign_id",
    LinkID: "link_id",
    SubscriberHash: "subscriber_hash",
}
client.Reports.GetClickDetailMember(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**linkID:** `string` — The id for the link.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListDomainPerformance(CampaignID) -> *mailchimpmarketinggosdk.ListDomainPerformanceReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get statistics for the top-performing email domains in a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListDomainPerformanceReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListDomainPerformance(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListEcommerceProductActivity(CampaignID) -> *mailchimpmarketinggosdk.ListEcommerceProductActivityReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get breakdown of product activity for a campaign
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListEcommerceProductActivityReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListEcommerceProductActivity(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListEcommerceProductActivityReportsRequestSortField` — Returns files sorted by the specified field.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListEepurl(CampaignID) -> *mailchimpmarketinggosdk.ListEepurlReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a summary of social activity for the campaign, tracked by EepURL.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListEepurlReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListEepurl(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListEmailActivity(CampaignID) -> *mailchimpmarketinggosdk.ListEmailActivityReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of member's subscriber activity in a specific campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListEmailActivityReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListEmailActivity(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**since:** `*string` — Restrict results to email activity events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**filterBots:** `*bool` — When true, exclude automated bot and Apple Mail Privacy Protection (MPP) proxy activity so the returned activity reflects human-only opens and clicks, matching the in-app Recipient Activity view. Filtering removes events from a member's activity, but never removes the member from the response. Defaults to false (all activity).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.GetEmailActivity(CampaignID, SubscriberHash) -> *mailchimpmarketinggosdk.EmailActivity</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific list member's activity in a campaign including opens, clicks, and bounces.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetEmailActivityReportsRequest{
    CampaignID: "campaign_id",
    SubscriberHash: "subscriber_hash",
}
client.Reports.GetEmailActivity(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**since:** `*string` — Restrict results to email activity events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**filterBots:** `*bool` — When true, exclude automated bot and Apple Mail Privacy Protection (MPP) proxy activity so the returned activity reflects human-only opens and clicks, matching the in-app Recipient Activity view. Filtering removes events from a member's activity, but never removes the member from the response. Defaults to false (all activity).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListLocations(CampaignID) -> *mailchimpmarketinggosdk.ListLocationsReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get top open locations for a specific campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListLocationsReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListLocations(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListOpenDetails(CampaignID) -> *mailchimpmarketinggosdk.ListOpenDetailsReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get detailed information about any campaign emails that were opened by a list member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListOpenDetailsReportsRequest{
    CampaignID: "campaign_id",
    Since: mailchimpmarketinggosdk.String(
        "2016-04-12 12:00:00",
    ),
}
client.Reports.ListOpenDetails(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**since:** `*string` — Restrict results to campaign open events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListOpenDetailsReportsRequestSortField` — Returns open reports sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListOpenDetailsReportsRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>

<dl>
<dd>

**filterBots:** `*bool` — When true, exclude automated (proxy/bot) opens so the returned open counts reflect human opens only, matching the in-app Recipient Activity view. A member whose opens are all automated is excluded from the human-only view. Defaults to false (all opens).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.GetOpenDetail(CampaignID, SubscriberHash) -> *mailchimpmarketinggosdk.OpenActivity</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific subscriber who opened a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetOpenDetailReportsRequest{
    CampaignID: "campaign_id",
    SubscriberHash: "subscriber_hash",
}
client.Reports.GetOpenDetail(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**filterBots:** `*bool` — When true, exclude automated (proxy/bot) opens so the returned open counts reflect human opens only, matching the in-app Recipient Activity view. A member whose opens are all automated is excluded from the human-only view. Defaults to false (all opens).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListSentTo(CampaignID) -> *mailchimpmarketinggosdk.ListSentToReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about campaign recipients.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSentToReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListSentTo(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.GetSentTo(CampaignID, SubscriberHash) -> *mailchimpmarketinggosdk.SentTo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific campaign recipient.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetSentToReportsRequest{
    CampaignID: "campaign_id",
    SubscriberHash: "subscriber_hash",
}
client.Reports.GetSentTo(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListSubReports(CampaignID) -> *mailchimpmarketinggosdk.ListSubReportsReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of reports with child campaigns for a specific parent campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSubReportsReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListSubReports(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.ListUnsubscribed(CampaignID) -> *mailchimpmarketinggosdk.ListUnsubscribedReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about members who have unsubscribed from a specific campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListUnsubscribedReportsRequest{
    CampaignID: "campaign_id",
}
client.Reports.ListUnsubscribed(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reports.GetUnsubscribed(CampaignID, SubscriberHash) -> *mailchimpmarketinggosdk.Unsubscribes</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific list member who unsubscribed from a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetUnsubscribedReportsRequest{
    CampaignID: "campaign_id",
    SubscriberHash: "subscriber_hash",
}
client.Reports.GetUnsubscribed(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**campaignID:** `string` — The unique id for the campaign.
    
</dd>
</dl>

<dl>
<dd>

**subscriberHash:** `string` — The MD5 hash of the lowercase version of the list member's email address.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SearchCampaigns
<details><summary><code>client.SearchCampaigns.List() -> *mailchimpmarketinggosdk.ListSearchCampaignsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search all campaigns for the specified query terms.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSearchCampaignsRequest{
    Query: "query",
}
client.SearchCampaigns.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**query:** `string` — The search query used to filter results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SmsCampaigns
<details><summary><code>client.SmsCampaigns.List() -> *mailchimpmarketinggosdk.ListSmsCampaignsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all SMS campaigns in an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSmsCampaignsRequest{}
client.SmsCampaigns.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SmsCampaigns.Create(request) -> *mailchimpmarketinggosdk.SmsCampaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new SMS campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateSmsCampaignsRequest{
    Name: "name",
}
client.SmsCampaigns.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the campaign.
    
</dd>
</dl>

<dl>
<dd>

**listID:** `*int` — The numeric ID of the list to send the campaign to.
    
</dd>
</dl>

<dl>
<dd>

**folderID:** `*string` — The ID of the folder to place this campaign in.
    
</dd>
</dl>

<dl>
<dd>

**segments:** `[]int` — The segment IDs to target for this campaign.
    
</dd>
</dl>

<dl>
<dd>

**excludedSegments:** `[]int` — The segment IDs to exclude from this campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SmsCampaigns.Get(SmsCampaignID) -> *mailchimpmarketinggosdk.SmsCampaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the details for a single SMS campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetSmsCampaignsRequest{
    SmsCampaignID: "sms_campaign_id",
}
client.SmsCampaigns.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**smsCampaignID:** `string` — The unique id for the SMS campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SmsCampaigns.Delete(SmsCampaignID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a campaign from your Mailchimp account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteSmsCampaignsRequest{
    SmsCampaignID: "sms_campaign_id",
}
client.SmsCampaigns.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**smsCampaignID:** `string` — The unique id for the SMS campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SmsCampaigns.Update(SmsCampaignID, request) -> *mailchimpmarketinggosdk.SmsCampaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an SMS campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateSmsCampaignsRequest{
    SmsCampaignID: "sms_campaign_id",
}
client.SmsCampaigns.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**smsCampaignID:** `string` — The unique id for the SMS campaign.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the campaign.
    
</dd>
</dl>

<dl>
<dd>

**folderID:** `*string` — The ID of the folder to place this campaign in.
    
</dd>
</dl>

<dl>
<dd>

**segments:** `[]int` — The segment IDs to target for this campaign.
    
</dd>
</dl>

<dl>
<dd>

**excludedSegments:** `[]int` — The segment IDs to exclude from this campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SmsCampaigns.CreateActionCancelSend(SmsCampaignID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancel a scheduled or sending SMS campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionCancelSendSmsCampaignsRequest{
    SmsCampaignID: "sms_campaign_id",
}
client.SmsCampaigns.CreateActionCancelSend(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**smsCampaignID:** `string` — The unique id for the SMS campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SmsCampaigns.CreateActionSchedule(SmsCampaignID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Schedule an SMS campaign for delivery.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionScheduleSmsCampaignsRequest{
    SmsCampaignID: "sms_campaign_id",
    ScheduleTime: mailchimpmarketinggosdk.MustParseDateTime(
        "2024-01-15T09:30:00Z",
    ),
}
client.SmsCampaigns.CreateActionSchedule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**smsCampaignID:** `string` — The unique id for the SMS campaign.
    
</dd>
</dl>

<dl>
<dd>

**scheduleTime:** `time.Time` — The UTC date and time to schedule the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SmsCampaigns.CreateActionSend(SmsCampaignID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Send an SMS campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionSendSmsCampaignsRequest{
    SmsCampaignID: "sms_campaign_id",
}
client.SmsCampaigns.CreateActionSend(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**smsCampaignID:** `string` — The unique id for the SMS campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SmsCampaigns.GetContent(SmsCampaignID) -> *mailchimpmarketinggosdk.SmsCampaignContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the content for an SMS campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetContentSmsCampaignsRequest{
    SmsCampaignID: "sms_campaign_id",
}
client.SmsCampaigns.GetContent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**smsCampaignID:** `string` — The unique id for the SMS campaign.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SmsCampaigns.UpsertContent(SmsCampaignID, request) -> *mailchimpmarketinggosdk.SmsCampaignContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set the content for an SMS campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpsertContentSmsCampaignsRequest{
    SmsCampaignID: "sms_campaign_id",
    MessageBody: "message_body",
}
client.SmsCampaigns.UpsertContent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**smsCampaignID:** `string` — The unique id for the SMS campaign.
    
</dd>
</dl>

<dl>
<dd>

**messageBody:** `string` — The SMS message body.
    
</dd>
</dl>

<dl>
<dd>

**media:** `[]*mailchimpmarketinggosdk.UpsertContentSmsCampaignsRequestMediaItem` — Attached images or files. Limited to one item. Omitting this field or sending an empty array removes any existing media; to keep the current media while updating other fields, re-send the media array.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SearchMembers
<details><summary><code>client.SearchMembers.List() -> *mailchimpmarketinggosdk.ListSearchMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search for list members. This search can be restricted to a specific list, or can be used to search across all lists in an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListSearchMembersRequest{
    Query: "query",
}
client.SearchMembers.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**query:** `string` — The search query used to filter results. Query should be a valid email, or a string representing a contact's first or last name.
    
</dd>
</dl>

<dl>
<dd>

**listID:** `*string` — The unique id for the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## TemplateFolders
<details><summary><code>client.TemplateFolders.List() -> *mailchimpmarketinggosdk.ListTemplateFoldersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all folders used to organize templates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListTemplateFoldersRequest{}
client.TemplateFolders.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TemplateFolders.Create(request) -> *mailchimpmarketinggosdk.CreateTemplateFoldersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new template folder.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateTemplateFoldersRequest{
    Name: "name",
}
client.TemplateFolders.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the folder.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TemplateFolders.Get(FolderID) -> *mailchimpmarketinggosdk.GetTemplateFoldersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific folder used to organize templates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetTemplateFoldersRequest{
    FolderID: "folder_id",
}
client.TemplateFolders.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `string` — The unique id for the template folder.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TemplateFolders.Delete(FolderID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific template folder, and mark all the templates in the folder as 'unfiled'.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteTemplateFoldersRequest{
    FolderID: "folder_id",
}
client.TemplateFolders.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `string` — The unique id for the template folder.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TemplateFolders.Update(FolderID, request) -> *mailchimpmarketinggosdk.UpdateTemplateFoldersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific folder used to organize templates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateTemplateFoldersRequest{
    FolderID: "folder_id",
    Name: "name",
}
client.TemplateFolders.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `string` — The unique id for the template folder.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the folder.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## templates
<details><summary><code>client.Templates.List() -> *mailchimpmarketinggosdk.ListTemplatesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of an account's available templates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListTemplatesRequest{}
client.Templates.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — The number of records to return. Default value is 10. Maximum value is 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
    
</dd>
</dl>

<dl>
<dd>

**createdBy:** `*string` — The Mailchimp account user who created the template.
    
</dd>
</dl>

<dl>
<dd>

**sinceDateCreated:** `*string` — Restrict the response to templates created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**beforeDateCreated:** `*string` — Restrict the response to templates created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Limit results based on template type.
    
</dd>
</dl>

<dl>
<dd>

**category:** `*string` — Limit results based on category.
    
</dd>
</dl>

<dl>
<dd>

**folderID:** `*string` — The unique folder id.
    
</dd>
</dl>

<dl>
<dd>

**sortField:** `*mailchimpmarketinggosdk.ListTemplatesRequestSortField` — Returns user templates sorted by the specified field.
    
</dd>
</dl>

<dl>
<dd>

**contentType:** `*mailchimpmarketinggosdk.ListTemplatesRequestContentType` — Limit results based on how the template's content is put together. Only templates of type `user` can be filtered by `content_type`. If you want to retrieve saved templates created with the legacy email editor, then filter `content_type` to `template`. If you'd rather pull your saved templates for the new editor, filter to `multichannel`. For code your own templates, filter to `html`.
    
</dd>
</dl>

<dl>
<dd>

**sortDir:** `*mailchimpmarketinggosdk.ListTemplatesRequestSortDir` — Determines the order direction for sorted results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Templates.Create(request) -> *mailchimpmarketinggosdk.TemplateInstance</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new template for the account. Only Classic templates are supported.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateTemplatesRequest{
    HTML: "html",
    Name: "Freddie's Jokes",
}
client.Templates.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folderID:** `*string` — The id of the folder the template is currently in.
    
</dd>
</dl>

<dl>
<dd>

**html:** `string` — The raw HTML for the template. We  support the Mailchimp [Template Language](https://mailchimp.com/help/getting-started-with-mailchimps-template-language/) in any HTML code passed via the API.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the template.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Templates.Get(TemplateID) -> *mailchimpmarketinggosdk.TemplateInstance</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetTemplatesRequest{
    TemplateID: "template_id",
}
client.Templates.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateID:** `string` — The unique id for the template.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Templates.Delete(TemplateID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteTemplatesRequest{
    TemplateID: "template_id",
}
client.Templates.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateID:** `string` — The unique id for the template.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Templates.Update(TemplateID, request) -> *mailchimpmarketinggosdk.TemplateInstance</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the name, HTML, or `folder_id` of an existing template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.UpdateTemplatesRequest{
    TemplateID: "template_id",
}
client.Templates.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateID:** `string` — The unique id for the template.
    
</dd>
</dl>

<dl>
<dd>

**folderID:** `*string` — The id of the folder the template is currently in.
    
</dd>
</dl>

<dl>
<dd>

**html:** `*string` — The raw HTML for the template. We  support the Mailchimp [Template Language](https://mailchimp.com/help/getting-started-with-mailchimps-template-language/) in any HTML code passed via the API.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the template.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Templates.ListDefaultContent(TemplateID) -> *mailchimpmarketinggosdk.ListDefaultContentTemplatesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the sections that you can edit in a template, including each section's default content.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.ListDefaultContentTemplatesRequest{
    TemplateID: "template_id",
}
client.Templates.ListDefaultContent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateID:** `string` — The unique id for the template.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>

<dl>
<dd>

**excludeFields:** `*string` — A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## VerifiedDomains
<details><summary><code>client.VerifiedDomains.List() -> *mailchimpmarketinggosdk.ListVerifiedDomainsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all of the sending domains on the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.VerifiedDomains.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VerifiedDomains.Create(request) -> *mailchimpmarketinggosdk.CreateVerifiedDomainsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a domain to the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateVerifiedDomainsRequest{
    VerificationEmail: "verification_email",
}
client.VerifiedDomains.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**verificationEmail:** `string` — The e-mail address at the domain you want to verify. This will receive a two-factor challenge to be used in the verify action.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VerifiedDomains.Get(DomainName) -> *mailchimpmarketinggosdk.GetVerifiedDomainsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the details for a single domain on the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.GetVerifiedDomainsRequest{
    DomainName: "domain_name",
}
client.VerifiedDomains.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domainName:** `string` — The domain name.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VerifiedDomains.Delete(DomainName) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a verified domain from the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.DeleteVerifiedDomainsRequest{
    DomainName: "domain_name",
}
client.VerifiedDomains.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domainName:** `string` — The domain name.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VerifiedDomains.CreateActionVerify(DomainName, request) -> *mailchimpmarketinggosdk.CreateActionVerifyVerifiedDomainsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Verify a domain for sending.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &mailchimpmarketinggosdk.CreateActionVerifyVerifiedDomainsRequest{
    DomainName: "domain_name",
    Code: "code",
}
client.VerifiedDomains.CreateActionVerify(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domainName:** `string` — The domain name.
    
</dd>
</dl>

<dl>
<dd>

**code:** `string` — The code that was sent to the email address provided when adding a new domain to verify.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

