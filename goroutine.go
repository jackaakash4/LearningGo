package main 

import ( 
    "fmt"
    "time"
)

func printNumbers(){
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        
        time.Sleep( 500 * time.Millisecond)
        fmt.Println("Sleeping time", (500 * time.Millisecond))

    }
}

func main(){
    go printNumbers()
    //started a go routine
    time.Sleep( 3 * time.Second)
    
    fmt.Println("Main function")
    
}

    
