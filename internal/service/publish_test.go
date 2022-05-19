package service_test

import (
	"context"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"moul.io/zapgorm2"

	"github.com/golang/mock/gomock"
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
	mock_s3 "github.com/yunyandz/tiktok-demo-backend/internal/s3/mock"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PublishTestService *service.Service
var cfg config.Config
var mdb *gorm.DB
var mylogger *zap.Logger
var mockS3API *mock_s3.MockS3ObjectAPI

func init() {
	cfg = config.Config{
		S3: config.S3{
			AccessKey: "",
			SecretKey: "",
			Region:    "",
		},
		Mysql: config.Mysql{
			Host:     "localhost",
			Port:     "3306",
			User:     "tiktok",
			Password: "tiktok",
			Database: "tiktok",
		},
	}
	var err error
	mylogger, _ = zap.NewProduction()
	mdb, err = gorm.Open(mysql.Open(cfg.Mysql.User+":"+cfg.Mysql.Password+"@tcp("+cfg.Mysql.Host+":"+cfg.Mysql.Port+")/"+cfg.Mysql.Database), &gorm.Config{
		Logger: zapgorm2.New(mylogger),
	})
	if err != nil {
		panic(err)
	}
}

func TestService_PublishVideo(t *testing.T) {
	mockctl := gomock.NewController(t)
	mockS3API := mock_s3.NewMockS3ObjectAPI(mockctl)
	PublishTestService = service.New(&cfg, mdb.Debug(), nil, mylogger, nil, mockS3API)
	um := model.NewUserModel(mdb, nil)
	uid, err := um.CreateUser(&model.User{
		Username: "test",
		Password: "test",
	})
	if err != nil {
		t.Error(err)
	}

	testvideo, err := os.OpenFile("../../public/bear.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening example video", err)
	}
	defer testvideo.Close()
	filename := "bear.mp4"

	for i := 0; i < 10; i++ {
		// TODO: 设置期望值
		response := PublishTestService.PublishVideo(context.TODO(), uid, filename, testvideo)
		if response.StatusCode != 0 {
			t.Fatalf("an error '%s' was not expected when opening example video", response.StatusMsg)
		}
	}
}
