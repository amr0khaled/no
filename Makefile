default:
	go build -C build .. && echo "---------" 
	./build/lang ./build/index.no
