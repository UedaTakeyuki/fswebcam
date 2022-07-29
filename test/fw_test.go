package fswebcam

import (
	"testing"
	fw "github.com/UedaTakeyuki/fswebcam"
)

func Test(t *testing.T) {
	fw.Capture([]string{"fswebcam", "--no-timestamp", "--title", "©Atelier UEDA", "-d",  "v4l2:/dev/video0", "kerokero.jpg"})
}
