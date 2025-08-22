package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeleteFolderResponse represents the DeleteFolderResponse schema from the OpenAPI specification
type DeleteFolderResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// GetSharedLinksResponse represents the GetSharedLinksResponse schema from the OpenAPI specification
type GetSharedLinksResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []SharedLink `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// CreateFolderRequest represents the CreateFolderRequest schema from the OpenAPI specification
type CreateFolderRequest struct {
	Description string `json:"description,omitempty"` // Optional description of the folder.
	Drive_id string `json:"drive_id,omitempty"` // ID of the drive to create the folder in.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Name string `json:"name"` // The name of the folder.
	Parent_folder_id string `json:"parent_folder_id"` // The parent folder to create the new file within.
}

// GetSharedLinkResponse represents the GetSharedLinkResponse schema from the OpenAPI specification
type GetSharedLinkResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data SharedLink `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// FilesSort represents the FilesSort schema from the OpenAPI specification
type FilesSort struct {
	By string `json:"by,omitempty"` // The field on which to sort the Files
	Direction string `json:"direction,omitempty"` // The direction in which to sort the results
}

// LinkedFolder represents the LinkedFolder schema from the OpenAPI specification
type LinkedFolder struct {
	Name string `json:"name,omitempty"` // The name of the folder
	Id string `json:"id"` // A unique identifier for an object.
}

// FilesFilter represents the FilesFilter schema from the OpenAPI specification
type FilesFilter struct {
	Drive_id string `json:"drive_id,omitempty"` // ID of the drive to filter on
	Folder_id string `json:"folder_id,omitempty"` // ID of the folder to filter on. The root folder has an alias "root"
	Shared bool `json:"shared,omitempty"` // Only return files and folders that are shared
}

// NotImplementedResponse represents the NotImplementedResponse schema from the OpenAPI specification
type NotImplementedResponse struct {
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
}

// UpdateSharedLinkResponse represents the UpdateSharedLinkResponse schema from the OpenAPI specification
type UpdateSharedLinkResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// GetFileResponse represents the GetFileResponse schema from the OpenAPI specification
type GetFileResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedFile `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// UnauthorizedResponse represents the UnauthorizedResponse schema from the OpenAPI specification
type UnauthorizedResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// DeleteFileResponse represents the DeleteFileResponse schema from the OpenAPI specification
type DeleteFileResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// NotFoundResponse represents the NotFoundResponse schema from the OpenAPI specification
type NotFoundResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// GetFoldersResponse represents the GetFoldersResponse schema from the OpenAPI specification
type GetFoldersResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Folder `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// PaymentRequiredResponse represents the PaymentRequiredResponse schema from the OpenAPI specification
type PaymentRequiredResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// TooManyRequestsResponse represents the TooManyRequestsResponse schema from the OpenAPI specification
type TooManyRequestsResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail map[string]interface{} `json:"detail,omitempty"`
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 6585)
}

// CopyFolderRequest represents the CopyFolderRequest schema from the OpenAPI specification
type CopyFolderRequest struct {
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Name string `json:"name,omitempty"` // The name of the folder.
	Parent_folder_id string `json:"parent_folder_id"` // The parent folder to create the new file within.
}

// DeleteDriveResponse represents the DeleteDriveResponse schema from the OpenAPI specification
type DeleteDriveResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
	Current string `json:"current,omitempty"` // Link to navigate to the current page through the API
}

// UploadSession represents the UploadSession schema from the OpenAPI specification
type UploadSession struct {
	Part_size float64 `json:"part_size,omitempty"` // Size in bytes of each part of the file that you will upload. Uploaded parts need to be this size for the upload to be successful.
	Success bool `json:"success,omitempty"` // Indicates if the upload session was completed successfully.
	Uploaded_byte_range string `json:"uploaded_byte_range,omitempty"` // The range of bytes that was successfully uploaded.
	Expires_at string `json:"expires_at,omitempty"`
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Parallel_upload_supported bool `json:"parallel_upload_supported,omitempty"` // Indicates if parts of the file can uploaded in parallel.
}

// UpdateFileResponse represents the UpdateFileResponse schema from the OpenAPI specification
type UpdateFileResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// GetDriveGroupsResponse represents the GetDriveGroupsResponse schema from the OpenAPI specification
type GetDriveGroupsResponse struct {
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []DriveGroup `json:"data"`
}

// CreateDriveGroupResponse represents the CreateDriveGroupResponse schema from the OpenAPI specification
type CreateDriveGroupResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// UnexpectedErrorResponse represents the UnexpectedErrorResponse schema from the OpenAPI specification
type UnexpectedErrorResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// UpdateDriveResponse represents the UpdateDriveResponse schema from the OpenAPI specification
type UpdateDriveResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// SharedLinkTarget represents the SharedLinkTarget schema from the OpenAPI specification
type SharedLinkTarget struct {
	Id string `json:"id"` // A unique identifier for an object.
	Name string `json:"name,omitempty"` // The name of the file
	TypeField string `json:"type,omitempty"` // The type of resource. Could be file, folder or url
}

// Owner represents the Owner schema from the OpenAPI specification
type Owner struct {
	Name string `json:"name,omitempty"` // Name of the owner
	Email string `json:"email,omitempty"` // Email of the owner
	Id string `json:"id,omitempty"` // ID of the owner
}

// UnifiedFile represents the UnifiedFile schema from the OpenAPI specification
type UnifiedFile struct {
	Description string `json:"description,omitempty"` // Optional description of the file
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Permissions map[string]interface{} `json:"permissions,omitempty"` // Permissions the current user has on this file.
	Path string `json:"path,omitempty"` // The full path of the file or folder (includes the file name)
	Name string `json:"name"` // The name of the file
	TypeField string `json:"type"` // The type of resource. Could be file, folder or url
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Parent_folders []LinkedFolder `json:"parent_folders,omitempty"` // The parent folders of the file, starting from the root
	Parent_folders_complete bool `json:"parent_folders_complete,omitempty"` // Whether the list of parent folders is complete. Some connectors only return the direct parent of a file
	Size int `json:"size,omitempty"` // The size of the file in bytes
	Downstream_id string `json:"downstream_id,omitempty"` // The third-party API ID of original entity
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Exportable bool `json:"exportable,omitempty"` // Whether the current file is exportable to other file formats. This property is relevant for proprietary file formats such as Google Docs or Dropbox Paper.
	Owner Owner `json:"owner,omitempty"`
	Id string `json:"id"` // A unique identifier for an object.
	Downloadable bool `json:"downloadable,omitempty"` // Whether the current user can download this file
	Export_formats []string `json:"export_formats,omitempty"` // The available file formats when exporting this file.
	Mime_type string `json:"mime_type,omitempty"` // The MIME type of the file.
}

// CreateUploadSessionResponse represents the CreateUploadSessionResponse schema from the OpenAPI specification
type CreateUploadSessionResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// CreateFolderResponse represents the CreateFolderResponse schema from the OpenAPI specification
type CreateFolderResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// CreateDriveResponse represents the CreateDriveResponse schema from the OpenAPI specification
type CreateDriveResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// UpdateDriveGroupResponse represents the UpdateDriveGroupResponse schema from the OpenAPI specification
type UpdateDriveGroupResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// UpdateUploadSessionResponse represents the UpdateUploadSessionResponse schema from the OpenAPI specification
type UpdateUploadSessionResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// DeleteSharedLinkResponse represents the DeleteSharedLinkResponse schema from the OpenAPI specification
type DeleteSharedLinkResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// DriveGroupsFilter represents the DriveGroupsFilter schema from the OpenAPI specification
type DriveGroupsFilter struct {
	Parent_group_id string `json:"parent_group_id,omitempty"` // ID of the drive group to filter on
}

// UpdateFolderResponse represents the UpdateFolderResponse schema from the OpenAPI specification
type UpdateFolderResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// SharedLink represents the SharedLink schema from the OpenAPI specification
type SharedLink struct {
	Password_protected bool `json:"password_protected,omitempty"` // Indicated if the shared link is password protected.
	Download_url string `json:"download_url,omitempty"` // The URL that can be used to download the file.
	Target SharedLinkTarget `json:"target,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Url string `json:"url,omitempty"` // The URL that can be used to view the file.
	Password string `json:"password,omitempty"` // Optional password for the shared link.
	Target_id string `json:"target_id"` // The ID of the file or folder to link.
	Expires_at string `json:"expires_at,omitempty"`
	Scope string `json:"scope,omitempty"` // The scope of the shared link.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
}

// CreateUploadSessionRequest represents the CreateUploadSessionRequest schema from the OpenAPI specification
type CreateUploadSessionRequest struct {
	Size int `json:"size"` // The size of the file in bytes
	Drive_id string `json:"drive_id,omitempty"` // ID of the drive to upload to.
	Name string `json:"name"` // The name of the file.
	Parent_folder_id string `json:"parent_folder_id"` // The parent folder to create the new file within.
}

// FilesSearch represents the FilesSearch schema from the OpenAPI specification
type FilesSearch struct {
	Drive_id string `json:"drive_id,omitempty"` // ID of the drive to filter on
	Query string `json:"query"` // The query to search for. May match across multiple fields.
}

// UpdateFolderRequest represents the UpdateFolderRequest schema from the OpenAPI specification
type UpdateFolderRequest struct {
	Description string `json:"description,omitempty"` // Optional description of the folder.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Name string `json:"name,omitempty"` // The name of the folder.
	Parent_folder_id string `json:"parent_folder_id,omitempty"` // The parent folder to create the new file within.
}

// DriveGroup represents the DriveGroup schema from the OpenAPI specification
type DriveGroup struct {
	Description string `json:"description,omitempty"` // A description of the object.
	Name string `json:"name"` // The name of the drive group
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Display_name string `json:"display_name,omitempty"` // The display name of the drive group
	Id string `json:"id"` // A unique identifier for an object.
}

// FileStorageWebhookEvent represents the FileStorageWebhookEvent schema from the OpenAPI specification
type FileStorageWebhookEvent struct {
	Unified_api string `json:"unified_api,omitempty"` // Name of Apideck Unified API
	Execution_attempt float64 `json:"execution_attempt,omitempty"` // The current count this request event has been attempted
	Service_id string `json:"service_id,omitempty"` // Service provider identifier
	Entity_id string `json:"entity_id,omitempty"` // The service provider's ID of the entity that triggered this event
	Entity_type string `json:"entity_type,omitempty"` // The type entity that triggered this event
	Entity_url string `json:"entity_url,omitempty"` // The url to retrieve entity detail.
	Occurred_at string `json:"occurred_at,omitempty"` // ISO Datetime for when the original event occurred
	Consumer_id string `json:"consumer_id,omitempty"` // Unique consumer identifier. You can freely choose a consumer ID yourself. Most of the time, this is an ID of your internal data model that represents a user or account in your system (for example account:12345). If the consumer doesn't exist yet, Vault will upsert a consumer based on your ID.
	Event_id string `json:"event_id,omitempty"` // Unique reference to this request event
	Event_type string `json:"event_type,omitempty"`
}

// GetDrivesResponse represents the GetDrivesResponse schema from the OpenAPI specification
type GetDrivesResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Drive `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// PassThroughQuery represents the PassThroughQuery schema from the OpenAPI specification
type PassThroughQuery struct {
	Example_downstream_property string `json:"example_downstream_property,omitempty"` // All passthrough query parameters are passed along to the connector as is (?pass_through[search]=leads becomes ?search=leads)
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// UnifiedId represents the UnifiedId schema from the OpenAPI specification
type UnifiedId struct {
	Id string `json:"id"` // The unique identifier of the resource
}

// UpdateFileRequest represents the UpdateFileRequest schema from the OpenAPI specification
type UpdateFileRequest struct {
	Description string `json:"description,omitempty"` // Optional description of the file.
	Name string `json:"name,omitempty"` // The name of the file.
	Parent_folder_id string `json:"parent_folder_id,omitempty"` // The parent folder to create the new file within.
}

// Drive represents the Drive schema from the OpenAPI specification
type Drive struct {
	Description string `json:"description,omitempty"` // A description of the object.
	Id string `json:"id"` // A unique identifier for an object.
	Name string `json:"name"` // The name of the drive
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
}

// UnprocessableResponse represents the UnprocessableResponse schema from the OpenAPI specification
type UnprocessableResponse struct {
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// CustomMappings represents the CustomMappings schema from the OpenAPI specification
type CustomMappings struct {
}

// GetDriveGroupResponse represents the GetDriveGroupResponse schema from the OpenAPI specification
type GetDriveGroupResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data DriveGroup `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// DrivesFilter represents the DrivesFilter schema from the OpenAPI specification
type DrivesFilter struct {
	Group_id string `json:"group_id,omitempty"` // ID of the drive group to filter on
}

// GetFilesResponse represents the GetFilesResponse schema from the OpenAPI specification
type GetFilesResponse struct {
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []UnifiedFile `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
}

// GetUploadSessionResponse represents the GetUploadSessionResponse schema from the OpenAPI specification
type GetUploadSessionResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UploadSession `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// CreateFileRequest represents the CreateFileRequest schema from the OpenAPI specification
type CreateFileRequest struct {
	Parent_folder_id string `json:"parent_folder_id"` // The parent folder to create the new file within.
	Description string `json:"description,omitempty"` // Optional description of the file.
	Drive_id string `json:"drive_id,omitempty"` // ID of the drive to upload to.
	Name string `json:"name"` // The name of the file.
}

// DeleteDriveGroupResponse represents the DeleteDriveGroupResponse schema from the OpenAPI specification
type DeleteDriveGroupResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// Folder represents the Folder schema from the OpenAPI specification
type Folder struct {
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Parent_folders []LinkedFolder `json:"parent_folders"` // The parent folders of the file, starting from the root
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Owner Owner `json:"owner,omitempty"`
	Parent_folders_complete bool `json:"parent_folders_complete,omitempty"` // Whether the list of parent folder is complete. Some connectors only return the direct parent of a folder
	Path string `json:"path,omitempty"` // The full path of the folder (includes the folder name)
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Name string `json:"name"` // The name of the folder
	Size int `json:"size,omitempty"` // The size of the folder in bytes
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Description string `json:"description,omitempty"` // Optional description of the folder
}

// GetFolderResponse represents the GetFolderResponse schema from the OpenAPI specification
type GetFolderResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Folder `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// GetDriveResponse represents the GetDriveResponse schema from the OpenAPI specification
type GetDriveResponse struct {
	Data Drive `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CreateSharedLinkResponse represents the CreateSharedLinkResponse schema from the OpenAPI specification
type CreateSharedLinkResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// DeleteUploadSessionResponse represents the DeleteUploadSessionResponse schema from the OpenAPI specification
type DeleteUploadSessionResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// CreateFileResponse represents the CreateFileResponse schema from the OpenAPI specification
type CreateFileResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}
