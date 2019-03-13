# RabbitMQ Initial
&emsp;&emsp;Headers exchange \
A headers exchange is designed to for routing on multiple attributes that are more easily expressed as message headers than a routing key. Headers exchanges ignore the routing key attribute. Instead, the attributes used for routing are taken from the headers attribute. A message is considered matching if the value of the header equals the value specified upon binding.

It is possible to bind a queue to a headers exchange using more than one header for matching. In this case, the broker needs one more piece of information from the application developer, namely, should it consider messages with any of the headers matching, or all of them? This is what the "x-match" binding argument is for. When the "x-match" argument is set to "any", just one matching header value is sufficient. Alternatively, setting "x-match" to "all" mandates that all the values must match.

Headers exchanges can be looked upon as "direct exchanges on steroids". Because they route based on header values, they can be used as direct exchanges where the routing key does not have to be a string; it could be an integer or a hash (dictionary) for example.

&emsp;&emsp;Default Exchange

...

#### rabbitmqdmin

```bash
rabbitmqadmin declare binding --vhost=Host_NAME \
              source=Name_Exchange \
              destination_type=queue \
              destination=Name_Queue \
              routing_key="#"
              

rabbitmqadmin declare exchange --vhost=Host_NAME \
              name=Name_Exchange \
              type=topic \
              durable=true
```