.PHONY: clean

all:
	go build -o hymn-sheet-generator main.go
clean:
	rm -rf *.aux *.log *.pdf hymn-sheet-generator