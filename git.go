package main  

import (  
  git "github.com/libgit2/git2go"  
  log "github.com/sirupsen/logrus"
  "time"
)  

// func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {  
//     ret, cred := git.NewCredSshKey("git", "/var/www/testgo/git/id_rsa.pub", "/var/www/testgo/git/id_rsa", "")  
//     return git.ErrorCode(ret), &cred  
// }  

// func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {  
//     return 0  
// }  

func main() {
  
  //Repository 打开本地代码仓库 返回一个 *Repository
  Repository, err := git.OpenRepository("/tmp/code")
  if err != nil{
	  log.Println(err)
  }
  sig := &git.Signature{
		Name:  "huangby",
		Email: "huangby@imooc.com",
		When:  time.Now(),
  }
  message := "This123123123 is a commit\n"

  
  idx, err := Repository.Index()
  if err != nil {
    log.Println(err)
  }
  treeId, err := idx.WriteTree()
  if err != nil {
    log.Println(err)
  }
	tree, err := Repository.LookupTree(treeId)
  
  if err != nil {
    log.Println(err)
  }
	commitId, err := Repository.CreateCommit("HEAD", sig, sig, message, tree)
  if err != nil {
    log.Println(err)
  }
  log.Println(commitId)
  
}