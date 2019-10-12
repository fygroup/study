git reflog   #你所有的操作
git cherry-pick  0cba81c   #会退到这个状态
//---初键库---------------------------------
(1)创建云端仓库
    git init --bare --share  //bare表示空库
(2)克隆云端仓库
    git clone ./**
    git add -A 
    git commit -m 'dadad'
    git push origin master
(3)连接github的操作
git init命令把这个目录变成Git可以管理的仓库
git config --global user.email ""
git config --global user.name ""
git remote add origin https://github.com/BingGostar/program_study
git add -A 
git commit -m ''
git push -u master origin
(2)
git clone https://github.com/BingGostar/program_study
git add -A 
git commit -m 'dsad'
git push -u master origin
//--------------------------------------
git config format.pretty oneline #显示历史记录时，每个提交的信息只显示一行
git config color.ui true
关于本地误删的后悔药:
git reflog #查看所有分支的所有操作记录
git cherry-pick 147b3b5  #恢复
如何让本地回归到先前的版本号:
git log 查看之前的版本信息获取版本号
git checkout 版本id号
例如:git  checkout b59aa374658e105d8c3be04d1d3ed931ae8a0c48
当pull冲突时
git fetch --all
git reset --hard origin/master
//---权限控制--------------------------------
你只有git pull的权利，并没有git push的权利
url = https://github.com/BingGostar/c_project.git   //会提示你用账号密码登陆

url = git@github.com:BingGostar/c_project.git       //会用密钥登陆
ssh-keygen -t rsa -C "you@example.com" //生成密钥，然后将公钥复制到github

//示例
账号
[core]
	repositoryformatversion = 0
	filemode = false
	bare = false
	logallrefupdates = true
	symlinks = false
	ignorecase = true
[remote "origin"]
	url = https://github.com/BingGostar/study
	fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
	remote = origin
	merge = refs/heads/master

密钥
[core]
	repositoryformatversion = 0
	filemode = false
	bare = false
	logallrefupdates = true
	symlinks = false
	ignorecase = true
[remote "origin"]
	url = git@github.com:BingGostar/program_study.git
	fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
	remote = origin
	merge = refs/heads/master
[user]
	name = malx
	email = guduqiubai.mlx@163.com 







