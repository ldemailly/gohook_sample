package main

import (
	"fortio.org/fortio/log"
	hook "github.com/robotn/gohook"
)

func Events() {
	evChan := hook.Start()
	defer hook.End()
	for ev := range evChan {
		log.LogVf("hook: %v", ev)
		if ev.Kind == hook.MouseMove || ev.Kind == hook.MouseDrag {
			continue
		}
		if ev.Kind == hook.MouseHold {
			log.Infof("Mouse Button %v held/pressed", ev.Button)
		}
		// https://github.com/robotn/gohook/issues/32
		// MouseUp if using the fork github.com/ldemailly/gohook v0.41.0-pre1
		// Otherwise need to use MouseDown
		if ev.Kind == hook.MouseUp {
			log.Infof("Mouse Button %v released", ev.Button)
		}
	}
}

func main() {
	Events()
}
