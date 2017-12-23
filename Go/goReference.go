package main // needs to be included where your main function is

/* * * * * * * * * * * * * * * * * * * * * * * * *
 * Go References:
 * https://gobyreference.com
 * https://gobyexample.com
 * https://github.com/ardanlabs/gotraining
 *
 * * * * * * * * * * * * * * * * * * * * * * * * */



import (
    "fmt" // for printing to standard out
    "runtime" // for Concurrency
    "sync"
    "time"
    "math/rand"
)


// This is our entrypoint
func main() {
    // variadic function example
    n := average(10, 20, 30, 40, 50, 60)
    fmt.Println("variadic function average: ", n)

    // variadic arguments with slice
    data := []float64{1, 2, 3, 4, 5} // <-- this right here is a slice, denoted with [] before array
    // elipses indicate variadic arguments
    // this will pull each element from data and run it through
    // our avg function
    p := avg(data...)
    fmt.Println("Variadic argument average: ", p)

    // anonymous function
    notAnonymous()

    // returning a function: returns a hex value
    fmt.Print("Print address of function:  ")
    fmt.Println(returnFunction())
    t := returnFunction()
    //  this will print the actual value returned
    fmt.Print("Print return value of func: ")
    fmt.Println(t())

    // pass functions as arguments
    fmt.Println("Passing. . . ")
    passFunc([]int{1,2,3,4,5}, func(n int) {
        fmt.Println("n is: ", n)
    })

    // example of deferring
    toDefer()

    // Slice examples
    slicing()
    makeSlices()

    // examples of concurrency
    Concurrency()
    mutexing()
    // example of channel
    goChannel()
    goChannel1()
}

// variadic function returning float64
func average (array ...float64) float64 {
    var total float64 // automatically assigns total = 0
    for _, num := range array { // range iterates for all values in array
        total += num
    }
    return total / float64(len(array))
}

// just to demonstrate a slice
func avg (array ...float64) float64 {
    var total float64
    for _, num := range array {
        total += num
    }
    return total / float64(len(array))
}

// playing with anonymous functions
func notAnonymous() {
    // func expression
    anonymous := func() {
        // very javascript, no?
        fmt.Println("func expression:\t    Such anonymous")
    }
    anonymous()
}

// returning a function
func returnFunction() func() string {
    return func() string {
        return "Function returned!"
    }
}

                    // slice of int
func passFunc(numbers []int, callback func(int)) {
    for _, n := range numbers {
        callback(n)
    }
}

// example of defer
func toDefer() {
    t := func() {
        fmt.Println("defer me!")
    }
    defer t() // this will happen right before toDefer() exits
    fmt.Println("Me first!")
    // An example where this might be useful is you can open a file,
    // and then on the next line defer closing the file. This way
    // the file will close after the function ends.
}

// examples of slices and slicing
func slicing() {
    fmt.Println("Initiate Slicing: ")
    mySlice := []int{1, 9, 4, 22, 67, 23}
    fmt.Println(mySlice)
    fmt.Println(mySlice[2:4])
    fmt.Println(mySlice[2])
    // casting because it'll print the ASCII value
    fmt.Println(string("myString"[2]))
}

func makeSlices() {
    fmt.Println("Lets get slicing: ")
    // make int slice with 0 len, capacity of 5
    mySlice := make([]int, 0, 5)
    fmt.Println("leng is: ", len(mySlice))
    fmt.Println("capacitiy is: ", cap(mySlice))
    fmt.Println("--------------------")
    fmt.Println("demonstrate dynamic resizing: ")
    for i := 0; i < 30; i++ {
        mySlice = append(mySlice, i)
        fmt.Println("len is ", len(mySlice), "cap is: ", cap(mySlice))
    }
    // Go doubles the capacity of the slice once the number of elements
    // in the slice exceed its capacity
}

// Concurrency examples
// ** note ** Concurrency != Parallel. Parallel = Concurrency
var waitGroup sync.WaitGroup

// this is gonna happen at the beginning
func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func Concurrency() {
    waitGroup.Add(2)
    go printOdd()
    go printEven()
    waitGroup.Wait()
}

func printOdd() {
    for i := 1; i < 10; i++ {
        fmt.Print("Odd:", i, " ")
        time.Sleep(time.Duration(3 * time.Millisecond))
    }
    waitGroup.Done()
}

func printEven() {
    for i := 0; i < 10; i++ {
        fmt.Print("Even:", i, " ")
        time.Sleep(time.Duration(3 * time.Millisecond))
    }
    waitGroup.Done()
}
// end concurrency example

// begin Mutex example
var mutex sync.Mutex
var counter int
func mutexing() {
    waitGroup.Add(2)
    go incrementor("one")
    go incrementor("two")
    waitGroup.Wait()
}

func incrementor(input string) {
    for i := 0; i < 10; i++ {
        time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
        mutex.Lock()
        x := counter
        x++
        counter = x
        fmt.Println(input, i, "counter:", counter)
        mutex.Unlock()
    }
    waitGroup.Done()
}
// end mutex example

// begin channel examples
func goChannel() {
    fmt.Println("Channel")
    // make the channel (unbuffered channel)
    c := make(chan int)
    go func(){
        // put things into the channel
        for i :=0; i < 10; i++ {
            fmt.Println("Putting on")
            c <- i
        }
        // close(c)     Once you do this, no more values can be put into the channel.
        //              Values can be taken out, but they cannot be added.
    }()
    go func(){
        for {
        // take things out of the channel
            fmt.Println(<-c)
        }
    }()
    time.Sleep(time.Second)
}

func goChannel1() {
    fmt.Println("Channel1")
    c := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        close(c)
    }()

    for n := range c {
        fmt.Println(n)
    }
}
