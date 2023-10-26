# 2 Boiler-Plate
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
Git ist ein verteiltes Versionskontrollsystem, das es Einzelpersonen und Teams ermöglicht, den Verlauf und die Entwicklung von Code (oder anderen Dateitypen) über die Zeit zu verwalten. Hier sind einige der wichtigsten Hauptfunktionalitäten von Git:

1. **Verteilte Versionierung**: Jeder, der eine Kopie eines Git-Repositories klont, erhält die gesamte Historie des Projekts. Dies ermöglicht eine dezentrale Arbeit und Kollaboration.

2. **Commits**: Ein Commit speichert eine Momentaufnahme der Änderungen. Jeder Commit hat eine eindeutige ID (SHA-1-Hash), die ihn identifiziert.

3. **Branching**: Git ermöglicht die Erstellung von Zweigen (Branches), sodass verschiedene Features oder Experimente isoliert voneinander entwickelt werden können.

4. **Merging**: Zweige können wieder zusammengeführt werden, wodurch die Änderungen aus einem Branch in einen anderen integriert werden.

5. **Stash**: Mit dem Befehl `git stash` können Sie temporäre Änderungen speichern und später wiederherstellen, ohne einen Commit zu erstellen.

6. **Remote Repositories**: Git unterstützt die Zusammenarbeit mit Remote-Repositories (z. B. auf Plattformen wie GitHub oder GitLab), wodurch Code gemeinsam mit anderen entwickelt und geteilt werden kann.

7. **Pull und Push**: Änderungen können von einem Remote-Repository abgerufen (`pull`) oder zu einem Remote-Repository geschickt (`push`) werden.

8. **Cloning**: Ein Remote-Repository kann geklont werden, um eine lokale Kopie auf dem eigenen Computer zu erstellen.

9. **Conflict Resolution**: Bei widersprüchlichen Änderungen bietet Git Werkzeuge an, um diese Konflikte manuell aufzulösen.

10. **Tags**: Git ermöglicht das Markieren von spezifischen Punkten in der Historie als "Tags", oft zur Kennzeichnung von Release-Versionen.

11. **Log**: Mit dem Befehl `git log` können Sie den Verlauf von Commits anzeigen, um die Entwicklungsgeschichte eines Projekts zu verfolgen.

12. **Checkout**: Der `git checkout` Befehl ermöglicht das Wechseln zwischen verschiedenen Commits, Tags oder Branches.

13. **Rebase**: Mit Rebase können Sie die Basis eines Branches ändern und/oder die Historie linearer gestalten.

14. **Bisect**: Ein Tool zur Hilfe bei der Fehlersuche, mit dem man systematisch nach dem Commit suchen kann, der einen Fehler eingeführt hat.

Diese Kurzbeschreibung skizziert nur die Hauptfunktionalitäten von Git. Es gibt viele weitere Befehle und Konzepte, die das Arbeiten mit Git effizienter und leistungsfähiger machen. Git hat eine steile Lernkurve, bietet aber enormen Nutzen, sobald man sich damit vertraut gemacht hat.
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
Postman ist ein beliebtes Werkzeug für Entwickler, um APIs zu entwickeln, zu testen, zu dokumentieren und zu überwachen. Es bietet eine benutzerfreundliche grafische Oberfläche, die die Arbeit mit APIs deutlich vereinfacht, insbesondere im Vergleich zu reinen Kommandozeilenwerkzeugen. Hier sind einige der Hauptfunktionalitäten von Postman:

1. **API-Anfragen erstellen**: Mit Postman können Benutzer HTTP-Anfragen erstellen und senden, um die Antworten von APIs zu überprüfen. Es werden alle gängigen HTTP-Methoden wie GET, POST, PUT, DELETE usw. unterstützt.

2. **Geschichtsfunktion**: Jede Anfrage, die Sie mit Postman senden, wird in der History gespeichert, sodass Sie frühere Anfragen leicht erneut aufrufen können.

3. **Umgebungen und Variablen**: Postman ermöglicht es, verschiedene "Umgebungen" zu definieren, sodass Sie leicht zwischen verschiedenen Konfigurationen (z.B. Entwicklung, Test, Produktion) wechseln können. Sie können auch Variablen verwenden, um wiederkehrende Werte in Ihren Anfragen zu speichern und wiederzuverwenden.

4. **Authentifizierung**: Postman unterstützt eine Vielzahl von Authentifizierungsmethoden, darunter Basic Auth, Bearer Token, OAuth 1.0 und 2.0 und viele andere.

5. **Tests**: Sie können Tests schreiben und diese direkt in Postman ausführen, um die Korrektheit der Antworten Ihrer API zu überprüfen. Dies ist besonders nützlich für Regressionstests.

6. **Testautomatisierung**: Mit der Collection Runner-Funktion können Sie eine Sammlung von Anfragen automatisch nacheinander ausführen, um Bulk-Tests oder Langlebigkeitstests durchzuführen.

7. **Dokumentation**: Sie können eine API-Dokumentation direkt aus Ihren Postman-Sammlungen generieren, die dann mit anderen geteilt werden kann.

8. **Mock-Server**: Mit Postman können Sie Mock-Server erstellen, um die API-Entwicklung und das Testing zu simulieren, bevor die eigentliche API implementiert wird.

9. **API-Überwachung**: Mit Postman Monitoring können Sie Ihre APIs in regelmäßigen Abständen überprüfen, um sicherzustellen, dass sie ordnungsgemäß funktionieren.

10. **Kollaboration**: In den Team-Versionen von Postman können Sie Sammlungen, Umgebungen und andere Postman-Ressourcen mit Ihren Teamkollegen teilen.

11. **Integrationen**: Postman bietet Integrationen mit anderen beliebten Tools und Plattformen, um den Workflow zu verbessern.

12. **Import/Export**: Sie können API-Definitionen im OpenAPI, RAML oder WADL Format importieren und Postman-Sammlungen und Umgebungen exportieren.

Dies ist nur ein kurzer Überblick über die Fähigkeiten von Postman. Das Tool hat im Laufe der Zeit viele zusätzliche Funktionen erhalten und wird weiterhin regelmäßig aktualisiert und erweitert. Es bietet auch eine umfangreiche Dokumentation und eine aktive Community, die bei Fragen und Problemen hilfreich sein kann.
### 2.5.Virtualisierungs Plattformen
tbd  
#### 2.5.1 Docker (B)
Docker ist ein Open-Source-Tool zur Automatisierung des Bereitstellens von Anwendungen in leichtgewichtigen, portablen Containern. Es ermöglicht Entwicklern, Anwendungen und ihre Abhängigkeiten in einem einheitlichen Paket zu kapseln, das überall – sowohl auf dem lokalen Entwicklungssystem als auch in der Cloud – ausgeführt werden kann. Hier sind einige der wichtigsten Hauptfunktionalitäten von Docker:

1. **Container**: Docker ermöglicht das Packen und Isolieren von Anwendungen mit ihren Abhängigkeiten in Containern.

2. **Images**: Ein Docker-Image ist eine leichtgewichtige, eigenständige, ausführbare Softwarepaketeinheit, die alles enthält, was eine Software zum Laufen benötigt: Code, Laufzeit, Systemtools, Bibliotheken und Einstellungen.

3. **Netzwerk und Vernetzung**: Docker bietet verschiedene Netzwerkmodi, mit denen Benutzer die Kommunikation zwischen Containern und der Außenwelt steuern können.

4. **Volumen und Datenpersistenz**: Volumen sind der bevorzugte Mechanismus zum Speichern von Daten, die von Docker-Containern erzeugt oder verwendet werden.

5. **Docker CLI & API**: Docker bietet eine Befehlszeilenschnittstelle (CLI) und eine API zur Verwaltung von Containern, Images und mehr.

6. **Isolation und Sicherheit**: Docker-Container laufen isoliert voneinander, wodurch sie zusätzliche Sicherheitsebenen bieten.

7. **Skalierbarkeit**: Docker-Container können leicht horizontal skaliert werden, sowohl auf einem einzelnen Host als auch über mehrere Hosts hinweg.

8. **Integrationsmöglichkeiten**: Docker lässt sich in viele CI/CD-Tools und -Plattformen integrieren und bietet eine Vielzahl von Plugins und Erweiterungen.

9. **Portabilität**: Anwendungen, die in Docker-Containern verpackt sind, können konsistent über mehrere Umgebungen hinweg bereitgestellt und ausgeführt werden.

Docker hat die Art und Weise revolutioniert, wie Entwickler Software erstellen, testen und verteilen, und bietet eine konsistente Umgebung von der Entwicklung bis zur Produktion. Es ist zu einem wesentlichen Tool für DevOps und Container-Orchestrierung geworden und spielt oft eine Schlüsselrolle in modernen Cloud-nativen Architekturen.

#### 2.5.2 Podman (C)

Podman und Docker sind beides Containerisierungstools, die es Benutzern ermöglichen, Container zu erstellen, auszuführen und zu verwalten. Sie bieten ähnliche Funktionen und sind in vielen Aspekten kompatibel, doch es gibt einige wichtige Unterschiede zwischen den beiden. Hier sind die Hauptunterscheidungsmerkmale von Podman im Vergleich zu Docker:

1. **Daemonless Architektur**: Einer der auffälligsten Unterschiede ist, dass Podman ohne einen langlaufenden Daemon funktioniert. Jeder `podman`-Befehl ist ein eigenständiger Prozess. Docker hingegen verwendet einen Client-Server-Ansatz, bei dem der Docker-Client mit dem Docker-Daemon kommuniziert.

2. **Rootless Containers**: Podman kann Container als normaler Benutzer ausführen, ohne Root-Rechte zu benötigen. Dies wird als "rootless" bezeichnet und verbessert die Sicherheit, da kein Daemon mit erhöhten Privilegien erforderlich ist.

3. **Kompatibilität mit Docker-CLI**: Podman bietet eine Docker-CLI-kompatible Schnittstelle, was bedeutet, dass die meisten Befehle, die mit `docker` funktionieren, auch mit `podman` funktionieren sollten, ohne Änderungen.

4. **SystemD Integration**: Podman bietet bessere Integration mit `systemd`, was die Verwaltung von Containern und Pods als Systemdienste ermöglicht.

5. **Pods**: Neben Containern unterstützt Podman auch "Pods", eine Gruppe von einem oder mehreren Containern, die Netzwerkressourcen teilen. Dieses Konzept wurde von Kubernetes inspiriert.

6. **Keine zentrale Registry**: Im Gegensatz zu Docker, das eine zentrale Registry wie Docker Hub verwendet, legt Podman nicht fest, von wo Images geholt werden. Stattdessen verwendet es Container-Image-Spezifikationen.

7. **OCI-kompatibel**: Sowohl Docker als auch Podman unterstützen den Open Container Initiative (OCI) Standard für Container-Images und Laufzeit, was die Portabilität von Containern zwischen verschiedenen Tools gewährleistet.

8. **Kompatibilität mit anderen Tools**: Durch die Ähnlichkeit in der Befehlszeilenschnittstelle kann Podman oft als Drop-in-Ersatz für Docker in Skripten oder CI/CD-Pipelines verwendet werden.

9. **Sicherheit**: Da Podman keinen Daemon verwendet und rootless-Container unterstützt, gibt es weniger Angriffspunkte im Vergleich zu einem Docker-Setup.

Zusammengefasst: Podman und Docker ähneln sich in ihrer Kernfunktionalität stark, aber ihre Architekturen und ihr Umgang mit Sicherheit unterscheiden sich. Podman bietet Vorteile in Bezug auf Sicherheit und Flexibilität, insbesondere in Umgebungen, in denen das Fehlen eines Daemons und die Möglichkeit, rootless zu laufen, wünschenswert sind.

#### 2.5.3 Multipass (B)
Multipass ist ein Open-Source-Tool, das von Canonical, dem Unternehmen hinter Ubuntu, entwickelt wurde. Es ermöglicht Nutzern, schnell und einfach virtuelle Ubuntu-Instanzen (und andere unterstützte Betriebssysteme) auf ihrem Desktop oder in der Cloud zu starten und zu verwalten. Hier sind einige der Hauptfunktionalitäten von Multipass:

1. **Schnelle Instanzerstellung**: Mit einem einfachen Befehl können Benutzer eine neue virtuelle Maschine basierend auf den neuesten verfügbaren Ubuntu-Images (oder anderen unterstützten Images) erstellen.

2. **Leichte Verwaltung**: Multipass bietet Befehle zum Starten, Stoppen, Löschen und Auflisten von Instanzen.

3. **Shell-Zugriff**: Benutzer können direkt auf eine Shell einer laufenden Instanz zugreifen, was die Interaktion und das Debugging erleichtert.

4. **Befehlsausführung**: Sie können spezifische Befehle direkt in einer virtuellen Maschine ausführen, ohne sich zuerst in diese einzuloggen.

5. **Automatische Aktualisierungen**: Multipass aktualisiert automatisch die VM-Images, sodass Benutzer immer Zugriff auf die neuesten Versionen haben.

6. **Cloud-Integration**: Obwohl Multipass in erster Linie für den lokalen Einsatz entwickelt wurde, unterstützt es auch die Erstellung von Instanzen in unterstützten Cloud-Umgebungen.

7. **Resource Management**: Sie können Speicher, CPU und Disk-Spezifikationen für Ihre VMs festlegen.

8. **Snapshots**: Multipass unterstützt das Erstellen von Snapshots von VMs, sodass Benutzer zu einem früheren Zustand zurückkehren können.

9. **Netzwerkeinstellungen**: Es ermöglicht Benutzern, Netzwerkeinstellungen für VMs anzupassen, einschließlich Portweiterleitungen.

10. **Mounten von Verzeichnissen**: Benutzer können lokale Verzeichnisse in ihre Multipass-Instanzen einhängen, wodurch die Dateiübertragung und die Arbeit mit Dateien innerhalb von VMs vereinfacht wird.

11. **Plattformübergreifend**: Multipass ist für verschiedene Betriebssysteme verfügbar, einschließlich Windows, macOS und Linux, und passt seine Hypervisor-Interaktionen entsprechend an.

Multipass zielt darauf ab, Entwicklern und Systemadministratoren eine einfache und konsistente Umgebung zur Verfügung zu stellen, um Ubuntu (und andere unterstützte Betriebssysteme) in einer isolierten, virtuellen Umgebung schnell zu testen, zu entwickeln oder einfach nur zu verwenden. Es bietet viele der Vorteile von containerisierten Lösungen, bietet jedoch eine vollständige VM-Umgebung.

#### 2.5.4 Kubernetes (B)
Kubernetes, oft einfach als "K8s" abgekürzt, ist ein Open-Source-Container-Orchestrierungssystem zur Automatisierung der Bereitstellung, Skalierung und Verwaltung von containerisierten Anwendungen. Hier sind einige der wichtigsten Hauptfunktionalitäten von Kubernetes:

1. **Pods**: Die kleinste und einfachste Einheit in Kubernetes. Ein Pod kann einen oder mehrere Container beinhalten, die gemeinsam verwaltet werden.

2. **Services**: Ein Kubernetes-Service ist ein abstrahiertes Interface zu Pods, das eine Netzwerkschnittstelle und eine stabile IP-Adresse für den Zugriff bietet.

3. **ReplicaSets und Deployments**: Diese sorgen dafür, dass eine bestimmte Anzahl von Pod-Kopien (Replikaten) laufen. Deployments bieten zudem fortschrittliche Update-Mechanismen für Pods.

4. **ConfigMaps und Secrets**: Erlauben es, Konfigurationsdaten und geheime Daten von der Anwendungs-Image-Konfiguration zu trennen und so die Wiederverwendbarkeit von Container-Images zu erhöhen.

5. **Autoscaling**: Kubernetes kann automatisch Pods hinzufügen oder entfernen, basierend auf CPU-Nutzung oder anderen ausgewählten Metriken.

6. **Persistent Storage**: Mit Persistent Volumes (PVs) und Persistent Volume Claims (PVCs) bietet Kubernetes eine Schnittstelle für Speicherressourcen.

7. **Ingress**: Ein Objekt, das den Zugriff auf die Services von außerhalb des Kubernetes-Clusters, oft von öffentlichen IP-Adressen, regelt.

8. **Network Policies**: Bestimmen, wie Pods innerhalb eines Clusters kommunizieren dürfen.

9. **Node Management**: Selbst wenn Knoten (Nodes) ausfallen, kann Kubernetes die darauf laufenden Aufgaben woanders neu starten.

10. **Cluster Federation**: Unterstützt das Management von mehreren Kubernetes-Clustern, indem es diese verknüpft.

11. **Batch Execution**: Neben Diensten kann Kubernetes auch Batch-Jobs verwalten, die einmalig oder periodisch ausgeführt werden.

12. **Health Checks**: Kubernetes überprüft regelmäßig den Zustand von Anwendungen und Services und kann bei Bedarf Neustarts durchführen.

13. **Service Discovery**: Pods erhalten automatisch DNS-Namen und IP-Adressen, sodass sie sich problemlos gegenseitig finden können.

14. **Rollout und Rollback**: Kubernetes ermöglicht es, Änderungen an Anwendungen schrittweise auszurollen und bei Problemen auf eine vorherige Version zurückzukehren.

Kubernetes bietet eine umfangreiche und modulare Plattform, die es Entwicklern und Administratoren ermöglicht, containerisierte Anwendungen in verschiedensten Umgebungen, von On-Premises-Infrastrukturen bis hin zu Cloud-Anbietern, zu orchestrieren. Es hat eine aktive Gemeinschaft und wird ständig weiterentwickelt und erweitert.

2.5.6 kubectl (B)
2.5.7 wget (A)
2.5.8 protoc (A)

### 2.6.Cloud-Infrastruktur
#### 2.6.1 AWS (B)
AWS (Amazon Web Services) ist die Cloud-Computing-Plattform von Amazon, die eine breite Palette von IT-Diensten bietet. AWS ermöglicht es Unternehmen, Rechenkapazität, Speicher und andere Ressourcen über das Internet zu mieten, anstatt physische Server zu kaufen und in eigenen Rechenzentren zu betreiben. Hier ist eine kurze Erklärung zu AWS Cloud Services:

**AWS Cloud Services** sind eine Sammlung von Remote-Computing-Diensten (auch als Webdienste bezeichnet) von Amazon. Diese Dienste bieten eine Suite von Cloud-basierten Produkten, die Rechenleistung, Speicheroptionen, Netzwerking und Datenbanken, Entwickler-Tools, Maschinenlernen, Identitäts- und Zugriffsmanagement, Sicherheits-Tools und viele andere Funktionen bieten, um Unternehmen bei der Skalierung und Wachstum zu unterstützen. 

Einige der bekanntesten AWS-Dienste umfassen:

- **EC2 (Elastic Compute Cloud)**: Ein Dienst, der virtuelle Maschinen in der Cloud bereitstellt.
  
- **S3 (Simple Storage Service)**: Ein skalierbarer Speicherdienst.
  
- **RDS (Relational Database Service)**: Ein Dienst zum Betrieb und Skalieren von relationalen Datenbanken in der Cloud.

- **Lambda**: Ein serverloser Computing-Dienst.

- **IAM (Identity and Access Management)**: Ein Dienst, der die Authentifizierung und Autorisierung für AWS-Ressourcen steuert.

Dank dieser Dienste können Unternehmen Ressourcen je nach Bedarf bereitstellen oder reduzieren (ein Konzept, das oft als "Elastizität" bezeichnet wird) und nur für das bezahlen, was sie tatsächlich verwenden. Das Modell von AWS hat dazu beigetragen, die Art und Weise, wie Unternehmen IT-Ressourcen betrachten und nutzen, zu revolutionieren. Es ermöglicht schnelle Innovationen, Skalierbarkeit und Flexibilität.
#### 2.6.2 GCP (C)
Die Google Cloud Platform (GCP) ist die Cloud-Computing-Plattform von Google, die eine Vielzahl von Diensten für Rechenleistung, Datenspeicherung, Datenanalyse, Maschinenlernen und andere Cloud-basierte Aufgaben bietet. Mit GCP können Unternehmen, Entwickler und andere Nutzer IT-Ressourcen von Google's hochmodernen Rechenzentren nutzen. Hier ist eine kurze Erklärung zu GCP:

**Google Cloud Platform (GCP)** bietet eine Suite von Cloud-basierten Diensten und Produkten, die Unternehmen dabei helfen, in der Cloud zu bauen, zu betreiben und zu skalieren. Die Plattform nutzt die gleiche Infrastruktur, die Google verwendet, um seine eigenen Produkte wie Google Suche, YouTube und Gmail zu betreiben.

Einige der Hauptdienste und Produkte von GCP umfassen:

- **Compute Engine**: Ein Infrastructure-as-a-Service (IaaS) Dienst, der virtuelle Maschinen in der Cloud bereitstellt.

- **App Engine**: Eine Platform-as-a-Service (PaaS) Lösung, die Entwicklern das Erstellen und Hosten von Webanwendungen in von Google verwalteten Rechenzentren ermöglicht.

- **Kubernetes Engine**: Ein verwalteter Kubernetes-Dienst zur Orchestrierung von Docker-Containern.

- **Cloud Storage**: Ein skalierbarer Objektspeicherdienst.

- **BigQuery**: Ein vollständig verwalteter Data Warehouse-Dienst für große Datenanalysen.

- **Cloud Spanner**: Eine vollständig verwaltete, skalierbare und konsistente relationale Datenbank.

- **Cloud Machine Learning Engine**: Ein Dienst, der Maschinenlernen in der Cloud ermöglicht und auf TensorFlow basiert.

- **Pub/Sub**: Ein Echtzeit-Nachrichtenübermittlungsdienst.

Die Google Cloud Platform zeichnet sich durch ihre leistungsstarken Maschinenlernalgorithmen, ihre Datenanalysefunktionen und ihre offenen APIs aus. Sie ist für Unternehmen konzipiert, die hochverfügbare und skalierbare Cloud-Computing-Lösungen suchen und dabei von der Sicherheit und Effizienz profitieren möchten, die Google's Infrastruktur bietet. 
#### 2.6.3 Azure (C)
Microsoft Azure, oft einfach als Azure bezeichnet, ist die Cloud-Computing-Plattform von Microsoft, die eine Vielzahl von Diensten und Lösungen in den Bereichen Computing, Analytik, Speicherung und Netzwerk bietet. Azure ermöglicht es Unternehmen und Entwicklern, ihre Anwendungen in Microsoft-eigenen und -verwalteten Rechenzentren zu erstellen, zu hosten und zu skalieren. Hier ist eine kurze Erklärung zu Microsoft Azure:

**Microsoft Azure Cloud Platform** ist eine stetig wachsende Sammlung von Cloud-Diensten, mit denen Unternehmen jede Art von Lösung – von einfachen Web-Apps bis zu komplexen, verteilt arbeitenden Anwendungen – konfrontieren, erstellen und verwalten können. Azure stellt die notwendige Infrastruktur bereit, damit Unternehmen Anwendungen in der Cloud betreiben und Ressourcen nach Bedarf skalieren können.

Einige der Hauptdienste und Produkte von Azure umfassen:

- **Azure Virtual Machines (VMs)**: Ein Infrastructure-as-a-Service (IaaS) Dienst, der es Benutzern ermöglicht, virtuelle Maschinen in der Azure-Cloud zu erstellen und zu betreiben.

- **Azure Blob Storage**: Ein Objektspeicherdienst für unstrukturierte Daten.

- **Azure SQL Database**: Ein verwalteter relationaler Datenbankdienst basierend auf SQL Server.

- **Azure Kubernetes Service (AKS)**: Ein verwalteter Kubernetes-Container-Orchestrierungsdienst.

- **Azure Active Directory**: Ein Identitätsdienst, der Single Sign-On und Multi-Factor-Authentication bietet.

- **Azure Cosmos DB**: Eine NoSQL-Datenbank für globale Anwendungen, die horizontale Skalierung und Multi-Model-Datenbankunterstützung bietet.

- **Azure App Service**: Ein Platform-as-a-Service (PaaS) Angebot, das es Entwicklern ermöglicht, Webanwendungen und APIs in der Azure-Cloud zu erstellen und zu hosten.

- **Azure DevOps Services**: Tools für Softwareentwicklung, Zusammenarbeit und CI/CD (Continuous Integration/Continuous Deployment).

- **Azure Functions**: Ein serverloser Computing-Dienst, der es Entwicklern ermöglicht, Code in Reaktion auf Ereignisse auszuführen, ohne eine vollständige Anwendung oder Infrastruktur zu verwalten.

Microsoft Azure bietet eine breite Palette von Diensten und Lösungen, um unterschiedliche Geschäftsanforderungen zu erfüllen. Es unterstützt mehrere Programmiersprachen, Frameworks, Betriebssysteme, Datenbanken und Geräte, wodurch Azure zu einer vielseitigen Cloud-Plattform für verschiedene Anwendungs- und Entwicklungsszenarien wird.

