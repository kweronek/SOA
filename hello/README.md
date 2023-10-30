# Hello World in Go

Das Programm hello.go dient dazu, zu testen, ob die Go-Installation prinzipiell functioniert hat. Es zeigt den grundätzlichen Aufbau eines Go-Programm mit package, include und der Ausgabe eines Textes. Es zeigt auch, dass Go sehr einfach mit Unicode-Zeichen umgegehen kann.  
Zunächts versuche man '''go run hello.go''' Das compiliert das Programm und führt es bei fehlerfreiheit unmittelbar aus.  
Als nächstes baue man ein executable durch ausführen von '''go build hello.go"'''.  Dies erzeugt dann ein ausführbares Programm mit dem Nahmen hello.  
Darüber hinaus kann das Programm installiert werden mit '''go install hello.go'''. Indiesem Fall wird ein ausführbares PRogramm erzeugt und unter $GOBIN gespeichert. In der Regel ist dies $GOPATH/bin.
