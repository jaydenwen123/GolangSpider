package leetcode

import "testing"

func TestGetAnswersByQuestion(t *testing.T) {

	GetAnswersByQuestion("two-sum")
	GetAnswersByQuestion("add-two-numbers")
	GetAnswersByQuestion("longest-substring-without-repeating-characters")
}

func TestGetAnswerDetail(t *testing.T)  {
	GetAnswerDetail("two-sum","liang-shu-zhi-he-by-gpe3dbjds1")
	GetAnswerDetail("two-sum","liang-shu-zhi-he-by-leetcode-2")
	GetAnswerDetail("two-sum","cyu-yan-yi-bian-ha-xi-kong-jian-onhuan-shi-jian-on")

	GetAnswerDetail("add-two-numbers","liang-shu-xiang-jia-by-leetcode")
	GetAnswerDetail("add-two-numbers","zhi-xing-yong-shi-8-ms-nei-cun-xiao-hao-406-mb-by-")
	GetAnswerDetail("add-two-numbers","liang-shu-xiang-jia-by-gpe3dbjds1")

	GetAnswerDetail("longest-substring-without-repeating-characters","wu-zhong-fu-zi-fu-de-zui-chang-zi-chuan-by-leetcod")
	GetAnswerDetail("longest-substring-without-repeating-characters","hua-dong-chuang-kou-by-powcai")
	GetAnswerDetail("longest-substring-without-repeating-characters","javati-jie-3wu-zhong-fu-zi-fu-de-zui-chang-zi-chua")
}
