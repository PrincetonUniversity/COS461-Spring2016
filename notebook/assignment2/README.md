### Assignment 2

#### Due 6:00 pm. Thursday, March 31, 2016

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

The new Vagrantfile will install several packages on your VM. You can take a look at the Vagrantfile on the folder COS461-Spring2016 on your host machine to see the packages that are required for the assignment. Make sure that all the packages were successfully installed.

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
```

If the packages were not installed correctly, you can reprovision your VM or install the packages manually by executing the commands that are in the Vagrantfile. For example, to install mininet, you can execute the following command on your VM:
```bash
$ sudo apt-get install mininet
```


This assignment has two parts: active measurement (Bufferbloat), and passive network data analysis with Spark. The instructions for the assignment are in the form of iPython notebooks. To access the iPython notebooks, you can use a web browser on your host machine and type the URL http://localhost:8888/

#### Part 1 - Bufferbloat (10 points)

The instructions for this part of the assignment are in the form of an iPython notebook in the *Bufferbloat* folder. The notebook has step-by-step instructions on how to create a simple network in a simulation tool called mininet, and see the effect of a phenomena called Bufferbloat.

#### Part 2 - Passive Network Data Analysis with Spark (10 points)
The instructions for this part of the assignment are in the form of an iPython notebook in the *Passive Measurement* folder. 
 
#### Submission and Grading
You should submit your completed proxy and web server by the date posted on the course website to [CS Dropbox](https://dropbox.cs.princeton.edu/COS461_S2016/Assignment2). You will need to submit only the two iPython notebooks. Each part of your assignment is worth ten points.

