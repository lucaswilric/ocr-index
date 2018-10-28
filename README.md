# ocr-index

*An extremely naive search engine for text in images*

- - - - -

This, friends, is purely a learning project. I hoped to learn a few things by
doing this:

* How to start a Go project from scratch
* How to use packages/libraries in Go
* How to do OCR (though basically what I learned here was "Tesseract is OK, I
  guess")
* How a really basic search index works

What I've come out with is a really good idea of what a basic search index
_doesn't_ do:

* Case insensitivity (or rather, getting to choose between
  sensitive/insensitive)
* Full text search (lol)
* Searching for multiple terms
* Ranking, at all
* Spell checking/correction
* Word stemming (e.g. returning results for "running" when I type "run")
* Un-indexing stuff (though I don't really know whether that's a thing
  anywhere - why not just re-index with the unwanted data removed?)

I also got a better idea of Go's opinions about workspaces, how to import
packages, and a little about making apps that have subcommands using
[Cobra](https://github.com/spf13/cobra). So that was nice, and not too bad for a
Sunday afternoon's work.
