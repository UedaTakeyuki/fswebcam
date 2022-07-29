# fswebcam
Go wrapper of github.com/fsphil/fswebcam inspired by github.com/nimbleindustry/fswebcam.

## How to use

### Set Up
go get this module, then execute **setup.sh** on the downloaded mod folder as follows
```
go get github.com/UedaTakeyuki/fswebcam
cd ~/go/pkg/mod/github.com/\!ueda\!takeyuki/fswebcam@*
sudo sh setup.sh
```

### Code
Import this module and call **Capture()** with fswebcam command line params start with **fswebcam** as follows.


```
import (
  fw "github.com/UedaTakeyuki/fswebcam"
)

func main(){
  // call Capture with fswebcam command line string
  fw.Capture([]string{"fswebcam", "-d", "v4l2:/dev/video0", "kerokero.jpg") 
}
```

# History
- 1.0.0  2022.07.29  first version.
