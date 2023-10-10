package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"template/internal/usecase"
)

type handler struct {
	User      usecase.UserUcase
	Gathering usecase.GatheringUcase
}

func Run(u usecase.UserUcase, t usecase.GatheringUcase) {
	e := gin.Default()

	h := handler{
		User:      u,
		Gathering: t,
	}

	e.POST("/register", h.RegisterUser)
	e.POST("/login", h.LoginUser)

	member := e.Group("/member")
	{
		member.Use(JWTMiddleware("secret")) // still default,can change anytime (i suggest i should placed in  .env)
		member.POST("/gathering", h.CreateGathering)
		member.POST("/gathering/:id/send", h.SendInvitation)
		member.PUT("/gathering/:id/reject", h.RejectInvitation)
		member.PUT("/gathering/:id/approve", h.ApproveInvitation)
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
