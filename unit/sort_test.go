package unit

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSortChinesebyUTF8(t *testing.T) {
	// input: 一乙丁七乃九二人八力十三丈久元六公四黑暗
	followingList := []User{{Name: "一"}, {Name: "乙"}, {Name: "丁"}, {Name: "七"}, {Name: "乃"}, {Name: "九"}, {Name: "二"}, {Name: "人"}, {Name: "八"}, {Name: "力"}, {Name: "十"}, {Name: "三"}, {Name: "丈"}, {Name: "久"}, {Name: "元"}, {Name: "六"}, {Name: "公"}, {Name: "四"}, {Name: "黑"}, {Name: "暗"}}
	followingList = sortChinesebyUTF8(followingList)
	str := ""
	for _, s := range followingList {
		str += s.Name
	}
	assert.Equal(t, "一丁七丈三乃久乙九二人元八公六力十四暗黑", str)

}
