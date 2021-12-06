package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	size = 12
)

func count(nums []uint64) []uint64 {
	counters := make([]uint64, size)
	for _, num := range nums {
		for i := 0; i < size; i++ {
			counters[i] += num % 2
			num = num / 2
		}
	}
	return counters
}

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	nums := make([]uint64, 0)
	for fscanner.Scan() {
		num, _ := strconv.ParseUint(fscanner.Text(), 2, 64)
		nums = append(nums, num)
	}

	counters := count(nums)
	gamma := uint64(0)
	epsilon := uint64(0)
	half := uint64(len(nums) / 2)

	for i := len(counters) - 1; i > -1; i-- {
		to_add := uint64(math.Pow(2, float64(i)))
		if counters[i] >= half {
			gamma += to_add
		} else {
			epsilon += to_add
		}
	}

	// P1
	fmt.Println(gamma)
	fmt.Println(epsilon)

	oxygen := nums
	for index := size - 1; len(oxygen) > 1; index-- {
		oxygen_half := uint64((len(oxygen) + 1) / 2)
		counters := count(oxygen)
		bit_mask := uint64(math.Pow(2, float64(index)))
		new_oxygen := make([]uint64, 0)
		for _, num := range oxygen {
			// mask relevant bit and compare to
			if (num & bit_mask) == counters[index]/oxygen_half*bit_mask {
				new_oxygen = append(new_oxygen, num)
			}
		}
		oxygen = new_oxygen
	}

	co2 := nums
	for index := size - 1; len(co2) > 1; index-- {
		co2_half := uint64((len(co2) + 1) / 2)
		counters := count(co2)
		bit_mask := uint64(math.Pow(2, float64(index)))
		new_co2 := make([]uint64, 0)
		for _, num := range co2 {
			// mask relevant bit and compare
			if (num & bit_mask) != counters[index]/co2_half*bit_mask {
				new_co2 = append(new_co2, num)
			}
		}
		co2 = new_co2
	}

	// P2
	fmt.Println(oxygen[0])
	fmt.Println(co2[0])
}
