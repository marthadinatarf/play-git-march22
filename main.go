package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Print("hello dunia\n")
	fmt.Println("hello")
	fmt.Printf("%s | %d\n", "ini nomor", 1)

	nama := fmt.Sprint("Ucup", 123)
	fmt.Println(nama)

	var namaku string = "ucok"
	fmt.Println("hallo nama saya", namaku)

	fmt.Println("==============================")
	//multiple declare

	var nama1, nama2 string = "bambang", "esan"
	day, angka2 := "senin", 2
	fmt.Println(nama1, nama2, day, angka2)
	fmt.Println("==============================")

	var data1 int = 50
	data2 := strconv.Itoa(data1)
	fmt.Println(reflect.TypeOf(data2))
	//aritmatika
	//data1 = data1 + 25
	//fmt.Println(data1 + 10)
	//data1 += 10
	//fmt.Println(data1)
	data1 = data1 + 25
	fmt.Println("update :", data1+10)

	fmt.Println("==============================")

	nama3 := fmt.Sprint(nama2, " punya uang ", 50000)
	fmt.Println(nama3)

	fmt.Println("==============================")

}
