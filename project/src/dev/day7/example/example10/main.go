package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path      string
	op        string
	creatTime string
	message   string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s op=%s createTime=%s message=%s", p.path, p.op,
		p.creatTime, p.message)
}

func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:      filename,
			op:        "read",
			message:   err.Error(),
			creatTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}

func main() {
	err := Open("C:/asdfasdfasdf.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error, ", v)
	default:

	}
}
