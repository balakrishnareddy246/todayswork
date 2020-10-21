package BinarySearch
 
import (
	"math/rand"
	"testing"
	"time"
)
 
func BinarySearch(array []int, number int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex := int((maxIndex + minIndex) / 2)
		midItem := array[midIndex]
		if number == midItem {
			return midIndex
		}
		if midItem < number {
			minIndex = midIndex + 1
		} else if midItem > number {
			maxIndex = midIndex - 1
		}
	}
	return -1
}
 
func SortArray(array []int) {
	for itemIndex, itemValue := range array {
		for itemIndex != 0 && array[itemIndex-1] > itemValue {
			array[itemIndex] = array[itemIndex-1]
			itemIndex -= 1
		}
		array[itemIndex] = itemValue
	}
}
 
func TestBinarySearch(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := make([]int, random.Intn(100-10)+10)
	for i := range array {
		array[i] = random.Intn(100)
	}
	SortArray(array)
	for _, value := range array {
		result := BinarySearch(array, value)
		if result == -1 {
			t.Fail()
		}
	}
}