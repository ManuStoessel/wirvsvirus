# LadenLiebe.org

### Werde zum Stadtretter für Deine Innenstadt


<br>

<p align="center" style="margin-bottom: 4em">
  <img src="https://raw.githubusercontent.com/Solidaric-org/ladenliebe-org/master/public/images/ladenliebe-org-screenshot.jpg" alt="LadenLiebe.org Website" width="100%"/>
</p>

<br>


## Installation


#### 1. Repository klonen (stable)

```bash
  git clone -b stable https://github.com/Solidaric-org/ladenliebe-org.git
```

<br>


#### 2. Backend (Vorbereitung)

```bash
# Mariadb / Mysql Datenbank anlegen
CREATE DATABASE IF NOT EXISTS ladenliebe_db CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

# Zugriff gewähren
GRANT ALL PRIVILEGES ON ladenliebe_db.* TO 'll_user'@'localhost' IDENTIFIED BY 'll_password' WITH GRANT OPTION;

# DB aktualisieren
FLUSH PRIVILEGES;
```

<br>


#### 3. Backend (Go Service)

* Golang packages installieren und Service binary erstellen

> Golang 1.14 sollte verwendet werden

```bash
# in den pfad mit den backend sourcen
cd backend

# go packages auflösen und installieren
GO111MODULE=on go get ./...

# Binary erstellen
GO111MODULE=on GOCGO_ENABLED=0 go build -o ladenliebe_backend

# Backend Service starten
./ladenliebe_backend --port 8080 --db ll_user:ll_password@(localhost)/ladenliebe_db?charset=utf8&parseTime=True&loc=Local
```

<br>


#### 4. Frontend (Vue Application)

* Nodejs Vue Frontend erstellen

```bash
# npm module installieren
npm install

# Vue application erstellen mit DEV Server
npm run serve
```

<br>


#### 5. Zugang testen

* Browser öffnen mit Adresse: https://localhost:8081

<br>


## Documentation

**[Project homepage](https://ladenliebe.org)** - Live Prototype at [ladenliebe.org](https://ladenliebe.org)

**[#WirVsVirusHack](https://devpost.com/software/1_d_16_lokaleunternehmen-plus-tmpname)** - you will find a detailed description at [DevPost](https://devpost.com/software/1_d_16_lokaleunternehmen-plus-tmpname#updates)

<br>



## Author & Credits

Authors: [Team LadenLiebe](https://devpost.com/software/1_d_16_lokaleunternehmen-plus-tmpname)

Copyright (c) 2020 Team LadenLiebe, released under the MIT license
