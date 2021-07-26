# Local-Alignment
## Introduction
A Go implementation of the Smith–Waterman algorithm. This program will find the optimal global alignments for two DNA, RNA or amino acid sequences.
# Installation
Install Go
Clone the repository
`$ git clone "https://github.com/Brandonag1997/Local-Alignment"`

Build localalign.go

`$ cd Local-Alignment; go build localalign.go`
## Usage
The program takes at least 2 arguments, the locations of the sequences you want to align. For example

`$ ./localalign -seq1=examples/seq1 -seq2=examples/seq2`

## Options
- `-seq1`
  - The location of a text file containing the first sequence to align
- `-seq2`
  - The location of a text file containing the second sequence to align
- `-similarity_matrix`
  - This optional parameter specifies the similarity matrix if no matrix is specied a simple match/mismatch/gap scoring matrix will be used
  - Available similarity matrices include
    - BLOSUM62 
    - BLOSUM30 increased by 5 up to BLOSUM90
    - BLUSOM100
    - PAM10 increased by 10 up to PAM500
    - NUC.4.2 and NUC.4.4
    - DAYHOFF
    - GONNET
    - IDENTITY
    - MATCH (gives the same results as the default)
- `-match`
  - The score that is assigned to 2 nucleotides or amino acids that match. This value defaults to 2. This is only used when a similarity matrix is not specified.
- `-mismatch`
  - The score that is assigned to 2 nucleotides or amino acids that do not match. This value defaults to -2. This is only used when a similarity matrix is not specified.
- `-gapopen`
  - The score that is assigned to a nucleotide or amino acid that is aligned to a gap opening. This value defaults to -1.
- `-gapextend`
  - The score that is assigned to a nucleotide or amino acid that is aligned to a gap extension. This value defaults to the same value is gapopen.
