package main

// package main

import (
"context"
"fmt"
"github.com/kr/pretty"
"googlemaps.github.io/maps"
)

func main() {
	// der apiKey dient der Autentisierung und Authorisierung bei Google
	apiKey := "AIzaSyDTWWzLzqHPwVgnjc18l8KHw2IMSLW4myY"
	// eine Instanz des clients mit dem apiKey wird zum Ansprechen der Google-maps-API erzeugt
	client, _ := maps.NewClient(maps.WithAPIKey(apiKey))

	// ein leerer Request für die Abfrage der Distance-Matrix wird erzeugt
	req := maps.DistanceMatrixRequest{}
	// im request werden die Ausgangspunkte definiert
	// die Ausgangspunkte sind ein Array of strings
	req.Origins = []string{"Hauptbahnhof, Frankfurt am Main"}
	// im request werden die Zielpunkte definiert
	// die Zielpunkte sind ein Array of strings
	req.Destinations = []string{"Nibelungenplatz	, Frankfurt am Main"}
	// der Abfahrtszeitpunkt wird als aktueller Zeitpunkt des Requests definiert
	req.DepartureTime = "now"

	// der Request wird versendet und der Response gespeichert
	// der context.Background ist ein leerer Context, da die Parameter im context opional sind
	resp, err := client.DistanceMatrix(context.Background(), &req)
	if err != nil {fmt.Println(err)}
	// Einfache und schöne Ausgabe des Responses
	pretty.Println(resp)
}
