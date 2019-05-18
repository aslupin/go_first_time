package main

import "fmt"

func typeCheck(x interface{}) {
	_, ok := x.(int)
	if ok {
		fmt.Println("OK int")
		return
	}

	_, ok2 := x.(string)
	if ok2 {
		fmt.Println("OK string")
		return
	}
}

type Animal interface {
	Cry() string
}

func Print(x interface{}) {
	fmt.Println(x)
}
func CryThreeTime(a Animal) string {
	return a.Cry() + a.Cry() + a.Cry()
}

type Dog struct {
	Name string
}

func (d Dog) Cry() string {
	return "Hong, My Name is " + d.Name + "!"
}

type Cat struct {
	// letter for private (Can't Export)
	// Capital for public (Can Export)
	Name string `json:"name"`
	Age  int    `json:"-"`
}

func (c Cat) Cry() string {
	return "Meow, My Name is " + c.Name + "!"
}

func (c *Cat) Rename(name string) {
	c.Name = name
}

type Computer struct {
	Age int
	Ip  string
}

func main() {
	bobo := Cat{
		Name: "Bobo",
		Age:  4,
	}

	bobo.Rename("Poon")
	fmt.Println(bobo.Cry())
	fmt.Println(CryThreeTime(bobo))

	Alice := Dog{"MyDog"}
	fmt.Println(Alice.Cry())
	Print(bobo)
	typeCheck(12)
	// DAlice := Dog

	// alice := Cat{"Alice", 2}
	// fmt.Println(alice.Name)
	// fmt.Printf("%+v", bobo)
	// fmt.Println(bobo.Name)

	// // jsonByte, _ := json.Marshal(bobo) // return's second  for error
	// jsonByte, err := json.Marshal(bobo)
	// fmt.Println(string(jsonByte))
	// fmt.Println(err)

	// var newCat Cat
	// // var newCC Computer
	// // if err = json.Unmarshal(jsonByte, &newCat): err != nil
	// // {
	// err = json.Unmarshal(jsonByte, &newCat)
	// fmt.Println(newCat)
	// println(err) // 0x0 is Nil
	// // }

}
