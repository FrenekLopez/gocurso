package users

import (
	"fmt"
	"time"

	"github.com/freneklopez/gocurso/modelos"
)

func AltaUsusario() {
	u := new(modelos.User)
	u.AddUser(10, "eric", time.Now(), true)

	fmt.Println(u)
}
