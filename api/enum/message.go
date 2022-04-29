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
	MessageTypeTable    MessageType = "table"
	MessageTypeDigit    MessageType = "digit"
	MessageTypeOkr      MessageType = "okr"
	MessageTypeInfo     MessageType = "info"
	MessageTypeTodo     MessageType = "todo"
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

const (
	MessageCreatedStatus = iota
	MessageSendingStatus
	MessageSentSuccessStatus
	MessageSentFailedStatus
	MessageReadStatus
	MessageActionedStatus
	MessageArchivedStatus
)
