package token

import (
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

const PublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu1SU1LfVLPHCozMxH2Mo
4lgOEePzNm0tRgeLezV6ffAt0gunVTLw7onLRnrq0/IzW7yWR7QkrmBL7jTKEn5u
+qKhbwKfBstIs+bMY2Zkp18gnTxKLxoS2tFczGkPLPgizskuemMghRniWaoLcyeh
kd3qqGElvW/VDL5AaWTg0nLVkjRo9z+40RQzuVaE8AkAFmxZzow3x+VJYKdjykkJ
0iT9wCS0DRTXu269V264Vf/3jvredZiKRkgwlL9xNAwxXFg0x/XFw005UWVRIkdg
cKWTjpBP2dPwVZ4WWC+9aGVd+Gyn1o0CLelf4rEjGoXbAAEgAqeGUxrcIlbjXfbc
mwIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pk, err := jwt.ParseRSAPublicKeyFromPEM([]byte(PublicKey))
	if err != nil {
		t.Fatalf("cant parse public key: %v", err)
	}
	verifier := &JWTTokenVerifier{PublicKey: pk}
	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name: "valid_token",
			tkn:  "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzIyZTM1MWEzIn0.jPVRIZXsNz08OCudP4cC8KGzVEIWC42TOMHpc6cN-_3yUgbPcrhuJL6C27fzoxt0j8J3L0z6nv0ni_713fzYjo1Y_b4Axxz4sI5bz-b9O1BziFU1NC9t3IJbwFsF2Svz2OpG3aY388rTZ4orHShfRbrzGnzK8NbNXIZ7CcCvEznHiJEmSgqSZSYeZVjjid2p2l_T_eTQxJTkHi9LE-3g_AfLKLXXmqLlXYpurTGMWEBkJq51uNs6MnESi4pEwbLviTmZTTtC6qAhkVmeJh7QUZA8BPKoxSbNEYQxYYQK1aiRGyrrONsK1etXW6JG2F4x0wiNjTKMvQSAsq7GnWvkoQ",
			now:  time.Unix(1516239122, 0),
			want: "5f7c3168e2283aa722e351a3",
		},
		{
			name:    "token_expired",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzIyZTM1MWEzIn0.jPVRIZXsNz08OCudP4cC8KGzVEIWC42TOMHpc6cN-_3yUgbPcrhuJL6C27fzoxt0j8J3L0z6nv0ni_713fzYjo1Y_b4Axxz4sI5bz-b9O1BziFU1NC9t3IJbwFsF2Svz2OpG3aY388rTZ4orHShfRbrzGnzK8NbNXIZ7CcCvEznHiJEmSgqSZSYeZVjjid2p2l_T_eTQxJTkHi9LE-3g_AfLKLXXmqLlXYpurTGMWEBkJq51uNs6MnESi4pEwbLviTmZTTtC6qAhkVmeJh7QUZA8BPKoxSbNEYQxYYQK1aiRGyrrONsK1etXW6JG2F4x0wiNjTKMvQSAsq7GnWvkoQ",
			now:     time.Unix(1517239122, 0),
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "bad_token",
			now:     time.Unix(1517239122, 0),
			wantErr: true,
		},
		{
			name:    "wrong_signature",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzIyZTM1MWE0In0.jPVRIZXsNz08OCudP4cC8KGzVEIWC42TOMHpc6cN-_3yUgbPcrhuJL6C27fzoxt0j8J3L0z6nv0ni_713fzYjo1Y_b4Axxz4sI5bz-b9O1BziFU1NC9t3IJbwFsF2Svz2OpG3aY388rTZ4orHShfRbrzGnzK8NbNXIZ7CcCvEznHiJEmSgqSZSYeZVjjid2p2l_T_eTQxJTkHi9LE-3g_AfLKLXXmqLlXYpurTGMWEBkJq51uNs6MnESi4pEwbLviTmZTTtC6qAhkVmeJh7QUZA8BPKoxSbNEYQxYYQK1aiRGyrrONsK1etXW6JG2F4x0wiNjTKMvQSAsq7GnWvkoQ",
			now:     time.Unix(1516239122, 0),
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}
			accountID, err := verifier.Verify(c.tkn)

			if !c.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}

			if c.wantErr && err == nil {
				t.Errorf("want error; got no error")
			}

			if accountID != c.want {
				t.Errorf("wrong account id. want: %q, got: %q", c.want, accountID)
			}
		})
	}
}
