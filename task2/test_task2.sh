#! /bin/bash

printf "\n\n"
echo "*****************************"
echo "*  Installed docker version *"
echo "*****************************"
sudo apt list --installed | grep docker-ce

printf "\n\n"
echo "****************************"
echo "*  Show docker daemon.json *"
echo "****************************"
sudo cat /etc/docker/daemon.json

printf "\n\n"
echo "*********************************"
echo "*  List custom data-root folder *"
echo "*********************************"
sudo ls -la /data/docker