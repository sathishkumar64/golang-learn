package main 

import "fmt"

func main (){

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
	i := 0
	/*for   {
		fmt.Println(i)
		i++
		if i==5{
			break
		}
	}*/

	j := 1
	Loop:
	 for i < 3{
		fmt.Println(i)
		 for j:=1;j <= i ;j++ {
			 if i * j >=3{
				 fmt.Println(i,j)
				 break 
			 }
		 } 
		 i++
	 }
}

