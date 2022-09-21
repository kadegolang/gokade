package errors

type Errors struct {
	errors map[string][]string
}

func (e *Errors) Add(key, err string) { //增加错误
	if _, ok := e.errors[key]; !ok {
		e.errors[key] = make([]string, 0, 5)
	}
	e.errors[key] = append(e.errors[key], err)
}

func (e *Errors) Errors() map[string][]string { //获取错误
	return e.errors
}

func (e *Errors) ErrorsByKey(key string) []string { //通过key获取错误，返回切片
	return e.errors[key]
}

func (e *Errors) HasErrors() bool { //判断是否错误
	return len(e.errors) != 0 //不等于0 返回错误
}

func New() *Errors {
	return &Errors{
		errors: make(map[string][]string), //errors 初始化
	}
}
