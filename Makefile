TALKSOURCE = wtf_is_cf_and_bosh.md

.PHONY: pdf
pdf: $(TALKSOURCE)
	pandoc -s -t beamer -H preamble.tex $(TALKSOURCE) -o wtf_is_cf_and_bosh.pdf

.PHONY: clean
clean:
	rm *.pdf
