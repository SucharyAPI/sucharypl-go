package main

import (
	"errors"
	"fmt"
	"log"

	sucharygo "github.com/SucharyAPI/sucharypl-go"
)

func main() {
	fmt.Println("=== Pobieranie losowego kawału ===")
	zart, err := sucharygo.DajKawal()
	if err != nil {
		log.Fatalf("Błąd: %v", err)
	}
	fmt.Println(zart)

	fmt.Println("\n=== Pobieranie kawału z kategorii: wulgarne ===")
	zartProg, err := sucharygo.DajKawal(sucharygo.Wulgarne)
	if err != nil {
		// Przykład obsługi konkretnego błędu
		var httpErr *sucharygo.HTTPError
		if errors.As(err, &httpErr) {
			fmt.Printf("Błąd HTTP (%d): %s\n", httpErr.Kod, httpErr.Tresc)
		} else {
			log.Fatalf("Inny błąd: %v", err)
		}
	} else {
		fmt.Println(zartProg)
	}

	fmt.Println("\n=== Pobieranie wielu kawałów (3) ===")
	zarty, err := sucharygo.WieleKawalow(3)
	if err != nil {
		log.Fatalf("Błąd: %v", err)
	}
	fmt.Println(zarty)
	for _, z := range zarty {
		fmt.Println(z)
	}

	fmt.Println("\n=== Pobieranie masakry ===")
	mas, err := sucharygo.Masakra()
	if err != nil {
		log.Fatalf("Błąd: %v", err)
	}
	fmt.Println(mas)

	fmt.Println("\n=== Pobieranie wielu masakr (3) ===")
	masakry, err := sucharygo.WieleMasakr(3)
	if err != nil {
		log.Fatalf("Błąd: %v", err)
	}
	fmt.Println(masakry)
	for _, z := range masakry {
		fmt.Println(z)
	}
}
