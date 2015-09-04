package statusserver

import (
  "strconv"
  "fmt"
  "log"
  "net/http"
  "os"
  "regexp"
  "github.com/haikulearning/mysql_probe/mysqltest"
	"github.com/haikulearning/mysql_probe/jsonlog"
)

var required_up_checks = []string{"connect", "threads_connected_count_lte_2400"}

type StatuServer struct {
	reportdir     string
	port          int
	jsonlog       *jsonlog.JsonLog
}

func StartStatuServer(reportdir string, port int, log_path string) *StatuServer {

	jsonlog := jsonlog.Init(log_path)
	defer jsonlog.LogFile().Close()

	s := StatuServer{reportdir: reportdir, port: port, jsonlog: jsonlog}

  s.Start()

	return &s
}

// Start up the status server
func (s *StatuServer) Start() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    s.handler(w, r)
  })
  //fs := http.FileServer(http.Dir("tmp"))
  //http.Handle("/", fs)

  log.Println("Listening to port " + strconv.Itoa(s.port))
  http.ListenAndServe(":" + strconv.Itoa(s.port), nil)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func (s *StatuServer) handler(w http.ResponseWriter, r *http.Request) {
  is_up := true

  for _,testname := range required_up_checks {
    if is_up {
      is_up = s.testResultIsUp(testname, r)
    }
  }

  if is_up {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "up\n")
  } else {
    w.WriteHeader(http.StatusServiceUnavailable)
    fmt.Fprintf(w, "down\n")
  }
}

func (s *StatuServer) testResultIsUp(testname string, r *http.Request) bool {
  testpath := mysqltest.TestResultPath(s.reportdir, testname)
  match := false

  defer func() {
    if r := recover(); r != nil {
      log.Println("down via failure checking " + testpath + " (skipping all subsequent checks)")
    }
  }()

  f, err := os.Open(testpath)
  check(err)

  b1 := make([]byte, 10)
  _, err = f.Read(b1)
  check(err)

  match, err = regexp.MatchString("up", string(b1))
  check(err)

  if match {
    s.Log("up via " + testpath, r)
  } else {
    s.Log("down via " + testpath + " (skipping all subsequent checks)", r)
  }

  return match
}

func (s *StatuServer) Log(msg string, r *http.Request) {
	json_msg := map[string]interface{}{
		"Method": 		r.Method,
		"RequestURI": r.RequestURI,
		"RemoteAddr": r.RemoteAddr,
	}

	s.jsonlog.Log(msg, json_msg)
}
