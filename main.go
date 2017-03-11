package main

import (
	"flag"
	"fmt"
	"os"
	"log"
)



func main() {

	fs := flag.NewFlagSet("gomoji", flag.ExitOnError)
	version := fs.Bool("version", false, "Print version and exit")

	commands := map[string]command {
		"generate": generateCmd(),
	}

	fs.Usage = func() {
		fmt.Println("Usage: gomoji [global flags] <command> [command flags]")
		fmt.Printf("\nglobal flags:\n")
		fs.PrintDefaults()
	}

	fs.Parse(os.Args[1:])

	if *version {
		fmt.Println(Version)
		return
	}

	args := fs.Args()
	if len(args) == 0 {
		fs.Usage()
		os.Exit(1)
	}

	if cmd, ok := commands[args[0]]; !ok {
		log.Fatalf("Unknown command: %s", args[0])
	} else if err := cmd.fn(args[1:]); err != nil {
		log.Fatal(err)
	}


}

var Version = "???"

type command struct {
	fs *flag.FlagSet
	fn func(args []string) error
}
