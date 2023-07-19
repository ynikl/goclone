# goclone

A enhanced git clone tool, `git clone` GitHub/GitLab repositories into proper destination.

## Example

You can use `goclone` like one of below:

```bash
goclone https://github.com/keaising/goclone
goclone git@github.com:keaising/goclone.git
goclone https://github.com/keaising/goclone.git
```

And `goclone` will clone the repo to proper destination by

```bash
git clone git@github.com:keaising/goclone.git /path/pre-set/github.com/keaising/goclone
```

## How to pre-set path

```bash
git config --global clone.directory /path/pre-set
```

## In a more fluent way

`goclone` runs in a child process so that it is impossible to `cd` to the directory where repo is. 

But we can seek help from shell. 

Please add code below to your shell scripts such as `.bash_profile` or `.zshrc`, then reopen shell or terminal.

```bash
glone () {
	goclone $1 | tee /tmp/goclone
	cd $(cat /tmp/goclone | head -n 1 | awk '{print $4}')
}
```

And enjoy a wonderful workflow like:

```bash
~/code/github.com
Δ glone git@github.com:keaising/goclone.git
git clone git@github.com:keaising/goclone.git /Users/taiga/code/github.com/keaising/goclone
Cloning into '/Users/taiga/code/github.com/keaising/goclone'...
remote: Enumerating objects: 50, done.
remote: Counting objects: 100% (50/50), done.
remote: Compressing objects: 100% (30/30), done.
remote: Total 50 (delta 19), reused 34 (delta 11), pack-reused 0
Receiving objects: 100% (50/50), 11.92 KiB | 3.97 MiB/s, done.
Resolving deltas: 100% (19/19), done.

~/code/github.com/keaising/goclone master 5s
Δ
```

## CHANGELOG

修改代码调整，适应以下使用方式, 默认拷贝到 `GOPATH` 路径下

``` bash
go run . https://github.com/keaising/goclone 
```

结果路径为:

``` bash
~/go/src/github.com/keaising/goclone/  
```
