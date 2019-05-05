package dsa

import (
	"github.com/pkg/errors"
)

var (
	errSubTooLong = errors.New("sub string is too long")
	errNotFound   = errors.New("not found sub string")
)

// 简单模式匹配
func simplePattern(main, sub string) (int, error) {
	if len(sub) > len(main) {
		return -1, errSubTooLong
	}
	for i := 0; i < len(main); i++ {
		var j int
		for j = 0; j < len(sub); j++ {
			// 从i开始，逐个往后匹配
			if main[i+j] != sub[j] {
				break
			}
		}
		if j == len(sub) {
			return i, nil
		}
	}
	return -1, errNotFound
}

// kmp算法
func kmp(main, sub string) (int, error) {
	if len(sub) > len(main) {
		return -1, errSubTooLong
	}
	next := next(sub)
	var i, j int
	ml := len(main)
	sl := len(sub)
	for i < ml && j < sl {
		// 如果子串"准备"从第一位开始 或 主串和子串当前字符匹配，则移向下一位
		if j == -1 || main[i] == sub[j] {
			i++
			j++
		} else { // 不匹配则子串从j的最长前后缀处开始匹配
			j = next[j] // j = next[j] = k
		}
	}
	if j == sl {
		return i - j, errNotFound
	}
	return i, nil
}

// 求每一位字符的最长前后缀
func next(str string) map[int]int {
	next := make(map[int]int)
	next[0] = -1 // 表示第一个字符
	l := len(str)
	k := -1 // 重复的前缀最后一位字符
	j := 0  // 向后遍历字符串
	for j < l {
		// 没有最长前后缀(k+1=0) 或 在之前的最长前后缀基础上又多一位
		// 如果没有最长前后缀，则重新跟第一个字符对比
		if k == -1 || str[j] == str[k] {
			k++
			j++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}
