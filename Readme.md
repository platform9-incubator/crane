Crane is a simple system that helps you work with docker containers.

Crane is a set of command line and a set of github repos.

Git Repos can either 
a) contain content of a container image.
or
b) Reference to images and an overlay filesystem that just has your application.

## 
How do I use Crane ##

There are two ways:

### Start from an existing git repo with runtime already installed ###

Following will initalize a container for you with the content of the 'node_base' repo. Containing your favorite runtime node
```
#!bash

crane init --src github.com/platform9/node_base
```


### Start from an existing docker image ###

If you are not a fan of adding binaries to git repo, just reference an existing git repo or docker image

```
#!bash

crane init --src http://dockerhub.com/node/base
```

After the initialization you are running into the container, work as you please in the newly created container, creating and deleting files as needed.


```
#!bash

git add /opt/pf9/hello_world.js
```

Add a small manifest to specify the container 'command'


```
#!bash

echo 'cmd=node /opt/pf9/hello_world.js' > /.crane.env
git add .crane.env

```

Once you are satisfied, git commit the changes because the container is nothing but a git repo.


```
#!bash

git commit -am "My hello world application with crane"
```

Push your changes 

```
#!bash

crane push github.com/roopakparikh/helloworld
OR simply use git
git push github.com/roopakparikh/helloworld

```

You can run it yourself locally or ask someone else to run it


```
#!bash

crane run github.com/roopakparikh/helloworld

```