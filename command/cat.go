package command

import (
	"fmt"
	"fblviewer/logdata"
	"fblviewer/logic"
	"github.com/spf13/cobra"
)

var catCommand = &cobra.Command{
	Use:   "cat",
	Short: "打印日志",
	Long:  "打印日志",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("ERROR: Please enter the file path to view")
			return
		}
		filePath := args[0]
		logic, err := logic.NewLogic(filePath)
		if err != nil {
			fmt.Println("new logic error:", err.Error())
		}
		logs := logic.List()
		for _, log := range logs {
			fmt.Printf("%s [%s] %s %v %s %d\n", log.Time.Format("2006-01-02 15:04:05"), logdata.LevelName[log.Level], log.Message, log.Fields, log.File, log.Line)
		}
	},
}

var (
	catLevel int8
	catStartAt string
	catEndAt string
)
func init()  {
	root.AddCommand(catCommand)
}

