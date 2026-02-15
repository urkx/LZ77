package lz77

import (
	"os"
	"io"
	"bufio"
	"bytes"
)

func WriteResultFile(name string, data []Result) error {
	buff := make([]byte, 0)
	for _, d := range data {
		for _, dd := range d.Byte() {
			buff = append(buff, dd)
		}
	}
	err := os.WriteFile(name, buff, 0666)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(name string) ([]byte, error) {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(f)
	var res bytes.Buffer
	buffer := make([]byte, 1024)
	for {
		n, err := r.Read(buffer)
	    if err != nil && err != io.EOF {
	        return nil, err
	    }
	    if n == 0 {
	        break
	    }
	    res.Write(buffer[:n])
	}
	return res.Bytes(), nil
}
