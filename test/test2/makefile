.PHONY: all clean

all:
	gogll -p "github.com/goccmack/gogll/test/test2" test2.bnf
	go test -v

clean:
	rm *.txt; \
	rm -rf goutil; \
	rm -rf parser