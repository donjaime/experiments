package main

import (
       "fmt"
       "sort"
)

func main() {
     simulationLimit := 5
     boardProbabilitiesByRoll := map[int]board{}
     for numRolls := 1; numRolls <= simulationLimit; numRolls++ {
     	 boardProbabilitiesByRoll[numRolls] = buildBoard(numRolls)
	 }

	 // We now have probabilities of being on a given spot after 1,2,3,4,5 rolls.
	 // In order to figure out what spaces you are most likely to have hit at some point, we must
	 // compute the sum of the complement weighted probabilities and then sort them
	 spaceProbabilities := map[int]float64{}
	 for numRolls := 1; numRolls <= simulationLimit; numRolls++ {
	     b := boardProbabilitiesByRoll[numRolls]
	       for _, s := range b {
	       	      spaceProbabilities[s.id] += (1.0 - spaceProbabilities[s.id]) * s.prob
		      			       }
					       }

					       // Let's sort this and see what we have.
					       var b byProb
					       for k, v := range spaceProbabilities {
					       	   b = append(b, space{id: k, prob: v})
						   }
						   sort.Sort(b)
						   fmt.Println(fmt.Sprintf("Most likely spaces (after simulating 5 rolls) ->\n %+v", b))
}

func buildBoard(numRolls int) board {
     hits := map[int]int{}
     permute(0, numRolls, hits)

     // Compute the totally number of possible combinations encountered.
     var b board
     total := 0.0
     for _, v := range hits {
     	 total += float64(v)
	 }
	 // Now sort.
	 for k, v := range hits {
	     b = append(b, space{k, v, float64(v) / total})
	     }
	     sort.Sort(b)
	     return b
}

func permute(prevSpace int, rollsLeft int, hits map[int]int) {
     if rollsLeft == 0 {
     	return
	}
	for i := 1; i <= 6; i++ {
	    hits[prevSpace+i] += 1
	    		      permute(prevSpace+i, rollsLeft-1, hits)
			      }
}

type space struct {
     id   int
     hits int
     prob float64
}

type board []space

func (p board) Len() int           { return len(p) }
func (p board) Less(i, j int) bool { return p[i].hits >= p[j].hits }
func (p board) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type byProb []space

func (p byProb) Len() int           { return len(p) }
func (p byProb) Less(i, j int) bool { return p[i].prob >= p[j].prob }
func (p byProb) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
