[![Build Status](https://travis-ci.org/BytemarkHosting/docker-machine-driver-bytemark.svg?branch=master)](https://travis-ci.org/BytemarkHosting/docker-machine-driver-bytemark)

# docker-machine-driver-bytemark

One-command provisioning for bytemark docker-machines!

```console
$ BYTEMARK_AUTH_TOKEN="abcdefg..." docker-machine create -d bytemark --engine-storage-driver overlay2 example-docker-machine
Running pre-create checks...
Creating machine...
(example-docker-machine) Generating SSH Key
(example-docker-machine) Creating host...
Waiting for machine to be running, this may take a few minutes...
Detecting operating system of created instance...
Waiting for SSH to be available...
Detecting the provisioner...
Provisioning with debian...
Copying certs to the local machine directory...
Copying certs to the remote machine...
Setting Docker configuration on the remote daemon...
Checking connection to Docker...
Docker is up and running!
To see how to connect your Docker Client to the Docker Engine running on this virtual machine, run: docker-machine env example-docker-machine

$ eval "$(docker-machine env example-docker-machine)"
$ docker-machine ls
NAME                     ACTIVE   DRIVER     STATE     URL                                SWARM   DOCKER        ERRORS
example-docker-machine   *        bytemark   Running   tcp://[2001:41c9:1:426::88]:2376           v18.06.1-ce

$ docker run busybox echo hello world
Unable to find image 'busybox:latest' locally
latest: Pulling from library/busybox
90e01955edcd: Pull complete
Digest: sha256:2a03a6059f21e150ae84b0973863609494aad70f0a80eaeb64bddd8d92465812
Status: Downloaded newer image for busybox:latest
hello world

$ docker ps -a
CONTAINER ID        IMAGE               COMMAND              CREATED             STATUS                     PORTS               NAMES
95e4570d7f6e        busybox             "echo hello world"   4 seconds ago       Exited (0) 3 seconds ago                       happy_williams
```

## Installation

Download the correct build for your OS from our
[github releases][github releases] and install the binary into your PATH
(/usr/local/bin tends to work well)

For Homebrew users on MacOS
```console
$ brew install BytemarkHosting/tools/docker-machine-driver-bytemark
```

## Usage

The driver provides the following flags to docker-machine create:
```console
$ docker-machine create -d bytemark
   --bytemark-cores "1"         Number of CPU cores [$BYTEMARK_CORES]
   --bytemark-disk-grade "sata" Disk storage grade (sata / archive)
   --bytemark-disk-size "25"    Disk size in GiB
   --bytemark-memory "1024"     Memory in MiB [$BYTEMARK_MEMORY]
   --bytemark-zone "york"       Zone [$BYTEMARK_ZONE]
```
Right now `docker-machine-driver-bytemark` requires that the
`--engine-storage-driver` flag is set to `overlay2` as per the example at the
beginning of this README. This is due to docker-machine's `debian` provisioner
defaulting `--engine-storage-driver` to `aufs` - which is not supported on
Debians with a 4.0+ Linux kernel (Debian Stretch and beyond)

The machine name is used to provide the group and account which the machine will
belong to. Defaults are chosen in the same way as bytemark-client; a request is
sent to our billing system to determine your default account (which is the first
account associated with your login user) and the group named 'default' is the
default group within that account. Here are some examples:

```console
# creates a server called 'docker' in the 'default' group in your default account
$ docker-machine create -d bytemark --engine-storage-driver overlay2 docker

# creates a server called 'master' in the 'swarm' group in your default account
$ docker-machine create -d bytemark --engine-storage-driver overlay2 master.swarm

# creates a server called 'master' in the 'swarm' group in the '[honeyiscool][honeyiscool]' account
$ docker-machine create -d bytemark --engine-storage-driver overlay2 master.swarm.honeyiscool
```


## Authentication

Authentication is done using a token obtained from Bytemark's auth server as
described in [our API documentation][authapi].

To provide the token to docker-machine you must either set the
`BYTEMARK_AUTH_TOKEN` environment variable, or store the token in
`~/.bytemark/token`. This is also the location that
[bytemark-client][bytemark-client] uses to store the token, so you can, for
example, set up a swarm cluster without having to go get a token by yourself.

```console
$ bytemark create group swarm
$ docker-machine create -d bytemark --engine-storage-driver overlay2 manager
$ docker-machine create -d bytemark --engine-storage-driver overlay2 worker
$ docker-machine ssh manager "docker swarm init"
$ docker-machine ssh worker "docker swarm join --token [worker-token]"
```

[authapi](https://docs.bytemark.co.uk/article/about-the-cloud-server-api/#authentication)
[bytemark-client](https://github.com/BytemarkHosting/bytemark-client)
[honeyiscool](https://www.youtube.com/watch?v=NxNCWogS-SI)
