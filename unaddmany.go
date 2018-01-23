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

	//执行 add 的反向操作
    files := []string{"smarty_resource_custom.php","smarty_resource_recompiled.php","smarty_resource_uncompiled.php"}
	err = idx.RemoveAll( files,nil)
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
	
	log.Print("unadd 3files")

}  
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
  files := []string{"smarty_resource_custom.php","smarty_resource_recompiled.php","smarty_resource_uncompiled.php"}
  err = idx.RemoveAll( files,nil)
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
  
  log.Print("unadd 3files")

}  
