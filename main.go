package main

import (
	"TFP/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println("Starting Application")
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.SellerRoutes(router)
	routes.AdminRoutes(router)
	//router.Use(middleware.Authentication())

	routes.SecuredUserRoutes(router)
	routes.SecuredSellerRoutes(router)
	routes.SecuredAdminRoutes(router)
	routes.CategoryRoutes(router)
	routes.ProductsRoutes(router)
	routes.AddToCartRoutes(router)
	routes.WishlistRoutes(router)
	routes.BuyRoutes(router)

	router.Run(":" + port)

	//http.HandleFunc("/socket", routes.SocketRoute(test))
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

//func test(w http.ResponseWriter, r *http.Request) {
//	ws, err := websocket.Upgrade(w, r)
//	if err != nil {
//		fmt.Println(err)
//	}
//	go websocket.Writer(ws)
//}
