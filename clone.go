package main  

import (  
  git "github.com/libgit2/git2go"  
  log "github.com/sirupsen/logrus" 

)

func main(){
	cloneOptions := &git.CloneOptions{}
    repo, err := git.Clone("https://github.com/Devying/test.git", "/tmp/code", cloneOptions)
            if err != nil {
                    log.Println(err)
            }
    log.Println(repo)
}