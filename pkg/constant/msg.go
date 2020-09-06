package constant

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	SERVER_ERROR:   "fail",
	INVALID_PARAMS: "error parameters",
	ERROR_GET_GAME_FAIL: "fail to get the game"
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[SERVER_ERROR]
}
