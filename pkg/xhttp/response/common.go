package response

import (
	"go.uber.org/zap"
)

const (
	httpCodeTag = "HTTP_CODE"
	rqUrlTag    = "RQ_URL"
)

type HttpResponseWrapper struct {
	logger *zap.Logger
}

func NewHttpResponseWrapper(logger *zap.Logger) *HttpResponseWrapper {
	return &HttpResponseWrapper{logger: logger}
}
