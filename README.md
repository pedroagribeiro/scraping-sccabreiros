# Scraping Service

The classifications that are shown in
https://dev.sccabreiros.org/games/classification are supposed to be as accurate
as possible. In order to achieve this goal there would be two possible solutions
- either update the table at the end of each championship round or rely on a
table that an external service provides. We chose to rely on the classification
tables that are available in https://zerozero.pt.

However, https://zerozero.pt does not make its API public so simply requesting
the updated table to the API is not an option. Therefore we chose to develop a
web scraping service that gathers the table information whenever its API
endpoint is called.

Since web scraping takes a long time, generaly several seconds, the Flask
service will conserve the web scraping result for a week and returned the cached
version. When it receives a request more than a week has passed since the last
scraping the information will be updated as a result of the new scraping
process.
