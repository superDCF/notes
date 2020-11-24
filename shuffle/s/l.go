package shuffle
import (
	"math/rand"
	"time"
)

type Solution struct {
    original []int
    target []int
}


func Constructor(nums []int) Solution {
    return Solution{
        original: nums,
        target: nums,
    }
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    this.target = this.original
    return this.target
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.target); i++ {
		index := random(i,len(this.target))
		tem := this.target[i]
		this.target[i] = this.target[index]
		this.target[index] = tem
	}
    return this.target
}

func random(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(b-a) + a
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */