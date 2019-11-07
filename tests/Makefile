all:
	../../bin/srctrace --language Ada
	gnatmake test
	./test
	../../bin/srctrace --language C -m 3 -n 4 -b 292
	gcc ctest.c -o ctest
	./ctest
clean:
	$(RM) revisions.ads
	$(RM) revisions.h