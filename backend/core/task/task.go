package task

import (
	"fmt"
)

// 测试无参数
func Test1() {
	fmt.Println("无参测试")
}

// 测试传参数
func Test2() {
	// 获取参数
	task := GetByName("test2")
	if task == nil {
		return
	}
	fmt.Println(task.Param)
}
