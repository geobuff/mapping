# geobuff/mapping
[![Go Report Card](https://goreportcard.com/badge/github.com/geobuff/mapping)](https://goreportcard.com/report/github.com/geobuff/mapping)
[![Go Reference](https://pkg.go.dev/badge/github.com/geobuff/mapping.svg)](https://pkg.go.dev/github.com/geobuff/mapping)

Package for mapping location data.

## Install
```
go get github.com/geobuff/mapping
```

## Usage

```
package main

import "github.com/geobuff/mapping"

func main() {
  mappings := mapping.Mappings["world-countries"]
}
```

## Available Mappings

| Name | Key |
| --- | --- |
| World, Countries | "world-countries" |
| World, Capitals | "world-capitals" |
| Africa, Countries | "africa-countries" |
| Asia, Countries | "asia-countries" |
| Europe, Countries | "europe-countries" |
| North America, Countries | "north-america-countries" |
| South America, Countries | "south-america-countries" |
| Oceania, Countries | "oceania-countries" |
| Argentina, Provinces | "argentina-provinces" |
| Australia, States and Territories | "australia-states-and-territories" |
| Brazil, States | "brazil-states" |
| Canada, Provinces and Territories | "canada-provinces-and-territories" |
| France, Regions | "france-regions" |
| Germany, States | "germany-states" |
| India, States and Union Territories | "india-states-and-union-territories" |
| Japan, Prefectures | "japan-prefectures" |
| Mexico, States | "mexico-states" |
| New Zealand, Regions | "new-zealand-regions" |
| Nigeria, States | "nigeria-states" |
| Turkey, Provinces | "turkey-provinces" |
| US, States | "us-states" |
| Uganda, Districts | "uganda-districts" |
| Zambia, Provinces | "zambia-provinces" |
