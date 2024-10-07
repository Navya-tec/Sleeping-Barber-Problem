## Sleeping Barber Problem
In computer science, the sleeping barber problem is a classic inter-process communication and synchronization problem that illustrates the complexities that arise when there are multiple operating system processes.
Imagine a quirky barbershop with one barber, one chair for cutting hair, and a waiting area that has a variable number of chairs (let’s call it n chairs). Here’s how things work:
### **Description**

This project simulates the classic barber shop problem with the following scenarios:

### **Scenarios**

* **Barber's Nap Time:**  
  If there are no customers, the barber snoozes in his chair.

* **Waking the Barber:**  
  If a customer walks in while the barber is asleep, they have to wake him up for a haircut.

* **Busy Barber:**  
  If a customer arrives when the barber is busy and all waiting chairs are full, the customer has to leave. If there’s an empty chair, they’ll take a seat.

* **Finishing Up:**  
  Once the barber finishes a haircut, he checks the waiting area. If he sees no customers waiting, he’ll go back to sleep.

**Race Conditions:**  
Imagine a scenario where a customer arrives while the barber is cutting hair. They might see the barber working and head to the waiting area for a seat. But what if the barber finishes just as the customer is walking back? The barber checks the waiting room, sees it empty, and dozes off again! This is a classic race condition, where actions happen out of sync, leading to potential confusion.

**Simultaneous Arrivals:**  
Now, let’s say two customers arrive at the same time, and there’s only one empty chair. Both rush for it, but only the one who gets there first can sit down. This introduces a competitive scenario where timing matters, adding complexity to the situation.

### Why It Matters
This barber shop scenario is a fun way to illustrate common problems in concurrent programming—specifically, issues like race conditions and resource allocation. Just like in coding, where multiple processes may try to access the same resource, managing access and timing is crucial to avoid chaos!

 **Output**
 ### ![](https://github.com/user-attachments/assets/c965999c-133d-42be-8183-3df4e843e98b)

