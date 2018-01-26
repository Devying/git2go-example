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
	remoteCollection := repo.Remotes
	//remoteList,err := remoteCollection.List()//可以获取远程库的列表
	remote,err := remoteCollection.Lookup("origin")
	if err != nil{
		log.Println(err)
	}
	//log.Println(remoteList)
	fetchOptions := &git.FetchOptions {}
	err = remote.Fetch([]string{}, fetchOptions , "")
	if err != nil{
		log.Println(err)
	}

}