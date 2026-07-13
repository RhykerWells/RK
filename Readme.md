# Records Keeper 
Records Keeper is a self-hosted documentation and personal records platform.  
RK is designed to provide a central place for managing documents, records, knowledge, and structured information across multiple portals.  
The platform consists of a Go backend API, a web frontend, PostgreSQL storage, TipTap-based document editing, and integrations such as Discord.

## Project setup
### Standalone
#### Pre-requisites
```shell
sudo apt update && sudo apt upgrade
sudo apt install git
git clone https://github.com/RhykerWells/RK
```

#### Database
```shell
sudo apt update && sudo apt upgrade
sudo apt install postgresql
sudo -u postgres psql
```

```shell
CREATE DATABASE rk;
CREATE USER rk WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE rk TO rk;
GRANT USAGE, CREATE ON SCHEMA public TO rk;
\q
```
### Docker
#### Prerequisites
- docker
```
sudo apt update && sudo apt upgrade
sudo apt install git
git clone https://github.com/RhykerWells/RK
```

#### Database
```shell
cd docker/
mv example.db.env db.env
vi db.env
```
```shell
docker compose up 
```