package util

import "fmt"

func Addr(host string, port int) string {
	return fmt.Sprintf("%s:%v", host, port)
}
