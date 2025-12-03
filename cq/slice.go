p//ackage main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	size := len(a)

	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	end := size

	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	if start < 0 {
		start = size + start
	}
	if end < 0 {
		end = size + end
	}

	if start < 0 {
		start = 0
	}
	if start > size {
		start = size
	}
	if end < 0 {
		end = 0
	}
	if end > size {
		end = size
	}

	if start > end {
		return nil
	}

	return a[start:end]
}

func main() {

	a := []string{"coding", "algorithm", "ascii", "package", "golang"}

	fmt.Printf("%#v\n", Slice(a, 1))

	fmt.Printf("%#v\n", Slice(a, 2, 4))

	fmt.Printf("%#v\n", Slice(a, -3))

	fmt.Printf("%#v\n", Slice(a, -2, -1))

	fmt.Printf("%#v\n", Slice(a, 2, 0))

}
