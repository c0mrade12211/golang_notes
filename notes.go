Расс Кокс, один из ведущих разработчиков Go, написал статью о внутреннем представлении интерфейсов и объяснил, что интерфейс состоит из двух слов:

указатель на информацию о сохраненном типе
указатель на данные



package main

func main() {
    var myVar int32 = 1
    printInterface(myVar)
}

//go:noinline
func printInterface(val interface{}) {
    println(val)
}


(0x459dc0,0xc00003476c)





func main() {
	var values []interface{}
	values = append(values, 42, "Hello, World!", true)
	fmt.Printf("Type of value: %T\n", values[0])
	fmt.Printf("Type of value: %T\n", values[1])
	fmt.Printf("Type of value: %T\n", values[2])
	/*
	Type of value: int
	Type of value: string
	Type of value: bool
	*/
}
