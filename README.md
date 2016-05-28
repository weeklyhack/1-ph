# Ph: Add some chemistry to your git push.
A simple way to reduce keystrokes and make git push faster.

![Introduction](https://cdn.rawgit.com/weeklyhack/1-ph/master/img/intro.svg)

## Features
- Scans your current repository, intelligently finding git remotes and branches.
- Doesn't require any runtime installed on the remote system
- Dead simple to use

## Download
[https://github.com/1egoman/1-ph/tree/master/compiled](Grab a precompiled binary here)
or, compile it yourself by cloning the repository and running `go build ph` in
the root. Either way, copy the binary to somewhere in your $PATH, then run `ph
help`.

## Usage
Run `ph help`.

## A few real world examples
![A demo](http://weeklyhack.github.io/assets/images/posts/ph.gif)

## Why did I make this?
Every day, I push code with git at least 50 times. Usually, I run something
like `git push origin master`, a full 22 characters. There had to be a more
efficient way to push code. Am I lazy? Maybe. But, efficiency matters.
I tried some fancy shell aliases, and while they
were ok they never really worked quite right for my needs. I figured this
would be a perfect opportunity to write a solution.

