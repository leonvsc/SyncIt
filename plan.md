# Plan

## Algemeen

- Bestand in delen sturen door middel van een `int64` waardoor er bestanden groter dan 4gb verzonden kunnen worden. Dit word gedaan in een loop.
  - De server ontvangt een GET request van de client -> Server leest leest de metadata van het bestand en vult de header -> Server stuurt een response naar de client welke in delen worden ontvangen door de client -> client zet ontvangen data om in een bestand.
- Na een succesvolle sync moeten de bestanden dezelfde info bevatten. Hier mag geen afwijking in zitten.
- End to end test: Het testen van de volledige flow tussen server en client.

## Client

- Optie om een server URL in te voeren.
- Optie om een map of bestand te kiezen die gesync moet worden.
  - Door middel van een ingebouwde file manager?
  - Selecteren in een menu in de client? Deze zou dan de volledige map moeten lezen en de gebruiker moet dan in een menu selecteren wat die wilt syncen.
- Testen schrijven

## Server

- Server in de cloud laten draaien.
  - Server laten draaien in een docker container.
- Zodra een bestand word ontvangen moet deze direct doorgestuurd worden naar de op dat moment geconnecte clients.
- Testen schrijven
