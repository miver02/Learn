# Git 较新命令
## 1, git restore - 恢复文件
git restore main.go  # 恢复工作区的main.go（丢弃未暂存的修改）
git restore --staged main.go  # 取消暂存（将已add的文件撤回工作区，等价于git reset main.go）
git restore --source=a1b2c3d main.go  # 从指定提交ID恢复main.go

## 2, git sparse-checkout - 稀疏检出（只拉取仓库部分文件 / 目录）
### 1. 初始化稀疏检出
git init my-repo && cd my-repo
git remote add origin https://github.com/username/large-repo.git
git sparse-checkout init --cone  # --cone=只拉取根目录文件+指定目录

### 2. 指定要拉取的目录
git sparse-checkout set "src/" "docs/"  # 只拉取src和docs目录

### 3. 拉取代码
git pull origin main

## 3, git worktree - 多工作区管理（无需多次克隆仓库）
### 1. 在当前仓库外创建一个新工作区（关联feature/bugfix分支）
git worktree add ../repo-bugfix feature/bugfix

### 2. 进入新工作区开发（修改不会影响原仓库）
cd ../repo-bugfix

### 3. 查看所有工作区
git worktree list

### 4. 删除工作区
git worktree remove ../repo-bugfix

## 4, git bisect - 二分查找定位 bug 提交（快速找哪次提交引入 bug）
git bisect start  # 开始二分查找
git bisect bad  # 标记当前版本为“有bug”
git bisect good v1.0  # 标记v1.0版本为“无bug”

### Git自动切换到中间版本，手动测试后标记good/bad
git bisect good  # 若当前版本无bug
### 或 git bisect bad  # 若当前版本有bug

### 找到bug提交后，结束查找
git bisect reset