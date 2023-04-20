# MyGram RESTFull API 

## Description 
This is a project of RESTFull API using Go language and detailed tech stack below 
| Tech          | Used as       |
| ------------- | ------------- |
| Gin           | web framework in Go     |
| Postgressql   | Dataabase     |
| Swagger       | API Spec Docs |
| ORM Gorm      | ORM Library in Go |
| Json Web Token (JWT)           | encrypt auth  |
| Docker        | Container Package to run[soon] |
| Terraform     | IaaC to create resource in cloud[soon] |
| Cloud Run     | Deploy the api in production[soon] |

## Run this App 
1. clone in your local ``````
2. Run Using docker | 

## API List 

http://localhost:8081

    User Routes :
	    POST("/users/register") register user
	    POST("/users/login")    login
	
	Photo Routes :
		GET("/photos")              getting all photos
		GET("/photos/:photoId")     getting photo by Id
		POST("/photos")             Create photo
		PUT("/photos/:photoId")     uptade photo (using authorization)
		DELETE("/photos/:photoId")  delete photo (using authorization)

	Comment Router :
		GET("/comments")                getting all comments
		GET("/comments")                getting comment by Id
		POST("/comments")               create comment
		PUT("/comments/:commentId")     uptade comment (using authorization)
		DELETE("/comments/:commentId")  delete comment (using authorization)

	Social Media Router :
		GET("/socialmedias")             getting all social medias
		GET("/socialmedias/:mediaId")    getting social media by Id
		POST("/socialmedias")            create social media
		PUT("/socialmedias/:mediaId")    update social media (using authorization)
		DELETE("/socialmedias/:mediaId") delete social media (using authorization)
