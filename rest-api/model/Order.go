package model

import "time"

type Order struct {
	ID            string    `json:"order_id"`
	Customer_name string    `json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	Items_list    []Item    `json:"items"`
}
