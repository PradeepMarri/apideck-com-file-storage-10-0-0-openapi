package main

import (
	"github.com/file-storage-api/mcp-server/config"
	"github.com/file-storage-api/mcp-server/models"
	tools_drive_groups "github.com/file-storage-api/mcp-server/tools/drive_groups"
	tools_upload_sessions "github.com/file-storage-api/mcp-server/tools/upload_sessions"
	tools_files "github.com/file-storage-api/mcp-server/tools/files"
	tools_folders "github.com/file-storage-api/mcp-server/tools/folders"
	tools_drives "github.com/file-storage-api/mcp-server/tools/drives"
	tools_shared_links "github.com/file-storage-api/mcp-server/tools/shared_links"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_drive_groups.CreateDrivegroupsallTool(cfg),
		tools_drive_groups.CreateDrivegroupsaddTool(cfg),
		tools_upload_sessions.CreateUploadsessionsdeleteTool(cfg),
		tools_upload_sessions.CreateUploadsessionsoneTool(cfg),
		tools_upload_sessions.CreateUploadsessionsuploadTool(cfg),
		tools_files.CreateFilesdownloadTool(cfg),
		tools_files.CreateFilesexportTool(cfg),
		tools_folders.CreateFolderscopyTool(cfg),
		tools_upload_sessions.CreateUploadsessionsaddTool(cfg),
		tools_upload_sessions.CreateUploadsessionsfinishTool(cfg),
		tools_drives.CreateDrivesallTool(cfg),
		tools_drives.CreateDrivesaddTool(cfg),
		tools_files.CreateFilessearchTool(cfg),
		tools_folders.CreateFoldersaddTool(cfg),
		tools_drive_groups.CreateDrivegroupsoneTool(cfg),
		tools_drive_groups.CreateDrivegroupsupdateTool(cfg),
		tools_drive_groups.CreateDrivegroupsdeleteTool(cfg),
		tools_shared_links.CreateSharedlinksdeleteTool(cfg),
		tools_shared_links.CreateSharedlinksoneTool(cfg),
		tools_shared_links.CreateSharedlinksupdateTool(cfg),
		tools_files.CreateFilesallTool(cfg),
		tools_files.CreateFilesuploadTool(cfg),
		tools_folders.CreateFoldersoneTool(cfg),
		tools_folders.CreateFoldersupdateTool(cfg),
		tools_folders.CreateFoldersdeleteTool(cfg),
		tools_drives.CreateDrivesoneTool(cfg),
		tools_drives.CreateDrivesupdateTool(cfg),
		tools_drives.CreateDrivesdeleteTool(cfg),
		tools_files.CreateFilesdeleteTool(cfg),
		tools_files.CreateFilesoneTool(cfg),
		tools_files.CreateFilesupdateTool(cfg),
		tools_shared_links.CreateSharedlinksaddTool(cfg),
		tools_shared_links.CreateSharedlinksallTool(cfg),
	}
}
