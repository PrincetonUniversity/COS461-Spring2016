# COS461-Spring2016
## Assignment0: Virtual Machine Setup

###Setting up the course VM

####Prerequisite

To get started install these softwares on your host machine:

1. Install ***Vagrant***, it is a wrapper around virtualization softwares like VirtualBox, VMWare etc.: https://www.vagrantup.com/downloads.html

2. Install ***VirtualBox***, this would be your VM provider: https://www.virtualbox.org/wiki/Downloads

3. Install ***Git***, it is a distributed version control system: https://git-scm.com/downloads

4. Install X Server and SSH capable terminal
    * For Windows install [Xming](http://sourceforge.net/project/downloading.php?group_id=156984&filename=Xming-6-9-0-31-setup.exe) and [Putty](http://the.earth.li/~sgtatham/putty/latest/x86/putty.exe).
    * For MAC OS install [XQuartz](http://xquartz.macosforge.org/trac/wiki) and Terminal.app (builtin)
    * Linux comes pre-installed with X server and Gnome terminal + SSH (buitlin)   

####Basics

* Clone the course repository from Github:
```bash 
$ git clone https://github.com/PrincetonUniversity/COS461-Spring2016.git
```

* Change the directory to COS461-Spring2016:
```bash
$ cd COS461-Spring2016
```

* Now run the vagrant up command. This will read the Vagrantfile from the current directory and provision the VM accordingly:
```bash
$ vagrant up
```

If you want to tear down your vagrant session, you have multiple options to do so, each has its pros and cons. These options are as follows: 
* **vagrant suspend**: With this option you will be able to save the state of the VM and stop it. 
* **vagrant halt**: This will gracefully shutdown the guest operating system and power down the guest machine. 
* **vagrant destroy**: If you want to remove all traces of the guest machine from your system, this is the command you should use. It will power down the machine and remove all guest hard disks

Go [here](http://docs.vagrantup.com/v2/getting-started/teardown.html) for more information about vagrant teardown. 

* Now SSH into the VM:
``` bash
$ vagrant ssh
```

* Programming assignments: You will find the programming assignments in the vm under the directory: /vagrant/notebook/assignments.
``` bash
vagrant@cos461:~$ ls /vagrant/notebook/assignments
assignment0
```

* Getting started with `notebook`: On your host machine, start the web browser. Now type the url, `http://127.0.0.1:8888/tree`. You will see the list of all the assignments for the course. Click on `assignment0/` and then click on `Instructions.ipynb` to start the notebook for `assignment0`. To run this notebook, refer to the set of basic instructions [here](https://jupyter-notebook.readthedocs.org/en/latest/examples/Notebook/rstversions/Notebook%20Basics.html). 

#### Notes and Tips for running vagrant and the course VM
- Follow the instructions closely
- If you are running a 64bit OS, run a 64 bit VM
- If you are running a 32bit OS, run a 32 bit VM
- Disable Hyper-V on Windows
- Enable Virtualization support on the host (BIOS setup)
- The VM runs 192.168.0.0/24 as default network, if you use that network locally you need to change it (edit Vagrantfile: config.vm.network :private_network, ip: "192.168.0.100”)
- The host machine needs to run an X server (it’s native on Linux; OS X and Windowns require the installation of an X Server)
- Alternative methods to ssh to the VM: run "ssh -X vagrant@192.168.0.100" password is vagrant
- If you see "ssh_exchange_identification: read: Connection reset by peer" when trying to connect using "vagrant ssh" use the alternative method provided above since it has been reported as a valid workaround, this issue is under investigation
- In past, some students reported that on old/slower machines some assignments might fail because of extra delay on the computation wich translates to increaced delay on ping request/response. The original default vagrant configuration for VBox was 2Gb RAM and 1 CPU with 50% cap. We changed this to 1GB RAM and 1 CPU with no cap. We think this new configuration will be good for most, but if you want to you can always tweek this to your requirements or likings (edit the line that starts with "# CPU & RAM" on Vagrantfile)

#### Sample setup output for reference
- Following output is from MacOS

```

$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Checking if box 'ubuntu/trusty64' is up to date...
==> default: A newer version of the box 'ubuntu/trusty64' is available! You currently
==> default: have version '20150821.0.1'. The latest is version '20160120.0.1'. Run
==> default: `vagrant box update` to update.
==> default: Clearing any previously set forwarded ports...
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 8888 => 8888 (adapter 1)
    default: 22 => 2222 (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Connection timeout. Retrying...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /Users/glex/Documents/cos461
==> default: Machine already provisioned. Run `vagrant provision` or use the `--provision`
==> default: to force provisioning. Provisioners marked to run always will still run.
==> default: Running provisioner: shell...
    default: Running: inline script
==> default: stdin: is not a tty


$ vagrant ssh
Welcome to Ubuntu 14.04.3 LTS (GNU/Linux 3.13.0-62-generic x86_64)

 * Documentation:  https://help.ubuntu.com/

  System information as of Wed Jan 27 01:16:32 UTC 2016

  System load:  0.81              Processes:           80
  Usage of /:   3.5% of 39.34GB   Users logged in:     0
  Memory usage: 49%               IP address for eth0: 10.0.2.15
  Swap usage:   0%

  Graph this data and manage this system at:
    https://landscape.canonical.com/

  Get cloud support with Ubuntu Advantage Cloud Guest:
    http://www.ubuntu.com/business/services/cloud

Last login: Wed Jan 27 01:08:11 2016 from 10.0.2.2


vagrant@cos461:~$ ls /vagrant/
notebook  README.md	Vagrantfile
```
