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



### 2.3. IDE
#### 2.3.1 Goland (B)
#### 2.3.2 Git (B)
#### 2.3.3 GitLab FRA-UAS (B)
#### 2.3.4 GitLab Cloud (C)
#### 2.3.5 GitHub (C)

### 2.4. Tools
#### 2.4.1 cURL (A)
#### 2.4.2 Postman (B)

### 2.5.Virtualisierungs Plattformen
#### 2.5.1 Docker/Podman (B)
#### 2.5.2 Multipass (B)
#### 2.5.3 Kubernetes (B)

### 2.6.Cloud-Infrastruktur
#### 2.6.1 AWS (B)
#### 2.6.2 GCP (C)
#### 2.6.3 Azure (C)

