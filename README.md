# go-vte
This is a binding for libvte

The binding acts as an extension to github.com/gotk3/gotk3/gtk

## Example


```go
term, err := vte.TerminalNew()
if err != nil {
	log.Fatal("Unable to create terminal:", err)
}
wd, _ := os.Getwd()
term.SpawnAsyncSimple(wd, []string{"/bin/bash"}, []string{"TERM=xterm"})
```
