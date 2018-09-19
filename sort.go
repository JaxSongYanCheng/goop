package sort

func BubbleSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func SelectSort(arr []int) {
	for i, _ := range arr {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func InsertSort(arr []int) {
	for i, _ := range arr {
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			} else {
				break
			}
		}
	}
}

func QuickSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	index := 0
	for i := 1; i < length; i++ {
		if arr[i] < arr[index] {
			if i > index+1 {
				arr[index], arr[index+1], arr[i] = arr[i], arr[index], arr[index+1]
			} else {
				arr[index], arr[i] = arr[i], arr[index]
			}
			index++
		}
	}
	QuickSort(arr[:index])
	QuickSort(arr[index+1:])
}
