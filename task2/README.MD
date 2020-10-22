# TASK 2 - ANSIBLE

Write an Ansible playbook to:

1. Create the /etc/docker/ directory
2. Create a docker configuration file as template (/etc/docker/daemon.json)
3. Generate the configuration file using data-root=/data/docker as variable
4. Install docker 19.03.1 in a linux system.
5. Make sure docker will auto start after reboot.
  

# Solution notes

As part of the Vagrant provisioning script, the ansible playbook file docker.yml is executed to complete the task 2.

After the VM is up, you can execute to check task 2 completion using:

    vagrant ssh  
    ./tasks_points/task2/check_task2.sh  