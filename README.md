# htf
htf (Host That File) is a tool to make serving up your favorite pentest tools simpler and faster

![image](https://user-images.githubusercontent.com/44281620/140176817-edf9a3c0-5106-4cb8-a15d-3573d5f45a63.png)

### Usage
```
htf -r -p 8000 linpeas chisel pspy
htf linpeas
```

### Example Configuration File
```
{
  "files": [
    {
      "name": "linpeas",
      "path": "/opt/peass-ng/linPEAS/linpeas.sh"
    },
    {
      "name": "pspy",
      "path": "/opt/pspy/pspy32"
    },
    {
      "name": "chisel",
      "path": "/opt/chisel/chisel"
    }
  ]
}
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
