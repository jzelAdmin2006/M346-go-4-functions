.PHONY: all clean

all: grade.png pythagoras.png hypotenuse.png quadratic-formula.png discriminant.png \
	quadratic-formula-discriminant.png celsius-to-fahrenheit.png fahrenheit-to-celsius.png

.SUFFIXES: .pdf .png
.pdf.png:
	convert -density 300 $< $@

.SUFFIXES: .tex .pdf
.tex.pdf:
	xelatex $<

clean:
	rm -f *.pdf *.aux *.log
