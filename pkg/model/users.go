package model

import "github.com/google/uuid"

type Group struct {
	ID uuid.UUID
	// Discord role for members of the group
	RoleID string
	// Translation key for the name of the group
	Name uuid.UUID
	// Translation key for the group's description
	Description uuid.NullUUID
	// Category id for the group
	TicketCategoryID string
}

type User struct {
	ID        uuid.UUID
	DiscordID string
	AvatarURL string
	Group     uuid.UUID
	Admin     bool
}

type PermissionBinding struct {
	ID    uuid.UUID
	Group uuid.UUID
	User  uuid.UUID

	Admin bool
}
