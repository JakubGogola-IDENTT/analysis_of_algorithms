# Sprawozdanie [lista 1]

**Jakub Gogola**  
236412
___

Lista 1. na laboratoria polegała na zaimplementowaniu algorytmu **wyboru lidera** oraz jego analizie w ramach zadań nr 2, 3 oraz 4.

## Zadanie 2
W niniejszym zadaniu należało zobrazować rozkład empiryczny zmiennej losowej **L** definiowanej jako liczba slotów potrzebnych do wyboru lidera. Poniżej przedstawiono histogramy dla odpowiednich eksperymentów:

![with_nodes](hist_with_nodes_n_100.png)  
_(Histogram dla scenariusza ze znaną liczbą węzłów n = 100)_

![with_upper_limit](hist_with_upper_limit_u_100_n_2.png)  
_(Histogram dla scenariusza z zadanym ogarniczniem górnym u = 100, n = 2)_

![with_upper_limit](hist_with_upper_limit_u_100_n_50.png)  
_(Histogram dla scenariusza z zadanym ogarniczniem górnym u = 100, n = 50)_

![with_upper_limit](hist_with_upper_limit_u_100_n_100.png)  
_(Histogram dla scenariusza z zadanym ogarniczniem górnym u = 100, n = 100)_

Powyższe wykresy prezentują wyniki jedynie dla wybranych wartości zmiennych **n** oraz **u**, ponieważ ciężko byłoby zmieścić więcej w tak krótkim raporcie, jednak przeprowadzone zostaly testy dla różnych wartości wspomnianych zmiennych oraz różnej liczbi powtórzeń eksprymentu (iteracji) i za każdym razem otrzymywano zbliżoy kształ wykresów do tych prezentowanych w tym dokumencie.

Na każdym z prezentowanych wykresów można zauważyć, że rozkład zmiennej losowej układa się w **rozkład geometryczny**. Rozkład ten opisuje prawdopodobieństwo zdarzenia, że pierwszy sukces (w przypadku analizowanego algorytmu za sukces przyjmuje się wybór lidera) zostanie osiągnięty w k-tej próbie (tutaj - slocie). Jest on zatem odwzorowaniem wyników uzyskanych podczas przeprowadzonych eksperymentów - w k-tym slocie został wybrany lider (tylko jeden z węzłów nadawał).

## Zadanie 3

W zadaniu tym należlo dla scenariusza ze znaną liczbą węzłów policczyć (eksperymentalnie) wariancję oraz wartość oczekiwanę. Przeprowadzono eksperymenty dla wartości _n = 1, ..., 100_ i otrzymano następujące wyniki:

| _n_ | _EX_ | _Var_ | 
| --- | --- | --- |
| 20 | 2.623800 | 4.262274 |
| 50 | 2.706200 | 4.649482 |
| 75 | 2.712600 | 4.674601 |
| 100 | 2.711800 | 4.753741 |

Ze względów praktycznych w powyższej tabeli zaprezentowano jedynie wybrane wyniki.

Wiadomo, że wartośc oczekiwana dla zmiennej losowej _L_ w rozkładzie geometrycznym _EX[L] = 1/p_ oraz wariancja _Var[L] = (1 - p) / p<sup>2</sup>_ oraz (z lematu 2.) wiadomo, że _E[L] = 1/p < e_. Stąd można oszacować, na podstawie podanych wzorów, że _Var[L] = (1 - p) / p<sup>2</sup> < e<sup>2</sup> - e_ (~ 4.67), co zgadza się z wynikami otrzymanymi w ramach eksperymentu.

Warto zauważyć, że w przypadku eksperymentu wartość oczekiwana jest średnią arytmetyczną, a wariancja dla poszczególnych zmiennych (zmienna to numer slotu, w którym wybrano lidera) to średnia arytmetyczna kwadratów odchyleń (_EX_) od ich średniej arytmetycznej.

## Zadanie 4
W zadaniu 4. należało, w sposób eksperymentalny, wyznaczyć wartość λ (ograniczenie z twierdzenia 1.). W tym celu przeprowadzono symulację dla scenariusza z zadanym ogarniczeniem górnym i zliczano, w której rundzie nastapił sukces. Zgodnie z treścią twierdzenia interesowały nas takie zdarzenia, gdzie lider został wybrany w pierwszej rundzie. W tym celu zliczano liczbę rund potrzebnych do wyboru lidera i za sukces uznawano takie zdarzenie, gdzie lider został wybrany w pierwszej rundzie. Otrzymano następujące wyniki:

| _n_ | _u_ | _λ_ |
| --- | --- | --- |
| 2 | 100 | 0.809800 |
| 50 | 100 | 0.724800 | 
| 100 | 100 | 0.628000 |

Zgodnie z treścią twierdzenia _λ_ jest w przybliżeniu równa 0.579. Twierdzenie to mówi, że prawdopodobieństwo wyboru lidera w pierwszej rundzie jest większe od tej wartości i prezentowane wyniki to potwierdzają.


