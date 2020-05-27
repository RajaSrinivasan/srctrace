SOURCES=$(wildcard *.go)
EXEC=srctrace
BINARIES=./bin
all: 
	go build -o $(BINARIES)/$(EXEC)
	GOOS=linux GOARCH=amd64 GOARM=6 go build -o $(BINARIES)/linux64/$(EXEC) $(SOURCES)
	GOOS=windows GOARCH=amd64 GOARM=6 go build -o $(BINARIES)/win64/$(EXEC).exe $(SOURCES)

clean:
	-$(RM) $(BINARIES)/$(EXEC)
	-$(RM) $(BINARIES)/linux64/$(EXEC)
	-$(RM) $(BINARIES)/win64/$(EXEC).exe

generate:	
ifneq ("$(wildcard $(BINARIES)/$(EXEC))","")
	cd tests; ../$(BINARIES)/$(EXEC) --language go --output versions; mv versions.go ../versions/genversions.go
endif

dependencies:
	go get -u -v github.com/akamensky/argparse