# command and control (C2)
* know how in movies computers are *"hacked"*? attackers somehow get access to a machine remotely and *"brinng the house down"* so to speak...
* well, it may be possible to do that in real life. considr the following:

    ```mermaid
        ---
        title: command and control (c2)
        ---
        flowchart RL
            A[admin client] e1@==> |admin API|B[C2 gRPC API server]
            B e1@==> |embed API|C[victim client]
            C e1@==> B
            B e1@==> A
            e1@{ animate: true, animation: fast }
    ```

* attacker uses  `admin client` to execute commands in `victim client` through `C2 gRPC API server`: the server is the *command and control* server
* the `victim client` returns the results of said commands to `admin client` through the same channel: `victim client` &rarr; `C2 server` &rarr; `admin client`
* why TF do we need two clients and one server? why can `admin client` not connect directly to `victim client`?
    * glad you asked...
    * the answer is scale
    * assume we want to control/access the nodes in an organisation (size of org is irrelevant): we need a way to talk to all of them individually and simultaneously at the same time
        * wait. how TF do you communicate individually and simultaneously at the same time?
            - maintain an open channel with all victim clients
            - communicate with any one(s) you want: machine A, for example. machines B and E perhaps. all the victim clients. world's your oyster...
    * there could be one client in the org; there could be a million: the number of clients is in the range \[0, &infin;\)
    * the method of communication must be IP agnostic yet capable of using IP addresses to speak to individual clients when that is required
    * also, we do not want the `admin client` to run as a server 24/7
        - some processes may take forever; there are other things the attacker could do in the meantime: climb tall trees, play tic-tac-toe, feed the cows/chickens
        - these things may require computing power that the attacker may not afford to acquire and maintain
        - there is also the part where the attacker does not want to get geo-located, therefore, "spreads his forces" (thank you Sun Tzu)
* why two APIs?
    * because
    * what!?
    * yes, because
    * 