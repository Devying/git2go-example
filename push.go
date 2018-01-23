package main  

import (  
  git "github.com/libgit2/git2go"  
  log "github.com/sirupsen/logrus"
) 


func main() {
	repo, err := git.OpenRepository("/tmp/code")
	if err != nil{
	  log.Println(err)
	}
	remote, err := repo.Remotes.Lookup("origin")
	if err != nil {
		remote, err = repo.Remotes.Create("origin", repo.Path())
		if err != nil {
			log.Println(err)
		}
	}

	// Get the branch
	branch, err := repo.Branch()
	if err != nil {
		log.Println(err)
	}

	// Get the name
	branchName, err := branch.Name()
	if err != nil {
		log.Println(err)
	}

	if err := remote.Push([]string{"refs/heads/"+branchName}, &git.PushOptions{}); err != nil {
		log.Println(err)
	}
}