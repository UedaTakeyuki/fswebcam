package fswebcam

import (
	"testing"

	fw "local.packages/fswebcam"
)

func Test(t *testing.T) {
	fw.Capture("fswebcam -d v4l2:/dev/video0 kerokero.jpg")
}
