## Flyer Project

#### How to run
1. Install go
2. Set your $GOPATH
3. Clone project outside of $GOPATH

```sh
cd WORKING_DIRECTORY
git clone github.com/hahooh/flyerAPI
cd flyerAPI
ln -s WORKING_DIRECTORY/flyerAPI $GOPATH/github/hahooh/fleryAPI
go build ./...
go build
go ./flyerAPI
```