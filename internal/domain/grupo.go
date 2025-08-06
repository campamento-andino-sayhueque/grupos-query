package domain

// Grupo represents the projection of a group in the read database.
type Grupo struct {
	ID                 string   `firestore:"id"`
	Nombre             string   `firestore:"nombre"`
	Dirigentes         []string `firestore:"dirigentes"`
	CantidadAcampantes int      `firestore:"cantidadAcampantes"`
}
