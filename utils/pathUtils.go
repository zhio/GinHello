package utils

import (
	"os"
)

func RootPath() string {
	//s,err := exec.LookPath(os.Args[0])
	str, _ := os.Getwd()
	//fmt.Println(str)
	//fmt.Println(s)
	//if err != nil{
	//	log.Panicln("发生错误",err.Error())
	//}
	//i := strings.LastIndex(s,"\\")
	//fmt.Println(i)
	//path := s[0:i+1]
	//fmt.Println(path)
	return str
}
