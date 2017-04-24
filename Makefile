
compile:
	go get github.com/gorilla/mux
	go build -o TaskAPI
	
run:
	./TaskAPI &
	
test:
	python tests/main.py