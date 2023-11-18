# dockerdemo1

Das Verzeichnis enthält 3 essentielle Dateien:

1. dockerdemo.go
2. Dockerfile
3. make

Die Datei dockerdemo.go ist das Sourcefile für eine Bake. Diese sendet im 10-Sekundenabständen einen Timestamp nach stdout. Die Parameter für die Anzahl und den Zeittakt können natürlich angepasst werden. Die Funktionsweise ist in den Kommentarzeilen im Programm beschrieben. Zum Ausprobieren kann dieses, wie gewohnt, mit ``go rund dockerdemo.go`` ausgeführt werden.  
Die Datei Dockerfile ist ein Dockerfile, welche das Programm dockerdemo1.go compiliert und zu einem ausführbaren Programm bindet. Hierbei handetl es sich um ein one-stage-Dockerfile, d.h. alles wird in einem Basisimage erledigt. Als Basis-Image wird golang:alpine verwendet, d.h. ein Alpine Linux mit go-Installation. Compile, Build und Run erfolgen in diesem Container. Er ist daher relative groß.



Zusätzlich gibt es ein Makefile für die Kommandozeile.
