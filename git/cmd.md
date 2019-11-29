# 常用命令

git clone
# 远程仓库
git remote -v
git remote add shortname url # 添加远程仓库
git remote add origin git@github.com:xxx/learngit.git
git remote rm origin：删除已有的远程仓库；
git fetch shortname #从远程仓库中获得数据，
git remote show origin #查看某个远程仓库
git remote rename origin new_origin #重命名
git remote rm paul #删除


# git push
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

# git log
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


# git reset
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

# Git revert
Git revert用来撤销某次操作，此次操作之前和之后的commit和history都会保留，并且把这次撤销作为一次最新的提交。git revert是提交一个新的版本，将需要revert的版本的内容再反向修改回去，版本会递增，不影响之前提交的内容。
Git revert和git reset都可以进行版本的回退，将工作区回退到历史的某个状态，二者有如下的区别：
- git revert是用一次新的commit来回滚之前的commit，而git reset是直接删除指定的commit（并没有真正的删除，通过git reflog可以找回），这是二者最显著的区别；
- git reset 是把HEAD向后移动了一下，而git revert是HEAD继续前进，只是新的commit的内容和要revert的内容正好相反，能够抵消要被revert的内容；
- 在回滚这一操作上，效果差不多。但是在日后继续merge以前的老版本时有区别。因为git revert是用一次逆向的commit"中和"之前的提交，因此日后合并老的branch时，导致这部分改变不会再次出现；但是git reset是之间把某些commit在某个branch上删除，因而和老的branch再次merge时，这些被回滚的commit应该还会被引入。

https://blog.csdn.net/yxlshk/article/details/79944535
https://git-scm.com/docs/git-revert
# Git stash
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


git rm
git checkout -b dev：git checkout命令加上-b参数表示创建并切换，相当于以下两条命令：
git branch dev
git checkout dev
git branch：查看当前分支，git branch命令会列出所有分支，当前分支前面会标一个*号。
git branch -d dev：删除dev分支了
git merge
git merge dev：git merge命令用于合并指定分支到当前分支。
# git tag
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

# rebase
- 合并多次提交纪录，减少了无用的提交信息
- 分支合并
把一个分支中的修改整合到另一个分支的办法有两种，
第一种是我们常用的git merge操作，
而第二种便是本节要讲的rebase（中文翻译为衍合）。
该命令的原理是，回到两个分支最近的共同祖先，根据当前分支（也就是要进行衍合的分支experiment）后续的历次提交对象（这里只有一个 C3），生成一系列文件补丁，然后以基底分支（也就是主干分支master）最后一个提交对象（C4）为新的出发点，逐个应用之前准备好的补丁文件，最后会生成一个新的合并提交对象（C3'），从而改写 experiment 的提交历史，使它成为 master 分支的直接下游。
一般我们使用rebase的目的，是想要得到一个能在远程分支上干净应用的补丁，比如某些项目你不是维护者，但想帮点忙的话，最好用rebase：先在自己的一个分支里进行开发，当准备向主项目提交补丁的时候，根据最新的 origin/master 进行一次衍合操作然后再提交，这样维护者就不需要做任何整合工作（实际上是把解决分支补丁同最新主干代码之间冲突的责任，化转为由提交补丁的人来解决），只需根据你提供的仓库地址作一次快进合并，或者直接采纳你提交的补丁。

在rebase的过程中，也许会出现冲突。在这种情况，Git会停止rebase并会让你去解决冲突；在解决完冲突后，用git add命令去更新这些内容的索引， 然后，你无需执行git-commit，只要执行git rebase –continue，这样git会继续应用(apply)余下的补丁。如果要舍弃本次衍合，只需要git rebase --abort即可。**切记，一旦分支中的提交对象发布到公共仓库，就千万不要对该分支进行rebase操作**。

我们在使用git pull命令的时候，可以使用--rebase参数，即git pull --rebase。这里表示把你的本地当前分支里的每个提交取消掉，并且把它们临时保存为补丁（这些补丁放到.git/rebase目录中），然后把本地当前分支更新为最新的origin分支，最后把保存的这些补丁应用到本地当前分支上。在使用tortoise的pull的过程中，如果你留意tortoiseGit的日志的话，你就会发现，它使用的就是这种方式来pull最新的提交的。

# git grep

