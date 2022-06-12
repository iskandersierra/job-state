migrate -source file://./db/migrations/sqlserver -database $env:SQLSERVER_CONNECTION_STRING -verbose down $args[0]
