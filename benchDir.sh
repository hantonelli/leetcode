pushd $(pwd)
go test -bench=. -benchmem
popd