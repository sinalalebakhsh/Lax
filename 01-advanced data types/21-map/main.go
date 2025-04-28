package main


func main()  {
	var ages = make(map[string]int)
	// خوانده میشود مپ استرینگ به اینت
	another_map := map[string]int{}



	// add value
	ages["sina"] = 100
	another_map["lina"] = 150

	

	// an error 
	// panic: assignment to entry in nil map
	var another_one map[string]int  
	another_one["something"] = 10

	
}