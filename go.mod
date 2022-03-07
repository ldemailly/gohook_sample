module github.com/ldemailly/gohook_sample

go 1.17

require (
	fortio.org/fortio v1.21.1
	github.com/robotn/gohook v0.31.3
)

require github.com/vcaesar/keycode v0.10.0 // indirect

replace github.com/robotn/gohook => github.com/ldemailly/gohook v0.41.0-pre1
