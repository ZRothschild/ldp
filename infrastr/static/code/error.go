package code

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	MinCode                codes.Code = 100000
	PwdConfirmNotMatchCode codes.Code = 100010
	PwdConfirmNotMatchErr             = status.Error(PwdConfirmNotMatchCode, "Passwords do not match")
	PwdOrNicknameMatchCode codes.Code = 100020
	PwdOrNicknameMatchErr             = status.Error(PwdOrNicknameMatchCode, "user passwords or nickname do not match")
)
