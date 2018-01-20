# from http://blog.jgc.org/2011/07/gnu-make-recursive-wildcard-function.html
rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

default: ${GOPATH}/bin/hwtd ${GOPATH}/bin/hwtc
	
${GOPATH}/bin/hwtd: $(call rwildcard,,*.go) hwt.go vendor
	@go install -v mcquay.me/hwt/cmd/hwtd

${GOPATH}/bin/hwtc: $(call rwildcard,,*.go) hwt.go vendor
	@go install -v mcquay.me/hwt/cmd/hwtc

hwt.go: rpc/hwt/service.twirp.go rpc/hwt/service.pb.go

rpc/hwt/service.twirp.go: rpc/hwt/service.proto
	@echo "generating twirp file"
	@protoc --proto_path=${GOPATH}/src:. --twirp_out=. --go_out=. rpc/hwt/service.proto

rpc/hwt/service.pb.go: rpc/hwt/service.proto
	@echo "generating pb file"
	@protoc --proto_path=${GOPATH}/src:. --twirp_out=. --go_out=. rpc/hwt/service.proto

vendor: Gopkg.toml Gopkg.lock
	dep ensure

.PHONY: clean
clean:
	@rm -f rpc/hwt/service.{twirp,pb}.go
	@rm -rf vendor
