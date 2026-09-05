// Command quick-txt-save-alfred is the binary the packaged Alfred Workflow
// invokes (see workflow/info.plist).
//
// Subcommands, each corresponding to one workflow/info.plist node's script:
//
//	list [filename]  — the "save" Script Filter keyword node; previews the
//	                    resolved destination path
//	write <path>      — the Run Script action after Enter; writes the text
//	                    carried in the $text env var (an Arguments and
//	                    Variables node sets it to Alfred's {clipboard}
//	                    placeholder) to path and prints the outcome as an
//	                    Alfred workflow variable for the downstream native
//	                    Post Notification node to display
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/y-marui/alfred-quick-txt-save/internal/quicksave"
	"github.com/y-marui/alfred-quick-txt-save/internal/quicksavecmd"
	"github.com/y-marui/alfred-quick-txt-save/internal/scriptfilter"
)

func main() {
	if len(os.Args) < 2 {
		writeErrorResponse(fmt.Errorf("missing subcommand"))
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		runList()
	case "write":
		runWrite()
	default:
		writeErrorResponse(fmt.Errorf("unknown subcommand %q", os.Args[1]))
		os.Exit(1)
	}
}

func runList() {
	filename := ""
	if len(os.Args) > 2 {
		filename = os.Args[2]
	}
	writeResponse(dispatch(filename))
}

// dispatch recovers from any panic in quicksavecmd, mirroring the Python
// workflow's safe_run: an unhandled failure must still produce a visible
// Script Filter error item rather than empty/invalid output.
func dispatch(filename string) (resp scriptfilter.Response) {
	defer func() {
		if r := recover(); r != nil {
			resp = errorResponse(fmt.Sprintf("%v", r))
		}
	}()
	return quicksavecmd.List(filename)
}

func runWrite() {
	if len(os.Args) < 3 || os.Args[2] == "" {
		fmt.Fprintln(os.Stderr, "quick-txt-save-alfred: write: path argument is required")
		os.Exit(1)
	}
	text := os.Getenv("text")
	_, message, err := quicksave.SaveText(os.Args[2], text)
	if err != nil {
		fmt.Fprintln(os.Stderr, "quick-txt-save-alfred:", err)
	}
	writeVariables(map[string]string{"message": message})
}

// writeVariables prints Alfred's workflow-variables JSON envelope
// (https://www.alfredapp.com/help/workflows/advanced/variables/) to stdout.
// A Run Script action has no way to branch a downstream connection on its
// own exit code, but any node wired after it — here, a native Post
// Notification node whose text is "{message}" — can read a variable set
// this way regardless of whether the write succeeded.
func writeVariables(vars map[string]string) {
	payload := struct {
		Alfredworkflow struct {
			Variables map[string]string `json:"variables"`
		} `json:"alfredworkflow"`
	}{}
	payload.Alfredworkflow.Variables = vars
	if err := json.NewEncoder(os.Stdout).Encode(payload); err != nil {
		fmt.Fprintln(os.Stderr, "quick-txt-save-alfred: writing variables:", err)
	}
}

func writeResponse(resp scriptfilter.Response) {
	if err := resp.Write(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "quick-txt-save-alfred: writing response:", err)
		os.Exit(1)
	}
}

// writeErrorResponse always emits valid Script Filter JSON — even on an
// unexpected internal failure — so Alfred shows a readable error row
// instead of an empty/unparseable result.
func writeErrorResponse(err error) {
	_ = errorResponse(err.Error()).Write(os.Stdout)
}

func errorResponse(message string) scriptfilter.Response {
	return scriptfilter.Response{
		Items: []scriptfilter.Item{
			{Title: "Workflow Error", Subtitle: message, Arg: message, Valid: scriptfilter.BoolPtr(false)},
		},
	}
}
