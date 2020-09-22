package serve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func api(writer http.ResponseWriter, request *http.Request) {
	var temp string
	temp = request.FormValue("level")
	var level int
	var err error
	if temp != "" {
		level,err = strconv.Atoi(temp)
		if err != nil {
			data, _ := json.Marshal(Response{
				Status:  1000,
				Message: "level错误",
			})
			writer.Write(data)
			return
		}
	}else{
		level = -1
	}
	fmt.Println(level)

	//startAt := request.FormValue("start_at")
	//
	//endAt := request.FormValue("end_at")
	//
	//
	//page := request.FormValue("page")
	//limit := request.FormValue("limit")
	//logic.NewFilter()
	//if level == "" {
	//
	//}
}
