package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpcProtobuf/common/auth/token"
	"grpcProtobuf/common/id"
	"io"
	"os"
	"strings"
)

const (
	// ImpersonateAccountHeader defines the header for account
	// id impersonation.
	ImpersonateAccountHeader = "impersonate-account-id"
	authorizationHeader      = "authorization"
	bearerPrefix             = "Bearer "
)

func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cant open publicKeyFile : %v", err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cant read publicKeyFile : %v", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cant parse publicKeyFile : %v", err)
	}
	i := &interceptor{verifier: &token.JWTTokenVerifier{PublicKey: pubKey}}
	return i.HandleReq, nil

}

func tokenFromContext(c context.Context) (tkn string, err error) {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", status.Errorf(codes.Unavailable, "")
	}
	for _, v := range m["authorization"] {
		if strings.HasPrefix(v, "Bearer ") {
			tkn = v[len("Bearer "):]
		}

	}
	if tkn == "" {
		return "", status.Errorf(codes.Unavailable, "")
	}
	return tkn, nil
}

type accountIDKey struct{}

// ContextWithAccountID creates a context with given account ID.
func ContextWithAccountID(c context.Context, aid id.AccountID) context.Context {
	return context.WithValue(c, accountIDKey{}, aid)
}

// AccountIDFromContext gets account id from context.
// Returns unauthenticated error if no account id is available.
func AccountIDFromContext(c context.Context) (id.AccountID, error) {
	v := c.Value(accountIDKey{})
	aid, ok := v.(id.AccountID)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}

type tokenVerifier interface {
	Verify(token string) (string, error)
}

type interceptor struct {
	verifier tokenVerifier
}

func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	aid, err := i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "token verify failed")
	}
	return handler(ContextWithAccountID(ctx, id.AccountID(aid)), req)
}
