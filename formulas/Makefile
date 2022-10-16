.PHONY: all clean

all: grade.pdf pythagoras.pdf hypotenuse.pdf quadratic-formula.pdf discriminant.pdf \
	quadratic-formula-discriminant.pdf celsius-to-fahrenheit.pdf fahrenheit-to-celsius.pdf

.SUFFIXES: .tex .pdf
.tex.pdf:
	xelatex $<

clean:
	rm -f *.pdf *.aux *.log
