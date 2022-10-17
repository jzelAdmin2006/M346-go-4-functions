.PHONY: all clean

all: grade.png pythagoras.png hypotenuse.png quadratic-formula.png discriminant.png \
	quadratic-formula-discriminant.png celsius-to-fahrenheit.png fahrenheit-to-celsius.png

.SUFFIXES: .pdf .png
.pdf.png:
	convert -density 200 -background white -flatten $< $@

.SUFFIXES: .tex .pdf
.tex.pdf:
	xelatex $<

clean:
	rm -f *.pdf *.aux *.log
