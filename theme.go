package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
)

// setting toml struct for color-theme
type Skel struct {
	Bg   string `toml:"client.background"`
	Fo   string `toml:"client.focused"`
	Unfo string `toml:"client.unfocused"`
	Foia string `toml:"client.focused_inactive"`
	Ur   string `toml:"client.urgent"`
	Ph   string `toml:"client.placeholder"`
}

/* setting toml struct for bar-color-theme
type Bar struct {
	Bg string `toml:"background"`
	Sl string `toml:"statusline"`
	Sp string `toml:"separator"`
	Fw string `toml:"focused_workspace"`
	Aw string `toml:"active_workspace"`
	Ia string `toml:"inactive_workspace"`
	Ur string `toml:"urgent_workspace"`
	Bm string `toml:"binding_mode"`
}
*/

type I3wm struct {
	Skel Skel `toml:"i3wm"`
	//Bar  Bar  `toml:"bar"`
}

func themeCreate(a string) {
	dir := a
	os.Chdir(dir)

	t1n := "000.toml"
	t1 := I3wm{Skel{
		Bg:   "#162025",
		Fo:   "#bfbfbf #162025 #bfbfbf #39402e #39402e",
		Unfo: "#bfbfbf #162025 #bfbfbf #75404b #75404b",
		Foia: "#bfbfbf #162025 #bfbfbf #75404b #75404b",
		Ur:   "#bfbfbf #162025 #bfbfbf #75404b #75404b",
		Ph:   "#bfbfbf #162025 #bfbfbf #75404b #75404b"}}

	t2n := "001.toml"
	t2 := I3wm{Skel{
		Bg:   "#a",
		Fo:   "#b #162025 #bfbfbf #39402e #39402e",
		Unfo: "#c #162025 #bfbfbf #75404b #75404b",
		Foia: "#d #162025 #bfbfbf #75404b #75404b",
		Ur:   "#e #162025 #bfbfbf #75404b #75404b",
		Ph:   "#f #162025 #bfbfbf #75404b #75404b"}}

	b0, err := toml.Marshal(t1)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(t1n, b0, os.ModePerm)

	b1, err := toml.Marshal(t2)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(t2n, b1, os.ModePerm)

}

func themeSet(d string, t string) {
	dir := d
	theme := t + ".toml"
	os.Chdir(dir)

	tFile, err := os.Open(theme)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened theme file.")

	defer tFile.Close()

	bytV, _ := ioutil.ReadAll(tFile)
	var cfg I3wm

	toml.Unmarshal(bytV, &cfg)

	c := filepath.Join(d, "config")
	input, err := ioutil.ReadFile(c)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "client.background") {
			lines[i] = "client.background " + "\t" + cfg.Skel.Bg
		}
		if strings.Contains(line, "client.focused ") {
			lines[i] = "client.focused " + "\t" + cfg.Skel.Fo
		}
		if strings.Contains(line, "client.focused_inactive") {
			lines[i] = "client.focused_inactive " + "\t" + cfg.Skel.Foia
		}
		if strings.Contains(line, "client.unfocused") {
			lines[i] = "client.unfocused " + "\t" + cfg.Skel.Unfo
		}
		if strings.Contains(line, "client.urgent") {
			lines[i] = "client.urgent " + "\t" + cfg.Skel.Ur
		}
		if strings.Contains(line, "client.placeholder") {
			lines[i] = "client.placeholder " + "\t" + cfg.Skel.Ph
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(c, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

}
