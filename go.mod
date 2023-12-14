module github.com/mikedotexe/dota-go-matchmaking

go 1.20

require (
	github.com/paralin/go-dota2 v0.0.0-20231212221913-75cf2224dbfd
	github.com/paralin/go-steam v0.0.0-20231025185642-e7c8d97e052a
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/paralin/go-steam => github.com/mikedotexe/go-steam v0.0.0-20231214192746-30e6c8764839
