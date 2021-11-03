windows:
	GOOS=windows GOARCH=amd64 go build -o ./htf.exe htf.go
windows-release:
	GOOS=windows GOARCH=amd64 go build -o ./htf.exe htf.go
	tar -czvf htf-windows-amd64.tar.gz ./htf.exe ./README.md
	sha256sum htf-windows-amd64.tar.gz > htf-windows-amd64.tar.gz.sha256sum
	GOOS=windows GOARCH=386 go build -o ./htf.exe htf.go
	tar -czvf htf-windows-386.tar.gz ./htf.exe ./README.md
	sha256sum htf-windows-386.tar.gz > htf-windows-386.tar.gz.sha256sum

darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./htf htf.go
darwin-release:
	GOOS=darwin GOARCH=amd64 go build -o ./htf htf.go
	tar -czvf htf-darwin-amd64.tar.gz ./htf ./README.md
	sha256sum htf-darwin-amd64.tar.gz > htf-darwin-amd64.tar.gz.sha256sum

linux:
	GOOS=linux GOARCH=amd64 go build -o ./htf htf.go
linux-release:
	GOOS=linux GOARCH=amd64 go build -o ./htf htf.go
	tar -czvf htf-linux-amd64.tar.gz ./htf ./README.md
	sha256sum htf-linux-amd64.tar.gz > htf-linux-amd64.tar.gz.sha256sum
	GOOS=linux GOARCH=386 go build -o ./htf htf.go
	tar -czvf htf-linux-386.tar.gz ./htf ./README.md
	sha256sum htf-linux-386.tar.gz > htf-linux-386.tar.gz.sha256sum

releases: linux-release windows-release darwin-release

clean:
	@if [ -f htf ]; then rm htf; fi
	@if [ -f htf.exe ]; then rm htf.exe; fi

