# Git 常用命令速查表（美化版）
本文整理 Git 日常开发核心命令，按功能模块分类，搭配清晰格式和实用示例，方便快速查阅使用。

---

## 📁 一、仓库初始化与克隆
| 命令 | 说明 | 示例 |
|------|------|------|
| `git init <目录名>` | 初始化本地 Git 仓库（默认当前目录） | `git init my_project`（创建并初始化 `my_project` 仓库） |
| `git clone -b <分支名> <仓库地址>` | 克隆远程仓库指定分支到本地 | `git clone -b develop https://github.com/username/repo.git` |
| `git clone <仓库地址>` | 克隆远程仓库默认分支（通常是 `main`） | `git clone https://github.com/username/repo.git` |

---

## 📤 二、代码暂存与提交
| 命令 | 说明 | 示例 |
|------|------|------|
| `git add <文件>` | 将指定文件添加到暂存区 | `git add main.go`（添加单个文件） |
| `git add .` | 将当前目录所有修改文件添加到暂存区 | `git add .`（批量添加，开发常用） |
| `git commit -m "<提交信息>"` | 提交暂存区文件到本地仓库（规范填写信息） | `git commit -m "feat: 新增用户登录功能"` |
| `git status` | 查看工作区/暂存区状态（红色=未暂存，绿色=已暂存） | `git status`（快速排查未提交文件） |
| `git reset <文件>` | 取消指定文件的暂存状态 | `git reset main.go` |
| `git reset .` | 取消当前目录所有文件的暂存状态 | `git reset .`（批量取消暂存） |
| `git commit --amend` | 修正上一次提交（补充文件/修改描述） | `git add 漏提交文件 && git commit --amend` |

---

## 🌿 三、分支管理（协作核心）
| 命令 | 说明 | 示例 |
|------|------|------|
| `git branch` | 查看本地所有分支（当前分支标 `*`） | `git branch` |
| `git branch -a` | 查看本地+远程所有分支 | `git branch -a` |
| `git branch <分支名>` | 创建新分支（基于当前分支） | `git branch feature/login`（创建登录功能分支） |
| `git checkout <分支名>` | 切换到指定分支（Git 2.23+ 推荐用 `switch`） | `git checkout main`（切换到主分支） |
| `git switch <分支名>` | 切换分支（语法更直观，替代 `checkout`） | `git switch feature/login` |
| `git switch -c <分支名>` | 创建并切换到新分支（一步完成） | `git switch -c feature/pay`（创建支付分支并切换） |
| `git merge <分支名>` | 将指定分支合并到当前分支 | `git switch main && git merge feature/login`（合并功能分支到主分支） |
| `git branch -d <分支名>` | 删除本地已合并的分支（安全删除） | `git branch -d feature/login` |
| `git branch -D <分支名>` | 强制删除本地分支（未合并也可删除） | `git branch -D feature/old`（删除废弃分支） |

---

## 🌐 四、远程仓库同步（协作必备）
| 命令 | 说明 | 示例 |
|------|------|------|
| `git remote -v` | 查看远程仓库绑定关系（显示拉取/推送地址） | `git remote -v` |
| `git remote add <远程名> <仓库地址>` | 绑定本地仓库到远程仓库（首次关联） | `git remote add origin https://github.com/username/repo.git` |
| `git pull <远程名> <分支名>` | 拉取远程分支更新并合并到本地 | `git pull origin main`（拉取远程主分支更新） |
| `git pull` | 拉取当前分支绑定的远程分支更新（已绑定 `upstream`） | `git pull`（简化写法） |
| `git push <远程名> <分支名>` | 推送本地分支到远程仓库 | `git push origin feature/login`（推送功能分支） |
| `git push -u <远程名> <分支名>` | 绑定本地与远程分支（后续可直接 `git push`） | `git push -u origin feature/login` |
| `git fetch` | 拉取远程所有分支更新但不合并（先查看差异） | `git fetch origin`（安全同步远程更新） |

---

## ⏳ 五、日志与版本回溯
| 命令 | 说明 | 示例 |
|------|------|------|
| `git log` | 查看提交日志（按时间倒序，详细信息） | `git log` |
| `git log --oneline` | 简洁单行显示日志（提交ID前7位+描述） | `git log --oneline`（快速浏览提交记录） |
| `git log --graph` | 查看分支合并图（直观展示分支流向） | `git log --graph --oneline`（简洁合并图） |
| `git log --graph --all` | 查看所有分支的合并图 | `git log --graph --oneline --all` |
| `git reset --hard <提交ID>` | 回溯到指定版本（覆盖本地修改，谨慎使用） | `git reset --hard a1b2c3d`（`a1b2c3d` 为提交ID前7位） |
| `git reflog` | 查看所有操作记录（含 `reset` 后的版本，用于恢复） | `git reflog`（找回误删/误回溯的版本） |

---

## 🛠️ 六、其他常用辅助命令
| 命令 | 说明 | 示例 |
|------|------|------|
| `git stash` | 暂存工作区未提交的修改（切换分支时用） | `git stash`（暂存当前修改） |
| `git stash pop` | 恢复最近一次暂存并删除暂存记录 | `git stash pop`（开发中断后恢复修改） |
| `git stash apply` | 恢复最近一次暂存（保留暂存记录） | `git stash apply`（需多次恢复时使用） |
| `git stash list` | 查看所有暂存记录 | `git stash list`（查看暂存ID和描述） |
| `git clean -fd` | 删除工作区未跟踪的文件/目录（如编译产物） | `git clean -fd`（`-f`=文件，`-d`=目录） |
| `git diff` | 查看工作区与暂存区的文件差异 | `git diff main.go`（查看单个文件修改） |
| `git diff --cached` | 查看暂存区与本地仓库的差异 | `git diff --cached`（提交前确认修改内容） |
| `git restore <文件>` | 丢弃工作区未暂存的修改 | `git restore main.go` |
| `git restore .` | 丢弃当前目录所有未暂存的修改 | `git restore .` |
| `git restore --staged <文件>` | 取消指定文件的暂存（等价于 `git reset <文件>`） | `git restore --staged main.go` |

---

### ✨ 使用说明
1. 命令示例均为开发高频场景，可直接复制使用
2. 尖括号 `<>` 中的内容为占位符，需替换为实际值（如分支名、仓库地址）
3. 标注「谨慎使用」的命令（如 `git reset --hard`），执行前建议用 `git status` 确认本地修改是否已备份
4. 推荐配合 Git 别名使用（如 `git lg = git log --graph --oneline --all`），提升效率