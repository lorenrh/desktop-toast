package main

import (
	"time"

	"github.com/ulyssessouza/desktop-toast/toast"
)

func main() {
	c := toast.NewClient(50 * time.Millisecond)
	err := c.SendToast(toast.NativeNotificationModel{
		Title: "MyToastTitle",
		Body:  "My Toast Body",
		Level: toast.Warning,
	})
	if err != nil {
		panic(err)
	}
}
