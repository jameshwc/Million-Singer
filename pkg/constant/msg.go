package constant

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	SERVER_ERROR:          "fail",
	INVALID_PARAMS:        "error parameters",
	ERROR_GET_TOUR_FAIL:   "fail to get the tour",
	ERROR_UPLOAD_SRT_FILE: "fail to upload the srt file",
	ERROR_SRT_FILE_FORMAT: "wrong srt file format",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[SERVER_ERROR]
}
