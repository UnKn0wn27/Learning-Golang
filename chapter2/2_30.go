package main

import "fmt"

func doPanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recover  with: ", e)
		}
	}()
	panic("Just panicking for the sake of demo")
	fmt.Println("This will never be called")
}

func main()  {
	fmt.Println("Starting to panic")
	doPanic()
	fmt.Println("Program regains control after panic recover.")
}
