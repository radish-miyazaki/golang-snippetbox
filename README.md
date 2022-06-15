### Snippetbox

This repo is unofficial sample code in LET'S GO! 1st edition (by **[Alex Edwards](https://www.alexedwards.net/)**).

For practice purposes, it differs slightly from the book.

Book available for purchase here: https://lets-go.alexedwards.net/sample/00.00-front-matter.html

### Environment
- Go 1.18
- MySQL 8.0
- go-chi/chi v5.0.7
- go-sql-driver/mysql v1.6.0
- alexedwards/scs v2.5.0

### How to run
1. run `make build` command to build MySQL image 
2. run query in `schema.sql` & `insert_snippets.sql` to setup database
3. run `go cmd/web` command to run HTTP server
4. Visit http://localhost:4000/
