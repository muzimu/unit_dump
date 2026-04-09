# unit_dump

扫描指定目录下所有 Go 测试文件，输出每个 `Test*` 函数的名称（每行一个）。

## 安装

```bash
go install github.com/lizijun/unit_dump@latest
```

或从源码构建：

```bash
cd /Library/workspace/unit_dump
go build -o unit_dump .
```

## 用法

```bash
# 扫描当前目录
unit_dump

# 扫描指定目录
unit_dump ./path/to/project
```

## 配合 fzf 使用

```bash
alias gout='go test -v -run $(unit_dump | fzf)'
```

执行 `gout` 后，通过 fzf 交互式选择要运行的测试函数。
