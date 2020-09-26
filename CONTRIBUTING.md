# Contributing to Pluggabl

First off, thanks for taking the time to contribute!

## Table of contents

1. [Code style](#code-style)
1. [Commit messages](#commit-messages)
1. [Branches](#branches)
1. [Pull requests](#pull-requests)
1. [Tests](#tests)

---

## Code style

Read [Effective Go](https://golang.org/doc/effective_go.html) and try to check if everything is written according to standards. Make sure that your commits passed `gofmt -e -d .` in the GitHub workflow.

---

## Commit messages

As for now, I am trying to be as consistient with commit messages as it is possible. As you can see in commit history, it was not the case since the begining. Now the encouraged way to commit into this repository, is to use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

Examples:

```git
feat: added new type of error code
fix: removed typo in UmServer struct comment
docs: created contributing file
```

---

## Branches

As you can see, there are not so many branches (at the time when I'm writing it, there are three - [main](https://github.com/Shanduur/pluggabl), [develop](https://github.com/Shanduur/pluggabl/tree/develop) and [2020/09/21](https://github.com/Shanduur/pluggabl/tree/2020/09/21)). The convention used in this project is to work with [Trunk Based Developement](https://trunkbaseddevelopment.com/), where **develop** branch acts as the **trunk**, **main** is the current release.

Every other release branch is named using [this pattern](https://regex101.com/r/NbX5nY/2):
```regexp
R_[0-9]\.[0-9]{0,3}
```

For example:

```
R_1.02
```

Branches containing your work that will be merged into the **develop** should be named with pattern:

```regexp
YYYY/MM/DD-YourName
```

For example:


```
2006/12/27-Shanduur
```

---

## Pull requests

Every pull request **must be** reviewed before merging, and pass all the tests.

---

## Tests

You don't have to write tests for every part of your code, but **it's good practice to do so**. If you are adding part of code that does not interact of database I strongly encourage you to write at least **smoke** tests and **sanity** tests for the most important functionality you added.