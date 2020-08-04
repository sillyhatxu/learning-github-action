# learning-github-actions

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Go-Version](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/sillyhatxu/learning-github-actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/sillyhatxu/learning-github-actions)](https://pkg.go.dev/github.com/sillyhatxu/learning-github-actions)
[![Build and Test](https://github.com/sillyhatxu/learning-github-actions/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)](https://github.com/sillyhatxu/learning-github-actions/actions)
[![codecov](https://codecov.io/gh/sillyhatxu/learning-github-actions/branch/master/graph/badge.svg)](https://codecov.io/gh/sillyhatxu/learning-github-actions)

## 新增workflows

新增文件 `.github/workflows/master.yml`

> 创建后，系统会在Actions中显示build

1. 直接创建

```yaml
name: Build and Test

on:
  push:
    branches: [ master ]

jobs:

  build:
    name: Build
    runs-on: ubuntu-latest
    steps:

      - name: Set up Go 1.x
        uses: actions/setup-go@v2
        with:
          go-version: ^1.14
        id: go

      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: Get dependencies
        run: go mod vendor

      - name: Test
        run: go test -v .

# If it does not has the main function and doesn't need to build.
#      - name: Build
#        run: go build -v .
```

2. 页面操作新增文件

* page 1

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-01.png)

* page 2

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-02.png)

* page 3

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-03.png)


## 添加 badges

### 添加 Made with Go

```yaml
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
```

### 添加 Go Version

```yaml
[![Go-Version](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/<OWNER>/<REPOSITORY>)
```

### 添加 Go Reference

需要release new version

> git tag -a v1.0.0
>
> git push origin v1.0.1
```yaml
[![PkgGoDev](https://pkg.go.dev/badge/github.com/<OWNER>/<REPOSITORY>)](https://pkg.go.dev/github.com/<OWNER>/<REPOSITORY>)
```
github.com/sillyhatxu/learning-github-actions/@latest

https://proxy.golang.org/github.com/sillyhatxu/learning-github-actions/@v/v1.0.0.info

### 添加 build and test badge

```yaml
[![Build and Test](https://github.com/sillyhatxu/learning-github-actions/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)](https://github.com/sillyhatxu/learning-github-actions/actions)
```

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/workflow-name.png)

