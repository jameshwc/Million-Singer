package constant

const (
	SUCCESS        = 200
	INVALID_PARAMS = 400
	UNAUTHORIZED   = 401
	SERVER_ERROR   = 500
	// 100xx tour
	ERROR_GET_TOUR_FAIL               = 10000
	ERROR_GET_TOUR_ID_NOT_NUM         = 10001
	ERROR_GET_TOUR_NO_RECORD          = 10002
	ERROR_ADD_TOUR_FORMAT_INCORRECT   = 10011
	ERROR_ADD_TOUR_NO_COLLECTS_RECORD = 10012
	ERROR_DEL_TOUR_ID_INCORRECT       = 10020
	ERROR_DEL_TOUR_DELETED            = 10021

	// 200xx collect
	ERROR_ADD_COLLECT_UNKNOWN            = 20001
	ERROR_ADD_COLLECT_SONG_NAN           = 20002
	ERROR_ADD_COLLECT_NO_SONGID_OR_TITLE = 20003
	ERROR_ADD_COLLECT_NO_SONGID_RECORD   = 20004
	ERROR_ADD_COLLECT_SERVER_ERROR       = 20005
	ERROR_GET_COLLECT_FAIL_UNKNOWN       = 20010
	ERROR_GET_COLLECT_ID_NAN             = 20011
	ERROR_GET_COLLECT_NO_RECORD          = 20012
	ERROR_DEL_COLLECT_ID_INCORRECT       = 20020
	ERROR_DEL_COLLECT_DELETED            = 20021
	ERROR_DEL_COLLECT_FOREIGN_KEY        = 20022

	// 300xx song
	ERROR_ADD_SONG_FAIL_UNKNOWN                   = 30001
	ERROR_ADD_SONG_FORMAT_INCORRECT               = 30002
	ERROR_ADD_SONG_LYRICS_FILE_TYPE_NOT_SUPPORTED = 30003
	ERROR_ADD_SONG_PARSE_LYRICS_ERROR             = 30004
	ERROR_ADD_SONG_MISS_LYRICS_ERROR              = 30005
	ERROR_ADD_SONG_SERVER_ERROR                   = 30006
	ERROR_ADD_SONG_DUPLICATE                      = 30007
	ERROR_ADD_SONG_URL_INCORRECT                  = 30008
	ERROR_GET_SONG_ID_NAN                         = 30011
	ERROR_GET_SONG_NO_RECORD                      = 30012
	ERROR_GET_SONG_SERVER_ERROR                   = 30013
	ERROR_DEL_SONG_ID_INCORRECT                   = 30020
	ERROR_DEL_SONG_DELETED                        = 30021
	ERROR_DEL_SONG_FOREIGN_KEY                    = 30022

	// 500xx user
	ERROR_REGISTER_FAIL                   = 50000
	ERROR_REGISTER_USERNAME_CONFLICT      = 50001
	ERROR_REGISTER_EMAIL_CONFLICT         = 50002
	ERROR_REGISTER_FAIL_SERVER_ERROR      = 50003
	ERROR_REGISTER_FORMAT_INCORRECT       = 50004
	ERROR_LOGIN_FAIL_UNKNOWN              = 50010
	ERROR_LOGIN_FAIL_FORMAT_INCORRECT     = 50011
	ERROR_LOGIN_FAIL_AUTHENTICATION       = 50012
	ERROR_LOGIN_FAIL_JWT_TOKEN_GENERATION = 50013
	ERROR_LOGIN_FAIL_UPDATE_LOGIN_STATUS  = 50014
	ERROR_CHECK_PARAM_INCORRECT           = 50021
	ERROR_CHECK_FORMAT_INCORRECT          = 50022
	ERROR_CHECK_NAME_CONFLICT             = 50023
	ERROR_CHECK_EMAIL_CONFLICT            = 50024
	ERROR_AUTH_TOKEN_TIMEOUT              = 50031
	ERROR_AUTH_TOKEN_FAIL                 = 50032

	// 900xx subtitle
	ERROR_UPLOAD_SRT_FILE                 = 90002
	ERROR_SRT_FILE_FORMAT                 = 90003
	ERROR_GET_CAPTION                     = 90004
	ERROR_CONVERT_FILE_PARSE              = 90010
	ERROR_CONVERT_FILE_TYPE_NOT_SUPPORTED = 90011
	ERROR_DOWNLOAD_YOUTUBE_SUBTITLE       = 90020
	ERROR_GET_YOUTUBE_TITLE               = 90021
)
