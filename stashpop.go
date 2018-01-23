package main  

import (  
  git "github.com/libgit2/git2go"  
  log "github.com/sirupsen/logrus" 
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

  opts, err := git.DefaultStashApplyOptions()
  if err != nil{
	  log.Println(err)
  }
  err = Repository.Stashes.Pop(1, opts)

  if err != nil{
	  log.Println(err)
  }
}