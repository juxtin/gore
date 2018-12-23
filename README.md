# GORE

The GOlang Repository Explorer will splay out your program's internals for the
world to see.


## WIP

This is nowhere near done, it barely does anything, and what it does it probably
does poorly. Use at your own risk, or better yet not at all.

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
