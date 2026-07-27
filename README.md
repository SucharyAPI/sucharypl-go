# SucharyPL - oficjalna biblioteka SucharyAPI dla Go
Biblioteka Go do pobierania kawałów z ogólnopolskiej bazy SucharyAPI

[![Go Version](https://img.shields.io/badge/go-1.18%2B-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

---

## 🚀 Instalacja
### Za pomocą go get
```bash
go get github.com/SucharyAPI/sucharypl-go
```

---

## ⚡ Szybki start
### Pobranie losowego suchara
```go
package main

import (
    "fmt"
    "log"

    sucharygo "github.com/SucharyAPI/sucharypl-go"
)

func main() {
    zart, err := sucharygo.DajKawal()
    if err != nil {
        log.Fatalf("Błąd: %v", err)
    }
    fmt.Println(zart)
}
```

### Pobranie suchara z konkretnej kategorii
```go
package main

import (
    "fmt"
    "log"

    sucharygo "github.com/SucharyAPI/sucharypl-go"
)

func main() {
    zart, err := sucharygo.DajKawal(sucharygo.Suchary)
    if err != nil {
        log.Fatalf("Błąd: %v", err)
    }
    fmt.Println(zart)
}
```

---

## Kategorie
- `programowanie` - Suchary o programowaniu
- `dwuznaczne` - Suchary dwuznaczne
- `szkolne` - Suchary szkolne
- `suchary` - Suchary ogólne
- `wulgarne` - Suchary wulgarne
- `czarnyhumor` - Czarny humor
- `inne` - Inne kategorie

W kodzie dostępne jako typ `Kategoria`.

---

## 🛡️ Obsługa błędów
Biblioteka definiuje sformalizowane błędy:
- `SucharyError` - bazowy błąd
- `ConnectionError` - problemy z połączeniem
- `HTTPError` - błędy serwera (z kodem HTTP)
- `JSONError` - błąd dekodowania odpowiedzi

### Przykład obsługi błędów
```go
package main

import (
    "errors"
    "fmt"
    "log"

    sucharygo "github.com/SucharyAPI/sucharypl-go"
)

func main() {
    zart, err := sucharygo.DajKawal(sucharygo.Wulgarne)
    if err != nil {
        var httpErr *sucharygo.HTTPError
        if errors.As(err, &httpErr) {
            fmt.Printf("Błąd HTTP (%d): %s\n", httpErr.Kod, httpErr.Tresc)
            return
        }
        log.Fatalf("Inny błąd: %v", err)
    }
    fmt.Println(zart)
}
```

---

## 📡 API
Biblioteka korzysta z oficjalnej bazy SucharyAPI:
🔗 https://sucharyapi.onrender.com

---

## 📝 Licencja
MIT

---
