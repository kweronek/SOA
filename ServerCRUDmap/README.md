# CRUD-API in Go

## Einleitung
Das vorliegende CRUD-API ist zum Lernen gedacht. Deshalb wurde z.B. auf den Einsatz eines Multiplexers bewusst verzichtet. 
Es dient dazu, die http-MEchanismen auszuprobieren und zu verstehen. Die Daten werden aktuell in einer Map gespeicher und
sind daher nicht persistent. Der Anschluss einer MySQL-Datenbank ist vorbereiten aber nicht vollständig implementiert.
Auch hiefür wir man in der Praxis ein GORM-Framework einsetzen.

## Starten und Stoppen des CRUD-APIs
Das API kann direkt mit `go run main.go` gestartet werden.  
Das API kann in der Konsole mit Ctr-C gestoppt werden.  
Beim Neustart kann es vorkommen, dass der Port 8080 weiterhin blockiert ist.  
In diesem Fall müssen Sie eingreifen und das alte Programm entfernen.
Hinweise hierzu finden Sie unter:  
[How to properly close a port?](https://dev.to/sylwiavargas/how-to-properly-close-a-port-2p36)
Danach kann das API wie oben neu gestartet werden.

## Lesen der Daten
Beim Starten des Server werden einige wenige Datensätze vorerzeugt!  
Der Webserver lauscht dann auf Port 8080.  
Zum Testen kann direkt im Browser eingegeben werden:  
```http://localhost::8080/home```  
```http://localhost::8080/about```  
```http://localhost::8080/resources```  
```http://localhost::8080/resource/1```

## Hinzufügen von Daten
Während beim Lesen http-GET vom Browser verwendet werden kann, müssen beim Hinzufügen von Daten http-PUT Requests gesendet werden.
Aus diesem Grunde verwenden Sie ab hier den Postman!  
Holen Sie sich eine Datensatz mittels 
```http://localhost::8080/resource/1```
Kopieren Sie diesen in das "Body"-Feld des Postmans in einem http POST http://localhost::8080 request.
Passen Sie diesen Datensatz etwas an und schicken ihn ab.
Rufen Sie danach mit
```http://localhost::8080/resources```
alle Datensätze ab und suchen Sie den neu erzeugten Datensatz.

## Löschen von Datensätzen
Versuchen Sie einen Datensatz zu löschen. Stellen Sie zunächst sicher, dass es der Datensatz existiert mit:  
```http GET localhost::8080/resource/1```.  
Löschen Sie den Datensatz mit:  
```http DELETE localhost::8080/resource/1```.  
Prüfen Sie danach, ob der Datensatz noch vorhanden ist oder gelöscht ist mit:  
```http GET localhost::8080/resource/1```.  

## Ändern von Datensätzen
Versuchen Sie einen Datensatz zu ändern. Stellen Sie zunächst sicher, dass es der Datensatz existiert mit:  
```http GET localhost::8080/resource/1```.  
Holen Sie sich eine Datensatz mittels 
```http GET localhost::8080/resource/1```
Kopieren Sie diesen in das "Body"-Feld des Postmans in einem  
'http PUT http://localhost::8080/resource/1` request.  
Passen Sie diesen Datensatz etwas an und schicken ihn ab.
prüfen Sie danach mit
```http GET localhost::8080/resource/1```
ob der Datensätze noch existiert und ob dieser die ursprünglichen Daten oder die neuen Daten enthält.

## "Accept"-Parameter
In HTTP gibt es einen Header-Parameter Accept. Hiermit kann der Requester dem Provider mitteilen, welchen Antwortformat er erwarten.
In unserem Fall gibt es die Möglichkeiten `text/html` und `text/plain`.  
Versuchen Sie daher:
`http GET http://localhost:8080/home`
Einmal mit `Accept=text/html` und einmal mit `Accept=text/plain" im Header und vergleichen Sie die Ergebnisse.  
Das funktioniert natürlich nur, wenn der Server diesen Header-Key auch auswertet.  
Suchen Sie die entsprechende Stelle im Code!


