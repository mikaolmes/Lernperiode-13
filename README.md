# Lernperiode 13

## Übersicht
In dieser Lernperiode konzentriere ich mich gezielt auf den Tech-Stack meiner zukünftigen Praktumsfirma. Ziel ist es, ein solides Grundverständnis für die Web- und Backend-Entwicklung mit modernen Technologien zu erlangen.

---

## Technologien

### 1. Next.js (Webentwicklung)
**Begründung:** Meine Praktumsfirma wird Next.js für die Frontend-Entwicklung einsetzen.
* **Projekt:** Erstellung einer Web-App zur Notizverwaltung.
* **Ressource:** [Next.js Tutorial - Full Course](https://www.youtube.com/watch?v=__mSgDEOyv8)

### 2. TypeScript (Webentwicklung & Backend)
**Begründung:** TypeScript wird in der Firma sowohl im Frontend als auch im Backend genutzt.
* **Projekt:** Entwicklung einer funktionalen To-Do Liste.
* **Fokus:** Typisierung, Interfaces und die Vorteile gegenüber Standard-JavaScript verstehen.
* **Ressourcen:** * [TypeScript Grundlagen](https://www.youtube.com/watch?v=H91aqUHn8sE)
    * [Praxisprojekt: To-Do Liste](https://www.youtube.com/watch?v=jBmrduvKl5w)

### 3. Go Programming (Backend)
**Begründung:** Go ist Neuland für mich, wird aber von der Firma für performante Backend-Systeme gefordert. Daher werde ich hier die meiste Zeit investieren.
* **Ziel:** Syntax-Grundlagen verstehen und einfache Logik implementieren.
* **Ressourcen:** * [Go Beginner Tutorial](https://www.youtube.com/watch?v=446E-r0rXHI)
    * [Vertiefungsprojekt](https://www.youtube.com/watch?v=d_L64KT3SFM)

---

## Weiteres Vorgehen
1. **Woche 1-2:** Fokus auf Next.js & TypeScript (Grundlagen des Web-Frontends).
2. **Woche 3-4:** Intensiver Einstieg in Go (Syntax & erste kleine Skripte).
3. **Abschluss:** Tutorials für jede Technologie schreiben


## 09.01.2026

Heute habe ich in Next.js die Grundlagen repetiert und danach einige Anpassungen gemacht um zu schauen was ich alles noch präsent habe, nach den Ferien usw. Hat sich herausgestellt, das eines noch ziemlich da war, danach habe ich auch schon angefangen das Tutorial zu Programmieren.

## 16.01.2026
- [X] Pocketbase Datenbank aufsetzten (verlangt das Tutorial, was sehr passend ist, da es meine Firma ebenfalls verlangt)
- [X] Ordnerstruktur wie im Tutorial fertigstellen
- [ ] Projekt nachbauen wie auf Tutorial
- [ ] Zusatz zu Tutorial Projekt: mit Tailwind CSS verbinden und schönes Design machen.

Heute habe ich zuerst das Tutorial für Next.js gemacht, doch mit einem Gespräch mit meinem Lehrer (Herr Colic), habe ich gemerkt, das ich nicht drei Sprachen machen muss, sondern nur eine. Deshalb habe ich mir Go vogenommen. Ich habe ein wenig herumprobiert und die Datei ebenfalls hochgeladen (HelloWorld). Danach habe ich eine Rest-API mithilfe eines Tutorials gecoded. Ich habe es sogar geschafft es fertig zu machen. Ich kann jetzt mit verschiedenen Endpunkten Daten abrufen, bis auf die Löschung von Daten, das hat das Tutorial nicht vorgestellt. Den Code vom Tutorial habe ich hochgeladen, Sie müssen den Code einfach mit diesem Befehl im Terminal starten: ```go run main.go```

## 23.01.2026
- [X] Löschung von Daten in Code hinzufügen (ohne Hilfe)
- [X] Herausfinden wie ich Fotos speichern kann mit Go
- [X] Herausfinden wie ich eine Datenbank (z.B. mit pocketbase) mit Go verbinden kann und dort Fotos speichern kann
- [X] Fotodatenbank Projekt anfangen --> Datenbank aufsetzten

Heute habe ich zuerst nochmal den Code vom Tutorial angeschaut und habe die Delete Funktion hinzugefügt. Am Anfang hatte ich ein paar Probleme, weil ich vergessen hatte ganz am Schluss den Endpoint in die Main Funktion hinzuzufügen. Danach ging es zwar immer noch nicht, aber das war, weil ich das Programm nochmal neu starten musste, danach ging es. Hier den Code der neu dazugefügt wurde.

```
func deleteTodo(context *gin.Context) {
    id := context.Param("id")

    for i, t := range todos {
        if t.ID == id {
            todos = append(todos[:i], todos[i+1:]...)

            context.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted"})
            return
        }
    }

    context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
```

Für die Fotodatenbank wpllte ich PocketBase verwenden, da ich das schon kannte. Ich wollte testen, ob ich es verwenden könnte. Und da PocketBase in Go geschrieben ist, konnte ich es tatsächlich sehr einfach importieren und verwenden. Anschließend habe ich die Datenbank eingerichtet und die Collection 'photos' inklusive 'image'-Feld erstellt."

## 30.01.2026

- [X] html File erstellen, eine einfache Benutzeroberfläche um Bilder hochladen zu können.

Heute habe ich eine HTML Datei erstellt, die es vereinfacht, Bilder hochzuladen, und man kann auch alle Bilder anschauen, leider aber nur den Namen und die ID. Danach habe ich das Tutorial dazu geschrieben.


