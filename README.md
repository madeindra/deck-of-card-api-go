# Toggl Backend Unattended Programming Test

The aim of this project is to create REST API that can simulate a deck of cards.

## How To Run
1. Create `.env` by copying values from `.env.example`
2. Edit the values inside `.env` to reflect to your environment
3. Build using `make` or `go`
```
make build
```

or

```
go build -o main
```

4. Run the app (make sure you already have the required database & tables)
```
./main
```

## Database Preparation
1. Create a database in PostgreSQL
```
CREATE DATABASE deckofcards;
```

2. Create table for decks data
```
CREATE TABLE decks (
  uuid VARCHAR(255) NOT NULL,
  shuffled BOOLEAN,
  PRIMARY KEY (uuid)
);
```

3. Create table for cards data
```
CREATE TABLE cards (
	uuid VARCHAR(255) NOT NULL,
  deck_uuid VARCHAR(255) NOT NULL,
	value VARCHAR(255) NOT NULL,
	suit VARCHAR(255) NOT NULL,
	code VARCHAR(255) NOT NULL,
  PRIMARY KEY (uuid),
  FOREIGN KEY (deck_uuid) REFERENCES decks(uuid)
);
```


## APIs

### Create a new Deck
Path: `/v1/decks`

Method: `POST`

Query Parameters:
  * shuffled (optional, boolean, default to false)
  * cards (optional, comma separated string, 2-3 chars of value [2-10, A, J, Q, K] and suit [H, D, S, C], example 10C)

Path Parameters: -

Request Body: -

Request Header: -

### Open a Deck
Path: `/v1/decks/{deck_id}`

Method: `GET`

Query Parameters: -

Path Parameters: 
   * deck_id (required, uuid of the deck)

Request Body: -

Request Header: -

### Draw a Card (or some Cards)
Path: `/v1/decks/{deck_id}/draw`

Method: `GET`

Query Parameters: 
   * count (required, int, minimum of 1)

Path Parameters: 
   * deck_id (required, uuid of the deck)

Request Body: -

Request Header: -

## Open API Specification
Please refer to [openapi.yml](./docs/openapi.yml) inside `docs` directory.

## FAQ
Q: I can't run the built file

A: Give permission to execute the file by running this command
```
chmod +x ./main
```

Q: What will happen if I draw more cards than the remaining cards in the deck?

A: It will draw the remaining cards

Q: What will happen if I draw card(s) from empty deck?

A: It will return an error

## Further Development
- Unit tests
- Build dockerimage for simpler deployment
- Refactoring & cleaning