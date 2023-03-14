package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) averageScore() float64 {
	var totalScore int
	for _, val := range s.score {
		totalScore += val
	}
	return float64(totalScore) / float64(len(s.score))
}

func (s Student) minScore() (string, int) {
	minScore := s.score[0]
	minName := s.name[0]
	for i, val := range s.score {
		if val < minScore {
			minScore = val
			minName = s.name[i]
		}
	}
	return minName, minScore
}

func (s Student) maxScore() (string, int) {
	maxScore := s.score[0]
	maxName := s.name[0]
	for i, val := range s.score {
		if val > maxScore {
			maxScore = val
			maxName = s.name[i]
		}
	}
	return maxName, maxScore
}

func main() {
	// inisialisasi slice siswa
	siswa := Student{
		name:  []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"},
		score: []int{90, 75, 70, 75, 50},
	}

	// menampilkan hasil
	fmt.Printf("Average Score : %.2f\n", siswa.averageScore())

	minName, minScore := siswa.minScore()
	fmt.Printf("Min Score of Students : %s (%v)\n", minName, minScore)

	maxName, maxScore := siswa.maxScore()
	fmt.Printf("Max Score of Students : %s (%v)\n", maxName, maxScore)
}
