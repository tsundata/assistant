package exception

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrGrpcUnauthenticated = status.Error(codes.Unauthenticated, "grpc unauthenticated")
