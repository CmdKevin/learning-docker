# Practice
## ContainerVolume V1
![Create Volume](./../run-existing-docker-image/create%20volume.PNG)
```sh
docker run --name my-postgres-kevin -e POSTGRES_USER=kevin -e POSTGRES_PASSWORD=mysecretpassword -v my-postgres-kevin:/var/lib/postgresql/data -p 5432:5432 -d postgres:latest
```
## Hasil Container V1
![Container V1](./../run-existing-docker-image/containerV1.PNG)

## Hasil Volume Postgress
![Volume Postgres](./../run-existing-docker-image/volume-postgres.PNG)

## Create Table di pgAdmin
![Create Table](./../run-existing-docker-image/create%20table.PNG)

## Stop Hapus Container Lama

## Buat Container V2
![Create Container V2](./../run-existing-docker-image/create-volumeV2.PNG)
```sh
 docker run --name my-postgres-V2-CmdKevin -e POSTGRES_USER=kevin -e POSTGRES_PASSWORD=mysecretpassword -v my-postgres-kevin:/var/lib/postgresql/data -p 5432:5432 -d postgres:latest
 ```
 
## Hasil Container V2
![Container V2](./../run-existing-docker-image/containerV2.PNG)

## Table Masih Ada
![Table masih ada](./../run-existing-docker-image/tablemasih.PNG)
