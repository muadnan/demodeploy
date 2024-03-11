#!/bin/bash


server_user=demodeploy   ##Write server name Here
server_ip=34.93.58.8     ##Write Server IP Here

##Server Password Provided through terminal Input

docker_image=adnanshahm/demodeploy    #Write Image name here

sudo apt update -y                    #Updating

sudo apt upgrade -y                   #Upgrading Packages


#Assuming that the server Initailly does not have docker installed

sudo apt install docker.io -y       ##Installing docker


ssh $server_user@$server_ip <<  EOF            ##Extablishing conection to remote server and pulling and running image

    sudo docker pull $docker_image            
    sudo docker run -d -p 80:8501 $docker_image

EOF

#Running container on port 80 in detach mode...
demo

