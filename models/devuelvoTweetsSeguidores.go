package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omotempty"`
	UsuarioID         string             `bson:"usuarioid" json:"userId,omotempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelationId,omotempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omotempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omotempty"`
		ID      string    `bson:"_id" json:"_id,omotempty"`
	}
}
