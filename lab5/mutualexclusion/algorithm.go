package mutualexclusion

import (
	"container/list"
	"fmt"
	"math"
)

type path struct {
	state []int
	perm  []int
	step  int
}

func getStateId(state []int) (idx int) {
	n := len(state)

	for i := n - 1; i >= 0; i-- {
		idx += state[i] * int(math.Pow(float64(n+1), float64(n-i-1)))
	}

	return idx
}

func getNextState(state []int) []int {
	n := len(state)

	isComplete := true
	for _, v := range state {
		if v != n {
			isComplete = false
			break
		}
	}

	if isComplete {
		return nil
	}

	for i := range state {
		state[i] = (state[i] + 1) % (n + 1)

		if state[i] != 0 {
			break
		}
	}

	return state
}

func isStable(state []int) bool {
	for i := 0; i < len(state)-1; i++ {
		if state[i] != state[i+1] {
			return false
		}
	}

	return true
}

func mutate(id int, result []int) {
	n := len(result)

	// first processor with distinct action
	if id == 0 {
		if result[0] == result[n-1] {
			result[0] = (result[n-1] + 1) % (n + 1)
		}
		return
	}

	if result[id] != result[id-1] {
		result[id] = result[id-1]
	}
}

func getNextPerm(perm []int) {
	n := len(perm)
	for i := n - 1; i >= 0; i-- {
		if i == 0 || perm[i] < n-i-1 {
			perm[i]++
			return
		}

		perm[i] = 0
	}
}

func getPerm(rng, perm []int) (result []int) {
	result = copyArray(rng)

	for i, v := range perm {
		result[i], result[i+v] = result[i+v], result[i]
	}

	return result
}

func buildPerms(n int) (arr [][]int) {
	rng := getRange(0, n-1)
	arr = make([][]int, fact(n))

	idx := 0
	for perm := make([]int, len(rng)); perm[0] < len(perm); getNextPerm(perm) {
		arr[idx] = getPerm(rng, perm)
		idx++
	}

	return arr
}

func areStatesEqual(state1, state2 []int) bool {
	if len(state1) != len(state2) {
		return false
	}

	for i := range state1 {
		if state1[i] != state2[i] {
			return false
		}
	}

	return true
}

func visit(curr, visited []int, value int) {
	id := getStateId(curr)
	visited[id] = int(math.Max(float64(visited[id]), float64(value)))
}

func isVisited(curr, visited []int) bool {
	id := getStateId(curr)
	return visited[id] != -1
}

func process(originState, visitedStates []int, perms [][]int, steps, n int) {
	queue := list.New()
	for _, perm := range perms {
		queue.PushBack(path{copyArray(originState), perm, steps})
	}

	for queue.Len() > 0 {
		next := queue.Front()
		currPath := next.Value.(path)

		flag := true
		for _, id := range currPath.perm {
			if isStable(currPath.state) {
				visit(originState, visitedStates, currPath.step)
				flag = false
				break
			}

			if !areStatesEqual(currPath.state, originState) && isVisited(currPath.state, visitedStates) {
				id := getStateId(currPath.state)
				visit(originState, visitedStates, currPath.step+visitedStates[id])
				flag = false
				break
			}

			prevState := copyArray(currPath.state)
			mutate(id, currPath.state)

			if !areStatesEqual(prevState, currPath.state) {
				currPath.step++
			}
		}

		if flag {
			for _, perm := range perms {
				queue.PushBack(path{copyArray(currPath.state), perm, currPath.step})
			}
		}

		queue.Remove(next)
	}
}

func Simulate(n int) {
	nf := float64(n)
	statesCount := int(math.Pow(nf+1., nf))

	visitedStates := make([]int, statesCount)
	for i := range visitedStates {
		visitedStates[i] = -1
	}

	perms := buildPerms(n)
	for state := make([]int, n); state != nil; state = getNextState(state) {
		process(copyArray(state), visitedStates, perms, 0, n)
	}

	fmt.Printf("Longest path: %d\n", max(visitedStates))
}
