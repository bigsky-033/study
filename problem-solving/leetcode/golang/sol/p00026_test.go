package sol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP00026_01(t *testing.T) {
	nums := []int{1, 1, 2}
	expectedNums := []int{1, 2}

	k := p00026_removeDuplicates(nums)

	assert.Equal(t, len(expectedNums), k)
	for i := 0; i < k; i++ {
		assert.Equal(t, expectedNums[i], nums[i])
	}
}

func TestP00026_02(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedNums := []int{0, 1, 2, 3, 4}

	k := p00026_removeDuplicates(nums)

	assert.Equal(t, len(expectedNums), k)
	for i := 0; i < k; i++ {
		assert.Equal(t, expectedNums[i], nums[i])
	}
}

func TestP00026_03(t *testing.T) {
	nums := []int{-1, 0, 0, 0, 0, 3, 3}
	expectedNums := []int{-1, 0, 3}

	k := p00026_removeDuplicates(nums)

	assert.Equal(t, len(expectedNums), k)
	for i := 0; i < k; i++ {
		assert.Equal(t, expectedNums[i], nums[i])
	}
}
