package main

import (
	"fmt"
	"os"
	"flag"
	"path/filepath"
	"bufio"
	"strings"
	"strconv"
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

func reverse(str string) (result string) {
	for _, i := range str {
		result = string(i) + result
	}
	return
}

func main() {
	//parameters
	var sequence1, sequence2 string
	flag.StringVar(&sequence1, "seq1", "", "Sequence 1")
	flag.StringVar(&sequence2, "seq2", "", "Sequence 2")
	gopen := flag.Int("gapopen", 2, "Penalty for opening a gap")
	gextend := flag.Int("gapextend", -1, "Penalty for extending a gap, default is linear")
	match := flag.Int("match", 3, "Score for a match")
	mismatch := flag.Int("mismatch", -3, "Score for a mismatch")
	smat := flag.String("similarity_matrix", "", "similarity matrix")
	flag.Parse()

	gapopen := *gopen
	gapextend := *gextend
	simmat := *smat
	matchScore := *match
	mismatchScore := *mismatch
	linear := false
	if gapextend == -1 {
		linear = true
	}

	//map to hold the specified similarity matrix (if any)
	type Key struct {
		X, Y string
	}
	m := map[Key]int{}

	// if a similarity matrix is chosen, open it and add the values to the map
	if len(simmat) != 0 {
		dirname, _ := os.Getwd()
		file, err := os.Open(filepath.Join(dirname,"matrices",simmat))
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(file)
		var ls []string
		for scanner.Scan() {
			if len(scanner.Text()) == 0 {
				continue
			} else if string(scanner.Text()[0]) == "#"{
				continue
			} else if string(scanner.Text()[0]) == " " {
				ls = strings.Fields(scanner.Text())
				continue
			} else {
				vals := strings.Fields(scanner.Text())
				if len(vals) == 0 {
					continue
				}
				letter := vals[0]
				for index,  item := range ls {
					val, err := strconv.Atoi(vals[index+1])
					if err != nil {
						fmt.Println(err)
					}
					m[Key{letter, item}] = val
				}
			}
		}
		file.Close()
	}
	
	//if a similarity matrix was specified use the map, if not then use match/mismatch scores
	similarityMatrix := func (a, b, name string) int {
		if len(name) == 0 {
			if a == b{
				return matchScore
			}else{
				return mismatchScore
			}
		} else {
			return m[Key{a, b}]
		}
		 
	}


	seq1 := readSeq(sequence1)
	seq2 := readSeq(sequence2)
	fmt.Println("Sequence 1")
	fmt.Println(seq1)
	fmt.Println("Sequence 2")
	fmt.Println(seq2)

	//initialize Fmatrix
	rows := len(seq2) + 1
	cols := len(seq1) + 1
	Fmatrix := make([][]int, rows)
	for i := range Fmatrix {
		Fmatrix[i] = make([]int, cols)
	}

	//variables to keep track of the highest score for traceback
	highestScore := 0
	hi := 0
	hj := 0

	//fill in Fmatrix
	gepLength := 0
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			M := Fmatrix[i-1][j-1] + similarityMatrix(string(seq1[j-1]),string(seq2[i-1]),simmat)
			var D, I int
			if linear {
				D = Fmatrix[i-1][j] - gapopen
				I = Fmatrix[i][j-1] - gapopen
			} else {
				D = Fmatrix[i-1][j] - (gapextend * gepLength) - gapopen
				I = Fmatrix[i][j-1] - (gapextend * gepLength) - gapopen
			}

			if (M > 0 || D > 0 || I > 0) {
				if M > D {
					if M > I {
						Fmatrix[i][j] = M
						gepLength = 0
					} else {
						Fmatrix[i][j] = I
						gepLength += 1
					}
				} else {
					if D > I {
						Fmatrix[i][j] = D
						gepLength += 1
					} else {
						Fmatrix[i][j] = I
						gepLength += 1
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

