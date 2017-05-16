package abstract

import "gopkg.in/oauth2.v3"

const(
    // DiNameConfig is the configuration dependency for the di contaienr
    DiNameConfig = "config"
)

type (

    ClientStore interface {

    }

    TokenStore interface {
        // create and store the new token information
        Create(info oauth2.TokenInfo) error

        // delete the authorization code
        RemoveByCode(code string) error

        // use the access token to delete the token information
        RemoveByAccess(access string) error

        // use the refresh token to delete the token information
        RemoveByRefresh(refresh string) error

        // use the authorization code for token information data
        GetByCode(code string) (oauth2.TokenInfo, error)

        // use the access token for token information data
        GetByAccess(access string) (oauth2.TokenInfo, error)

        // use the refresh token for token information data
        GetByRefresh(refresh string) (oauth2.TokenInfo, error)
    }

    UserStore interface {

    }
)
