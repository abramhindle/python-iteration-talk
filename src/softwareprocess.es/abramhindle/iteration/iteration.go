package main

import "fmt"
import "os"
import "bufio"

type Tree struct {
    Value int
    Branches map[string]Tree
}


func TreeSum( tree Tree ) int  {
	sum := tree.Value;
	for _, val := range tree.Branches {
		sum += TreeSum(val)
	}
	return sum;
}

type iInt int

func (max iInt) Iter () <-chan iInt {
	ch := make(chan iInt);
	go func () {
		m := int(max)
		for i := 0; i <= m; i++ {
			ch <- iInt(i)
		}
		close(ch)
	} ();
	return ch
}

func TreeWalker( tree Tree, f func(Tree) )  {
	f(tree)
	for _, val := range tree.Branches {
		TreeWalker(val,f)
	}
}

func TreeIter( tree Tree ) <-chan Tree {
	ch := make(chan Tree);
	go func () {
		TreeWalker( tree, func(t Tree) {
			ch <- t
		})
		close(ch)
	} ();
	return ch
}

func (tree Tree) Iter() <-chan Tree {
	return TreeIter( tree )
}

type StringIterator struct {
	current int
	string s
}



func main() {
	fmt.Printf("Hello, world.\n")
	condition := false

	for condition {
		// ...
		condition = !condition
	}
	reader := bufio.NewReader(os.Stdin)
	reads := 0
	for {
		reads++
		_, err := reader.ReadString('\n');
		if err!=nil {
			break;
		}
	}
	fmt.Printf("Reads: %d\n", reads)

	x := 10
	for x > 0 {
		x = x - 1
	}
	// x is 0
	fmt.Printf("x is %d\n", x)
	
	// maybe you're not sure how many
	// iterations you need?
	y := 100.0
	for y > 1 {
		y = y / 3;
	}
	fmt.Printf("y is %f\n", y)
	// y is 0.41152263374485604

	sum := 0;
	for i := 0 ; i < 10; i++ {
		sum += i;
	}
	fmt.Printf("Sum is %d\n", sum)

	s := ""
	s2 := "❤☀☆☂☻♞☯☭☢"
	v := []string{"a","b","c"}
	u := map[string]string{"A":"a", "B":"b","C":"c"}
	for i, val := range v {
		for j := 0; j <= i; j++ {
			s += val; 
		}
	}
	fmt.Printf("s is [%s]\n",s)
	for key, val := range u {
		s += key;
		s += val; //  over keys
	}
	fmt.Printf("s is [%s]\n",s)
	for i, val := range s2 {
		// i is the byte location
		for j := 0; j <= i; j++ {
			s += string(val);
		}
	}
	fmt.Printf("s is (note that the unique was base-4) [%s]\n",s)

	stump := map[string]Tree{}	

	tree := Tree{0,
		map[string]Tree{
			"a":Tree{1,stump},
			"b":Tree{2, map[string]Tree{
					"h":Tree{8,stump},
					"i":Tree{9,stump},
					"j":Tree{10,stump},
				},
			},
			"c":Tree{3,stump},
			"d":Tree{4, map[string]Tree{
					"e":Tree{5,stump},
					"f":Tree{6,stump},
					"g":Tree{7,stump},
				},
			},
		},
	}	

	fmt.Printf("Treesum %d\n", TreeSum( tree ))

	for i := range iInt(6).Iter() {
		fmt.Printf("Wow! %v\n", i)
	}

	// Call Back walker
	TreeWalker(tree, func(t Tree) { fmt.Printf("Node value %d\n",t.Value) } )

	for tree := range tree.Iter() {
		fmt.Printf("Now via Iter Node value %d\n",tree.Value)
	}
}
