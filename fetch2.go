package main
import (
	"github.com/runemadsen/easygit"
    "fmt"
)

func main() {

	  // Clones a repo to a local path. Credentials are not used if
	    // it's a public repo.
		//   err := easygit.Clone("https://github.com/runemadsen/testrepo.git", "/my/repo/path", "user", "password")
		//   fmt.Println(err.(*git.GitError).Code) // if -7, it was wrong credentials for private repo

		//   // Add all files to index. Similar to 'git add .'
		//   err := easygit.AddAll("path/to/repo")

		//   // Commit files in the index. If repo is commit, it will also create HEAD
		//   err := easygit.Commit("path/to/repo", "My commit message", "Git Name", "Git Email")

		//   // List all local branches. Similar to 'git branch'
		//   branchNames := easygit.ListBranches("path/to/repo")

		//   // Get the current local branch
		//   currentBranch := easyGit.CurrentBranch("path/to/repo")

		//   // Deletes a branch
		//   err := easygit.DeleteBranch("path/to/repo", "mybranch")

		//   // Creates a branch from another branch.
		//   err := easygit.CreateBranch("path/to/repo", "master", "newbranch")

		//   // Pushes a branch to a HTTPS remote. Similar to 'git push origin master'
		//   err := easygit.PushBranch("path/to/repo", "origin", "master", "httpsuser", "httpspassword")
		//   fmt.Println(err.(*git.GitError).Code) // if -7, it was wrong credentials

		  // Pulls and merges a branch from a remote. This function is a bit opinionated: If the pull
		    // results in a conflict, it will reset back to the current state and return an error saying "conflict".
			  // This will not leave the repo in a conflict state, but leave the branch in the state that it was before the pull.
			    // If there is no conflict, it will create a merge commit and merge the local and remote branch
				  // (which is why it needs the last signature params)
				    err := easygit.PullBranch("/tmp/code", "origin", "master", "Devying", "0299891a", "hby", "hby@imooc.com")
					  fmt.Println(err)
					    // Checks out a branch. Similar to 'git checkout slave'
						  //err := easygit.CheckoutBranch("path/to/repo", "mybranch")
					  }
