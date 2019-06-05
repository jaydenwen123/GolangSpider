package gifshow

import (
	"testing"
)

func TestDownloadHotVideo(t *testing.T) {
	DownloadHotVideo(1,nil)
	//time.Sleep(time.Second*30)

}

func TestDownloadWithBatch(t *testing.T) {
	DownloadWithBatch(1)
}
