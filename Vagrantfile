# -*- mode: ruby -*-
# vi: set ft=ruby :

GO_VERSION = "go1.15.1"

VAGRANT_BOX = "roboxes/ubuntu2004"

$go_install = <<~"SCRIPT"
echo "#### INSTALLING GO ####"
curl -OL https://golang.org/dl/go#{GO_VERSION}.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go#{GO_VERSION}.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/vagrant/.bashrc
echo 'export GOPATH=$HOME/go' >> /home/vagrant/.bashrc
SCRIPT

Vagrant.configure("2") do |config|

  config.vm.define "gobox" do |node|

    node.vm.box = VAGRANT_BOX

    node.vm.synced_folder "./go-code/", "/home/vagrant/go-code/", type: "rsync"

    node.vm.provider :libvirt do |libvirt|
      libvirt.cpus = 2
      libvirt.memory = 2048
    end

    node.vm.provision "shell", inline: $go_install

  end

end
