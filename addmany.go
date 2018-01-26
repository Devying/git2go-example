package main  

import (  
  git "github.com/libgit2/git2go"  
  log "github.com/sirupsen/logrus"
)
func main() {
	repo, err := git.OpenRepository("/tmp/code")
	if err != nil {
		log.Println(err)
		return
	}

	idx, err := repo.Index()
	if err != nil {
		log.Println(err)
		return
	}
	var path []string
	//path := []string{'a.txt','b.txt','c.txt'}
	pathAll := append(path,"1.txt","2.txt","README.md","demo.php","fuck.php")
	log.Println(pathAll)
	err = idx.AddAll(pathAll, git.IndexAddDefault, nil)
	if err != nil {
		log.Println(err)
		return
	}
	err = idx.Write()
	if err != nil {
		log.Println(err)
		return
	}
}