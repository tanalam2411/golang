

```bash
cobra_cli$ go get -u github.com/spf13/cobra/cobra
```
```bash
cobra_cli$ cobra init --pkg-name github.com/tanalam2411/golang/cobra_cli .
Your Cobra applicaton is ready at
/home/tan/tanveer/golang/src/golang/cobra_cli

cobra_cli$ go mod init github.com/tanalam2411/golang/cobra_cli
go: creating new go.mod: module github.com/tanalam2411/golang/cobra_cli

cobra_cli$ cat go.mod 
module github.com/tanalam2411/golang/cobra_cli

go 1.13

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.6.1
)

```

```bash
cobra_cli$ cobra add serve
cobra_cli$ cobra add config
cobra_cli$ cobra add create -p 'configCmd'
```

```bash

```