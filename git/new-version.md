为了让新命令更易读、更具实用性，我用「功能说明+分步示例+效果标注」的结构美化，搭配清晰排版和图标，方便快速上手：

# Git 较新命令速查表（Git 2.0+ 推荐）
整理 Git 后续版本新增的高效命令，解决旧命令痛点（如功能过载、操作繁琐），提升开发效率。

---

## 🔄 `git restore` - 精准恢复文件（替代 `git checkout -- 文件`）
### 核心作用
专门用于恢复文件状态（丢弃修改、取消暂存），语法直观，避免 `git checkout` 的功能混淆。

| 命令 | 说明 | 示例 |
|------|------|------|
| `git restore <文件>` | 丢弃工作区未暂存的修改 | `git restore main.go`（恢复 `main.go` 到最近一次提交状态） |
| `git restore --staged <文件>` | 取消文件暂存（等价于 `git reset <文件>`） | `git restore --staged main.go`（将已 `add` 的文件撤回工作区） |
| `git restore --source=<提交ID> <文件>` | 从指定提交版本恢复文件 | `git restore --source=a1b2c3d main.go`（从提交 `a1b2c3d` 恢复 `main.go`） |

---

## 📁 `git sparse-checkout` - 稀疏检出（只拉取仓库部分目录）
### 核心作用
针对大型仓库（含大量历史资源/无关目录），仅拉取需要的目录，节省磁盘空间和拉取时间。

### 分步操作示例
```bash
# 1. 初始化本地仓库并关联远程
git init my-repo && cd my-repo
git remote add origin https://github.com/username/large-repo.git

# 2. 开启稀疏检出模式（--cone 表示仅拉取根目录+指定目录，提速）
git sparse-checkout init --cone

# 3. 指定需要拉取的目录（多个目录用空格分隔）
git sparse-checkout set "src/" "docs/" "config/"

# 4. 拉取远程代码（仅下载指定目录）
git pull origin main
```

### 关键说明
- 查看已配置的稀疏目录：`git sparse-checkout list`
- 新增拉取目录：`git sparse-checkout add "test/"`
- 取消稀疏检出（拉取全量）：`git sparse-checkout disable`

---

## 🌿 `git worktree` - 多工作区管理（无需重复克隆仓库）
### 核心作用
同时在多个分支开发（如一边开发新功能，一边修复旧版本 Bug），无需多次克隆仓库，共享本地仓库的提交历史。

### 常用命令示例
| 操作 | 命令 | 说明 |
|------|------|------|
| 创建新工作区 | `git worktree add <工作区路径> <分支名>` | `git worktree add ../repo-bugfix feature/bugfix`（在仓库外创建 `repo-bugfix` 工作区，关联 `feature/bugfix` 分支） |
| 进入工作区 | `cd <工作区路径>` | `cd ../repo-bugfix`（进入新工作区，修改不影响原仓库） |
| 查看所有工作区 | `git worktree list` | 显示所有关联的工作区路径、分支名、提交ID |
| 删除工作区 | `git worktree remove <工作区路径>` | `git worktree remove ../repo-bugfix`（安全删除工作区，不影响原仓库） |
| 清理失效工作区 | `git worktree prune` | 移除已删除的工作区关联记录 |

---

## 🔍 `git bisect` - 二分查找定位 Bug 提交
### 核心作用
当代码突然出现 Bug 但不确定引入源头时，用二分法快速缩小范围，定位到首次引入 Bug 的提交（比逐行查看日志高效10倍）。

### 分步操作示例
```bash
# 1. 开始二分查找流程
git bisect start

# 2. 标记当前版本为「有Bug」（当前分支的最新提交）
git bisect bad

# 3. 标记一个已知「无Bug」的版本（如历史稳定版本、发布版本）
git bisect good v1.0  # 或用提交ID：git bisect good a1b2c3d

# 4. Git 自动切换到中间版本，手动测试代码后标记
git bisect good  # 若当前版本无Bug，继续查找后续提交
# 或 git bisect bad  # 若当前版本有Bug，查找之前的提交

# 5. 重复步骤4，直到 Git 定位到 Bug 提交（显示类似：b3c4d5e is the first bad commit）
# 6. 结束查找，回到原分支
git bisect reset
```

### 关键说明
- 标记后 Git 会自动计算中间版本，无需手动切换
- 若测试过程中想暂停：`git bisect suspend`，恢复：`git bisect resume`
- 适合 Bug 可复现、且能快速验证的场景

---

## ✨ 使用建议
1. 优先用 `git restore`/`git switch` 替代 `git checkout`，降低学习成本
2. 大型仓库必用 `git sparse-checkout`，避免拉取无关资源
3. 多分支并行开发时，`git worktree` 比频繁切换分支更高效
4. 定位历史 Bug 时，`git bisect` 是核心工具，大幅节省排查时间

所有命令均兼容 Git 2.23+ 版本（目前主流环境已默认支持），可通过 `git --version` 查看本地 Git 版本。