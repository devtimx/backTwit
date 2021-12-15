package bd

import (
	"context"
	"time"

	"github.com/devtimx/backTwit/models"
)

/*BorroRelacion elimina la relacion entre usuarios de la DB*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("backTwit")
	col := db.Collection("relations")
	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
