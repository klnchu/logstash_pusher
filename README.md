# Logstash Pusher
Prometheus Pusher srcape from the metrics available in Logstash since version 5.0,  and push prometheus PushGateway.

## Usage

```bash
go get -u github.com/klnchu/logstash_pusher
cd /github.com/klnchu/logstash_pusher
make
./logstash_pusher -exporter.bind_address :1234 -logstash.endpoint http://localhost:1235
```

### Flags
Flag | Description | Default
-----|-------------|---------
-exporter.bind_address | Exporter bind address | :9198
-logstash.endpoint | Metrics endpoint address of logstash | http://localhost:9600
-intervel.scrape | Intervel Scrape, when less 0, stop the processing | 10


### Environment

PUSH_GATEWAYE_EDNPOINT  = 'http://pushgateway.simple.com'

## Implemented metrics
* Node metrics
