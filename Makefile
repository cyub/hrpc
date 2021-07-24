install:
	cd cmd/protoc-gen-hrpc && go install

test-protoc:
	/usr/bin/find testdata/proto/ -name '*.proto' -exec bash -c 'protoc --hrpc_out=`dirname {}` --go_out=`dirname {}` {}' \;

clean:
	rm -rf testdata/proto/*.go