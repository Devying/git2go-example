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
  //返回一个 *Reference
  // head, err := Repository.Head()
  // if err != nil {
	//   log.Println(err)
  // }
  // //head.Target() 返回一个*Oid
  // //LookupCommit 返回一个*Commit
  // headCommit, err := Repository.LookupCommit(head.Target())
  // if err != nil {
	//   log.Println(err,headCommit)
  // }
  //

  //我们需要当前的 branch 遍历一下找到当前的分支
  BranchIterator,err := Repository.NewBranchIterator(git.BranchLocal)
  if err != nil{
	  log.Println(err)
  }
  log.Println("BranchIterator is ",BranchIterator)
  var checkBranch *git.Branch
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
    if name == "helloimooc" {
      checkBranch = b
    }
	  return nil
  }

  err = BranchIterator.ForEach(f)
  if err != nil  {
	  log.Println("err is ",err)
  }
  log.Println("Name is",branchs)
  log.Println("checkBranch is",checkBranch)
  

  checkCommit, err := Repository.LookupCommit(checkBranch.Target())
	if err != nil {
		log.Println("Failed to lookup for commit in local branch helloimooc",err)
		
  }
  
  log.Println("checkCommit is",checkCommit)


	defer checkCommit.Free()

	tree, err := Repository.LookupTree(checkCommit.TreeId())
	if err != nil {
		log.Println("Failed to lookup for tree helloimooc",err)
		
	}
	defer tree.Free()

  log.Println("tree is",tree)

  // Checkout the tree
  checkoutOpts := &git.CheckoutOpts{
		Strategy: git.CheckoutForce,
	}
	err = Repository.CheckoutTree(tree, checkoutOpts)
	if err != nil {
		log.Println("Failed to checkout tree helloimooc",err)
		
  }
  err = Repository.SetHead("refs/heads/helloimooc")

  if err != nil {
		log.Println("checkout to  helloimooc",err)
		
  }


  //删除这个分支
  // err = checkBranch.Delete()
  // if err != nil {
	// 	log.Println("Failed to del helloimooc",err)
		
	// }

}