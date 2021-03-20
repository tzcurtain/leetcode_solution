import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=1487 lang=golang
 *
 * [1487] 保证文件名唯一
 */

// @lc code=start
func getFolderNames(names []string) []string {
	dict := make(map[string]int, 1)
	var res []string

	for i := range names {
		now := names[i]
		if dict[now] != 0 {
			suffixnum := dict[names[i]]
			nowname := names[i] + "(" + strconv.Itoa(suffixnum) + ")"
			for dict[nowname] != 0 {
				suffixnum++
				nowname = names[i] + "(" + strconv.Itoa(suffixnum) + ")"
			}
			dict[now] = suffixnum
			dict[nowname] = 1
			now = nowname
		} else {
			dict[now] = 1
		}
		// fmt.Println(now, dict)
		res = append(res, now)
	}

	return res
}

// @lc code=end

