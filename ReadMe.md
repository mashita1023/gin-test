# 環境
- gin
- mariaDB or mySQL
- flutter

# はじめる
```
$ make up
```

# adminer
server: `db`
user: `root`
password: `pwd`
databasse: `test`

# goのmodule追加
```
$ make get GET=${module}
$ make tidy
```

# Makefileについて
proxy環境の場合のみ以下のように実行する
```
$ make build -f proxy.mk
```
