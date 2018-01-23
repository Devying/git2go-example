package main  

import (  
  git "github.com/libgit2/git2go"  
  log "github.com/sirupsen/logrus" 

)  

func main() {  
  
  //Repository 打开本地代码仓库 返回一个 *Repository
  Repository, err := git.OpenRepository("/tmp/code")
  if err != nil{
	  log.Print(err)
  }
  idx, err := Repository.Index()
  if err != nil{
	  log.Print(err)
  }
  log.Print("idx is ",idx)
  path := idx.Path()
  log.Print("path is ",path)

  //执行 add 的方向操作
  //删除对应于磁盘上文件的索引条目
  err = idx.RemoveByPath("src/Lang/Php.php")
  if err != nil{
	  log.Print(err)
  }
  treeId, err := idx.WriteTree()
  if err != nil {
	  log.Print(err)
  }
  
  err = idx.Write()
  if err != nil {
	  log.Print(err)
  }
  log.Print("treeId is",treeId)
  
  log.Print("remove added smarty_security.php")

}  
