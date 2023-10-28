## Docker
Dieser Bereich dient dazu, zu zeigen, wie Go-Programme in Images eingearbeitet werden können um anschließend auf einer Containerplattform laufen gelassen zu werden. Wegen der Einfachheit kommt hier Docker zum Einsatz.
Die zwei Beispiele sind
1. dockerdemo
   dockerdemo verucht es so einfach wie möglich. Das bring gewisse Nachteile mit sich. Diese Nachteile werden in dockerdemo2 behoben. Hier kommt ein Doublestage-Image-Build zum Einsatz. Darüber hinause wird beim Bauen und beim Starten der Images noch Parameter mitgegeben die die Orgestrierung vereinfachen.
2. dockertimesrv
   Bei dockertimesrv handelt es sich um ein Beispiel aus dem Internet. Hier wird ein Zeitserver als Container auf Docker bereitgestellt. Er ermöchtlicht mittels http-Anfragen die Zeit als Antwort zu bekommen.
