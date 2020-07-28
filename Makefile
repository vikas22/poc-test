.PHONY: help

install-go:
	cd /tmp;\
	wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz;\
	sudo tar -xvf go1.13.3.linux-amd64.tar.gz
	sudo mv go /usr/local
	echo "Add below lines in profile"
	echo "export GOROOT=/usr/local/go"
	echo "export PATH=$GOPATH/bin:$GOROOT/bin:$PATH"
	echo "source ~/.profile"

