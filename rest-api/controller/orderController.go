package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"rest-api/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetDBInstance(database *sql.DB) {
	db = database
}

func CreateOrder(ctx *gin.Context) {
	// tampung data
	var orderData model.Order

	// baca data dari json, param -> tampung ke orderData
	err := ctx.ShouldBindJSON(&orderData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		fmt.Println("error should bind json")
		return // force exit
	}

	// insert order ID
	sqlStatementOrder := `INSERT INTO orders(customer_name, ordered_at, created_at, updated_at) VALUES(?, ?, ?, ?)`
	result, err := db.Exec(sqlStatementOrder, orderData.Customer_name, orderData.Ordered_at, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}

	// get orderID
	orderID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	orderData.ID = strconv.Itoa(int(orderID)) // convert int to string

	// insert items ke table item
	sqlStatementItem := `INSERT INTO items(name, description, quantity, order_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)`
	for index, value := range orderData.Items_list {

		res, err := db.Exec(sqlStatementItem, value.Name, value.Description, value.Quantity, orderData.ID, time.Now(), time.Now())

		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		orderData.Items_list[index].ID = strconv.Itoa(int(lastInsertID)) // convert int to string
		orderData.Items_list[index].Order_id = orderData.ID              // assign order_id to the item list
		orderData.Items_list[index].Created_at = time.Now()
		orderData.Items_list[index].Updated_at = time.Now()

		fmt.Printf("\n[Success] Added new item: %v", orderData.Items_list[index])
	}

	// add created and updated time for order table
	orderData.Created_at = time.Now()
	orderData.Updated_at = time.Now()

	ctx.JSON(http.StatusOK, gin.H{
		"data":    orderData,
		"message": "Added New Order to Database.",
	})
}

func GetAllOrders(ctx *gin.Context) {

	// ambil data order table
	sqlRetrieveOrder := `SELECT * FROM orders`
	rows, err := db.Query(sqlRetrieveOrder)
	if err != nil {
		panic(err)
	}

	var ordersList = []model.Order{}
	for rows.Next() {
		var order = model.Order{}

		// coba scan/ambil value order table tiap column nya
		err = rows.Scan(&order.ID, &order.Customer_name, &order.Ordered_at, &order.Created_at, &order.Updated_at)
		if err != nil {
			panic(err)
		}

		// ambil data item table dgn order_id current rows
		sqlRetrieveItem := `SELECT * FROM items WHERE order_id = ?`
		rowsItem, err := db.Query(sqlRetrieveItem, order.ID)
		if err != nil {
			panic(err)
		}

		for rowsItem.Next() {
			var item = model.Item{}

			// coba scan value item table tiap column nya
			err = rowsItem.Scan(&item.ID, &item.Name, &item.Description, &item.Quantity, &item.Order_id, &item.Created_at, &item.Updated_at)
			if err != nil {
				panic(err)
			}

			// masukkin item ke item list
			order.Items_list = append(order.Items_list, item)
		}
		// masukkin order ke order list
		ordersList = append(ordersList, order)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    ordersList,
		"message": "Successfully retrieve data.",
	})
}

func GetOrderDetailByID(ctx *gin.Context) {
	orderID := ctx.Param("id")

	sqlRetrieveOrder := `SELECT * FROM orders WHERE id = ?`
	rows, err := db.Query(sqlRetrieveOrder, orderID)
	if err != nil {
		panic(err)
	}

	var order model.Order
	if rows.Next() {
		err = rows.Scan(&order.ID, &order.Customer_name, &order.Ordered_at, &order.Created_at, &order.Updated_at)
		if err != nil {
			panic(err)
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("order with id %v not found", orderID),
		})
		return
	}

	sqlRetrieveItem := `SELECT * FROM items WHERE order_id = ?`
	rowsItem, err := db.Query(sqlRetrieveItem, order.ID)
	if err != nil {
		panic(err)
	}

	for rowsItem.Next() {
		var item = model.Item{}

		err = rowsItem.Scan(&item.ID, &item.Name, &item.Description, &item.Quantity, &item.Order_id, &item.Created_at, &item.Updated_at)
		if err != nil {
			panic(err)
		}

		order.Items_list = append(order.Items_list, item)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": "Successfully retrieve data.",
	})
}

func UpdateOrderByID(ctx *gin.Context) {
	orderID := ctx.Param("id")

	var updatedOrder model.Order
	updatedOrder.ID = orderID

	// baca data dari json, return error
	err := ctx.ShouldBindJSON(&updatedOrder)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// update order by ID
	sqlStatementOrder := `UPDATE orders SET ordered_at = ?, customer_name = ? WHERE id = ?`
	result, err := db.Exec(sqlStatementOrder, updatedOrder.Ordered_at, updatedOrder.Customer_name, orderID)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	if count == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("order with id %v not found", orderID),
		})
		return
	}

	sqlRetrieveItem := `SELECT * FROM items WHERE order_id = ?`
	rowsItem, err := db.Query(sqlRetrieveItem, orderID)
	if err != nil {
		panic(err)
	}

	var indexItem = []string{}
	for rowsItem.Next() {
		var item = model.Item{}

		err = rowsItem.Scan(&item.ID, &item.Name, &item.Description, &item.Quantity, &item.Order_id, &item.Created_at, &item.Updated_at)
		if err != nil {
			panic(err)
		}

		indexItem = append(indexItem, item.ID)
	}

	sqlStatementItem := `UPDATE items SET name = ?, description = ?, quantity = ? WHERE id = ?`

	for index, value := range updatedOrder.Items_list {
		res, err := db.Exec(sqlStatementItem, value.Name, value.Description, value.Quantity, indexItem[index])

		if err != nil {
			panic(err)
		}

		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}

		updatedOrder.Items_list[index].ID = indexItem[index]
		updatedOrder.Items_list[index].Order_id = orderID

		fmt.Println("Updated data amount: ", count)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    updatedOrder,
		"message": "Successfully update the data.",
	})
}

func DeleteOrderByID(ctx *gin.Context) {
	orderID := ctx.Param("id")

	sqlDeleteStatementItem := `DELETE FROM items WHERE order_id = ?`
	resItem, err := db.Exec(sqlDeleteStatementItem, orderID)
	if err != nil {
		panic(err)
	}

	countItem, err := resItem.RowsAffected()
	if err != nil {
		panic(err)
	}

	sqlDeleteStatementOrder := `DELETE FROM orders WHERE id = ?`
	resOrder, err := db.Exec(sqlDeleteStatementOrder, orderID)
	if err != nil {
		panic(err)
	}

	countOrder, err := resOrder.RowsAffected()
	if err != nil {
		panic(err)
	}

	if countOrder == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("order with id %v not found", orderID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": fmt.Sprintf("Successfully delete order with id %v, and deleted %d items", orderID, countItem),
	})
}
