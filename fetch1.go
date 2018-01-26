package main  

import (  
  git "github.com/libgit2/git2go"  
  log "github.com/sirupsen/logrus"
) 


func main() {
	repo, err := git.OpenRepository("/tmp/code")
	if err != nil{
	  log.Println(err)
	  return
	}

	// branch, err := repo.Branch()
	// if err != nil {
	// 	log.Println(err)
	// }

	// // Get the name
	// name, err := branch.Name()
	// if err != nil {
	// 	log.Println(err)
	// }

	// Locate remote
	remote, err := repo.Remotes.Lookup("origin")
	if err != nil {
		log.Println(err)
		return
	}

	// Fetch changes from remote
	if err := remote.Fetch([]string{}, nil, ""); err != nil {
		log.Println(err)
		return
	}

	// Get remote master
	remoteBranch, err := repo.References.Lookup("refs/remotes/origin/master")
	if err != nil {
		log.Println(err)
		return
	}

	remoteBranchID := remoteBranch.Target()
	// Get annotated commit
	annotatedCommit, err := repo.AnnotatedCommitFromRef(remoteBranch)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(annotatedCommit)
	// Do the merge analysis
	mergeHeads := make([]*git.AnnotatedCommit, 1)
	mergeHeads[0] = annotatedCommit
	analysis, _, err := repo.MergeAnalysis(mergeHeads)
	if err != nil {
		log.Println(err)
		return
	}

	// Get repo head
	head, err := repo.Head()
	if err != nil {
		log.Println(err)
		return
	}

	if analysis & git.MergeAnalysisUpToDate != 0 {
		log.Println("Already up-to-date.")
		return
	}  else if analysis & git.MergeAnalysisNormal != 0 {
		// Just merge changes
		dftMgOpt,err :=  git.DefaultMergeOptions()
		if err != nil {
			log.Println(err)
			return
		}
		if err := repo.Merge([]*git.AnnotatedCommit{annotatedCommit}, &dftMgOpt , nil); err != nil {
			log.Println("xxxxxxx",err)
			// Check for conflicts
			index, err := repo.Index()
			if err != nil {
				log.Println(err)
				return
			}
			if index.HasConflicts() {
				log.Println("Conflicts encountered. Please resolve them.")
				return
			}
			return
		}
		// Check for conflicts
		index, err := repo.Index()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("0000")
		if index.HasConflicts() {
			log.Println("Conflicts encountered. Please resolve them.")
			return
		}
		log.Println("1111")
		// Make the merge commit
		sig, err := repo.DefaultSignature()
		if err != nil {
			log.Println(err)
			return
		}

		// Get Write Tree
		treeId, err := index.WriteTree()
		if err != nil {
			log.Println(err)
			return
		}

		tree, err := repo.LookupTree(treeId)
		if err != nil {
			log.Println(err)
			return
		}

		localCommit, err := repo.LookupCommit(head.Target())
		if err != nil {
			log.Println(err)
			return
		}

		remoteCommit, err := repo.LookupCommit(remoteBranchID)
		if err != nil {
			log.Println(err)
			return
		}

		repo.CreateCommit("HEAD", sig, sig, "", tree, localCommit, remoteCommit)

		// Clean up
		repo.StateCleanup()
	} else if analysis & git.MergeAnalysisFastForward != 0 {
		// Fast-forward changes
		// Get remote tree
		remoteTree, err := repo.LookupTree(remoteBranchID)
		if err != nil {
			log.Println(err)
			return
		}

		// Checkout
		if err := repo.CheckoutTree(remoteTree, nil); err != nil {
			log.Println(err)
			return
		}

		branchRef, err := repo.References.Lookup("refs/heads/master")
		if err != nil {
			log.Println(err)
			return
		}

		// Point branch to the object
		branchRef.SetTarget(remoteBranchID, "")
		if _, err := head.SetTarget(remoteBranchID, ""); err != nil {
			log.Println(err)
			return
		}

	} else {
		log.Println("Unexpected merge analysis result %d", analysis)
		return
	}
}