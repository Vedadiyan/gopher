package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
	flaggy "github.com/vedadiyan/flaggy/pkg"
	gopher "github.com/vedadiyan/gopher/internal"
)

type Flags struct {
	Create  Create  `long:"create" short:"" help:"Used for creating a new project based on an existing template"`
	Setup   bool    `long:"setup" short:"" help:"Setups gopher in the system"`
	Init    Init    `long:"init" short:"" help:"Initializes a new project"`
	Install Install `long:"install" short:"" help:"Installs a dependency"`
	Remove  Remove  `long:"remove" short:"" help:"Removes an existing dependency"`
	Restore Restore `long:"restore" short:"" help:"Restores dependencies in an existing project"`
	Clear   bool    `long:"clear" short:"" help:"Removes go.mod and go.sum files"`
	Publish Publish `long:"publish" short:"" help:"Builds the project"`
	Help    bool    `long:"help" short:"" help:"Shows gopher help"`
}

func (f Flags) Run() error {
	if !flaggy.Parsed() {
		color.Hex(gopher.YELLOW).Println("Falling back to `go`")
		gopher.Run("go", strings.Join(os.Args[1:], " "), nil)
		return nil
	}
	if f.Setup {
		gopher.Setup()
		return nil
	}
	if f.Help {
		flaggy.PrintHelp()
		return nil
	}
	return nil
}

type Create struct {
	Template string `long:"--template" short:"-t" help:"Specifies the template url"`
	Name     string `long:"--name" short:"-n" help:"Specifies the project name"`
}

func (c Create) Run() error {
	failing := false
	if len(c.Template) == 0 {
		failing = true
		fmt.Println("template is required")
	}
	if len(c.Name) == 0 {
		failing = true
		fmt.Println("name is required")
	}
	if failing {
		flaggy.PrintHelp()
		return nil
	}
	gopher.CreateFromTemplate(c.Template, c.Name)
	return nil
}

type Init struct {
	Name    string `long:"--name" short:"-n" help:"Specifies project name"`
	Version string `long:"--version" short:"-v" help:"Specifies project version"`
}

func (i Init) Run() error {
	failing := false
	if len(i.Name) == 0 {
		failing = true
		fmt.Println("name is required")
	}
	if len(i.Version) == 0 {
		failing = true
		fmt.Println("version is required")
	}
	if failing {
		flaggy.PrintHelp()
		return nil
	}
	gopher.PkgFileCreate(i.Name, i.Version)
	gopher.ModFileCreate(i.Name, "")
	return nil
}

type Install struct {
	Url       string `long:"--url" short:"-u" help:"Specifies dependency URL"`
	Name      string `long:"--name" short:"-n" help:"Specifies dependency name"`
	Private   bool   `long:"--private" short:"-p" help:"Used for installing from private repositories"`
	Recursive bool   `long:"--recursive" short:"-r" help:"Used for recursively installing depdendencies"`
	Update    bool   `long:"--update" short:"-u" help:"Used for updating previously downloaded dependencies"`
}

func (i Install) Run() error {
	failing := false
	if len(i.Url) == 0 {
		failing = true
	}
	if len(i.Name) == 0 {
		failing = true
	}
	if failing {
		color.Hex(gopher.YELLOW).Println("Falling back to `go`")
		gopher.Run("go", fmt.Sprintf("install %s", strings.Join(os.Args[2:], "")), nil)
		return nil
	}
	gopher.PkgFileLoad()
	gopher.PkgAdd(i.Url, i.Name, i.Private, i.Update, i.Recursive)
	gopher.Write()
	return nil
}

type Remove struct {
	Name string `long:"--name" short:"-n" help:"Specifies dependency name"`
}

func (r Remove) Run() error {
	failing := false
	if len(r.Name) == 0 {
		failing = true
		fmt.Println("name is required")
	}
	if failing {
		flaggy.PrintHelp()
		return nil
	}
	gopher.PkgDelete(r.Name)
	return nil
}

type Restore struct {
	Tidy   bool `long:"--tidy" short:"-t" help:"Runs go mod tidy after restoring the project"`
	Update bool `long:"--update" short:"-u" help:"Used for updating previously downloaded dependencies"`
}

func (r Restore) Run() error {
	gopher.PkgFileLoad()
	gopher.PkgRestore(true, r.Update)
	gopher.Write()
	if r.Tidy {
		gopher.Tidy()
	}
	return nil
}

type Publish struct {
	Runtime      string `long:"--runtime" short:"-r" help:"Specifies the runtime"`
	Architecture string `long:"--architecture" short:"-a" help:"Specifies build architecture"`
	Output       string `long:"--output" short:"-o" help:"Specifies the output"`
	Target       string `long:"--target" short:"-t" help:"Specifies the target build path"`
}

func (p Publish) Run() error {
	failing := false
	if len(p.Architecture) == 0 {
		failing = true
		fmt.Println("architecture is required")
	}
	if len(p.Output) == 0 {
		failing = true
		fmt.Println("output is required")
	}
	if len(p.Runtime) == 0 {
		failing = true
		fmt.Println("runtime is required")
	}
	if len(p.Target) == 0 {
		failing = true
		fmt.Println("target is required")
	}
	if failing {
		flaggy.PrintHelp()
		return nil
	}
	gopher.Build(p.Runtime, p.Architecture, p.Output, p.Target)
	return nil
}
