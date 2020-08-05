# learning-github-actions

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Go-Version](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/sillyhatxu/learning-github-actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/sillyhatxu/learning-github-actions)](https://pkg.go.dev/github.com/sillyhatxu/learning-github-actions)
[![Build and Test](https://github.com/sillyhatxu/learning-github-actions/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)](https://github.com/sillyhatxu/learning-github-actions/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/sillyhatxu/learning-github-actions)](https://goreportcard.com/report/github.com/sillyhatxu/learning-github-actions)
[![codecov](https://codecov.io/gh/sillyhatxu/learning-github-actions/branch/master/graph/badge.svg)](https://codecov.io/gh/sillyhatxu/learning-github-actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://choosealicense.com/licenses/mit/)
[![Release](https://img.shields.io/github/release/sillyhatxu/learning-github-actions.svg?style=flat-square)](https://github.com/sillyhatxu/learning-github-actions/releases)

## [中文文档](https://github.com/sillyhatxu/learning-github-actions/blob/master/README-zh.md)

## 1. Add workflows

create file `.github/workflows/master.yml`

> After the first creation, The Actions will be auto build.

1) Create file `master.yml` directly

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

2) Create file `master.yml` by github page

* page 1

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-01.png)

* page 2

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-02.png)

* page 3

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-03.png)


## 2. Add badges

### 1) Add Made with Go

```yaml
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
```

### 2) Add Go Version

```yaml
[![Go-Version](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/<OWNER>/<REPOSITORY>)
```

### 4) Add Go Reference

*You must be created with License file*

* page 1

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-license-01.png)

* page 2

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-license-02.png)

* page 3

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-license-03.png)

* page 4

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-license-04.png)

> Sometimes it is not automatically created, and officials offer two solutions.

* First：`I used this method, but still not working`

> Making a request to proxy.golang.org for the module version, to any endpoint specified by the Module proxy protocol. 
> For example: https://proxy.golang.org/example.com/my/module/@v/v1.0.0.info

    curl https://proxy.golang.org/github.com/sillyhatxu/learning-github-actions/@latest
    {"Version":"v1.0.0","Time":"2020-08-04T15:50:54Z"}
    curl https://proxy.golang.org/github.com/sillyhatxu/learning-github-actions/@v/v1.0.0.info
    {"Version":"v1.0.0","Time":"2020-08-04T15:50:54Z"}

* Second：`Create another golang project，after go mod init. Used go get command create Go Reference`

> Downloading the package via the go command. 
> For example: GOPROXY=https://proxy.golang.org GO111MODULE=on 
> go get example.com/my/module@v1.0.0
    
    create new project xxxxxx
    go mod init xxxxxx
    go get github.com/sillyhatxu/learning-github-actions/@v1.0.2

```yaml
[![PkgGoDev](https://pkg.go.dev/badge/github.com/<OWNER>/<REPOSITORY>)](https://pkg.go.dev/github.com/<OWNER>/<REPOSITORY>)
```

### 4) Add build and test badge

```yaml
[![Build and Test](https://github.com/sillyhatxu/learning-github-actions/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)](https://github.com/sillyhatxu/learning-github-actions/actions)
```

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/workflow-name.png)

### 5) Add coverage badge

* Login [codecov](https://codecov.io/)
* Add new repository

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-coverage-01.png)

* Choose a new repository below

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-coverage-02.png)

* Copy token

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/create-coverage-03.png)

* Return Github Page `project->Settings->Secrets->New Secret`

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

### 6) Add go report

```yaml
[![Go Report Card](https://goreportcard.com/badge/github.com/<OWNER>/<REPOSITORY>)](https://goreportcard.com/report/github.com/<OWNER>/<REPOSITORY>)
```

### 7) Add release version

```yaml
[![Release](https://img.shields.io/github/release/<OWNER>/<REPOSITORY>.svg?style=flat-square)](https://github.com/<OWNER>/<REPOSITORY>/releases)
```

### 8) Add MIT License

```yaml
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://choosealicense.com/licenses/mit/)
```