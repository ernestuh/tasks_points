# TASK 1 - DESIGN

Your popular web application is deployed in a Kubernetes cluster. For HA (High Availability), you've decided to deploy your web application to another Kubernetes cluster in the same on-prem co-located environment. You would now like to set up both clusters in an active-active state.

Design this architecture, and make considerations for:

    - Load Balancing - how would you distribute requests to your web application

    A: Round Robin

    - SSL support - where would you terminate SSL

    A: Load balancing Offloading

    - Database - how would you ensure data consistency from both active instances

    A: Both intances point to an external DB Cluster, sharing the Database(s).


See the PDF or PNG file in this folder for ARCHITECTURE DIAGRAM DESIGN.



