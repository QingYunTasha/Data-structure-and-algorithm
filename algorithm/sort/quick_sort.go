package sort

func qSortOutMem(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	} else if len(nums) == 1 {
		return nums
	}

	l, r, pivot := []int{}, []int{}, nums[0]
	for _, num := range nums[1:] {
		if num < pivot {
			l = append(l, num)
		} else {
			r = append(r, num)
		}
	}

	l = qSortOutMem(l)
	l = append(l, pivot)
	return append(l, qSortOutMem(r)...)
}

func qSortInMem(nums []int, l, r int) {
	if l >= r {
		return
	}

	i, j, pivot := l+1, r, nums[(l+r)>>1]
	for i < j {
		for nums[i] < pivot && i < r {
			i++
		}
		for nums[j] > pivot && j > l {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
		j--
	}

	nums[l], nums[r] = nums[r], nums[l]
	qSortInMem(nums, l, j-1)
	qSortInMem(nums, j+1, r)
}

func qSortInPartition()
