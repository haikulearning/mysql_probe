package jsonlog

import (
  "fmt"
	"os"
	"encoding/json"
	"time"
  //"github.com/haikulearning/mysql_probe/mysqltest"
)

type JsonLog struct {
	jsonlog    *os.File
}

func Init(log_path string) *JsonLog {
	open_file, err := os.OpenFile(log_path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("Couldn't open log_path \"%s\" for writing: %s", log_path, err.Error()))
	}

	l := JsonLog{jsonlog: open_file}

	return &l
}

func (l *JsonLog) LogFile() *os.File {
	return l.jsonlog
}

func (l *JsonLog) Log(msg string, other_data interface{}) {
	data, err := json.Marshal(other_data)
	if err != nil {
		panic(err.Error())
	}

	l.jsonlog.WriteString(
		fmt.Sprintf("{\"@timestamp\":\"%s\",\"type\":\"mysql_probe\",\"message\":\"%s\",\"data\":",
			time.Now().Format(time.RFC3339), msg))
	l.jsonlog.Write(data)
	l.jsonlog.WriteString("}\n")
}
