package db


type ObjectID struct{
	ID string
}

func (oid *ObjectID) SetOID(id string){
	oid.ID = id
}