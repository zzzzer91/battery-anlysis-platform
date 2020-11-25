package main

import (
	"battery-analysis-platform/app/web"
)

func main() {
	if err := web.Run(); err != nil {
		panic(err)
	}
}
