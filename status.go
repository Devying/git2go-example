package main  

import (  
	git "github.com/libgit2/git2go"  
	log "github.com/sirupsen/logrus"
	"encoding/json"

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
  flag := git.StatusOptIncludeUntracked | git.StatusOptRenamesHeadToIndex | git.StatusOptSortCaseSensitively
  StatusOptions:=&git.StatusOptions{Show:git.StatusShowIndexAndWorkdir,Flags:flag};
  StatusList, err := Repository.StatusList(StatusOptions)  
  if err != nil {  
		  log.Print(err)
		  return 
  }
  //修改的文件数量
  count, err:=StatusList.EntryCount()
  if err != nil {  
	  log.Print(err)
  }
  log.Print("the cout is ",count)

  var statusFiles []map[string]interface{}
  var statusFilesAll []map[string]interface{}
  for i := 0;i < count; i++ {
	  //状态条目entry
	  entry, err := StatusList.ByIndex(i)
	  if err != nil{
		  log.Println(err)
	  }
	  //log.Println("Status is ",entry.Status)
	  //entry的结构
	// type StatusEntry struct {
	// 	Status         Status //一个值
	// 	HeadToIndex    DiffDelta
	// 	IndexToWorkdir DiffDelta
	// }
	//DiffDelta的结构
	// type DiffDelta struct {
	// 	Status     Delta//一个值 0-10
	// 	Flags      DiffFlag//一个值 1,2,4,8
	// 	Similarity uint16
	// 	OldFile    DiffFile //旧
	// 	NewFile    DiffFile //新
	// }
	//DiffFile 的结构

	// type DiffFile struct {
	// 	Path  string //文件路径
	// 	Oid   *Oid //hash 值
	// 	Size  int  //大小
	// 	Flags DiffFlag //一个值 1,2,4,8
	// 	Mode  uint16 //文件权限
	// }
	statusFilesItem := make(map[string]interface{})
	statusFilesItemAll := make(map[string]interface{})

	headToIndex := entry.HeadToIndex;
	indexToWorkdir := entry.IndexToWorkdir;
	// Index 中的状态
	switch headToIndex.Status {
		case git.DeltaUnmodified:
			statusFilesItem["hi_status"] = git.DeltaUnmodified
			statusFilesItem["hi_status_msg"] = nil
		case git.DeltaAdded:
			statusFilesItem["hi_status"] = git.DeltaAdded
			statusFilesItem["hi_status_msg"] = "Added"
		case git.DeltaDeleted:
			statusFilesItem["hi_status"] = git.DeltaDeleted
			statusFilesItem["hi_status_msg"] = "Deleted"
		case git.DeltaModified:
			statusFilesItem["hi_status"] = git.DeltaModified
			statusFilesItem["hi_status_msg"] = "Modified"
		case git.DeltaRenamed:
			statusFilesItem["hi_status"] = git.DeltaRenamed
			statusFilesItem["hi_status_msg"] = "Renamed"
		case git.DeltaCopied:
			statusFilesItem["hi_status"] = git.DeltaCopied
			statusFilesItem["hi_status_msg"] = "Copied"
		case git.DeltaIgnored:
			statusFilesItem["hi_status"] = git.DeltaIgnored
			statusFilesItem["hi_status_msg"] = "Ignored"
		case git.DeltaUntracked:
			statusFilesItem["hi_status"] = git.DeltaUntracked
			statusFilesItem["hi_status_msg"] = "Untracked"
		case git.DeltaTypeChange:
			statusFilesItem["hi_status"] = git.DeltaTypeChange
			statusFilesItem["hi_status_msg"] = "TypeChange"
		case git.DeltaUnreadable:
			statusFilesItem["hi_status"] = git.DeltaUnreadable
			statusFilesItem["hi_status_msg"] = "Unreadable"
		case git.DeltaConflicted:
			statusFilesItem["hi_status"] = git.DeltaConflicted
			statusFilesItem["hi_status_msg"] = "Conflicted"
		default:
			statusFilesItem["hi_status"] = -1
			statusFilesItem["hi_status_msg"] = "Unknown"
	}
	// 工作区中的状态
	switch indexToWorkdir.Status {
		case git.DeltaUnmodified:
			statusFilesItem["iw_status"] = git.DeltaUnmodified
			statusFilesItem["iw_status_msg"] = nil
		case git.DeltaAdded:
			statusFilesItem["iw_status"] = git.DeltaAdded
			statusFilesItem["iw_status_msg"] = "Added"
		case git.DeltaDeleted:
			statusFilesItem["iw_status"] = git.DeltaDeleted
			statusFilesItem["iw_status_msg"] = "Deleted"
		case git.DeltaModified:
			statusFilesItem["iw_status"] = git.DeltaModified
			statusFilesItem["iw_status_msg"] = "Modified"
		case git.DeltaRenamed:
			statusFilesItem["iw_status"] = git.DeltaRenamed
			statusFilesItem["iw_status_msg"] = "Renamed"
		case git.DeltaCopied:
			statusFilesItem["iw_status"] = git.DeltaCopied
			statusFilesItem["iw_status_msg"] = "Copied"
		case git.DeltaIgnored:
			statusFilesItem["iw_status"] = git.DeltaIgnored
			statusFilesItem["iw_status_msg"] = "Ignored"
		case git.DeltaUntracked:
			statusFilesItem["iw_status"] = git.DeltaUntracked
			statusFilesItem["iw_status_msg"] = "Untracked"
		case git.DeltaTypeChange:
			statusFilesItem["iw_status"] = git.DeltaTypeChange
			statusFilesItem["iw_status_msg"] = "TypeChange"
		case git.DeltaUnreadable:
			statusFilesItem["iw_status"] = git.DeltaUnreadable
			statusFilesItem["iw_status_msg"] = "Unreadable"
		case git.DeltaConflicted:
			statusFilesItem["iw_status"] = git.DeltaConflicted
			statusFilesItem["iw_status_msg"] = "Conflicted"
		default:
			statusFilesItem["iw_status"] = -1
			statusFilesItem["iw_status_msg"] = "Unknown"
	}
	statusFilesItem["hi_file_old"]=headToIndex.OldFile.Path
	statusFilesItem["hi_file_new"]=headToIndex.NewFile.Path
	statusFilesItem["iw_file_old"]=indexToWorkdir.OldFile.Path
	statusFilesItem["iw_file_new"]=indexToWorkdir.NewFile.Path
	statusFiles = append(statusFiles,statusFilesItem)

	if statusFilesItem["hi_status"] != git.DeltaUnmodified {
		statusFilesItemAll["status"] = statusFilesItem["hi_status"]
		statusFilesItemAll["status_msg"] = statusFilesItem["hi_status_msg"]
		statusFilesItemAll["file"] = statusFilesItem["hi_file_new"]
	}
	if statusFilesItem["iw_status"] != git.DeltaUnmodified {
		statusFilesItemAll["status"] = statusFilesItem["iw_status"]
		statusFilesItemAll["status_msg"] = statusFilesItem["iw_status_msg"]
		statusFilesItemAll["file"] = statusFilesItem["iw_file_new"]
	}

	if statusFilesItem["hi_status"] != git.DeltaUnmodified && statusFilesItem["iw_status"] != git.DeltaUnmodified {
		statusFilesItemAll["status"] = statusFilesItem["iw_status"]
		statusFilesItemAll["status_msg"] = statusFilesItem["iw_status_msg"]
		statusFilesItemAll["file"] = statusFilesItem["iw_file_new"]
	}
	statusFilesAll = append(statusFilesAll,statusFilesItemAll)

	//   var S string
	//   var file git.DiffFile = entry.IndexToWorkdir.OldFile
	//   if entry.Status == git.StatusIndexNew {
	// 	  //git add 以后就要去索引区了
	// 	  file = entry.IndexToWorkdir.NewFile
	//   }	
	//   if entry.Status == git.StatusWtNew {
	// 	  S += "A" 
	//   } else if entry.Status == git.StatusWtModified {
	// 	  S += "M"
	//   } else if entry.Status == git.StatusWtDeleted {
	// 	  S += "D"
	//   } else {
	// 	  S += "?"
	//   }
	  
	//   log.Println(S,file.Path,file.Flags)
	//   log.Println("Status is ",entry.Status)
	  
	//   log.Println("HeadToIndex is ",entry.HeadToIndex)//git add 后进了索引区
	//   log.Println("IndexToWorkdir is ",entry.IndexToWorkdir)// 在工作目录
	  
	  // log.Println("Status",entry.IndexToWorkdir.Status)
	  // log.Println("Flags",entry.IndexToWorkdir.Flags)
	  // log.Println("Similarity",entry.IndexToWorkdir.Similarity)
	  // log.Println("OldFile",entry.IndexToWorkdir.OldFile)
	  // log.Println("NewFile",entry.IndexToWorkdir.NewFile)
  }

  b, _ := json.Marshal(statusFiles)
  c, _ := json.Marshal(statusFilesAll)
  log.Println(string(b))
  log.Println(string(c))
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
  //什么状态没有
  log.Println("StatusCurrent         = ",git.StatusCurrent )
  
  //索引中的状态
  log.Println("StatusIndexNew        = ",git.StatusIndexNew )//新建文件执行git add 后的状态
  log.Println("StatusIndexModified   = ",git.StatusIndexModified )//修改旧文件执行git add 后的状态
  log.Println("StatusIndexDeleted    = ",git.StatusIndexDeleted )//删除旧文件执行git add 后的状态
  log.Println("StatusIndexRenamed    = ",git.StatusIndexRenamed )
  log.Println("StatusIndexTypeChange = ",git.StatusIndexTypeChange )
  //工作目录中的状态
  log.Println("StatusWtNew           = ",git.StatusWtNew )//新建文件后啥都不执行的状态
  log.Println("StatusWtModified      = ",git.StatusWtModified )//修改旧文件后啥都不执行的状态
  log.Println("StatusWtDeleted       = ",git.StatusWtDeleted )//删除旧文件后啥都不执行的状态
  log.Println("StatusWtTypeChange    = ",git.StatusWtTypeChange )
  log.Println("StatusWtRenamed       = ",git.StatusWtRenamed )
  //忽略的文件
  log.Println("StatusIgnored         = ",git.StatusIgnored )
  //冲突的文件
  log.Println("StatusConflicted      = ",git.StatusConflicted )

  log.Println("===================head--->index&&index--->workdir======================")
  log.Println("DeltaUnmodified  Delta = ",git.DeltaUnmodified)
  log.Println("DeltaAdded       Delta = ",git.DeltaAdded     )
  log.Println("DeltaDeleted     Delta = ",git.DeltaDeleted   )
  log.Println("DeltaModified    Delta = ",git.DeltaModified  )
  log.Println("DeltaRenamed     Delta = ",git.DeltaRenamed   )
  log.Println("DeltaCopied      Delta = ",git.DeltaCopied    )
  log.Println("DeltaIgnored     Delta = ",git.DeltaIgnored   )
  log.Println("DeltaUntracked   Delta = ",git.DeltaUntracked )
  log.Println("DeltaTypeChange  Delta = ",git.DeltaTypeChange)
  log.Println("DeltaUnreadable  Delta = ",git.DeltaUnreadable)
  log.Println("DeltaConflicted  Delta = ",git.DeltaConflicted)



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
  log.Println("==================diff flag ======================")
  log.Println("DiffFlagBinary  Delta = ",git.DiffFlagBinary)
  log.Println("DiffFlagNotBinary  Delta = ",git.DiffFlagNotBinary)
  log.Println("DiffFlagValidOid  Delta = ",git.DiffFlagValidOid)
  log.Println("DiffFlagExists  Delta = ",git.DiffFlagExists)
//   const (
// 	DiffFlagBinary    DiffFlag = C.GIT_DIFF_FLAG_BINARY
// 	DiffFlagNotBinary DiffFlag = C.GIT_DIFF_FLAG_NOT_BINARY
// 	DiffFlagValidOid  DiffFlag = C.GIT_DIFF_FLAG_VALID_ID
// 	DiffFlagExists    DiffFlag = C.GIT_DIFF_FLAG_EXISTS
//   )
  StatusList.Free()
}  
