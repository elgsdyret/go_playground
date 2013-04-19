package main

import (
    "io"
    "os"
    "strings"
)

func Rot13(b byte) byte {	
	switch {
        case 'A' <= b && b <= 'M':
                b = (b - 'A')+'N'
        case 'N' <= b && b <= 'Z':
                b = (b - 'N')+'A'
        case 'a' <= b && b <= 'm':
                b = (b - 'a')+'n'
        case 'n' <= b && b <= 'z':
                b = (b - 'n')+'a'
    }  
	return b
}

type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(data []byte) (n int, err error) {
	bytesRead, err := r.r.Read(data)
	
    for i := 0; i < bytesRead; i++ {
        data[i] = Rot13(data[i])   
    }

	return bytesRead, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
