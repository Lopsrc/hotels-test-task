package tasks

// CalculaeCommonDivisors calculates the common divisors of a list of integers.
func CalculaeCommonDivisors(nums []int) []int{
	var res []int
	for i := 2; i <= min(nums); i++ {
		counter := 0
		for _, v := range nums {
			if v % i == 0 {
				counter++
			}
		}
		if counter == len(nums) {
            res = append(res, i)
        }
	}
	return res
}
// min function calculates the minimum value in a slice of integers.
func min(nums []int) int{
	min := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] < min {
            min = nums[i]
        }
    }
    return min
}