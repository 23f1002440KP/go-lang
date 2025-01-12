package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// f,err := os.Open("example.txt")
	// defer f.Close()
	
	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo,err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File Name: ",fileInfo.Name())
	// fmt.Println("File Or Folder: ",fileInfo.IsDir())
	// fmt.Println("File Size: ",fileInfo.Size())
	// fmt.Println("File Permission: ",fileInfo.Mode())
	// fmt.Println("File Modified at: ",fileInfo.ModTime())
	// fmt.Println("File Some: ",fileInfo)


	// buf := make([]byte, 1024)
	// n, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(buf[:n]),n,buf)


	// f , err:= os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(f))


	// dir , err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo,err := dir.Readdir(-1)

	// if err != nil {
	// 	panic(err)
	// }

	// for _, file := range fileInfo {
	// 	fmt.Println(file.Name())
	// }

	// file,err := os.Create("example2.txt")
	// defer file.Close()

	// if err != nil {
	// 	panic(err)
	// }

	// file.WriteString("Hello World1")
	// file.WriteString("Hello World2")
	// file.Sync()

	// bytes := []byte("Hello fo")
	// file.Write(bytes)



	src, err := os.Open("example.txt")  // Source File
	defer src.Close()
	if err != nil {
		panic(err)
	}

	dst, err := os.Create("example2.txt") // Destination File  // if file not exist then creates  new file automatically
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(src)  // Reader
	writer := bufio.NewWriter(dst)  // Writer

	for {
		read,err:=reader.ReadByte() // Read Byte
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		writer.WriteByte(read) // Write Byte
		fmt.Println(string(read))
	}

	writer.Flush() // Flush Writer
	fmt.Println("File Copied Successfully")
	dst.Close()

	
	time.Sleep(5*time.Second)


	//delete file
	err = os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File Deleted Successfully")

}