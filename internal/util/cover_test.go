package util_test

import (
	"io"
	"os"
	"testing"

	"github.com/yunyandz/tiktok-demo-backend/internal/util"
)

func Test_cover(t *testing.T) {
	testvideo := "../../public/bear.mp4"
	testcover := "../../public/bear.jpg"
	cover, err := util.ReadFrameAsJpeg(testvideo)
	if err != nil {
		t.Error(err)
	}
	f, err := os.Create(testcover)
	if err != nil {
		t.Error(err)
	}
	_, err = io.Copy(f, cover)
	if err != nil {
		t.Error(err)
	}
}
