package main

import (
	"fmt"
	"time"
)

func main() {
	// switch cơ bản
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	/*Bạn có thể sử dụng dấu phẩy để phân tách nhiều biểu thức trong cùng một câu lệnh case.
	Chúng tôi cũng sử dụng trường hợp mặc định tùy chọn trong ví dụ này.*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	/*
		Việc sử dụng interface{} trong hàm whatAmI là để cho phép hàm này nhận
		và xử lý các giá trị thuộc bất kỳ kiểu dữ liệu cụ thể nào.
		Trong Go, interface{} là một kiểu dữ liệu rỗng có thể chứa giá trị của bất kỳ kiểu dữ liệu nào.*/
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
