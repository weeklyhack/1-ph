# PH: Add some chemistry to your git push.

Every day, I push code with git at least 50 times. Usually, I run something
like `git push origin master`, a full 22 characters. There had to be a more
efficient way to push code. Am I lazy? Maybe. But, efficiency matters.
I tried some fancy shell aliases, and while they
were ok they never really worked quite right for my needs. I figured this
would be a perfect opportunity to write a solution.

PH is an app that makes git push easier. Remember `git push origin master`?
Instead, do `ph om`. A full 18 characters shorter. Need to do something a little
more complex, like `git pull origin branch -v`? Simple, just do `ph
lobv`. Enough with my convoluted examples, here's it in action:

![http://weeklyhack.github.io/assets/images/posts/ph.gif](A demo)

### Download
[https://github.com/1egoman/1-ph/tree/master/compiled](Grab a precompiled binary here)
or, comile it yourself by cloning the repository and running `go build ph` in
the root.

## Usage
Run `ph help` for some assistance.
