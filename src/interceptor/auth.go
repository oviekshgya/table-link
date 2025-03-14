package interceptor

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
	"table-link/pkg/helper"
)

const SECRET_KEY = "rahasia_super_rahasia"

type AuthInterceptor struct {
	secretKey string
}

func NewAuthInterceptor(secretKey string) *AuthInterceptor {
	return &AuthInterceptor{secretKey: secretKey}
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		if strings.Contains(info.FullMethod, "/Login") || strings.Contains(info.FullMethod, "/CreateUser") {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("metadata not provided")
		}

		authHeader, exists := md["authorization"]
		if !exists || len(authHeader) == 0 {
			return nil, errors.New("missing authorization token")
		}

		tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
		_, _, _, err := helper.ValidateToken(tokenString)
		if err != nil {
			return nil, err
		}

		if err != nil {
			return nil, fmt.Errorf("Unauthorized: %v", err)
		}

		return handler(ctx, req)
	}
}
