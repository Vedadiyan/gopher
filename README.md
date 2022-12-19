# Gopher (as an extension to the original `go` CLI)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![Go report](https://goreportcard.com/badge/github.com/vedadiyan/gopher)](https://goreportcard.com/report/github.com/vedadiyan/gopher)

Gopher is a simple package manager that resembles core features and mechanism of npm. It manages and maintains all project's dependencies in a `.json` file allowing them to be restored when required without relying on the `go.mod` and `go.sum` files. When a restore operation is requested, gopher will automatically create `go.mod` and `go.sum` files both for the current project and all its dependencies. Accordingly, `go.mod` and `go.mod` can be added to the `.gitignore` file. 

## 🚀 Installation Guide
`go build -o gopher ./cmd/` 
or for Windows or Mac
`go build -o gopher.(exe|dmg) ./cmd/`
then you can easily run the executable with the `setup` command: 
`./gopher setup`.

You need to add the path `HomeDirectory/gopher/bin` to your path variables.

**In Linux you should assign executable permissions as well. For example `chmod 777 ~/gopher/bin/gopher`**

## 💡 Commands 

**go-painess falls through the original `go` command when a command is not available. This means that you can use it as a drop-in replacement for the original `go` command.**

*all UPPERCARE flags staring with a single `-` are required*

|Command| Description  | Example | Notes |
|--|--|--|--|
| init | creates a new go project  | gopher init -n demo -v v1.0.0| --name or -n = the name of the project <br /> --version or -v = the version of the project 
| create | creates a project based on a template project | gopher create -t github.com/abc/efg.git -n github.com/abc/xyz | --template or -t = the template repository url <br/> --name or -n = the name of the project
|install| installs a go dependency | gopher install -u https://github.com/abc/efg.git -n custom_dependency_name --private --recursive | --url or -u = the URL of the dependency (whether private or public) <br /> --name or -n = the name used to reference the dependency. This name is used for referencing private packages.  <br />  --private = used for installing private packages <br /> --recursive = used for installing nested dependencies in gopher maintained packages <br /> --update = used for updating existing packages 
| remove | removes a go dependency | gopher remove -n custom_dependency_name | --name or -n = the name of the dependency to be removed
| restore | restores all dependencies | gopher restore | --update = used for updating existing dependency <br /> --update-global = used for updating global dependencies (Experimental) <br /> --tidy = runs `go mod tidy` after the restore has completed 
| clear | clear the project | gopher clear | -
| publish | creates binaries for a go project | gopher publish -r linux -a amd64 -o ./ -T ./cmd/ | --runtime or -r = specifies target OS <br /> --architecture or -a = specifies target architecture <br /> --output or -o = specifies the output directory <br /> --template or -t = specifies the go file or folder to build 

### Tips
It is recommended to always use `gopher tidy` after `gopher restore`.  This is because `gopher` is a complement to the original `go` command. 
When using `gopher`, the project has to be initiated using `gopher initialize` and all dependencies have to be maintained using the `gopher install` commands.

## 🤝 Contibution
Feel free to contibute by forking the project! All new features are welcome. For any issues please open an issue so we can find and fix the problem. 

## 💫 Show your support

Give a ⭐️ if it kills the pain of managing dependencies in your go projects!

## 📝 License

Copyright © 2022 [Pouya Vedadiyan](https://github.com/vedadiyan).

This project is [MIT](./LICENSE) licensed.

