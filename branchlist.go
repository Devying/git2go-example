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
  BranchIterator,err := Repository.NewBranchIterator(git.BranchLocal)
  if err != nil{
	  log.Println(err)
  }
  log.Println("BranchIterator is ",BranchIterator)
  
  var branchs []map[string]interface{} 
  //函数原型   type BranchIteratorFunc func(*Branch, BranchType) error
  f := func( b *git.Branch,t git.BranchType) error {
	  branch := make(map[string]interface{})
	  name, err := b.Name()
	  if err != nil {
		  return err
	  }
	  branch["name"] = name;
	  branch["head"] = false
	  head, err := b.IsHead()
	  if err != nil {
		  return err
	  }
	  if head == true {
		  branch["head"] = true
	  }
	  branchs = append(branchs, branch)
	  return nil
  }

  err = BranchIterator.ForEach(f)
  if err != nil  {
	  log.Println("err is ",err)
  }
  log.Println("Name is",branchs)
  


  //创建新的分支 三个参数 分支名,目标commit,是否强制创建,
  head, err := Repository.Head()
  if err != nil {
	  log.Println(err)
  }
  
  headCommit, err := Repository.LookupCommit(head.Target())
  if err != nil {
	  log.Println(err)
  }
  
  branchNew, err := Repository.CreateBranch("helloimooc1", headCommit, false)
  if err != nil {
	  log.Println(err)
  }
  log.Println("created new Branch")
  branchNewName,err := branchNew.Name()
  if err != nil {
		log.Println(err)
	}
	log.Println("new branch id ",branchNewName);
	//根据分支名字拿到branch 指针
	branchDel,err := Repository.LookupBranch("helloimooc1",BranchLocal)
	if err != nil {
		log.Println(err)
	}
	log.Println("del branch pointer ",branchDel)

	log.Println("branch helloimooc1 will be deleted")
	err = branchDel.Delete()
	if err != nil {
		log.Println(err)
	}
	log.Println("branch helloimooc1 deleted success")

}