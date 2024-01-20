package main

import "fmt"

func main() {
	/*câu lệnh init: được thực thi trước lần lặp đầu tiên
	biểu thức điều kiện:được đánh giá trước mỗi lần lặp
	câu lệnh post: được thực thi ở cuối mỗi lần lặp*/
	// vòng lặp for được hợp nhất từ for và while không có do while
	i := 1
	//giong while
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	//giống for c
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("Dao Hai")
		break
	}
	/* Nếu bạn đang lặp qua một mảng, lát,
	chuỗi hoặc bản đồ hoặc đọc từ một kênh thì mệnh đề phạm vi có thể quản lý vòng lặp.*/
	// Tạo một bản đồ cũ
	oldMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Tạo một bản đồ mới
	newMap := make(map[string]int)

	// Lặp qua mỗi cặp khóa-giá trị trong bản đồ cũ và sao chép sang bản đồ mới
	for key, value := range oldMap {
		newMap[key] = value
	}

	// In ra bản đồ mới
	fmt.Println("New Map:", newMap)

	/*nếu lặp map,mảng hay slice mà chỉ muốn lấy value*/
	numbers := []int{1, 2, 3, 4, 5}

	// Biến để lưu tổng
	sum := 0

	// Lặp qua mỗi giá trị trong mảng và cộng vào biến tổng
	for _, value := range numbers {
		sum += value
	}

	// In ra tổng
	fmt.Println("Tổng các số trong mảng là:", sum)
	//đảo ngược chuỗi golang. Không có toán tử ++ và --
	// Reverse a
	a := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println("reserve ::", a)

}
