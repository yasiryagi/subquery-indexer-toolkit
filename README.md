# Swiss knife for Subquery indexers

This project collects the optimal infrastructure that makes it easier to work with the infrastructure and improve the quality of its indexer.


# Infrastructure
![dgrm_1](https://github.com/web3cdnservices/subquery-indexer-toolkit/assets/115787312/7e5e1bbc-cbb1-4b1a-acfb-0ebce31bac07)



## What is inside?
Inside the package, you can always support the indexer on the current version without the risk of damaging the configuration :)

The solution is based on
- Postgres
- Subquerynetwork/indexer-coordinator
- Redis
- Subquerynetwork/indexer-proxy
- Trafik
- Grafana
- Node exporter
- Wireguard

## Requirements

 - Any Linux dist
 - docker-compose (https://github.com/docker/compose)
 - docker (https://www.docker.com/)
 - Hard drive NVME FROM 3TB
 - 64GB memory
 - Fast and stable ethernet connection

## New setup
```
cd /opt
git clone https://github.com/web3cdnservices/subquery-indexer-toolkit.git
cd subquery-indexer-toolkit
bash ./tools/generate_initial_configuration
```
Script will generate keys and request for domain name for indexer and email.
![image](https://github.com/web3cdnservices/subquery-indexer-toolkit/assets/115787312/02422cd4-626a-4c53-b923-a19bc0203aae)

Next step preparing system and deploy all requrement software.
```
bash ./tools/start_indexer_with_wireguard 
```
![image](https://github.com/web3cdnservices/subquery-indexer-toolkit/assets/115787312/d0a372c3-9d5a-45c3-8fc2-17f83df748aa)

**After this command all services succesfully initialized. Next step - Connect indexer wallet to coordinator and generate operator wallet.** 

## Upgrade
```
cd /opt/subquery-indexer-toolkit
git pull
bash ./tools/start_indexer_with_wireguard 
```
**All your configs will be safe. You not need manually edit configurations anymore.**


## Connecting to wireguard network
**You shuld install wireguard client on your home pc**
- Macos (https://apps.apple.com/us/app/wireguard/id1451685025?mt=12)
- Linux (see repos. yum install wireguard, emerge -av wireguard, apt install wireguard)
- Windows (https://www.wireguard.com/install/)
  
Archive **SubqueryIndexerHandbook.tar.gz** containes wireguard folder with 2 configuration files.
You need only one. Secondary for secondary device or for team.

Fell free copy paste configuratin and connect with it. Nothing nedd to change.

## Internal addresses
**By default Wireguard mask 10.253.1.0/24, but you can set you own in .env file**
All addresses and services accessible via private network!

|  Service |  Address | Annotation |
| :------------ | :------------ | :------------ |
|  Postgres |  10.253.1.2 | |
|  Redis | 10.253.1.4  | |
| Indexer Coordinator  |  [http://10.253.1.3:8000](http://10.253.1.3:8000 "http://10.253.1.3:8000") | |
| Indexer Proxy  |  [http://10.253.1.5:8375](http://10.253.1.5:8375 "http://10.253.1.5:8375") | |
| Traefik  |  10.253.1.6 | ports 8082(metrics), 80(public), 443(public with Letsencrypt) |
| Prometheus  |  [http://10.253.1.8:9090](http://10.253.1.8:9090 "http://10.253.1.8:9090") | |
| Node Exporter  |  [http://10.253.1.9:9100](http://10.253.1.9:9100 "http://10.253.1.9:9100") | |
| Grafana  |  [http://10.253.1.10:3000](http://10.253.1.10:3000 "http://10.253.1.10:3000") | |

Wireguard by default use 816 port. But you can set it manually in .env file. This file NEVER overwrites, Script preventing override.


# Inbox monitoring
- Postgres statistic
- Internal metrics from Subquery
- Traefik metrics (for load counting and error counter)
- Node exporter

  Some screenshots from just created instance.

Node exporter.
![image](https://github.com/web3cdnservices/subquery-indexer-toolkit/assets/115787312/48db3d6a-387a-4711-8b8a-be0dfae1559d)

Postgres.
![image](https://github.com/web3cdnservices/subquery-indexer-toolkit/assets/115787312/0badd3e8-ca1e-447d-86e6-572c7694d1b2)

Traefik.
![image](https://github.com/web3cdnservices/subquery-indexer-toolkit/assets/115787312/48aae68f-7f4f-4314-aeb7-c77cca6bb816)

Basic indexer stat.
![image](https://github.com/web3cdnservices/subquery-indexer-toolkit/assets/115787312/94730cae-0c2c-4cae-8fc4-621c9808aed0)


  #### Guys, project from indexer for indexers. Fill free for Issues.
  ```
  Very soon postgres migration for low space servers! Be patient.
```
