package fswebcam

import (
	"testing"
	"github.com/UedaTakeyuki/fswebcam"
//	fw "local.packages/fswebcam"
)

func Test(t *testing.T) {
	fw.Capture("fswebcam -d v4l2:/dev/video0 kerokero.jpg")
}
