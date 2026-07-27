package sucharygo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Struktura zwrotu z API
type ZwrotSerwera struct {
	Zart string `json:"zart"`
}

// Custom errory
type SucharyError struct {
	Tresc string
}

func (e *SucharyError) Error() string {
	return e.Tresc
}

type ConnectionError struct {
	Tresc string
}

func (e *ConnectionError) Error() string {
	return "Błąd połączenia: " + e.Tresc
}

type HTTPError struct {
	Kod   int
	Tresc string
}

func (e *HTTPError) Error() string {
	if e.Kod == 404 {
		return "Nie znaleziono kawału w podanej kategorii (kod 404)"
	}
	return fmt.Sprintf("Serwer zwrócił kod %d: %s", e.Kod, e.Tresc)
}

type JSONError struct {
	Tresc string
}

func (e *JSONError) Error() string {
	return "Błąd dekodowania danych: " + e.Tresc
}

// Kategorie kawałów
type Kategoria string

const (
	Domyslna      Kategoria = ""
	Programowanie Kategoria = "programowanie"
	Dwuznaczne    Kategoria = "dwuznaczne"
	Szkolne       Kategoria = "szkolne"
	Suchary       Kategoria = "suchary"
	Wulgarne      Kategoria = "wulgarne"
	CzarnyHumor   Kategoria = "czarnyhumor"
	Inne          Kategoria = "inne"
)

// Funkcja fetchująca
func DajKawal(kategoria ...Kategoria) (string, error) {
	klient := http.Client{
		Timeout: 6500 * time.Millisecond,
	}
	adres := "https://sucharyapi.onrender.com"
	if len(kategoria) > 0 && kategoria[0] != Domyslna {
		adres = fmt.Sprintf("https://sucharyapi.onrender.com/%s", kategoria[0])
	}
	odp, err := klient.Get(adres)
	if err != nil {
		return "", &ConnectionError{Tresc: err.Error()}
	}
	defer odp.Body.Close()
	if odp.StatusCode != http.StatusOK {
		return "", &HTTPError{
			Kod:   odp.StatusCode,
			Tresc: "Nie udało się pobrać suchara",
		}
	}
	var kawal ZwrotSerwera
	err = json.NewDecoder(odp.Body).Decode(&kawal)
	if err != nil {
		return "", &JSONError{Tresc: err.Error()}
	}
	return kawal.Zart, nil
}
