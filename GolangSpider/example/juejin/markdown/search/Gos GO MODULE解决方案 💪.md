# Gos: GO MODULE解决方案 💪 #

( https://link.juejin.im?target=https%3A%2F%2Fcircleci.com%2Fgh%2Fstoryicon%2Fgos%2Ftree%2Fmaster )

![CircleCI](https://user-gold-cdn.xitu.io/2019/5/21/16ada7ce041fb556?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Fcircleci.com%2Fgh%2Fstoryicon%2Fgos%2Ftree%2Fmaster ) ![Go Report Card](https://user-gold-cdn.xitu.io/2019/5/21/16ada7cdee05b95b?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Fgoreportcard.com%2Freport%2Fgithub.com%2Fstoryicon%2Fgos ) ![Build Status](https://user-gold-cdn.xitu.io/2019/5/21/16ada7ce02511d58?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Ftravis-ci.org%2Fstoryicon%2Fgos ) ![GoDoc](https://user-gold-cdn.xitu.io/2019/5/21/16ada7cd8864a39a?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Fgodoc.org%2Fgithub.com%2Fstoryicon%2Fgos ) ![Gitter chat](https://user-gold-cdn.xitu.io/2019/5/21/16ada7cdf1888537?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Fgitter.im%2Fstoryicon%2FLobby )

![gos](https://user-gold-cdn.xitu.io/2019/5/21/16ada7cd955f0ee8?imageView2/0/w/1280/h/960/ignore-error/1)

## Project Address: [github.com/storyicon/g…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fstoryicon%2Fgos ) ##

The current gos is still an alpha version, welcome more people to comment and improve it 🍓, you can add more commands to it, or modify something to make it perform better.

You can download the compiled binary program here: [Release Page]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fstoryicon%2Fgos%2Freleases%2F )

* [Brief introduction]( #brief-introduction )
* [How to start]( #how-to-start )
* [What GOS can do:]( #what-gos-can-do )
- [1. Fully compatible with Go native commands]( #1-fully-compatible-with-go-native-commands )
- [2. Simpler Cross-Compilation]( #2-simpler-cross-compilation )
- [3. Rapid generation of .proto]( #3-rapid-generation-of-proto )
- [4. Go proxy solution]( #4-go-proxy-solution )

## 🦄 Brief introduction ##

from now on, use gos instead of go:

` go get => gos get go build => gos build go run => gos run go ... => gos ... 复制代码`

gos is compatible with all go commands and has go mod/get equipped with smart ` GOPROXY` , it automatically distinguishes between private and public repositories and uses ` GOPROXY` to download your lost package when appropriate.

gos has a few extra commands to enhance your development experience:

` cross agile and fast cross compiling proto quick and easy compilation of proto files 复制代码`

You can use ` -h` on these sub commands to get more information.

## 🐋 How to start ##

This can't be simpler.
According to your system type, download the zip file from the [release page]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fstoryicon%2Fgos%2Freleases%2F ) , unzip, rename the binaries to ` gos` and put it in your ` $PATH`. Then use ` gos` as if you were using the ` go` command.
You can also download the source code and compile it using ` go build -o gos main.go`

Note: The prerequisite for gos to work properly is that the [go binary]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fdl%2F ) is in your $PATH. If you need to use the ` gos proto` command, you need the [protoc binary]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fprotocolbuffers%2Fprotobuf%2Freleases ) too.

## :tangerine: What GOS can do: ##

### 1. Fully compatible with Go native commands ###

You can use ` gos` just like you would with the ` go` command. Compatible with all flags and arguments, such as the following:

` go get -u -v github.com/xxxx/xxxx => gos get -u -v github.com/xxxx/xxxx 复制代码`

### 2. Simpler Cross-Compilation ###

You can use ` gos cross` command for simpler cross-compilation:

` # Compile Linux platform binaries for the current system architecture # For example, if your computer are amd64, it will compile main.go into the binary of linux/amd64 architecture. gos cross main.go linux # Specify the build platform and architecture gos cross main.go linux amd64 gos cross main.go linux arm gos cross main.go linux 386 gos cross main.go windows amd64 gos cross main.go darwin 386 # Compiling binary files for all architectures on the specified platform gos cross main.go linux all gos cross main.go windows all # Compiling binary files for all platforms on the specified architecture gos cross main.go all amd64 # Trying to compile binary files for all platforms and architectures gos cross all all 复制代码`

Gos uses parallel compilation, very fast 🚀, but still depends on the configuration of your operating system.

more information: ` gos cross -h`

### 3. Rapid generation of .proto ###

This feature may only be useful to RPC developers. You can compile proto files more easily, as follows:

` # Compile a single file gos proto helloworld.proto # Compile all proto files under the current folder (excluding subfolders) gos proto all # Compile all proto files in the current directory and all subdirectories gos proto all/all 复制代码`

Of course, the precondition is that you have a [protoc binary]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fprotocolbuffers%2Fprotobuf%2Freleases ) in your $PATH.

more information: ` gos proto -h`

### 4. Go proxy solution ###

There is a dilemma here. If you don't use ` GOPROXY` , there may be a large number of Package pull timeouts (network reasons) or non-existence (repository rename, delete or migrate), like the following:

` unrecognized import path "golang.org/x/net" (https fetch: Get https://golang.org/x/net?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.) 复制代码`

If use ` GOPROXY` , you will not be able to pull the private repositories (github, gitlab, etc) properly, like that:

` go get github.com/your_private_repo: unexpected status (https://athens.azurefd.net/github.com/your_private_repo/@v/list): 500 Internal Server Error 复制代码`

GOS strengthens all of GO's native commands, no matter it's go mod/get/build/run/....Any situation that might cause a package pull, gos will intelligently determine whether the current repository to be pulled needs to use ` GOPROXY`.

**Now, live your thug life 😎**