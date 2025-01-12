package opt

import (
	"context"
	"fmt"
	"github.com/ZRothschild/ldp/gen/common"
	"github.com/ZRothschild/ldp/infrastr/static/code"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"net/textproto"
	"strings"
)

func CustomForwardResponseRewriter(ctx context.Context, response proto.Message) (any, error) {
	return response, nil
}

func CustomOutgoingHeaderMatcher(key string) (string, bool) {
	return fmt.Sprintf("%s%s", runtime.MetadataHeaderPrefix, key), true
}

func CustomOutgoingTrailerMatcher(key string) (string, bool) {
	return fmt.Sprintf("%s%s", runtime.MetadataTrailerPrefix, key), true
}

func CustomHTTPErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	// return Internal when Marshal failed
	const fallback = `{"errorCode": 13, "errorMessage": "failed to marshal error message"}`
	const fallbackRewriter = `{"errorCode": 13, "errorMessage": "failed to rewrite error message"}`

	var customStatus *runtime.HTTPStatusError
	if errors.As(err, &customStatus) {
		err = customStatus.Err
	}

	s := status.Convert(err)
	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	respRw, err := CustomForwardResponseRewriter(ctx, s.Proto())
	if err != nil {
		grpclog.Errorf("Failed to rewrite error message %q: %v", s, err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallbackRewriter); err != nil {
			grpclog.Errorf("Failed to write response: %v", err)
		}
		return
	}

	contentType := marshaler.ContentType(respRw)
	w.Header().Set("Content-Type", contentType)

	if s.Code() == codes.Unauthenticated {
		w.Header().Set("WWW-Authenticate", s.Message())
	}
	ss := &common.CommonResp{
		ErrorCode:    int64(s.Code()),
		ErrorMessage: s.Message(),
		ShowType:     common.ErrorShowType_ERROR_MESSAGE,
	}
	buf, merr := marshaler.Marshal(ss)
	if merr != nil {
		grpclog.Errorf("Failed to marshal error message %q: %v", s, merr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			grpclog.Errorf("Failed to write response: %v", err)
		}
		return
	}

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Error("Failed to extract ServerMetadata from context")
	}

	// handleForwardResponseServerMetadata(w, mux, md)
	for k, vs := range md.HeaderMD {
		if h, ok := CustomOutgoingHeaderMatcher(k); ok {
			for _, v := range vs {
				w.Header().Add(h, v)
			}
		}
	}

	// RFC 7230 https://tools.ietf.org/html/rfc7230#section-4.1.2
	// Unless the request includes a TE header field indicating "trailers"
	// is acceptable, as described in Section 4.3, a server SHOULD NOT
	// generate trailer fields that it believes are necessary for the user
	// agent to receive.

	//doForwardTrailers := requestAcceptsTrailers(r)
	doForwardTrailers := strings.Contains(strings.ToLower(r.Header.Get("TE")), "trailers")

	if doForwardTrailers {
		//handleForwardResponseTrailerHeader(w, mux, md)
		for k := range md.TrailerMD {
			if h, ok := CustomOutgoingTrailerMatcher(k); ok {
				w.Header().Add("Trailer", textproto.CanonicalMIMEHeaderKey(h))
			}
		}
		w.Header().Set("Transfer-Encoding", "chunked")
	}

	st := runtime.HTTPStatusFromCode(s.Code())
	if customStatus != nil {
		st = customStatus.HTTPStatus
	}
	// 如果是自定错误就使用 http status 200
	if s.Code() > code.MinCode {
		st = http.StatusOK
	}
	w.WriteHeader(st)

	if _, err := w.Write(buf); err != nil {
		grpclog.Errorf("Failed to write response: %v", err)
	}

	if doForwardTrailers {
		// handleForwardResponseTrailer
		for k, vs := range md.TrailerMD {
			if h, ok := CustomOutgoingTrailerMatcher(k); ok {
				for _, v := range vs {
					w.Header().Add(h, v)
				}
			}
		}
	}
}
