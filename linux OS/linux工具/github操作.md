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
(3)
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


### 操作
```
// github上新建一个新库，然后将本地库推上去
	在github上面创建一个新库
	git init
	git add README.md
	git commit -m "first commit"
	git remote add origin https://github.com/BingGostar/imageToLatex.git
	git push -u origin master
	之后每次提交 git push master origin

// github上存在一个库
	git clone https://github.com/BingGostar/study
	...

// 

```

### config
```
// 对于账号登陆
[core]
	repositoryformatversion = 0
	filemode = false
	bare = false
	logallrefupdates = true
	symlinks = false
	ignorecase = true
[remote "origin"]
	url = https://github.com/BingGostar/study		<<<重要
	fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
	remote = origin
	merge = refs/heads/master

//对于密钥登陆
[core]
	repositoryformatversion = 0
	filemode = false
	bare = false
	logallrefupdates = true
	symlinks = false
	ignorecase = true
[remote "origin"]
	url = git@github.com:BingGostar/program_study.git  <<< 重要
	fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
	remote = origin
	merge = refs/heads/master
[user]
	name = malx
	email = guduqiubai.mlx@163.com 
```

### 相关概念
```
(1) 工作区（Working Directory）
	就是你在电脑里能看到的目录

(1) 版本库（Repository）
	工作区有一个隐藏目录.git，这个不算工作区，而是Git的版本库
	里面包含:
	1) stage暂存区
		git add把文件添加进去，实际上就是把文件修改添加到暂存区
	2) master
		创建Git版本库时，Git自动为我们创建了唯一一个master分支，git commit就是往master分支上提交更改。
	3) HEAD

	4) fetch


```

### 操作命令
```
// git status
获得仓库当前的状态，可以告诉我们哪些文件被修改过，那些文件还没有提交

// git diff <file>
查看文件修改的内容

// git log
查看提交历史，以便确定要回退到哪个版本

// git reflog
查看命令历史，以便确定要回到未来的哪个版本

// git reset --hard <版本号>  （版本号没必要全写，前几位就可以）
回溯版本

// git reset HEAD <file>
把暂存区的修改撤销掉（unstage），重新放回工作区
场景：当你不但改乱了工作区某个文件的内容，还添加到了暂存区时，想丢弃修改


// git checkout -- <file>		（注意 --）
场景：当你改乱了工作区某个文件的内容，想直接丢弃工作区的修改时
情形一：文件没有被放到暂存区，撤销修改就回到和版本库一模一样的状态；
情形二：文件已经添加到暂存区后，又作了修改，撤销修改就回到添加到暂存区后的状态。

// git checkout

// git rm 
删除文件，流程如下
rm <file> && git rm <file> && git commit 

// git remote 


```




