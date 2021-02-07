# 😎 Gitmoji

## Table of Contents
- [💭 Background](##💭-Background)
- [⚡ Quickstart](##⚡-Quickstart)
- [📥 Installation](##📥-Installation)
  * [Linux/Mac](###Linux/Mac)
  * [Windows](###Windows)
  * [Self Build](###Self-Build)
- [📄 Documentation](##📄-Documentation)
- [🙌 Contributing](##🙌-Contributing)

## 💭 Background
I was on the wave of emoji commit messages, but after a while, having to keep going to
a website and then think about which emoji you want to use, copy it, and then paste it
along with my commit message was just not worth it.  

So how about I just type something like `gitm fix It should be all good now` where fix
declares the type of emoji I want to use and "It should be all good now" is my commit
message? Here, I'll help you scroll down: [click me](##⚡-Quickstart).

## ⚡ Quickstart
```bash
$ git add -A
$ gitm new Implemented new security features
Successfully commited: ✨ Implemented new security features

# OR

$ gitm new -a Implemented new security features
Successfully commited: ✨ Implemented new security features
```

## 📥 Installation
### Linux/Mac
```bash
$ wget xxxxxxxxxxx.tar.gz && tar -xzf xxxxxxxxxxx.tar.gz
$ sudo mv gitmoji/xxxx/gitm /usr/local/bin
$ gitm -v
gitm version 0.1.0
```

### Windows
You're using WSL right? 😅  
If so refer to the [Linux/Mac](###Linux/Mac) Section.  
If not, [click this](https://docs.microsoft.com/en-us/windows/wsl/install-win10).

### Self Build
Requirements: [Go](https://golang.org/)!
```bash
$ git clone https://github.com/VijayStroup/gitmoji && cd gitmoji
$ go test tests/commit_test.go
ok      command-line-arguments  3.766s
$ go build -o gitm .
$ chmod +x gitm && sudo mv gitm /usr/local/bin
$ cd .. && rm gitmoji
$ gitm -v
gitm version 0.1.0
```

## 📄 Documentation
### First Steps
Make sure where ever you choose to install the gitm binary, that location is in
your `$PATH`.  

### Modularity
When adding a new command, or overwriting one of the three default commands, all
you have to do is add a new environment variable with a prefix of `GITM_`.  
Here is an example:
```bash
export GITM_BUILD="build:🏗️"
```
The suffix of the variable name is insignificant, however for every command you
wish to make, it would be wise to make them all different so that they do not
overwrite each other.  
Notice the format of the variable: `command:emoji`. Gitmoji commands will always
be in lowercase, even if in this case you were to set the variable above to
`BuiLD:🏗️`. The correct way to use this command would be `gitm build A fresh build`.  
For the `emoji`, any text can be represented here (it does not actually have to
be an emoji), and will be prepended to your git commit message.  

For lasting effect, make sure to add Gitmoji commands to your `.bashrc`.
Here is an example `.bashrc`:
```bash
export GITM_BUILD="build:🏗️"
export GITM_NEW="new:🌟"
export GITM_DELETE="del:❌"
```
Now the following new commands would be available: `build, del` and `new` would
be overwritten from the default emoji of ✨ to 🌟.


## 🙌 Contributing
I'm not that quite sure on how much more you can do with just git commits but if
you know something, don't hesitate to open a new Pull Request!  

Only one thing though (well a few):  
PRs are broken down into 3 sections: new, update, fix.  
- ✨ New: Any new features or completely new logic into a pre-existing function.
- ☝️ Update: Edits to a pre-existing function or logic.
- 🔧 Fix: Fixing bug found in program (overtakes Update if Update is a fix to a bug).

When submitted PRs, please have your commit messages in the form of `EMOJI MESSAGE`
with `EMOJI` being the corresponding emoji from the list above and `MESSAGE` being
the commit message. We must set the example!
