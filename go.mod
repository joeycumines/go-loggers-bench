module github.com/imkira/go-loggers-bench

go 1.19

replace (
	github.com/joeycumines/go-utilpkg/logiface => ./../go-utilpkg/logiface
	github.com/joeycumines/go-utilpkg/logiface/logrus => ./../go-utilpkg/logiface/logrus
	github.com/joeycumines/go-utilpkg/logiface/zerolog => ./../go-utilpkg/logiface/zerolog
)

require (
	fortio.org/fortio v1.40.1
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/go-kit/kit v0.12.0
	github.com/joeycumines/go-utilpkg/logiface v0.0.0-20230209142644-7492764539e0
	github.com/joeycumines/go-utilpkg/logiface/logrus v0.0.0-20230211053058-0f1429ef7ffc
	github.com/joeycumines/go-utilpkg/logiface/zerolog v0.0.0-20230209142644-7492764539e0
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/rs/zerolog v1.29.0
	github.com/sirupsen/logrus v1.9.0
	gopkg.in/inconshreveable/log15.v2 v2.16.0
)

require (
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	golang.org/x/exp v0.0.0-20230210204819-062eb4c674ab // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/term v0.5.0 // indirect
)
