# htf
htf (Host That File) is a tool to make serving up your favorite pentest tools simpler and faster


### Usage
```
htf -r -p 8000 linpeas chisel pspy
htf linpeas
```

### Installation (Linux Example)
> With go v1.16+ installed
```
go install -v github.com/binexisHATT/htf
```

> From releases page
```
wget htf-linux-amd64.tar.gz -O ./htf.tar.gz
tar xvzf htf.tar.gz
sudo mv htf /usr/local/bin/
```

> From source
```
git clone https://github.com/binexisHATT/htf
go build -o htf htf.go
sudo mv htf /usr/local/bin/
```
