package controllers

import (
	"fmt"
	"net/http"
	"v1/models"

	"v1/config"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	product := models.Product{}
	var resp models.Product
	c.ShouldBindJSON(&product)
	var stroreIDs []int
	for _, store := range product.Stores {
		var respstore models.Store
		err := config.DB.QueryRow("insert into stores (name) values($1) returning id, name", store.Name).Scan(&respstore.ID, &respstore.Name)
		if err != nil {
			fmt.Println("error while insert into stores", err)
			return
		}
		fmt.Println(respstore.Name)
		stroreIDs = append(stroreIDs, respstore.ID)
		for _, address := range store.Addresses {
			var respAddress models.Address
			err = config.DB.QueryRow("insert into addresses (street, district, store_id) values($1,$2, $3) returning id, district, street, store_id", address.Street, address.District, store.ID).Scan(&respAddress.ID, &respAddress.District, &respAddress.Street, &respAddress.StoreID)
			if err != nil {
				fmt.Println("error while insert into addresses", err)
				return
			}
			fmt.Println(respAddress)
			respstore.Addresses = append(respstore.Addresses, respAddress)
		}

		resp.Stores = append(resp.Stores, respstore)
	}

	fmt.Println("this is product name", product.Category, product.Name, product.Type, product.Price)
	err := config.DB.QueryRow("insert into product (name,category,type, price) values($1,$2,$3,$4) returning id, name, price, category, type", product.Name, product.Category, product.Type, product.Price).Scan(&resp.ID, &resp.Name, &resp.Price, &resp.Category, &resp.Type)
	if err != nil {
		fmt.Println(err)
	}
	for _, sID := range stroreIDs {
		config.DB.Exec("insert into store_products (store_id, product_id) values($1,$2)", sID, resp.ID)
	}
	fmt.Println(resp)
	c.JSON(http.StatusAccepted, gin.H{
		"product": resp,
	})
}

func UpdateProduct(c *gin.Context) {

	product := models.Product{}
	id := c.Param("id")
	c.ShouldBindJSON(&product)

	_, err := config.DB.Exec("update product set name=$1, price=$2, type=$3, category=$4 where id=$5", product.Name, product.Price, product.Type, product.Category, id)

	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"err":   err,
			"error": "error while update product",
		})
	}

	if len(product.Stores) > 0 {
		for _, store := range product.Stores {
			_, err := config.DB.Exec("update stores set name=$1 where id=$2", store.Name, store.ID)
			if err != nil {
				fmt.Println("error while updating stores", err)
				return
			}
			if len(store.Addresses) > 0 {
				for _, address := range store.Addresses {
					_, err := config.DB.Exec("update addresses set street=$1, district=$2 where id=$3", address.Street, address.District, store.ID)
					if err != nil {
						fmt.Println("error while insert into addresses", err)
						return
					}
				}
			}
		}
	}

	c.JSON(http.StatusAccepted, gin.H{
		"product": product,
	})
}

func GetProduct(c *gin.Context) {

}
