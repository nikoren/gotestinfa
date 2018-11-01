VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.provider "virtualbox" do |v|
    v.memory = 1048
    v.cpus = 2
    config.ssh.forward_agent = true
  end

  config.vm.box = "bento/centos-7.2"
  config.vm.network :private_network, ip: "10.0.99.4"
  
config.vm.provision :ansible do |ansible|
    ansible.playbook = "vagrant-provision.yml"
  end
end

