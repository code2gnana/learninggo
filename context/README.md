# Understanding Context in go

- Contexts are mainly used as a bucket to hold the key/value pair to be transported between each layer of an application.
- Contexts can also be used to define the time sensitive deadlines like handling transaction timeout deadlines in a payment system. 
- You can define how long a process can take to execute and once it exceeds, you can control the deadline.
    
    For Example:
    > If there is a transaction of 100 records to process, after processing 50 records if there are issues identied and if we need to terminate the remaining 50 records. The context would help us to handle this situation. 

- It should be noted that, you can certainly not use context for passing information between layers of your application. you shoule absolutely use it for only things that propogate through.

- You shouldn't be using context as buckets for all your information. It should be used as a supplementry object to store object to store request IDs or trace ids



