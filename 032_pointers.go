package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i, p는 포인터 변수
	fmt.Println(*p) // read i through the pointer, *p는 p가 가리키는 값
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i, i는 21로 변경됨

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	p = &i
	fmt.Println(*p)
	p = &j
	fmt.Println(*p)
}
