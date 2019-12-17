package main

import (
	"H-2/P1/product"
	"H-2/P1/profile"
	"H-2/P1/blog"
	"H-2/P1/store"
)

func main() {
	var name = "Dzakiy"
	profile := profile.User{}

	profile.SetProfile(name, "Kamu", "Dimana?")
	profile.GetProfile()


	product := product.New("Aqua", 100)
	product.GetProduct()

	blog := blog.New("Tutorial golang", "oop in", "Linux")
	blog.GetBlog()

	brancStore := store.BranchStore{StoreName :"cabang", OwnerBranch: "Dzakiy"}
	store := store.Store{BranchStore : brancStore, StoreName: "cabang", Owner : "Dzakiy"}

	store.GetBranchStore()
}
