# Tasks Points

For Task 2 and 3 could be used Vagrant to try the solution.

Software Required:

- Vagrant
- Virtualbox or Hyper-V

Execute:

    vagrant up (for creating the VM - Tested in Windows (Hyper-V) and Mac (Virtualbox))
    vagrant ssh - Access the VM
    vagrant destroy - Delete the VM 

Note: The VM have a copy of the repo in the path /home/vagrant/task-points         

# Task 1

See folder.

# Task 2

As part of the Vagrant provisioning script, the ansible playbook file docker.yml is executed to complete the task 2.

After the VM is up, you can execute to check task 2 completion using:

    ./tasks_points/task2/check_task2.sh

# Task 3

Using the same Vagrant VM from task 2 to run the go program.

    cd task_points/task3
    go run ager.go .......
    go test 

