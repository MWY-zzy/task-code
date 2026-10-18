package main

import "fmt"

const thisyear = 2026

func main() {
	fmt.Println("你好，GO！")
	fmt.Println("这是我建立的第一个GO程序，它会先打印文字，再和你聊两句。")

	var name string
	fmt.Print("输入你的名字:")
	fmt.Scanln(&name)

	var birthyear int
	fmt.Print("请输入你的出生年月（比如2007）：")
	fmt.Scanln(&birthyear)

	var city string
	fmt.Print("输入你来自的城市：")
	fmt.Scanln(&city)

	age := thisyear - birthyear

	fmt.Printf("你好，%s\n", name)
	fmt.Printf("%d 出生的来自 %s 的你，今年 %d 岁了。\n", birthyear,city, age)
}
