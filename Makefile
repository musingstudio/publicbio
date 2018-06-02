GOPATH := ${PWD}:${GOPATH}
export GOPATH

build: ui

install: build-go
	cd less/; $(MAKE) install $(MFLAGS)

ui: 
	cd less/; $(MAKE) $(MFLAGS)

build-go:
	go get -d
	go install ./cmd/publicbio

clean: 
	cd less/; $(MAKE) clean $(MFLAGS)
