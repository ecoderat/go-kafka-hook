# go-kafka-hook

A web application is written in go lang; that logs the HTTP requests it responds; to a file, Kafka, and a database.


![](https://github.com/ecoderat/go-kafka-hook/blob/main/ui/static/img/readme-grafana-graphic.png?raw=true)


## Prerequisites
- docker
- docker-compose


## Installation
Clone the repo
```bash
git clone https://github.com/ecoderat/go-kafka-hook.git
```
Change to the directory where you cloned this repository
```
cd go-kafka-hook
```
Run docker-compose with build tag
```bash
sudo docker-compose up --build
```


## Monitoring with Grafana
#### Usage Grafana
1. Open a new tab.
1. Browse to **localhost:3000**.

#### Add a metrics data source
1. In the side bar, hover your cursor over the **Configuration** (gear) icon, and then click **Data Sources**.
1. Click **Add data source**. 
1. In the list of data sources, click **Postgresql**.
1. In the URL box, enter **http://172.17.0.1:5432**.
1. In the Database box, enter **test**.
1. In the User box, enter **postgres**.
1. In the Password box, enter **password**.
1. In the SSL Mode box, turn **disable**.
1. In the Version box, enter **12**.
1. Click **Save & Test**.

#### Build a dashboard
1. In the side bar, hover your cursor over the **Create** (plus sign) icon and then click **Dashboard**.
1. Click **Add new panel**.
2. Configire like that:
 
![](https://github.com/ecoderat/go-kafka-hook/blob/main/ui/static/img/readme-grafana-query.png?raw=true)


## Remove
Stops containers and removes containers with data
```
docker-compose down --remove-orphans --volumes
```
## Contact

Murat Güngör - [@ecoderat](https://twitter.com/ecoderat) - muratgungor.dev@gmail.com

Project Link: [https://github.com/ecoderat/go-kafka-hook](https://github.com/ecoderat/go-kafka-hook)


## License

Distributed under the Apache License. See `LICENSE` for more information.
