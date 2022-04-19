// func function_name(Parameter-list)(Return_type){
//     // function body.....
// }

package main
import "fmt"

func calculate(a, b int) (int, int){
    return a + b, a - b
}

func calc(a, b int) (x, y int){
    x = a * b
	y = a / b
	return
}
 
func main() {
   
	res1, res2 := calculate(10, 5)
   	fmt.Printf("Sum = %d\nDiff = %d\n", res1, res2)

	res3, res4 := calc(15, 3)
   	fmt.Printf("Mul = %d\nDiv = %d\n", res3, res4)
}