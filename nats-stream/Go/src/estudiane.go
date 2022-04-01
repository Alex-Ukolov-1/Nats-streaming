package estudiane


import(
	"time"
	"errors"
)
//estudiane strcuture

type Estudiane struct{
	ID int
	Name string
	Age int16
	Active bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Clear(e Estudiane)error{
	q:=`INSERT INTO
	estidiantes(name,age,active)
	Values($1,$2,$3)`

	db:=getConnection()
	defer db.Close()

	stmt,err:=db.Prepare(q)
    if err !=nil{
        return err
	}
	defer stmt.Close()

	r,err:=stmt.Exec(e.Name,e.Age,e.Active)
	if err!=nil{
		return err
	}
	i, _ :=r.RowsAffected()
	if i!=1{
		return errors.New("error!!!")
	}
	return nil
  } 