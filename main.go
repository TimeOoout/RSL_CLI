package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/imroc/req"
	"os"
)

func main() {
	client()
}

func client() {
	GlobalInfo()
	_, err := req.Post("http://127.0.0.1:30713/is_exist")
	if err != nil {
		Error(err.Error() + "\nUnable to connect to RedStone Launcher. Please check your Launcher to make sure it is started. " +
			"If it still cannot be solved, please contact the administrator to solve the problem." +
			"\n| Email: hempflower@outlook.com")
	} else {
		var command_line string
		for {
			fmt.Print("RSL> ")
			command_line = ""
			fmt.Scanln(&command_line)
			//fmt.Print("\n")
			if command_line == "/h" || command_line == "/help" {
				show_command_list()
			}
			if command_line == "/exit" {
				Simple("Press Any Key To Quit. . .")
				fmt.Scanln(&command_line)
				os.Exit(0)
			}
		}
	}

}

func show_command_list() {
	Result("" +
		"| Command List                                                         |\n" +
		"| Commands                | Explain                                    |\n" +
		"| /h or /help             | Get command list.                          |\n" +
		"")
}

func GlobalInfo() {
	Simple("Welcome to RedStone Launcher!")
	Simple("Launcher version: 0.0.1_220901_Alpha.")
	Simple("Type '/help' or '/h' for help. Type '/c' to clear the current input statement.\n")
}

func Result(format string, args ...interface{}) {
	fmt.Printf(color.CyanString("[Result] \n"+format+"\n", args...))
}

func Info(format string, args ...interface{}) {
	fmt.Printf(color.GreenString("[INFO] "+format+"\n", args...))
}

func Warn(format string, args ...interface{}) {
	fmt.Printf(color.YellowString("[Warning] "+format+"\n", args...))
}

func Error(format string, args ...interface{}) {
	fmt.Printf(color.RedString("[Error] "+format+"\n", args...))
}

func Debug(format string, args ...interface{}) {
	fmt.Printf(color.BlueString("[Debug] "+format+"\n", args...))
}

func Simple(format string, args ...interface{}) {
	fmt.Printf(color.WhiteString(format+"\n", args...))
}
