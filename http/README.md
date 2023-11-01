## HTTP-Server in GO

HTTP-Server in GO sind sehr einfach zu implementieren. Den einfachsten Fall gibt das Beispiel unter http0Srv.  
http1Srv ist etwas aufwändiger und senden selbst ein HTTP GET auf http://country.io.  
http1SrvTLS stellt einen HTTPS-Server (HTTP1.1) dar. Hierzu sind Zertifikate erforderliche. Diese sind im Unterordner certs.  
http2SrvTLS ist identisch mit http1SrvTLS, jedoch wird in diesem Fall HTTPS2 verwendent.

## HTTP clients

Das Programm httpGETclient ist ein Beispiel for ein go-Programm, welches ein HTTP GET Request schickt und die Antwort eines Servers entgegennimmt. httpPOSTclient hat dieselbe Funktion, benutzt aber als Verb HTTP PUT.

## Multiplexer

Das Programm mux.go gibt ein Beispiel, wie das Routing innerhalb eines Webservers implementiert werden kann.
