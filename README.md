# Hrpc

Hrpc copied from Project [twitchtv/twirp](https://github.com/twitchtv/twirp), with rewritten the part of the generated code.

## Usage

Please install the following software first.

- [protoc](https://github.com/protocolbuffers/protobuf/releases)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go/releases)

```bash
go get github.com/cyub/hrpc/cmd/protoc-gen-hrpc # install protoc-gen-hrpc
protoc --hrpc_out=. --go_out=. test.proto # generate hrpc code
```