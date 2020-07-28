.PHONY: help

install-go:
	cd /tmp;\
	wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz;\
	sudo tar -xvf go1.13.3.linux-amd64.tar.gz;\
	sudo mv go /usr/local;\
	echo "Add below lines in profile";\
	echo "export GOROOT=/usr/local/go";\
	echo "export PATH=\$GOPATH/bin:\$GOROOT/bin:\$PATH";\
	echo "source ~/.profile"

install-dc:
	sudo curl -L "https://github.com/docker/compose/releases/download/1.26.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose;

install-k6:
	sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 379CE192D401AB61;\
    echo "deb https://dl.bintray.com/loadimpact/deb stable main" | sudo tee -a /etc/apt/sources.list;\
    sudo apt-get update;\
    sudo apt-get install k6;
