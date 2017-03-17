package main

import (

	"log"
	"net/http"
	"strings"
	"os/exec"
	"os"
)

func main () {
	http.HandleFunc("/it", handler)
	err := http.ListenAndServe(":8787", nil)
	log.Fatal(err)
}

func handler (resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("<script>window.close();</script>"))
	script := `tell application "iTerm2"
    set newWindow to (create window with default profile)
        tell current session of newWindow
        write text "COMMAND"
    end tell
end tell`

	input := req.URL.Query()

	script = strings.Replace(script, "COMMAND", input["q"][0], -1)

	cmd := *exec.Command("osascript")
	cmd.Stdin = strings.NewReader(script)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil{
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

}
