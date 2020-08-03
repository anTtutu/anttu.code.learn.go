package main

import "fmt"

// 1、定义一个结构体 demo1
type Person struct {
	name string
	city string
	age  int8
}

// 2、定义一个结构体 demo2
type Person1 struct {
	name, city string
	age        int8
}


/**
 * @Author: Anttu
 * @Date: 2020-02-12 09:47
 */
func main() {
	// 3、结构体实例化
	var p1 Person
	p1.name = "张三"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={张三 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"张三", city:"北京", age:18}

	// 4、匿名结构体
	var user struct { Name string; Age int }
	user.Name = "李四"
	user.Age = 20
	fmt.Printf("%#v\n", user)

	// 5、创建指针类型的结构体  我们通过使用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(Person)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}

	// 6、给结构体属性赋值    对结构体指针直接使用.来访问结构体的成员
	var p3 = new(Person)
	p3.name = "小王子"
	p3.age = 28
	p3.city = "上海"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"小王子", city:"上海", age:28}

	// 7、取结构体的地址实例化    使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	p4 := &Person{}
	fmt.Printf("%T\n", p4)     //*main.person
	fmt.Printf("p4=%#v\n", p4) //p4=&main.person{name:"", city:"", age:0}
	p4.name = "赵六"
	p4.age = 30
	p4.city = "成都"
	fmt.Printf("p4=%#v\n", p4) //p4=&main.person{name:"赵六", city:"成都", age:30}

	// 8、结构体初始化
	var p5 Person
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"", city:"", age:0}

	// 9、使用键值对初始化结构体
	p6 := Person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=main.person{name:"小王子", city:"北京", age:18}

	// 10、也可以对结构体指针进行键值对初始化
	p7 := &Person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"小王子", city:"北京", age:18}

	// 11、当某些字段没有初始值的时候，该字段可以不写
	p8 := &Person{
		city: "北京",
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"", city:"北京", age:0}

	// 12、使用值的列表初始化结构体
	// 需要注意：
	//必须初始化结构体的所有字段。
	//初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	//该方式不能和键值初始化方式混用。
	p9 := &Person{
		"王校长",
		"北京",
		28,
	}
	fmt.Printf("p9=%#v\n", p9) //p9=&main.person{name:"王校长", city:"北京", age:28}

	// 13、结构体内存布局探讨   结构体占用一块连续的内存
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}
