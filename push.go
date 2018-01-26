package main  

import (  
  git "github.com/libgit2/git2go"  
  log "github.com/sirupsen/logrus"
) 


func main() {
	repo, err := git.OpenRepository("/tmp/code")
	if err != nil{
	  log.Println(err)
	  return
	}
	remote, err := repo.Remotes.Lookup("origin")
	if err != nil {
		remote, err = repo.Remotes.Create("origin", repo.Path())
		if err != nil {
			log.Println(err)
			return
		}
	}

	// ref
	ref, err := repo.Head()
	if err != nil {
		log.Println(err)
		return
	}

	// Get the name
	branch := ref.Branch()
	branchName, err := branch.Name()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(branchName)
	
	if err := remote.Push([]string{"refs/heads/"+branchName}, &git.PushOptions{}); err != nil {
		log.Println(err)
		return
	}
}