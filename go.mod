module github.com/gotzmann/grpc

go 1.16

require (
	github.com/ericlagergren/decimal v0.0.0-20181231230500-73749d4874d5
	// too modern verstion! github.com/ericlagergren/decimal v0.0.0-20210307182354-5f8425a47c58 // indirect
	// it generates the bug => pq: encode: unknown type types.NullDecimal
	github.com/friendsofgo/errors v0.9.2
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/kat-co/vala v0.0.0-20170210184112-42e1d8b61f12
	github.com/lib/pq v1.10.2
	github.com/spf13/viper v1.8.0
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/randomize v0.0.1
	github.com/volatiletech/sqlboiler/v4 v4.6.0
	github.com/volatiletech/strmangle v0.0.1
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)
