package main

import "github.com/zserge/webview"

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("recycle",
		"https://jagkadmin.pythonanywhere.com/", 607, 500, false)
}