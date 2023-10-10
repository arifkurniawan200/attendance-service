package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"template/internal/usecase"
)

type handler struct {
	User        usecase.UserUcase
	Transaction usecase.TransactionUcase
}

func Run(u usecase.UserUcase, t usecase.TransactionUcase) {
	e := gin.Default()

	h := handler{
		User:        u,
		Transaction: t,
	}

	e.POST("/register", h.RegisterUser)
	e.POST("/login", h.LoginUser)

	customer := e.Group("/member")
	{
		customer.Use(JWTMiddleware("secret")) // still default,can change anytime (i suggest i should placed in  .env)
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := e.Run(":8080"); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	wg.Wait()
}
