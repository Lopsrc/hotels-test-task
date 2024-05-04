package tasks

func FindPrimes(min, max int) []int{
	var res []int
    for i := min; i <= max; i++{
        if isPrime(i) {
            res = append(res, i)
        }
    }
    return res
}

func isPrime(n int) bool{
	for i := 2; i < n; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}