# Aufgaben

1. **Erstellen Sie einen Fork von diesem Repository!**
2. **Klonen Sie Ihren Fork, nicht das Original-Repository!**
3. **Reichen Sie Ihre Lösungen per Pull Request ein!**

Die folgenden Aufgaben können Sie direkt in die angegebene Datei lösen. Beachten
Sie hierzu die `// TODO: `-Kommentare im Code und die folgenden Instruktionen:

## 1) Note berechnen

Eien Note von 1 (schlechteste) bis 6 (beste) kann linear anhand der erreichten
und maximalen Punktezahl berechnet werden:

    Note = (erreichte Punktzahl / maximale Punktzahl) * 5 + 1

Schreiben Sie eine Funktion namens `computeGrade`, welche die beiden Parameter
`gotPoints` (erreichte Punktzahl) und `maxPoints` (maximale Punktzahl) erwartet
und eine Note im Bereich `[1.0;6.0]` zurückgibt.

Rufen Sie die Funktion dreimal auf und schreiben Sie das Ergebnis als Kommentar
hinter den jeweiligen Funktionsaufruf, z.B.:

    gotPoints(17.5, 28.0) // 4.125

### Zusatzaufgabe

Erweitern Sie die Funktion, sodass sie auf sinnlose Argumente reagiert (z.B.
negative Punktzahlen, mehr Punkte erreicht als möglich) und einen entsprechenden
`error` zurückgibt. Passen Sie Ihre Funktionsaufrufe entsprechend an.

## 2) Hypotenuse berechnen

Der Satz von Pythagoras ist definiert als:

    a² + b² = c²

Stellt man die Formel auf `c` um, kann man mit der Formel die Länge der
Hypotenuse bei gegebenen Kathetenlängen berechnen:

    c = Quadratwurzel von (a² + b²)

Schreiben Sie eine Funktion namens `computeHypotenuse`, welche die beiden
Kathetenlängen `a` und `b` als Parameter erwartet und die Länge der Hypotenuse
zurückgibt.

Tipp: Verwenden Sie die Funktion [`math.Pow`](https://pkg.go.dev/math#Pow) und
[`math.Sqrt`](https://pkg.go.dev/math#Sqrt) für die Berechnung.

Testen Sie die Funktion mit verschiedenen Kathetenlängen und dokumentieren Sie
die Aufrufe wie bei Aufgabe 1 mit Kommentaren.

### Zusatzaufgabe

Die beiden Kathetenlängen `a` und `b` können zu einem neuen Typ namens
`ShortSides` gruppiert werden:

```go
type ShortSides struct {
    a float64
    b float64
}
```

Schreiben Sie eine _Methode_ `Hypotenuse`, welche auf `ShortSides` operiert und
keine weiteren Parameter erwartet.

Testen Sie die Methode mit den gleichen Kathetenpaaren wie `computeHypotenuse`
und notieren Sie sich die Ergebnisse als Kommentare hinter den Methodenaufrufen.

## 3) Quadratische Gleichungen lösen

Quadratische Gleichungen der Form `ax² + bx + c = 0` lassen sich mithilfe der
sogenannten _Mitternachtsformel_ lösen:

    x = -b ± Quadratwurzel von (b² - 4ac) geteilt durch 2a

Den Teil unter der Wurzel bezeichnet man als _Diskriminante_:

    b² - 4ac

Je nach Wert der Diskriminante `D` gibt es folgende Lösungen:

- `D > 0`: zwei Lösungen
- `D = 0`: eine Lösung
- `D < 0`: keine Lösung

Schreiben Sie eine Funktion namens `computeQuadraticFormula`, welche die
Parameter `a`, `b` und `c` erwartet, und ein Slice mit den berechneten Lösungen
zurückgibt.

Tipp: Verwenden Sie die Funktion [`math.Pow`](https://pkg.go.dev/math#Pow) und
[`math.Sqrt`](https://pkg.go.dev/math#Sqrt) für die Berechnung.

Rufen Sie die Funktion dreimal auf, sodass je ein Ergebnis mit zwei Lösungen,
mit einer Lösung und ohne Lösung entsteht. Dokumentieren Sie die Ergebnisse
wiederum als Kommentare.

### Zusatzaufgabe

Lagern Sie die Berechnung der Diskriminanten in eine eigene Funktion
namens `computeDiscriminant` aus. Testen Sie die Funktion separat mit den
Testdaten von vorher und dokumentieren Sie die Ergebnisse als Kommentar.

Schreiben Sie anschliessend die Funktion `computeQuadraticFormula` so um, dass
sie zur Berechnung die Funktion `computeDiscriminant` verwendet.

## 4) Temperaturumrechnungen durchführen
