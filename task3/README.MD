# TASK 3 - DEVELOPMENT

Write a python (or language of your choice) CLI (Command Line Interface) which takes a name
and age as an argument and starts a docker container for X seconds, where X is the age
argument value.

    E.g- $ python ager.py -n <NAME> -a <AGE>

The docker container should log the following output, which the python program should capture:

    Hi <NAME>, age: <AGE>
    The container should be stopped after X seconds. An example of this would be:

$ python ager.py -n DevOps -a 3
    docker run ...  
    <after 3 seconds>  
    Hi DevOps, age: 3  
  
  
# Solution notes

Using the same Vagrant VM from task 2 to run the go program.

    cd task_points/task3  
    go run ager.go -n DevOps -a 3  
    go test  
