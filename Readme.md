### Code for "How to dockerize your gRPC service" tutorial

### Steps
    - Build docker image
    - Run docker container
    - Configure client ports to be sure requests will be send to the right destination

### How to build image
    - docker image build -t <IMAGE_NAME>:<TAG> --build-arg SSH_PRIVATE_KEY=$HOME/.ssh id_rsa .

### How to run container
    - docker container run -p <EXTERNAL_PORT>:<INTERNAL_PORT> findservice:1.0

    Before running this command think for a while what port do you want to use 
    for communication with the service