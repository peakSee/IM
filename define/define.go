package define

// var MailPassword = os.Getenv("MailPassword")
var MailPassword = "YHWBBHWBIZRABUEB"

type MessageStruct struct {
	Message      string `json:"message"`
	RoomIdentity string `json:"room_identity"`
}

var RegisterPrefix = "TOKEN_"
var ExpireTime = 300
