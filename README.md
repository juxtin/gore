# GORE

The GOlang Repository Explorer will splay out your program's internals for the
world to see, exposing inconvenient truths and telling the story of your code.


## WIP

This is currently POC/tech demo-quality software. Use at your own risk, if at
all. That said, neither the program itself nor the accompanying scripts do
anything particularly dangerous like delete files, so it should be perfectly
safe when used as directed.

## Usage

Clone this repo, then open a terminal in the project root and run
`script/bootstrap` followed by `script/build`. That should leave a new
executable at `bin/gore`.

The easiest way to test it out is to run this project on itself with `bin/gore
.`, then paste the output into [Webgraphviz](http://www.webgraphviz.com/) to
visualize it.

Note that `gore` currently just spits its results to stdout, so if you'd like to
save them to a file then please use shell redirection like so: `bin/gore . >
gore.dot`.

## purpose and rationale

This is primarily a hack project for me to get to know Go in two ways:

1. by actually writing a Go program
2. by giving me a tool to explore a Go codebase in a way that makes sense to me.

The former is self-explanatory, so I'll describe the latter. When I'm
approaching a new codebase, I often have a difficult time deciding which files
to read and in which order. I end up with the same questions over and over
again:

* which file has the main program logic?
* which files have commonly-used abstractions, types, utility functions, etc?
* if I wanted to read every function in implementation order, where would I
  start?

What I often want, or at least think I want, is a
[DAG](https://en.wikipedia.org/wiki/Directed_acyclic_graph) showing all the
source files in the project and how they relate to one another. That way, I can
choose to start reading, say, a/the file that *no other file* requires (probably
main or something). I can also see the project at a glance to get an idea of the
high-level architecture.

GORE is my first attempt to implement such a tool.
