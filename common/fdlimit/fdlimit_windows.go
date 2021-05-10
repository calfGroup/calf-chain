
package fdlimit

import "fmt"

// hardlimit is the number of file descriptors allowed at max by the kernel.
const hardlimit = 16384


func Raise(max uint64) (uint64, error) {

	if max > hardlimit {
		return hardlimit, fmt.Errorf("file descriptor limit (%d) reached", hardlimit)
	}
	return max, nil
}

// Current retrieves the number of file descriptors allowed to be opened by this
// process.
func Current() (int, error) {
	// Please see Raise for the reason why we use hard coded 16K as the limit
	return hardlimit, nil
}


func Maximum() (int, error) {
	return Current()
}
