package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/imroc/req"
	"github.com/tidwall/sjson"
	"os"
	"time"
)

var host = "http://127.0.0.1:30713/"

var debug = true

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
		if debug == true {
			Debug("Debug:True")
		}
		var command_line string
		fmt.Print("\n")
		for {
			fmt.Print("Launcher> ")
			command_line = ""
			fmt.Scanln(&command_line)
			//fmt.Print("\n")
			if command_line == "/h" || command_line == "/help" {
				Help()
			} else if command_line == "/exit" {
				Exit(0)
			} else if command_line == "/add_dir" {
				Select("Path>")
				path := ""
				fmt.Scanln(&path)
				Select("Name>")
				name := ""
				fmt.Scanln(&name)
				Add_dir(path, name)
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
		"| /add_dir                | Add a new game directory.                  |\n" +
		"")
}

func Exit(code int) {
	Info("Press Any Key To Quit. . .")
	input := ""
	fmt.Scanln(&input)
	os.Exit(code)
}

func Add_dir(path string, name string) {
	head, _ := sjson.Set("", "Path", path)
	head, _ = sjson.Set(head, "Name", name)
	client_ := req.New()
	content, err := client_.Post(host+"add_game_dir", head)
	Debug(head)
	if err != nil {
		Error("Failed to add a new game directory. Error:" + err.Error())
	} else {
		if content.String() == "" {
			Info("Successfully added a new game directory!")
		} else {
			Error(content.String())
		}
	}
}

func GlobalInfo() {
	fun := color.New(color.Bold, color.FgCyan).PrintfFunc()
	fun("RedStone Launcher [v0.0.1_220901_Alpha]\n" +
		"Powered by TimeOut.\n" +
		"Copyright (c) 2022 TimeOut. All rights reserved.\n" +
		"Launcher version: v0.0.1_220901_Alpha.\n" +
		"Type '/help' or '/h' for help. Type '/c' to clear the current input statement.\n")
}

func Backstage_check_connect() {
	_, err := req.Post(host + "is_exist")
	if err != nil {
		Error(err.Error() + "\nUnable to connect to RedStone Launcher. Please check your Launcher to make sure it is started. " +
			"If it still cannot be solved, please contact the administrator to solve the problem." +
			"\n| Email: hempflower@outlook.com")
		Exit(-2)
	}
	time.Sleep(time.Second * 5)
}

func Select(format string, args ...interface{}) {
	fun := color.New(color.Bold, color.FgCyan).PrintfFunc()
	fun(format, args...)
}

func Result(format string, args ...interface{}) {
	color.Cyan("[Result] \n"+format+"\n", args...)
}

func Info(format string, args ...interface{}) {
	color.Green("[INFO] "+format+"\n", args...)
}

func Warn(format string, args ...interface{}) {
	color.Yellow("[Warning] "+format+"\n", args...)
}

func Error(format string, args ...interface{}) {
	color.Red("[Error] "+format+"\n", args...)
}

func Debug(format string, args ...interface{}) {
	if debug == true {
		color.Blue("[Debug] "+format+"\n", args...)
	}
}

func Simple(format string, args ...interface{}) {
	color.White(format+"\n", args...)
}
