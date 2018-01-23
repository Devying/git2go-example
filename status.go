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
  Repository, err := git.OpenRepository("/tmp/code");
  if err != nil{
	  log.Print(err)
  }
  //

  // type StatusOpt int
  
  // const (
  // 	StatusOptIncludeUntracked      StatusOpt = C.GIT_STATUS_OPT_INCLUDE_UNTRACKED
  // 	StatusOptIncludeIgnored        StatusOpt = C.GIT_STATUS_OPT_INCLUDE_IGNORED
  // 	StatusOptIncludeUnmodified     StatusOpt = C.GIT_STATUS_OPT_INCLUDE_UNMODIFIED
  // 	StatusOptExcludeSubmodules     StatusOpt = C.GIT_STATUS_OPT_EXCLUDE_SUBMODULES
  // 	StatusOptRecurseUntrackedDirs  StatusOpt = C.GIT_STATUS_OPT_RECURSE_UNTRACKED_DIRS
  // 	StatusOptDisablePathspecMatch  StatusOpt = C.GIT_STATUS_OPT_DISABLE_PATHSPEC_MATCH
  // 	StatusOptRecurseIgnoredDirs    StatusOpt = C.GIT_STATUS_OPT_RECURSE_IGNORED_DIRS
  // 	StatusOptRenamesHeadToIndex    StatusOpt = C.GIT_STATUS_OPT_RENAMES_HEAD_TO_INDEX
  // 	StatusOptRenamesIndexToWorkdir StatusOpt = C.GIT_STATUS_OPT_RENAMES_INDEX_TO_WORKDIR
  // 	StatusOptSortCaseSensitively   StatusOpt = C.GIT_STATUS_OPT_SORT_CASE_SENSITIVELY
  // 	StatusOptSortCaseInsensitively StatusOpt = C.GIT_STATUS_OPT_SORT_CASE_INSENSITIVELY
  // 	StatusOptRenamesFromRewrites   StatusOpt = C.GIT_STATUS_OPT_RENAMES_FROM_REWRITES
  // 	StatusOptNoRefresh             StatusOpt = C.GIT_STATUS_OPT_NO_REFRESH
  // 	StatusOptUpdateIndex           StatusOpt = C.GIT_STATUS_OPT_UPDATE_INDEX
  // )
  
  // type StatusShow int
  
  // const (
  // 	StatusShowIndexAndWorkdir StatusShow = C.GIT_STATUS_SHOW_INDEX_AND_WORKDIR
  // 	StatusShowIndexOnly       StatusShow = C.GIT_STATUS_SHOW_INDEX_ONLY
  // 	StatusShowWorkdirOnly     StatusShow = C.GIT_STATUS_SHOW_WORKDIR_ONLY
  // )
  
  // type StatusOptions struct {
  // 	Show     StatusShow
  // 	Flags    StatusOpt
  // 	Pathspec []string
  // }
  StatusOptions:=&git.StatusOptions{Show:git.StatusShowIndexAndWorkdir,Flags:git.StatusOptIncludeUntracked};
  StatusList, err := Repository.StatusList(StatusOptions)  
  if err != nil {  
		  log.Print(err)  
  }

  //修改的文件数量
  count, err:=StatusList.EntryCount()
  if err != nil {  
	  log.Print(err)
  }
  log.Print("the cout is ",count)

  for i := 0;i < count; i++ {

	  entry, err := StatusList.ByIndex(i)
	  if err != nil{
		  log.Println(err)
	  }
	  var S string
	  var file git.DiffFile = entry.IndexToWorkdir.OldFile
	  if entry.Status == git.StatusIndexNew {
		  //git add 以后就要去索引区了
		  file = entry.IndexToWorkdir.NewFile
	  }	
	  if entry.Status == git.StatusWtNew {
		  S = "A" 
	  } else if entry.Status == git.StatusWtModified {
		  S = "M"
	  } else if entry.Status == git.StatusWtDeleted {
		  S = "D"
	  } else {
		  S = "?"
	  }
	  log.Printf("%s %s %d %d",S,file.Path,file.Flags,file.Mode)
	  
	  log.Println("Status is ",entry.Status)
	  log.Println("HeadToIndex is ",entry.HeadToIndex)//git add 后进了索引区
	  log.Println("IndexToWorkdir is ",entry.IndexToWorkdir)// 在工作目录
	  
	  // log.Println("Status",entry.IndexToWorkdir.Status)
	  // log.Println("Flags",entry.IndexToWorkdir.Flags)
	  // log.Println("Similarity",entry.IndexToWorkdir.Similarity)
	  // log.Println("OldFile",entry.IndexToWorkdir.OldFile)
	  // log.Println("NewFile",entry.IndexToWorkdir.NewFile)
  }

  // type DiffFile struct {
  // 	Path  string
  // 	Oid   *Oid
  // 	Size  int
  // 	Flags DiffFlag
  // 	Mode  uint16
  // }
  
  // type DiffDelta struct {
  // 	Status     Delta
  // 	Flags      DiffFlag
  // 	Similarity uint16
  // 	OldFile    DiffFile
  // 	NewFile    DiffFile
  // }

  log.Println("StatusCurrent         = ",git.StatusCurrent )
  log.Println("StatusIndexNew        = ",git.StatusIndexNew )
  log.Println("StatusIndexModified   = ",git.StatusIndexModified )
  log.Println("StatusIndexDeleted    = ",git.StatusIndexDeleted )
  log.Println("StatusIndexRenamed    = ",git.StatusIndexRenamed )
  log.Println("StatusIndexTypeChange = ",git.StatusIndexTypeChange )
  log.Println("StatusWtNew           = ",git.StatusWtNew )
  log.Println("StatusWtModified      = ",git.StatusWtModified )
  log.Println("StatusWtDeleted       = ",git.StatusWtDeleted )
  log.Println("StatusWtTypeChange    = ",git.StatusWtTypeChange )
  log.Println("StatusWtRenamed       = ",git.StatusWtRenamed )
  log.Println("StatusIgnored         = ",git.StatusIgnored )
  log.Println("StatusConflicted      = ",git.StatusConflicted )

  // log.Println("DeltaUnmodified  Delta = ",git.DeltaUnmodified)
  // log.Println("DeltaAdded       Delta = ",git.DeltaAdded     )
  // log.Println("DeltaDeleted     Delta = ",git.DeltaDeleted   )
  // log.Println("DeltaModified    Delta = ",git.DeltaModified  )
  // log.Println("DeltaRenamed     Delta = ",git.DeltaRenamed   )
  // log.Println("DeltaCopied      Delta = ",git.DeltaCopied    )
  // log.Println("DeltaIgnored     Delta = ",git.DeltaIgnored   )
  // log.Println("DeltaUntracked   Delta = ",git.DeltaUntracked )
  // log.Println("DeltaTypeChange  Delta = ",git.DeltaTypeChange)
  // log.Println("DeltaUnreadable  Delta = ",git.DeltaUnreadable)
  // log.Println("DeltaConflicted  Delta = ",git.DeltaConflicted)
  // const (
  // 	DeltaUnmodified Delta = C.GIT_DELTA_UNMODIFIED
  // 	DeltaAdded      Delta = C.GIT_DELTA_ADDED
  // 	DeltaDeleted    Delta = C.GIT_DELTA_DELETED
  // 	DeltaModified   Delta = C.GIT_DELTA_MODIFIED
  // 	DeltaRenamed    Delta = C.GIT_DELTA_RENAMED
  // 	DeltaCopied     Delta = C.GIT_DELTA_COPIED
  // 	DeltaIgnored    Delta = C.GIT_DELTA_IGNORED
  // 	DeltaUntracked  Delta = C.GIT_DELTA_UNTRACKED
  // 	DeltaTypeChange Delta = C.GIT_DELTA_TYPECHANGE
  // 	DeltaUnreadable Delta = C.GIT_DELTA_UNREADABLE
  // 	DeltaConflicted Delta = C.GIT_DELTA_CONFLICTED
  // )

}  
