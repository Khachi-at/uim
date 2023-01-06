package wire

import "time"

// Algorithm in routing.
const (
	AlgorithmHashSlots = "hashslots"
)

// Command defined data type between client and server.
const (
	// login.
	CommandLoginSignIn  = "login.signin"
	CommandLoginSignOut = "login.signout"

	// chat.
	CommandChatUserTalk  = "chat.user.talk"
	CommandChatGroupTalk = "chat.group.talk"
	CommandChatTalkAck   = "chat.talk.ack"

	// offline.
	CommandOfflineIndex   = "chat.offline.index"
	CommandOfflineContent = "chat.offline.content"

	// Group.
	CommandGroupCreate  = "chat.group.crate"
	CommandGroupJoin    = "chat.group.join"
	CommandGroupQuit    = "chat.group.quit"
	CommandGroupMembers = "chat.group.members"
	CommandGroupDetail  = "chat.group.detail"
)

// Meta key of a packet
const (
	// The packet destination's serviceName.
	MetaDestServer = "dest.server"
	// The packet destination's channels.
	MetaDestChannels = "dest.channels"
)

type Protocol string

const (
	ProtocolTCP       Protocol = "tcp"
	ProtocolWebsocket Protocol = "websocket"
)

const (
	SNWGateway = "wgateway"
	SNTGateway = "tgateway"
	SNLogin    = "chat"
	SNChat     = "chat"
	SNService  = "royal"
)

type Magic [4]byte

var (
	MagicLogicPkt = Magic{0xc3, 0x11, 0xa3, 0x65}
	MagicBasicPkt = Magic{0xc3, 0x15, 0xa7, 0x65}
)

const (
	OfflineReadIndexExpiresIn = time.Hour * 24 * 30
	OfflineSyncIndexCount     = 2000
	OfflineMessageExpiresIn   = 15
	MessageMaxCountPerPage    = 200
)
