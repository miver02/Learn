# 1, 常用基础命令
1, git init # 初始化本地git仓库 eg: git init my_project
2, git clone -b <分支名> <url> # 克隆远程分支

# 2, 代码暂存与提交
1, git add . or <文件> # 将文件加入暂存区
2, git commit -m "<信息>" # 提交暂存区文件到本地仓库
3, git status # 查看工作区 / 暂存区状态
4, git reset . or <文件> # 取消暂存

# 3, 分支管理
1, git branch # 查看本地分支
2, git branch <分支名> # 创建新分支
3, git checkout <分支名> # 切换分支
4, git switch <分支名> # 切换分支
5, git switch -c <分支名> # 创建并切换到新分支
6, git merge <分支名> # 合并分支到当前分支
7, git branch -d <分支名> # 删除本地分支
8, git branch -D <分支名> # 强制删除本地分支

# 4, 远程仓库同步
1, git remote -v # 	查看远程仓库地址
2, git remote add <远程名> <url> # 绑定远程仓库（本地仓库首次关联远程）
3, git pull # 	拉取远程分支更新并合并到本地
4, git push <远程名> <分支名> # 推送本地分支到远程仓库
5, git push -u <远程名> <分支名> # 绑定本地与远程分支（后续可直接git push）
6, git fetch # 拉取远程分支更新但不合并

# 5, 日志与版本回溯
1, git log # 查看提交日志（按时间倒序）
2, git log --graph # 查看分支合并图（直观看到分支流向） eg: git log --graph --oneline（简洁合并图）
3, git reset --hard <提交ID> # 回溯到指定版本（谨慎使用，会覆盖本地修改）
4, git reflog # 查看所有操作记录（包括 reset 后的版本，用于恢复

# 6, 其他常用辅助命令
1, git stash # 暂存工作区未提交的修改（切换分支时用） eg: 	git stash（暂存）、git stash pop（恢复并删除暂存）、git stash apply（恢复最近一次暂存并删除暂存）
2, git stash list # 查看所有暂存记录
3, git clean -fd # 删除工作区未跟踪的文件
4, git diff # 查看文件修改差异（工作区 vs 暂存区）
5, git diff --cached # 查看暂存区 vs 本地仓库的差异
5, git restore . or <文件> # 丢弃未暂存的修改
