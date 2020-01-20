package hw10

import (
	"fmt"
	"io"
	"os"
)

//Copy will copy From (string) file destination To (string) file destination Limit (int) bytes with Offset (int) bytes
func Copy(from string, to string, limit int64, offset int64) error {
	file, err := os.Open(from)
	if err != nil {

		return fmt.Errorf("No file or access denied: %s", err.Error())
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Can't read file information: %s", err.Error())
	}
	var total int64 = info.Size()

	var input io.Reader = file
	if limit != -1 {
		input = io.LimitReader(file, limit)
		total = limit
	}

	if offset > 0 {
		pos, err := file.Seek(offset, 0)
		if err != nil || pos != offset {
			return fmt.Errorf("Can't set offset for reading")
		}
		if offset+total >= info.Size() {
			total = offset + total - info.Size()
		}
	}

	output, err := os.Create(to)
	if err != nil {
		return fmt.Errorf("Can't create file: %s", err.Error())
	}
	defer output.Close()
	defer println()
	for totalWritten := int64(0); totalWritten < total; {
		written, err := io.CopyN(output, input, 1024)
		totalWritten += written
		if err != nil {
			if err == io.EOF {
				printBar(totalWritten, totalWritten)
				break
			}
			return fmt.Errorf("Error: %s", err.Error())
		}
		printBar(totalWritten, total)
	}

	return nil
}

func printBar(written int64, total int64) {
	fmt.Printf(
		"%%\rProcessing... [%d / %d] %.2f",
		written,
		total,
		float64(written)/float64(total)*100,
	)
}
