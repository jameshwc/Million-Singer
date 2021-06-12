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
	ErrTourAddCollectsDuplicate      = errors.New("collects id duplicate")
	ErrTourAddCollectsRecordNotFound = errors.New("collects record not found")
	ErrTourDelIDIncorrect            = errors.New("tour id negative or not found")
	ErrTourDelDeleted                = errors.New("tour has been deleted (or database error)")

	ErrCollectIDNotNumber        = errors.New("param id is not a number")
	ErrCollectNotFound           = errors.New("collect record not found")
	ErrCollectAddFormatIncorrect = errors.New("songs not provided or title is empty string")
	ErrCollectAddSongsDuplicate  = errors.New("songs id duplicate")
	ErrCollectDelIDIncorrect     = errors.New("collect id negative or not found")
	ErrCollectDelDeleted         = errors.New("collect has been deleted (or database error)")
	ErrCollectDelForeignKey      = errors.New("foreign key constraint: some tours are using this collect")

	ErrCollectAddSongsRecordNotFound = errors.New("songs record not found")

	ErrSongFormatIncorrect               = errors.New("song param format incorrect")
	ErrSongIDNotNumber                   = errors.New("param id is not a number")
	ErrSongNotFound                      = errors.New("song record not found")
	ErrSongAddDuplicate                  = errors.New("video duplicate in database")
	ErrSongAddURLIncorrect               = errors.New("youtube url not correct. Do you forgot to add protocol (http/https)?")
	ErrSongAddMissLyricsIncorrect        = errors.New("miss lyrics id negative or exceed total num of lyrics")
	ErrSongAddParseLyrics                = errors.New("parse lyrics file error")
	ErrSongAddLyricsFileTypeNotSupported = errors.New("lyrics file type not supported")
	ErrSongAddLyricsIndexDuplicate       = errors.New("lyrics index duplicate")
	ErrSongDelIDIncorrect                = errors.New("song id negative or not found")
	ErrSongDelDeleted                    = errors.New("song has been deleted (or database error)")
	ErrSongDelForeignKey                 = errors.New("foreign key constraint: some collects are using this song")

	ErrCaptionError                          = errors.New("caption error")
	ErrConvertFileToSubtiteParse             = errors.New("convert file to subtitle parse error")
	ErrConvertFileToSubtitleTypeNotSupported = errors.New("convert file to subtitle file type incorrect")
)
