// To run this code
// $ cat example.txt | go run find-duplicates.go

package main

import(
    "fmt"
    "os"
    "bufio"
)

func main() {
    // 19: make is a built in function that creates a new, empty map
    //     A map is a data type that stores key/value pairs
    // 20: reads a file from standard input (using cat)
    // 22: input.Scan() scans each line in the file
    // 24: input.Text() retreives the line of text being scanned currently
    
    counts := make(map[string]int)
    input  := bufio.NewScanner(os.Stdin)
    
    for input.Scan() {
        // counts[line] = counts[line] + 1
        line := input.Text()
        counts[line]++
    }
    
    for line, n := range counts {
        fmt.Printf("%d %s \n", n, line)
    }
    
    if err := input.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}