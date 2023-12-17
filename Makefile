test:
	echo "running test for search data race"
	go test -race 
benchmark:
	echo "running benchmark"
	go test -bench=. -benchtime=10s