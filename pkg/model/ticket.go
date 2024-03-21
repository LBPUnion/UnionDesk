package model

import "github.com/google/uuid"

type Category struct {
	// The ID of the Category
	ID uuid.UUID
	// The Category's parent, null if it's a root category
	Parent uuid.NullUUID
	// TranslationKey for the Category's name
	Name uuid.UUID
	// If a ticket can be assigned to the Category
	Assignable bool
}

type Ticket struct {
	// The ID of the ticket
	ID uuid.UUID
	// Numerical ID of the ticket
	TicketID uint64
	// Prefix for the ticket, like "MOS" for a moderation ticket or "R&D" for a developer ticket, etc.
	IDPrefix string
	// Group that the ticket is assigned to
	AssignedGroup uuid.NullUUID
	// Discord ID of assigned user
	Assignee string
	// Category of the ticket
	Category uuid.UUID
	// Title of the ticket
	Title string
}

type Comment struct {
	// ID is the ID of a Comment
	ID uuid.UUID
	// TicketID is the ID of the parent Ticket
	TicketID uuid.UUID
	// Content is the text content of the ticket
	Content string
	// Sender is the discord id of the person who sent the message
	Sender string
}

type Attachment struct {
	// ID of the Attachment
	ID uuid.UUID
	// Discord ID of the file uploader
	Uploader string
	// SHA-256 hash of the file
	Hash string
	// Human-readable name of the file
	FileName string
	// Size of the file in bytes
	Size uint64
}
