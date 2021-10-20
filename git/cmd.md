[TOC]
# git
https://git-scm.com/docs/
## 相关名词解释
master: 默认开发分支
origin: 默认远程版本库
Index / Stage：暂存区
Workspace：工作区
Repository：仓库区（或本地仓库）
Remote：远程仓库

## 常用命令

git clone

### Git 全局设置

```
git config --global user.name "kingreatwill"
git config --global user.email "350840291@qq.com"
```
配置存放在.gitconfig文件中
```
# 显示当前的Git配置
$ git config --list
# 编辑Git配置文件
$ git config -e [--global]
```

### 远程仓库
git remote -v
git remote add shortname url # 添加远程仓库
git remote add origin git@github.com:xxx/learngit.git
git remote rm origin：删除已有的远程仓库；
git fetch shortname #从远程仓库中获得数据，
git remote show origin #查看某个远程仓库
git remote rename origin new_origin #重命名
git remote rm paul #删除
```
# 下载远程仓库的所有变动
$ git fetch [remote]
# 取回远程仓库的变化，并与本地分支合并
$ git pull [remote] [branch]
# 显示所有远程仓库
$ git remote -v
# 显示某个远程仓库的信息
$ git remote show [remote]
# 增加一个新的远程仓库，并命名
$ git remote add [shortname] [url]
# 上传本地指定分支到远程仓库
$ git push [remote] [branch]
# 强行推送当前分支到远程仓库，即使有冲突
$ git push [remote] --force
# 推送所有分支到远程仓库
$ git push [remote] --all
# 删除远程分支或标签 
$ git push <remote> :<branch/tag-name> 
# 上传所有标签
$ git push --tags 
```

### 增加/删除/修改文件
```
# 查看状态
$ git status
# 查看变更内容
$ git diff
# 添加指定文件到暂存区
$ git add [file1] [file2]
# 添加指定目录到暂存区，包括子目录
$ git add [dir]
# 添加当前目录的所有文件到暂存区
$ git add
# 添加每个变化前，都会要求确认，对于同一个文件的多处变化，可以实现分次提交
$ git add -p
# 删除工作区文件，并且将这次删除放入暂存区
$ git rm [file1] [file2]
# 停止追踪指定文件，但该文件会保留在工作区
$ git rm --cached [file]
# 改名文件，并且将这个改名放入暂存区
$ git mv [file-original] [file-renamed
```
### git commit
```
# 提交暂存区到仓库区
$ git commit -m [message]
# 提交暂存区的指定文件到仓库区
$ git commit [file1] [file2] ... -m [message]
# 提交工作区自上次commit之后的变化，直接到仓库区
$ git commit -a
# 提交时显示所有diff信息
$ git commit -v
# 使用一次新的commit，替代上一次提交
# 如果代码没有任何新变化，则用来改写上一次commit的提交信息
$ git commit --amend -m [message]
# 重做上一次commit，并包括指定文件的新变化
$ git commit --amend [file1] [file2]
```

### git pull在指定位置的本地仓库
`git --work-tree=E:/xxx --git-dir=E:/xxx/.git pull`
`git -C "E:/xxx" pull`


### 新建仓库

创建一个新仓库
```
# 在当前目录新建一个Git代码库
$ git init
# 新建一个目录，将其初始化为Git代码库
$ git init [project-name]
# 下载一个项目和它的整个代码历史
$ git clone [url]
```

```
git clone git@codechina.csdn.net:kingreatwill/open.git
cd open
touch README.md
git add README.md
git commit -m "add README"
git push -u origin master
```

#### 推送现有文件夹
```
cd existing_folder
git init
git remote add origin git@codechina.csdn.net:kingreatwill/open.git
git add .
git commit -m "Initial commit"
git push -u origin master
```
`git add xx`命令可以将xx文件添加到暂存区，如果有很多改动可以通过 
`git add -A .`来一次添加所有改变的文件。
注意 `-A` 选项后面还有一个句点。
`git add -A`表示添加所有内容， 
`git add .` 表示添加新文件和编辑过的文件不包括删除的文件; 
`git add -u` 表示添加编辑或者删除的文件，不包括新添加的文件
`git clone https://git.coding.net/gamedaybyday/HelloGit.git D:\Git\HelloGit`
`git pull https://zhangsan:123456@github.com`

git clone 带用户名密码的形式但包含@等特殊符号无法正常解析
```
正常使用git clone 的方式

git clone https：//remote

使用带用户名密码的方式（可以避免后续每次都要输入用户名密码）

git clone https://[username]:[password]@/remote

但有时会出现用户名或密码中含有像@这样的特殊符号，而不能被正常解析

我们需要通过下面方式进行重新编码

String c = URLEncoder.encode("@","utf-8");
System.out.println(c);

console -> %40
所有这样就可以知道@在url中需要写成%40的形式
```

推送现有的 Git 仓库
```
cd existing_repo
git remote rename origin old-origin
git remote add origin git@codechina.csdn.net:kingreatwill/open.git
git push -u origin --all
git push -u origin --tags
```

### git 查看最近或某一次提交修改的文件列表相关命令整理
```
git log --name-status #每次修改的文件列表, 显示状态
git log --name-only #每次修改的文件列表
git log --name-only -1 #查看最近一次提交的文件列表
git log --stat #每次修改的文件列表, 及文件修改的统计
git log --pretty=format:"" -1 # --pretty=format:"" 格式化commit message 这里什么都不显示
git log --pretty=format:"" --name-only -1 # --pretty=format:"" 格式化commit message 这里什么都不显示

git whatchanged #每次修改的文件列表
git whatchanged --stat #每次修改的文件列表, 及文件修改的统计

git show #显示最后一次的文件改变的具体内容
git show -5 #显示最后 5 次的文件改变的具体内容
git show <commit-id> #显示某个 commitid 改变的具体内容
git show --name-only #显示最后一次的文件列表
git show --pretty=format:"" --name-only #显示最后一次的文件列表 # --pretty=format:"" 格式化commit message 这里什么都不显示
git show --pretty=format:"" --name-only <commit-id> #显示某个 commitid 的文件列表

git diff --name-only <commit-id-1> <commit-id-2> #比较两次commit修改的文件列表


```

### git push
git push [remote-name] [branchname]。 
git push -u origin master：由于远程库是空的，我们第一次推送master分支时，加上了-u参数，Git不但会把本地的master分支内容推送的远程新的master分支，还会把本地的master分支和远程的master分支关联起来，在以后的推送或者拉取时就可以简化命令。
git push origin master
git push origin tagname：推送某个标签到远程
git push origin --tags：一次性推送全部尚未推送到远程的本地标签

git init
git add file/directory
git add -A 添加所有

git commit -m <注释>
git commit 加上 -a 跳过 git add 步骤
git commit -a -m 'added new benchmarks'
git commit --amend #撤消提交
例如，你提交后发现忘记了暂存某些需要的修改，可以像下面这样操作：
```
$ git commit -m 'initial commit'
$ git add forgotten_file
$ git commit --amend
```
最终你只会有一个提交——第二次提交将代替第一次提交的结果。


git status # git status -s ：--short 
git diff

### git 分支


git branch google  #创建分支
git checkout google # 切换分支
git checkout -b google # 创建并切换
git push origin newbranch # 推送本地分支到远程分支（远程没有会自动创建）
git push origin newbranch:newbranch # 推送本地分支到远程分支（远程没有会自动创建，：前面的是本地分支，：后面是远程分支）
git branch -d xx # 删除本地分支
git pull origin dev # 拉取远程分支到本地

有时在远程创建的分支本地没有：
git branch -a # 查看所有分支
git fetch --all # 拉取所有
git checkout -b 0.1 origin/0.1  # 拉取远程分支0.1 并创建本地分支0.1
git branch --set-upstream-to origin/0.1 # 关联本地当前分支到远程分支

```
# 显示所有本地分支
$ git branch
# 列出所有远程分支
$ git branch -r
# 列出所有本地分支和远程分支
$ git branch -a
# 新建一个分支，但依然停留在当前分支
$ git branch [branch-name]
# 新建一个分支，与指定的远程分支建立追踪关系
$ git branch --track [branch] [remote-branch]
# 删除分支
$ git branch -d [branch-name]
# 删除远程分支
$ git push origin --delete [branch-namel
$ git branch -dr [remote/branch]
# 新建一个分支，并切换到该分支
$ git checkout -b [branch]
# 切换到指定分支，并更新工作区
$ git checkout [branch-name]
# 切换到上一个分支
$ git checkout -
# 建立追踪关系，在现有分支与指定的远程分支之间
$ git branch --set-upstream [branch] [remote-branch]
# 合并指定分支到当前分支
$ git merge [branch]
# 衍合指定分支到当前分支
$ git rebase <branch>
# 选择一个commit，合并进当前分支
$ git cherry-pick [commit]
```

git rm
git checkout -b dev：git checkout命令加上-b参数表示创建并切换，相当于以下两条命令：
git branch dev
git checkout dev
git push origin dev # 推送本地分支到远程分支（远程没有会自动创建）
git branch：查看当前分支，git branch命令会列出所有分支，当前分支前面会标一个*号。
git branch -d dev：删除dev分支了
git merge
git merge dev：git merge命令用于合并指定分支到当前分支。
### git tag
git tag：查看所有标签
git tag -a v1.4 -m "my version 1.4" # 附注标签（annotated）
git tag v1.4-lw #轻量标签（lightweight）
git tag v1.0：在当前分支的HEAD上打标签
git tag v0.9 f52c633：#后期打标签，在当前分支的指定commit上打标签
git tag -a v0.9 f52c633
git tag -d v0.1：删除指定标签
git tag -l 'v1.8.5*' #只对 1.8.5系列感兴趣
git show v1.4
git push origin v1.4
git push origin --tags：一次性推送全部尚未推送到远程的本地标签
git checkout v1.4

```
# 列出所有本地标签 
$ git tag 
# 基于最新提交创建标签 
$ git tag <tagname> 
# 删除标签
$ git tag -d <tagname> 
# 删除远程tag
$ git push origin :refs/tags/[tagName]
# 查看tag信息
$ git show [tag]
# 提交指定tag
$ git push [remote] [tag]
# 提交所有tag
$ git push [remote] --tags
# 新建一个分支，指向某个tag
$ git checkout -b [branch] [tag]
```

### 查看信息
```
# 显示有变更的文件
$ git status
# 显示当前分支的版本历史
$ git log
# 显示commit历史，以及每次commit发生变更的文件
$ git log --stat
# 搜索提交历史，根据关键词
$ git log -s [keyword]
# 显示某个commit之后的所有变动，每个commit占据一行
$ git log [tag] HEAD --pretty=format:%s
# 显示某个commit之后的所有变动，其"提交说明"必须符合搜索条件
$ git log [tag] HEAD --grep feature
# 显示某个文件的版本历史，包括文件改名
$ git log --follow [file]
$ git whatchanged [file]
# 显示指定文件相关的每一次diff
$ git log -p [file]
# 显示过去5次提交
$ git log -5 --pretty--oneline
# 显示所有提交过的用户，按提交次数排序
$ git shortlog -sn
# 显示指定文件是什么人在什么时间修改过
$ git blame [file]
# 显示暂存区和工作区的差异
$ git diff
# 显示暂存区和上一个commit的差异
$ git diff --cached [file]
# 显示工作区与当前分支最新commit之间的差异
$ git diff HEAD
# 显示两次提交之间的差异
$ git diff [first-branch]...[second-branch]
# 显示今天你写了多少行代码
$ git diff --shortstat "@{0 day ago}"
# 显示某次提交的元数据和内容变化
$ git show [commit]
# 显示某次提交发生变化的文件
$ git show --name-only [commit]
# 显示某次提交时，某个文件的内容
$ git show [commit]:[filename]
# 显示当前分支的最近几次提交
$ git reflog
```

#### git log
git log：提交历史
git log --pretty=oneline
 git log --pretty=format:"%h - %an, %ar : %s"
另外一个常用的选项是 --pretty。 这个选项可以指定使用不同于默认格式的方式展示提交历史。 这个选项有一些内建的子选项供你使用。 比如用 oneline 将每个提交放在一行显示，查看的提交数很大时非常有用。 另外还有 short，full 和 fuller 可以用，展示的信息或多或少有些不同，请自己动手实践一下看看效果如何。但最有意思的是 format，可以定制要显示的记录格式。 这样的输出对后期提取分析格外有用 — 因为你知道输出的格式不会随着 Git 的更新而发生改变：

git log --pretty=format 常用的选项
选项 | 说明
--|--
%H |提交对象（commit）的完整哈希字串
%h |提交对象的简短哈希字串
%T |树对象（tree）的完整哈希字串
%t| 树对象的简短哈希字串
%P |父对象（parent）的完整哈希字串
%p |父对象的简短哈希字串
%an |作者（author）的名字
%ae |作者的电子邮件地址
%ad |作者修订日期（可以用 --date= 选项定制格式）
%ar |作者修订日期，按多久以前的方式显示
%cn |提交者（committer）的名字
%ce |提交者的电子邮件地址
%cd |提交日期
%cr |提交日期，按多久以前的方式显示
%s |提交说明
你一定奇怪 作者 和 提交者 之间究竟有何差别， 其实作者指的是实际作出修改的人，提交者指的是最后将此工作成果提交到仓库的人。 所以，当你为某个项目发布补丁，然后某个核心成员将你的补丁并入项目时，你就是作者，而那个核心成员就是提交者。

git log --graph：可以看到分支合并图。
 git log --pretty=format:"%h %s" --graph

git log -p -2
一个常用的选项是 -p，用来显示每次提交的内容差异。 你也可以加上 -2 来仅显示最近两次提交

git log --stat
如果你想看到每次提交的简略的统计信息，你可以使用 --stat 选项


git log 的常用选项
选项 |说明
--|--
-p |按补丁格式显示每个更新之间的差异。
--stat |显示每次更新的文件修改统计信息。
--shortstat |只显示 --stat 中最后的行数修改添加移除统计。
--name-only |仅在提交信息后显示已修改的文件清单。
--name-status |显示新增、修改、删除的文件清单。
--abbrev-commit |仅显示 SHA-1 的前几个字符，而非所有的 40 个字符。
--relative-date |使用较短的相对时间显示（比如，“2 weeks ago”）。
--graph |显示 ASCII 图形表示的分支合并历史。
--pretty |使用其他格式显示历史提交信息。可用的选项包括 oneline，short，full，fuller 和
format（后跟指定格式）。

查看文件历史
```shell
# see the changes of a file, works even
# if the file was deleted
git log --full-history -- x\xx\.gitkeep

# limit the output of Git log to the
# last commit, i.e. the commit which delete the file
# -1 to see only the last commit
# use 2 to see the last 2 commits etc
git log --full-history -1 -- x\xx\.gitkeep


# include stat parameter to see
# some statics, e.g., how many files were
# deleted
git log --full-history -1 --stat -- x\xx\.gitkeep
```
### 撤销
```
# 撤销工作目录中所有未提交文件的修改内容 
$ git reset --hard HEAD 
# 撤销指定的未提交文件的修改内容 
$ git checkout HEAD <file> 
# 撤销指定的提交 
$ git revert <commit> 
# 退回到之前1天的版本
$ git log --before="1 days"
# 恢复暂存区的指定文件到工作区
$ git checkout [file]
# 恢复某个commit的指定文件到暂存区和工作区
$ git checkout [commit] [file]
# 恢复暂存区的所有文件到工作区
$ git checkout.
# 重置暂存区的指定文件，与上一次commit保持一致，但工作区不变
$ git reset [file]
# 重置暂存区与工作区，与上一次commit保持一致
$ git reset --hard
# 重置当前分支的指针为指定commit，同时重置暂存区，但工作区不变
$ git reset [commit]
# 重置当前分支的HEAD为指定commit，同时重置暂存区和工作区，与指定commit一致
$ git reset --hard [commit]
# 重置当前HEAD为指定commit，但保持暂存区和工作区不变
$ git reset --keep [commit]
# 新建一个commit，用来撤销指定commit
# 后者的所有变化都将被前者抵消，并且应用到当前分支
$ git revert [commit]
# 暂时将未提交的变化移除，稍后再移入 
$ git stash
$ git stash pop
```
#### git reset
在使用Git的过程中，由于操作不当，作为初学者的我们可能经常要去解决冲突。某些时候，当你不小心改错了内容，或者错误地提交了某些commit，我们就需要进行版本的回退。版本回退最常用的命令包括git reset和git revert。这两个命令允许我们在版本的历史之间穿梭。

下面就几种比较经典的场景进行总结：
- 场景1：当你改乱了工作区某个文件的内容，想直接丢弃工作区的修改时，用命git checkout -- filename；
- 场景2：当你不但改乱了工作区某个文件的内容，还添加到了暂存区时，想丢弃修改，分两步，第一步用命令git reset HEAD file，就回到了场景1，第二步按场景1操作；
- 场景3：已经提交了不合适的修改到版本库时，想要撤销本次提交，使用git reset --hard commit_id，不过前提是没有推送到远程库。

穿梭前，用git log可以查看提交历史，以便确定要回退到哪个版本；要重返未来，用git reflog查看命令历史，以便确定要回到未来的哪个版本。

git reset HEAD~3 #当前分支相当于回滚了3个提交点，HEAD^ 一个 HEAD^^ 两个
git reset --hard HEAD^
git reset --hard 1094a
git reset HEAD <file>：git reset命令既可以回退版本，也可以把暂存区的修改回退到工作区。当我们用HEAD时，表示最新的版本。
reset有3种常用的模式：
- soft，只改变提交点，暂存区和工作目录的内容都不改变
- mixed，改变提交点，同时改变暂存区的内容。这是默认的回滚方式
- hard，暂存区、工作目录的内容都会被修改到与提交点完全一致的状态

git reflog：命令历史
git checkout
git checkout -- file：撤消对文件的修改:丢弃工作区的修改，让这个文件回到最近一次git commit或git add时的状态。命令中的--很重要，没有--，就变成了“切换到另一个分支”的命令。

#### Git revert
Git revert用来撤销某次操作，此次操作之前和之后的commit和history都会保留，并且把这次撤销作为一次最新的提交。git revert是提交一个新的版本，将需要revert的版本的内容再反向修改回去，版本会递增，不影响之前提交的内容。
Git revert和git reset都可以进行版本的回退，将工作区回退到历史的某个状态，二者有如下的区别：
- git revert是用一次新的commit来回滚之前的commit，而git reset是直接删除指定的commit（并没有真正的删除，通过git reflog可以找回），这是二者最显著的区别；
- git reset 是把HEAD向后移动了一下，而git revert是HEAD继续前进，只是新的commit的内容和要revert的内容正好相反，能够抵消要被revert的内容；
- 在回滚这一操作上，效果差不多。但是在日后继续merge以前的老版本时有区别。因为git revert是用一次逆向的commit"中和"之前的提交，因此日后合并老的branch时，导致这部分改变不会再次出现；但是git reset是之间把某些commit在某个branch上删除，因而和老的branch再次merge时，这些被回滚的commit应该还会被引入。

https://blog.csdn.net/yxlshk/article/details/79944535
https://git-scm.com/docs/git-revert

### Git stash
Git stash用来暂存当前正在进行的工作， 将工作区还没加入索引库的内容压入本地的Git栈中，在需要应用的时候再弹出来。比如想pull 最新代码，又不想加新commit；或者为了修复一个紧急的bug，先stash，使返回到自己上一个commit，改完bug之后再stash pop，继续原来的工作。Git stash可以让本地仓库返回到上一个提交状态，而本地的还未提交的内容则被压入Git栈。Git stash的基本使用流程如下：
```
git stash #暂存工作区尚未提交的内容
Do your work #在上一个提交的状态之上完成你的操作
git stash pop #将暂存的内容弹出并应用
```
当你多次使用git stash命令后，你的栈里将充满了未提交的代码，这时候你会对将哪个版本应用回来有些困惑，
这时git stash list命令可以将当前的Git栈信息打印出来，你只需要将找到对应的版本号，
例如使用 git stash apply stash@{1} 就可以将你指定版本号为stash@{1}的暂存内容取出来，当你将所有的栈都应用回来的时候，可以使用git stash clear来将栈清空。TortoiseGit中的stash save菜单就对应该命令。
https://git-scm.com/docs/git-stash

### 合并代码和提交
#### git merge
https://git-scm.com/docs/git-merge
```
# 合并指定分支到当前分支
$ git merge [branch]
```
#### rebase
- 合并多次提交纪录，减少了无用的提交信息
- 分支合并
把一个分支中的修改整合到另一个分支的办法有两种，
第一种是我们常用的git merge操作，
而第二种便是本节要讲的rebase（中文翻译为衍合）。
该命令的原理是，回到两个分支最近的共同祖先，根据当前分支（也就是要进行衍合的分支experiment）后续的历次提交对象（这里只有一个 C3），生成一系列文件补丁，然后以基底分支（也就是主干分支master）最后一个提交对象（C4）为新的出发点，逐个应用之前准备好的补丁文件，最后会生成一个新的合并提交对象（C3'），从而改写 experiment 的提交历史，使它成为 master 分支的直接下游。
一般我们使用rebase的目的，是想要得到一个能在远程分支上干净应用的补丁，比如某些项目你不是维护者，但想帮点忙的话，最好用rebase：先在自己的一个分支里进行开发，当准备向主项目提交补丁的时候，根据最新的 origin/master 进行一次衍合操作然后再提交，这样维护者就不需要做任何整合工作（实际上是把解决分支补丁同最新主干代码之间冲突的责任，化转为由提交补丁的人来解决），只需根据你提供的仓库地址作一次快进合并，或者直接采纳你提交的补丁。

在rebase的过程中，也许会出现冲突。在这种情况，Git会停止rebase并会让你去解决冲突；在解决完冲突后，用git add命令去更新这些内容的索引， 然后，你无需执行git-commit，只要执行git rebase –continue，这样git会继续应用(apply)余下的补丁。如果要舍弃本次衍合，只需要git rebase --abort即可。**切记，一旦分支中的提交对象发布到公共仓库，就千万不要对该分支进行rebase操作**。

我们在使用git pull命令的时候，可以使用--rebase参数，即git pull --rebase。这里表示把你的本地当前分支里的每个提交取消掉，并且把它们临时保存为补丁（这些补丁放到.git/rebase目录中），然后把本地当前分支更新为最新的origin分支，最后把保存的这些补丁应用到本地当前分支上。在使用tortoise的pull的过程中，如果你留意tortoiseGit的日志的话，你就会发现，它使用的就是这种方式来pull最新的提交的。

### git grep

### git blame 找到谁动了某行代码
git blame -L 50,50 init/main.c


### git cherry-pick 合并指定的commitid
`git cherry-pick commitID`
https://www.cnblogs.com/yychuyu/p/13751649.html

### git gc
`git gc [--aggressive] [--auto] [--quiet] [--prune=<date> | --no-prune] [--force]`

1、简单指令：
`git gc`

2、 --aggressive：仔细检查并清理，犹如电脑的全部杀毒，用时较久，一般上100个commit后可以执行，经常执行区别不大：
`git gc --aggressive`

3、 --auto：大概看一下仓库有没有需要整理，如果情况良好，不执行gc：
`git gc --auto`

4、—no-prune：不要整理任何零散的文件：
`git gc -no-prune`

5、–quiet：取消所有进度报告：
`git gc --quiet`

运行git prune命令即可，或者直接运行git gc --prune=now把所有的悬空对象都清空
git gc --prune=now

### 其它
#### 大文件
##### 大文件清理
1. 找出大文件
`git rev-list --objects --all | grep -a "$(git verify-pack -v .git/objects/pack/*.idx | sort -k 3 -n | tail -5 | awk '{print $1}')"`

2. 删除文件

删文件，将 bigfile 换成上面找出的文件名

git filter-branch --force --index-filter \
  'git rm --cached --ignore-unmatch "bigfile"' \
  --prune-empty -- --all

删文件夹，将 wrongdir 换成上面找出的文件夹

git filter-branch --force --index-filter \
  'git rm -r --cached --ignore-unmatch "wrongdir"' \
  --prune-empty -- --all

3. 强制更新远程仓库
(这一步执行了，就真没救了。请确认已备份。)

git push --force --verbose --dry-run
git push --force

##### 大文件清理2

> https://gitee.com/help/articles/4232#article-header2

1. 查看存储库中的大文件
```
git rev-list --objects --all | grep -E `git verify-pack -v .git/objects/pack/*.idx | sort -k 3 -n | tail -10 | awk '{print$1}' | sed ':a;N;$!ba;s/\n/|/g'`

或

git rev-list --objects --all | grep "$(git verify-pack -v .git/objects/pack/*.idx | sort -k 3 -n | tail -15 | awk '{print$1}')"
```

2. 改写历史，去除大文件
> 注意：下方命令中的 path/to/large/files 是大文件所在的路径，千万不要弄错！
```
git filter-branch --tree-filter 'rm -f path/to/large/files' --tag-name-filter cat -- --all
git push origin --tags --force
git push origin --all --force
```
如果在 `git filter-branch` 操作过程中遇到如下提示，需要在 `git filter-branch` 后面加上参数 `-f`
```
Cannot create a new backup.
A previous backup already exists in refs/original/
Force overwriting the backup with -f
```

并告知所有组员，push 代码前需要 pull rebase，而不是 merge，否则会从该组员的本地仓库再次引入到远程库中，导致仓库在此被 Gitee 系统屏蔽。

更加具体的操作可以点击文章 [改写历史，永久删除git库的物理文件](https://my.oschina.net/jfinal/blog/215624?fromerr=ZTZ6c38X) 查看

##### 大文件存储
突破GitHub的限制，使用 git-lfs(Git Large File Storage) 支持单个文件超过100M

1. 下载Git Large File Storage (LFS)：https://git-lfs.github.com/
2. `git lfs install`
3. `git lfs track "*.pdf"` and `git lfs track "*.ppt"`
会出现.gitattributes文件
4. 正常的提交

对这些大文件（默认10M）在git的repository里面只存一个小的文本文件，这个文本文件描述了要去哪里下载对应的二进制文件。

[git lfs migrate](https://github.com/git-lfs/git-lfs/blob/main/docs/man/git-lfs-migrate.1.ronn)

#### git archive
生成一个可供发布的压缩包