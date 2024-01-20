package main

import "fmt"

func main() {
	//kiểu string có thể bảo string trình biên dịch cũng tự hiểu là string
	var a string = "Dao Hai"
	fmt.Println(a)
	//var có thể khai báo 1 hoặc nhiều biến
	var b, c int = 1, 2
	fmt.Println(b, c)
	//boolean
	var d = true
	fmt.Println(d)
	// khai bao kiểu số nguyên chưa gán giá trị mặc định
	var e int
	fmt.Println(e)
	// Cú pháp := là cách viết tắt để khai báo và khởi tạo một biến
	f := "Vip pro"
	fmt.Println(f)
}
