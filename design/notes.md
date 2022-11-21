Checkpoint 1

Design
- Static list of all server addresses
- Discard cluster logic. Only an array of addresses.
- Leader sends sync to every server except itself. Two arrays, one with keys, the other with data.

Steps

- Send hello get request to both servers. Configure docker, simplify.
- Logic to convert map to 2 arrays and back.
- Postman requests to leader to post and get data from database.

Checkpoint 2

- Each time data is received, it sends copies to half+1 nodes which are randomly selected.
    - To get, delete or modify data, an UDP ping is sent to check who has the data. Then a TCP request is sent.
- Fault tolerance, needs to UDP ping periodically and see if any servers are down
    - http endpoint to kill a server (by id)
    - minions periodically ping the leader, if he is dead, they ping the whole network. The server with lowest id becomes the leader.
- Log every network message
---
    ToDo
- log of communication
- udp periodic ping to check on the network

Checkpoint 3
- data integrity synchronisation
    - add hashmap to database
    - add timestamp to data
    - minions ping leader from time to time to check integrity
- load balancing
    - leader has a list who has what data
    - he randomly selects other servers to handle requests