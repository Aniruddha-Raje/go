package main
  
import "fmt"
  
func main() {
    var value int = 4
      
   switch {
       case value == 1:
       fmt.Println("One")
       case value == 2:
       fmt.Println("Two")
       case value == 3:
       fmt.Println("Three")
       default: 
       fmt.Println("Invalid")
   }
  
}