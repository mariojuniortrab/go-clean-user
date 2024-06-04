package default_mysql_repository

const rawSelectQuery = "SELECT ##fields## FROM ##table## "
const rawDeleteQuery = "DELETE FROM ##table## WHERE id = ?"
const rawInsertQuery = "INSERT INTO ##table## (##fields##) VALUES (##values##)"
const rawUpdateQuery = "UPDATE ##table## SET ##set## WHERE id = ?"
