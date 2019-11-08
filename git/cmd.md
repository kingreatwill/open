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


# 取消暂存的文件
git reset
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

# Git stash
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