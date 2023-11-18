# dockerdemo1

Das Verzeichnis enthält 3 essentielle Dateien:

1. dockerdemo.go
2. Dockerfile
3. make

## dockerdemo.go
Die Datei dockerdemo.go ist das Sourcefile für eine Bake. Diese sendet im 10-Sekundenabständen einen Timestamp nach stdout. Die Parameter für die Anzahl und den Zeittakt können natürlich angepasst werden. Die Funktionsweise ist in den Kommentarzeilen im Programm beschrieben. Zum Ausprobieren kann dieses, wie gewohnt, mit ``go run dockerdemo.go`` ausgeführt werden.  

## Prerequisites
Für den folgenden Teil müssen neben einer funktionierenden Go-Installation drei Bedingungen bzgl Docker bzw. Dockerdesktop erfüllt sein:
1. Docker bzw. Dockerdesktop muss installiert sein
1. Der Docker-Daeman muss als Service laufen und erreichbar sein
1. Sie benötigen eine Registrierung bei Dockerhub
Wie Sie das erreichen können, finden Sie unter "Boilerplate".

## Erstellen eines Dockerimages
Zum Erstellen eines Containers-Images wird ein Programm für den Containerbau benötigt. Für den Bau eines Container-Images gibt es den sog. OCI-Standard (Open-Container-Initiativ). Docker, Podman und rkt erfüllen diesen Standard. Während Docker und Podman weitegehend identisch sind (Unterschiede gibt es in der Laufzeitarchitektur) benötigt rkt teilweise Anpassungen. Durch die starke Verbreitung werden wir aktuell Docker einsetzen, auch wenn Docker teilweise durch Podman ersetzt wird.  
Hinweis: Bei korrekter Sprechweise ist ein Container-Image ein executable, welches in einer Container-Runtime-Umgebung ausgeführt werden kann. Ein Container-Image, welches ausgeführt wird heißt dann Container. Leider wird häufig der Begriff "Container" und "Image" synonym verwendet. Manchmal wird auch nur von Dockercontainern oder Dockerimages gesprochen.
Die Datei Dockerfile ist ein Dockerfile. Ein Dockerfile ist eine Skriptdatei, welche dem Dockerprogramm vorgibt, wie und in welcher Reihenfolge des Image gebaut wird. welche das Programm dockerdemo1.go compiliert und zu einem ausführbaren Programm bindet. Hierbei handelt es sich um ein one-stage-Dockerfile, d.h. alles wird in einem Basisimage erledigt. Als Basis-Image wird golang:alpine verwendet, d.h. ein Alpine Linux mit go-Installation. Compile, Build und Run erfolgen in diesem Container. Er ist daher relativ groß.



Zusätzlich gibt es ein Makefile für die Kommandozeile.
