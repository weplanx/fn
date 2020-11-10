package utils

type TaskMap struct {
	hashMap map[string]*TaskOption
}

func NewTaskMap() *TaskMap {
	c := new(TaskMap)
	c.hashMap = make(map[string]*TaskOption)
	return c
}

func (c *TaskMap) Put(taskId string, option *TaskOption) {
	c.hashMap[taskId] = option
}

func (c *TaskMap) Get(taskId string) (option *TaskOption, found bool) {
	found = c.hashMap[taskId] != nil
	if found {
		option = c.hashMap[taskId]
	}
	return
}

func (c *TaskMap) Remove(taskId string) (err error) {
	if option, found := c.hashMap[taskId]; found {
		err = option.StreamWriterMap.Clear()
	}
	return
}
