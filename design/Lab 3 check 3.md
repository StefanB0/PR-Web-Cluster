Lab 3 goals

1. **Re-election**

	1.1. **Find if leader is dead**
		
		1.1.1 Ping leader every 2 seconds (gorutine)
		1.1.2 Record last time leader responded (mutex)
		1.1.3 Declare leader dead if wait-time more than 10 sec.

	1.2. **Choose new leader**
		
		1.2.1 Ping all nodes
		1.2.2 Update cluster list
		1.2.3 Choose node with least id
	
	1.3. **Promote to leader**
		
		1.3.1 New leader declares himself leader
		1.3.2 New leader tells the plebenians he is new leader
		1.3.3 Business as usual

2. **Load balancing**
	
	2.1 Implement read request proxy (?), redirect (?)

3. **Periodic Synchronisation**
	
	3.1 Every time data is updated add the time to a hash table.
	3.2 Send a ping every 30 seconds to check other people's timestamps
	3.3 If someone has latest timestamp, send it to those who have lesser.
	
4. **Implement aditional protocols**

	4.1 FTP file storage support for backups and synching. (will need a simple back-up at master where he dumps everything)
	4.2 Websockets for syncs. 
	4.3 TCP/IP for load balancing. 
	4.4 HTTP for client communication.
	4.1 UDP for pinging.
	Broadcast udp??