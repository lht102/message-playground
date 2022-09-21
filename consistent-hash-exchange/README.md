# Consistent Hash Exchange

```
./bin/setup-consistent-hash-exchange -uri=amqp://guest:guest@localhost:5672/ -exchange=event -queue-prefix=event-update -hash-header=hash-on -number-of-queues=4 
./bin/producer -uri=amqp://guest:guest@localhost:5672/ -exchange=event -hash-header=hash-on -number-of-messages=20
./bin/consumer -uri=amqp://guest:guest@localhost:5672/ -queue=event-update001 -consumer-tag=c1
./bin/consumer -uri=amqp://guest:guest@localhost:5672/ -queue=event-update002 -consumer-tag=c2
./bin/consumer -uri=amqp://guest:guest@localhost:5672/ -queue=event-update003 -consumer-tag=c3
./bin/consumer -uri=amqp://guest:guest@localhost:5672/ -queue=event-update004 -consumer-tag=c4
```

```
./bin/setup-queue -uri=amqp://guest:guest@localhost:5672/ -queue=event-update-all
./bin/producer -uri=amqp://guest:guest@localhost:5672/ -routing-key=event-update-all -number-of-messages=20
./bin/consumer -uri=amqp://guest:guest@localhost:5672/ -queue=event-update-all -consumer-tag=c1
./bin/consumer -uri=amqp://guest:guest@localhost:5672/ -queue=event-update-all -consumer-tag=c2
./bin/consumer -uri=amqp://guest:guest@localhost:5672/ -queue=event-update-all -consumer-tag=c3
./bin/consumer -uri=amqp://guest:guest@localhost:5672/ -queue=event-update-all -consumer-tag=c4
```
