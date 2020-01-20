# String slicer

## Usage
`$ go get github.com/lenniDespero/otus-golang/hw10`

    package main

	import (
    	"fmt"
    
    	"github.com/lenniDespero/otus-golang/hw8"
    )

	func main() {
			from := flag.String("from", "", "limit bytes to write")
            to := flag.String("to", "", "offset bytes to write")
            limit := flag.Int64("limit", -1, "limit bytes to write")
            offset := flag.Int64("offset", 0, "offset bytes to write")
            flag.Parse()
            if *from == "" || *to == "" {
            	fmt.Println("Nothing to copy. Use -h to help.")
            	os.Exit(1)
            }
            err := hw10.Copy(*from, *to, *limit, *offset)
            if err != nil {
            	fmt.Printf("An error occurred while file copying: %s\n", err.Error())
            	os.Exit(1)
            }
            fmt.Println("File copied successfully")
	}

Output will be like: 

    Processing... [104857600 / 104857600] 100.00%
    File copied successfully


For test run `$ go test [path_to_module]`  

