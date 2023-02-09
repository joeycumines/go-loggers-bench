GOTEST_FLAGS=-cpu=1,2,4 -benchmem -benchtime=5s

# TODO delete this note: (a=Zerolog && make test TEXT_PKGS=Logiface"$a" JSON_PKGS=Logiface"$a" && echo --- && make test TEXT_PKGS="$a" JSON_PKGS="$a")
TEXT_PKGS=Gokit Logrus Log15 Gologging Seelog Zerolog Fortiolog \
		LogifaceZerolog
JSON_PKGS=Gokit Logrus Log15 Zerolog LogifaceZerolog

TEXT_PKG_TARGETS=$(addprefix test-text-,$(TEXT_PKGS))
JSON_PKG_TARGETS=$(addprefix test-json-,$(JSON_PKGS))

.PHONY: all deps test test-text test-json $(TEXT_PKG_TARGETS) $(JSON_PKG_TARGETS)

all: deps test

deps:
	go get -u -t ./...
	go mod tidy

test: test-text test-json

test-text: $(TEXT_PKG_TARGETS)

$(TEXT_PKG_TARGETS): test-text-%:
	go test $(GOTEST_FLAGS) -bench "^Benchmark$*.*Text"

test-json: $(JSON_PKG_TARGETS)

$(JSON_PKG_TARGETS): test-json-%:
	go test $(GOTEST_FLAGS) -bench "^Benchmark$*.*JSON"
