# goclone

A enhanced git clone tool, `git clone` GitHub/GitLab repositories into proper destination.

## Example

You can use `goclone` like one of below:

```bash
goclone https://github.com/keaising/goclone
goclone git@github.com:keaising/goclone.git
goclone https://github.com/keaising/goclone.git
```

And `goclone` will do

1. clone the repo to proper destination

```bash
git clone git@github.com:keaising/goclone.git /path/pre-set/github.com/keaising/goclone
```

2. `cd` to the directory where repo is

```base
cd /path/pre-set/github.com/keaising/goclone
```

## How to pre-set path

```bash
git config --global clone.directory /path/pre-set
```
