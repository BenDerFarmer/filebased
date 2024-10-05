package main

import (
	"fmt"
	"github.com/ChaotenHG/filebased/internal"
	flag "github.com/spf13/pflag"
	"os"
)

const version = "0.1.0"

func main() {

	outfile_name := flag.StringP("output", "o", "routes", "Output file name (without .go)")
	routes_dir := flag.StringP("input", "i", "./routes", "directory of your route files")
	pkg_name := flag.StringP("package", "p", "main", "the name of your package where the server is defined")

	showHelp := flag.BoolP("help", "h", false, "Show usage summary")
	showVersion := flag.BoolP("version", "v", false, "Show version")
	flag.Parse()

	if *showHelp {
		fmt.Println("Usage: filebased [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("Filebased: version %s\n", version)
		os.Exit(0)
	}

	if err := internal.Build(*routes_dir, *outfile_name, *pkg_name); err != nil {
		fmt.Printf("Filebased Error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Filebased: %s.go was successfully generated.\n", *outfile_name)
}
