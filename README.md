# Local-Alignment
## Introduction
A Go implementation of the Smith–Waterman algorithm. This program will find the optimal global alignments for two DNA, RNA or amino acid sequences.
# Installation
Install Go

`$ git clone "https://github.com/Brandonag1997/Local-Alignment"`

Build localalign.go

`$ cd Local-Alignment; go build localalign.go`
## Usage
The program takes at least 2 arguments, the locations of the sequences you want to align. For example

`$ ./localalign -seq1=examples/seq1 -seq2=examples/seq2`
