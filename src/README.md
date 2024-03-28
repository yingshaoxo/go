## Prepare a base golang compiler
Download a c_code 'go1.4-bootstrap-20171003.tar.gz' file, extract it, I assume it is `$HOME/CS/go2`

In `~/.bashrc`:
```bash
export GOROOT_BOOTSTRAP="$HOME/CS/go2"
```

## Clone the source code
Clone this repository, extract it, I assume it is `$HOME/CS/go`

In `~/.bashrc`:
```bash
export PATH="$PATH:$HOME/CS/go/bin"
```

Then, compile it by:
```bash
cd ~/CS/go/src
./all.bash
```

## Check the version
```bash
source ~/.bashrc
go version
```

> Great, now you have your own golang compiler.
