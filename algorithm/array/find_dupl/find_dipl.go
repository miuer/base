package find_dupl

/*

1-1000放在含有1001个元素的数组中，只有唯一的一个元素值重复，其它均只出现
一次。每个数组元素只能访问一次，设计一个算法，将它找出来；不用辅助存储空
间，能否设计一个算法实现？

*/

func miuerImpl(arr []int) int {
	// 求和 求差

	dupl := 0

	for _, v := range arr {
		dupl += v
	}

	sum := 1000 * (1 + 1000) / 2
	return dupl - sum
}
