package main

import (
	"fmt"
	//"reflect"
	"golang.org/x/tour/tree"
)

/*
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//flatt1 := []int{}
	//flatt2 := []int{}
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		v1 := <-ch1
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
	/*for i := 0; i < 10; i++ {
		flatt1 = append(flatt1, <-ch1)
		flatt2 = append(flatt2, <-ch2)
	}
	return reflect.DeepEqual(flatt1, flatt2)
	*/
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	go Walk(tree.New(2), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}



