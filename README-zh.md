# learning-github-actions

[![Build and Test](https://github.com/sillyhatxu/learning-github-actions/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)](https://github.com/sillyhatxu/learning-github-actions/actions)

## 新增workflows

新增文件 `.github/workflows/master.yml`

创建后，系统会在Actions中显示build

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

      - name: Set up Go 1.14.x
        uses: actions/setup-go@v2
        with:
          go-version: ^1.14
        id: go

      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: Get dependencies
        run: go mod vendor

      - name: Build
        run: go build -v .

      - name: Test
        run: go test -v .
```

2. 页面操作新增文件

* page 1

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-01.png)

* page 2

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-02.png)

* page 3

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/page-add-workflows-03.png)





## 添加 badges

### 添加 build and test badge

![](https://github.com/sillyhatxu/learning-github-actions/blob/master/asset/workflow-name.png)
