# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.

Vagrant.configure(2) do |config|
  
  # 64 bit Vagrant Box
  config.vm.box = "ubuntu/trusty64"

  ## Guest Config
  config.vm.hostname = "cos461"
  config.vm.network "forwarded_port", guest: 8888, host: 8888
  config.ssh.forward_x11 = true
 
  ## Provisioning
  config.vm.provision "shell", inline: <<-SHELL
    #sudo apt-get update
    #sudo apt-get install -y python-dev
    #sudo apt-get install -y python-pip
    #sudo pip install --upgrade ipython[all]
    #sudo apt-get install -y gccgo-go
    sudo apt-get install -y apache2-utils
  SHELL

  
  
  ## Notebook
  config.vm.provision "shell", run: "always", inline: <<-SHELL
    ipython notebook --notebook-dir=/vagrant/notebook --no-browser --ip=0.0.0.0 &
  SHELL
  
  ## CPU & RAM
  config.vm.provider "virtualbox" do |vb|
    vb.customize ["modifyvm", :id, "--cpuexecutioncap", "100"]
    vb.memory = 2048
    vb.cpus = 1
  end

end
