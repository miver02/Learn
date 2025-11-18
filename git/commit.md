# Commit 提交规范

## feat 最实用的延伸价值
通过工具（如 standard-version、semantic-release），可以扫描所有含 feat 的提交，自动生成结构化的变更日志


## 关联语义化版本
主版本号（1.x.x）：不兼容的 API 变更（通常配合 feat! 或 BREAKING CHANGE）。
次版本号（x.2.x）：向后兼容的功能新增（对应 feat 提交）。
修订号（x.x.3）：向后兼容的问题修复（对应 fix 提交）。

## 提交类型
feat	新增功能（向后兼容）	feat: 支持微信登录
fix	修复 bug	fix: 修复登录失败的问题
docs	仅修改文档（如 README）	docs: 更新API文档
refactor	重构代码（不新增功能不修复 bug）	refactor: 优化用户模块代码结构
style	代码格式调整（如缩进、空格）	style: 格式化代码（ESLint）
test	新增 / 修改测试用例	test: 补充登录功能单元测试
chore	构建 / 依赖 / 工具配置修改	chore: 升级npm依赖到v10