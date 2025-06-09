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
* `embed.proto` file *viz*:

    ```proto
        syntax = "proto3"
        package c2grpcapi

        service Embed {
            rpc GetCommand (Empty) returns (Command);
            rpc SendCommand (Command) returns (Empty);
        }

        service Admin {
            rpc ExecuteCommand (Command) returns (Command);
        }

        message Command {
            String Input = 1;
            String Output = 2;
        }

        // Empty message to use in place of null
        message Empty {
        }
    ```

* WTF is all this?
    * eea...sy!
    * `GetCommand` RPC is called from  `victim client` onto `C2 gRPC server`. it is a polling command: said server polls said client as soon as client establishes a connects with server
        - victim client to C2 server: Hi. Do you have any commands I should run?
        - C2 server to victim client: No.
        - victim client to C2 server: *(a few seconds later)* How about now?
        - C2 server to victim client: Yes. Here... `sudo ...`
        - *(this goes on ad infinitum)*
    * message `Command` is what `GetCommand` returns. contains the commands that `victim client` will execute. `Input` field is where the command actually resides
    * `SendResult` RPC is used to send the result of the commands executed by `victim client` to `C2 gRPC server`