package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dir() (bool, string) {

	// i3wm has two possible config locations
	// finding and confirming confg location
	c := make([]string, 2)
	c[0] = ".i3"
	c[1] = ".config/i3a" // just for testing

	dir := os.Getenv("HOME")
	e1 := false

	for _, i := range c {
		d := filepath.Join(dir, i)
		if stat, err := os.Stat(filepath.Join(dir, i)); err == nil && stat.IsDir() {
			e1 = true
			if e1 == true {
				dir = d
				break
			} else {
				continue
			}
		}
	}

	if e1 == false {
		dir = "i3 config dir not found"
	}
	return e1, dir

}

func tCreate(b string) {

	// checking for theme dir in i3 config dir and creating if missing
	tDir := filepath.Join(b, "theme")
	if stat, err := os.Stat(tDir); err == nil && stat.IsDir() {
		fmt.Println("Creating theme files.")
	} else {
		fmt.Println("No theme folder...")
		os.Chdir(b)
		err := os.Mkdir("theme", 0755)
		check(err)
		fmt.Println("...created")
	}

	themeCreate(tDir)

	fmt.Println("Completed")
}

func i3c() string {
	a, b := dir()

	if a != true {
		fmt.Println(b)
		os.Exit(1)
	}

	return b
}

func main() {

	lsT := flag.NewFlagSet("-l", flag.ExitOnError)
	help := flag.NewFlagSet("-h", flag.ExitOnError)
	setTheme := flag.NewFlagSet("-t", flag.ExitOnError)
	stName := setTheme.String("name", "", "name")
	createTheme := flag.NewFlagSet("-c", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("i3goT v0.1 - for more info use -h")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "-c":
		createTheme.Parse(os.Args[2:])
		b := i3c()
		tCreate(b)
	case "-l":
		lsT.Parse(os.Args[2:])
		fmt.Println("Available Themes:")
		createTheme.Parse(os.Args[2:])
		b := i3c()
		os.Chdir(filepath.Join(b, "theme"))

		r, err := ioutil.ReadDir(".")
		check(err)

		for _, entry := range r {
			fmt.Println(" ", entry.Name())
		}
	case "-t":
		setTheme.Parse(os.Args[2:])
		b := i3c()
		if len(*stName) == 0 {
			fmt.Println("setting default theme")
			themeSet(b, "000")
		} else {
			fmt.Printf("setting %s", *stName)
			themeSet(b, *stName)
		}
	case "-h":
		help.Parse(os.Args[2:])
		fmt.Printf("i3goT v0.1 - for more info use -h\n")
		fmt.Printf("other options:\n")
		fmt.Printf("-c Create a few default themes\n")
		fmt.Printf("-l List available themes for use\n")
		fmt.Printf("-t set default theme\n")
		fmt.Printf("\t-name='theme name'\n")
		fmt.Println("-h print this menu")
	default:
		fmt.Println("i3goT v0.1 - for more info use -h")
		os.Exit(1)
	}
}
