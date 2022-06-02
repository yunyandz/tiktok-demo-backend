package util

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ReadFrameAsJpeg(inFileName string) (io.Reader, error) {
	frameNum := rand.Int()%10 + 1
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	return buf, nil
}
