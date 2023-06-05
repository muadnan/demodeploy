This is demo deployement of docker

Here First we make a simple streamlit app and also write docker file
to create our own image.

Then I push the docker image to docker hub  you can access my image through:
docker pull adnanshah/demo

Then I create a virtual machine on Azue to deploy the contianer.
After installing docker on it I pulled the image from docker hub and run it.

I mapped the port specify in image 8501 to port 80.

You can access the site by:

104.211.200.206

