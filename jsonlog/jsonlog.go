package jsonlog

import (
  "fmt"
	"os"
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

// This is a very dumb json func. If more interesting stuff needs to be logged,
// pass it in as a map[string]interface{} and then detect value as int, string, w/e
// before marshaling json.
func (l *JsonLog) Log(msg string, host string, iteration uint64) {
	l.jsonlog.WriteString(fmt.Sprintf("{\"@timestamp\":\"%s\",\"type\":\"mysql_probe\",\"host\":\"%s\",\"iteration\":%v,\"message\":\"%s\"}\n",
		time.Now().Format(time.RFC3339), host, iteration, msg))
}
