// To run this code, pass the text file as an arg
// $ go run find-dups-args.go example.txt

package main

import(
    "fmt"
    "os"
    "log"
    "bufio"
)

func main() {
    // 20: make is a built in function that creates a new, empty map
    //     A map is a data type that stores key/value pairs
    // 23: reads a file from command line arg
    // 33: input.Scan() scans each line in the file
    // 35: input.Text() retreives the line of text being scanned currently
    
    counts := make(map[string]int)
    
    // Read the file from os
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    
    // Create a buffered read, then a scan
    read := bufio.NewReader(file)
    scan  := bufio.NewScanner(read)
    
    // Iterate over each scanned line, and set k/v's in our map
    for scan.Scan() {
        // counts[line] = counts[line] + 1
        line := scan.Text()
        counts[line]++
    }
    
    for line, num := range counts {
        fmt.Printf("%d %s \n", num, line)
    }
    
}