
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	result := make([]int, len(nums1)+len(nums2))

	var i int
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] < nums2[0] {
			result[i] = nums1[0]
			nums1 = nums1[1:]
		} else {
			result[i] = nums2[0]
			nums2 = nums2[1:]
		}
		i++
	}

	for _, el := range nums1 {
		result[i] = el
		i++
	}
	for _, el := range nums2 {
		result[i] = el
		i++
	}
	
	mid := float64(len(result)) / 2.0
	
	if len(result) % 2 == 0 {
		a := float64(result[int(mid-1)])
		b := float64(result[int(mid)])
		median := (a + b) / 2.0
		return median
	}

	return float64(result[int(mid)])
}
