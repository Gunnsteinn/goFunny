package app

import "github.com/Gunnsteinn/goFunny/controllers/ping"

// mapUrls is used to mapping urls
func mapUrls() {
	router.GET("/ping", ping.Ping)
}
