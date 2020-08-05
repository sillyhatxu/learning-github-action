# learning-github-actions

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Go-Version](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/sillyhatxu/learning-github-actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/sillyhatxu/learning-github-actions)](https://pkg.go.dev/github.com/sillyhatxu/learning-github-actions)
[![Build and Test](https://github.com/sillyhatxu/learning-github-actions/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)](https://github.com/sillyhatxu/learning-github-actions/actions)
[![codecov](https://codecov.io/gh/sillyhatxu/learning-github-actions/branch/master/graph/badge.svg)](https://codecov.io/gh/sillyhatxu/learning-github-actions)
[![Release](https://img.shields.io/github/release/sillyhatxu/learning-github-actions.svg?style=flat-square)](https://github.com/sillyhatxu/learning-github-actions/releases)

## 1. 新增workflows

新增文件 `.github/workflows/master.yml`

> 创建后，系统会在Actions中显示build

1) 直接创建

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

2) 页面操作新增文件

* page 1

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-01.png)

* page 2

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-02.png)

* page 3

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-03.png)


## 2. 添加 badges

### 1) 添加 Made with Go

```yaml
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
```

### 2) 添加 Go Version

```yaml
[![Go-Version](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/<OWNER>/<REPOSITORY>)
```

### 4) 添加 Go Reference

*需要创建License*

* page 1

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-license-01.png)

* page 2

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-license-02.png)

* page 3

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-license-03.png)

* page 4

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-license-04.png)

> 有时不会自动创建，官方给出两种解决方案。

* 第一种：`不知名原因，只返回了json，但没有更新Go Reference Doc`

> Making a request to proxy.golang.org for the module version, to any endpoint specified by the Module proxy protocol. 
> For example: https://proxy.golang.org/example.com/my/module/@v/v1.0.0.info

    curl https://proxy.golang.org/github.com/sillyhatxu/learning-github-actions/@latest
    {"Version":"v1.0.0","Time":"2020-08-04T15:50:54Z"}
    curl https://proxy.golang.org/github.com/sillyhatxu/learning-github-actions/@v/v1.0.0.info
    {"Version":"v1.0.0","Time":"2020-08-04T15:50:54Z"}

* 第二种：`需要使用另一个golang的项目，在go mod init 后，使用 go get 命令来发布`
> Downloading the package via the go command. 
> For example: GOPROXY=https://proxy.golang.org GO111MODULE=on 
> go get example.com/my/module@v1.0.0
    
    create new project xxxxxx
    go mod init xxxxxx
    go get github.com/sillyhatxu/learning-github-actions/@v1.0.2

```yaml
[![PkgGoDev](https://pkg.go.dev/badge/github.com/<OWNER>/<REPOSITORY>)](https://pkg.go.dev/github.com/<OWNER>/<REPOSITORY>)
```

### 4) 添加 build and test badge

```yaml
[![Build and Test](https://github.com/sillyhatxu/learning-github-actions/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)](https://github.com/sillyhatxu/learning-github-actions/actions)
```

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/workflow-name.png)

### 5) 添加 coverage badge

* 登陆 [codecov](https://codecov.io/)
* Add new repository

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-coverage-01.png)

* Choose a new repository below

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-coverage-02.png)

* Copy token

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-coverage-03.png)

* 回到Github `project->Settings->Secrets->New Secret`

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-coverage-04.png)

* Create Secret `CODECOV_TOKEN=xxxxx` paste codecov token

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-coverage-05.png)

* Create workflows `coverage.yml`

```yaml
name: Coverage

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

      - name: Create coverage file
        run: |
          set -e
          echo "" > coverage.txt

          for d in $(go list ./... | grep -v vendor); do
              go test -race -coverprofile=profile.out -covermode=atomic "$d"
              if [ -f profile.out ]; then
                  cat profile.out >> coverage.txt
                  rm profile.out
              fi
          done

      - name: Coverage
        run: bash <(curl -s https://codecov.io/bash) -t ${{ secrets.CODECOV_TOKEN }}
```

### 6) 添加 go report


### 6) 添加 release version

```yaml
[![Release](https://img.shields.io/github/release/<OWNER>/<REPOSITORY>.svg?style=flat-square)](https://github.com/<OWNER>/<REPOSITORY>/releases)
```