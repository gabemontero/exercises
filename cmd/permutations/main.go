package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type void struct{}
var val void

func main() {
	stringsToPermutate := []string{"a", "b", "c", "ab", "ac", "bc", "abc", "bac", "acb", "bca", "cab", "cba", "abcd", "efju"}
	bufStdOut := bufio.NewWriter(os.Stdout)
	defer bufStdOut.Flush()
	fmt.Fprintln(bufStdOut)
	fmt.Fprintln(bufStdOut)
	for _, str := range stringsToPermutate {
		fmt.Fprintf(bufStdOut, "handling case: %s\n", str)

		list := map[string]void{}
		permutate2(str, 0, list)
		fmt.Fprintf(bufStdOut, "permutations: %v", reflect.ValueOf(list).MapKeys())
		fmt.Fprintln(bufStdOut)
		

		Perm([]rune(str), func(a []rune) {
			fmt.Fprintf(bufStdOut,"permutation: %s\n", string(a))
		})
		fmt.Fprintln(bufStdOut)

	}
}

func permutate2(str string, pos int, list map[string]void) map[string]void {
	if pos > len(str) {
		list[str] = val
		return list
	}
	permutate2(str, pos+1, list)
	for i := pos+1; i < len(str); i++ {
		stringAsArray := strings.Split(str, "")
		stringAsArray[pos], stringAsArray[i] = stringAsArray[i], stringAsArray[pos]
		swappedStr := strings.Join(stringAsArray, "")
		permutate2(swappedStr, pos+1, list)
	}

	return list
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) || len(a) == 1 {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}