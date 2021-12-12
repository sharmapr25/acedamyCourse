package main

import "sample/acedamyCourse/app"

func main(){
	ap := &app.App{}
	ap.Initialize()
	ap.Run()

}

