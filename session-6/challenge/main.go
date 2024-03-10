package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user      = "root"
	password  = ""
	dbname    = "go_sql_challenge"
	parseTime = "true"
)

var (
	db  *sql.DB
	err error
)

type Product struct {
	ID         int
	Name       string
	Created_at time.Time
	Updated_at time.Time
}

type Variant struct {
	ID          int
	Varian_name string
	Product_id  int
	Created_at  time.Time
	Updated_at  time.Time
}

func main() {
	mysqlInfo := fmt.Sprintf("%s:%s@/%s?parseTime=%s", user, password, dbname, parseTime)

	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close() // tutup di akhir (defer)

	err = db.Ping() // test connection masih jalan atau engga
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	menu()
}

func inputName(table string) string {
	fmt.Printf("Masukkan nama %s: ", table)

	scanner := bufio.NewReader(os.Stdin)
	name, _ := scanner.ReadString('\n')
	// to remove new line
	name = strings.TrimSpace(name)

	return name
}

func inputID(table string) int {
	fmt.Printf("Masukkan ID %s: ", table)

	var id int
	fmt.Scanln(&id)

	return id
}

func menu() {
loop:
	for {
		menuOption()
		fmt.Print("Masukkan opsi: ")
		var opsiMenu int
		// fmt.Scan()
		fmt.Scanln(&opsiMenu)

		switch opsiMenu {
		case 1:
			createProduct()
		case 2:
			updateProduct()
		case 3:
			id := inputID("product")
			getProductById(id)
		case 4:
			createVariant()
		case 5:
			updateVariantById()
		case 6:
			deleteVariantById()
		case 7:
			getProductWithVariant()
		case 0:
			fmt.Println("Exit...")
			break loop
		default:
			fmt.Println("Opsi tidak valid")
		}
		time.Sleep(1 * time.Second)
	}
}

func menuOption() {
	fmt.Println("===============")
	fmt.Println("1. Create Product")
	fmt.Println("2. Update Product By ID")
	fmt.Println("3. Get Product By ID")
	fmt.Println("4. Create Variant")
	fmt.Println("5. Update Variant By ID")
	fmt.Println("6. Delete Variant By ID")
	fmt.Println("7. Get Product with Variant")
	fmt.Println("===============")
}

func createProduct() {
	name := inputName("product") // get product name from user

	var product = Product{}

	sqlStatement := `INSERT INTO products(name, created_at, updated_at) VALUES(?, ?, ?)`
	res, err := db.Exec(sqlStatement, name, time.Now(), time.Now())

	if err != nil {
		panic(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	sqlRetrieve := `SELECT * FROM products WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&product.ID, &product.Name, &product.Created_at, &product.Updated_at)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n[Success] Added new product data:\n%+v\n", product)
}

func updateProduct() {
	id := inputID("product")     // get product id from user
	name := inputName("product") // get product name from user

	sqlStatement := `UPDATE products SET name = ? WHERE id = ?`

	result, err := db.Exec(sqlStatement, name, id)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("\nUpdated data amount: ", count)
}

func getProductById(id int) {
	sqlRetrieve := `SELECT * FROM products WHERE id = ?`

	rows, err := db.Query(sqlRetrieve, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var product = Product{}
	if rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Created_at, &product.Updated_at)
		if err != nil {
			panic(err)

		}
	}

	fmt.Println("\nProduct data: ", product)
}

// variant
func createVariant() {
	name := inputName("variant") // get variant name from user
	id := inputID("product")     // get product ID from user

	sqlStatement := `INSERT INTO variants(varian_name, product_id, created_at, updated_at) VALUES(?, ?, ?, ?)`
	res, err := db.Exec(sqlStatement, name, id, time.Now(), time.Now())

	if err != nil {
		panic(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	sqlRetrieve := `SELECT * FROM variants WHERE id = ?`

	var variant = Variant{}
	err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&variant.ID, &variant.Varian_name, &variant.Product_id, &variant.Created_at, &variant.Updated_at)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New variant data: %+v\n", variant)
}

func updateVariantById() {
	name := inputName("variant") // get product name from user
	id := inputID("variant")     // get variant ID from user

	sqlStatement := `UPDATE variants SET varian_name = ? WHERE id = ?`

	result, err := db.Exec(sqlStatement, name, id)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount: ", count)
}

func deleteVariantById() {
	id := inputID("variant") // get variant ID from user

	sqlStatement := `DELETE from variants WHERE id = ?`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount: ", count)
}

func getProductWithVariant() {
	// join query
	id := inputID("product") // get product ID from user

	sqlJoinRetrieve := `SELECT products.id, products.name, variants.id, variants.varian_name
	FROM products
	INNER JOIN variants ON products.id = variants.product_id
	WHERE products.id = ?`

	rows2, err := db.Query(sqlJoinRetrieve, id)
	if err != nil {
		panic(err)
	}
	defer rows2.Close()

	type Product_Variant struct {
		Product_ID   string
		Product_name string
		Variant_ID   string
		Variant_name string
	}

	var listProductVariant []Product_Variant
	for rows2.Next() {
		var product_variant = Product_Variant{}

		err = rows2.Scan(&product_variant.Product_ID, &product_variant.Product_name, &product_variant.Variant_ID, &product_variant.Variant_name)
		if err != nil {
			panic(err)
		}

		listProductVariant = append(listProductVariant, product_variant)
	}

	fmt.Println("Product with Variant Data:")
	if listProductVariant != nil {
		for _, item := range listProductVariant {
			fmt.Println(item)
		}
	} else {
		fmt.Println("No data")
	}
}
