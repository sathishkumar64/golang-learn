package main

import "fmt"

func main() {

	/*	for i, j :=0,0 ;i<5 ;i,j =i+1,j+1{
			fmt.Println(i,j)
		}
	*/
	/*i := 0
	for i < 5  {
		fmt.Println(i)
		i++
	}
	*/
	/*i := 0
	for   {
		fmt.Println(i)
		i++
		if i==5{
			break
		}
	}*/

	/*Loop:
	 for i < 3{
		fmt.Println(i)
		 for j:=1;j <= 3 ;j++ {
			 if i * j >=5{
				 fmt.Println(i,j)
				break Loop
			 }
		 }
		 i++
	 }*/

	// working with collections :=  it will work sclies ,array & map .
	//	s:= []int {1,2,3}

	s := "Hello GO!" // it will work string
	for k, v := range s {
		fmt.Printf(k, string(v))
	}

	studentsDetails := make(map[string]int)
	studentsDetails = map[string]int{
		"Anbu":    24,
		"Babu":    32,
		"Chandru": 38,
	}
	for _, v := range studentsDetails {
		fmt.Println(v)
	}

	for k, v := range studentsDetails {
		fmt.Println(k, v)
	}
}
