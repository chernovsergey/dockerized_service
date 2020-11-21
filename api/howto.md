 - Enter api/ folder
 - export GOPATH=$(go env GOPATH)
 - export PATH=$PATH:$GOPATH/bin
 - export PATH=$PATH:<path_to_protoc>
 
 - go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
 - go get -u google.golang.org/grpc
 - protoc --go_out=plugins=grpc:. *.proto
    
 - go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
 - protoc -I /usr/local/include -I. -I $GOPATH/src -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --grpc-gateway_out=logtostderr=true:./api/ api/api.proto