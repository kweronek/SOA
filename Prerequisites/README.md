## 2 Boiler-Plate
Unter Boiler-plate (Herdplatte) wir alles das verstanden, was benötigt wird um Anwendungssoftware entwickeln zu können. Da SOA zunächst Technlogie-offen ist, kann mann alles verwenden, wird aber dann nie fertig und kann auch keinen fragen, da jeder etwas anderes benutzt. Der Autor hat daher versucht, die Technologien auszuwählen, die einerseits die Grundlagen gut vermitteln können, wahrscheinlich später im Beruf gebraucht werden und die schnell und einfach zu installieren und zu benutzen sind. Außerdem müssen aller Dinge ja auch getestet werden und im Zweifel Unterstützung geleistet zu werden.  
Wie erläutert basiert das Gesamtkonzept auf der Basis von Go als Programmiersprache auf Linux. Einer der Hauptgründe ist, dass die meisten der aktuellen Cloud-/Microservices Applikationen aber auch Technologien mit Go auf Linux entwickelt sind oder werden.

### 2.1 Betriebssystem für Entwicklung und Testen
#### 2.1.1 Linux (A)
In vielen Fällen werden moderen Service-Architekturen auf Basis von Open-Source-Produkten gebaut. Aus diesem Grunde ist ein Linux-Betriebssytem die erste Wahl. Dabei ists es praktisch unerheblich, welche Distritbution eingesetzt wird. In unserer Hochschule kommt auf den Laborrechner OpenSUSE zum Einsatz, welches ich empfehle, wobei auch Ubuntu LTS eine gute Wahl ist. Der entscheidende Unterschied zwischen den Linux-Distributionen ist in der Regel bei der Installation und dem Update von Software, sowie die Bootmechnismen.  

#### 2.1.2 MacOS (C)
MacOS ist Linux sehr ähnlich. Es wird mit der MAC-Hardware mitgelifert. Die Installation der Software unterscheidet sich von den Linux-Systemen. Der Betrieb und die Funktionsweise ist aber praktisch gleich, so dass alles, was in diesem Context gemacht wird auch auf MacOS läuft.  

#### 2.1.3 Andere Betriebssysteme (C)
Prinzipiell eignet sich auch Microsoft-Windows. Windows ist in diesem Sinne aber Proprietär und ggf. kostenpflichtig. Es ist für Microservice-Architekturen und Cloud-native-Technologien die erste Wahl. Aus diesem Grunde wird Window hier nicht weiter unterstützt. Eine Variante ist der Einsatz von WSL2 (Windows-Subsystem für Linux), um ein Ubuntu auf Windows laufen zu lassen. Die Erfahrung zeigt, dass das in vielen Fällen ausreicht. Gleichzeitig tauchen aber Inkompatibilitäten auf, die unnötigen Aufwand erfordern. Auch Oracle-Virtual-Box ist möglich, um eine Linux auf einem Windows-Rechner laufen zu lassen. Hier gibt es ggf. Einschränkungen für die Virtualisierung innerhalb des Unix in der Virtual-Box.  

### 2.2 Go-Compiler/-Environment (A)
Go ist unabhängig vom Betriebssystem sehr leicht zu installiern. Sie hierzu die Originalquelle unter  
[https://go.dev/doc/install].
Überprüfen der Installation mit:  
```go version```  
Erwartete Antwort (kann leicht abweichen):    
```go version go1.21.3 darwin/amd64```  
Überprüfen der Installation mit:  
```which go```  
Erwartete Antwort (kann leicht abweichen):   
```/usr/local/go```  
Überprüfen der Installation mit:  
```echo $GOROOT```  
Erwartete Antwort (kann leicht abweichen):   
```/usr/local/go```  
Überprüfen der Installation mit:  
```echo $GOPATH```  
Erwartete Antwort (kann leicht abweichen):  
```user\me\go```  
Überprüfen der Installation mit:  
```echo $GOPATH```  
Erwartete Antwort (kann leicht abweichen):  
```go version go1.21.3 darwin/amd64```  

### 2.3. Entwicklungsumgebung
Zur Entwicklungsumgebung gehören zunächst alle Tools zu Entwickung der Software benötigt wird. Das sind im einfachsten Fall ein Editor und der Compiler. In der professionellen Softwareentwicklung zeigt sich schnell, dass das sehr bald nicht mehr ausreicht. Intelligente Editoren, die Syntax-fehler bereits beim Eintippen erkennen oder Befehle und Parameter vorschlagen können, einen Linter enthalten und auch gleich die Software testen und starten können, werden daher in sog. integrierten Entwicklungsumgebungen (IDE) zusammengefasst.
Für die Entwicklung im Team wir ein Code-Repository für das Versionsmanagement nötig.  

#### 2.3.1 Goland (B)
Der Einsatz einer IDE ist hoch sinnvoll. Die Auswahlkriterien sind, eine dedizierte Unterstützung von Go, ein professionell Tool, aber ohne Kosten. Auf Basis der Kriterien ist die Entscheidung auf Goland der Firma JetBrains gefallen. Es erfüllt alle Kriterien, da JetBrains für Studentinnen und Studenten ihre professionelle Variante kostenlos zur Verfügung stellt. Darüber hinaus besitzt die Hochschule einen Lizenz-Server, die Lizenzen für die Laborumgebung bereitstellt.  
Die Installation ist in der Regel unkompliziert und erfolgt entweder über die sog. JetBrains Toolbox, snap (Ubuntu), .tar-file (andere Linux) oder .dmg-file (MAC).  
Die Details sind beschrieben unter:  
[https://www.jetbrains.com/help/go/installation-guide.html]  

#### 2.3.2 Git (B)
#### 2.3.3 GitLab FRA-UAS (B)
tbd  
#### 2.3.4 GitLab Cloud (C)
tbd  
#### 2.3.5 GitHub (C)
tbd  
#### 2.3.6 GIT: So geht's
[https://github.com/kweronek/github-demo/tree/master]

### 2.4. Tools
tbd  
#### 2.4.1 cURL (A)
Als einfachster Client für unsere Implementierung bietet sich der Kommandozeilen-Browser cURL an.
cURL (Client URL) ist ein vielseitiges Werkzeug zur Übertragung von Daten mit und ohne Protokollspezifikation. Der Kommandozeilenbrowser cURL unterstützt viele Protokolle, darunter HTTP, HTTPS, FTP, FTPS, SCP, SFTP, LDAP, LDAPS, SMTP, POP3 und viele mehr. cURL wird häufig für Entwicklungs- und Debugging-Zwecke verwendet, um HTTP-Anfragen zu erstellen und daraufhin die Antworten zu inspizieren.

Hier sind einige der wichtigsten Funktionalitäten, die cURL bietet:

1. **Datenabfrage**: Mit cURL können Sie Daten von Servern über verschiedene Protokolle abrufen.
  
2. **Datenübertragung**: Sie können Daten zu einem Server senden, z.B. durch POST-Anfragen bei Webformularen oder durch Dateiuploads.

3. **Header-Manipulation**: Sie können benutzerdefinierte Header zu Ihren Anfragen hinzufügen oder die Antwortheader anzeigen.

4. **Authentifizierung**: cURL unterstützt verschiedene Authentifizierungsmethoden, darunter Basic, Digest, NTLM und weitere.

5. **Folgen von Weiterleitungen**: cURL kann automatisch HTTP-Weiterleitungen (z.B. 302 Found) folgen.

6. **Cookie-Handling**: Mit cURL können Sie Cookies setzen, auslesen und speichern.

7. **Zertifikats- und SSL/TLS-Verwaltung**: Sie können cURL anweisen, SSL-Zertifikate zu überprüfen, Zertifikate zu verwenden oder die Überprüfung zu umgehen.

8. **Benutzeragenten**: Sie können den User-Agent-String anpassen, um sich z.B. als bestimmter Webbrowser auszugeben.

9. **Proxy-Unterstützung**: Wenn Sie über einen Proxy-Server arbeiten, kann cURL entsprechend konfiguriert werden.

10. **Begrenzung der Übertragungsrate**: cURL ermöglicht es Ihnen, die Bandbreite für Up- oder Downloads zu begrenzen.

11. **Zeitüberschreitungen**: Sie können eine maximale Zeit für die Anfrage festlegen, nach deren Ablauf cURL abbricht.

12. **FTP-Kommandos**: Bei Verwendung von FTP können Sie zusätzliche Kommandos senden und verschiedene Methoden für den Dateitransfer auswählen.

Das ist nur ein grober Überblick über die Fähigkeiten von cURL. Um alle Optionen und Möglichkeiten zu sehen, können Sie `curl --help` oder `man curl` (sofern die man-Seiten installiert sind) in der Kommandozeile eingeben. Das Tool hat eine umfangreiche Dokumentation und eine aktive Gemeinschaft, die bei Problemen hilfreich sein kann.
#### 2.4.2 Postman (B)

### 2.5.Virtualisierungs Plattformen
tbd  
#### 2.5.1 Docker/Podman (B)
tbd  
#### 2.5.2 Multipass (B)
tbd  
#### 2.5.3 Kubernetes (B)

### 2.6.Cloud-Infrastruktur
#### 2.6.1 AWS (B)
tbd  
#### 2.6.2 GCP (C)
tbd  
#### 2.6.3 Azure (C)
tbd  

