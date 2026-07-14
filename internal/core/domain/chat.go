package domain

import "time"

type Chat struct {
	ID        int64
	CreatedAt time.Time
}

type ChatParticipant struct {
	UserID      int64
	DisplayName string
}

type ChatWithParticipant struct {
	Chat
	Participant ChatParticipant
}

type ChatListItem struct {
	Chat
	CompanionID          int64
	CompanionDisplayName string
}
