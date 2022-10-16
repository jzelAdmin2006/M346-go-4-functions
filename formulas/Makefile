.PHONY: all clean

all: grade.pdf

.SUFFIXES: .tex .pdf
.tex.pdf:
	xelatex $<

clean:
	rm -f *.pdf *.aux *.log
