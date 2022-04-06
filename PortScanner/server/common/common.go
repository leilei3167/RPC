package common

type Task struct {
	dest map[string][]string //ip:port...
}

//server遍历map获取ip和对应的端口,每一个ip开协程扫描

/*
t:=&Task{
dest:map[string][]string{
"127.0.0.1":["8080","27017","8088"],
}
}

*/

type TaskRes struct {
}
