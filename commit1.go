package main  

import (  
  git "github.com/libgit2/git2go"  
  log "github.com/sirupsen/logrus"
  "time"
) 


func main() {
	signature := &git.Signature{
		Name:  "huangby",
		Email: "huangby@imooc.com",
		When:  time.Now(),
	}
	repo, err := git.OpenRepository("/tmp/code")
	if err != nil{
	  log.Println(err)
	}
	
	head, err := repo.Head()
	if err != nil {
	  log.Println(err)
	}

	headCommit, err := repo.LookupCommit(head.Target())
	if err != nil {
	  log.Println(err)
	}
	log.Println(headCommit)

	idx, err := repo.Index()
	if err != nil {
		log.Println(err)
	}

	treeId,err := idx.WriteTreeTo(repo)
	if err != nil {
	  log.Println(err)
	}
	tree,err := repo.LookupTree(treeId)
	if err != nil {
		log.Println(err)
	}
	message := "now i will write tree to  modifed read.md"
	commitId,err := repo.CreateCommit("HEAD", signature, signature, message, tree, headCommit)
	if err != nil {
	  log.Println(err)
	}
	log.Println(commitId)
	// commitId is log hash value
	
	
}