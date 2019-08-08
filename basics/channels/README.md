#### Here channels can send and recive.

go func(){
		i := <- ch
		fmt.Println(i)
		ch <- 64
		wg.Done()
	}()
	go func(){	
		ch <- 42
		fmt.Println(<- ch)		
		wg.Done()
	}()
#### so we have to restirct read and write channels
go func(ch <- chan int){
		i := <- ch
		fmt.Println(i)
		ch <- 64
		wg.Done()
	}(ch)
	go func(ch chan <- int){	
		ch <- 42
		fmt.Println(<- ch)		
		wg.Done()
	}(ch)