package exercise

import "runtime"

func Version() string {
	return runtime.Version()
}
