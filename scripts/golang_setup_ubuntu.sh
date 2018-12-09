#!/usr/bin/env bash


read -p "Enter $USER sudo user's password: " sudo_password


cd /tmp

echo $sudo_password | sudo apt-get update
echo $sudo_password | sudo apt-get -y upgrade

curl -O https://storage.googleapis.com/golang/go1.11.2.linux-amd64.tar.gz

tar -xvf go1.11.2.linux-amd64.tar.gz


echo $sudo_password | sudo mv go /usr/local

export PATH=$PATH:/usr/local/go/bin

export PATH=$PATH:$HOME/go/bin