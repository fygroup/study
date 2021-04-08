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

### 认证方式
```
(1) 基于HTTPS
	选择这种，我们每次向该库push代码的时候，都要输入用户名和密码(当然，我自然不愿意将密码告诉别人啦)。
	url = https://github.com/BingGostar/c_project.git   //会提示你用账号密码登陆

(2)基于SSH
	选择这种，我们就可以通过公钥密钥的身份来验证自己的权限，下面重点介绍的就是这个。
	url = git@github.com:BingGostar/c_project.git       //会用密钥登陆
	ssh-keygen -t rsa -C "you@example.com" //生成密钥，然后将公钥复制到github
```

### 权限控制
```
git 权限控制就是linux的权限控制
```

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

// git clone 用户 密码
	git clone https://diaohaiyong:939791536a-@gitlab.sz.sensetime.com/senseNebula-m/nebula-mini.git
	
// pull request
	https://www.zhihu.com/question/21682976
	1) 先 fork 别人的仓库到自己的仓库
		在浏览器上fork https://github.com/twbs/bootstrap.git
	2) clone到本地
		git clone https://github.com/BingGostar/bootstrap.git
	3) 创建切换分支
		git checkout -b test-pr
	4) 添加修改
		git add . && git commit -m 'test-pr'
	5) 推送分支到远端							
		git push origin test-pr
	6) 等待作者merge
```

### config
```
git config --global user.name "John Doe"
git config --global user.email johndoe@example.com

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

(2) 版本库（Repository）
	工作区有一个隐藏目录.git，这个不算工作区，而是Git的版本库
	里面包含:
	1) stage暂存区
		git add把文件添加进去，实际上就是把文件修改添加到暂存区
	2) master
		创建Git版本库时，Git自动为我们创建了唯一一个master分支，git commit就是往master分支上提交更改。
	3) HEAD

	4) 分支

(3) 分支策略
	master分支应该是非常稳定的
	干活都在dev分支上
	每个人都有自己的分支，时不时地往dev分支上合并就可以了


```

### 操作命令
```
// git status
获得仓库当前的状态，可以告诉我们哪些文件被修改过，那些文件还没有提交

// git diff <file>
查看文件修改的内容

// git log
查看提交历史，以便确定要回退到哪个版本

// git log --graph
分支合并图。

// git reflog
查看命令历史，以便确定要回到未来的哪个版本

// git reset --hard <版本号>  （版本号没必要全写，前几位就可以）
回溯版本

// git reset HEAD <file>
把暂存区的修改撤销掉（unstage），重新放回工作区
场景：当你不但改乱了工作区某个文件的内容，还添加到了暂存区时，想丢弃修改

// git rm 
删除文件，流程如下
rm <file> && git rm <file> && git commit 

// git checkout -- <file>		（注意 --）
撤销文件修改
场景：当你改乱了工作区某个文件的内容，想直接丢弃工作区的修改时
情形一：文件没有被放到暂存区，撤销修改就回到和版本库一模一样的状态；
情形二：文件已经添加到暂存区后，又作了修改，撤销修改就回到添加到暂存区后的状态。

// git checkout -b <分支>
我们创建分支，然后切换到分支
相当于git branch <分支> && git checkout <分支>

// git checkout -b <分支> origin/<分支>   （origin/<分支> 已存在？？？）
在本地创建和远程分支对应的分支

// git branch
查看分支

// git branch <分支>
新建分支

// git branch -d <分支>
删除分支

// git branch -D <分支>
强行删除分支

// git merge <分支>
用于合并"指定分支"到"当前分支"
当Git无法自动合并分支时，就必须首先解决冲突。解决冲突后，再提交，合并完成。
解决冲突就是把Git合并失败的文件手动编辑为我们希望的内容，再提交。

// git switch <分支>
切换分支

// git switch -c <分支>
新建并切换分支

// git remote 
查看远程库的信息

// git remote -v
查看远程库的更详细信息

// git push origin master
推送主分支

// git push origin <分支>
推送其他分支

// git pull


```

### .gitignore
```
https://zhuanlan.zhihu.com/p/52885189

git为我们提供了一个.gitignore文件，只要在这个文件中声明哪些文件你不希望添加到git中去，这样当你使用git add .的时候这些文件就会被自动忽略掉。
```

### 远端分支
```
//查看远端分支
git branch -r
//创建本地分支
git branch current
//切换分支
git checkout current
// 推送分支
git push origin current	
// 更新分支
git pull origin current	
// 删除分支
git push origin :current

```

### stash
```
// 情形
我们往往会建一个自己的分支去修改和调试代码, 如果别人或者自己发现原有的分支上有个不得不修改的bug，我们往往会把完成一半的代码commit提交到本地仓库，然后切换分支去修改bug，改好之后再切换回来。这样的话往往log上会有大量不必要的记录

// git stash
会把所有未提交的修改（包括暂存的和非暂存的）都保存起来，用于后续恢复当前工作目录

// git stash pop
将缓存堆栈中的第一个stash删除，并将对应修改应用到当前的工作目录下

// git stash apply
不会删除堆栈中的第一个stash，并将对应修改应用到当前的工作目录下

// git stash show
查看stash

// git stash drop
移除stash
```

### git pull push 带用户 密码
```
git push https://diaohaiyong:939791536a-@gitlab.sz.sensetime.com/senseNebula-m/nebula-m.git develop-v2.1.2_jcv
git pull https://diaohaiyong:939791536a-@gitlab.sz.sensetime.com/senseNebula-m/nebula-m.git develop-v2.1.2_jcv

```