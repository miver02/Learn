# 新建项目开发
git init → git add . → git commit -m "init" → git remote add origin → git push -u origin main

# 参与已有项目
git clone <url> → git switch -c feature/xxx → 开发 → git add . → git commit → git push

# 合并功能到主分支
git switch main → git pull origin main → git merge feature/xxx → git push origin main

# 修复线上 bug
git switch main → git switch -c hotfix/bug123 → 修复 → git commit → git push → 合并到 main

# 定位 bug 提交
git bisect start → git bisect bad → git bisect good <无bug版本> → 测试标记 → git bisect reset