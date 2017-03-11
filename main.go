package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("mojiemoji!!")
	fs := flag.NewFlagSet("mojiemoji", flag.ExitOnError)
	version := fs.Bool("version", false, "Print version and exit")

	fs.Usage = func() {
		fmt.Println("Usage: mojiemoji [global flags] <command> [command flags]")
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

}

var Version = "???"

// type command struct {
// 	fs *flag.FlagSet
// 	fn func(args []string) error
// }
