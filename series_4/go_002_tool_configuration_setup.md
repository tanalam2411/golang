GO Artifacts

- GO Source to Binary

go_source_binary.png

- GO Code Organization
  - A package contains one or more GO source 
    - Note: A GO source file must be in exactly one package
    - Packages are reusable pieces of code, and works to modularize a large project
  - A typical GO project (application or package), is developed in modules
  - A module provides versioning for pakcages used in your code


- Go modules
  - go help modules
  - https://github.com/golang/go/wiki/Modules#go-111-modules
  - go mod init golang/series_4


- Your GO Source Dir
  - Is any directory on your filesystem
  - Typically, you will have one directory into which you will create several modules
    - Within which is one directory per module


- GOBIN Directory
  - GOBIN is the directory GO uses to store binaries
    - These are applications installed using:
      - go get
      - go install

- Setting GOBIN & PATH
  - Set GOBIN Variable
    - GOBIN=${HOME}/go/bin
      - Mac/Linux Users
        - Set GOBIN in your shell login file
          - Eg: ~/.zshrc, ~/.bashrc, ~/.cshrc
      $ echo 'GOBIN=${HOME}/go/bin' >> ~/.bashrc
  - Set PATH variable
    - PATH=${GOBIN}:${PATH}
    - $ echo 'PATH=${GOBIN}:${PATH}' >> ~/.bashrc

- Source Dir
  - Get the source 

        