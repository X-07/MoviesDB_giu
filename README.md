# MoviesDB_giu
Test GIU library for GO 

Je développe en **GO** un gestionnaire de collection de films (façon **GCStar** pour les connaisseurs).
Il n'est pas question, ici, de visualiser les films, d'autres le font très bien (**VLC**...) mais juste de gérer une collection.
Ce logiciel sera composé à 90% d'une interface graphique (**UI**) et pour 10% d'une BDD (ici **SQLite**).

Pour l'**UI** après les résultats mitigés avec **Fyne** : https://github.com/fyne-io/fyne
J'ai donc exploré ce que propose **GIU** : https://github.com/AllenDang/giu

Voici le résultat obtenu ...

### Onglet 'Fiche'
![Image 1](/ScreenShots/MoviesDB-1.jpg)

- Syntaxe particulière (déroutante pour moi, mais pas insurmontable) donc assez complexe
- Code très concis
- Look! Chacun se fera son avis
- IMPOSSIBLE d'avoir un retour à la ligne automatique pour les très long texte même dans un champ multiligne!
 Un bug qui sera peut être corrigé un jour. 
- Prometteur, à surveiller.

## Le problème du champ multiligne est rédhibitoire pour moi
j'ai donc abandonné ce projet (que je laisse ici comme example) et le l'ai continué en GOTK3 : https://github.com/gotk3/gotk3
- Beaucoup, beaucoup plus verbeux

### Interface principale Gtk3 à un stade beaucoup plus avancé!
![Image Gtk3](/ScreenShots/MoviesDB-Gotk3.jpg)
