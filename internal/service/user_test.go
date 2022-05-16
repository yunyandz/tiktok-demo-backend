package service

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/dao/mysql"
	"github.com/yunyandz/tiktok-demo-backend/internal/dao/redis"
	"github.com/yunyandz/tiktok-demo-backend/internal/logger"
)

var UserTestSvr *Service

func init() {
	cfg, err := config.Phase()
	if err != nil {
		panic(err)
	}
	mylogger := logger.New(cfg)
	db := mysql.New(cfg, mylogger)
	rds := redis.New(cfg)
	UserTestSvr = New(db, rds, mylogger)
}

func TestService_Register(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		response, err := UserTestSvr.Register(s, s)
		require.NoError(t, err)
		require.NotEmpty(t, response)
		require.NotEmpty(t, response.Token)
		require.Equal(t, uint64(i+1), response.UserID)
		require.Equal(t, int32(0), response.Response.StatusCode)
		require.Equal(t, "ok", response.Response.StatusMsg)
	}

}

func TestService_Login(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		response, err := UserTestSvr.Login(s, s)
		require.NoError(t, err)
		require.NotEmpty(t, response)
		require.NotEmpty(t, response.Token)
		require.Equal(t, uint64(i+1), response.UserID)
		require.Equal(t, int32(0), response.Response.StatusCode)
		require.Equal(t, "ok", response.Response.StatusMsg)
	}

}

func TestService_GetUserInfo(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		response, err := UserTestSvr.GetUserInfo(uint64(i + 1))
		require.NoError(t, err)
		require.NotEmpty(t, response)
		require.Equal(t, uint64(i+1), response.User.ID)
		require.Equal(t, s, response.User.Username)

		// have not give real value which is init value now
		require.Zero(t, response.User.FollowCount)
		require.Zero(t, response.User.FollowerCount)
		require.Equal(t, false, response.User.IsFollow)

		require.Equal(t, int32(0), response.Response.StatusCode)
		require.Equal(t, "ok", response.Response.StatusMsg)
	}
}
