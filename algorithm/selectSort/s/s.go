package s

func SelectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	min := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
	return arr
}