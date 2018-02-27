package main

import "fmt"

type Product struct {
	id    string
	name  string
	price int
}

func main() {
	fmt.Println("--- 8 STRUCTS ---")

	var sqp = Product{"sqp", "Square Print", 19}
	fmt.Println(sqp)
	fmt.Println(fmt.Sprintf("%s costs %d$", sqp.name, sqp.price))

	sqp.price = 20
	fmt.Println(sqp)

	var p *Product
	p = &sqp

	p.price += 40

	fmt.Println(p)
	fmt.Println(*p)

	art := Product{id: "artprint", name: "Art Print"}
	panorama := Product{name: "Panorama Print", price: 22}

	fmt.Println(art)
	fmt.Println(panorama)
	panorama.id = "panorama-print"
	fmt.Println(panorama)

	pointer := &Product{id: "temp", name: "Temporary Product"}
	fmt.Println(pointer)

	var products [5]Product
	products[0] = sqp
	products[1] = art
	products[2] = panorama

	for index := 0; index < 5; index++ {
		fmt.Println(products[index])
	}

	names := [5]string{"Ales", "Petr", "Ala", "Terka", "Mira"}
	fmt.Println(names)

	var newNames []string = names[2:5]
	fmt.Println(newNames)

	a := names[0:3]
	b := names[2:5]
	fmt.Println(a, b)

	names[2] = "XXX"
	fmt.Println(a, b)

	fmt.Println(newNames)

	var graph = []struct {
		up    bool
		value int
	}{
		{true, 12},
		{false, 10},
		{true, 8},
		{up: false, value: 20},
		{false, 19},
	}
	fmt.Println(graph)

	var subjects = []string{
		"UMTE",
		"ZMAT",
		"OA2",
	}
	fmt.Println(subjects)

	fmt.Println(subjects[:2])
	fmt.Println(subjects[2:])
	fmt.Println(subjects[:])

	fmt.Println(fmt.Sprintf("Number of subjects: %d", len(subjects)))
	fmt.Println(fmt.Sprintf("Capacity of product array: %d", cap(products)))

	var students = []struct {
		name string
		age  int
	}{
		{name: "Ales", age: 20},
	}
	fmt.Println(students, len(students), cap(students))
	if students == nil {
		fmt.Println("NIL!")
	} else {
		fmt.Println("NOT NIL!")
	}

	var arr = make([]string, 5, 20)
	fmt.Println(arr, len(arr), cap(arr))

	var mat = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	fmt.Println("MAT!")
	mat[1][1] = 0
	printMat(mat)

	fmt.Println("! NEW MAT !")
	printMat(append(mat, []int{0, 0, 0, 0}))

	for i, v := range mat {
		fmt.Println(fmt.Sprintf("%d. %v %v", i, v, mat[i]))
	}
}

func printMat(mat [][]int) {
	for x := 0; x < len(mat); x++ {
		fmt.Println(mat[x])
	}
}
