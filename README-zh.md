# learning-github-action

[![Build and Test](https://github.com/sillyhatxu/learning-github-action/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)](https://github.com/sillyhatxu/learning-github-action/actions)

## 新增workflows

### 新增文件 `.github/workflows/master.yml`

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

### 或者你也可以在页面操作



