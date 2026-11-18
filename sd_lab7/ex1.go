package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Message reprezinta un mesaj transmis intre noduri in Runda 1 sau Runda 2.
type Message struct {
	SenderID int
	Round    int // 1 pentru starea curenta, 2 pentru valoarea Tie
	Value    int // 0 sau 1
}

type Node struct {
	id          int
	x           int // starea curenta x_i (t)
	s           int
	n           int // numarul de noduri din sistem
	isByzantine bool // comportamentul la iteratia curenta
	inbox       chan Message
}

func (node *Node) Bcast(nodes []*Node, round int, val int) {
	for _, dest := range nodes {
		dest.inbox <- Message{
			SenderID: node.id,
			Round:    round,
			Value:    val,
		}
	}
}

func (node *Node) MaliciousBcastState(nodes []*Node) {
	for _, dest := range nodes {
		dest.inbox <- Message{
			SenderID: node.id,
			Round:    1,
			Value:    rand.Intn(2),
		}
	}
}

func (node *Node) MaliciousBcastTie(nodes []*Node) {
	for _, dest := range nodes {
		dest.inbox <- Message{
			SenderID: node.id,
			Round:    2,
			Value:    rand.Intn(2),
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(firstLine[0])
	s, _ := strconv.Atoi(firstLine[1])

	if n <= 4*s {
		fmt.Println("Conditia n > 4s nu este indeplinita!")
		return
	}

	scanner.Scan()
	valFields := strings.Fields(scanner.Text())

	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		v, _ := strconv.Atoi(valFields[i])
		nodes[i] = &Node{
			id:    i,
			x:     v,
			s:     s,
			n:     n,
			inbox: make(chan Message, n*2),
		}
	}

	// Algoritmul ruleaza s + 2 faze (t = 0 .. s + 1)
	for t := 0; t <= s+1; t++ {
		// Selectia aleatorie a nodurilor malitioase pentru iteratia t
		for i := 0; i < n; i++ {
			nodes[i].isByzantine = false
		}
		f := rand.Intn(s + 1)
		perm := rand.Perm(n)
		for i := 0; i < f; i++ {
			nodes[perm[i]].isByzantine = true
		}

		// Runda 1: Agregarea voturilor
		for _, node := range nodes {
			if node.isByzantine {
				node.MaliciousBcastState(nodes)
			} else {
				node.Bcast(nodes, 1, node.x)
			}
		}

		majority := make([]int, n)
		mult := make([]int, n)

		for i, node := range nodes {
			count0, count1 := 0, 0
			for j := 0; j < n; j++ {
				msg := <-node.inbox
				if msg.Value == 1 {
					count1++
				} else {
					count0++
				}
			}

			if count1 >= (n/2)+1 {
				majority[i] = 1
				mult[i] = count1
			} else {
				majority[i] = 0
				mult[i] = count0
			}
		}

		// Runda 2: Spargerea simetriei prin Leader/King (P_t)
		king := nodes[t]
		if king.isByzantine {
			king.MaliciousBcastTie(nodes)
		} else {
			king.Bcast(nodes, 2, majority[t])
		}

		for i, node := range nodes {
			tieMsg := <-node.inbox
			if !node.isByzantine {
				if float64(mult[i]) > float64(n)/2.0+float64(s) {
					node.x = majority[i]
				} else {
					node.x = tieMsg.Value
				}
			}
		}
	}

	// Afisam valoarea finala a unui nod corect (la ultima iteratie)
	for _, node := range nodes {
		if !node.isByzantine {
			fmt.Printf("Valoarea finala de consens: %d\n", node.x)
			break
		}
	}
} 

