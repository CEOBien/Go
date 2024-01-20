package main

//math là thư viện toán học
import (
	"fmt"
	"math"
)

const s string = "constant"

// kiểu dữ liệu
/*
KIỂU	GIỚI HẠN
uint8	0 – 255
uint16	0 – 65535
uint32	0 – 4294967295
uint64	0 – 18446744073709551615
int8	-128 – 127
int16	-32768 – 32767
int32	-2147483648 – 2147483647
int64	-9223372036854775808 – 9223372036854775807
*/
func main() {
	//const khai báo một giá trị không đổi
	//Câu lệnh const có thể xuất hiện ở bất cứ đâu mà câu lệnh var có thể.
	fmt.Println(s)
	const n = 50000000
	const d = 3e+20 / n
	//Một hằng số không có kiểu cho đến khi nó được cung cấp một kiểu,
	//chẳng hạn như bằng một chuyển đổi rõ ràng.
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	var num1 int16 = 20132
	var num2 int16 = 23244
	//ket qua -22160
	//bởi vì nó tràn bộ nhớ kiểu int16 chỉ giới hạn tới 32767
	fmt.Println("Tong 2 so num1 và num2 là:", num1+num2)
}
