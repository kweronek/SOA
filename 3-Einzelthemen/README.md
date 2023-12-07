# Dieses Kapitel enthält Einzelthemen
## Docker
### Unix in Containers
Interaktives Linux (hier alpine) im Dockercontainer
Hole das aktuelle alpine-Image in die lokale Docker-Registry:
```
docker pull alpine
```
Prüfe ob das alpine-Image angekommen ist:
```
docker images
```
Starte das alpine-Image als Container mit dem Namen myalpine.
```
docker run -it --name myalpine -d alpine
```
Prüfe, ob der Container myalpine läuft:
```
docker ps
```
Prüfe, welche alpine-Version im Container läuft:
```
cat /etc/os-release
```
Optional: Passwort für root setzen mit:
```
passwd
```
Verlassen der bash-Shell des Containers:
```
exit
```

Ref.: [https://blog.jetbrains.com/dotnet/2023/03/27/connecting-to-a-running-docker-container-shell/](https://blog.jetbrains.com/dotnet/2023/03/27/connecting-to-a-running-docker-container-shell/)
## Serialisierung
### XML
### JSON
### YAML
### ProtoBuf
