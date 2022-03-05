package loggerratelimiter

import (
	"testing"
)

func Test(t *testing.T) {
	logger := Constructor()
	logger.ShouldPrintMessage(1, "foo")  // return true, next allowed timestamp for "foo" is 1 + 10 = 11
	logger.ShouldPrintMessage(2, "bar")  // return true, next allowed timestamp for "bar" is 2 + 10 = 12
	logger.ShouldPrintMessage(3, "foo")  // 3 < 11, return false
	logger.ShouldPrintMessage(8, "bar")  // 8 < 12, return false
	logger.ShouldPrintMessage(10, "foo") // 10 < 11, return false
	logger.ShouldPrintMessage(11, "foo") // 11 >= 11, return true, next allowed timestamp for "foo" is 11 + 10 = 21
	logger.Close()
}
