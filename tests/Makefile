all:
	../../bin/srctrace --language Ada
	gnatmake test
	./test
clean:
	$(RM) revisions.ads