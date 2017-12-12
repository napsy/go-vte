package vte

// #cgo pkg-config: vte-2.91
// #include <vte/vte.h>
// #include "vte.go.h"
// #include <stdlib.h>
import "C"

import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	. "github.com/gotk3/gotk3/gtk"
)

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

type Terminal struct {
	Widget
}

func wrapWidget(obj *glib.Object) *Widget {
	return &Widget{glib.InitiallyUnowned{obj}}
}

func wrapTerminal(obj *glib.Object) *Terminal {
	w := &Widget{glib.InitiallyUnowned{obj}}
	return &Terminal{*w}
}

func (t *Terminal) native() *C.VteTerminal {
	p := unsafe.Pointer(t.GObject)
	return C.toVteTerminal(p)
}
func TerminalNew() (*Terminal, error) {
	t := C.vte_terminal_new()
	if t == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(t))
	return wrapTerminal(obj), nil
}

func (t *Terminal) Feed(text string) {
	str := C.CString(text)
	defer C.free(unsafe.Pointer(str))
	C.vte_terminal_feed(t.native(), str, C.gssize(len(text)))
}

func (t *Terminal) FeedChild(text string) {
	str := C.CString(text)
	defer C.free(unsafe.Pointer(str))
	C.vte_terminal_feed_child(t.native(), str, C.gssize(len(text)))
}

func makeStrings(array []string) **C.char {
	cArray := C.make_strings(C.int(len(array)))
	for i, e := range array {
		cstr := C.CString(e)
		C.set_string(cArray, C.int(i), (*C.char)(cstr))
	}
	C.set_string(cArray, C.int(len(array)), nil)
	return cArray
}

func destroyStrings(strings **C.char, count int) {
	C.destroy_strings(strings, C.int(count))
}

func (t *Terminal) SpawnAsyncSimple(workingDirectory string, argv, envv []string) error {
	wd := C.CString(workingDirectory)
	defer C.free(unsafe.Pointer(wd))

	argvStrings := makeStrings(argv)
	envvStrings := makeStrings(envv)
	defer destroyStrings(argvStrings, len(argv))
	defer destroyStrings(envvStrings, len(envv))

	timeout := -1 // in ms (wait indefinitely)

	C.vte_terminal_spawn_async(
		t.native(),     // terminal
		0,              // VTE_PTY_DEFAULT
		wd,             // working directory
		argvStrings,    // argv
		envvStrings,    // env
		0,              // G_SPAWN_DEFAULT
		nil,            // child_setup
		nil,            // child_setup_data
		nil,            // child_setup_data_destroy
		C.int(timeout), // timeout
		nil,            // cancellable
		nil,            // command_spawned
		nil,            // user_data
	)
	return nil
}
