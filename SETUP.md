# Setting Up Your Environment
### Installing Go on Mac OSX 
#### Install brew
If you do not already have brew installed, you can install 
Brew [HERE](http://brew.sh/).
You can see if Brew is installed and the current version by running
```bash
brew config
```

#### Install Git
Check to see that you have git installed by running `git version`. If you
do not have git installed, install git by running 
```bash
brew install git 
```

#### Install Go v1.12+
The Go binary can be installed using Brew as well. This material 
is written under the assumptions of Go v1.12. You can read the Golang
release notes [HERE](https://golang.org/doc/devel/release.html).
```bash
brew install go
```

#### Set GOPATH
Go modules will eventually completely replace the need for a GOPATH. In the meantime,
Go searches in two locations for executables and code source `$GOPATH` and `$GOROOT`
 - `$GOROOT` will be set by the installer for your system and the location of GOROOT will
be the location for builtin packages and the compiler.
 - `$GOPATH` will be the location of your Go workspace. To set your "workspace" run:
 ```bash
 mkdir -p ~/Code/go
 
 echo "export GOPATH=~/Code/go" >> ~/.bash_profile
 # OR
 vi ~/.bash_profile
 export GOPATH=~/Code/go

 source ~/.bash_profile
 ```

#### Testing Your Go Install
```bash
go env
go version
```
The output of `go env` shows all of the environment information regarding your Golang
install. The important things to check here are that `$GOPATH` and `$GOROOT` are set
correctly. Your Go ENV should look _similar_ to this:
```bash
...
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/josh5276/Code/go" 
GOPROXY=""
GORACE=""
GOROOT="/usr/local/Cellar/go/1.12.5/libexec"
....
```

#### Pull Down Course Repo
The Go workspace mimics the file structure that github uses for repo and source
code management. We will create the directories to keep you local code consistent
with git hierarchy.
```bash
mkdir -p ~/Code/go/src/github.com/josh5276
cd ~/Code/go/src/github.com/josh5276
git clone https://github.com/josh5276/go-course
```

### 3rd Party Packages
There are a large amount of builtin packages already in the Go ecosystem within the
standard library. Occasionally (or often), you will need to pull down 3rd party packages
into your application for a specific feature. 

You will find some packages that you used consistently and start to rely on as you
develop more applications. For the purposes of this course, I use the logrus package
to print std logging to stdout/stderr. You can pull down 3rd party packages using
`go get`:
```bash
go get github.com/sirupsen/logrus
go get golang.org/x/tools/cmd/goimports
```

### IDE / Editors
I use IntelliJ (Goland) for my personal IDE, however, there are many Go IDE's out
there that you can use. The Go Docs offer a live document of common IDE's that
are being used for Go

[Golang Editors](https://golang.org/doc/editors.html)