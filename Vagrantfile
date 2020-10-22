# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.

## Vars
PROJECT_NAME = "tasks_points"


Vagrant.configure("2") do |config|
  config.vm.box = "generic/ubuntu1804"
  # uncomment next line to force hyper-v in Windows 
  # config.vm.provider "hyperv"
  config.vm.synced_folder ".", "/vagrant", disabled: true
  config.ssh.insert_key = false
  config.vm.provider "hyperv" do |h|
    h.enable_virtualization_extensions = true
    h.linked_clone = true
  end
  
  # Script to inject to Vagrant VM
  $script = <<-SCRIPT
   echo "Project name: $PROJECT_NAME"
   add-apt-repository ppa:longsleep/golang-backports -y 
   apt update
   apt install software-properties-common
   apt-add-repository --yes --update ppa:ansible/ansible
   apt install ansible --yes
   apt install python-apt --yes
   git clone https://github.com/ernestuh/$PROJECT_NAME.git
   chown -R vagrant:vagrant $PROJECT_NAME
   ansible-playbook $PROJECT_NAME/task2/docker.yml -vvv >> docker-ansible-output.log
   chmod a+x $PROJECT_NAME/task2/check_task2.sh
   docker pull busybox
   apt install golang-go --yes
  SCRIPT
  
  config.vm.provision "shell", inline: $script, env: {"PROJECT_NAME"=> PROJECT_NAME}
end