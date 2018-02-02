package main  

import (  
  git "github.com/libgit2/git2go"  
  log "github.com/sirupsen/logrus" 

)

func main(){
	
	
	user := "Devying"
	email := "huangby19890920@gmail.com"
	pass := "xxxxxxxxx"
	gitPath := "1.git"
	gitServer := "http://github.com/"
	target := "/tmp/huangby"
	//不能初始化裸仓库因为我们需要commit等操作
	repo, err := git.InitRepository(target,false)
	if err != nil {
		log.Warn(err)
		return
	}
	log.Println("git init finished");

	//git config

	config, err := repo.Config()
	if err != nil{
		log.Warn(err)
		return
	}
	err = config.SetString("user.name",user)

	if err != nil{
		log.Warn(err)
		return
	}
	err = config.SetString("user.email",email)
	if err != nil{
		log.Warn(err)
		return
	}

	err = config.SetString("pull.rebase","true")
	if err != nil{
		log.Warn(err)
		return
	}


	log.Println("git config finished");
	//git remote add

	remote, err := repo.Remotes.Create("origin", gitServer+user+"/"+gitPath)
	if err != nil {
		log.Warn(err)
		return
	}

	//git add
	//拿到INDEX
	idx, err := repo.Index()
	if err != nil {
		log.Warn(err)
		return
	}
	
	//将文件add写到INDEX中
	err = idx.AddAll([]string {"."}, git.IndexAddDefault, nil)
	if err != nil {
		log.Warn(err)
		return
	}
	//更新INDEX值
	err = idx.Write()
	if err != nil {
		log.Warn(err)
		return
	}
	log.Println("git add finished");

	//git commit
	//提交人信息
	signature := &git.Signature{
		Name:  user,
		Email: email,
		When:  time.Now(),
	}


	//创建tree组件
	treeId,err := idx.WriteTreeTo(repo)
	if err != nil {
		log.Warn(err)
		return
	}
	//
	tree,err := repo.LookupTree(treeId)
	if err != nil {
		log.Warn(err)
		return
	}
	message := "init"
	//创建commit组件返回一个hash码   这里初始化的时候最后一个参数可以不填写
	commitId,err := repo.CreateCommit("HEAD", signature, signature, message, tree)
	if err != nil {
		log.Warn(err)
		return
	}
	log.Println(commitId)
	log.Println("git commit finished");

	//git push
	
	called := false
	if err := remote.Push([]string{"refs/heads/master"}, &git.PushOptions{
		RemoteCallbacks: git.RemoteCallbacks{
			CredentialsCallback: func(url string, username_from_url string, allowed_types git.CredType) (git.ErrorCode, *git.Cred) {
				if called {
					return git.ErrUser, nil
				}
				called = true
				ret, creds := git.NewCredUserpassPlaintext(user, pass)
				return git.ErrorCode(ret), &creds
			},
		},
	}); err != nil {
		log.Warn("push err ",err)
		return
	}
	log.Println("git push finished");
	return
}