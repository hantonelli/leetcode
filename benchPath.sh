pushd $(pwd)
cd $(pwd)/${1}
go test -bench=. -benchmem
popd