# Cosmic Horizon Testnets

You will find all testnet implementations of the Cosmic Horizon Network in this repository.

# Installing CoHo

## Quick Links
Genesis: **coming soon**

Git tag: coho v0.1

Block explorer: **coming soon**

Seeds: TBA

## Hardware Requirements
Here are the minimal hardware configs required for running a validator/sentry node
 - 8GB RAM
 - 4vCPUs
 - 200GB Disk space

## Software Requirements

- Ubuntu 18.04 or higher
- [Go v1.17.1](https://golang.org/doc/install)
- [Starport](https://docs.starport.network/guide/install.html)

# Install CoHo, Generate Wallet and Start your Node

You have two options for installing the cohod binary. First, our team will be providing releases of cohod in our github repository, please check the [releases page](https://github.com/cosmic-horizon/coho/releases) for the latest version of cohod.  Secondly, you can follow the steps below to compile coho yourself.

## Install Go version 1.17.1

```
sudo apt update  
sudo apt install build-essential jq wget git -y

wget https://dl.google.com/go/go1.17.1.linux-amd64.tar.gz
tar -xvf go1.17.1.linux-amd64.tar.gz
sudo mv go /usr/local
```

Now add go to your bashrc -
```
echo "" >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.bashrc

# use this new bashrc configuration
source ~/.bashrc
```

## Install Starport

This command invokes curl to download the install script and pipes the output to bash to perform the installation. The starport binary is installed in /usr/local/bin.

To learn more or customize the installation process, see [Starport installer docs](https://github.com/allinbits/starport-installer) on GitHub.

```
curl https://get.starport.network/starport! | bash
```

**(Optional)**
Starport installation requires write permission to the /usr/local/bin/ directory. If the installation fails because you do not have write permission to /usr/local/bin/, run the following command:

```
curl https://get.starport.network/starport | bash
```
Then run this command to move the starport executable to /usr/local/bin/:
```
sudo mv starport /usr/local/bin/
```

## Build CoHo

These next steps will install the cohod binary which is used to run your chain.
```
# Clone the Repo
git clone https://github.com/cosmic-horizon/coho.git

# Install CoHo
cd ~/coho
starport chain build
```