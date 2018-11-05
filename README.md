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

## `--engine-storage-driver overlay2`?

Yeah, sorry... docker-machine's default driver for debian right now is aufs,
which makes sense on debian versions prior to stretch. Bytemark's stretch image
does not have aufs installed, so --engine-storage-driver must be set when
creating your docker-machine to `overlay2` - this is the replacement for aufs
and is generally recommended.

[authapi](https://docs.bytemark.co.uk/article/about-the-cloud-server-api/#authentication)
[bytemark-client](https://github.com/BytemarkHosting/bytemark-client)
