// GOLANG PROGRAM TO PRINT DIAMOND STAR PATTERN
package main

// fmt package provides the function to print anything
import "fmt"

// start the function main ()
// this function is the entry point of the executable program
func main() {

   // Declare the integer variables
   var x, y, z, rows int
   
   // initialize the rows variable
   rows = 5
   
   // Use of For Loop for the upper half
   for x = 1; x <= rows; x++ {
      // This loop starts when y = 1
      // executes till y<=rows condition is true
      // post statement is y++
      for y = 1; y <= rows-x; y++ {
         fmt.Printf(" ")
      }
      for z = 1; z <= x*2-1; z++ {
         fmt.Printf("*")
      }
      fmt.Println()
   }
   // Use of For Loop for the lower half
   for x = rows - 1; x > 0; x-- {
   
      // This loop starts when y = 1
      // executes till y<=rows condition is true
      // post statement is y++
      for y = 1; y <= rows-x; y++ {
         fmt.Printf(" ")
      }
      for z = 1; z <= x*2-1; z++ {
         fmt.Printf("*")
      }
      fmt.Println()
      // print the result using fmt.Println () function
   }
}
