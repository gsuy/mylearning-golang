## Start with command ...
- go mod init <projectname> // to create go.mod, like node init to create package.json ( but more betterrrr!!! hahaha \('')/ )

## Semantic versioning
- version of package e.g. x.xx.xxx (Major.Minor.Patch)
- x is Patch/Build => fix bug of system or little bug (low risk) 
- xx is Minor => new feature (medium risk)
- xxx is Major => big impack of package update (often followed by RC(Release Candidate), Preview, Beta for listen the problem) (high risk)

## Interesting command
- go get => install package follow dot go files
- go mod tidy => delete dependencies isn't active
- go mod vendor

## Executables for Multiple Platforms
- [How To Build Go Executables for Multiple Platforms on Ubuntu 16.04](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)