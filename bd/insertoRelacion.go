package bd

import (
	"context"
	"time"

	"github.com/devtimx/backTwit/models"
)

/*InsertoRelacion inserta la relacion entre usuarios en la DB*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("backTwit")
	col := db.Collection("relations")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
