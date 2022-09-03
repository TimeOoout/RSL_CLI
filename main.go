package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/imroc/req"
	"os"
	"time"
)

func main() {
	client()
	os.Exit(-1)
}

func client() {
	GlobalInfo()
	_, err := req.Post("http://127.0.0.1:30713/is_exist")
	if err != nil {
		Error(err.Error() + "\nUnable to connect to RedStone Launcher. Please check your Launcher to make sure it is started. " +
			"If it still cannot be solved, please contact the administrator to solve the problem." +
			"\n| Email: hempflower@outlook.com")
	} else {
		go Backstage_check_connect()
		var command_line string
		for {
			fmt.Print("RSL> ")
			command_line = ""
			fmt.Scanln(&command_line)
			//fmt.Print("\n")
			if command_line == "/h" || command_line == "/help" {
				Help()
			}
			if command_line == "/exit" {

			}
		}
	}

}

func Help() {
	Result("" +
		"| Command List                                                         |\n" +
		"| Commands                | Explain                                    |\n" +
		"| /h or /help             | Get command list.                          |\n" +
		"| /c                      | Clear the current input statement.         |\n" +
		"| /exit                   | Exit the program.                          |\n" +
		"| /add_dir                 | Add a new game dir.                        |\n" +
		"")
}

func Exit(code int) {
	Info("Press Any Key To Quit. . .")
	input := ""
	fmt.Scanln(&input)
	os.Exit(code)
}

func Add_dir() {

}

func Check_connect() {
	_, err := req.Post("http://127.0.0.1:30713/is_exist")
	if err != nil {
		Error(err.Error() + "\nUnable to connect to RedStone Launcher. Please check your Launcher to make sure it is started. " +
			"If it still cannot be solved, please contact the administrator to solve the problem." +
			"\n| Email: hempflower@outlook.com")
		Exit(-2)
	}
}

func GlobalInfo() {
	fmt.Print(color.CyanString("Welcome to RedStone Launcher!\n" +
		"Powered by TimeOut | TimeOoout@outlook.com\n" +
		"Launcher version: 0.0.1_220901_Alpha.\n" +
		"Type '/help' or '/h' for help. Type '/c' to clear the current input statement.\n"))
}

func Backstage_check_connect() {
	_, err := req.Post("http://127.0.0.1:30713/is_exist")
	if err != nil {
		Error(err.Error() + "\nUnable to connect to RedStone Launcher. Please check your Launcher to make sure it is started. " +
			"If it still cannot be solved, please contact the administrator to solve the problem." +
			"\n| Email: hempflower@outlook.com")
		Exit(-2)
	}
	time.Sleep(time.Second * 5)
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
