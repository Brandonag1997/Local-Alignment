package main

import (
	"fmt"
	"os"
)

func readSeq(fileName string) string {
	file, ferr := os.Open(fileName)
	if ferr != nil {
		fmt.Println(ferr)
		return ""
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(fileName,"bytes read:",bytesread)
	return string(buffer)
}

func similarityMatrix(a, b byte) int {
	if a == b{
		return 3
	}else{
		return -3
	}
}

func reverse(str string) (result string) {
	for _, i := range str {
		result = string(i) + result
	}
	return
}

func main() {
	argLength := len(os.Args[1:])
	if argLength < 2 {
		fmt.Println("requires 2 files")
		return
	}
	sequence1 := os.Args[1]
	sequence2 := os.Args[2]
	seq1 := readSeq(sequence1)
	seq2 := readSeq(sequence2)
	fmt.Println(seq1)
	fmt.Println(seq2)
	rows := len(seq2) + 1
	cols := len(seq1) + 1

	Fmatrix := make([][]int, rows)
	for i := range Fmatrix {
		Fmatrix[i] = make([]int, cols)
	}

	highestScore := 0
	hi := 0
	hj := 0
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			M := Fmatrix[i-1][j-1] + similarityMatrix(seq1[j-1],seq2[i-1])
			D := Fmatrix[i-1][j] - 2
			I := Fmatrix[i][j-1] - 2
			if (M > 0 || D > 0 || I > 0) {
				if M > D {
					if M > I {
						Fmatrix[i][j] = M
					} else {
						Fmatrix[i][j] = I
					}
				} else {
					if D > I {
						Fmatrix[i][j] = D
					} else {
						Fmatrix[i][j] = I
					}
				}
				
			} else {
				Fmatrix[i][j] = 0
			} 
			if Fmatrix[i][j] > highestScore {
				highestScore = Fmatrix[i][j]
				hi = i
				hj = j
			}
		}

	}
	// fmt.Println(Fmatrix)
	all1 := ""
	all2 := ""

	//Traceback
	end := false
	for !end {
		M := Fmatrix[hi-1][hj-1]
		G1 := Fmatrix[hi-1][hj]
		G2 := Fmatrix[hi][hj-1]
		if M == 0 {
			all1 += string(seq1[hj-1])
			all2 += string(seq2[hi-1])
			end = true
			break
		}
		if M > G1 {
			if M > G2 {
				all1 += string(seq1[hj-1])
				all2 += string(seq2[hi-1])
				hi -= 1
				hj -= 1
			} else {
				all1 += string(seq1[hj-1])
				all2 += "-"	
				hj -= 1							
			}
		} else {
			if G1 > G2 {
				all1 += "-"
				all2 += string(seq2[hi-1])
				hi -= 1
			} else {
				all1 += string(seq1[hj-1])
				all2 += "-"
				hj -= 1
			}
		}
	}
	alignment1 := reverse(all1)
	alignment2 := reverse(all2)
	fmt.Println("Local Alignment")
	fmt.Println(alignment1)
	fmt.Println(alignment2)
}

