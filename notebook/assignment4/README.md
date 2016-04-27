### Assignment 4: Web Performance

#### Due 5:00 p.m. Tuesday, May 10, 2016 (Dean's Date)

#### Getting Started

On your host machine (laptop), go to the course directory:

```bash 
$ cd COS461-Spring2016
```

* Now, pull the latest update from Github.
```bash
$ git pull
```

* Reprovision your VM as follows: 
```bash
$ vagrant reload --provision
```

The new Vagrantfile will install a couple of packages on your VM. You can take a look at the Vagrantfile on the folder COS461-Spring2016 on your host machine to see the packages that are required for the assignment. Make sure that all the packages were successfully installed.

* Verify that the required packages were installed, as follows:
```bash
$ vagrant ssh
* Inside the VM, execute the following commands:
$ mn
* You should see:
*** Mininet must run as root.

$ java
* You should see the Usage message from the Java compiler.

* Check if the following python packages can be imported:
$ python
>>> import dateutil
>>> import test_helper
>>> import matplotlib
>>> import mininet
>>> import scapy
```

If the packages were not installed correctly, you can reprovision your VM or install the packages manually by executing the commands that are in the Vagrantfile. For example, to install Mininet, you can execute the following command on your VM:
```bash
$ sudo apt-get install mininet
```

This assignment is entirely described in the iPython notebook Assignment4-Web-Performance.ipynb. To access the iPython notebook, you can use a web browser on your host machine and type the URL http://localhost:8888/

