SOURCES=$(wildcard *.go)
EXEC=srctrace
BINARIES=../bin
all: $(EXEC)

$(EXEC):
	go build -o $(BINARIES)/$(EXEC)

clean:
	$(RM) $(BINARIES)/$(EXEC)
	
dependencies:
	go get -u -v github.com/akamensky/argparse