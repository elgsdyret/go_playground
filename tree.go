package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
    "sort"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {	
    if t == nil {
		return     	   
    }
    
    ch <- t.Value
    Walk(t.Left, ch)
    Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {    
    ch1 := make(chan int)
    ch2 := make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)

    s1 := make([]int, 10)
    s2 := make([]int, 10)

    for i := 0; i < 10; i++ {
        a := <- ch1
        b := <- ch2
        s1[i] = a
        s2[i] = b
    }    
    sort.Ints(s1)
    sort.Ints(s2)

    for i := 0; i < 10; i++ {
        if (s1[i] != s2[i]) {
            return false
        }
    }
    return true;
}

func main() {
    fmt.Println(Same(tree.New(10), tree.New(10))) 
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
