package electionday

import "fmt"

func main() {
	var initialVotes int
	initialVotes = 2

	var counter *int
	counter = NewVoteCounter(initialVotes)
	fmt.Println(*counter)

	var votes int
	votes = 3

	var voteCounter *int
	voteCounter = &votes

	fmt.Println(VoteCount(voteCounter))
	// => 3

	var nilVoteCounter *int
	fmt.Println(VoteCount(nilVoteCounter))
	// => 0

	IncrementVoteCount(voteCounter, 2)

	fmt.Printf("El valor debe de ser (5): %v and %v\n", votes, *voteCounter) // == 5          // true

	var result *ElectionResult
	result = NewElectionResult("Peter", 3)

	fmt.Printf("Name %v, Votes %v\n", result.Name, result.Votes) //== "Peter" // true  == 3      // true

	var result1 *ElectionResult
	result1 = &ElectionResult{
		Name:  "John",
		Votes: 32,
	}

	fmt.Println(DisplayResult(result1))

	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}

	DecrementVotesOfCandidate(finalResults, "Mary")

	fmt.Println(finalResults["Mary"])

}

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) (voteCounts int) {
	if counter != nil {
		voteCounts = *counter
	}
	return voteCounts
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	if counter != nil {
		*counter += increment
	}
}

type ElectionResult struct {
	Name  string // Name of the candidate
	Votes int    // Number of votes the candidate had
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{Name: candidateName, Votes: votes}
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%v (%v)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}
