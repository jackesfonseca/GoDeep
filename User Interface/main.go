package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func soma()int{
	var soma int
	soma=10+10
	return soma
}





func main() {
	
	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("UserInterface.qml", 0))

	gui.QGuiApplication_Exec()

}