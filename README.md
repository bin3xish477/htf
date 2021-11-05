# htf
**htf** (Host That File) is a tool to make serving up your favorite pentest tools simpler and faster. All you need to do is populate the htf configuration file (see example below) with the nickname and path to tools you want to host and the pass the nickname you gave to the file as an argument to **htf**. [Medium 
Blog](https://binexishatt.medium.com/htf-host-that-file-e75c8daae80f)

![image](https://user-images.githubusercontent.com/44281620/140176817-edf9a3c0-5106-4cb8-a15d-3573d5f45a63.png)

### Example Usage
```
$ htf -r -p 8000 linpeas chisel pspy
$ htf linpeas
```

### Help
```
$ htf -h
Usage: htf [--random] [--port PORT] [FILES [FILES ...]]

Positional arguments:
  FILES                  files to host by nickname

Options:
  --random, -r           generate random file name for hosted file
  --port PORT, -p PORT   the port to listen on [default: 7000]
  --help, -h             display this help and exit
```

### Example Configuration File

> Must save this file as `~/.htf.json`

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
