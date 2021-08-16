# Blog

An API for a blog written in GO with ElasticSearch.

## Installation and Usage
1. Clone the repository
```bash
git clone https://github.com/afifialaa/blog.git
```
2. Change directory
```bash
cd blog
```
3. Install dependencies
```bash
go get
```

### Add config functions
```go
func SetEnv() {
	fmt.Println("Setting env")
	os.Setenv("CLOUD_MONGO", "<YOUR_CONNECTION_STRING>")
	os.Setenv("PORT", "8000")
}

func InitApp() {
	fmt.Println("Initializing app")
	ES, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic("Client failed")
	}
	fmt.Println(ES.Ping())
}
```


Server will be running on port 8000

## Articles Routes
    
| Method  | Route | description |
| ------------- |:-------------:| ------------|
| GET      | /blog    | Fetches all articles |
| POST      | /blog     | Creates new article |
| DELETE     | /blog     | Deletes an article |
| PUT     | /blog     | Updates an article |

## Comments Routes

| Method  | Route | Description |
| ------------- |:-------------:| ----------|
| GET      | /comment    | Fetches comments of an article
| POST      | /comment     | Creates new comment |
| DELETE     | /comment     | Deletes a comment |
| PUT     | /comment     | Update a comment |

## Pending
* User authentication
