Crane is a simple system that helps you work with docker containers.

Crane is a set of command line and a set of github repos.

Git Repos contain content of a container image.

## 
How do I use Crane ##


Following will initalize a container for you with the content of the 'node_base' repo. Containing your favorite runtime node
```
#!bash

crane init --src github.com/platform9/node_base
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







