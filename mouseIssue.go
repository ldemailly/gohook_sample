package main

import (
	"flag"

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
		switch ev.Kind {
		case hook.MouseHold:
			log.Infof("Mouse Button %v down (hold)", ev.Button)
			// https://github.com/robotn/gohook/issues/32
			// MouseUp if using the fork github.com/ldemailly/gohook v0.41.0-pre1
			// Otherwise (with original code) need to use MouseDown
		case hook.MouseUp:
			log.Infof("Mouse Button %v released", ev.Button)
		case hook.MouseDown:
			log.Infof("Mouse Button %v down event", ev.Button)
		}
	}
}

func main() {
	flag.Parse()
	Events()
}
