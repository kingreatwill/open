常用命令
git clone
git remote add origin git@github.com:xxx/learngit.git
git remote rm origin：删除已有的远程仓库；
git push
git push -u origin master：由于远程库是空的，我们第一次推送master分支时，加上了-u参数，Git不但会把本地的master分支内容推送的远程新的master分支，还会把本地的master分支和远程的master分支关联起来，在以后的推送或者拉取时就可以简化命令。
git push origin master
git init
git add <file/directory>
git add -A 添加所有
git commit -m <注释>
git status
git diff
git log：提交历史
git log --pretty=oneline
git log --graph：可以看到分支合并图。
git reset
git reset --hard HEAD^
git reset --hard 1094a
git reset HEAD <file>：git reset命令既可以回退版本，也可以把暂存区的修改回退到工作区。当我们用HEAD时，表示最新的版本。
git reflog：命令历史
git checkout
git checkout -- file：丢弃工作区的修改，让这个文件回到最近一次git commit或git add时的状态。命令中的--很重要，没有--，就变成了“切换到另一个分支”的命令。
git rm
git checkout -b dev：git checkout命令加上-b参数表示创建并切换，相当于以下两条命令：
git branch dev
git checkout dev
git branch：查看当前分支，git branch命令会列出所有分支，当前分支前面会标一个*号。
git branch -d dev：删除dev分支了
git merge
git merge dev：git merge命令用于合并指定分支到当前分支。
git tag
git tag：查看所有标签
git tag v1.0：在当前分支的HEAD上打标签
git tag v0.9 f52c633：在当前分支的指定commit上打标签
git tag -d v0.1：删除指定标签
git push origin <tagname>：推送某个标签到远程
git push origin --tags：一次性推送全部尚未推送到远程的本地标签