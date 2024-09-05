package objectidpkg

import "go.mongodb.org/mongo-driver/bson/primitive"

func NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func ObjectIDFromHex(hex string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return primitive.NilObjectID
	}
	return objID
}

func Cast2ObjectID(data interface{}) primitive.ObjectID {
	if objID, ok := data.(primitive.ObjectID); ok {
		return objID
	}
	return primitive.NilObjectID
}
