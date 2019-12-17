package store

import "fmt"

type Store struct {
	BranchStore //class store extends
	StoreName   string
	Owner       string
}

type BranchStore struct {
	StoreName   string
	OwnerBranch string
}

func (a Store) GetStore() {
	fmt.Println(a)
}

func (b BranchStore) GetBranchStore() {
	fmt.Println(b)
}
