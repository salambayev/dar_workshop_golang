package main

import (
	"net/http"

	"io/ioutil"

	"encoding/json"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/salambayev/workshop_go_dar/detector"
)

type Command struct {
	Cmd string `json:"cmd"`
}

type CommandExecutor struct {
	detector detector.DetectorCommand
}

//func ExecuteCommand(command string, wg *sync.WaitGroup, ch *chan string) {
//	fmt.Println("Executing command: " + command)
//	parts := strings.Split(command, " ")
//	head := parts[0]
//	tail := parts[1:len(parts)]
//	out, err := exec.Command(head, tail...).Output()
//	if err != nil {
//		*ch <- err.Error()
//	}
//	*ch <- string(out)
//	wg.Done()
//}

func main() {
	//ch := make(chan string)
	//wg := new(sync.WaitGroup)
	//wg.Add(3)
	//go ExecuteCommand("brew services list", wg, &ch)
	//go ExecuteCommand("ls ../../", wg, &ch)
	//go ExecuteCommand("ls", wg, &ch)
	//
	////ind := 0
	//var ind int
	//for item := range ch {
	//	ind++
	//	fmt.Println(item)
	//	if ind == 3 {
	//		close(ch)
	//	}
	//}
	r := mux.NewRouter()
	r.HandleFunc("/kubernetes", ExecCommand).Methods("POST")
	http.ListenAndServe("localhost:8080", r)
}

func ExecCommand(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	data := &Command{}
	err = json.Unmarshal(body, data)
	fmt.Println(*data)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(data)
	dtc := &CommandExecutor{detector.NewDetectorCommand(detector.NewExecuteCommand())}
	output, err := dtc.detector.ExecuteCommand(data.Cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(output)
	fmt.Println(string(output))
	w.WriteHeader(http.StatusOK)
	w.Write(output)
	return
}
