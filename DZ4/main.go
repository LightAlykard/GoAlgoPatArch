package main

import (
	"flag"
	"fmt"
	"github.com/guptarohit/asciigraph"
)

func Search(mass []int, chislo int) (result int, count float64) {
	mid := len(mass) / 2
	switch {
	case len(mass) == 0:
		result = -1
	case mass[mid] > chislo:
		result, count = Search(mass[:mid], chislo)
	case mass[mid] < chislo:
		result, count = Search(mass[mid+1:], chislo)
		result += mid + 1
	default:
		result = mid
	}
	count++
	return
}

func main() {
	//Сложность log n

	var num int
	var counts []float64	

	searchArray0 := []int{10, 20, 30, 40, 50}
	searchArray1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	searchArray2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150}
	searchArray3 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200}	
	searchArray4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	flag.IntVar(&num, "n", 0, "Search number")

	flag.Parse()

	fmt.Printf("Поиск числа %d\n", num)

	_, count := Search(searchArray0, num)
	counts = append(counts, count)

	_, count = Search(searchArray1, num)
	counts = append(counts, count)

	_, count = Search(searchArray2, num)
	counts = append(counts, count)

	_, count = Search(searchArray3, num)
	counts = append(counts, count)

	_, count = Search(searchArray4, num)
	counts = append(counts, count)	

	graph := asciigraph.Plot(counts)

	fmt.Println(graph)
}
