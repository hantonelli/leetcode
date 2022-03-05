package insertdeletegetrandom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	randomizedSet := Constructor()
	require.Equal(t, true, randomizedSet.Insert(1))  // Inserts 1 to the set. Returns true as 1 was inserted successfully.
	require.Equal(t, false, randomizedSet.Remove(2)) // Returns false as 2 does not exist in the set.
	require.Equal(t, true, randomizedSet.Insert(2))  // Inserts 2 to the set, returns true. Set now contains [1,2].
	randomizedSet.GetRandom()                        // getRandom() should return either 1 or 2 randomly.
	require.Equal(t, true, randomizedSet.Remove(1))  // Removes 1 from the set, returns true. Set now contains [2].
	require.Equal(t, false, randomizedSet.Insert(2)) // 2 was already in the set, so return false.
	require.Equal(t, 2, randomizedSet.GetRandom())   // Since 2 is the only number in the set, getRandom() will always return 2.
}

func Test2(t *testing.T) {
	randomizedSet := Constructor()
	randomizedSet.Remove(0)
	randomizedSet.Remove(0)
	randomizedSet.Insert(0)
	randomizedSet.GetRandom()
	randomizedSet.Remove(0)
	randomizedSet.Insert(0)
}

func Test3(t *testing.T) {
	randomizedSet := Constructor()
	randomizedSet.Insert(3)
	randomizedSet.Insert(3)
	randomizedSet.GetRandom()
	randomizedSet.GetRandom()
	randomizedSet.Insert(1)
	randomizedSet.Remove(3)
	randomizedSet.GetRandom()
	randomizedSet.GetRandom()
	randomizedSet.Insert(0)
	randomizedSet.Remove(0)
}
