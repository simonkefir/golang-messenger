package domain

import "time"

type Chat struct {
	ID        int64
	CreatedAt time.Time
}

type ChatParticipant struct {
	UserID   int64
	Username string
}

type ChatWithParticipants struct {
	Chat
	Participants []ChatParticipant
}

type ChatListItem struct {
	Chat
	CompanionID   int64
	CompanionName string
}
