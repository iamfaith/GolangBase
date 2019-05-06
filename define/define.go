package define

const (
	CodeOk        = 0
	CodeBadParam  = -1
	CodeServerErr = -2
)

const (
	UploadFile = "upload_file"
)

type AuthStatus int

const (
	AuthErr AuthStatus = iota
	AuthNoLogin
	AuthNoUser
	AuthNoPerm
	AuthHasReadPerm
	AuthHasWritePerm
	AuthSuccess
)

const (
	LinkTbl = "link"
)
