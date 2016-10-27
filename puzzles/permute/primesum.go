package main

import (
	"fmt"
	"sync"
)

var (
	// N is less than 250. And 1<= k <=5
	// We can compute a table of primes < 250
	// Computed using a simple program.
	primes = []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241,
	}
)

func primeSum(n int, k int) bool {
	// Naively we can linear scan to find the candidate set of primes to permute
	largestPrimeIdx := 0
	for i, p := range primes {
		// The smallest prime is 2. So we know that the smallest
		// sum we can have is (k-1)*2 + p.
		if (k-1)*2+p > n {
			largestPrimeIdx = i - 1
			break
		}
	}
	if largestPrimeIdx < 0 {
		return false
	}

	// Now.. we must find all permutations of length k using the primes we
	// figured out above, whose sum append is equal to n.
	for _, permutationOfLengthK := range permuteK(primes[:largestPrimeIdx+1], k) {
		if sum(permutationOfLengthK) == n {
			fmt.Println(permutationOfLengthK)
			return true
		}
	}
	return false
}

func sum(vals []int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}

// Returns a channel that will send permutation of the numbers in vals that are of size K.
// Allowing for repetitions.
func permuteK(vals []int, k int) [][]int {
	if k == 0 {
		return [][]int{nil}
	}
	if len(vals) == 0 {
		return nil
	}
	perm := permuteK(vals[1:], k)
	for _, subPerm := range permuteK(vals, k-1) {
		perm = append(perm, append(subPerm, vals[0]))
	}
	return perm
}

func main() {
	//   fmt.Println(primeSum(22, 2))
	//   fmt.Println(primeSum(5, 1))
	wg := &sync.WaitGroup{}
	for n := 1; n <= 250; n++ {
		for k := 1; k <= 5; k++ {
			wg.Add(1)
			go func(n, k int) {
				fmt.Println("Computed n", n, " k", k, primeSum(n, k))
				wg.Done()
			}(n, k)
		}
	}
	wg.Wait()
}
