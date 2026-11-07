package sucharygo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Struktura zwrotu z API
type ZwrotSerwera struct {
	Zart string `json:"tresc"`
}

type ZwrotMasakra struct {
	Masakra string `json:"masakra"`
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

// Funkcja fetchująca kawał
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

// Funkcja fetchująca wiele kawałów
func WieleKawalow(ilosc int) ([]string, error) {
	klient := http.Client{Timeout: 6500 * time.Millisecond}

	adres := fmt.Sprintf("https://sucharyapi.onrender.com/zart/%d", ilosc)

	odp, err := klient.Get(adres)
	if err != nil {
		return nil, &ConnectionError{Tresc: err.Error()}
	}
	defer odp.Body.Close()

	if odp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			Kod:   odp.StatusCode,
			Tresc: "Nie udało się pobrać sucharów",
		}
	}

	var kawalyR []ZwrotSerwera
	err = json.NewDecoder(odp.Body).Decode(&kawalyR)
	if err != nil {
		return nil, &JSONError{Tresc: err.Error()}
	}

	var kawaly []string

	for _, i := range kawalyR {
		kawaly = append(kawaly, i.Zart)
	}

	return kawaly, nil
}

// Funkcja fetchująca masakrę
func Masakra() (string, error) {
	klient := http.Client{Timeout: 6500 * time.Millisecond}

	adres := "https://sucharyapi.onrender.com/masakra"

	odp, err := klient.Get(adres)
	if err != nil {
		return "", &ConnectionError{Tresc: err.Error()}
	}
	defer odp.Body.Close()

	if odp.StatusCode != http.StatusOK {
		return "", &HTTPError{
			Kod:   odp.StatusCode,
			Tresc: "Nie udało się pobrać masakry",
		}
	}

	var masakra ZwrotMasakra
	err = json.NewDecoder(odp.Body).Decode(&masakra)
	if err != nil {
		return "", &JSONError{Tresc: err.Error()}
	}

	return masakra.Masakra, nil
}

// Funkcja fetchująca wiele masakr
func WieleMasakr(ilosc int) ([]string, error) {
	klient := http.Client{Timeout: 6500 * time.Millisecond}

	adres := fmt.Sprintf("https://sucharyapi.onrender.com/masakra/%d", ilosc)

	odp, err := klient.Get(adres)
	if err != nil {
		return nil, &ConnectionError{Tresc: err.Error()}
	}
	defer odp.Body.Close()

	if odp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			Kod:   odp.StatusCode,
			Tresc: "Nie udało się pobrać masakr",
		}
	}

	var masakryR []ZwrotMasakra
	err = json.NewDecoder(odp.Body).Decode(&masakryR)
	if err != nil {
		return nil, &JSONError{Tresc: err.Error()}
	}

	var masakry []string

	for _, i := range masakryR {
		masakry = append(masakry, i.Masakra)
	}

	return masakry, nil
}
