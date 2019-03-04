package main

import ("fmt"; "math/rand"; "time")

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var a[10]int
	var b[10]int
	var c[10]int
	for i := 0; i < 10; i++ {
	a[i] = rand.Intn(100)
	}
	fmt.Println("Array #1", a)
	for i := 0; i < 10; i++ {
	b[i] = rand.Intn(100)
	}
	fmt.Println("Array #2", b)
	for i := 0; i < 10; i++ {
	if a[i] > b[i] {
        c[i] = a[i] 
	} else { 
	c[i] = b[i]
	}
	} 
	fmt.Println("Array #3", c)
}