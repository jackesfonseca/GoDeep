package geneticextractor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func score(s1 []string, s2 []string, match int, mismatch int, gap int, i int, j int) int {
	var score = 0
	if s1[i-1] != "_" && s2[j-1] != "_" {
		if s1[i-1] == s2[j-1] {
			score = match
		} else {
			score = mismatch
		}
	} else {
		score = gap
	}
	return score
}
func max3(a int, b int, c int) int {
	if b > c {
		if b > a {
			return b
		}
		return a
	} else {
		if c > a {
			return c
		}
		return a
	}
}
func getF(s1 string, s2 string, match int, mismatch int, gap int) [][]int {
	var r1 = s1
	var r2 = s2
	s := strings.Split(r1, "")
	t := strings.Split(r2, "")
	d := gap

	F := make([][]int, len(r1)+10)
	for i := 0; i < len(r1)+10; i++ {
		F[i] = make([]int, len(r2)+10)
	}

	F[0][0] = 0
	if len(r1) == len(r2) {
		for i := 0; i <= len(r1); i++ {
			F[i][0] = d * i
		}
		for j := 0; j <= len(r2)+1; j++ {
			F[0][j] = d * j
		}
	}

	var c = 0
	for i := 1; i <= len(r2); i++ {
		c++
		for j := 1; j <= len(r1); j++ {
			Choice1 := F[i-1][j-1] + score(s, t, 1, -1, -1, j, c)
			Choice2 := F[i-1][j] + d
			Choice3 := F[i][j-1] + d
			F[i][j] = max3(Choice1, Choice2, Choice3)
		}
	}
	return F
}

func get_score(s1 string, s2 string, match int, mismatch int, gap int) int {
	var r1 = s1
	var r2 = s2
	s := strings.Split(r1, "")
	t := strings.Split(r2, "")
	d := gap

	F := make([][]int, len(r1)+10)
	for i := 0; i < len(r1)+10; i++ {
		F[i] = make([]int, len(r2)+10)
	}

	F[0][0] = 0
	if len(r1) == len(r2) {
		for i := 0; i <= len(r1); i++ {
			F[i][0] = d * i
		}
		for j := 0; j <= len(r2)+1; j++ {
			F[0][j] = d * j
		}
	}

	var c = 0
	for i := 1; i <= len(r2); i++ {
		c++
		for j := 1; j <= len(r1); j++ {
			Choice1 := F[i-1][j-1] + score(s, t, 1, -1, -1, j, c)
			Choice2 := F[i-1][j] + d
			Choice3 := F[i][j-1] + d
			F[i][j] = max3(Choice1, Choice2, Choice3)
		}
	}

	if r1 == r2 {
		return F[len(r1)][len(r2)]
	} else if r2 > r1 {
		return F[len(r2)][len(r1)]
	} else {
		return F[len(r1)][len(r2)]
	}
}

func Alimented(A string, B string, gap int) (string, string) {
	var AlignmentA = ""
	var AlignmentB = ""
	var i int = len(A)
	var j int = len(B)
	var r1 = A
	var r2 = B
	s := strings.Split(r1, "")
	t := strings.Split(r2, "")
	d := gap
	var F [][]int = getF(A, B, 1, -1, -2)
	for i > 0 && j > 0 {
		var Score int = F[i][j]
		var ScoreDiag int = F[i-1][j-1]
		var ScoreUp int = F[i][j-1]
		var ScoreLeft int = F[i-1][j]
		if Score == ScoreDiag+score(s, t, 1, -1, -1, i, j) {
			AlignmentA = s[i-1] + AlignmentA
			AlignmentB = t[j-1] + AlignmentB
			i = i - 1
			j = j - 1
		} else if Score == ScoreLeft+d {
			AlignmentA = s[i-1] + AlignmentA
			AlignmentB = "-" + AlignmentB
			i = i - 1
		} else if Score == ScoreUp+d {
			AlignmentA = "-" + AlignmentA
			AlignmentB = t[j-1] + AlignmentB
			j = j - 1
		}
	}
	for i > 0 {
		AlignmentA = s[i-1] + AlignmentA
		AlignmentB = "-" + AlignmentB
		i = i - 1
	}
	for j > 0 {
		AlignmentA = "-" + AlignmentA
		AlignmentB = t[j-1] + AlignmentB
		j = j - 1
	}
	return AlignmentA, AlignmentB
}
