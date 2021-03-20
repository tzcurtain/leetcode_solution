/*
 * @lc app=leetcode.cn id=365 lang=golang
 *
 * [365] 水壶问题
 */

// @lc code=start

// better solution
// 裴蜀定理
// ax+by=m 有整数解时当且仅当m是a及b的最大公约数d的倍数

type intpair struct {
	a, b int
}

func canMeasureWater(x int, y int, z int) bool {
	m := make(map[intpair]bool, 10)

	return canMeasureWaterRecur(0, 0, x, y, z, &m)
}

func canMeasureWaterRecur(a, b, x, y, z int, m *map[intpair]bool) bool {
	if a == z || b == z || a+b == z {
		return true
	}
	ab := intpair{a, b}
	if (*m)[ab] {
		return false
	}

	(*m)[ab] = true

	var atob, btoa bool
	if a+b > x {
		atob = canMeasureWaterRecur(x, b-(x-a), x, y, z, m)
	} else {
		atob = canMeasureWaterRecur(a+b, 0, x, y, z, m)
	}

	if atob {
		return true
	}

	if a+b > y {
		btoa = canMeasureWaterRecur(a-(y-b), y, x, y, z, m)
	} else {
		btoa = canMeasureWaterRecur(0, a+b, x, y, z, m)
	}

	return btoa || canMeasureWaterRecur(0, b, x, y, z, m) ||
		canMeasureWaterRecur(a, 0, x, y, z, m) || // 清空任意一个水壶
		canMeasureWaterRecur(x, b, x, y, z, m) ||
		canMeasureWaterRecur(a, y, x, y, z, m) // 倒满任意个水壶

}

// @lc code=end

