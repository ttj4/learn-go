package main
import "fmt"

func main() {
	a := 1
	b := 2

	fmt.Println(a,b)
	swap(&a, &b)
	fmt.Println(a,b)

}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
