package constant

import "errors"

var (
	// Login Error
	ErrUserLoginFormat             = errors.New("incorrect format of username or password")
	ErrUserLoginAuthentication     = errors.New("incorrect username or password")
	ErrUserLoginJwtTokenGeneration = errors.New("fail to generate jwt token")
	ErrUserLoginUpdateUserStatus   = errors.New("fail to update user's login status")

	ErrUserCheckParamIncorrect = errors.New("no username or email in params")
	ErrUserCheckFormat         = errors.New("incorrect format of username or email")
	ErrUserCheckNameConflict   = errors.New("username conflict")
	ErrUserCheckEmailConflict  = errors.New("email conflict")

	ErrUserRegisterFailServerError = errors.New("fail to create the user")
	ErrUserRegisterFormat          = errors.New("incorrect format of username or email or password")
	ErrUserRegisterNameConflict    = errors.New("username conflict")
	ErrUserRegisterEmailConflict   = errors.New("email conflict")

	ErrRedisSetKeyJsonMarshal = errors.New("fail to marshal struct to json")

	ErrDatabase                      = errors.New("database error")
	ErrTourIDNotNumber               = errors.New("param id is not a number")
	ErrTourNotFound                  = errors.New("tour record not found")
	ErrTourAddFormatIncorrect        = errors.New("collects id null")
	ErrTourAddCollectsRecordNotFound = errors.New("collects record not found")

	ErrCollectIDNotNumber            = errors.New("param id is not a number")
	ErrCollectNotFound               = errors.New("collect record not found")
	ErrCollectAddFormatIncorrect     = errors.New("songs not provided or title is empty string")
	ErrCollectAddSongsRecordNotFound = errors.New("songs record not found")

	ErrSongFormatIncorrect            = errors.New("song param format incorrect")
	ErrSongLyricsFileTypeNotSupported = errors.New("lyrics file type not supported")
	ErrSongParseLyrics                = errors.New("parse lyrics file error")
	ErrSongMissLyricsIncorrect        = errors.New("miss lyrics id negative or exceed total num of lyrics")
	ErrSongIDNotNumber                = errors.New("param id is not a number")
	ErrSongNotFound                   = errors.New("song recordnot found")
)
