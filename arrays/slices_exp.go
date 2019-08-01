package main

import "fmt"

func main (){

	a:= [] int {1,2,3,4}

	fmt.Println(a)
	fmt.Printf("Lenght %v\n",len(a))
	fmt.Printf("Capacity %v \n",cap(a))

	b:= a
	b[1]=6

	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("Lenght %v\n",len(a))
	fmt.Printf("Capacity %v \n",cap(a)) //If will change any value reference array, also will be update source array

	//Creating different ways 

	//aa :=[] int {1,2,3,4,5,6,7,8,9,0}
	//aa :=[...] int {1,2,3,4,5,6,7,8,9,0}  ///aslo work as array
	//bb := aa[:]
	//c := aa[:3]
//	d := aa[4:]
//	e := aa[1:5]

	//fmt.Printf("a %v\n",aa)
	//fmt.Printf("bb %v\n",bb)
	//fmt.Printf("c %v\n",c)
	//fmt.Printf("d %v\n",d)
	//fmt.Printf("e %v\n",e)


	//Make predefiend oper

	f := make([]int, 3,100) // slice does nt have fixed array size

	fmt.Println(f)
	fmt.Printf("Lenght: %v\n",len(f))
	fmt.Printf("Capacity: %v\n",cap(f))

	g := [] int{}

	fmt.Println(g)
	fmt.Printf("Lenght g : %v\n",len(g))
	fmt.Printf("Capacity g: %v\n",cap(g))

	g =append(g,1)

	fmt.Println(g)
	fmt.Printf("Lenght g : %v\n",len(g))
	fmt.Printf("Capacity g: %v\n",cap(g))

	//g =append(g,1,2,3,4)

	//fmt.Println(g)
	//fmt.Printf("Lenght g : %v\n",len(g))
	//fmt.Printf("Capacity g: %v\n",cap(g))

	//want to add one slice to another

	//g =append(g,[] int {1,2,3,4}) // It wont work
	g =append(g,[] int {1,2,3,4}...)  // ==g =append(g,1,2,3,4)

	fmt.Println(g)
	fmt.Printf("Lenght g : %v\n",len(g))
	fmt.Printf("Capacity g: %v\n",cap(g))


}