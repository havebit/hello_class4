package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"
	_ "github.com/pallat/hello_api_class4/effect"
)

func main() {
	fmt.Println(add([]int{1, 2, 3, 4, 5}))

}

func add[T int | float64](s []T) T {
	var sum T

	for _, v := range s {
		sum += v
	}
	return sum
}

func mainGo() {
	start := time.Now()
	ch := make(chan struct{}, 3)

	go slowPrint(ch, "1")
	go slowPrint(ch, "2")
	go slowPrint(ch, "3")

	for i := 0; i < 3; i++ {
		<-ch
	}

	fmt.Println(time.Since(start))
}

func slowPrint(ch chan struct{}, s string) {
	time.Sleep(time.Second)
	fmt.Println(s)
	ch <- struct{}{}
}

type xhandler interface {
	Serve(int)
}

type x struct{}

func (x) Serve(i int) {
	fmt.Println("x.Serve", i)
}

func handleX(fn xhandler) {
	fn.Serve(10)
}

func xFunc(n int) {
	fmt.Println("xFunc", n)
}

type tFunc func(int)

func (fn tFunc) Serve(i int) {
	fn(i)
}

func mainMethodOnFunc() {
	handleX(tFunc(xFunc))
}

func abc() func() int {
	i := 10
	return func() int {
		defer func() { i++ }()

		return i
	}
}

func mainDB() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:mysecretpassword@localhost:5432/myapp")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	if _, err := conn.Exec(context.Background(), "INSERT INTO TODOS(title) VALUES($1)", "Hello db"); err != nil {
		// Handling error, if occur
		fmt.Println("Unable to insert due to: ", err)
		return
	}
	fmt.Println("Insertion Succesfull")
}

type cat struct{}

func (c cat) Color() string {
	return "black"
}

type dog struct{}

func (d dog) Color() string {
	return "white"
}

type color interface {
	Color() string
}

func printAnimalColor(c color) {
	fmt.Println(c.Color())
}

func mainInterface() {
	c := cat{}
	printAnimalColor(c)
	d := dog{}
	printAnimalColor(d)
}

func mainAny() {
	var i any

	i = 10
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = "ten"
	fmt.Printf("type is %T, value is %v\n", i, i)
	if s, ok := i.(string); ok {
		fmt.Printf("type is %T, value is %v\n", s, s)
	}

	i = struct {
		number int
		text   string
	}{
		number: 10,
		text:   "ten",
	}
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = func() string {
		return "10"
	}
	fmt.Printf("type is %T, value is %v\n", i, i)
}

func mainRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("ok")
		}
	}()

	log.Fatal("fatal")
}

func mainDefer() {
	fmt.Println(number(9))
}

func number(n int) int {
	defer func() {
		n += 1
	}()
	return n
}

type rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

var jsonString = `{"width":10,"height":20}`

func mainJSONStruct() {
	var rec1 rectangle
	if err := json.Unmarshal([]byte(jsonString), &rec1); err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v\n", rec1)

}

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

type String string

func (s *String) toUpper() {
	*s = String(strings.ToUpper(string(*s)))
}

func (s String) toInt() int {
	i, _ := strconv.Atoi(string(s))
	return i
}

type my bool

func (my) String() string {
	return "my type"
}

func mainMethod() {
	fmt.Println(my(true))

	// var i Int = 10
	// fmt.Printf("%T %q\n", i, i)

	// var t String = "abc"
	// t.toUpper()
	// fmt.Println(t)
}

func mainCoule() {
	fmt.Println(couple("abcde"))
	fmt.Println(couple("abcdef"))
	fmt.Println(couple("กขคงจ"))

}

func couple(str string) []string {
	var r []string
	s := []rune(str)
	for s = append(s, []rune("*")...); len(s) > 1; s = s[2:] {
		r = append(r, string(s[:2]))
	}
	return r
}

func mainSlice() {
	// var a []string
	// fmt.Println(a == nil)

	// a := make([]string, 1, 2)

	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	a := array[0:5:5]
	a = append(a, 99)
	a[0] = 100

	fmt.Println(array[:])
	fmt.Println(a[:], cap(a), len(a))
}

func mainArray() {
	primes := [...]int{2, 3, 5, 7, 11, 13}

	for _, prime := range primes {
		fmt.Println(prime)
	}

	fmt.Printf("%T len: %d\n", primes, len(primes))

}

var name string = "Pallat"

func mainENV() {
	n := os.Getenv("NAME")
	if n != "" {
		name = n
	}
	fmt.Println("Hello,", name)
}

func power(b, x int) int {
	r := 1
	for i := 0; i < x; i++ {
		r *= b
	}
	return r
}

func printPrimes(n int) {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func mainBasic() {
	name := "John"
	println("Hello,", name)

	fmt.Println("Square area of 3 is", squareArea(3))

	a, b := swap(1, 2)
	fmt.Println("Swap 1 and 2:", a, b)

	if ok := IsCorrect(); ok {
		println("It's correct")
	}

}

func IsCorrect() bool {
	return true
}

func swap(a, b int) (int, int) {
	return b, a
}

func squareArea(a float64) float64 {
	return a * a
}
