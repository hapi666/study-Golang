package main

import "os"

/**
* os.Args[0]是当前所执行的文件路径
* os.Args[1],os.Args[2]等等...是命令行第一个参数，第二个参数...
**/
func main() {
	println("I am ", os.Args[0])
}
