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
	config, err := repo.Config()
	if err != nil{
		log.Println(err)
	}
	err = config.SetString("user.name","imooc")
	// commitId is log hash value
	if err != nil{
		log.Println(err)
	}

	err = config.SetString("remote.origin.url","https://github.com/Devying/test.git")
	// commitId is log hash value
	if err != nil{
		log.Println(err)
	}
	
}