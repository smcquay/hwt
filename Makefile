# from http://blog.jgc.org/2011/07/gnu-make-recursive-wildcard-function.html
rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

default: ${GOPATH}/bin/hwtd ${GOPATH}/bin/hwtc
	
${GOPATH}/bin/hwtd: vendor $(call rwildcard,,*.go) hwt.go
	go install -v mcquay.me/hwt/cmd/hwtd

${GOPATH}/bin/hwtc: vendor $(call rwildcard,,*.go) hwt.go
	go install -v mcquay.me/hwt/cmd/hwtc

hwt.go: service.twirp.go service.pb.go

service.twirp.go: service.proto
	protoc --proto_path=${GOPATH}/src:. --twirp_out=. --go_out=. ./service.proto

service.pb.go: service.proto
	protoc --proto_path=${GOPATH}/src:. --twirp_out=. --go_out=. ./service.proto

vendor: Gopkg.toml Gopkg.lock service.twirp.go service.pb.go
	dep ensure

.PHONY: clean
clean:
	@rm -f service.{twirp,pb}.go
	@rm -rf vendor
