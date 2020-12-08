package task

import "log"

func init() {
	log.Println("初始化任务。。。")

	addTask("test1", Test1)
	addTask("test2", Test2)
}

//任务信息
type TaskEntity struct {
	FuncName string
	Param    []string
	Run      func()
}

var taskList = make([]TaskEntity, 0)

// 添加任务
func addTask(funcName string, ifunc func()) {
	var tast TaskEntity
	tast.FuncName = funcName
	tast.Param = nil
	tast.Run = ifunc
	Add(tast)
}

//检查方法名是否存在
func GetByName(funcName string) *TaskEntity {
	for _, task := range taskList {
		if task.FuncName == funcName {
			return &task
		}
	}
	return nil
}

//增加Task方法
func Add(task TaskEntity) {
	if task.FuncName == "" || task.Run == nil {
		return
	}

	taskList = append(taskList, task)
}

//修改参数
func EditParams(funcName string, params []string) {
	for index := range taskList {
		if taskList[index].FuncName == funcName {
			taskList[index].Param = params
			break
		}
	}
}
