# go-datadog-poc

This is just a plain go program which runs some custom metrics.

### Docker-agent

please use the respective api key when creating this dd-agent

docker service create -d --replicas 1 --name dd-agent --mount type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock,ro --mount type=bind,src=/proc/,dst=/host/proc/,ro --mount type=bind,src=/sys/fs/cgroup/,dst=/host/sys/fs/cgroup,ro -e DD_API_KEY="<DATADOG_API_KEY>"  -e DD_LOGS_ENABLED=true -e DOCKER_CONTENT_TRUST=1 -e DD_DOGSTATSD_NON_LOCAL_TRAFFIC=true -e SD_BACKEND=docker -p 8125:8125/udp datadog/agent:6

Once the DD-agent is created if you run the code you will be seeing the metrics in metrics explorer.