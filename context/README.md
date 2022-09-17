# Understanding Context in go

- Contexts are mainly used as a bucket to hold the key/value pair to be transported between each layer of the workflow
- Contexts can also be used to define the time sensitive deadlines like handling transaction timeout deadlines in a payment system. 
- You can define how long a process can take to execute and once it exceeds, you can control the deadline.
