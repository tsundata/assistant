package enum

type MessageType string

const (
	MessageTypeText     MessageType = "text"
	MessageTypeAudio    MessageType = "audio"
	MessageTypeImage    MessageType = "image"
	MessageTypeFile     MessageType = "file"
	MessageTypeLocation MessageType = "location"
	MessageTypeVideo    MessageType = "video"
	MessageTypeLink     MessageType = "link"
	MessageTypeScript   MessageType = "script"
	MessageTypeAction   MessageType = "action"
	MessageTypeForm     MessageType = "form"
)

const (
	ActionScript = "action"
)

const (
	MessageUserType  = "user"
	MessageGroupType = "group"
	MessageBotType   = "bot"
)

const (
	MessageIncomingDirection = "incoming"
	MessageOutgoingDirection = "outgoing"
)

const (
	InboxCreate = iota + 1
	InboxSend
	InboxRead
)
