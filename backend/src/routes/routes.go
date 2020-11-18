package routes

import (
	"github.com/fullstacktf/fitness-backend/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter Function to initialize and configure gin-gonic
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/", controller.GetWelcome)

		// Users

		v1.POST("/users/", controller.CreateUser)

		v1.GET("/user/:id", controller.GetUser)

		v1.GET("/users/", controller.GetUsers)

		v1.PUT("/users/:id", controller.UpdateUser)

		v1.DELETE("/users/:id", controller.DeleteUser)

		// Paswords. Only accesible to admins

		v1.PUT("/password/:userId", controller.UpdateUserPass)

		// User stats

		v1.PUT("/userStat/:userId", controller.UpdateUserStat)

		// User weight history

		v1.POST("/userWeight/:userId", controller.CreateUserWeightHistory)

		v1.PUT("/userWeight/:userId", controller.UpdateUserWeightHistory)

		// Assigned routines. Used to asign custom routines to users by a trainer

		v1.POST("/assignedRoutine/", controller.CreateAssignedRoutine)

		v1.GET("/assignedRoutine/:id", controller.GetAssignedRoutine)

		v1.GET("/assignedRoutines/", controller.GetAssignedRoutines)

		v1.PUT("/assignedRoutine/:id", controller.UpdateAssignedRoutine)

		v1.DELETE("/assignedRoutine/:id", controller.DeleteAssignedRoutine)

		// History. Update available to the trainer only

		v1.POST("/history/", controller.CreateHistory)

		v1.PUT("/history/:id", controller.UpdateHistory)

		// Routine specific exercises. Used to store specific exercises defined within a custom routine.

		v1.POST("/routineSpecificExercises/", controller.CreateRoutineSpecificExercise)

		v1.GET("/routineSpecificExercises/:id", controller.GetRoutineSpecificExercise)

		v1.PUT("/routineSpecificExercises/:id", controller.UpdateRoutineSpecificExercise)

		v1.DELETE("/routineSpecificExercises/:id", controller.DeleteRoutineSpecificExercise)

		//Permissions

		v1.POST("/permission/", controller.CreatePermission)

		v1.GET("/permission/:id", controller.GetPermission)

		v1.GET("/permission/", controller.GetPermissions)

		v1.PUT("/permission/:id", controller.UpdatePermission)

		v1.DELETE("/permission/:id", controller.DeletePermission)

		//Roles

		v1.POST("/role/", controller.CreateRole)

		v1.GET("/role/:id", controller.GetRole)

		v1.GET("/role/", controller.GetRoles)

		v1.PUT("role/:id", controller.UpdateRole)

		v1.DELETE("/role/:id", controller.DeleteRole)

		//BaseExercise

		v1.POST("/baseExercise/", controller.CreateBaseExercise)

		v1.GET("/baseExercise/:id", controller.GetBaseExercise)

		v1.GET("/baseExercise/", controller.GetBaseExercises)

		v1.PUT("baseExercise/:id", controller.UpdateBaseExercise)

		v1.DELETE("/baseExercise/:id", controller.DeleteBaseExercise)

		//BaseRoutine

		v1.POST("/baseRoutine/", controller.CreateBaseRoutine)

		v1.GET("/baseRoutine/:id", controller.GetBaseRoutine)

		v1.GET("/baseRoutine/", controller.GetBaseRoutines)

		v1.PUT("baseRoutine/:id", controller.UpdateBaseRoutine)

		v1.DELETE("/baseRoutine/:id", controller.DeleteBaseRoutine)

		//RoutineCategory

		v1.POST("/routineCategory/", controller.CreateRoutineCategory)

		v1.GET("/routineCategory/:id", controller.GetRoutineCategory)

		v1.GET("/routineCategory/", controller.GetRoutineCategories)

		v1.PUT("routineCategory/:id", controller.UpdateRoutineCategory)

		v1.DELETE("/routineCategory/:id", controller.DeleteRoutineCategory)

		//Muscles

		v1.POST("/muscle/", controller.CreateMuscle)

		v1.GET("/muscle/:id", controller.GetMuscle)

		v1.GET("/muscle/", controller.GetMuscles)

		v1.PUT("muscle/:id", controller.UpdateMuscle)

		v1.DELETE("/muscle/:id", controller.DeleteMuscle)
	}

	return r
}