"# projet-groupie-tracker" 

# 🌟 Amiibo Explorer

Projet fait en Go pour afficher une liste de figurines Amiibo grâce à l’API officielle.  
On peut chercher, trier, filtrer par série, passer en dark mode et voir les détails en cliquant sur une carte.

## 🔧 Ce que fait le projet

- Récupère les données de l’API [amiiboapi.com](https://www.amiiboapi.com)
- Affiche tous les Amiibo dans une interface web
- Permet :
  - 🔍 de rechercher par nom
  - 🅰️ de trier (A→Z / Z→A)
  - 🎮 de filtrer par série (Zelda, Mario, etc.)
  - 🌙 de basculer en mode sombre
  - 🔄 d’ouvrir une fiche détaillée en popup
  - 🎴 d’avoir un effet flip card au clic
  - 🌀 de voir un petit loader pendant le chargement

## 🚀 Lancer le projet

Assure-toi d’avoir Go installé, puis dans le terminal :

```bash
go run main.go

amiiboapi/
├── controllers/
│   └── listAmiibo.controllers.go
├── routes/
│   └── main.routes.go
├── services/
│   └── amiibo.services.go
├── templates/
│   └── listAmiibo.html
├── main.go
├── go.mod
