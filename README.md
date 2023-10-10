# dnsprev

#### A Script written in GO to resolve your DNS

### Install
````bash
go install -v github.com/eikehacker1/dnsprev@latest
````
### Move to binary (If your PC is Linux)
````bash
cp ~/go/bin/dnsprev /usr/bin
````
#### Usage:
````bash
cat ips.txt | dnsprev
cat domains.txt | dnsprev
````
Use it together with an http resolver
