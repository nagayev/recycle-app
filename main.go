package main

import "github.com/zserge/webview"

func main() {
	// Open our app url 
	webview.Open("recycle",
		"https://jagkadmin.pythonanywhere.com/", 607, 500, false)
}
