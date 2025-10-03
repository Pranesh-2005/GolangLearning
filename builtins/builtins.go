package main
import "fmt"
import "math"
func main(){
	num := -12
	abs := math.Abs(float64(num))
	fmt.Println("Abs of ",num," is ", abs)
	round := math.Round(52.64)
	fmt.Println("Round of 52.64 is ", round)
	cbrt := math.Cbrt(27)
	fmt.Println("Cube root of 27 is ", cbrt)
	mx := math.Max(5,2)
	fmt.Println("Max of 5 and 2 is ", mx)
	min := math.Min(6,4)
	fmt.Println("Min of 6 and 4 is ", min)

}