<img alt="Ph" src="https://cdn.rawgit.com/weeklyhack/1-ph/master/img/logo.svg" style="height: 100px;" />

## Ph is a simple way to reduce keystrokes and make git push faster.
- Scans your current repository, intelligently finding git remotes and branches.
- Doesn't require any runtime installed on the remote system.
- Dead simple to install and use.

## Usage
![Introduction](https://cdn.rawgit.com/weeklyhack/1-ph/master/img/intro.svg)

## A few real world examples
![A demo](http://weeklyhack.github.io/assets/images/posts/ph.gif)

## Install with one command
```
curl -L https://github.com/weeklyhack/1-ph/raw/master/compiled/ph-$(uname -s)-$(uname -p) > ph && chmod +x ph && sudo cp ph /usr/local/bin/ph
```
Then, run `ph help`.

## Why did I make this?
Every day, I push code with git at least 50 times. Usually, I run something
like `git push origin master`, a full 22 characters. There had to be a more
efficient way to push code. Am I lazy? Maybe. But, efficiency matters.
I tried some fancy shell aliases, and while they
were ok they never really worked quite right for my needs. I figured this
would be a perfect opportunity to write a solution.
