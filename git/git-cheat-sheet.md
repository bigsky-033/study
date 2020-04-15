# Cheat Sheet - Git

- unstaging file

```bash
git reset HEAD ${file-name}
```

- discard changes file

```bash
git checkout -- ${file-name}
```

- rename file

```bash
#1, using git mv
git mv ${source} ${destination}

#2, using bash's mv
mv ${source} ${destination}
git add -A

# revert rename file
git mv ${new-file-name} ${old-file-name}
```

- remove file

```bash
## workflow of remove & restore file using git rm

#1. remove file using git rm
# file will be deleting in filesystem
# & delete file will be added to staging area
git rm ${file-name}

#2. unstaging deletion
# after unstaging, file is not present on filesystem
git reset HEAD ${file-name}

#3. restore file on filesystem
git checkout -- ${file-name}

## workflow of remove & restore file using bash rm

#1. remove file
# file is deleting on filesystem but not staged
rm ${file-name}

# if you want to added deletion to staging area
git add -A
# or
git add -u ${file-name}


# unstaging & restore it to filesystem if same as above
```

- git log

```bash
# basic
git log

# short commit it
git log --abbrev-commit

# graph view
git log --oneline --graph --decorate

# range
git log aa6ce29...f2a1302

# data based search
git log --since="3 days ago"

# history of individual file
git log -- ${file-name}

# show commit details
git show ${commit-id}
```

- comparison

```bash
# compare between local & staging
git diff

# compare between local & HEAD
git diff HEAD

# compare betweeen stage & HEAD
git diff --staged HEAD

# compare just one file
git diff -- ${file-name}

# compare two commits
git diff ${commit-1-id} ${commit-2-id}
git diff HEAD HEAD^

# compare with remote (example)
git diff ${commit-id} origin/master
```

- branching

```bash
# list branch
git branch -a

# create branch
git branch ${branch-name}

# switch branch
git checkout ${target-branch-name}

# rename branch
# m: move
git branch -m ${old-branch-name} ${new-branch-name}

# delete branch
git branch -d ${branch-name}

# create branch & checkout
git checkout -b ${branch-name}
```

- merging

```bash
# process example for happy path
# 1. checkout to branch which wants to merge feature branch
# e.g. master
git checkout master

# 2. checkout differences between two branches
git diff master feature-branch

# 3-1. do merge (fast-forward-merge)
# e.g. feature-branch
git merge ${branch-name-want-merged-into-current-branch}

# 3-2. disable fast-forward merge
# keep branch history & create another commit for merge
git merge ${branch-name-want-merged-into-current-branch} --no-ff

# 4. delete merged branch
# e.g. feature-branch
git branch -d ${merged-branch}
```

- merging - resolve conflicts

```bash
# trying to merge & conflict happens
git merge ${conflicting-branch-name}

# conflict messages
Auto-merging ${conflict contents}
CONFLICT (content): Merge conflict ${conflict contents}
Automatic merge failed; fix conflicts and then commit the result.

# check how does the current status
git status

# show un merged paths
You have unmerged paths.
        (fix conflicts and run "git commit")
        (use "git merge --abort" to abort the merge)

Unmerged paths:
        (use "git add <file>..." to mark resolution)

        both modified:

# resolving conflicts anyway...

# add both modified files
git add ${both-modified-files}

# do commit
git commit
```

- rebase

```bash
# simple rebase scenario
# after this, feature-branch's base will be ${source-branch-we-want-to-rebase}
# simple rebase will result similiar with fast forward merge

git checkout ${feature-branch}
# mostly, source branch will master
git rebase ${source-branch-we-want-to-rebase}

# doing some work on featur-branch
git checkout master
git merge ${feature-branch}
```

- rebase - abort

```bash
# rebase conflict scenario
# let's assume master branch & feature branch have conflicts

git checkout ${feature-branch}
git rebase master

# error messsage example
First, rewinding head to replay your work on top of it...
Applying: edit simpe.html 2 - conflict
Using index info to reconstruct a base tree...
M	simple.html
Falling back to patching base and 3-way merge...
Auto-merging simple.html
CONFLICT (content): Merge conflict in simple.html
error: Failed to merge in the changes.
Patch failed at 0001 edit simpe.html 2 - conflict
hint: Use 'git am --show-current-patch' to see the failed patch
Resolve all conflicts manually, mark them as resolved with
"git add/rm <conflicted_files>", then run "git rebase --continue".
You can instead skip this commit: run "git rebase --skip".
To abort and get back to the state before "git rebase", run "git rebase --abort".

# abort rebase & approach another way
git rebase --abort
```

- rebase - resolve conflicts

```bash
# rebase conflict scenario
# let's assume master branch & feature branch have conflicts

git checkout ${feature-branch}
git rebase master

# error messsage example
First, rewinding head to replay your work on top of it...
Applying: edit simpe.html 2 - conflict
Using index info to reconstruct a base tree...
M	simple.html
Falling back to patching base and 3-way merge...
Auto-merging simple.html
CONFLICT (content): Merge conflict in simple.html
error: Failed to merge in the changes.
Patch failed at 0001 edit simpe.html 2 - conflict
hint: Use 'git am --show-current-patch' to see the failed patch
Resolve all conflicts manually, mark them as resolved with
"git add/rm <conflicted_files>", then run "git rebase --continue".
You can instead skip this commit: run "git rebase --skip".
To abort and get back to the state before "git rebase", run "git rebase --abort".

# conflicting states
git status

rebase in progress; onto cf6de71
You are currently rebasing branch 'bigtrouble' on 'cf6de71'.
        (fix conflicts and then run "git rebase --continue")
        (use "git rebase --skip" to skip this patch)
        (use "git rebase --abort" to check out the original branch)

...

# resolve conflicts on conflicted file
git add ${conflicted-file}
git rebase --continue
```

- rebase - pull with rebase

```bash
# git fetch is non-destructive command that simply updates 
# the references between remote repository and the local repository
git fetch origin master

# check local branch have diverged
git status
On branch master
Your branch and 'origin/master' have diverged,

#
git pull --rebase origin master
```

- stash

```bash
# stash is default command
git stash (save)

# do some work...

# apply stashed file
git stash apply

# list stash
git stash list

# drop last stash
git stash drop

# stash untracked file
git stash -u

# git stash apply & drop
git stash pop

## work with multiple stashes

# stash with message
git stash save "${messages}"

# show stash details
# e.g. git stash show stash@{1}
git stash show stash@{${index}}

# stash with index
git stash apply stash@{${index}}

# drop with index
git stash drop stash@{${index}}

# clear all stash
git stash clear

## stashing into a branch
# if you ack that you are working on a wrong branch,
# git stash can be solution to apply your changes
# on right branch

# add all things to stash
git stash -u

# 1. create branch
# 2. switched to branch
# 3. apply stash
# 4. drop stash
git stash branch newchanges
```

- tagging

```bash
## light-weight tag
# tags are just nothing more than lables

# create light-weight tag
git tag ${tag-name}

# list tag, caution: always need to specify two dashes
# otherwise, you will create new tag
git tag --list

# delete tag
git tag --delete ${tag-name}

## annotated tag: similiar with lightweight tag, except
# it has a little extra information
git tag -a ${tag-name}

git tag ${tag-name} -m "${message}"

# comparing tags
git diff ${tag-name-1} ${tag-name-2}

# tagging a specific commit
git tag -a ${tag-name} ${commit-id}

# updating tag
# change tag to another tag
git tag -a ${tag-name-on-wrong-position} -f ${right-commit-id}

# push one tag to remote
git push origin ${tag-name}

# push any tag to remote
git push origin master --tags

# delete tag on remote
git push origin :${tag-name}
```

- amend commit message

```bash
# ammend commit message
git commit --amend
```

