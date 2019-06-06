
package main


// Use the right libraries for the implementation
import (

	"database/sql"
	_"github.com/go-sql-driver/mysql"
	//"time"
	"math/rand"
	"fmt"
	"strings"
	//"log"

)


type Holland struct {

	realistic int
	investigative int
	artistic int
	social int
	enterprising int
	conventional int
	grid_point string
	Id int
}

func main() {

	// Create a Holland object
	H := Holland{}

	holland_code :=  []int{ 0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 10 }

	H.assign_Holland(holland_code)
}


/*
   This function is used to establish the database connection
*/
func db_connect() ( db *sql.DB ) {

	db_driver := "mysql"
	db_user := "root"
	db_pass := "password"
	db_name := "DBname"

	db, err := sql.Open( db_driver, db_user + ":" + db_pass + "@/" + db_name )

	if err != nil {

		panic( err.Error() );
	}

	return db

}

/*
   This function is used to shuffle a sequence of numbers 
   It is to make sure that different sequence of holland 
   code can be in the database table
*/

func ( H *Holland ) get_combination( combination []int ) []int {

	combination_values := make( []int, len(combination) )

	shuffle_codes:

		for i := 0; i < len(combination); i++ {

			combination_values[ i ] = combination[ rand.Intn( len(combination) ) ]
		}

	shuffle_code := combination_values[0:6]
	sum := 0

	for i := 0; i < 6; i++ {

		sum += shuffle_code[ i ]
	}

	if sum > 100 || sum < 100 {

		goto shuffle_codes
	}

	return shuffle_code

}


/*
   This function is used to update the database for the random
   sequence of numbers
  
*/
func ( H *Holland ) Update( realistic_, investigative_, artistic_, social_, enterprising_, conventional_ int , grid_point string, id int ) {

	db := db_connect()

	// sql statement to accordingly update the database table with the random sequence of holland code for 
	// each specific record in the table
	statement, err := db.Prepare( "UPDATE uww_uwp SET realistic=?, investigative=?, artistic=?, social=?, enterprising=?, conventional=?, grid_points=?, WHERE id=? AND subject !='COMPSCI' AND subject !='MAGD' ")

	// if an error happen 
	// display a panic message
	if err != nil {

		panic(err.Error)

	}

	// Set a maximum of overhead possible
	db.SetMaxIdleConns(10000)

	_, err = statement.Exec( realistic_, investigative_, artistic_, social_, enterprising_, conventional_, grid_point, id )

	defer db.Close()

	if err != nil {

		panic(err)
	}

}

/*
  This function is used to the right set 
  of sequences by updating  the database
  accordingly
*/
func ( H *Holland ) assign_Holland( holland_code_1 []int ) {

	size_courses := 5336

	for i := 0; i < size_courses; i++ {

		holland_code_2 := H.get_combination(holland_code_1)

		fmt.Println( H.grid_point )

		// assign the holland codes to be
		// added to the database table
		H.realistic = holland_code_2[ 0 ]
		H.investigative = holland_code_2[ 1 ]
		H.artistic = holland_code_2[ 2 ];
		H.social = holland_code_2[ 3 ]
		H.enterprising = holland_code_2[ 4 ]
		H.conventional = holland_code_2[ 5 ]

		H.Id = i

		H.grid_point = strings.Trim( strings.Join( strings.Fields( fmt.Sprint( holland_code_2 )), "-"), "[]" )
		fmt.Println( H.Id )
		H.Update( H.realistic, H.investigative, H.artistic, H.social, H.enterprising, H.conventional,  H.grid_point, H.Id )

	}
}






















