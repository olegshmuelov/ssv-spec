package qbft

//go:generate rm -f ./messages_encoding.go
//go:generate go run .../fastssz/sszgen --path messages.go --include ./types.go,../types/crypto.go,../types/operator.go --exclude-objs OperatorID