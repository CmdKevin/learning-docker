# learning-docker

## Docker Version

![Docker Version](./run-existing-docker-image/docker_version.PNG)

## Running Image Docker docker/welome-to-docker

Install menggunakan dengan perintah di bawah
1. -d untuk running background terminal
2. --name untuk memberikan nama Image
3. 8080:80 port (bebasoport)

```sh
docker run -d -p 8080:80 --name welcome1 docker/welcome-to-docker
```

## Running Preview
![DockerImageRunningPreview](./run-existing-docker-image/running_docker_page.PNG)

## Mengecek Log Running Container

1. welcome1 nama file container
2. tail untuk cek log terakhir
3. 10 jumlah log terakhir

```sh
docker logs -f --tail 10 welcome1
```

![Docker Logs Container](./run-existing-docker-image/docker_running_log.PNG)

## Mengecek Log Stop Container
![Docker Stop Logs Contaier](./run-existing-docker-image/latest_logs_after_stop.PNG)

## Docker Stop Container Service

```sh
docker stop welcome1
```
![Docker Stop Container](./run-existing-docker-image/docker_stop.PNG)