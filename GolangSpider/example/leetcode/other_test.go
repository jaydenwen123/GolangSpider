package leetcode

import "testing"

func TestInitLeetcodeDir(t *testing.T) {
	InitLeetcodeDir()
}

func TestSaveDatabaseData(t *testing.T) {
	SaveDatabaseData()
}

func TestSaveShellData(t *testing.T) {
	SaveShellData()
}

func TestSaveTagData(t *testing.T) {
	SaveTagData()
}

func TestSaveTagDataOfType(t *testing.T)  {
	SaveTagDataOfType("stack","栈")
	SaveTagDataOfType("heap","堆")
	SaveTagDataOfType("greedy","贪心算法")
}