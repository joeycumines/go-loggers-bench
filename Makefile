GOTEST_FLAGS=-cpu=1,2,4 -benchmem -benchtime=5s

# TODO delete this note: (a=Zerolog && make test TEXT_PKGS=Logiface"$a" JSON_PKGS=Logiface"$a" && echo --- && make test TEXT_PKGS="$a" JSON_PKGS="$a")
TEXT_PKGS=Gokit Logrus Log15 Gologging Seelog Zerolog Fortiolog \
		LogifaceZerologJSON LogifaceZerologText LogifaceLogrusJSON \
		LogifaceLogrusText
JSON_PKGS=Gokit Logrus Log15 Zerolog LogifaceZerolog LogifaceZerologJSON \
		LogifaceZerologText LogifaceLogrusJSON LogifaceLogrusText

TEXT_PKG_TARGETS=$(addprefix test-text-,$(TEXT_PKGS))
JSON_PKG_TARGETS=$(addprefix test-json-,$(JSON_PKGS))

GO_BENCH = go test $(GOTEST_FLAGS) -bench

.PHONY: all deps test test-text test-json $(TEXT_PKG_TARGETS) $(JSON_PKG_TARGETS)

all: deps test

deps:
	go get -u -t ./...
	go mod tidy

test: test-text test-json

test-text: $(TEXT_PKG_TARGETS)

$(TEXT_PKG_TARGETS): test-text-%:
	$(GO_BENCH) '^Benchmark$*Text'
	$(GO_BENCH) '^Benchmark$*$$/^Text'

test-json: $(JSON_PKG_TARGETS)

$(JSON_PKG_TARGETS): test-json-%:
	$(GO_BENCH) '^Benchmark$*JSON'
	$(GO_BENCH) '^Benchmark$*$$/^JSON'
