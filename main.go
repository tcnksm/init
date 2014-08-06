package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/dotcloud/docker/pkg/mflag"
)

var (
	flVersion = flag.Bool([]string{"v", "-version"}, false, "Print version information and quit")
	flHelp    = flag.Bool([]string{"h", "-help"}, false, "Print this message and quit")
	flDebug   = flag.Bool([]string{"-debug"}, false, "Run as DEBUG mode")
)

func main() {
	// call hpMain in a separate function
	// so that it can use defer and have them
	// run before the exit.
	os.Exit(hpMain())
}

func hpMain() int {

	flag.Parse()

	if *flDebug {
		os.Setenv("DEBUG", "1")
	}

	if *flVersion {
		showVersion()
		return 0
	}

	if *flHelp {
		showHelp()
		return 0
	}

	return 0
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func showVersion() {
	fmt.Fprintf(os.Stderr, "hp version %s, build %s \n", Version, GitCommit)
}

func showHelp() {
	fmt.Fprintf(os.Stderr, helpText)
}

const helpText = `Usage: hp [options] [command] [<room>] [<message>]

hp - Most simple HipChat command line client.

Commands:
    chat            Set initial settings
    clean           Clean up all settings

Options:
    -h, --help      Print this message and quit
    -v, --version   Show version information and quit
    --debug         Run as DEBUG mode

Example:
    $ hp A-team
    $ hp dev-team Hi, I'll try to deploy from now
`
