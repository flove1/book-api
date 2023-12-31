package handler

import (
	_ "one-lab-final/docs"
	"one-lab-final/internal/entity"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	v1 := router.Group("/api/v1")

	v1.GET("/healthcheck", h.healthcheck)

	userV1 := v1.Group("/users")
	bookV1 := v1.Group("/books")
	reviewV1 := v1.Group("/reviews")
	modV1 := v1.Group("/mod")

	userV1.GET("/:username", h.getUserByUsername)
	userV1.GET("/suspensions/:id", h.checkSuspension)
	userV1.POST("/register", h.createUser)
	userV1.POST("/login", h.login)
	userV1.PATCH("/update", h.requireAuthenticatedUser(), h.updateUser)
	userV1.DELETE("/delete", h.requireAuthenticatedUser(), h.deleteUser)

	bookV1.GET("", h.getBooks)
	bookV1.GET("/:id", h.getBookByID)
	bookV1.GET("/:id/reviews", h.getReviewsByBookID)

	bookV1.POST("/new", h.requireRole(entity.MODERATOR), h.createBook)
	bookV1.DELETE("/delete/:id", h.requireRole(entity.MODERATOR), h.deleteBook)
	bookV1.PATCH("/update/:id", h.requireRole(entity.MODERATOR), h.updateBook)

	reviewV1.POST("/new", h.requireAuthenticatedUser(), h.createReview)
	reviewV1.PATCH("/update/:id", h.requireAuthenticatedUser(), h.updateReview)
	reviewV1.DELETE("/delete/:id", h.requireAuthenticatedUser(), h.deleteReview)

	modV1.POST("/suspensions/new", h.requireRole(entity.MODERATOR), h.suspendUser)
	modV1.PATCH("/suspensions/update/:id", h.requireRole(entity.MODERATOR), h.updateSuspension)

	modV1.PATCH("/roles/:id", h.requireRole(entity.ADMIN), h.grantRoleToUser)

	return router
}
