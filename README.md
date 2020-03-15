Description
===================
Transform json string to yaml string without the type infomation.

Features
====================
* zero config.
* supports Windows, Linux, macOS, FreeBSD, NetBSD, OpenBSD, Plan 9 etc..

Binary installation and usage
====================
* Download a binary that match your operation system and platform.
* https://github.com/bronze1man/json2yaml/releases
* copy it to `/usr/local/bin` like (`cp ~/Downloads/json2yaml_darwin_amd64 /usr/local/bin/json2yaml` )
* Use `chmod +x /usr/local/bin/json2yaml` give running permission to it.

### mac/linux usage
* `echo '{"a": 1}' | json2yaml`
* `json2yaml < 1.json > 2.yaml`

### window usage
* windows 7 cmd.exe:
```
C:\tmp>more .\1.yaml
a: 1

C:\tmp>.\json2yaml_windows_amd64.exe < ./1.yaml > 2.json

C:\tmp>more .\2.json
{"a":1}
```

* windows 7 powerShell 6.1.7600 file example:
```
PS C:\tmp> more .\1.json
{"a":1}

PS C:\tmp> Get-Content .\1.json | .\json2yaml_windows_amd64.exe > .\2.yaml
PS C:\tmp> more .\2.yaml
a: 1
```

Library installation
====================
* `go get -v github.com/bronze1man/json2yaml/j2yLib`


Development
==================
* Follow example is for develop on mac and linux. It should work to windows too. (need change some command line and path to work on windows).
* You need a golang on your computer. https://golang.org
* Create a new directory as your workspace, like `~/work/json2yaml`
* Change your current work directory to that directory.

```
cd ~/work/json2yaml
GOPATH=`pwd` go get -v github.com/bronze1man/json2yaml
GOPATH=`pwd` go run github.com/bronze1man/json2yaml/j2yBuilder
```
* use the file at $GOPATH/tmp/file to distribute your binary to others.



Notice
=====================
* if you don't know whether your platform is 386 or amd64, use the 386 build...

Reference
====================
* https://github.com/peter-edge/go-json2yaml
* https://github.com/go-yaml/yaml
* https://github.com/bronze1man/json2yaml