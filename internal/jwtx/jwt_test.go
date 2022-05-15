package jwtx

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateUserClaims(t *testing.T) {
	token, err := CreateUserClaims("hello world")
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestParseUserClaims(t *testing.T) {
	username := "yongteng"
	token, err := CreateUserClaims(username)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := ParseUserClaims(token)
	require.NoError(t, err)
	require.NotEmpty(t, claims)
	require.Equal(t, username, claims.Username)
	require.Equal(t, ISSUER, claims.Issuer)
	require.WithinDuration(t, time.Unix(claims.IssuedAt, 0), time.Unix(claims.ExpiresAt, 0), time.Second*EXPIRES_TIME)

}
