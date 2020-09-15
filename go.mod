module github.com/ory/fosite

require (
	github.com/asaskevich/govalidator v0.0.0-20180720115003-f9ffefc3facf
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/mock v1.4.4
	github.com/gorilla/mux v1.7.0
	github.com/magiconair/properties v1.8.1
	github.com/mattn/goveralls v0.0.5
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/oleiade/reflections v1.0.0
	github.com/ory/go-acc v0.2.1
	github.com/ory/go-convenience v0.1.0
	github.com/parnurzeal/gorequest v0.2.15
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	gopkg.in/square/go-jose.v2 v2.5.0
)

replace golang.org/x/oauth2 => products.bosch-si.com/stash/bisa/golang-oauth2 v0.0.0-tokex-contrib

go 1.14
