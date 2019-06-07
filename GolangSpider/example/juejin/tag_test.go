package juejin

import (
	"github.com/jaydenwen123/go-util"
	"testing"
)

func TestGetAllTags(t *testing.T) {
	//GetAllTags(GetHotTagUrl)
	GetAllTags(GetNewTagUrl)
}

func TestSaveAllTagPageArticles(t *testing.T) {
	util.InitDir(MARKDOWN_TAG_DIR)
	SaveAllTagPageArticles()
}