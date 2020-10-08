package constant

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	SERVER_ERROR:          "fail",
	INVALID_PARAMS:        "error parameters",
	ERROR_GET_TOUR_FAIL:   "fail to get the tour",
	ERROR_UPLOAD_SRT_FILE: "fail to upload the srt file",
	ERROR_SRT_FILE_FORMAT: "wrong srt file format",
	// 100xx tour
	ERROR_GET_TOUR_ID_NOT_NUM:         "param id is not a number",
	ERROR_GET_TOUR_NO_RECORD:          "tour not found",
	ERROR_ADD_TOUR_FORMAT_INCORRECT:   "collects id null",
	ERROR_ADD_TOUR_NO_COLLECTS_RECORD: "collects record not found",

	// 200xx collect
	ERROR_ADD_COLLECT_UNKNOWN:            "unknown error",
	ERROR_ADD_COLLECT_SONG_NAN:           "song id(s) not number(s)",
	ERROR_ADD_COLLECT_NO_SONGID_OR_TITLE: "no song id or title provided",
	ERROR_ADD_COLLECT_NO_SONGID_RECORD:   "songs record not found",
	ERROR_ADD_COLLECT_SERVER_ERROR:       "server error when add collects",
	ERROR_GET_COLLECT_FAIL_UNKNOWN:       "unknown error to get collects",
	ERROR_GET_COLLECT_ID_NAN:             "param id is not a number",
	ERROR_GET_COLLECT_NO_RECORD:          "collects record not found",

	// 300xx song
	ERROR_ADD_SONG_FAIL_UNKNOWN:                   "unknown error to add song",
	ERROR_ADD_SONG_FORMAT_INCORRECT:               "song param format incorrect",
	ERROR_ADD_SONG_LYRICS_FILE_TYPE_NOT_SUPPORTED: "lyrics file type not supported",
	ERROR_ADD_SONG_PARSE_LYRICS_ERROR:             "parse lyrics file error",
	ERROR_ADD_SONG_MISS_LYRICS_ERROR:              "miss lyrics id negative or exceed total num of lyrics",
	ERROR_ADD_SONG_SERVER_ERROR:                   "server error when add song",
	ERROR_GET_SONG_ID_NAN:                         "param id is not a number",
	ERROR_GET_SONG_NO_RECORD:                      "song record not found",
	ERROR_GET_SONG_SERVER_ERROR:                   "server error when get song",
	// 500xx user
	ERROR_REGISTER_FAIL:                   "fail to create the user",
	ERROR_REGISTER_USERNAME_CONFLICT:      "username conflict",
	ERROR_REGISTER_EMAIL_CONFLICT:         "email conflict",
	ERROR_REGISTER_FAIL_SERVER_ERROR:      "server error when register the user",
	ERROR_REGISTER_FORMAT_INCORRECT:       "incorrect format of username or email or password",
	ERROR_LOGIN_FAIL_UNKNOWN:              "unknown error to login",
	ERROR_LOGIN_FAIL_FORMAT_INCORRECT:     "incorrect format of username or password",
	ERROR_LOGIN_FAIL_AUTHENTICATION:       "incorrect username or password",
	ERROR_LOGIN_FAIL_JWT_TOKEN_GENERATION: "fail to generate jwt token",
	ERROR_LOGIN_FAIL_UPDATE_LOGIN_STATUS:  "fail to update user's login status",
	ERROR_CHECK_PARAM_INCORRECT:           "no username or email in params",
	ERROR_CHECK_FORMAT_INCORRECT:          "incorrect format of username or email",
	ERROR_CHECK_NAME_CONFLICT:             "username conflict",
	ERROR_CHECK_EMAIL_CONFLICT:            "email conflict",
	ERROR_AUTH_TOKEN_TIMEOUT:              "auth token timeout",
	ERROR_AUTH_TOKEN_FAIL:                 "auth token incorrect",
}

var SuccessMsg = "Success"

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[SERVER_ERROR]
}
