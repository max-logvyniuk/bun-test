package shared

import "fmt"


// Data represent data Model.
type Data struct {
	ID      int64 `bun:",pk,autoincrement"`
	Message string
}

func (d Data) String() string {
	return fmt.Sprintf("Data<%d %s>", d.ID, d.Message)
}

// DataCreate is input to create Data object.
type DataCreate struct {
	Message string `bun:",string"`
}

