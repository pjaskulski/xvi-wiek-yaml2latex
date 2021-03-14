go build -o yaml2latex ./cmd/yaml2latex
mv yaml2latex ../xvi-wiek/
cd ../xvi-wiek/
./yaml2latex
mv xvi-wiek.tex ../tufte-latex/
cd ../tufte-latex/
pdflatex xvi-wiek.tex
makeindex xvi-wiek.tex
pdflatex xvi-wiek.tex