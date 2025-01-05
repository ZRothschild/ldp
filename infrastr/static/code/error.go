package code

import (
	"fmt"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

type (
	CustomError struct {
		RequestId string `json:"requestId"`
		Code      codes.Code
		Message   string
		Details   []*anypb.Any
	}
)

func (c CustomError) Error() string {
	return fmt.Sprintf("rpc error: code = %d desc = %s detail: %v requestId: %v", c.Code, c.Message, c.Details, c.RequestId)
}

func (c CustomError) GRPCStatus() *status.Status {
	return status.FromProto(&spb.Status{
		Code:    int32(c.Code),
		Message: c.Message,
		Details: c.Details,
	})
}
